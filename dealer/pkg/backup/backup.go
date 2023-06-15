package backup

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/common/oss"
	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"

	"github.com/filecoin-project/go-commp-utils/writer"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-car"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/web3eye-io/Web3Eye/common/unixfs"
)

type backup struct {
	host host.Host
	peer *peer.AddrInfo
	mock bool
}

func (b *backup) mockOne(ctx context.Context) (cid.Cid, string, error) {
	b1 := make([]byte, 1024*1024)
	copy(b1, []byte("0123456789abcdef0123456789abcdef"))
	mockSrcPath := "/tmp/mockOneSource.data"
	mockDstPath := "/tmp/mockOneDest.data"

	_ = os.Remove(mockSrcPath)
	_ = os.Remove(mockDstPath)

	src, err := os.OpenFile(mockSrcPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return cid.Undef, "", err
	}

	for i := 0; i < 8*1024*1024/len(b1); i++ {
		_, err := src.Write(b1)
		if err != nil {
			return cid.Undef, "", err
		}
	}

	_cid, err := unixfs.CreateFilestore(ctx, mockSrcPath, mockDstPath)
	if err != nil {
		return cid.Undef, "", err
	}

	return _cid, mockDstPath, nil
}

func (b *backup) backupOne(ctx context.Context, index uint64) error {
	snapshot, err := orbit.Snapshot().GetSnapshot(ctx, index)
	if err != nil {
		return err
	}

	switch snapshot.BackupState {
	case dealerpb.BackupState_BackupStateSuccess:
		fallthrough // nolint
	case dealerpb.BackupState_BackupStateFail:
		return nil
	default:
	}

	if snapshot.SnapshotCommP == "" {
		_, _ = orbit.Snapshot().UpdateSnapshot(ctx, index, dealerpb.BackupState_BackupStateFail)
		return fmt.Errorf("invalid snapshot commP")
	}
	if snapshot.SnapshotRoot == "" {
		_, _ = orbit.Snapshot().UpdateSnapshot(ctx, index, dealerpb.BackupState_BackupStateFail)
		return fmt.Errorf("invalid snapshot root")
	}
	if snapshot.SnapshotURI == "" {
		_, _ = orbit.Snapshot().UpdateSnapshot(ctx, index, dealerpb.BackupState_BackupStateFail)
		return fmt.Errorf("invalid snapshot uri")
	}

	// TODO: backup items to IPFS

	var rdr interface{}
	w := &writer.Writer{}

	if b.mock {
		_cid, carPath, err := b.mockOne(ctx)
		if err != nil {
			return err
		}
		snapshot.SnapshotRoot = _cid.String()
		rdr, err = os.Open(carPath)
		if err != nil {
			return err
		}
		if _, err := rdr.(*os.File).Seek(0, io.SeekStart); err != nil {
			return err
		}
	} else {
		buf, err := oss.GetObject(ctx, snapshot.SnapshotURI)
		if err != nil {
			return err
		}
		rdr = bytes.NewReader(buf)
	}

	r := bufio.NewReader(rdr.(io.Reader))
	if _, err := car.ReadHeader(r); err != nil {
		return err
	}

	if _, err := io.CopyBuffer(w, rdr.(io.Reader), make([]byte, writer.CommPBuf)); err != nil {
		return err
	}

	start := time.Now()
	commP, err := w.Sum()
	if err != nil {
		return err
	}

	if b.mock {
		snapshot.SnapshotCommP = commP.PieceCID.String()
	}

	if commP.PieceCID.String() != snapshot.SnapshotCommP {
		_, _ = orbit.Snapshot().UpdateSnapshot(ctx, index, dealerpb.BackupState_BackupStateFail)
		return fmt.Errorf("mismatched commp %v != %v", commP.PieceCID, snapshot.SnapshotCommP)
	}

	logger.Sugar().Infow(
		"backupOne",
		"Snapshot", snapshot,
		"Index", index,
		"Mock", b.mock,
		"CommP", commP,
		"Elapsed", time.Since(start).Seconds(),
	)

	proposal, err := b.dealProposal(ctx, snapshot.SnapshotRoot, snapshot.SnapshotCommP, uint64(commP.PieceSize.Unpadded()))
	if err != nil {
		return err
	}

	logger.Sugar().Infow(
		"backupOne",
		"Proposal", proposal,
	)

	signed, err := b.signDealProposal(ctx, proposal)
	if err != nil {
		return err
	}

	logger.Sugar().Infow(
		"backupOne",
		"Signed", signed,
	)

	_cid, err := b.sendDealProposal(ctx, signed, snapshot.SnapshotRoot)
	if err != nil {
		return err
	}

	_state, err := b.getDealStatus(ctx, _cid.String())
	if err != nil {
		return err
	}

	logger.Sugar().Infow(
		"backupOne",
		"Cid", _cid,
		"State", _state,
	)

	_, _ = orbit.Snapshot().UpdateSnapshot(ctx, index, dealerpb.BackupState_BackupStateSuccess)

	return nil
}

func (b *backup) backupAll(ctx context.Context) {
	waits, err := orbit.Backup().Waits(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"backupAll",
			"Error", err,
		)
		return
	}
	for _, wait := range waits {
		if err := b.backupOne(ctx, wait); err != nil {
			logger.Sugar().Errorw(
				"backupAll",
				"Wait", wait,
				"Error", err,
			)
		}
	}
}

var newSnapshot chan struct{}

func Watch(ctx context.Context) (err error) {
	backup := &backup{
		mock: true,
	}

	if err := backup.buildHost(ctx); err != nil {
		return err
	}

	if err := backup.connectMiner(ctx); err != nil {
		return err
	}

	newSnapshot = make(chan struct{})
	backup.backupAll(ctx)

	ticker := time.NewTicker(1 * time.Minute)

	for {
		select {
		case <-ticker.C:
			backup.backupAll(ctx)
		case <-newSnapshot:
			backup.backupAll(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func NewSnapshot() {
	go func() {
		newSnapshot <- struct{}{}
	}()
}

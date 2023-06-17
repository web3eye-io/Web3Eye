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

	metacli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/snapshot"
	metapb "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
)

type backup struct {
	host host.Host
	peer *peer.AddrInfo
	mock bool
}

func (b *backup) mockOne(ctx context.Context) (cid.Cid, string, error) {
	b1 := make([]byte, 1024*1024)
	copy(b1, []byte("0123456789abcdef0123456789abcdef123456789abcdef"))
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

func (b *backup) updateSnapshotState(ctx context.Context, index uint64, state dealerpb.BackupState) error {
	snapshot, err := orbit.Snapshot().UpdateSnapshotState(ctx, index, state)
	if err != nil {
		return err
	}
	for _, item := range snapshot.Items {
		uid := fmt.Sprintf("%v:%v", item.Contract, item.TokenID)
		if err := orbit.FileState().SetFileState(ctx, item.ChainType, uid, item.ChainID, state); err != nil {
			return err
		}
	}

	_state := snapshot.BackupState.String()
	if _, err := metacli.UpdateSnapshot(ctx, &metapb.SnapshotReq{
		ID:          &snapshot.ID,
		BackupState: &_state,
	}); err != nil {
		return err
	}
	return nil
}

func (b *backup) backupOne(ctx context.Context, index uint64) error {
	snapshot, err := orbit.Snapshot().GetSnapshot(ctx, index)
	if err != nil {
		return err
	}

	switch snapshot.BackupState {
	case dealerpb.BackupState_BackupStateCreated:
	default:
		logger.Sugar().Warnw(
			"checkAndUpdateOne",
			"Snapshot", snapshot,
		)
		return nil
	}

	if snapshot.SnapshotCommP == "" {
		_ = b.updateSnapshotState(ctx, index, dealerpb.BackupState_BackupStateFail)
		return fmt.Errorf("invalid snapshot commP")
	}
	if snapshot.SnapshotRoot == "" {
		_ = b.updateSnapshotState(ctx, index, dealerpb.BackupState_BackupStateFail)
		return fmt.Errorf("invalid snapshot root")
	}
	if snapshot.SnapshotURI == "" {
		_ = b.updateSnapshotState(ctx, index, dealerpb.BackupState_BackupStateFail)
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
		_ = b.updateSnapshotState(ctx, index, dealerpb.BackupState_BackupStateFail)
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
	_ = b.updateSnapshotState(ctx, index, dealerpb.BackupState_BackupStateProposed)
	_, _ = orbit.Snapshot().UpdateSnapshotProposalCID(ctx, index, _cid.String())
	_ = orbit.Backup().Wait(ctx, index)

	_state, err := b.getDealStatus(ctx, _cid.String())
	if err != nil {
		return err
	}

	logger.Sugar().Infow(
		"backupOne",
		"Cid", _cid,
		"State", _state,
	)

	return nil
}

func (b *backup) backupAll(ctx context.Context) {
	creates, err := orbit.Backup().Creates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"backupAll",
			"Error", err,
		)
		return
	}
	for _, _create := range creates {
		if err := b.backupOne(ctx, _create); err != nil {
			logger.Sugar().Errorw(
				"backupAll",
				"Create", _create,
				"Error", err,
			)
		}
	}
}

var newSnapshot chan struct{}

func Watch(ctx context.Context) (err error) {
	backup := &backup{
		mock: false,
	}

	if err := backup.buildHost(ctx); err != nil {
		return err
	}

	if err := backup.connectMiner(ctx); err != nil {
		return err
	}

	go backup.check(ctx)

	newSnapshot = make(chan struct{})
	backup.backupAll(ctx)

	ticker := time.NewTicker(time.Minute)

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

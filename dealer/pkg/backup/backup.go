package backup

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/common/oss"
	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"

	"github.com/filecoin-project/go-commp-utils/writer"
	"github.com/filecoin-project/go-fil-markets/storagemarket/network"
	"github.com/ipld/go-car"
	"github.com/libp2p/go-libp2p/core/host"
)

type backup struct {
	host   host.Host
	stream network.StorageDealStream
}

func (b *backup) backupOne(ctx context.Context, index uint64) error {
	snapshot, err := orbit.Snapshot().GetSnapshot(ctx, index)
	if err != nil {
		return err
	}

	logger.Sugar().Infow(
		"backupOne",
		"Snapshot", snapshot,
		"Index", index,
	)

	switch snapshot.BackupState {
	case dealerpb.BackupState_BackupStateSuccess:
		fallthrough // nolint
	case dealerpb.BackupState_BackupStateFail:
		return nil
	default:
	}

	if snapshot.SnapshotCommP == "" {
		return fmt.Errorf("invalid snapshot commP")
	}
	if snapshot.SnapshotRoot == "" {
		return fmt.Errorf("invalid snapshot root")
	}
	if snapshot.SnapshotURI == "" {
		return fmt.Errorf("invalid snapshot uri")
	}

	// TODO: backup items to IPFS
	// TODO: backup car to Filecoin

	buf, err := oss.GetObject(ctx, snapshot.SnapshotURI)
	if err != nil {
		return err
	}

	rdr := bytes.NewReader(buf)
	r := bufio.NewReader(rdr)
	if _, err := car.ReadHeader(r); err != nil {
		return err
	}

	w := &writer.Writer{}
	if _, err := io.CopyBuffer(w, rdr, make([]byte, writer.CommPBuf)); err != nil {
		return err
	}

	commP, err := w.Sum()
	if err != nil {
		return err
	}

	if commP.PieceCID.String() != snapshot.SnapshotCommP {
		return fmt.Errorf("mismatched commp %v != %v", commP.PieceCID, snapshot.SnapshotCommP)
	}

	proposal, err := b.dealProposal(ctx, snapshot.SnapshotRoot, snapshot.SnapshotCommP, uint64(commP.PieceSize.Unpadded()))
	if err != nil {
		return err
	}

	logger.Sugar().Infow(
		"backupOne",
		"Proposal", proposal,
	)

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
	backup := &backup{}

	if err := backup.buildHost(ctx); err != nil {
		return err
	}
	if err := backup.connectMiner(ctx); err != nil {
		return err
	}

	newSnapshot = make(chan struct{})
	backup.backupAll(ctx)

	for {
		select {
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

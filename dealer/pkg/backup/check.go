package backup

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	dealerpb "github.com/web3eye-io/Web3Eye/proto/web3eye/dealer/v1"

	"github.com/filecoin-project/go-fil-markets/storagemarket"
)

var newProposal chan struct{}

func (b *backup) checkAndUpdateOne(ctx context.Context, index uint64) error {
	snapshot, err := orbit.Snapshot().GetSnapshot(ctx, index)
	if err != nil {
		return err
	}

	switch snapshot.BackupState {
	case dealerpb.BackupState_BackupStateProposed:
	case dealerpb.BackupState_BackupStateAccepted:
	default:
		logger.Sugar().Warnw(
			"checkAndUpdateOne",
			"Snapshot", snapshot,
		)
		return nil
	}

	_state, err := b.getDealStatus(ctx, snapshot.ProposalCID)
	if err != nil {
		return err
	}
	logger.Sugar().Warnw(
		"checkAndUpdateOne",
		"Snapshot", snapshot,
		"State", _state,
		"_State", storagemarket.DealStatesDescriptions[_state.State],
	)

	switch _state.State {
	case storagemarket.StorageDealTransferQueued:
	case storagemarket.StorageDealAwaitingPreCommit:
	case storagemarket.StorageDealClientTransferRestart:
	case storagemarket.StorageDealProviderTransferAwaitRestart:
	case storagemarket.StorageDealPublishing:
	case storagemarket.StorageDealPublish:
	case storagemarket.StorageDealClientFunding:
	case storagemarket.StorageDealProviderFunding:
	case storagemarket.StorageDealReserveClientFunds:
	case storagemarket.StorageDealReserveProviderFunds:
	case storagemarket.StorageDealVerifyData:
	case storagemarket.StorageDealWaitingForData:
	case storagemarket.StorageDealTransferring:
	case storagemarket.StorageDealStartDataTransfer:
	case storagemarket.StorageDealAcceptWait:
	case storagemarket.StorageDealValidating:
	case storagemarket.StorageDealCheckForAcceptance:
	case storagemarket.StorageDealFundsReserved:
	case storagemarket.StorageDealActive:
		if err := b.updateSnapshotState(ctx, index, dealerpb.BackupState_BackupStateSuccess); err != nil {
			return err
		}
		_ = orbit.Backup().Done(ctx, index, false)
	case storagemarket.StorageDealFinalizing:
	case storagemarket.StorageDealSealing:
	case storagemarket.StorageDealStaged:
	case storagemarket.StorageDealProposalAccepted:
		_ = b.updateSnapshotState(ctx, index, dealerpb.BackupState_BackupStateAccepted)
	case storagemarket.StorageDealError:
		fallthrough //nolint
	case storagemarket.StorageDealFailing:
		fallthrough //nolint
	case storagemarket.StorageDealRejecting:
		fallthrough //nolint
	case storagemarket.StorageDealSlashed:
		fallthrough //nolint
	case storagemarket.StorageDealExpired:
		fallthrough //nolint
	case storagemarket.StorageDealProposalRejected:
		fallthrough //nolint
	case storagemarket.StorageDealProposalNotFound:
		fallthrough //nolint
	case storagemarket.StorageDealUnknown:
		if err := b.updateSnapshotState(ctx, index, dealerpb.BackupState_BackupStateFail); err != nil {
			return err
		}
		_ = orbit.Backup().Done(ctx, index, true)
		return fmt.Errorf("proposal error %v", storagemarket.DealStatesDescriptions[_state.State])
	}

	return nil
}

func (b *backup) checkAndUpdateAll(ctx context.Context) {
	waits, err := orbit.Backup().Waits(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"checkAndUpdateAll",
			"Error", err,
		)
		return
	}
	for _, _wait := range waits {
		if err := b.checkAndUpdateOne(ctx, _wait); err != nil {
			logger.Sugar().Errorw(
				"checkAndUpdateAll",
				"Wait", _wait,
				"Error", err,
			)
		}
	}
}

func (b *backup) check(ctx context.Context) {
	newProposal = make(chan struct{})
	ticker := time.NewTicker(time.Minute)

	for {
		select {
		case <-ticker.C:
			b.checkAndUpdateAll(ctx)
		case <-newProposal:
			b.checkAndUpdateAll(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func NewProposal() {
	go func() {
		newProposal <- struct{}{}
	}()
}

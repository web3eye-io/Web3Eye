package retrieve

import (
	"context"
	_ "encoding/base64"
	"fmt"
	"os"
	"os/exec"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	orbit "github.com/web3eye-io/Web3Eye/dealer/pkg/orbit"
	retrieverpb "github.com/web3eye-io/Web3Eye/proto/web3eye/retriever/v1"
)

const mock = true

func (h *Handler) StartRetrieve(ctx context.Context) (*retrieverpb.Retrieve, error) {
	os.Setenv("FULLNODE_API_INFO", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.pSLTE16BE7rYSzlpcAdPUqQz5V4tF1ksGIWZmAaB_Rc:/ip4/210.209.69.38/tcp/20803/http")

	uid := fmt.Sprintf("%v:%v", h.Contract, h.TokenID)
	index, err := orbit.FileState().GetFileSnapshot(ctx, h.ChainType, uid, h.ChainID)
	if err != nil {
		return nil, err
	}
	snapshot, err := orbit.Snapshot().GetSnapshot(ctx, index)
	if err != nil {
		return nil, err
	}
	if snapshot.DealID == 0 {
		if !mock {
			return nil, fmt.Errorf("file backup not exist")
		}
		snapshot.SnapshotRoot = "bafykbzacedl3t5ozakdth7uyhptc27efkdhqbhshelqqqlx2z2ijenli3giqk"
	}

	_ = orbit.FileState().SetFileRetrieve(ctx, h.ChainType, uid, h.ChainID, "Starting")
	go func() {
		_ = orbit.FileState().SetFileRetrieve(ctx, h.ChainType, uid, h.ChainID, "Running")
		fileName := fmt.Sprintf("/opt/%v.car", snapshot.SnapshotRoot)
		cmd := exec.Command("lotus", "client", "retrieve", "--provider", "t01002", snapshot.SnapshotRoot, fileName)
		if out, err := cmd.Output(); err != nil {
			_ = orbit.FileState().SetFileRetrieve(ctx, h.ChainType, uid, h.ChainID, "Fail")
			logger.Sugar().Errorw(
				"StartRetrieve",
				"Error", err,
				"Output", string(out),
			)
			return
		}
		_ = orbit.FileState().SetFileRetrieve(ctx, h.ChainType, uid, h.ChainID, "Success")
	}()

	return &retrieverpb.Retrieve{
		ChainType:        h.ChainType,
		ChainID:          h.ChainID,
		Contract:         h.Contract,
		TokenID:          h.TokenID,
		RetrieveState:    "Running",
		ProposalCID:      snapshot.ProposalCID,
		DealID:           snapshot.DealID,
		BackupPayloadCID: snapshot.SnapshotRoot,
	}, nil
}

package snapshot

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/snapshot"
)

func Ent2Grpc(row *ent.Snapshot) *npool.Snapshot {
	if row == nil {
		return nil
	}

	return &npool.Snapshot{
		ID:            row.ID.String(),
		Index:         row.Index,
		SnapshotCommP: row.SnapshotCommP,
		SnapshotRoot:  row.SnapshotRoot,
		SnapshotURI:   row.SnapshotURI,
		BackupState:   npool.BackupState(npool.BackupState_value[row.BackupState]),
	}
}

func Ent2GrpcMany(rows []*ent.Snapshot) []*npool.Snapshot {
	infos := []*npool.Snapshot{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}

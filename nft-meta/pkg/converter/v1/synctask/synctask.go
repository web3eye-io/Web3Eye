package synctask

import (
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent"
	"github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/cttype"
	npool "github.com/web3eye-io/cyber-tracer/proto/cybertracer/nftmeta/v1/synctask"
)

func Ent2Grpc(row *ent.SyncTask) *npool.SyncTask {
	if row == nil {
		return nil
	}
	chainT := cttype.ChainType(cttype.ChainType_value[row.ChainType])
	syncS := cttype.SyncState(cttype.SyncState_value[row.SyncState])

	return &npool.SyncTask{
		ID:          row.ID.String(),
		ChainType:   chainT,
		ChainID:     row.ChainID,
		Start:       row.Start,
		End:         row.End,
		Current:     row.Current,
		Topic:       row.Topic,
		Description: row.Description,
		SyncState:   syncS,
		Remark:      row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.SyncTask) []*npool.SyncTask {
	infos := []*npool.SyncTask{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}

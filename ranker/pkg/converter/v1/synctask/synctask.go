package synctask

import (
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/cttype"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/synctask"
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

func Ent2GrpcReq(row *ent.SyncTask) *npool.SyncTaskReq {
	if row == nil {
		return nil
	}
	chainT := cttype.ChainType(cttype.ChainType_value[row.ChainType])
	syncS := cttype.SyncState(cttype.SyncState_value[row.SyncState])

	id := row.ID.String()
	return &npool.SyncTaskReq{
		ID:          &id,
		ChainType:   &chainT,
		ChainID:     &row.ChainID,
		Start:       &row.Start,
		End:         &row.End,
		Current:     &row.Current,
		Topic:       &row.Topic,
		Description: &row.Description,
		SyncState:   &syncS,
		Remark:      &row.Remark,
	}
}

func Ent2GrpcMany(rows []*ent.SyncTask) []*npool.SyncTask {
	infos := []*npool.SyncTask{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}

func Grpc2Ent(row *npool.SyncTask) *ent.SyncTask {
	if row == nil {
		return nil
	}
	id, err := uuid.Parse(row.ID)
	if err != nil {
		id = uuid.New()
	}

	return &ent.SyncTask{
		ID:          id,
		ChainType:   row.ChainType.String(),
		ChainID:     row.ChainID,
		Start:       row.Start,
		End:         row.End,
		Current:     row.Current,
		Topic:       row.Topic,
		Description: row.Description,
		SyncState:   row.SyncState.String(),
		Remark:      row.Remark,
	}
}

func Grpc2EntMany(rows []*npool.SyncTask) []*ent.SyncTask {
	infos := []*ent.SyncTask{}
	for _, row := range rows {
		infos = append(infos, Grpc2Ent(row))
	}
	return infos
}

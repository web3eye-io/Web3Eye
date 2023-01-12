package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/cttype"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type SyncTask struct {
	ent.Schema
}

func (SyncTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (SyncTask) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("chain_type").Optional().Default(npool.ChainType_Unknown.String()),
		field.Int32("chain_id"),
		field.Uint64("start"),
		field.Uint64("end"),
		field.Uint64("current"),
		field.String("topic").Unique(),
		field.String("description").
			Optional(),
		field.String("sync_state").Optional().Default(npool.SyncState_Default.String()),
		field.Text("remark").Optional(),
	}
}

func (SyncTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("topic"),
	}
}

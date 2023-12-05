package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"

	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type SyncTask struct {
	ent.Schema
}

func (SyncTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

func (SyncTask) Fields() []ent.Field {
	return []ent.Field{
		field.String("chain_type").Optional().Default(basetype.ChainType_ChainUnkonwn.String()),
		field.String("chain_id"),
		field.Uint64("start"),
		field.Uint64("end"),
		field.Uint64("current"),
		field.String("topic").Unique(),
		field.String("description").
			Optional(),
		field.String("sync_state").Optional().Default(basetype.SyncState_Default.String()),
		field.Text("remark").Optional(),
	}
}

func (SyncTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("topic"),
	}
}

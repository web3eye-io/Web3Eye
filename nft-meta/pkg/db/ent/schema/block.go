package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type Block struct {
	ent.Schema
}

func (Block) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

func (Block) Fields() []ent.Field {
	return []ent.Field{
		field.String("chain_type"),
		field.String("chain_id"),
		field.Uint64("block_number"),
		field.String("block_hash"),
		field.Uint64("block_time"),
		field.String("parse_state"),
		field.String("remark"),
	}
}

func (Block) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chain_type", "chain_id", "block_number").
			Unique(),
	}
}

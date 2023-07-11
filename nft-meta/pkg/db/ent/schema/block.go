package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type Block struct {
	ent.Schema
}

func (Block) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Block) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("chain_type"),
		field.String("chain_id"),
		field.Uint64("block_number"),
		field.String("block_hash"),
		field.Int64("block_time"),
	}
}

func (Block) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chain_type", "chain_id", "block_number").
			Unique(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type Transfer struct {
	ent.Schema
}

func (Transfer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Transfer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("chain_type"),
		field.Int32("chain_id"),
		field.String("contract"),
		field.String("token_type"),
		field.String("token_id"),
		field.String("from"),
		field.String("to"),
		field.Uint64("amount"),
		field.Uint64("block_number"),
		field.String("tx_hash"),
		field.String("block_hash"),
		field.Uint32("tx_time").
			Optional(),
		field.String("remark").
			Optional(),
	}
}

func (Transfer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("contract", "token_id"),
		index.Fields("tx_hash", "token_id"),
	}
}

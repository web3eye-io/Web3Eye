package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type Order struct {
	ent.Schema
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("chain_type"),
		field.String("chain_id"),
		field.String("tx_hash").MaxLen(128),
		field.Uint64("block_number"),
		field.Uint32("tx_index"),
		field.Uint32("log_index"),
		field.String("recipient").MaxLen(128),
		field.Text("remark").
			Optional(),
	}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tx_hash", "recipient", "log_index").Unique(),
	}
}

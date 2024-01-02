package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

var (
	MaxTXHashLen    = 128
	MaxRecipientLen = 128
)

type Order struct {
	ent.Schema
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		mixin.TimeMixin{},
	}
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("chain_type"),
		field.String("chain_id"),
		field.String("tx_hash").MaxLen(MaxTXHashLen),
		field.Uint64("block_number"),
		field.Uint32("tx_index"),
		field.Uint32("log_index"),
		field.String("recipient").MaxLen(MaxRecipientLen),
		field.Text("remark").
			Optional(),
	}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tx_hash", "recipient", "log_index").Unique(),
	}
}

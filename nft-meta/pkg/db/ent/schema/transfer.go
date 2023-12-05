package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

var (
	MaxAddressLen = 128
)

type Transfer struct {
	ent.Schema
}

func (Transfer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		mixin.TimeMixin{},
	}
}

func (Transfer) Fields() []ent.Field {
	return []ent.Field{
		field.String("chain_type"),
		field.String("chain_id"),
		field.String("contract").MaxLen(MaxAddressLen),
		field.String("token_type"),
		field.String("token_id"),
		field.String("from").MaxLen(MaxAddressLen),
		field.String("to").MaxLen(MaxAddressLen),
		field.String("amount"),
		field.Uint64("block_number"),
		field.String("tx_hash"),
		field.String("block_hash"),
		field.Uint32("tx_time").
			Optional(),
		field.Text("remark").
			Optional(),
	}
}

func (Transfer) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("contract", "token_id", "tx_hash", "from").Unique(),
	}
}

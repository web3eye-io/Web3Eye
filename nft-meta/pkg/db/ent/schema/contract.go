package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type Contract struct {
	ent.Schema
}

func (Contract) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Contract) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("chain_type"),
		field.String("chain_id"),
		field.String("address"),
		field.String("name"),
		field.String("symbol"),
		field.Uint32("decimals").Default(0),
		field.String("creator").
			Optional(),
		field.Uint64("block_num").
			Optional(),
		field.String("tx_hash").
			Optional(),
		field.Uint32("tx_time").
			Optional(),
		field.String("profile_url").
			Optional(),
		field.String("base_url").
			Optional(),
		field.String("banner_url").
			Optional(),
		field.String("description").
			Optional(),
		field.Text("remark").
			Optional(),
	}
}

func (Contract) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chain_type", "chain_id", "address").
			Unique(),
	}
}

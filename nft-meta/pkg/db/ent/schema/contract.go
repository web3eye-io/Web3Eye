package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/mixin"
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
		field.Int32("chain_id"),
		field.String("address"),
		field.String("name"),
		field.String("symbol"),
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
		field.String("remark").
			Optional(),
	}
}

func (Contract) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chain_type", "chain_id", "address").
			Unique(),
	}
}

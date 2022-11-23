package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/mixin"
)

type BlockNumber struct {
	ent.Schema
}

func (BlockNumber) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (BlockNumber) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("chain_type"),
		field.Int32("chain_id"),
		field.String("identifier"),
		field.Uint64("current_num"),
		field.String("topic"),
		field.String("description").
			Optional(),
	}
}

func (BlockNumber) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("identifier", "topic"),
	}
}

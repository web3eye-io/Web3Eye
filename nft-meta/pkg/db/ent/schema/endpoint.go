package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type Endpoint struct {
	ent.Schema
}

func (Endpoint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Endpoint) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("chain_type"),
		field.String("chain_id").Optional(),
		field.String("address"),
		field.String("state").Optional(),
		field.String("remark").Optional(),
	}
}

func (Endpoint) Indexes() []ent.Index {
	return []ent.Index{}
}

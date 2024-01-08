package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type Endpoint struct {
	ent.Schema
}

func (Endpoint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		mixin.TimeMixin{},
	}
}

func (Endpoint) Fields() []ent.Field {
	return []ent.Field{
		field.String("chain_type"),
		field.String("chain_id").Optional(),
		field.String("address"),
		field.String("state").Optional(),
		field.Uint32("rps").Default(1),
		field.String("remark").Optional(),
	}
}

func (Endpoint) Indexes() []ent.Index {
	return []ent.Index{}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type OrderItem struct {
	ent.Schema
}

func (OrderItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("contract"),
		field.String("token_type"),
		field.String("token_id"),
		field.Uint64("amount"),
		field.Uint32("portion_num"),
		field.Text("remark").
			Optional(),
	}
}

func (OrderItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

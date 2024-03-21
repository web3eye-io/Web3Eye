package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type OrderItem struct {
	ent.Schema
}

func (OrderItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		mixin.TimeMixin{},
	}
}

func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("order_id", uuid.UUID{}),
		field.String("order_item_type"),
		field.String("contract"),
		field.String("token_type"),
		field.String("token_id"),
		field.String("amount"),
		field.Text("remark").
			Optional(),
	}
}

func (OrderItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type OrderPair struct {
	ent.Schema
}

func (OrderPair) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (OrderPair) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("tx_hash"),
		field.String("recipient"),
		field.String("target_id"),
		field.String("offer_id"),
		field.Text("remark").
			Optional(),
	}
}

func (OrderPair) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tx_hash", "recipient"),
	}
}

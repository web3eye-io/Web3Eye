package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	npool "github.com/web3eye-io/cyber-tracer/message/cybertracer/nftmeta/v1/token"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/mixin"
)

type Token struct {
	ent.Schema
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("chain_type"),
		field.Int32("chain_id"),
		field.String("contract"),
		field.String("token_type"),
		field.String("token_id"),
		field.String("owner").
			Optional(),
		field.String("uri").
			Optional(),
		field.String("uri_type").
			Optional(),
		field.Text("image_url").
			Optional(),
		field.Text("video_url").
			Optional(),
		field.Text("description").
			Optional(),
		field.String("name").
			Optional(),
		field.Int64("vector_id").
			Optional(),
		field.String("vector_state").
			Optional().Default(npool.ConvertState_Default.String()),
		field.String("remark").
			Optional(),
	}
}

func (Token) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("contract", "token_id").
			Unique(),
	}
}

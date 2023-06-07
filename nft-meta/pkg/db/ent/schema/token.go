package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
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
		field.String("chain_id"),
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
		field.Text("remark").
			Optional(),
		field.Text("ipfs_image_url").
			Optional(),
		field.Text("file_cid").
			Optional(),
	}
}

func (Token) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("contract", "token_id").
			Unique(),
	}
}

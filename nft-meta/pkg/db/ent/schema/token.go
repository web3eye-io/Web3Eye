package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
	npool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
)

type Token struct {
	ent.Schema
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		mixin.TimeMixin{},
	}
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("chain_type"),
		field.String("chain_id"),
		field.String("contract"),
		field.String("token_type"),
		field.String("token_id"),
		field.String("owner").
			Optional(),
		field.Text("uri").
			Optional(),
		field.Text("uri_state").
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
		field.Uint32("image_snapshot_id").
			Optional(),
	}
}

func (Token) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("contract", "token_id").
			Unique(),
	}
}

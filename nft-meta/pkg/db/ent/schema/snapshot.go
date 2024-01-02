package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/mixin"
)

type Snapshot struct {
	ent.Schema
}

func (Snapshot) Mixin() []ent.Mixin {
	return []ent.Mixin{
		crudermixin.AutoIDMixin{},
		mixin.TimeMixin{},
	}
}

func (Snapshot) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("index"),
		field.String("snapshot_comm_p"),
		field.String("snapshot_root"),
		field.String("snapshot_uri"),
		field.String("backup_state"),
	}
}

func (Snapshot) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("index", "backup_state").
			Unique(),
	}
}

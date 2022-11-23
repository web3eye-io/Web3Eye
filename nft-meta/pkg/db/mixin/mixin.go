package mixin

import (
	"entgo.io/ent"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent/privacy"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/rule"
)

func (TimeMixin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (TimeMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterTimeRule(),
		},
	}
}

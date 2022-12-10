// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/google/uuid"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent/contract"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent/schema"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent/synctask"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent/token"
	"github.com/web3eye-io/cyber-tracer/nft-meta/pkg/db/ent/transfer"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	contractMixin := schema.Contract{}.Mixin()
	contract.Policy = privacy.NewPolicies(contractMixin[0], schema.Contract{})
	contract.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := contract.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	contractMixinFields0 := contractMixin[0].Fields()
	_ = contractMixinFields0
	contractFields := schema.Contract{}.Fields()
	_ = contractFields
	// contractDescCreatedAt is the schema descriptor for created_at field.
	contractDescCreatedAt := contractMixinFields0[0].Descriptor()
	// contract.DefaultCreatedAt holds the default value on creation for the created_at field.
	contract.DefaultCreatedAt = contractDescCreatedAt.Default.(func() uint32)
	// contractDescUpdatedAt is the schema descriptor for updated_at field.
	contractDescUpdatedAt := contractMixinFields0[1].Descriptor()
	// contract.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	contract.DefaultUpdatedAt = contractDescUpdatedAt.Default.(func() uint32)
	// contract.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	contract.UpdateDefaultUpdatedAt = contractDescUpdatedAt.UpdateDefault.(func() uint32)
	// contractDescDeletedAt is the schema descriptor for deleted_at field.
	contractDescDeletedAt := contractMixinFields0[2].Descriptor()
	// contract.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	contract.DefaultDeletedAt = contractDescDeletedAt.Default.(func() uint32)
	// contractDescID is the schema descriptor for id field.
	contractDescID := contractFields[0].Descriptor()
	// contract.DefaultID holds the default value on creation for the id field.
	contract.DefaultID = contractDescID.Default.(func() uuid.UUID)
	synctaskMixin := schema.SyncTask{}.Mixin()
	synctask.Policy = privacy.NewPolicies(synctaskMixin[0], schema.SyncTask{})
	synctask.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := synctask.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	synctaskMixinFields0 := synctaskMixin[0].Fields()
	_ = synctaskMixinFields0
	synctaskFields := schema.SyncTask{}.Fields()
	_ = synctaskFields
	// synctaskDescCreatedAt is the schema descriptor for created_at field.
	synctaskDescCreatedAt := synctaskMixinFields0[0].Descriptor()
	// synctask.DefaultCreatedAt holds the default value on creation for the created_at field.
	synctask.DefaultCreatedAt = synctaskDescCreatedAt.Default.(func() uint32)
	// synctaskDescUpdatedAt is the schema descriptor for updated_at field.
	synctaskDescUpdatedAt := synctaskMixinFields0[1].Descriptor()
	// synctask.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	synctask.DefaultUpdatedAt = synctaskDescUpdatedAt.Default.(func() uint32)
	// synctask.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	synctask.UpdateDefaultUpdatedAt = synctaskDescUpdatedAt.UpdateDefault.(func() uint32)
	// synctaskDescDeletedAt is the schema descriptor for deleted_at field.
	synctaskDescDeletedAt := synctaskMixinFields0[2].Descriptor()
	// synctask.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	synctask.DefaultDeletedAt = synctaskDescDeletedAt.Default.(func() uint32)
	// synctaskDescChainType is the schema descriptor for chain_type field.
	synctaskDescChainType := synctaskFields[1].Descriptor()
	// synctask.DefaultChainType holds the default value on creation for the chain_type field.
	synctask.DefaultChainType = synctaskDescChainType.Default.(string)
	// synctaskDescSyncState is the schema descriptor for sync_state field.
	synctaskDescSyncState := synctaskFields[8].Descriptor()
	// synctask.DefaultSyncState holds the default value on creation for the sync_state field.
	synctask.DefaultSyncState = synctaskDescSyncState.Default.(string)
	// synctaskDescID is the schema descriptor for id field.
	synctaskDescID := synctaskFields[0].Descriptor()
	// synctask.DefaultID holds the default value on creation for the id field.
	synctask.DefaultID = synctaskDescID.Default.(func() uuid.UUID)
	tokenMixin := schema.Token{}.Mixin()
	token.Policy = privacy.NewPolicies(tokenMixin[0], schema.Token{})
	token.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := token.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	tokenMixinFields0 := tokenMixin[0].Fields()
	_ = tokenMixinFields0
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescCreatedAt is the schema descriptor for created_at field.
	tokenDescCreatedAt := tokenMixinFields0[0].Descriptor()
	// token.DefaultCreatedAt holds the default value on creation for the created_at field.
	token.DefaultCreatedAt = tokenDescCreatedAt.Default.(func() uint32)
	// tokenDescUpdatedAt is the schema descriptor for updated_at field.
	tokenDescUpdatedAt := tokenMixinFields0[1].Descriptor()
	// token.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	token.DefaultUpdatedAt = tokenDescUpdatedAt.Default.(func() uint32)
	// token.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	token.UpdateDefaultUpdatedAt = tokenDescUpdatedAt.UpdateDefault.(func() uint32)
	// tokenDescDeletedAt is the schema descriptor for deleted_at field.
	tokenDescDeletedAt := tokenMixinFields0[2].Descriptor()
	// token.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	token.DefaultDeletedAt = tokenDescDeletedAt.Default.(func() uint32)
	// tokenDescVectorState is the schema descriptor for vector_state field.
	tokenDescVectorState := tokenFields[14].Descriptor()
	// token.DefaultVectorState holds the default value on creation for the vector_state field.
	token.DefaultVectorState = tokenDescVectorState.Default.(string)
	// tokenDescID is the schema descriptor for id field.
	tokenDescID := tokenFields[0].Descriptor()
	// token.DefaultID holds the default value on creation for the id field.
	token.DefaultID = tokenDescID.Default.(func() uuid.UUID)
	transferMixin := schema.Transfer{}.Mixin()
	transfer.Policy = privacy.NewPolicies(transferMixin[0], schema.Transfer{})
	transfer.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := transfer.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	transferMixinFields0 := transferMixin[0].Fields()
	_ = transferMixinFields0
	transferFields := schema.Transfer{}.Fields()
	_ = transferFields
	// transferDescCreatedAt is the schema descriptor for created_at field.
	transferDescCreatedAt := transferMixinFields0[0].Descriptor()
	// transfer.DefaultCreatedAt holds the default value on creation for the created_at field.
	transfer.DefaultCreatedAt = transferDescCreatedAt.Default.(func() uint32)
	// transferDescUpdatedAt is the schema descriptor for updated_at field.
	transferDescUpdatedAt := transferMixinFields0[1].Descriptor()
	// transfer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	transfer.DefaultUpdatedAt = transferDescUpdatedAt.Default.(func() uint32)
	// transfer.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	transfer.UpdateDefaultUpdatedAt = transferDescUpdatedAt.UpdateDefault.(func() uint32)
	// transferDescDeletedAt is the schema descriptor for deleted_at field.
	transferDescDeletedAt := transferMixinFields0[2].Descriptor()
	// transfer.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	transfer.DefaultDeletedAt = transferDescDeletedAt.Default.(func() uint32)
	// transferDescID is the schema descriptor for id field.
	transferDescID := transferFields[0].Descriptor()
	// transfer.DefaultID holds the default value on creation for the id field.
	transfer.DefaultID = transferDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.2"                                         // Version of ent codegen.
	Sum     = "h1:UM2/BUhF2FfsxPHRxLjQbhqJNaDdVlOwNIAMLs2jyto=" // Sum of ent codegen.
)

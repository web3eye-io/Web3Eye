// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/contract"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/predicate"
)

// ContractUpdate is the builder for updating Contract entities.
type ContractUpdate struct {
	config
	hooks     []Hook
	mutation  *ContractMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ContractUpdate builder.
func (cu *ContractUpdate) Where(ps ...predicate.Contract) *ContractUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *ContractUpdate) SetCreatedAt(u uint32) *ContractUpdate {
	cu.mutation.ResetCreatedAt()
	cu.mutation.SetCreatedAt(u)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableCreatedAt(u *uint32) *ContractUpdate {
	if u != nil {
		cu.SetCreatedAt(*u)
	}
	return cu
}

// AddCreatedAt adds u to the "created_at" field.
func (cu *ContractUpdate) AddCreatedAt(u int32) *ContractUpdate {
	cu.mutation.AddCreatedAt(u)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ContractUpdate) SetUpdatedAt(u uint32) *ContractUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(u)
	return cu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cu *ContractUpdate) AddUpdatedAt(u int32) *ContractUpdate {
	cu.mutation.AddUpdatedAt(u)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *ContractUpdate) SetDeletedAt(u uint32) *ContractUpdate {
	cu.mutation.ResetDeletedAt()
	cu.mutation.SetDeletedAt(u)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableDeletedAt(u *uint32) *ContractUpdate {
	if u != nil {
		cu.SetDeletedAt(*u)
	}
	return cu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cu *ContractUpdate) AddDeletedAt(u int32) *ContractUpdate {
	cu.mutation.AddDeletedAt(u)
	return cu
}

// SetChainType sets the "chain_type" field.
func (cu *ContractUpdate) SetChainType(s string) *ContractUpdate {
	cu.mutation.SetChainType(s)
	return cu
}

// SetChainID sets the "chain_id" field.
func (cu *ContractUpdate) SetChainID(s string) *ContractUpdate {
	cu.mutation.SetChainID(s)
	return cu
}

// SetAddress sets the "address" field.
func (cu *ContractUpdate) SetAddress(s string) *ContractUpdate {
	cu.mutation.SetAddress(s)
	return cu
}

// SetName sets the "name" field.
func (cu *ContractUpdate) SetName(s string) *ContractUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetSymbol sets the "symbol" field.
func (cu *ContractUpdate) SetSymbol(s string) *ContractUpdate {
	cu.mutation.SetSymbol(s)
	return cu
}

// SetDecimals sets the "decimals" field.
func (cu *ContractUpdate) SetDecimals(u uint32) *ContractUpdate {
	cu.mutation.ResetDecimals()
	cu.mutation.SetDecimals(u)
	return cu
}

// SetNillableDecimals sets the "decimals" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableDecimals(u *uint32) *ContractUpdate {
	if u != nil {
		cu.SetDecimals(*u)
	}
	return cu
}

// AddDecimals adds u to the "decimals" field.
func (cu *ContractUpdate) AddDecimals(u int32) *ContractUpdate {
	cu.mutation.AddDecimals(u)
	return cu
}

// SetCreator sets the "creator" field.
func (cu *ContractUpdate) SetCreator(s string) *ContractUpdate {
	cu.mutation.SetCreator(s)
	return cu
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableCreator(s *string) *ContractUpdate {
	if s != nil {
		cu.SetCreator(*s)
	}
	return cu
}

// ClearCreator clears the value of the "creator" field.
func (cu *ContractUpdate) ClearCreator() *ContractUpdate {
	cu.mutation.ClearCreator()
	return cu
}

// SetBlockNum sets the "block_num" field.
func (cu *ContractUpdate) SetBlockNum(u uint64) *ContractUpdate {
	cu.mutation.ResetBlockNum()
	cu.mutation.SetBlockNum(u)
	return cu
}

// SetNillableBlockNum sets the "block_num" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableBlockNum(u *uint64) *ContractUpdate {
	if u != nil {
		cu.SetBlockNum(*u)
	}
	return cu
}

// AddBlockNum adds u to the "block_num" field.
func (cu *ContractUpdate) AddBlockNum(u int64) *ContractUpdate {
	cu.mutation.AddBlockNum(u)
	return cu
}

// ClearBlockNum clears the value of the "block_num" field.
func (cu *ContractUpdate) ClearBlockNum() *ContractUpdate {
	cu.mutation.ClearBlockNum()
	return cu
}

// SetTxHash sets the "tx_hash" field.
func (cu *ContractUpdate) SetTxHash(s string) *ContractUpdate {
	cu.mutation.SetTxHash(s)
	return cu
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableTxHash(s *string) *ContractUpdate {
	if s != nil {
		cu.SetTxHash(*s)
	}
	return cu
}

// ClearTxHash clears the value of the "tx_hash" field.
func (cu *ContractUpdate) ClearTxHash() *ContractUpdate {
	cu.mutation.ClearTxHash()
	return cu
}

// SetTxTime sets the "tx_time" field.
func (cu *ContractUpdate) SetTxTime(u uint32) *ContractUpdate {
	cu.mutation.ResetTxTime()
	cu.mutation.SetTxTime(u)
	return cu
}

// SetNillableTxTime sets the "tx_time" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableTxTime(u *uint32) *ContractUpdate {
	if u != nil {
		cu.SetTxTime(*u)
	}
	return cu
}

// AddTxTime adds u to the "tx_time" field.
func (cu *ContractUpdate) AddTxTime(u int32) *ContractUpdate {
	cu.mutation.AddTxTime(u)
	return cu
}

// ClearTxTime clears the value of the "tx_time" field.
func (cu *ContractUpdate) ClearTxTime() *ContractUpdate {
	cu.mutation.ClearTxTime()
	return cu
}

// SetProfileURL sets the "profile_url" field.
func (cu *ContractUpdate) SetProfileURL(s string) *ContractUpdate {
	cu.mutation.SetProfileURL(s)
	return cu
}

// SetNillableProfileURL sets the "profile_url" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableProfileURL(s *string) *ContractUpdate {
	if s != nil {
		cu.SetProfileURL(*s)
	}
	return cu
}

// ClearProfileURL clears the value of the "profile_url" field.
func (cu *ContractUpdate) ClearProfileURL() *ContractUpdate {
	cu.mutation.ClearProfileURL()
	return cu
}

// SetBaseURL sets the "base_url" field.
func (cu *ContractUpdate) SetBaseURL(s string) *ContractUpdate {
	cu.mutation.SetBaseURL(s)
	return cu
}

// SetNillableBaseURL sets the "base_url" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableBaseURL(s *string) *ContractUpdate {
	if s != nil {
		cu.SetBaseURL(*s)
	}
	return cu
}

// ClearBaseURL clears the value of the "base_url" field.
func (cu *ContractUpdate) ClearBaseURL() *ContractUpdate {
	cu.mutation.ClearBaseURL()
	return cu
}

// SetBannerURL sets the "banner_url" field.
func (cu *ContractUpdate) SetBannerURL(s string) *ContractUpdate {
	cu.mutation.SetBannerURL(s)
	return cu
}

// SetNillableBannerURL sets the "banner_url" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableBannerURL(s *string) *ContractUpdate {
	if s != nil {
		cu.SetBannerURL(*s)
	}
	return cu
}

// ClearBannerURL clears the value of the "banner_url" field.
func (cu *ContractUpdate) ClearBannerURL() *ContractUpdate {
	cu.mutation.ClearBannerURL()
	return cu
}

// SetDescription sets the "description" field.
func (cu *ContractUpdate) SetDescription(s string) *ContractUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableDescription(s *string) *ContractUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *ContractUpdate) ClearDescription() *ContractUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetRemark sets the "remark" field.
func (cu *ContractUpdate) SetRemark(s string) *ContractUpdate {
	cu.mutation.SetRemark(s)
	return cu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cu *ContractUpdate) SetNillableRemark(s *string) *ContractUpdate {
	if s != nil {
		cu.SetRemark(*s)
	}
	return cu
}

// ClearRemark clears the value of the "remark" field.
func (cu *ContractUpdate) ClearRemark() *ContractUpdate {
	cu.mutation.ClearRemark()
	return cu
}

// Mutation returns the ContractMutation object of the builder.
func (cu *ContractUpdate) Mutation() *ContractMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContractUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContractUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContractUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContractUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ContractUpdate) defaults() error {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		if contract.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized contract.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := contract.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *ContractUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ContractUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *ContractUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contract.Table,
			Columns: contract.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: contract.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.ChainType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldChainType,
		})
	}
	if value, ok := cu.mutation.ChainID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldChainID,
		})
	}
	if value, ok := cu.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldAddress,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldName,
		})
	}
	if value, ok := cu.mutation.Symbol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldSymbol,
		})
	}
	if value, ok := cu.mutation.Decimals(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldDecimals,
		})
	}
	if value, ok := cu.mutation.AddedDecimals(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldDecimals,
		})
	}
	if value, ok := cu.mutation.Creator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldCreator,
		})
	}
	if cu.mutation.CreatorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldCreator,
		})
	}
	if value, ok := cu.mutation.BlockNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: contract.FieldBlockNum,
		})
	}
	if value, ok := cu.mutation.AddedBlockNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: contract.FieldBlockNum,
		})
	}
	if cu.mutation.BlockNumCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: contract.FieldBlockNum,
		})
	}
	if value, ok := cu.mutation.TxHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldTxHash,
		})
	}
	if cu.mutation.TxHashCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldTxHash,
		})
	}
	if value, ok := cu.mutation.TxTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldTxTime,
		})
	}
	if value, ok := cu.mutation.AddedTxTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldTxTime,
		})
	}
	if cu.mutation.TxTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: contract.FieldTxTime,
		})
	}
	if value, ok := cu.mutation.ProfileURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldProfileURL,
		})
	}
	if cu.mutation.ProfileURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldProfileURL,
		})
	}
	if value, ok := cu.mutation.BaseURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldBaseURL,
		})
	}
	if cu.mutation.BaseURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldBaseURL,
		})
	}
	if value, ok := cu.mutation.BannerURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldBannerURL,
		})
	}
	if cu.mutation.BannerURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldBannerURL,
		})
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldDescription,
		})
	}
	if cu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldDescription,
		})
	}
	if value, ok := cu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldRemark,
		})
	}
	if cu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldRemark,
		})
	}
	_spec.Modifiers = cu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contract.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ContractUpdateOne is the builder for updating a single Contract entity.
type ContractUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ContractMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cuo *ContractUpdateOne) SetCreatedAt(u uint32) *ContractUpdateOne {
	cuo.mutation.ResetCreatedAt()
	cuo.mutation.SetCreatedAt(u)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableCreatedAt(u *uint32) *ContractUpdateOne {
	if u != nil {
		cuo.SetCreatedAt(*u)
	}
	return cuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cuo *ContractUpdateOne) AddCreatedAt(u int32) *ContractUpdateOne {
	cuo.mutation.AddCreatedAt(u)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ContractUpdateOne) SetUpdatedAt(u uint32) *ContractUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(u)
	return cuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cuo *ContractUpdateOne) AddUpdatedAt(u int32) *ContractUpdateOne {
	cuo.mutation.AddUpdatedAt(u)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *ContractUpdateOne) SetDeletedAt(u uint32) *ContractUpdateOne {
	cuo.mutation.ResetDeletedAt()
	cuo.mutation.SetDeletedAt(u)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableDeletedAt(u *uint32) *ContractUpdateOne {
	if u != nil {
		cuo.SetDeletedAt(*u)
	}
	return cuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cuo *ContractUpdateOne) AddDeletedAt(u int32) *ContractUpdateOne {
	cuo.mutation.AddDeletedAt(u)
	return cuo
}

// SetChainType sets the "chain_type" field.
func (cuo *ContractUpdateOne) SetChainType(s string) *ContractUpdateOne {
	cuo.mutation.SetChainType(s)
	return cuo
}

// SetChainID sets the "chain_id" field.
func (cuo *ContractUpdateOne) SetChainID(s string) *ContractUpdateOne {
	cuo.mutation.SetChainID(s)
	return cuo
}

// SetAddress sets the "address" field.
func (cuo *ContractUpdateOne) SetAddress(s string) *ContractUpdateOne {
	cuo.mutation.SetAddress(s)
	return cuo
}

// SetName sets the "name" field.
func (cuo *ContractUpdateOne) SetName(s string) *ContractUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetSymbol sets the "symbol" field.
func (cuo *ContractUpdateOne) SetSymbol(s string) *ContractUpdateOne {
	cuo.mutation.SetSymbol(s)
	return cuo
}

// SetDecimals sets the "decimals" field.
func (cuo *ContractUpdateOne) SetDecimals(u uint32) *ContractUpdateOne {
	cuo.mutation.ResetDecimals()
	cuo.mutation.SetDecimals(u)
	return cuo
}

// SetNillableDecimals sets the "decimals" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableDecimals(u *uint32) *ContractUpdateOne {
	if u != nil {
		cuo.SetDecimals(*u)
	}
	return cuo
}

// AddDecimals adds u to the "decimals" field.
func (cuo *ContractUpdateOne) AddDecimals(u int32) *ContractUpdateOne {
	cuo.mutation.AddDecimals(u)
	return cuo
}

// SetCreator sets the "creator" field.
func (cuo *ContractUpdateOne) SetCreator(s string) *ContractUpdateOne {
	cuo.mutation.SetCreator(s)
	return cuo
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableCreator(s *string) *ContractUpdateOne {
	if s != nil {
		cuo.SetCreator(*s)
	}
	return cuo
}

// ClearCreator clears the value of the "creator" field.
func (cuo *ContractUpdateOne) ClearCreator() *ContractUpdateOne {
	cuo.mutation.ClearCreator()
	return cuo
}

// SetBlockNum sets the "block_num" field.
func (cuo *ContractUpdateOne) SetBlockNum(u uint64) *ContractUpdateOne {
	cuo.mutation.ResetBlockNum()
	cuo.mutation.SetBlockNum(u)
	return cuo
}

// SetNillableBlockNum sets the "block_num" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableBlockNum(u *uint64) *ContractUpdateOne {
	if u != nil {
		cuo.SetBlockNum(*u)
	}
	return cuo
}

// AddBlockNum adds u to the "block_num" field.
func (cuo *ContractUpdateOne) AddBlockNum(u int64) *ContractUpdateOne {
	cuo.mutation.AddBlockNum(u)
	return cuo
}

// ClearBlockNum clears the value of the "block_num" field.
func (cuo *ContractUpdateOne) ClearBlockNum() *ContractUpdateOne {
	cuo.mutation.ClearBlockNum()
	return cuo
}

// SetTxHash sets the "tx_hash" field.
func (cuo *ContractUpdateOne) SetTxHash(s string) *ContractUpdateOne {
	cuo.mutation.SetTxHash(s)
	return cuo
}

// SetNillableTxHash sets the "tx_hash" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableTxHash(s *string) *ContractUpdateOne {
	if s != nil {
		cuo.SetTxHash(*s)
	}
	return cuo
}

// ClearTxHash clears the value of the "tx_hash" field.
func (cuo *ContractUpdateOne) ClearTxHash() *ContractUpdateOne {
	cuo.mutation.ClearTxHash()
	return cuo
}

// SetTxTime sets the "tx_time" field.
func (cuo *ContractUpdateOne) SetTxTime(u uint32) *ContractUpdateOne {
	cuo.mutation.ResetTxTime()
	cuo.mutation.SetTxTime(u)
	return cuo
}

// SetNillableTxTime sets the "tx_time" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableTxTime(u *uint32) *ContractUpdateOne {
	if u != nil {
		cuo.SetTxTime(*u)
	}
	return cuo
}

// AddTxTime adds u to the "tx_time" field.
func (cuo *ContractUpdateOne) AddTxTime(u int32) *ContractUpdateOne {
	cuo.mutation.AddTxTime(u)
	return cuo
}

// ClearTxTime clears the value of the "tx_time" field.
func (cuo *ContractUpdateOne) ClearTxTime() *ContractUpdateOne {
	cuo.mutation.ClearTxTime()
	return cuo
}

// SetProfileURL sets the "profile_url" field.
func (cuo *ContractUpdateOne) SetProfileURL(s string) *ContractUpdateOne {
	cuo.mutation.SetProfileURL(s)
	return cuo
}

// SetNillableProfileURL sets the "profile_url" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableProfileURL(s *string) *ContractUpdateOne {
	if s != nil {
		cuo.SetProfileURL(*s)
	}
	return cuo
}

// ClearProfileURL clears the value of the "profile_url" field.
func (cuo *ContractUpdateOne) ClearProfileURL() *ContractUpdateOne {
	cuo.mutation.ClearProfileURL()
	return cuo
}

// SetBaseURL sets the "base_url" field.
func (cuo *ContractUpdateOne) SetBaseURL(s string) *ContractUpdateOne {
	cuo.mutation.SetBaseURL(s)
	return cuo
}

// SetNillableBaseURL sets the "base_url" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableBaseURL(s *string) *ContractUpdateOne {
	if s != nil {
		cuo.SetBaseURL(*s)
	}
	return cuo
}

// ClearBaseURL clears the value of the "base_url" field.
func (cuo *ContractUpdateOne) ClearBaseURL() *ContractUpdateOne {
	cuo.mutation.ClearBaseURL()
	return cuo
}

// SetBannerURL sets the "banner_url" field.
func (cuo *ContractUpdateOne) SetBannerURL(s string) *ContractUpdateOne {
	cuo.mutation.SetBannerURL(s)
	return cuo
}

// SetNillableBannerURL sets the "banner_url" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableBannerURL(s *string) *ContractUpdateOne {
	if s != nil {
		cuo.SetBannerURL(*s)
	}
	return cuo
}

// ClearBannerURL clears the value of the "banner_url" field.
func (cuo *ContractUpdateOne) ClearBannerURL() *ContractUpdateOne {
	cuo.mutation.ClearBannerURL()
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *ContractUpdateOne) SetDescription(s string) *ContractUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableDescription(s *string) *ContractUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *ContractUpdateOne) ClearDescription() *ContractUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetRemark sets the "remark" field.
func (cuo *ContractUpdateOne) SetRemark(s string) *ContractUpdateOne {
	cuo.mutation.SetRemark(s)
	return cuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (cuo *ContractUpdateOne) SetNillableRemark(s *string) *ContractUpdateOne {
	if s != nil {
		cuo.SetRemark(*s)
	}
	return cuo
}

// ClearRemark clears the value of the "remark" field.
func (cuo *ContractUpdateOne) ClearRemark() *ContractUpdateOne {
	cuo.mutation.ClearRemark()
	return cuo
}

// Mutation returns the ContractMutation object of the builder.
func (cuo *ContractUpdateOne) Mutation() *ContractMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ContractUpdateOne) Select(field string, fields ...string) *ContractUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Contract entity.
func (cuo *ContractUpdateOne) Save(ctx context.Context) (*Contract, error) {
	var (
		err  error
		node *Contract
	)
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Contract)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ContractMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContractUpdateOne) SaveX(ctx context.Context) *Contract {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ContractUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContractUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ContractUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		if contract.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized contract.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := contract.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *ContractUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ContractUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *ContractUpdateOne) sqlSave(ctx context.Context) (_node *Contract, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   contract.Table,
			Columns: contract.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: contract.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Contract.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contract.FieldID)
		for _, f := range fields {
			if !contract.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contract.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.ChainType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldChainType,
		})
	}
	if value, ok := cuo.mutation.ChainID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldChainID,
		})
	}
	if value, ok := cuo.mutation.Address(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldAddress,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldName,
		})
	}
	if value, ok := cuo.mutation.Symbol(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldSymbol,
		})
	}
	if value, ok := cuo.mutation.Decimals(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldDecimals,
		})
	}
	if value, ok := cuo.mutation.AddedDecimals(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldDecimals,
		})
	}
	if value, ok := cuo.mutation.Creator(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldCreator,
		})
	}
	if cuo.mutation.CreatorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldCreator,
		})
	}
	if value, ok := cuo.mutation.BlockNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: contract.FieldBlockNum,
		})
	}
	if value, ok := cuo.mutation.AddedBlockNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: contract.FieldBlockNum,
		})
	}
	if cuo.mutation.BlockNumCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Column: contract.FieldBlockNum,
		})
	}
	if value, ok := cuo.mutation.TxHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldTxHash,
		})
	}
	if cuo.mutation.TxHashCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldTxHash,
		})
	}
	if value, ok := cuo.mutation.TxTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldTxTime,
		})
	}
	if value, ok := cuo.mutation.AddedTxTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: contract.FieldTxTime,
		})
	}
	if cuo.mutation.TxTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: contract.FieldTxTime,
		})
	}
	if value, ok := cuo.mutation.ProfileURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldProfileURL,
		})
	}
	if cuo.mutation.ProfileURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldProfileURL,
		})
	}
	if value, ok := cuo.mutation.BaseURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldBaseURL,
		})
	}
	if cuo.mutation.BaseURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldBaseURL,
		})
	}
	if value, ok := cuo.mutation.BannerURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldBannerURL,
		})
	}
	if cuo.mutation.BannerURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldBannerURL,
		})
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldDescription,
		})
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldDescription,
		})
	}
	if value, ok := cuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: contract.FieldRemark,
		})
	}
	if cuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: contract.FieldRemark,
		})
	}
	_spec.Modifiers = cuo.modifiers
	_node = &Contract{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contract.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

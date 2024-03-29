// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/block"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/predicate"
)

// BlockUpdate is the builder for updating Block entities.
type BlockUpdate struct {
	config
	hooks     []Hook
	mutation  *BlockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the BlockUpdate builder.
func (bu *BlockUpdate) Where(ps ...predicate.Block) *BlockUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetCreatedAt sets the "created_at" field.
func (bu *BlockUpdate) SetCreatedAt(u uint32) *BlockUpdate {
	bu.mutation.ResetCreatedAt()
	bu.mutation.SetCreatedAt(u)
	return bu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bu *BlockUpdate) SetNillableCreatedAt(u *uint32) *BlockUpdate {
	if u != nil {
		bu.SetCreatedAt(*u)
	}
	return bu
}

// AddCreatedAt adds u to the "created_at" field.
func (bu *BlockUpdate) AddCreatedAt(u int32) *BlockUpdate {
	bu.mutation.AddCreatedAt(u)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BlockUpdate) SetUpdatedAt(u uint32) *BlockUpdate {
	bu.mutation.ResetUpdatedAt()
	bu.mutation.SetUpdatedAt(u)
	return bu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (bu *BlockUpdate) AddUpdatedAt(u int32) *BlockUpdate {
	bu.mutation.AddUpdatedAt(u)
	return bu
}

// SetDeletedAt sets the "deleted_at" field.
func (bu *BlockUpdate) SetDeletedAt(u uint32) *BlockUpdate {
	bu.mutation.ResetDeletedAt()
	bu.mutation.SetDeletedAt(u)
	return bu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bu *BlockUpdate) SetNillableDeletedAt(u *uint32) *BlockUpdate {
	if u != nil {
		bu.SetDeletedAt(*u)
	}
	return bu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (bu *BlockUpdate) AddDeletedAt(u int32) *BlockUpdate {
	bu.mutation.AddDeletedAt(u)
	return bu
}

// SetEntID sets the "ent_id" field.
func (bu *BlockUpdate) SetEntID(u uuid.UUID) *BlockUpdate {
	bu.mutation.SetEntID(u)
	return bu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (bu *BlockUpdate) SetNillableEntID(u *uuid.UUID) *BlockUpdate {
	if u != nil {
		bu.SetEntID(*u)
	}
	return bu
}

// SetChainType sets the "chain_type" field.
func (bu *BlockUpdate) SetChainType(s string) *BlockUpdate {
	bu.mutation.SetChainType(s)
	return bu
}

// SetChainID sets the "chain_id" field.
func (bu *BlockUpdate) SetChainID(s string) *BlockUpdate {
	bu.mutation.SetChainID(s)
	return bu
}

// SetBlockNumber sets the "block_number" field.
func (bu *BlockUpdate) SetBlockNumber(u uint64) *BlockUpdate {
	bu.mutation.ResetBlockNumber()
	bu.mutation.SetBlockNumber(u)
	return bu
}

// AddBlockNumber adds u to the "block_number" field.
func (bu *BlockUpdate) AddBlockNumber(u int64) *BlockUpdate {
	bu.mutation.AddBlockNumber(u)
	return bu
}

// SetBlockHash sets the "block_hash" field.
func (bu *BlockUpdate) SetBlockHash(s string) *BlockUpdate {
	bu.mutation.SetBlockHash(s)
	return bu
}

// SetBlockTime sets the "block_time" field.
func (bu *BlockUpdate) SetBlockTime(u uint64) *BlockUpdate {
	bu.mutation.ResetBlockTime()
	bu.mutation.SetBlockTime(u)
	return bu
}

// AddBlockTime adds u to the "block_time" field.
func (bu *BlockUpdate) AddBlockTime(u int64) *BlockUpdate {
	bu.mutation.AddBlockTime(u)
	return bu
}

// SetParseState sets the "parse_state" field.
func (bu *BlockUpdate) SetParseState(s string) *BlockUpdate {
	bu.mutation.SetParseState(s)
	return bu
}

// SetRemark sets the "remark" field.
func (bu *BlockUpdate) SetRemark(s string) *BlockUpdate {
	bu.mutation.SetRemark(s)
	return bu
}

// Mutation returns the BlockMutation object of the builder.
func (bu *BlockUpdate) Mutation() *BlockMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := bu.defaults(); err != nil {
		return 0, err
	}
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlockUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlockUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlockUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BlockUpdate) defaults() error {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		if block.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized block.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := block.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (bu *BlockUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *BlockUpdate {
	bu.modifiers = append(bu.modifiers, modifiers...)
	return bu
}

func (bu *BlockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   block.Table,
			Columns: block.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: block.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldCreatedAt,
		})
	}
	if value, ok := bu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldCreatedAt,
		})
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldUpdatedAt,
		})
	}
	if value, ok := bu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldUpdatedAt,
		})
	}
	if value, ok := bu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldDeletedAt,
		})
	}
	if value, ok := bu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldDeletedAt,
		})
	}
	if value, ok := bu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: block.FieldEntID,
		})
	}
	if value, ok := bu.mutation.ChainType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldChainType,
		})
	}
	if value, ok := bu.mutation.ChainID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldChainID,
		})
	}
	if value, ok := bu.mutation.BlockNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockNumber,
		})
	}
	if value, ok := bu.mutation.AddedBlockNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockNumber,
		})
	}
	if value, ok := bu.mutation.BlockHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldBlockHash,
		})
	}
	if value, ok := bu.mutation.BlockTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockTime,
		})
	}
	if value, ok := bu.mutation.AddedBlockTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockTime,
		})
	}
	if value, ok := bu.mutation.ParseState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldParseState,
		})
	}
	if value, ok := bu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldRemark,
		})
	}
	_spec.Modifiers = bu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{block.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// BlockUpdateOne is the builder for updating a single Block entity.
type BlockUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *BlockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (buo *BlockUpdateOne) SetCreatedAt(u uint32) *BlockUpdateOne {
	buo.mutation.ResetCreatedAt()
	buo.mutation.SetCreatedAt(u)
	return buo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (buo *BlockUpdateOne) SetNillableCreatedAt(u *uint32) *BlockUpdateOne {
	if u != nil {
		buo.SetCreatedAt(*u)
	}
	return buo
}

// AddCreatedAt adds u to the "created_at" field.
func (buo *BlockUpdateOne) AddCreatedAt(u int32) *BlockUpdateOne {
	buo.mutation.AddCreatedAt(u)
	return buo
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BlockUpdateOne) SetUpdatedAt(u uint32) *BlockUpdateOne {
	buo.mutation.ResetUpdatedAt()
	buo.mutation.SetUpdatedAt(u)
	return buo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (buo *BlockUpdateOne) AddUpdatedAt(u int32) *BlockUpdateOne {
	buo.mutation.AddUpdatedAt(u)
	return buo
}

// SetDeletedAt sets the "deleted_at" field.
func (buo *BlockUpdateOne) SetDeletedAt(u uint32) *BlockUpdateOne {
	buo.mutation.ResetDeletedAt()
	buo.mutation.SetDeletedAt(u)
	return buo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (buo *BlockUpdateOne) SetNillableDeletedAt(u *uint32) *BlockUpdateOne {
	if u != nil {
		buo.SetDeletedAt(*u)
	}
	return buo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (buo *BlockUpdateOne) AddDeletedAt(u int32) *BlockUpdateOne {
	buo.mutation.AddDeletedAt(u)
	return buo
}

// SetEntID sets the "ent_id" field.
func (buo *BlockUpdateOne) SetEntID(u uuid.UUID) *BlockUpdateOne {
	buo.mutation.SetEntID(u)
	return buo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (buo *BlockUpdateOne) SetNillableEntID(u *uuid.UUID) *BlockUpdateOne {
	if u != nil {
		buo.SetEntID(*u)
	}
	return buo
}

// SetChainType sets the "chain_type" field.
func (buo *BlockUpdateOne) SetChainType(s string) *BlockUpdateOne {
	buo.mutation.SetChainType(s)
	return buo
}

// SetChainID sets the "chain_id" field.
func (buo *BlockUpdateOne) SetChainID(s string) *BlockUpdateOne {
	buo.mutation.SetChainID(s)
	return buo
}

// SetBlockNumber sets the "block_number" field.
func (buo *BlockUpdateOne) SetBlockNumber(u uint64) *BlockUpdateOne {
	buo.mutation.ResetBlockNumber()
	buo.mutation.SetBlockNumber(u)
	return buo
}

// AddBlockNumber adds u to the "block_number" field.
func (buo *BlockUpdateOne) AddBlockNumber(u int64) *BlockUpdateOne {
	buo.mutation.AddBlockNumber(u)
	return buo
}

// SetBlockHash sets the "block_hash" field.
func (buo *BlockUpdateOne) SetBlockHash(s string) *BlockUpdateOne {
	buo.mutation.SetBlockHash(s)
	return buo
}

// SetBlockTime sets the "block_time" field.
func (buo *BlockUpdateOne) SetBlockTime(u uint64) *BlockUpdateOne {
	buo.mutation.ResetBlockTime()
	buo.mutation.SetBlockTime(u)
	return buo
}

// AddBlockTime adds u to the "block_time" field.
func (buo *BlockUpdateOne) AddBlockTime(u int64) *BlockUpdateOne {
	buo.mutation.AddBlockTime(u)
	return buo
}

// SetParseState sets the "parse_state" field.
func (buo *BlockUpdateOne) SetParseState(s string) *BlockUpdateOne {
	buo.mutation.SetParseState(s)
	return buo
}

// SetRemark sets the "remark" field.
func (buo *BlockUpdateOne) SetRemark(s string) *BlockUpdateOne {
	buo.mutation.SetRemark(s)
	return buo
}

// Mutation returns the BlockMutation object of the builder.
func (buo *BlockUpdateOne) Mutation() *BlockMutation {
	return buo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlockUpdateOne) Select(field string, fields ...string) *BlockUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Block entity.
func (buo *BlockUpdateOne) Save(ctx context.Context) (*Block, error) {
	var (
		err  error
		node *Block
	)
	if err := buo.defaults(); err != nil {
		return nil, err
	}
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, buo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Block)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BlockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlockUpdateOne) SaveX(ctx context.Context) *Block {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlockUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlockUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BlockUpdateOne) defaults() error {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		if block.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized block.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := block.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (buo *BlockUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *BlockUpdateOne {
	buo.modifiers = append(buo.modifiers, modifiers...)
	return buo
}

func (buo *BlockUpdateOne) sqlSave(ctx context.Context) (_node *Block, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   block.Table,
			Columns: block.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: block.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Block.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, block.FieldID)
		for _, f := range fields {
			if !block.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != block.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldCreatedAt,
		})
	}
	if value, ok := buo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldCreatedAt,
		})
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldUpdatedAt,
		})
	}
	if value, ok := buo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldUpdatedAt,
		})
	}
	if value, ok := buo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldDeletedAt,
		})
	}
	if value, ok := buo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: block.FieldDeletedAt,
		})
	}
	if value, ok := buo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: block.FieldEntID,
		})
	}
	if value, ok := buo.mutation.ChainType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldChainType,
		})
	}
	if value, ok := buo.mutation.ChainID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldChainID,
		})
	}
	if value, ok := buo.mutation.BlockNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockNumber,
		})
	}
	if value, ok := buo.mutation.AddedBlockNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockNumber,
		})
	}
	if value, ok := buo.mutation.BlockHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldBlockHash,
		})
	}
	if value, ok := buo.mutation.BlockTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockTime,
		})
	}
	if value, ok := buo.mutation.AddedBlockTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockTime,
		})
	}
	if value, ok := buo.mutation.ParseState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldParseState,
		})
	}
	if value, ok := buo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldRemark,
		})
	}
	_spec.Modifiers = buo.modifiers
	_node = &Block{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{block.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

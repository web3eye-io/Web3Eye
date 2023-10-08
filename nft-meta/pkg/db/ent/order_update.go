// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/order"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/predicate"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetCreatedAt sets the "created_at" field.
func (ou *OrderUpdate) SetCreatedAt(u uint32) *OrderUpdate {
	ou.mutation.ResetCreatedAt()
	ou.mutation.SetCreatedAt(u)
	return ou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableCreatedAt(u *uint32) *OrderUpdate {
	if u != nil {
		ou.SetCreatedAt(*u)
	}
	return ou
}

// AddCreatedAt adds u to the "created_at" field.
func (ou *OrderUpdate) AddCreatedAt(u int32) *OrderUpdate {
	ou.mutation.AddCreatedAt(u)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrderUpdate) SetUpdatedAt(u uint32) *OrderUpdate {
	ou.mutation.ResetUpdatedAt()
	ou.mutation.SetUpdatedAt(u)
	return ou
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ou *OrderUpdate) AddUpdatedAt(u int32) *OrderUpdate {
	ou.mutation.AddUpdatedAt(u)
	return ou
}

// SetDeletedAt sets the "deleted_at" field.
func (ou *OrderUpdate) SetDeletedAt(u uint32) *OrderUpdate {
	ou.mutation.ResetDeletedAt()
	ou.mutation.SetDeletedAt(u)
	return ou
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDeletedAt(u *uint32) *OrderUpdate {
	if u != nil {
		ou.SetDeletedAt(*u)
	}
	return ou
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ou *OrderUpdate) AddDeletedAt(u int32) *OrderUpdate {
	ou.mutation.AddDeletedAt(u)
	return ou
}

// SetChainType sets the "chain_type" field.
func (ou *OrderUpdate) SetChainType(s string) *OrderUpdate {
	ou.mutation.SetChainType(s)
	return ou
}

// SetChainID sets the "chain_id" field.
func (ou *OrderUpdate) SetChainID(s string) *OrderUpdate {
	ou.mutation.SetChainID(s)
	return ou
}

// SetTxHash sets the "tx_hash" field.
func (ou *OrderUpdate) SetTxHash(s string) *OrderUpdate {
	ou.mutation.SetTxHash(s)
	return ou
}

// SetBlockNumber sets the "block_number" field.
func (ou *OrderUpdate) SetBlockNumber(u uint64) *OrderUpdate {
	ou.mutation.ResetBlockNumber()
	ou.mutation.SetBlockNumber(u)
	return ou
}

// AddBlockNumber adds u to the "block_number" field.
func (ou *OrderUpdate) AddBlockNumber(u int64) *OrderUpdate {
	ou.mutation.AddBlockNumber(u)
	return ou
}

// SetTxIndex sets the "tx_index" field.
func (ou *OrderUpdate) SetTxIndex(u uint32) *OrderUpdate {
	ou.mutation.ResetTxIndex()
	ou.mutation.SetTxIndex(u)
	return ou
}

// AddTxIndex adds u to the "tx_index" field.
func (ou *OrderUpdate) AddTxIndex(u int32) *OrderUpdate {
	ou.mutation.AddTxIndex(u)
	return ou
}

// SetLogIndex sets the "log_index" field.
func (ou *OrderUpdate) SetLogIndex(u uint32) *OrderUpdate {
	ou.mutation.ResetLogIndex()
	ou.mutation.SetLogIndex(u)
	return ou
}

// AddLogIndex adds u to the "log_index" field.
func (ou *OrderUpdate) AddLogIndex(u int32) *OrderUpdate {
	ou.mutation.AddLogIndex(u)
	return ou
}

// SetRecipient sets the "recipient" field.
func (ou *OrderUpdate) SetRecipient(s string) *OrderUpdate {
	ou.mutation.SetRecipient(s)
	return ou
}

// SetRemark sets the "remark" field.
func (ou *OrderUpdate) SetRemark(s string) *OrderUpdate {
	ou.mutation.SetRemark(s)
	return ou
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableRemark(s *string) *OrderUpdate {
	if s != nil {
		ou.SetRemark(*s)
	}
	return ou
}

// ClearRemark clears the value of the "remark" field.
func (ou *OrderUpdate) ClearRemark() *OrderUpdate {
	ou.mutation.ClearRemark()
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ou.defaults(); err != nil {
		return 0, err
	}
	if len(ou.hooks) == 0 {
		if err = ou.check(); err != nil {
			return 0, err
		}
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ou.check(); err != nil {
				return 0, err
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrderUpdate) defaults() error {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		if order.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized order.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := order.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.TxHash(); ok {
		if err := order.TxHashValidator(v); err != nil {
			return &ValidationError{Name: "tx_hash", err: fmt.Errorf(`ent: validator failed for field "Order.tx_hash": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Recipient(); ok {
		if err := order.RecipientValidator(v); err != nil {
			return &ValidationError{Name: "recipient", err: fmt.Errorf(`ent: validator failed for field "Order.recipient": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ou *OrderUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderUpdate {
	ou.modifiers = append(ou.modifiers, modifiers...)
	return ou
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ou.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ou.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ou.mutation.ChainType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldChainType,
		})
	}
	if value, ok := ou.mutation.ChainID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldChainID,
		})
	}
	if value, ok := ou.mutation.TxHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldTxHash,
		})
	}
	if value, ok := ou.mutation.BlockNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: order.FieldBlockNumber,
		})
	}
	if value, ok := ou.mutation.AddedBlockNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: order.FieldBlockNumber,
		})
	}
	if value, ok := ou.mutation.TxIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldTxIndex,
		})
	}
	if value, ok := ou.mutation.AddedTxIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldTxIndex,
		})
	}
	if value, ok := ou.mutation.LogIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldLogIndex,
		})
	}
	if value, ok := ou.mutation.AddedLogIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldLogIndex,
		})
	}
	if value, ok := ou.mutation.Recipient(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldRecipient,
		})
	}
	if value, ok := ou.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldRemark,
		})
	}
	if ou.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldRemark,
		})
	}
	_spec.Modifiers = ou.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ouo *OrderUpdateOne) SetCreatedAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetCreatedAt()
	ouo.mutation.SetCreatedAt(u)
	return ouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCreatedAt(u *uint32) *OrderUpdateOne {
	if u != nil {
		ouo.SetCreatedAt(*u)
	}
	return ouo
}

// AddCreatedAt adds u to the "created_at" field.
func (ouo *OrderUpdateOne) AddCreatedAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddCreatedAt(u)
	return ouo
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrderUpdateOne) SetUpdatedAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetUpdatedAt()
	ouo.mutation.SetUpdatedAt(u)
	return ouo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ouo *OrderUpdateOne) AddUpdatedAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddUpdatedAt(u)
	return ouo
}

// SetDeletedAt sets the "deleted_at" field.
func (ouo *OrderUpdateOne) SetDeletedAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetDeletedAt()
	ouo.mutation.SetDeletedAt(u)
	return ouo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDeletedAt(u *uint32) *OrderUpdateOne {
	if u != nil {
		ouo.SetDeletedAt(*u)
	}
	return ouo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ouo *OrderUpdateOne) AddDeletedAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddDeletedAt(u)
	return ouo
}

// SetChainType sets the "chain_type" field.
func (ouo *OrderUpdateOne) SetChainType(s string) *OrderUpdateOne {
	ouo.mutation.SetChainType(s)
	return ouo
}

// SetChainID sets the "chain_id" field.
func (ouo *OrderUpdateOne) SetChainID(s string) *OrderUpdateOne {
	ouo.mutation.SetChainID(s)
	return ouo
}

// SetTxHash sets the "tx_hash" field.
func (ouo *OrderUpdateOne) SetTxHash(s string) *OrderUpdateOne {
	ouo.mutation.SetTxHash(s)
	return ouo
}

// SetBlockNumber sets the "block_number" field.
func (ouo *OrderUpdateOne) SetBlockNumber(u uint64) *OrderUpdateOne {
	ouo.mutation.ResetBlockNumber()
	ouo.mutation.SetBlockNumber(u)
	return ouo
}

// AddBlockNumber adds u to the "block_number" field.
func (ouo *OrderUpdateOne) AddBlockNumber(u int64) *OrderUpdateOne {
	ouo.mutation.AddBlockNumber(u)
	return ouo
}

// SetTxIndex sets the "tx_index" field.
func (ouo *OrderUpdateOne) SetTxIndex(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetTxIndex()
	ouo.mutation.SetTxIndex(u)
	return ouo
}

// AddTxIndex adds u to the "tx_index" field.
func (ouo *OrderUpdateOne) AddTxIndex(u int32) *OrderUpdateOne {
	ouo.mutation.AddTxIndex(u)
	return ouo
}

// SetLogIndex sets the "log_index" field.
func (ouo *OrderUpdateOne) SetLogIndex(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetLogIndex()
	ouo.mutation.SetLogIndex(u)
	return ouo
}

// AddLogIndex adds u to the "log_index" field.
func (ouo *OrderUpdateOne) AddLogIndex(u int32) *OrderUpdateOne {
	ouo.mutation.AddLogIndex(u)
	return ouo
}

// SetRecipient sets the "recipient" field.
func (ouo *OrderUpdateOne) SetRecipient(s string) *OrderUpdateOne {
	ouo.mutation.SetRecipient(s)
	return ouo
}

// SetRemark sets the "remark" field.
func (ouo *OrderUpdateOne) SetRemark(s string) *OrderUpdateOne {
	ouo.mutation.SetRemark(s)
	return ouo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableRemark(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetRemark(*s)
	}
	return ouo
}

// ClearRemark clears the value of the "remark" field.
func (ouo *OrderUpdateOne) ClearRemark() *OrderUpdateOne {
	ouo.mutation.ClearRemark()
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	if err := ouo.defaults(); err != nil {
		return nil, err
	}
	if len(ouo.hooks) == 0 {
		if err = ouo.check(); err != nil {
			return nil, err
		}
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouo.check(); err != nil {
				return nil, err
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ouo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Order)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrderUpdateOne) defaults() error {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		if order.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized order.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := order.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.TxHash(); ok {
		if err := order.TxHashValidator(v); err != nil {
			return &ValidationError{Name: "tx_hash", err: fmt.Errorf(`ent: validator failed for field "Order.tx_hash": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Recipient(); ok {
		if err := order.RecipientValidator(v); err != nil {
			return &ValidationError{Name: "recipient", err: fmt.Errorf(`ent: validator failed for field "Order.recipient": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ouo *OrderUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderUpdateOne {
	ouo.modifiers = append(ouo.modifiers, modifiers...)
	return ouo
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ouo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ouo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ouo.mutation.ChainType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldChainType,
		})
	}
	if value, ok := ouo.mutation.ChainID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldChainID,
		})
	}
	if value, ok := ouo.mutation.TxHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldTxHash,
		})
	}
	if value, ok := ouo.mutation.BlockNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: order.FieldBlockNumber,
		})
	}
	if value, ok := ouo.mutation.AddedBlockNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: order.FieldBlockNumber,
		})
	}
	if value, ok := ouo.mutation.TxIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldTxIndex,
		})
	}
	if value, ok := ouo.mutation.AddedTxIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldTxIndex,
		})
	}
	if value, ok := ouo.mutation.LogIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldLogIndex,
		})
	}
	if value, ok := ouo.mutation.AddedLogIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldLogIndex,
		})
	}
	if value, ok := ouo.mutation.Recipient(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldRecipient,
		})
	}
	if value, ok := ouo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldRemark,
		})
	}
	if ouo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldRemark,
		})
	}
	_spec.Modifiers = ouo.modifiers
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

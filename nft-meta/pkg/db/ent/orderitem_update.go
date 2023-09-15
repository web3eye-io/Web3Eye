// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderitem"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/predicate"
)

// OrderItemUpdate is the builder for updating OrderItem entities.
type OrderItemUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiu *OrderItemUpdate) Where(ps ...predicate.OrderItem) *OrderItemUpdate {
	oiu.mutation.Where(ps...)
	return oiu
}

// SetCreatedAt sets the "created_at" field.
func (oiu *OrderItemUpdate) SetCreatedAt(u uint32) *OrderItemUpdate {
	oiu.mutation.ResetCreatedAt()
	oiu.mutation.SetCreatedAt(u)
	return oiu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableCreatedAt(u *uint32) *OrderItemUpdate {
	if u != nil {
		oiu.SetCreatedAt(*u)
	}
	return oiu
}

// AddCreatedAt adds u to the "created_at" field.
func (oiu *OrderItemUpdate) AddCreatedAt(u int32) *OrderItemUpdate {
	oiu.mutation.AddCreatedAt(u)
	return oiu
}

// SetUpdatedAt sets the "updated_at" field.
func (oiu *OrderItemUpdate) SetUpdatedAt(u uint32) *OrderItemUpdate {
	oiu.mutation.ResetUpdatedAt()
	oiu.mutation.SetUpdatedAt(u)
	return oiu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (oiu *OrderItemUpdate) AddUpdatedAt(u int32) *OrderItemUpdate {
	oiu.mutation.AddUpdatedAt(u)
	return oiu
}

// SetDeletedAt sets the "deleted_at" field.
func (oiu *OrderItemUpdate) SetDeletedAt(u uint32) *OrderItemUpdate {
	oiu.mutation.ResetDeletedAt()
	oiu.mutation.SetDeletedAt(u)
	return oiu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableDeletedAt(u *uint32) *OrderItemUpdate {
	if u != nil {
		oiu.SetDeletedAt(*u)
	}
	return oiu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (oiu *OrderItemUpdate) AddDeletedAt(u int32) *OrderItemUpdate {
	oiu.mutation.AddDeletedAt(u)
	return oiu
}

// SetOrderID sets the "order_id" field.
func (oiu *OrderItemUpdate) SetOrderID(s string) *OrderItemUpdate {
	oiu.mutation.SetOrderID(s)
	return oiu
}

// SetOrderItemType sets the "order_item_type" field.
func (oiu *OrderItemUpdate) SetOrderItemType(s string) *OrderItemUpdate {
	oiu.mutation.SetOrderItemType(s)
	return oiu
}

// SetContract sets the "contract" field.
func (oiu *OrderItemUpdate) SetContract(s string) *OrderItemUpdate {
	oiu.mutation.SetContract(s)
	return oiu
}

// SetTokenType sets the "token_type" field.
func (oiu *OrderItemUpdate) SetTokenType(s string) *OrderItemUpdate {
	oiu.mutation.SetTokenType(s)
	return oiu
}

// SetTokenID sets the "token_id" field.
func (oiu *OrderItemUpdate) SetTokenID(s string) *OrderItemUpdate {
	oiu.mutation.SetTokenID(s)
	return oiu
}

// SetAmount sets the "amount" field.
func (oiu *OrderItemUpdate) SetAmount(u uint64) *OrderItemUpdate {
	oiu.mutation.ResetAmount()
	oiu.mutation.SetAmount(u)
	return oiu
}

// AddAmount adds u to the "amount" field.
func (oiu *OrderItemUpdate) AddAmount(u int64) *OrderItemUpdate {
	oiu.mutation.AddAmount(u)
	return oiu
}

// SetRemark sets the "remark" field.
func (oiu *OrderItemUpdate) SetRemark(s string) *OrderItemUpdate {
	oiu.mutation.SetRemark(s)
	return oiu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableRemark(s *string) *OrderItemUpdate {
	if s != nil {
		oiu.SetRemark(*s)
	}
	return oiu
}

// ClearRemark clears the value of the "remark" field.
func (oiu *OrderItemUpdate) ClearRemark() *OrderItemUpdate {
	oiu.mutation.ClearRemark()
	return oiu
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiu *OrderItemUpdate) Mutation() *OrderItemMutation {
	return oiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oiu *OrderItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := oiu.defaults(); err != nil {
		return 0, err
	}
	if len(oiu.hooks) == 0 {
		affected, err = oiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oiu.mutation = mutation
			affected, err = oiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oiu.hooks) - 1; i >= 0; i-- {
			if oiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oiu *OrderItemUpdate) SaveX(ctx context.Context) int {
	affected, err := oiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oiu *OrderItemUpdate) Exec(ctx context.Context) error {
	_, err := oiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiu *OrderItemUpdate) ExecX(ctx context.Context) {
	if err := oiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oiu *OrderItemUpdate) defaults() error {
	if _, ok := oiu.mutation.UpdatedAt(); !ok {
		if orderitem.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderitem.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderitem.UpdateDefaultUpdatedAt()
		oiu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (oiu *OrderItemUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderItemUpdate {
	oiu.modifiers = append(oiu.modifiers, modifiers...)
	return oiu
}

func (oiu *OrderItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderitem.Table,
			Columns: orderitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: orderitem.FieldID,
			},
		},
	}
	if ps := oiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldCreatedAt,
		})
	}
	if value, ok := oiu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldCreatedAt,
		})
	}
	if value, ok := oiu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldUpdatedAt,
		})
	}
	if value, ok := oiu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldUpdatedAt,
		})
	}
	if value, ok := oiu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldDeletedAt,
		})
	}
	if value, ok := oiu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldDeletedAt,
		})
	}
	if value, ok := oiu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldOrderID,
		})
	}
	if value, ok := oiu.mutation.OrderItemType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldOrderItemType,
		})
	}
	if value, ok := oiu.mutation.Contract(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldContract,
		})
	}
	if value, ok := oiu.mutation.TokenType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldTokenType,
		})
	}
	if value, ok := oiu.mutation.TokenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldTokenID,
		})
	}
	if value, ok := oiu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderitem.FieldAmount,
		})
	}
	if value, ok := oiu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderitem.FieldAmount,
		})
	}
	if value, ok := oiu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldRemark,
		})
	}
	if oiu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderitem.FieldRemark,
		})
	}
	_spec.Modifiers = oiu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, oiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderItemUpdateOne is the builder for updating a single OrderItem entity.
type OrderItemUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (oiuo *OrderItemUpdateOne) SetCreatedAt(u uint32) *OrderItemUpdateOne {
	oiuo.mutation.ResetCreatedAt()
	oiuo.mutation.SetCreatedAt(u)
	return oiuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableCreatedAt(u *uint32) *OrderItemUpdateOne {
	if u != nil {
		oiuo.SetCreatedAt(*u)
	}
	return oiuo
}

// AddCreatedAt adds u to the "created_at" field.
func (oiuo *OrderItemUpdateOne) AddCreatedAt(u int32) *OrderItemUpdateOne {
	oiuo.mutation.AddCreatedAt(u)
	return oiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (oiuo *OrderItemUpdateOne) SetUpdatedAt(u uint32) *OrderItemUpdateOne {
	oiuo.mutation.ResetUpdatedAt()
	oiuo.mutation.SetUpdatedAt(u)
	return oiuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (oiuo *OrderItemUpdateOne) AddUpdatedAt(u int32) *OrderItemUpdateOne {
	oiuo.mutation.AddUpdatedAt(u)
	return oiuo
}

// SetDeletedAt sets the "deleted_at" field.
func (oiuo *OrderItemUpdateOne) SetDeletedAt(u uint32) *OrderItemUpdateOne {
	oiuo.mutation.ResetDeletedAt()
	oiuo.mutation.SetDeletedAt(u)
	return oiuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableDeletedAt(u *uint32) *OrderItemUpdateOne {
	if u != nil {
		oiuo.SetDeletedAt(*u)
	}
	return oiuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (oiuo *OrderItemUpdateOne) AddDeletedAt(u int32) *OrderItemUpdateOne {
	oiuo.mutation.AddDeletedAt(u)
	return oiuo
}

// SetOrderID sets the "order_id" field.
func (oiuo *OrderItemUpdateOne) SetOrderID(s string) *OrderItemUpdateOne {
	oiuo.mutation.SetOrderID(s)
	return oiuo
}

// SetOrderItemType sets the "order_item_type" field.
func (oiuo *OrderItemUpdateOne) SetOrderItemType(s string) *OrderItemUpdateOne {
	oiuo.mutation.SetOrderItemType(s)
	return oiuo
}

// SetContract sets the "contract" field.
func (oiuo *OrderItemUpdateOne) SetContract(s string) *OrderItemUpdateOne {
	oiuo.mutation.SetContract(s)
	return oiuo
}

// SetTokenType sets the "token_type" field.
func (oiuo *OrderItemUpdateOne) SetTokenType(s string) *OrderItemUpdateOne {
	oiuo.mutation.SetTokenType(s)
	return oiuo
}

// SetTokenID sets the "token_id" field.
func (oiuo *OrderItemUpdateOne) SetTokenID(s string) *OrderItemUpdateOne {
	oiuo.mutation.SetTokenID(s)
	return oiuo
}

// SetAmount sets the "amount" field.
func (oiuo *OrderItemUpdateOne) SetAmount(u uint64) *OrderItemUpdateOne {
	oiuo.mutation.ResetAmount()
	oiuo.mutation.SetAmount(u)
	return oiuo
}

// AddAmount adds u to the "amount" field.
func (oiuo *OrderItemUpdateOne) AddAmount(u int64) *OrderItemUpdateOne {
	oiuo.mutation.AddAmount(u)
	return oiuo
}

// SetRemark sets the "remark" field.
func (oiuo *OrderItemUpdateOne) SetRemark(s string) *OrderItemUpdateOne {
	oiuo.mutation.SetRemark(s)
	return oiuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableRemark(s *string) *OrderItemUpdateOne {
	if s != nil {
		oiuo.SetRemark(*s)
	}
	return oiuo
}

// ClearRemark clears the value of the "remark" field.
func (oiuo *OrderItemUpdateOne) ClearRemark() *OrderItemUpdateOne {
	oiuo.mutation.ClearRemark()
	return oiuo
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiuo *OrderItemUpdateOne) Mutation() *OrderItemMutation {
	return oiuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oiuo *OrderItemUpdateOne) Select(field string, fields ...string) *OrderItemUpdateOne {
	oiuo.fields = append([]string{field}, fields...)
	return oiuo
}

// Save executes the query and returns the updated OrderItem entity.
func (oiuo *OrderItemUpdateOne) Save(ctx context.Context) (*OrderItem, error) {
	var (
		err  error
		node *OrderItem
	)
	if err := oiuo.defaults(); err != nil {
		return nil, err
	}
	if len(oiuo.hooks) == 0 {
		node, err = oiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oiuo.mutation = mutation
			node, err = oiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oiuo.hooks) - 1; i >= 0; i-- {
			if oiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oiuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oiuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderItem)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderItemMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) SaveX(ctx context.Context) *OrderItem {
	node, err := oiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oiuo *OrderItemUpdateOne) Exec(ctx context.Context) error {
	_, err := oiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) ExecX(ctx context.Context) {
	if err := oiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oiuo *OrderItemUpdateOne) defaults() error {
	if _, ok := oiuo.mutation.UpdatedAt(); !ok {
		if orderitem.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderitem.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderitem.UpdateDefaultUpdatedAt()
		oiuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (oiuo *OrderItemUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderItemUpdateOne {
	oiuo.modifiers = append(oiuo.modifiers, modifiers...)
	return oiuo
}

func (oiuo *OrderItemUpdateOne) sqlSave(ctx context.Context) (_node *OrderItem, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderitem.Table,
			Columns: orderitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: orderitem.FieldID,
			},
		},
	}
	id, ok := oiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderitem.FieldID)
		for _, f := range fields {
			if !orderitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldCreatedAt,
		})
	}
	if value, ok := oiuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldCreatedAt,
		})
	}
	if value, ok := oiuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldUpdatedAt,
		})
	}
	if value, ok := oiuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldUpdatedAt,
		})
	}
	if value, ok := oiuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldDeletedAt,
		})
	}
	if value, ok := oiuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldDeletedAt,
		})
	}
	if value, ok := oiuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldOrderID,
		})
	}
	if value, ok := oiuo.mutation.OrderItemType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldOrderItemType,
		})
	}
	if value, ok := oiuo.mutation.Contract(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldContract,
		})
	}
	if value, ok := oiuo.mutation.TokenType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldTokenType,
		})
	}
	if value, ok := oiuo.mutation.TokenID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldTokenID,
		})
	}
	if value, ok := oiuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderitem.FieldAmount,
		})
	}
	if value, ok := oiuo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: orderitem.FieldAmount,
		})
	}
	if value, ok := oiuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldRemark,
		})
	}
	if oiuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderitem.FieldRemark,
		})
	}
	_spec.Modifiers = oiuo.modifiers
	_node = &OrderItem{config: oiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

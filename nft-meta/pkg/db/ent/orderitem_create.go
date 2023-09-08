// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderitem"
)

// OrderItemCreate is the builder for creating a OrderItem entity.
type OrderItemCreate struct {
	config
	mutation *OrderItemMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (oic *OrderItemCreate) SetCreatedAt(u uint32) *OrderItemCreate {
	oic.mutation.SetCreatedAt(u)
	return oic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableCreatedAt(u *uint32) *OrderItemCreate {
	if u != nil {
		oic.SetCreatedAt(*u)
	}
	return oic
}

// SetUpdatedAt sets the "updated_at" field.
func (oic *OrderItemCreate) SetUpdatedAt(u uint32) *OrderItemCreate {
	oic.mutation.SetUpdatedAt(u)
	return oic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableUpdatedAt(u *uint32) *OrderItemCreate {
	if u != nil {
		oic.SetUpdatedAt(*u)
	}
	return oic
}

// SetDeletedAt sets the "deleted_at" field.
func (oic *OrderItemCreate) SetDeletedAt(u uint32) *OrderItemCreate {
	oic.mutation.SetDeletedAt(u)
	return oic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableDeletedAt(u *uint32) *OrderItemCreate {
	if u != nil {
		oic.SetDeletedAt(*u)
	}
	return oic
}

// SetOrderID sets the "order_id" field.
func (oic *OrderItemCreate) SetOrderID(s string) *OrderItemCreate {
	oic.mutation.SetOrderID(s)
	return oic
}

// SetOrderItemType sets the "order_item_type" field.
func (oic *OrderItemCreate) SetOrderItemType(s string) *OrderItemCreate {
	oic.mutation.SetOrderItemType(s)
	return oic
}

// SetContract sets the "contract" field.
func (oic *OrderItemCreate) SetContract(s string) *OrderItemCreate {
	oic.mutation.SetContract(s)
	return oic
}

// SetTokenType sets the "token_type" field.
func (oic *OrderItemCreate) SetTokenType(s string) *OrderItemCreate {
	oic.mutation.SetTokenType(s)
	return oic
}

// SetTokenID sets the "token_id" field.
func (oic *OrderItemCreate) SetTokenID(s string) *OrderItemCreate {
	oic.mutation.SetTokenID(s)
	return oic
}

// SetAmount sets the "amount" field.
func (oic *OrderItemCreate) SetAmount(i int64) *OrderItemCreate {
	oic.mutation.SetAmount(i)
	return oic
}

// SetRemark sets the "remark" field.
func (oic *OrderItemCreate) SetRemark(s string) *OrderItemCreate {
	oic.mutation.SetRemark(s)
	return oic
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableRemark(s *string) *OrderItemCreate {
	if s != nil {
		oic.SetRemark(*s)
	}
	return oic
}

// SetID sets the "id" field.
func (oic *OrderItemCreate) SetID(u uuid.UUID) *OrderItemCreate {
	oic.mutation.SetID(u)
	return oic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableID(u *uuid.UUID) *OrderItemCreate {
	if u != nil {
		oic.SetID(*u)
	}
	return oic
}

// Mutation returns the OrderItemMutation object of the builder.
func (oic *OrderItemCreate) Mutation() *OrderItemMutation {
	return oic.mutation
}

// Save creates the OrderItem in the database.
func (oic *OrderItemCreate) Save(ctx context.Context) (*OrderItem, error) {
	var (
		err  error
		node *OrderItem
	)
	if err := oic.defaults(); err != nil {
		return nil, err
	}
	if len(oic.hooks) == 0 {
		if err = oic.check(); err != nil {
			return nil, err
		}
		node, err = oic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oic.check(); err != nil {
				return nil, err
			}
			oic.mutation = mutation
			if node, err = oic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oic.hooks) - 1; i >= 0; i-- {
			if oic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oic.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (oic *OrderItemCreate) SaveX(ctx context.Context) *OrderItem {
	v, err := oic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oic *OrderItemCreate) Exec(ctx context.Context) error {
	_, err := oic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oic *OrderItemCreate) ExecX(ctx context.Context) {
	if err := oic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oic *OrderItemCreate) defaults() error {
	if _, ok := oic.mutation.CreatedAt(); !ok {
		if orderitem.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderitem.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := orderitem.DefaultCreatedAt()
		oic.mutation.SetCreatedAt(v)
	}
	if _, ok := oic.mutation.UpdatedAt(); !ok {
		if orderitem.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderitem.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderitem.DefaultUpdatedAt()
		oic.mutation.SetUpdatedAt(v)
	}
	if _, ok := oic.mutation.DeletedAt(); !ok {
		if orderitem.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized orderitem.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := orderitem.DefaultDeletedAt()
		oic.mutation.SetDeletedAt(v)
	}
	if _, ok := oic.mutation.ID(); !ok {
		if orderitem.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized orderitem.DefaultID (forgotten import ent/runtime?)")
		}
		v := orderitem.DefaultID()
		oic.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (oic *OrderItemCreate) check() error {
	if _, ok := oic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderItem.created_at"`)}
	}
	if _, ok := oic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderItem.updated_at"`)}
	}
	if _, ok := oic.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "OrderItem.deleted_at"`)}
	}
	if _, ok := oic.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "OrderItem.order_id"`)}
	}
	if _, ok := oic.mutation.OrderItemType(); !ok {
		return &ValidationError{Name: "order_item_type", err: errors.New(`ent: missing required field "OrderItem.order_item_type"`)}
	}
	if _, ok := oic.mutation.Contract(); !ok {
		return &ValidationError{Name: "contract", err: errors.New(`ent: missing required field "OrderItem.contract"`)}
	}
	if _, ok := oic.mutation.TokenType(); !ok {
		return &ValidationError{Name: "token_type", err: errors.New(`ent: missing required field "OrderItem.token_type"`)}
	}
	if _, ok := oic.mutation.TokenID(); !ok {
		return &ValidationError{Name: "token_id", err: errors.New(`ent: missing required field "OrderItem.token_id"`)}
	}
	if _, ok := oic.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "OrderItem.amount"`)}
	}
	return nil
}

func (oic *OrderItemCreate) sqlSave(ctx context.Context) (*OrderItem, error) {
	_node, _spec := oic.createSpec()
	if err := sqlgraph.CreateNode(ctx, oic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (oic *OrderItemCreate) createSpec() (*OrderItem, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderItem{config: oic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: orderitem.FieldID,
			},
		}
	)
	_spec.OnConflict = oic.conflict
	if id, ok := oic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := oic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := oic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := oic.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderitem.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := oic.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldOrderID,
		})
		_node.OrderID = value
	}
	if value, ok := oic.mutation.OrderItemType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldOrderItemType,
		})
		_node.OrderItemType = value
	}
	if value, ok := oic.mutation.Contract(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldContract,
		})
		_node.Contract = value
	}
	if value, ok := oic.mutation.TokenType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldTokenType,
		})
		_node.TokenType = value
	}
	if value, ok := oic.mutation.TokenID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldTokenID,
		})
		_node.TokenID = value
	}
	if value, ok := oic.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: orderitem.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := oic.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderitem.FieldRemark,
		})
		_node.Remark = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderItem.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderItemUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (oic *OrderItemCreate) OnConflict(opts ...sql.ConflictOption) *OrderItemUpsertOne {
	oic.conflict = opts
	return &OrderItemUpsertOne{
		create: oic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oic *OrderItemCreate) OnConflictColumns(columns ...string) *OrderItemUpsertOne {
	oic.conflict = append(oic.conflict, sql.ConflictColumns(columns...))
	return &OrderItemUpsertOne{
		create: oic,
	}
}

type (
	// OrderItemUpsertOne is the builder for "upsert"-ing
	//  one OrderItem node.
	OrderItemUpsertOne struct {
		create *OrderItemCreate
	}

	// OrderItemUpsert is the "OnConflict" setter.
	OrderItemUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *OrderItemUpsert) SetCreatedAt(v uint32) *OrderItemUpsert {
	u.Set(orderitem.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderItemUpsert) UpdateCreatedAt() *OrderItemUpsert {
	u.SetExcluded(orderitem.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderItemUpsert) AddCreatedAt(v uint32) *OrderItemUpsert {
	u.Add(orderitem.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderItemUpsert) SetUpdatedAt(v uint32) *OrderItemUpsert {
	u.Set(orderitem.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderItemUpsert) UpdateUpdatedAt() *OrderItemUpsert {
	u.SetExcluded(orderitem.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderItemUpsert) AddUpdatedAt(v uint32) *OrderItemUpsert {
	u.Add(orderitem.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderItemUpsert) SetDeletedAt(v uint32) *OrderItemUpsert {
	u.Set(orderitem.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderItemUpsert) UpdateDeletedAt() *OrderItemUpsert {
	u.SetExcluded(orderitem.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderItemUpsert) AddDeletedAt(v uint32) *OrderItemUpsert {
	u.Add(orderitem.FieldDeletedAt, v)
	return u
}

// SetOrderID sets the "order_id" field.
func (u *OrderItemUpsert) SetOrderID(v string) *OrderItemUpsert {
	u.Set(orderitem.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderItemUpsert) UpdateOrderID() *OrderItemUpsert {
	u.SetExcluded(orderitem.FieldOrderID)
	return u
}

// SetOrderItemType sets the "order_item_type" field.
func (u *OrderItemUpsert) SetOrderItemType(v string) *OrderItemUpsert {
	u.Set(orderitem.FieldOrderItemType, v)
	return u
}

// UpdateOrderItemType sets the "order_item_type" field to the value that was provided on create.
func (u *OrderItemUpsert) UpdateOrderItemType() *OrderItemUpsert {
	u.SetExcluded(orderitem.FieldOrderItemType)
	return u
}

// SetContract sets the "contract" field.
func (u *OrderItemUpsert) SetContract(v string) *OrderItemUpsert {
	u.Set(orderitem.FieldContract, v)
	return u
}

// UpdateContract sets the "contract" field to the value that was provided on create.
func (u *OrderItemUpsert) UpdateContract() *OrderItemUpsert {
	u.SetExcluded(orderitem.FieldContract)
	return u
}

// SetTokenType sets the "token_type" field.
func (u *OrderItemUpsert) SetTokenType(v string) *OrderItemUpsert {
	u.Set(orderitem.FieldTokenType, v)
	return u
}

// UpdateTokenType sets the "token_type" field to the value that was provided on create.
func (u *OrderItemUpsert) UpdateTokenType() *OrderItemUpsert {
	u.SetExcluded(orderitem.FieldTokenType)
	return u
}

// SetTokenID sets the "token_id" field.
func (u *OrderItemUpsert) SetTokenID(v string) *OrderItemUpsert {
	u.Set(orderitem.FieldTokenID, v)
	return u
}

// UpdateTokenID sets the "token_id" field to the value that was provided on create.
func (u *OrderItemUpsert) UpdateTokenID() *OrderItemUpsert {
	u.SetExcluded(orderitem.FieldTokenID)
	return u
}

// SetAmount sets the "amount" field.
func (u *OrderItemUpsert) SetAmount(v int64) *OrderItemUpsert {
	u.Set(orderitem.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *OrderItemUpsert) UpdateAmount() *OrderItemUpsert {
	u.SetExcluded(orderitem.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *OrderItemUpsert) AddAmount(v int64) *OrderItemUpsert {
	u.Add(orderitem.FieldAmount, v)
	return u
}

// SetRemark sets the "remark" field.
func (u *OrderItemUpsert) SetRemark(v string) *OrderItemUpsert {
	u.Set(orderitem.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *OrderItemUpsert) UpdateRemark() *OrderItemUpsert {
	u.SetExcluded(orderitem.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *OrderItemUpsert) ClearRemark() *OrderItemUpsert {
	u.SetNull(orderitem.FieldRemark)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrderItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orderitem.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrderItemUpsertOne) UpdateNewValues() *OrderItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(orderitem.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderItem.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrderItemUpsertOne) Ignore() *OrderItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderItemUpsertOne) DoNothing() *OrderItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderItemCreate.OnConflict
// documentation for more info.
func (u *OrderItemUpsertOne) Update(set func(*OrderItemUpsert)) *OrderItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderItemUpsertOne) SetCreatedAt(v uint32) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderItemUpsertOne) AddCreatedAt(v uint32) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderItemUpsertOne) UpdateCreatedAt() *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderItemUpsertOne) SetUpdatedAt(v uint32) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderItemUpsertOne) AddUpdatedAt(v uint32) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderItemUpsertOne) UpdateUpdatedAt() *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderItemUpsertOne) SetDeletedAt(v uint32) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderItemUpsertOne) AddDeletedAt(v uint32) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderItemUpsertOne) UpdateDeletedAt() *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetOrderID sets the "order_id" field.
func (u *OrderItemUpsertOne) SetOrderID(v string) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderItemUpsertOne) UpdateOrderID() *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateOrderID()
	})
}

// SetOrderItemType sets the "order_item_type" field.
func (u *OrderItemUpsertOne) SetOrderItemType(v string) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetOrderItemType(v)
	})
}

// UpdateOrderItemType sets the "order_item_type" field to the value that was provided on create.
func (u *OrderItemUpsertOne) UpdateOrderItemType() *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateOrderItemType()
	})
}

// SetContract sets the "contract" field.
func (u *OrderItemUpsertOne) SetContract(v string) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetContract(v)
	})
}

// UpdateContract sets the "contract" field to the value that was provided on create.
func (u *OrderItemUpsertOne) UpdateContract() *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateContract()
	})
}

// SetTokenType sets the "token_type" field.
func (u *OrderItemUpsertOne) SetTokenType(v string) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetTokenType(v)
	})
}

// UpdateTokenType sets the "token_type" field to the value that was provided on create.
func (u *OrderItemUpsertOne) UpdateTokenType() *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateTokenType()
	})
}

// SetTokenID sets the "token_id" field.
func (u *OrderItemUpsertOne) SetTokenID(v string) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetTokenID(v)
	})
}

// UpdateTokenID sets the "token_id" field to the value that was provided on create.
func (u *OrderItemUpsertOne) UpdateTokenID() *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateTokenID()
	})
}

// SetAmount sets the "amount" field.
func (u *OrderItemUpsertOne) SetAmount(v int64) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *OrderItemUpsertOne) AddAmount(v int64) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *OrderItemUpsertOne) UpdateAmount() *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateAmount()
	})
}

// SetRemark sets the "remark" field.
func (u *OrderItemUpsertOne) SetRemark(v string) *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *OrderItemUpsertOne) UpdateRemark() *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *OrderItemUpsertOne) ClearRemark() *OrderItemUpsertOne {
	return u.Update(func(s *OrderItemUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *OrderItemUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderItemCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderItemUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrderItemUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: OrderItemUpsertOne.ID is not supported by MySQL driver. Use OrderItemUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrderItemUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrderItemCreateBulk is the builder for creating many OrderItem entities in bulk.
type OrderItemCreateBulk struct {
	config
	builders []*OrderItemCreate
	conflict []sql.ConflictOption
}

// Save creates the OrderItem entities in the database.
func (oicb *OrderItemCreateBulk) Save(ctx context.Context) ([]*OrderItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oicb.builders))
	nodes := make([]*OrderItem, len(oicb.builders))
	mutators := make([]Mutator, len(oicb.builders))
	for i := range oicb.builders {
		func(i int, root context.Context) {
			builder := oicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, oicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = oicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, oicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) SaveX(ctx context.Context) []*OrderItem {
	v, err := oicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oicb *OrderItemCreateBulk) Exec(ctx context.Context) error {
	_, err := oicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) ExecX(ctx context.Context) {
	if err := oicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderItem.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderItemUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (oicb *OrderItemCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrderItemUpsertBulk {
	oicb.conflict = opts
	return &OrderItemUpsertBulk{
		create: oicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderItem.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oicb *OrderItemCreateBulk) OnConflictColumns(columns ...string) *OrderItemUpsertBulk {
	oicb.conflict = append(oicb.conflict, sql.ConflictColumns(columns...))
	return &OrderItemUpsertBulk{
		create: oicb,
	}
}

// OrderItemUpsertBulk is the builder for "upsert"-ing
// a bulk of OrderItem nodes.
type OrderItemUpsertBulk struct {
	create *OrderItemCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrderItem.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orderitem.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrderItemUpsertBulk) UpdateNewValues() *OrderItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(orderitem.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderItem.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrderItemUpsertBulk) Ignore() *OrderItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderItemUpsertBulk) DoNothing() *OrderItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderItemCreateBulk.OnConflict
// documentation for more info.
func (u *OrderItemUpsertBulk) Update(set func(*OrderItemUpsert)) *OrderItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderItemUpsertBulk) SetCreatedAt(v uint32) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderItemUpsertBulk) AddCreatedAt(v uint32) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderItemUpsertBulk) UpdateCreatedAt() *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderItemUpsertBulk) SetUpdatedAt(v uint32) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderItemUpsertBulk) AddUpdatedAt(v uint32) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderItemUpsertBulk) UpdateUpdatedAt() *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderItemUpsertBulk) SetDeletedAt(v uint32) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderItemUpsertBulk) AddDeletedAt(v uint32) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderItemUpsertBulk) UpdateDeletedAt() *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetOrderID sets the "order_id" field.
func (u *OrderItemUpsertBulk) SetOrderID(v string) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderItemUpsertBulk) UpdateOrderID() *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateOrderID()
	})
}

// SetOrderItemType sets the "order_item_type" field.
func (u *OrderItemUpsertBulk) SetOrderItemType(v string) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetOrderItemType(v)
	})
}

// UpdateOrderItemType sets the "order_item_type" field to the value that was provided on create.
func (u *OrderItemUpsertBulk) UpdateOrderItemType() *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateOrderItemType()
	})
}

// SetContract sets the "contract" field.
func (u *OrderItemUpsertBulk) SetContract(v string) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetContract(v)
	})
}

// UpdateContract sets the "contract" field to the value that was provided on create.
func (u *OrderItemUpsertBulk) UpdateContract() *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateContract()
	})
}

// SetTokenType sets the "token_type" field.
func (u *OrderItemUpsertBulk) SetTokenType(v string) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetTokenType(v)
	})
}

// UpdateTokenType sets the "token_type" field to the value that was provided on create.
func (u *OrderItemUpsertBulk) UpdateTokenType() *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateTokenType()
	})
}

// SetTokenID sets the "token_id" field.
func (u *OrderItemUpsertBulk) SetTokenID(v string) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetTokenID(v)
	})
}

// UpdateTokenID sets the "token_id" field to the value that was provided on create.
func (u *OrderItemUpsertBulk) UpdateTokenID() *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateTokenID()
	})
}

// SetAmount sets the "amount" field.
func (u *OrderItemUpsertBulk) SetAmount(v int64) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *OrderItemUpsertBulk) AddAmount(v int64) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *OrderItemUpsertBulk) UpdateAmount() *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateAmount()
	})
}

// SetRemark sets the "remark" field.
func (u *OrderItemUpsertBulk) SetRemark(v string) *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *OrderItemUpsertBulk) UpdateRemark() *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *OrderItemUpsertBulk) ClearRemark() *OrderItemUpsertBulk {
	return u.Update(func(s *OrderItemUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *OrderItemUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrderItemCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderItemCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderItemUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

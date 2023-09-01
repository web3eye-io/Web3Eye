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
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderpair"
)

// OrderPairCreate is the builder for creating a OrderPair entity.
type OrderPairCreate struct {
	config
	mutation *OrderPairMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (opc *OrderPairCreate) SetCreatedAt(u uint32) *OrderPairCreate {
	opc.mutation.SetCreatedAt(u)
	return opc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (opc *OrderPairCreate) SetNillableCreatedAt(u *uint32) *OrderPairCreate {
	if u != nil {
		opc.SetCreatedAt(*u)
	}
	return opc
}

// SetUpdatedAt sets the "updated_at" field.
func (opc *OrderPairCreate) SetUpdatedAt(u uint32) *OrderPairCreate {
	opc.mutation.SetUpdatedAt(u)
	return opc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (opc *OrderPairCreate) SetNillableUpdatedAt(u *uint32) *OrderPairCreate {
	if u != nil {
		opc.SetUpdatedAt(*u)
	}
	return opc
}

// SetDeletedAt sets the "deleted_at" field.
func (opc *OrderPairCreate) SetDeletedAt(u uint32) *OrderPairCreate {
	opc.mutation.SetDeletedAt(u)
	return opc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (opc *OrderPairCreate) SetNillableDeletedAt(u *uint32) *OrderPairCreate {
	if u != nil {
		opc.SetDeletedAt(*u)
	}
	return opc
}

// SetTxHash sets the "tx_hash" field.
func (opc *OrderPairCreate) SetTxHash(s string) *OrderPairCreate {
	opc.mutation.SetTxHash(s)
	return opc
}

// SetRecipient sets the "recipient" field.
func (opc *OrderPairCreate) SetRecipient(s string) *OrderPairCreate {
	opc.mutation.SetRecipient(s)
	return opc
}

// SetTargetID sets the "target_id" field.
func (opc *OrderPairCreate) SetTargetID(s string) *OrderPairCreate {
	opc.mutation.SetTargetID(s)
	return opc
}

// SetOfferID sets the "offer_id" field.
func (opc *OrderPairCreate) SetOfferID(s string) *OrderPairCreate {
	opc.mutation.SetOfferID(s)
	return opc
}

// SetRemark sets the "remark" field.
func (opc *OrderPairCreate) SetRemark(s string) *OrderPairCreate {
	opc.mutation.SetRemark(s)
	return opc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (opc *OrderPairCreate) SetNillableRemark(s *string) *OrderPairCreate {
	if s != nil {
		opc.SetRemark(*s)
	}
	return opc
}

// SetID sets the "id" field.
func (opc *OrderPairCreate) SetID(u uuid.UUID) *OrderPairCreate {
	opc.mutation.SetID(u)
	return opc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (opc *OrderPairCreate) SetNillableID(u *uuid.UUID) *OrderPairCreate {
	if u != nil {
		opc.SetID(*u)
	}
	return opc
}

// Mutation returns the OrderPairMutation object of the builder.
func (opc *OrderPairCreate) Mutation() *OrderPairMutation {
	return opc.mutation
}

// Save creates the OrderPair in the database.
func (opc *OrderPairCreate) Save(ctx context.Context) (*OrderPair, error) {
	var (
		err  error
		node *OrderPair
	)
	if err := opc.defaults(); err != nil {
		return nil, err
	}
	if len(opc.hooks) == 0 {
		if err = opc.check(); err != nil {
			return nil, err
		}
		node, err = opc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderPairMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = opc.check(); err != nil {
				return nil, err
			}
			opc.mutation = mutation
			if node, err = opc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(opc.hooks) - 1; i >= 0; i-- {
			if opc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = opc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, opc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderPair)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderPairMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (opc *OrderPairCreate) SaveX(ctx context.Context) *OrderPair {
	v, err := opc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opc *OrderPairCreate) Exec(ctx context.Context) error {
	_, err := opc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opc *OrderPairCreate) ExecX(ctx context.Context) {
	if err := opc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opc *OrderPairCreate) defaults() error {
	if _, ok := opc.mutation.CreatedAt(); !ok {
		if orderpair.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderpair.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := orderpair.DefaultCreatedAt()
		opc.mutation.SetCreatedAt(v)
	}
	if _, ok := opc.mutation.UpdatedAt(); !ok {
		if orderpair.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderpair.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderpair.DefaultUpdatedAt()
		opc.mutation.SetUpdatedAt(v)
	}
	if _, ok := opc.mutation.DeletedAt(); !ok {
		if orderpair.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized orderpair.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := orderpair.DefaultDeletedAt()
		opc.mutation.SetDeletedAt(v)
	}
	if _, ok := opc.mutation.ID(); !ok {
		if orderpair.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized orderpair.DefaultID (forgotten import ent/runtime?)")
		}
		v := orderpair.DefaultID()
		opc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (opc *OrderPairCreate) check() error {
	if _, ok := opc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderPair.created_at"`)}
	}
	if _, ok := opc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderPair.updated_at"`)}
	}
	if _, ok := opc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "OrderPair.deleted_at"`)}
	}
	if _, ok := opc.mutation.TxHash(); !ok {
		return &ValidationError{Name: "tx_hash", err: errors.New(`ent: missing required field "OrderPair.tx_hash"`)}
	}
	if _, ok := opc.mutation.Recipient(); !ok {
		return &ValidationError{Name: "recipient", err: errors.New(`ent: missing required field "OrderPair.recipient"`)}
	}
	if _, ok := opc.mutation.TargetID(); !ok {
		return &ValidationError{Name: "target_id", err: errors.New(`ent: missing required field "OrderPair.target_id"`)}
	}
	if _, ok := opc.mutation.OfferID(); !ok {
		return &ValidationError{Name: "offer_id", err: errors.New(`ent: missing required field "OrderPair.offer_id"`)}
	}
	return nil
}

func (opc *OrderPairCreate) sqlSave(ctx context.Context) (*OrderPair, error) {
	_node, _spec := opc.createSpec()
	if err := sqlgraph.CreateNode(ctx, opc.driver, _spec); err != nil {
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

func (opc *OrderPairCreate) createSpec() (*OrderPair, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderPair{config: opc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderpair.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: orderpair.FieldID,
			},
		}
	)
	_spec.OnConflict = opc.conflict
	if id, ok := opc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := opc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpair.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := opc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpair.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := opc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpair.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := opc.mutation.TxHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderpair.FieldTxHash,
		})
		_node.TxHash = value
	}
	if value, ok := opc.mutation.Recipient(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderpair.FieldRecipient,
		})
		_node.Recipient = value
	}
	if value, ok := opc.mutation.TargetID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderpair.FieldTargetID,
		})
		_node.TargetID = value
	}
	if value, ok := opc.mutation.OfferID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderpair.FieldOfferID,
		})
		_node.OfferID = value
	}
	if value, ok := opc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderpair.FieldRemark,
		})
		_node.Remark = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderPair.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderPairUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (opc *OrderPairCreate) OnConflict(opts ...sql.ConflictOption) *OrderPairUpsertOne {
	opc.conflict = opts
	return &OrderPairUpsertOne{
		create: opc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderPair.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (opc *OrderPairCreate) OnConflictColumns(columns ...string) *OrderPairUpsertOne {
	opc.conflict = append(opc.conflict, sql.ConflictColumns(columns...))
	return &OrderPairUpsertOne{
		create: opc,
	}
}

type (
	// OrderPairUpsertOne is the builder for "upsert"-ing
	//  one OrderPair node.
	OrderPairUpsertOne struct {
		create *OrderPairCreate
	}

	// OrderPairUpsert is the "OnConflict" setter.
	OrderPairUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *OrderPairUpsert) SetCreatedAt(v uint32) *OrderPairUpsert {
	u.Set(orderpair.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderPairUpsert) UpdateCreatedAt() *OrderPairUpsert {
	u.SetExcluded(orderpair.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderPairUpsert) AddCreatedAt(v uint32) *OrderPairUpsert {
	u.Add(orderpair.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderPairUpsert) SetUpdatedAt(v uint32) *OrderPairUpsert {
	u.Set(orderpair.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderPairUpsert) UpdateUpdatedAt() *OrderPairUpsert {
	u.SetExcluded(orderpair.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderPairUpsert) AddUpdatedAt(v uint32) *OrderPairUpsert {
	u.Add(orderpair.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderPairUpsert) SetDeletedAt(v uint32) *OrderPairUpsert {
	u.Set(orderpair.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderPairUpsert) UpdateDeletedAt() *OrderPairUpsert {
	u.SetExcluded(orderpair.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderPairUpsert) AddDeletedAt(v uint32) *OrderPairUpsert {
	u.Add(orderpair.FieldDeletedAt, v)
	return u
}

// SetTxHash sets the "tx_hash" field.
func (u *OrderPairUpsert) SetTxHash(v string) *OrderPairUpsert {
	u.Set(orderpair.FieldTxHash, v)
	return u
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *OrderPairUpsert) UpdateTxHash() *OrderPairUpsert {
	u.SetExcluded(orderpair.FieldTxHash)
	return u
}

// SetRecipient sets the "recipient" field.
func (u *OrderPairUpsert) SetRecipient(v string) *OrderPairUpsert {
	u.Set(orderpair.FieldRecipient, v)
	return u
}

// UpdateRecipient sets the "recipient" field to the value that was provided on create.
func (u *OrderPairUpsert) UpdateRecipient() *OrderPairUpsert {
	u.SetExcluded(orderpair.FieldRecipient)
	return u
}

// SetTargetID sets the "target_id" field.
func (u *OrderPairUpsert) SetTargetID(v string) *OrderPairUpsert {
	u.Set(orderpair.FieldTargetID, v)
	return u
}

// UpdateTargetID sets the "target_id" field to the value that was provided on create.
func (u *OrderPairUpsert) UpdateTargetID() *OrderPairUpsert {
	u.SetExcluded(orderpair.FieldTargetID)
	return u
}

// SetOfferID sets the "offer_id" field.
func (u *OrderPairUpsert) SetOfferID(v string) *OrderPairUpsert {
	u.Set(orderpair.FieldOfferID, v)
	return u
}

// UpdateOfferID sets the "offer_id" field to the value that was provided on create.
func (u *OrderPairUpsert) UpdateOfferID() *OrderPairUpsert {
	u.SetExcluded(orderpair.FieldOfferID)
	return u
}

// SetRemark sets the "remark" field.
func (u *OrderPairUpsert) SetRemark(v string) *OrderPairUpsert {
	u.Set(orderpair.FieldRemark, v)
	return u
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *OrderPairUpsert) UpdateRemark() *OrderPairUpsert {
	u.SetExcluded(orderpair.FieldRemark)
	return u
}

// ClearRemark clears the value of the "remark" field.
func (u *OrderPairUpsert) ClearRemark() *OrderPairUpsert {
	u.SetNull(orderpair.FieldRemark)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrderPair.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orderpair.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrderPairUpsertOne) UpdateNewValues() *OrderPairUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(orderpair.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderPair.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrderPairUpsertOne) Ignore() *OrderPairUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderPairUpsertOne) DoNothing() *OrderPairUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderPairCreate.OnConflict
// documentation for more info.
func (u *OrderPairUpsertOne) Update(set func(*OrderPairUpsert)) *OrderPairUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderPairUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderPairUpsertOne) SetCreatedAt(v uint32) *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderPairUpsertOne) AddCreatedAt(v uint32) *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderPairUpsertOne) UpdateCreatedAt() *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderPairUpsertOne) SetUpdatedAt(v uint32) *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderPairUpsertOne) AddUpdatedAt(v uint32) *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderPairUpsertOne) UpdateUpdatedAt() *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderPairUpsertOne) SetDeletedAt(v uint32) *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderPairUpsertOne) AddDeletedAt(v uint32) *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderPairUpsertOne) UpdateDeletedAt() *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetTxHash sets the "tx_hash" field.
func (u *OrderPairUpsertOne) SetTxHash(v string) *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetTxHash(v)
	})
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *OrderPairUpsertOne) UpdateTxHash() *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateTxHash()
	})
}

// SetRecipient sets the "recipient" field.
func (u *OrderPairUpsertOne) SetRecipient(v string) *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetRecipient(v)
	})
}

// UpdateRecipient sets the "recipient" field to the value that was provided on create.
func (u *OrderPairUpsertOne) UpdateRecipient() *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateRecipient()
	})
}

// SetTargetID sets the "target_id" field.
func (u *OrderPairUpsertOne) SetTargetID(v string) *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetTargetID(v)
	})
}

// UpdateTargetID sets the "target_id" field to the value that was provided on create.
func (u *OrderPairUpsertOne) UpdateTargetID() *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateTargetID()
	})
}

// SetOfferID sets the "offer_id" field.
func (u *OrderPairUpsertOne) SetOfferID(v string) *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetOfferID(v)
	})
}

// UpdateOfferID sets the "offer_id" field to the value that was provided on create.
func (u *OrderPairUpsertOne) UpdateOfferID() *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateOfferID()
	})
}

// SetRemark sets the "remark" field.
func (u *OrderPairUpsertOne) SetRemark(v string) *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *OrderPairUpsertOne) UpdateRemark() *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *OrderPairUpsertOne) ClearRemark() *OrderPairUpsertOne {
	return u.Update(func(s *OrderPairUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *OrderPairUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderPairCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderPairUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrderPairUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: OrderPairUpsertOne.ID is not supported by MySQL driver. Use OrderPairUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrderPairUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrderPairCreateBulk is the builder for creating many OrderPair entities in bulk.
type OrderPairCreateBulk struct {
	config
	builders []*OrderPairCreate
	conflict []sql.ConflictOption
}

// Save creates the OrderPair entities in the database.
func (opcb *OrderPairCreateBulk) Save(ctx context.Context) ([]*OrderPair, error) {
	specs := make([]*sqlgraph.CreateSpec, len(opcb.builders))
	nodes := make([]*OrderPair, len(opcb.builders))
	mutators := make([]Mutator, len(opcb.builders))
	for i := range opcb.builders {
		func(i int, root context.Context) {
			builder := opcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderPairMutation)
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
					_, err = mutators[i+1].Mutate(root, opcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = opcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, opcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, opcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (opcb *OrderPairCreateBulk) SaveX(ctx context.Context) []*OrderPair {
	v, err := opcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opcb *OrderPairCreateBulk) Exec(ctx context.Context) error {
	_, err := opcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opcb *OrderPairCreateBulk) ExecX(ctx context.Context) {
	if err := opcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderPair.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderPairUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (opcb *OrderPairCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrderPairUpsertBulk {
	opcb.conflict = opts
	return &OrderPairUpsertBulk{
		create: opcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderPair.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (opcb *OrderPairCreateBulk) OnConflictColumns(columns ...string) *OrderPairUpsertBulk {
	opcb.conflict = append(opcb.conflict, sql.ConflictColumns(columns...))
	return &OrderPairUpsertBulk{
		create: opcb,
	}
}

// OrderPairUpsertBulk is the builder for "upsert"-ing
// a bulk of OrderPair nodes.
type OrderPairUpsertBulk struct {
	create *OrderPairCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrderPair.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orderpair.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrderPairUpsertBulk) UpdateNewValues() *OrderPairUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(orderpair.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderPair.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrderPairUpsertBulk) Ignore() *OrderPairUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderPairUpsertBulk) DoNothing() *OrderPairUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderPairCreateBulk.OnConflict
// documentation for more info.
func (u *OrderPairUpsertBulk) Update(set func(*OrderPairUpsert)) *OrderPairUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderPairUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderPairUpsertBulk) SetCreatedAt(v uint32) *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderPairUpsertBulk) AddCreatedAt(v uint32) *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderPairUpsertBulk) UpdateCreatedAt() *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderPairUpsertBulk) SetUpdatedAt(v uint32) *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderPairUpsertBulk) AddUpdatedAt(v uint32) *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderPairUpsertBulk) UpdateUpdatedAt() *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderPairUpsertBulk) SetDeletedAt(v uint32) *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderPairUpsertBulk) AddDeletedAt(v uint32) *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderPairUpsertBulk) UpdateDeletedAt() *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetTxHash sets the "tx_hash" field.
func (u *OrderPairUpsertBulk) SetTxHash(v string) *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetTxHash(v)
	})
}

// UpdateTxHash sets the "tx_hash" field to the value that was provided on create.
func (u *OrderPairUpsertBulk) UpdateTxHash() *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateTxHash()
	})
}

// SetRecipient sets the "recipient" field.
func (u *OrderPairUpsertBulk) SetRecipient(v string) *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetRecipient(v)
	})
}

// UpdateRecipient sets the "recipient" field to the value that was provided on create.
func (u *OrderPairUpsertBulk) UpdateRecipient() *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateRecipient()
	})
}

// SetTargetID sets the "target_id" field.
func (u *OrderPairUpsertBulk) SetTargetID(v string) *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetTargetID(v)
	})
}

// UpdateTargetID sets the "target_id" field to the value that was provided on create.
func (u *OrderPairUpsertBulk) UpdateTargetID() *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateTargetID()
	})
}

// SetOfferID sets the "offer_id" field.
func (u *OrderPairUpsertBulk) SetOfferID(v string) *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetOfferID(v)
	})
}

// UpdateOfferID sets the "offer_id" field to the value that was provided on create.
func (u *OrderPairUpsertBulk) UpdateOfferID() *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateOfferID()
	})
}

// SetRemark sets the "remark" field.
func (u *OrderPairUpsertBulk) SetRemark(v string) *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.SetRemark(v)
	})
}

// UpdateRemark sets the "remark" field to the value that was provided on create.
func (u *OrderPairUpsertBulk) UpdateRemark() *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.UpdateRemark()
	})
}

// ClearRemark clears the value of the "remark" field.
func (u *OrderPairUpsertBulk) ClearRemark() *OrderPairUpsertBulk {
	return u.Update(func(s *OrderPairUpsert) {
		s.ClearRemark()
	})
}

// Exec executes the query.
func (u *OrderPairUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrderPairCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderPairCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderPairUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

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
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/endpoint"
)

// EndpointCreate is the builder for creating a Endpoint entity.
type EndpointCreate struct {
	config
	mutation *EndpointMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ec *EndpointCreate) SetCreatedAt(u uint32) *EndpointCreate {
	ec.mutation.SetCreatedAt(u)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EndpointCreate) SetNillableCreatedAt(u *uint32) *EndpointCreate {
	if u != nil {
		ec.SetCreatedAt(*u)
	}
	return ec
}

// SetUpdatedAt sets the "updated_at" field.
func (ec *EndpointCreate) SetUpdatedAt(u uint32) *EndpointCreate {
	ec.mutation.SetUpdatedAt(u)
	return ec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ec *EndpointCreate) SetNillableUpdatedAt(u *uint32) *EndpointCreate {
	if u != nil {
		ec.SetUpdatedAt(*u)
	}
	return ec
}

// SetDeletedAt sets the "deleted_at" field.
func (ec *EndpointCreate) SetDeletedAt(u uint32) *EndpointCreate {
	ec.mutation.SetDeletedAt(u)
	return ec
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ec *EndpointCreate) SetNillableDeletedAt(u *uint32) *EndpointCreate {
	if u != nil {
		ec.SetDeletedAt(*u)
	}
	return ec
}

// SetChainType sets the "chain_type" field.
func (ec *EndpointCreate) SetChainType(s string) *EndpointCreate {
	ec.mutation.SetChainType(s)
	return ec
}

// SetChainID sets the "chain_id" field.
func (ec *EndpointCreate) SetChainID(s string) *EndpointCreate {
	ec.mutation.SetChainID(s)
	return ec
}

// SetNillableChainID sets the "chain_id" field if the given value is not nil.
func (ec *EndpointCreate) SetNillableChainID(s *string) *EndpointCreate {
	if s != nil {
		ec.SetChainID(*s)
	}
	return ec
}

// SetAddress sets the "address" field.
func (ec *EndpointCreate) SetAddress(s string) *EndpointCreate {
	ec.mutation.SetAddress(s)
	return ec
}

// SetState sets the "state" field.
func (ec *EndpointCreate) SetState(s string) *EndpointCreate {
	ec.mutation.SetState(s)
	return ec
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ec *EndpointCreate) SetNillableState(s *string) *EndpointCreate {
	if s != nil {
		ec.SetState(*s)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *EndpointCreate) SetID(u uuid.UUID) *EndpointCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EndpointCreate) SetNillableID(u *uuid.UUID) *EndpointCreate {
	if u != nil {
		ec.SetID(*u)
	}
	return ec
}

// Mutation returns the EndpointMutation object of the builder.
func (ec *EndpointCreate) Mutation() *EndpointMutation {
	return ec.mutation
}

// Save creates the Endpoint in the database.
func (ec *EndpointCreate) Save(ctx context.Context) (*Endpoint, error) {
	var (
		err  error
		node *Endpoint
	)
	if err := ec.defaults(); err != nil {
		return nil, err
	}
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EndpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Endpoint)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EndpointMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EndpointCreate) SaveX(ctx context.Context) *Endpoint {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EndpointCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EndpointCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EndpointCreate) defaults() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		if endpoint.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized endpoint.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := endpoint.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		if endpoint.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized endpoint.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := endpoint.DefaultUpdatedAt()
		ec.mutation.SetUpdatedAt(v)
	}
	if _, ok := ec.mutation.DeletedAt(); !ok {
		if endpoint.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized endpoint.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := endpoint.DefaultDeletedAt()
		ec.mutation.SetDeletedAt(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		if endpoint.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized endpoint.DefaultID (forgotten import ent/runtime?)")
		}
		v := endpoint.DefaultID()
		ec.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ec *EndpointCreate) check() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Endpoint.created_at"`)}
	}
	if _, ok := ec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Endpoint.updated_at"`)}
	}
	if _, ok := ec.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Endpoint.deleted_at"`)}
	}
	if _, ok := ec.mutation.ChainType(); !ok {
		return &ValidationError{Name: "chain_type", err: errors.New(`ent: missing required field "Endpoint.chain_type"`)}
	}
	if _, ok := ec.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Endpoint.address"`)}
	}
	return nil
}

func (ec *EndpointCreate) sqlSave(ctx context.Context) (*Endpoint, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
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

func (ec *EndpointCreate) createSpec() (*Endpoint, *sqlgraph.CreateSpec) {
	var (
		_node = &Endpoint{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: endpoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: endpoint.FieldID,
			},
		}
	)
	_spec.OnConflict = ec.conflict
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: endpoint.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: endpoint.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ec.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: endpoint.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := ec.mutation.ChainType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: endpoint.FieldChainType,
		})
		_node.ChainType = value
	}
	if value, ok := ec.mutation.ChainID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: endpoint.FieldChainID,
		})
		_node.ChainID = value
	}
	if value, ok := ec.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: endpoint.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := ec.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: endpoint.FieldState,
		})
		_node.State = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Endpoint.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EndpointUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ec *EndpointCreate) OnConflict(opts ...sql.ConflictOption) *EndpointUpsertOne {
	ec.conflict = opts
	return &EndpointUpsertOne{
		create: ec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Endpoint.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ec *EndpointCreate) OnConflictColumns(columns ...string) *EndpointUpsertOne {
	ec.conflict = append(ec.conflict, sql.ConflictColumns(columns...))
	return &EndpointUpsertOne{
		create: ec,
	}
}

type (
	// EndpointUpsertOne is the builder for "upsert"-ing
	//  one Endpoint node.
	EndpointUpsertOne struct {
		create *EndpointCreate
	}

	// EndpointUpsert is the "OnConflict" setter.
	EndpointUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *EndpointUpsert) SetCreatedAt(v uint32) *EndpointUpsert {
	u.Set(endpoint.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *EndpointUpsert) UpdateCreatedAt() *EndpointUpsert {
	u.SetExcluded(endpoint.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *EndpointUpsert) AddCreatedAt(v uint32) *EndpointUpsert {
	u.Add(endpoint.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EndpointUpsert) SetUpdatedAt(v uint32) *EndpointUpsert {
	u.Set(endpoint.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EndpointUpsert) UpdateUpdatedAt() *EndpointUpsert {
	u.SetExcluded(endpoint.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *EndpointUpsert) AddUpdatedAt(v uint32) *EndpointUpsert {
	u.Add(endpoint.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EndpointUpsert) SetDeletedAt(v uint32) *EndpointUpsert {
	u.Set(endpoint.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EndpointUpsert) UpdateDeletedAt() *EndpointUpsert {
	u.SetExcluded(endpoint.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *EndpointUpsert) AddDeletedAt(v uint32) *EndpointUpsert {
	u.Add(endpoint.FieldDeletedAt, v)
	return u
}

// SetChainType sets the "chain_type" field.
func (u *EndpointUpsert) SetChainType(v string) *EndpointUpsert {
	u.Set(endpoint.FieldChainType, v)
	return u
}

// UpdateChainType sets the "chain_type" field to the value that was provided on create.
func (u *EndpointUpsert) UpdateChainType() *EndpointUpsert {
	u.SetExcluded(endpoint.FieldChainType)
	return u
}

// SetChainID sets the "chain_id" field.
func (u *EndpointUpsert) SetChainID(v string) *EndpointUpsert {
	u.Set(endpoint.FieldChainID, v)
	return u
}

// UpdateChainID sets the "chain_id" field to the value that was provided on create.
func (u *EndpointUpsert) UpdateChainID() *EndpointUpsert {
	u.SetExcluded(endpoint.FieldChainID)
	return u
}

// ClearChainID clears the value of the "chain_id" field.
func (u *EndpointUpsert) ClearChainID() *EndpointUpsert {
	u.SetNull(endpoint.FieldChainID)
	return u
}

// SetAddress sets the "address" field.
func (u *EndpointUpsert) SetAddress(v string) *EndpointUpsert {
	u.Set(endpoint.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *EndpointUpsert) UpdateAddress() *EndpointUpsert {
	u.SetExcluded(endpoint.FieldAddress)
	return u
}

// SetState sets the "state" field.
func (u *EndpointUpsert) SetState(v string) *EndpointUpsert {
	u.Set(endpoint.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *EndpointUpsert) UpdateState() *EndpointUpsert {
	u.SetExcluded(endpoint.FieldState)
	return u
}

// ClearState clears the value of the "state" field.
func (u *EndpointUpsert) ClearState() *EndpointUpsert {
	u.SetNull(endpoint.FieldState)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Endpoint.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(endpoint.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EndpointUpsertOne) UpdateNewValues() *EndpointUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(endpoint.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Endpoint.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EndpointUpsertOne) Ignore() *EndpointUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EndpointUpsertOne) DoNothing() *EndpointUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EndpointCreate.OnConflict
// documentation for more info.
func (u *EndpointUpsertOne) Update(set func(*EndpointUpsert)) *EndpointUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EndpointUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *EndpointUpsertOne) SetCreatedAt(v uint32) *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *EndpointUpsertOne) AddCreatedAt(v uint32) *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *EndpointUpsertOne) UpdateCreatedAt() *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EndpointUpsertOne) SetUpdatedAt(v uint32) *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *EndpointUpsertOne) AddUpdatedAt(v uint32) *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EndpointUpsertOne) UpdateUpdatedAt() *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EndpointUpsertOne) SetDeletedAt(v uint32) *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *EndpointUpsertOne) AddDeletedAt(v uint32) *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EndpointUpsertOne) UpdateDeletedAt() *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetChainType sets the "chain_type" field.
func (u *EndpointUpsertOne) SetChainType(v string) *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.SetChainType(v)
	})
}

// UpdateChainType sets the "chain_type" field to the value that was provided on create.
func (u *EndpointUpsertOne) UpdateChainType() *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateChainType()
	})
}

// SetChainID sets the "chain_id" field.
func (u *EndpointUpsertOne) SetChainID(v string) *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.SetChainID(v)
	})
}

// UpdateChainID sets the "chain_id" field to the value that was provided on create.
func (u *EndpointUpsertOne) UpdateChainID() *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateChainID()
	})
}

// ClearChainID clears the value of the "chain_id" field.
func (u *EndpointUpsertOne) ClearChainID() *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.ClearChainID()
	})
}

// SetAddress sets the "address" field.
func (u *EndpointUpsertOne) SetAddress(v string) *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *EndpointUpsertOne) UpdateAddress() *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateAddress()
	})
}

// SetState sets the "state" field.
func (u *EndpointUpsertOne) SetState(v string) *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *EndpointUpsertOne) UpdateState() *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateState()
	})
}

// ClearState clears the value of the "state" field.
func (u *EndpointUpsertOne) ClearState() *EndpointUpsertOne {
	return u.Update(func(s *EndpointUpsert) {
		s.ClearState()
	})
}

// Exec executes the query.
func (u *EndpointUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EndpointCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EndpointUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EndpointUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: EndpointUpsertOne.ID is not supported by MySQL driver. Use EndpointUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EndpointUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EndpointCreateBulk is the builder for creating many Endpoint entities in bulk.
type EndpointCreateBulk struct {
	config
	builders []*EndpointCreate
	conflict []sql.ConflictOption
}

// Save creates the Endpoint entities in the database.
func (ecb *EndpointCreateBulk) Save(ctx context.Context) ([]*Endpoint, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Endpoint, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EndpointMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EndpointCreateBulk) SaveX(ctx context.Context) []*Endpoint {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EndpointCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EndpointCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Endpoint.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EndpointUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ecb *EndpointCreateBulk) OnConflict(opts ...sql.ConflictOption) *EndpointUpsertBulk {
	ecb.conflict = opts
	return &EndpointUpsertBulk{
		create: ecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Endpoint.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ecb *EndpointCreateBulk) OnConflictColumns(columns ...string) *EndpointUpsertBulk {
	ecb.conflict = append(ecb.conflict, sql.ConflictColumns(columns...))
	return &EndpointUpsertBulk{
		create: ecb,
	}
}

// EndpointUpsertBulk is the builder for "upsert"-ing
// a bulk of Endpoint nodes.
type EndpointUpsertBulk struct {
	create *EndpointCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Endpoint.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(endpoint.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EndpointUpsertBulk) UpdateNewValues() *EndpointUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(endpoint.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Endpoint.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EndpointUpsertBulk) Ignore() *EndpointUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EndpointUpsertBulk) DoNothing() *EndpointUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EndpointCreateBulk.OnConflict
// documentation for more info.
func (u *EndpointUpsertBulk) Update(set func(*EndpointUpsert)) *EndpointUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EndpointUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *EndpointUpsertBulk) SetCreatedAt(v uint32) *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *EndpointUpsertBulk) AddCreatedAt(v uint32) *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *EndpointUpsertBulk) UpdateCreatedAt() *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EndpointUpsertBulk) SetUpdatedAt(v uint32) *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *EndpointUpsertBulk) AddUpdatedAt(v uint32) *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EndpointUpsertBulk) UpdateUpdatedAt() *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EndpointUpsertBulk) SetDeletedAt(v uint32) *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *EndpointUpsertBulk) AddDeletedAt(v uint32) *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EndpointUpsertBulk) UpdateDeletedAt() *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetChainType sets the "chain_type" field.
func (u *EndpointUpsertBulk) SetChainType(v string) *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.SetChainType(v)
	})
}

// UpdateChainType sets the "chain_type" field to the value that was provided on create.
func (u *EndpointUpsertBulk) UpdateChainType() *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateChainType()
	})
}

// SetChainID sets the "chain_id" field.
func (u *EndpointUpsertBulk) SetChainID(v string) *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.SetChainID(v)
	})
}

// UpdateChainID sets the "chain_id" field to the value that was provided on create.
func (u *EndpointUpsertBulk) UpdateChainID() *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateChainID()
	})
}

// ClearChainID clears the value of the "chain_id" field.
func (u *EndpointUpsertBulk) ClearChainID() *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.ClearChainID()
	})
}

// SetAddress sets the "address" field.
func (u *EndpointUpsertBulk) SetAddress(v string) *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *EndpointUpsertBulk) UpdateAddress() *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateAddress()
	})
}

// SetState sets the "state" field.
func (u *EndpointUpsertBulk) SetState(v string) *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *EndpointUpsertBulk) UpdateState() *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.UpdateState()
	})
}

// ClearState clears the value of the "state" field.
func (u *EndpointUpsertBulk) ClearState() *EndpointUpsertBulk {
	return u.Update(func(s *EndpointUpsert) {
		s.ClearState()
	})
}

// Exec executes the query.
func (u *EndpointUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EndpointCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EndpointCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EndpointUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

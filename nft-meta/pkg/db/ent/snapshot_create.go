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
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/snapshot"
)

// SnapshotCreate is the builder for creating a Snapshot entity.
type SnapshotCreate struct {
	config
	mutation *SnapshotMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEntID sets the "ent_id" field.
func (sc *SnapshotCreate) SetEntID(u uuid.UUID) *SnapshotCreate {
	sc.mutation.SetEntID(u)
	return sc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (sc *SnapshotCreate) SetNillableEntID(u *uuid.UUID) *SnapshotCreate {
	if u != nil {
		sc.SetEntID(*u)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SnapshotCreate) SetCreatedAt(u uint32) *SnapshotCreate {
	sc.mutation.SetCreatedAt(u)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SnapshotCreate) SetNillableCreatedAt(u *uint32) *SnapshotCreate {
	if u != nil {
		sc.SetCreatedAt(*u)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SnapshotCreate) SetUpdatedAt(u uint32) *SnapshotCreate {
	sc.mutation.SetUpdatedAt(u)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SnapshotCreate) SetNillableUpdatedAt(u *uint32) *SnapshotCreate {
	if u != nil {
		sc.SetUpdatedAt(*u)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *SnapshotCreate) SetDeletedAt(u uint32) *SnapshotCreate {
	sc.mutation.SetDeletedAt(u)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *SnapshotCreate) SetNillableDeletedAt(u *uint32) *SnapshotCreate {
	if u != nil {
		sc.SetDeletedAt(*u)
	}
	return sc
}

// SetIndex sets the "index" field.
func (sc *SnapshotCreate) SetIndex(u uint64) *SnapshotCreate {
	sc.mutation.SetIndex(u)
	return sc
}

// SetSnapshotCommP sets the "snapshot_comm_p" field.
func (sc *SnapshotCreate) SetSnapshotCommP(s string) *SnapshotCreate {
	sc.mutation.SetSnapshotCommP(s)
	return sc
}

// SetSnapshotRoot sets the "snapshot_root" field.
func (sc *SnapshotCreate) SetSnapshotRoot(s string) *SnapshotCreate {
	sc.mutation.SetSnapshotRoot(s)
	return sc
}

// SetSnapshotURI sets the "snapshot_uri" field.
func (sc *SnapshotCreate) SetSnapshotURI(s string) *SnapshotCreate {
	sc.mutation.SetSnapshotURI(s)
	return sc
}

// SetBackupState sets the "backup_state" field.
func (sc *SnapshotCreate) SetBackupState(s string) *SnapshotCreate {
	sc.mutation.SetBackupState(s)
	return sc
}

// SetID sets the "id" field.
func (sc *SnapshotCreate) SetID(u uint32) *SnapshotCreate {
	sc.mutation.SetID(u)
	return sc
}

// Mutation returns the SnapshotMutation object of the builder.
func (sc *SnapshotCreate) Mutation() *SnapshotMutation {
	return sc.mutation
}

// Save creates the Snapshot in the database.
func (sc *SnapshotCreate) Save(ctx context.Context) (*Snapshot, error) {
	var (
		err  error
		node *Snapshot
	)
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SnapshotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Snapshot)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SnapshotMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SnapshotCreate) SaveX(ctx context.Context) *Snapshot {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SnapshotCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SnapshotCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SnapshotCreate) defaults() error {
	if _, ok := sc.mutation.EntID(); !ok {
		if snapshot.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized snapshot.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := snapshot.DefaultEntID()
		sc.mutation.SetEntID(v)
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if snapshot.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized snapshot.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := snapshot.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if snapshot.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized snapshot.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := snapshot.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.DeletedAt(); !ok {
		if snapshot.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized snapshot.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := snapshot.DefaultDeletedAt()
		sc.mutation.SetDeletedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SnapshotCreate) check() error {
	if _, ok := sc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "Snapshot.ent_id"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Snapshot.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Snapshot.updated_at"`)}
	}
	if _, ok := sc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Snapshot.deleted_at"`)}
	}
	if _, ok := sc.mutation.Index(); !ok {
		return &ValidationError{Name: "index", err: errors.New(`ent: missing required field "Snapshot.index"`)}
	}
	if _, ok := sc.mutation.SnapshotCommP(); !ok {
		return &ValidationError{Name: "snapshot_comm_p", err: errors.New(`ent: missing required field "Snapshot.snapshot_comm_p"`)}
	}
	if _, ok := sc.mutation.SnapshotRoot(); !ok {
		return &ValidationError{Name: "snapshot_root", err: errors.New(`ent: missing required field "Snapshot.snapshot_root"`)}
	}
	if _, ok := sc.mutation.SnapshotURI(); !ok {
		return &ValidationError{Name: "snapshot_uri", err: errors.New(`ent: missing required field "Snapshot.snapshot_uri"`)}
	}
	if _, ok := sc.mutation.BackupState(); !ok {
		return &ValidationError{Name: "backup_state", err: errors.New(`ent: missing required field "Snapshot.backup_state"`)}
	}
	return nil
}

func (sc *SnapshotCreate) sqlSave(ctx context.Context) (*Snapshot, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (sc *SnapshotCreate) createSpec() (*Snapshot, *sqlgraph.CreateSpec) {
	var (
		_node = &Snapshot{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: snapshot.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: snapshot.FieldID,
			},
		}
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: snapshot.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: snapshot.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: snapshot.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: snapshot.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := sc.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: snapshot.FieldIndex,
		})
		_node.Index = value
	}
	if value, ok := sc.mutation.SnapshotCommP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: snapshot.FieldSnapshotCommP,
		})
		_node.SnapshotCommP = value
	}
	if value, ok := sc.mutation.SnapshotRoot(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: snapshot.FieldSnapshotRoot,
		})
		_node.SnapshotRoot = value
	}
	if value, ok := sc.mutation.SnapshotURI(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: snapshot.FieldSnapshotURI,
		})
		_node.SnapshotURI = value
	}
	if value, ok := sc.mutation.BackupState(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: snapshot.FieldBackupState,
		})
		_node.BackupState = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Snapshot.Create().
//		SetEntID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SnapshotUpsert) {
//			SetEntID(v+v).
//		}).
//		Exec(ctx)
func (sc *SnapshotCreate) OnConflict(opts ...sql.ConflictOption) *SnapshotUpsertOne {
	sc.conflict = opts
	return &SnapshotUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SnapshotCreate) OnConflictColumns(columns ...string) *SnapshotUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SnapshotUpsertOne{
		create: sc,
	}
}

type (
	// SnapshotUpsertOne is the builder for "upsert"-ing
	//  one Snapshot node.
	SnapshotUpsertOne struct {
		create *SnapshotCreate
	}

	// SnapshotUpsert is the "OnConflict" setter.
	SnapshotUpsert struct {
		*sql.UpdateSet
	}
)

// SetEntID sets the "ent_id" field.
func (u *SnapshotUpsert) SetEntID(v uuid.UUID) *SnapshotUpsert {
	u.Set(snapshot.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateEntID() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldEntID)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SnapshotUpsert) SetCreatedAt(v uint32) *SnapshotUpsert {
	u.Set(snapshot.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateCreatedAt() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *SnapshotUpsert) AddCreatedAt(v uint32) *SnapshotUpsert {
	u.Add(snapshot.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SnapshotUpsert) SetUpdatedAt(v uint32) *SnapshotUpsert {
	u.Set(snapshot.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateUpdatedAt() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SnapshotUpsert) AddUpdatedAt(v uint32) *SnapshotUpsert {
	u.Add(snapshot.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SnapshotUpsert) SetDeletedAt(v uint32) *SnapshotUpsert {
	u.Set(snapshot.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateDeletedAt() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *SnapshotUpsert) AddDeletedAt(v uint32) *SnapshotUpsert {
	u.Add(snapshot.FieldDeletedAt, v)
	return u
}

// SetIndex sets the "index" field.
func (u *SnapshotUpsert) SetIndex(v uint64) *SnapshotUpsert {
	u.Set(snapshot.FieldIndex, v)
	return u
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateIndex() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldIndex)
	return u
}

// AddIndex adds v to the "index" field.
func (u *SnapshotUpsert) AddIndex(v uint64) *SnapshotUpsert {
	u.Add(snapshot.FieldIndex, v)
	return u
}

// SetSnapshotCommP sets the "snapshot_comm_p" field.
func (u *SnapshotUpsert) SetSnapshotCommP(v string) *SnapshotUpsert {
	u.Set(snapshot.FieldSnapshotCommP, v)
	return u
}

// UpdateSnapshotCommP sets the "snapshot_comm_p" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateSnapshotCommP() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldSnapshotCommP)
	return u
}

// SetSnapshotRoot sets the "snapshot_root" field.
func (u *SnapshotUpsert) SetSnapshotRoot(v string) *SnapshotUpsert {
	u.Set(snapshot.FieldSnapshotRoot, v)
	return u
}

// UpdateSnapshotRoot sets the "snapshot_root" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateSnapshotRoot() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldSnapshotRoot)
	return u
}

// SetSnapshotURI sets the "snapshot_uri" field.
func (u *SnapshotUpsert) SetSnapshotURI(v string) *SnapshotUpsert {
	u.Set(snapshot.FieldSnapshotURI, v)
	return u
}

// UpdateSnapshotURI sets the "snapshot_uri" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateSnapshotURI() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldSnapshotURI)
	return u
}

// SetBackupState sets the "backup_state" field.
func (u *SnapshotUpsert) SetBackupState(v string) *SnapshotUpsert {
	u.Set(snapshot.FieldBackupState, v)
	return u
}

// UpdateBackupState sets the "backup_state" field to the value that was provided on create.
func (u *SnapshotUpsert) UpdateBackupState() *SnapshotUpsert {
	u.SetExcluded(snapshot.FieldBackupState)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(snapshot.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SnapshotUpsertOne) UpdateNewValues() *SnapshotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(snapshot.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SnapshotUpsertOne) Ignore() *SnapshotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SnapshotUpsertOne) DoNothing() *SnapshotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SnapshotCreate.OnConflict
// documentation for more info.
func (u *SnapshotUpsertOne) Update(set func(*SnapshotUpsert)) *SnapshotUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SnapshotUpsert{UpdateSet: update})
	}))
	return u
}

// SetEntID sets the "ent_id" field.
func (u *SnapshotUpsertOne) SetEntID(v uuid.UUID) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateEntID() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SnapshotUpsertOne) SetCreatedAt(v uint32) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *SnapshotUpsertOne) AddCreatedAt(v uint32) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateCreatedAt() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SnapshotUpsertOne) SetUpdatedAt(v uint32) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SnapshotUpsertOne) AddUpdatedAt(v uint32) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateUpdatedAt() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SnapshotUpsertOne) SetDeletedAt(v uint32) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *SnapshotUpsertOne) AddDeletedAt(v uint32) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateDeletedAt() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetIndex sets the "index" field.
func (u *SnapshotUpsertOne) SetIndex(v uint64) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *SnapshotUpsertOne) AddIndex(v uint64) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateIndex() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateIndex()
	})
}

// SetSnapshotCommP sets the "snapshot_comm_p" field.
func (u *SnapshotUpsertOne) SetSnapshotCommP(v string) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetSnapshotCommP(v)
	})
}

// UpdateSnapshotCommP sets the "snapshot_comm_p" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateSnapshotCommP() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateSnapshotCommP()
	})
}

// SetSnapshotRoot sets the "snapshot_root" field.
func (u *SnapshotUpsertOne) SetSnapshotRoot(v string) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetSnapshotRoot(v)
	})
}

// UpdateSnapshotRoot sets the "snapshot_root" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateSnapshotRoot() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateSnapshotRoot()
	})
}

// SetSnapshotURI sets the "snapshot_uri" field.
func (u *SnapshotUpsertOne) SetSnapshotURI(v string) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetSnapshotURI(v)
	})
}

// UpdateSnapshotURI sets the "snapshot_uri" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateSnapshotURI() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateSnapshotURI()
	})
}

// SetBackupState sets the "backup_state" field.
func (u *SnapshotUpsertOne) SetBackupState(v string) *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetBackupState(v)
	})
}

// UpdateBackupState sets the "backup_state" field to the value that was provided on create.
func (u *SnapshotUpsertOne) UpdateBackupState() *SnapshotUpsertOne {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateBackupState()
	})
}

// Exec executes the query.
func (u *SnapshotUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SnapshotCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SnapshotUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SnapshotUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SnapshotUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SnapshotCreateBulk is the builder for creating many Snapshot entities in bulk.
type SnapshotCreateBulk struct {
	config
	builders []*SnapshotCreate
	conflict []sql.ConflictOption
}

// Save creates the Snapshot entities in the database.
func (scb *SnapshotCreateBulk) Save(ctx context.Context) ([]*Snapshot, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Snapshot, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SnapshotMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SnapshotCreateBulk) SaveX(ctx context.Context) []*Snapshot {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SnapshotCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SnapshotCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Snapshot.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SnapshotUpsert) {
//			SetEntID(v+v).
//		}).
//		Exec(ctx)
func (scb *SnapshotCreateBulk) OnConflict(opts ...sql.ConflictOption) *SnapshotUpsertBulk {
	scb.conflict = opts
	return &SnapshotUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SnapshotCreateBulk) OnConflictColumns(columns ...string) *SnapshotUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SnapshotUpsertBulk{
		create: scb,
	}
}

// SnapshotUpsertBulk is the builder for "upsert"-ing
// a bulk of Snapshot nodes.
type SnapshotUpsertBulk struct {
	create *SnapshotCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(snapshot.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SnapshotUpsertBulk) UpdateNewValues() *SnapshotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(snapshot.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Snapshot.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SnapshotUpsertBulk) Ignore() *SnapshotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SnapshotUpsertBulk) DoNothing() *SnapshotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SnapshotCreateBulk.OnConflict
// documentation for more info.
func (u *SnapshotUpsertBulk) Update(set func(*SnapshotUpsert)) *SnapshotUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SnapshotUpsert{UpdateSet: update})
	}))
	return u
}

// SetEntID sets the "ent_id" field.
func (u *SnapshotUpsertBulk) SetEntID(v uuid.UUID) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateEntID() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SnapshotUpsertBulk) SetCreatedAt(v uint32) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *SnapshotUpsertBulk) AddCreatedAt(v uint32) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateCreatedAt() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SnapshotUpsertBulk) SetUpdatedAt(v uint32) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SnapshotUpsertBulk) AddUpdatedAt(v uint32) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateUpdatedAt() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SnapshotUpsertBulk) SetDeletedAt(v uint32) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *SnapshotUpsertBulk) AddDeletedAt(v uint32) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateDeletedAt() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetIndex sets the "index" field.
func (u *SnapshotUpsertBulk) SetIndex(v uint64) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *SnapshotUpsertBulk) AddIndex(v uint64) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateIndex() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateIndex()
	})
}

// SetSnapshotCommP sets the "snapshot_comm_p" field.
func (u *SnapshotUpsertBulk) SetSnapshotCommP(v string) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetSnapshotCommP(v)
	})
}

// UpdateSnapshotCommP sets the "snapshot_comm_p" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateSnapshotCommP() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateSnapshotCommP()
	})
}

// SetSnapshotRoot sets the "snapshot_root" field.
func (u *SnapshotUpsertBulk) SetSnapshotRoot(v string) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetSnapshotRoot(v)
	})
}

// UpdateSnapshotRoot sets the "snapshot_root" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateSnapshotRoot() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateSnapshotRoot()
	})
}

// SetSnapshotURI sets the "snapshot_uri" field.
func (u *SnapshotUpsertBulk) SetSnapshotURI(v string) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetSnapshotURI(v)
	})
}

// UpdateSnapshotURI sets the "snapshot_uri" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateSnapshotURI() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateSnapshotURI()
	})
}

// SetBackupState sets the "backup_state" field.
func (u *SnapshotUpsertBulk) SetBackupState(v string) *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.SetBackupState(v)
	})
}

// UpdateBackupState sets the "backup_state" field to the value that was provided on create.
func (u *SnapshotUpsertBulk) UpdateBackupState() *SnapshotUpsertBulk {
	return u.Update(func(s *SnapshotUpsert) {
		s.UpdateBackupState()
	})
}

// Exec executes the query.
func (u *SnapshotUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SnapshotCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SnapshotCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SnapshotUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

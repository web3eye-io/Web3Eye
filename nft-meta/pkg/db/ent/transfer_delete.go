// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/predicate"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/transfer"
)

// TransferDelete is the builder for deleting a Transfer entity.
type TransferDelete struct {
	config
	hooks    []Hook
	mutation *TransferMutation
}

// Where appends a list predicates to the TransferDelete builder.
func (td *TransferDelete) Where(ps ...predicate.Transfer) *TransferDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TransferDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(td.hooks) == 0 {
		affected, err = td.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransferMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			td.mutation = mutation
			affected, err = td.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(td.hooks) - 1; i >= 0; i-- {
			if td.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = td.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, td.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TransferDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TransferDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: transfer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: transfer.FieldID,
			},
		},
	}
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, td.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TransferDeleteOne is the builder for deleting a single Transfer entity.
type TransferDeleteOne struct {
	td *TransferDelete
}

// Exec executes the deletion query.
func (tdo *TransferDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{transfer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TransferDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}

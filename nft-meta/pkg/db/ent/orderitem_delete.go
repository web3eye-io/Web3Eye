// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/orderitem"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/predicate"
)

// OrderItemDelete is the builder for deleting a OrderItem entity.
type OrderItemDelete struct {
	config
	hooks    []Hook
	mutation *OrderItemMutation
}

// Where appends a list predicates to the OrderItemDelete builder.
func (oid *OrderItemDelete) Where(ps ...predicate.OrderItem) *OrderItemDelete {
	oid.mutation.Where(ps...)
	return oid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oid *OrderItemDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(oid.hooks) == 0 {
		affected, err = oid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oid.mutation = mutation
			affected, err = oid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oid.hooks) - 1; i >= 0; i-- {
			if oid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (oid *OrderItemDelete) ExecX(ctx context.Context) int {
	n, err := oid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oid *OrderItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: orderitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: orderitem.FieldID,
			},
		},
	}
	if ps := oid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, oid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// OrderItemDeleteOne is the builder for deleting a single OrderItem entity.
type OrderItemDeleteOne struct {
	oid *OrderItemDelete
}

// Exec executes the deletion query.
func (oido *OrderItemDeleteOne) Exec(ctx context.Context) error {
	n, err := oido.oid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oido *OrderItemDeleteOne) ExecX(ctx context.Context) {
	oido.oid.ExecX(ctx)
}

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"htp-platform/app/machine/service/internal/data/ent/machine"
	"htp-platform/app/machine/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MachineDelete is the builder for deleting a Machine entity.
type MachineDelete struct {
	config
	hooks    []Hook
	mutation *MachineMutation
}

// Where appends a list predicates to the MachineDelete builder.
func (md *MachineDelete) Where(ps ...predicate.Machine) *MachineDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MachineDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(md.hooks) == 0 {
		affected, err = md.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MachineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			md.mutation = mutation
			affected, err = md.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(md.hooks) - 1; i >= 0; i-- {
			if md.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = md.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, md.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MachineDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MachineDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: machine.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: machine.FieldID,
			},
		},
	}
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, md.driver, _spec)
}

// MachineDeleteOne is the builder for deleting a single Machine entity.
type MachineDeleteOne struct {
	md *MachineDelete
}

// Exec executes the deletion query.
func (mdo *MachineDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{machine.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MachineDeleteOne) ExecX(ctx context.Context) {
	mdo.md.ExecX(ctx)
}

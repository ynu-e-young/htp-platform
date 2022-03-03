// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"htp-platform/app/machine/service/internal/biz"
	"htp-platform/app/machine/service/internal/data/ent/cronjob"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CronJobCreate is the builder for creating a CronJob entity.
type CronJobCreate struct {
	config
	mutation *CronJobMutation
	hooks    []Hook
}

// SetMachineID sets the "machine_id" field.
func (cjc *CronJobCreate) SetMachineID(i int64) *CronJobCreate {
	cjc.mutation.SetMachineID(i)
	return cjc
}

// SetCheckName sets the "check_name" field.
func (cjc *CronJobCreate) SetCheckName(s string) *CronJobCreate {
	cjc.mutation.SetCheckName(s)
	return cjc
}

// SetCronString sets the "cron_string" field.
func (cjc *CronJobCreate) SetCronString(s string) *CronJobCreate {
	cjc.mutation.SetCronString(s)
	return cjc
}

// SetCoordinates sets the "coordinates" field.
func (cjc *CronJobCreate) SetCoordinates(bc []*biz.CheckCoordinate) *CronJobCreate {
	cjc.mutation.SetCoordinates(bc)
	return cjc
}

// SetCreatedAt sets the "created_at" field.
func (cjc *CronJobCreate) SetCreatedAt(t time.Time) *CronJobCreate {
	cjc.mutation.SetCreatedAt(t)
	return cjc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cjc *CronJobCreate) SetNillableCreatedAt(t *time.Time) *CronJobCreate {
	if t != nil {
		cjc.SetCreatedAt(*t)
	}
	return cjc
}

// SetUpdatedAt sets the "updated_at" field.
func (cjc *CronJobCreate) SetUpdatedAt(t time.Time) *CronJobCreate {
	cjc.mutation.SetUpdatedAt(t)
	return cjc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cjc *CronJobCreate) SetNillableUpdatedAt(t *time.Time) *CronJobCreate {
	if t != nil {
		cjc.SetUpdatedAt(*t)
	}
	return cjc
}

// SetID sets the "id" field.
func (cjc *CronJobCreate) SetID(i int64) *CronJobCreate {
	cjc.mutation.SetID(i)
	return cjc
}

// Mutation returns the CronJobMutation object of the builder.
func (cjc *CronJobCreate) Mutation() *CronJobMutation {
	return cjc.mutation
}

// Save creates the CronJob in the database.
func (cjc *CronJobCreate) Save(ctx context.Context) (*CronJob, error) {
	var (
		err  error
		node *CronJob
	)
	cjc.defaults()
	if len(cjc.hooks) == 0 {
		if err = cjc.check(); err != nil {
			return nil, err
		}
		node, err = cjc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CronJobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cjc.check(); err != nil {
				return nil, err
			}
			cjc.mutation = mutation
			if node, err = cjc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cjc.hooks) - 1; i >= 0; i-- {
			if cjc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cjc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cjc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cjc *CronJobCreate) SaveX(ctx context.Context) *CronJob {
	v, err := cjc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cjc *CronJobCreate) Exec(ctx context.Context) error {
	_, err := cjc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cjc *CronJobCreate) ExecX(ctx context.Context) {
	if err := cjc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cjc *CronJobCreate) defaults() {
	if _, ok := cjc.mutation.CreatedAt(); !ok {
		v := cronjob.DefaultCreatedAt()
		cjc.mutation.SetCreatedAt(v)
	}
	if _, ok := cjc.mutation.UpdatedAt(); !ok {
		v := cronjob.DefaultUpdatedAt()
		cjc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cjc *CronJobCreate) check() error {
	if _, ok := cjc.mutation.MachineID(); !ok {
		return &ValidationError{Name: "machine_id", err: errors.New(`ent: missing required field "CronJob.machine_id"`)}
	}
	if _, ok := cjc.mutation.CheckName(); !ok {
		return &ValidationError{Name: "check_name", err: errors.New(`ent: missing required field "CronJob.check_name"`)}
	}
	if _, ok := cjc.mutation.CronString(); !ok {
		return &ValidationError{Name: "cron_string", err: errors.New(`ent: missing required field "CronJob.cron_string"`)}
	}
	if _, ok := cjc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CronJob.created_at"`)}
	}
	if _, ok := cjc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CronJob.updated_at"`)}
	}
	return nil
}

func (cjc *CronJobCreate) sqlSave(ctx context.Context) (*CronJob, error) {
	_node, _spec := cjc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cjc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (cjc *CronJobCreate) createSpec() (*CronJob, *sqlgraph.CreateSpec) {
	var (
		_node = &CronJob{config: cjc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cronjob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: cronjob.FieldID,
			},
		}
	)
	if id, ok := cjc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cjc.mutation.MachineID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: cronjob.FieldMachineID,
		})
		_node.MachineID = value
	}
	if value, ok := cjc.mutation.CheckName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cronjob.FieldCheckName,
		})
		_node.CheckName = value
	}
	if value, ok := cjc.mutation.CronString(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cronjob.FieldCronString,
		})
		_node.CronString = value
	}
	if value, ok := cjc.mutation.Coordinates(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: cronjob.FieldCoordinates,
		})
		_node.Coordinates = value
	}
	if value, ok := cjc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cronjob.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cjc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cronjob.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// CronJobCreateBulk is the builder for creating many CronJob entities in bulk.
type CronJobCreateBulk struct {
	config
	builders []*CronJobCreate
}

// Save creates the CronJob entities in the database.
func (cjcb *CronJobCreateBulk) Save(ctx context.Context) ([]*CronJob, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cjcb.builders))
	nodes := make([]*CronJob, len(cjcb.builders))
	mutators := make([]Mutator, len(cjcb.builders))
	for i := range cjcb.builders {
		func(i int, root context.Context) {
			builder := cjcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CronJobMutation)
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
					_, err = mutators[i+1].Mutate(root, cjcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cjcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cjcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cjcb *CronJobCreateBulk) SaveX(ctx context.Context) []*CronJob {
	v, err := cjcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cjcb *CronJobCreateBulk) Exec(ctx context.Context) error {
	_, err := cjcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cjcb *CronJobCreateBulk) ExecX(ctx context.Context) {
	if err := cjcb.Exec(ctx); err != nil {
		panic(err)
	}
}
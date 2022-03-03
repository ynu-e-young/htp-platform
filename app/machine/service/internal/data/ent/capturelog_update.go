// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"htp-platform/app/machine/service/internal/data/ent/capturelog"
	"htp-platform/app/machine/service/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CaptureLogUpdate is the builder for updating CaptureLog entities.
type CaptureLogUpdate struct {
	config
	hooks    []Hook
	mutation *CaptureLogMutation
}

// Where appends a list predicates to the CaptureLogUpdate builder.
func (clu *CaptureLogUpdate) Where(ps ...predicate.CaptureLog) *CaptureLogUpdate {
	clu.mutation.Where(ps...)
	return clu
}

// SetMachineID sets the "machine_id" field.
func (clu *CaptureLogUpdate) SetMachineID(i int64) *CaptureLogUpdate {
	clu.mutation.ResetMachineID()
	clu.mutation.SetMachineID(i)
	return clu
}

// AddMachineID adds i to the "machine_id" field.
func (clu *CaptureLogUpdate) AddMachineID(i int64) *CaptureLogUpdate {
	clu.mutation.AddMachineID(i)
	return clu
}

// SetPixels sets the "pixels" field.
func (clu *CaptureLogUpdate) SetPixels(i int64) *CaptureLogUpdate {
	clu.mutation.ResetPixels()
	clu.mutation.SetPixels(i)
	return clu
}

// AddPixels adds i to the "pixels" field.
func (clu *CaptureLogUpdate) AddPixels(i int64) *CaptureLogUpdate {
	clu.mutation.AddPixels(i)
	return clu
}

// SetArea sets the "area" field.
func (clu *CaptureLogUpdate) SetArea(f float64) *CaptureLogUpdate {
	clu.mutation.ResetArea()
	clu.mutation.SetArea(f)
	return clu
}

// AddArea adds f to the "area" field.
func (clu *CaptureLogUpdate) AddArea(f float64) *CaptureLogUpdate {
	clu.mutation.AddArea(f)
	return clu
}

// SetSrcName sets the "src_name" field.
func (clu *CaptureLogUpdate) SetSrcName(s string) *CaptureLogUpdate {
	clu.mutation.SetSrcName(s)
	return clu
}

// SetProcName sets the "proc_name" field.
func (clu *CaptureLogUpdate) SetProcName(s string) *CaptureLogUpdate {
	clu.mutation.SetProcName(s)
	return clu
}

// SetSrcOssURL sets the "src_oss_url" field.
func (clu *CaptureLogUpdate) SetSrcOssURL(s string) *CaptureLogUpdate {
	clu.mutation.SetSrcOssURL(s)
	return clu
}

// SetProcOssURL sets the "proc_oss_url" field.
func (clu *CaptureLogUpdate) SetProcOssURL(s string) *CaptureLogUpdate {
	clu.mutation.SetProcOssURL(s)
	return clu
}

// SetCreatedAt sets the "created_at" field.
func (clu *CaptureLogUpdate) SetCreatedAt(t time.Time) *CaptureLogUpdate {
	clu.mutation.SetCreatedAt(t)
	return clu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (clu *CaptureLogUpdate) SetNillableCreatedAt(t *time.Time) *CaptureLogUpdate {
	if t != nil {
		clu.SetCreatedAt(*t)
	}
	return clu
}

// SetUpdatedAt sets the "updated_at" field.
func (clu *CaptureLogUpdate) SetUpdatedAt(t time.Time) *CaptureLogUpdate {
	clu.mutation.SetUpdatedAt(t)
	return clu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (clu *CaptureLogUpdate) SetNillableUpdatedAt(t *time.Time) *CaptureLogUpdate {
	if t != nil {
		clu.SetUpdatedAt(*t)
	}
	return clu
}

// Mutation returns the CaptureLogMutation object of the builder.
func (clu *CaptureLogUpdate) Mutation() *CaptureLogMutation {
	return clu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (clu *CaptureLogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(clu.hooks) == 0 {
		affected, err = clu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CaptureLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			clu.mutation = mutation
			affected, err = clu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(clu.hooks) - 1; i >= 0; i-- {
			if clu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = clu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, clu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (clu *CaptureLogUpdate) SaveX(ctx context.Context) int {
	affected, err := clu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (clu *CaptureLogUpdate) Exec(ctx context.Context) error {
	_, err := clu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (clu *CaptureLogUpdate) ExecX(ctx context.Context) {
	if err := clu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (clu *CaptureLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   capturelog.Table,
			Columns: capturelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: capturelog.FieldID,
			},
		},
	}
	if ps := clu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := clu.mutation.MachineID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: capturelog.FieldMachineID,
		})
	}
	if value, ok := clu.mutation.AddedMachineID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: capturelog.FieldMachineID,
		})
	}
	if value, ok := clu.mutation.Pixels(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: capturelog.FieldPixels,
		})
	}
	if value, ok := clu.mutation.AddedPixels(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: capturelog.FieldPixels,
		})
	}
	if value, ok := clu.mutation.Area(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: capturelog.FieldArea,
		})
	}
	if value, ok := clu.mutation.AddedArea(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: capturelog.FieldArea,
		})
	}
	if value, ok := clu.mutation.SrcName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capturelog.FieldSrcName,
		})
	}
	if value, ok := clu.mutation.ProcName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capturelog.FieldProcName,
		})
	}
	if value, ok := clu.mutation.SrcOssURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capturelog.FieldSrcOssURL,
		})
	}
	if value, ok := clu.mutation.ProcOssURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capturelog.FieldProcOssURL,
		})
	}
	if value, ok := clu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: capturelog.FieldCreatedAt,
		})
	}
	if value, ok := clu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: capturelog.FieldUpdatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, clu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{capturelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CaptureLogUpdateOne is the builder for updating a single CaptureLog entity.
type CaptureLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CaptureLogMutation
}

// SetMachineID sets the "machine_id" field.
func (cluo *CaptureLogUpdateOne) SetMachineID(i int64) *CaptureLogUpdateOne {
	cluo.mutation.ResetMachineID()
	cluo.mutation.SetMachineID(i)
	return cluo
}

// AddMachineID adds i to the "machine_id" field.
func (cluo *CaptureLogUpdateOne) AddMachineID(i int64) *CaptureLogUpdateOne {
	cluo.mutation.AddMachineID(i)
	return cluo
}

// SetPixels sets the "pixels" field.
func (cluo *CaptureLogUpdateOne) SetPixels(i int64) *CaptureLogUpdateOne {
	cluo.mutation.ResetPixels()
	cluo.mutation.SetPixels(i)
	return cluo
}

// AddPixels adds i to the "pixels" field.
func (cluo *CaptureLogUpdateOne) AddPixels(i int64) *CaptureLogUpdateOne {
	cluo.mutation.AddPixels(i)
	return cluo
}

// SetArea sets the "area" field.
func (cluo *CaptureLogUpdateOne) SetArea(f float64) *CaptureLogUpdateOne {
	cluo.mutation.ResetArea()
	cluo.mutation.SetArea(f)
	return cluo
}

// AddArea adds f to the "area" field.
func (cluo *CaptureLogUpdateOne) AddArea(f float64) *CaptureLogUpdateOne {
	cluo.mutation.AddArea(f)
	return cluo
}

// SetSrcName sets the "src_name" field.
func (cluo *CaptureLogUpdateOne) SetSrcName(s string) *CaptureLogUpdateOne {
	cluo.mutation.SetSrcName(s)
	return cluo
}

// SetProcName sets the "proc_name" field.
func (cluo *CaptureLogUpdateOne) SetProcName(s string) *CaptureLogUpdateOne {
	cluo.mutation.SetProcName(s)
	return cluo
}

// SetSrcOssURL sets the "src_oss_url" field.
func (cluo *CaptureLogUpdateOne) SetSrcOssURL(s string) *CaptureLogUpdateOne {
	cluo.mutation.SetSrcOssURL(s)
	return cluo
}

// SetProcOssURL sets the "proc_oss_url" field.
func (cluo *CaptureLogUpdateOne) SetProcOssURL(s string) *CaptureLogUpdateOne {
	cluo.mutation.SetProcOssURL(s)
	return cluo
}

// SetCreatedAt sets the "created_at" field.
func (cluo *CaptureLogUpdateOne) SetCreatedAt(t time.Time) *CaptureLogUpdateOne {
	cluo.mutation.SetCreatedAt(t)
	return cluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cluo *CaptureLogUpdateOne) SetNillableCreatedAt(t *time.Time) *CaptureLogUpdateOne {
	if t != nil {
		cluo.SetCreatedAt(*t)
	}
	return cluo
}

// SetUpdatedAt sets the "updated_at" field.
func (cluo *CaptureLogUpdateOne) SetUpdatedAt(t time.Time) *CaptureLogUpdateOne {
	cluo.mutation.SetUpdatedAt(t)
	return cluo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cluo *CaptureLogUpdateOne) SetNillableUpdatedAt(t *time.Time) *CaptureLogUpdateOne {
	if t != nil {
		cluo.SetUpdatedAt(*t)
	}
	return cluo
}

// Mutation returns the CaptureLogMutation object of the builder.
func (cluo *CaptureLogUpdateOne) Mutation() *CaptureLogMutation {
	return cluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cluo *CaptureLogUpdateOne) Select(field string, fields ...string) *CaptureLogUpdateOne {
	cluo.fields = append([]string{field}, fields...)
	return cluo
}

// Save executes the query and returns the updated CaptureLog entity.
func (cluo *CaptureLogUpdateOne) Save(ctx context.Context) (*CaptureLog, error) {
	var (
		err  error
		node *CaptureLog
	)
	if len(cluo.hooks) == 0 {
		node, err = cluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CaptureLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cluo.mutation = mutation
			node, err = cluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cluo.hooks) - 1; i >= 0; i-- {
			if cluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cluo *CaptureLogUpdateOne) SaveX(ctx context.Context) *CaptureLog {
	node, err := cluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cluo *CaptureLogUpdateOne) Exec(ctx context.Context) error {
	_, err := cluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cluo *CaptureLogUpdateOne) ExecX(ctx context.Context) {
	if err := cluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cluo *CaptureLogUpdateOne) sqlSave(ctx context.Context) (_node *CaptureLog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   capturelog.Table,
			Columns: capturelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: capturelog.FieldID,
			},
		},
	}
	id, ok := cluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CaptureLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, capturelog.FieldID)
		for _, f := range fields {
			if !capturelog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != capturelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cluo.mutation.MachineID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: capturelog.FieldMachineID,
		})
	}
	if value, ok := cluo.mutation.AddedMachineID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: capturelog.FieldMachineID,
		})
	}
	if value, ok := cluo.mutation.Pixels(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: capturelog.FieldPixels,
		})
	}
	if value, ok := cluo.mutation.AddedPixels(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: capturelog.FieldPixels,
		})
	}
	if value, ok := cluo.mutation.Area(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: capturelog.FieldArea,
		})
	}
	if value, ok := cluo.mutation.AddedArea(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: capturelog.FieldArea,
		})
	}
	if value, ok := cluo.mutation.SrcName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capturelog.FieldSrcName,
		})
	}
	if value, ok := cluo.mutation.ProcName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capturelog.FieldProcName,
		})
	}
	if value, ok := cluo.mutation.SrcOssURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capturelog.FieldSrcOssURL,
		})
	}
	if value, ok := cluo.mutation.ProcOssURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: capturelog.FieldProcOssURL,
		})
	}
	if value, ok := cluo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: capturelog.FieldCreatedAt,
		})
	}
	if value, ok := cluo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: capturelog.FieldUpdatedAt,
		})
	}
	_node = &CaptureLog{config: cluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{capturelog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
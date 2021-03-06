// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/timing"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// TimingCreate is the builder for creating a Timing entity.
type TimingCreate struct {
	config
	mutation *TimingMutation
	hooks    []Hook
}

// SetCategory sets the category field.
func (tc *TimingCreate) SetCategory(s string) *TimingCreate {
	tc.mutation.SetCategory(s)
	return tc
}

// SetTimingLabel sets the timing_label field.
func (tc *TimingCreate) SetTimingLabel(s string) *TimingCreate {
	tc.mutation.SetTimingLabel(s)
	return tc
}

// SetUnit sets the unit field.
func (tc *TimingCreate) SetUnit(s string) *TimingCreate {
	tc.mutation.SetUnit(s)
	return tc
}

// SetVariable sets the variable field.
func (tc *TimingCreate) SetVariable(s string) *TimingCreate {
	tc.mutation.SetVariable(s)
	return tc
}

// SetValue sets the value field.
func (tc *TimingCreate) SetValue(f float64) *TimingCreate {
	tc.mutation.SetValue(f)
	return tc
}

// AddEventIDs adds the events edge to Event by ids.
func (tc *TimingCreate) AddEventIDs(ids ...uuid.UUID) *TimingCreate {
	tc.mutation.AddEventIDs(ids...)
	return tc
}

// AddEvents adds the events edges to Event.
func (tc *TimingCreate) AddEvents(e ...*Event) *TimingCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tc.AddEventIDs(ids...)
}

// Mutation returns the TimingMutation object of the builder.
func (tc *TimingCreate) Mutation() *TimingMutation {
	return tc.mutation
}

// Save creates the Timing in the database.
func (tc *TimingCreate) Save(ctx context.Context) (*Timing, error) {
	if err := tc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Timing
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TimingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TimingCreate) SaveX(ctx context.Context) *Timing {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TimingCreate) preSave() error {
	if _, ok := tc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New("ent: missing required field \"category\"")}
	}
	if _, ok := tc.mutation.TimingLabel(); !ok {
		return &ValidationError{Name: "timing_label", err: errors.New("ent: missing required field \"timing_label\"")}
	}
	if _, ok := tc.mutation.Unit(); !ok {
		return &ValidationError{Name: "unit", err: errors.New("ent: missing required field \"unit\"")}
	}
	if _, ok := tc.mutation.Variable(); !ok {
		return &ValidationError{Name: "variable", err: errors.New("ent: missing required field \"variable\"")}
	}
	if _, ok := tc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New("ent: missing required field \"value\"")}
	}
	return nil
}

func (tc *TimingCreate) sqlSave(ctx context.Context) (*Timing, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TimingCreate) createSpec() (*Timing, *sqlgraph.CreateSpec) {
	var (
		t     = &Timing{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: timing.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: timing.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Category(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: timing.FieldCategory,
		})
		t.Category = value
	}
	if value, ok := tc.mutation.TimingLabel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: timing.FieldTimingLabel,
		})
		t.TimingLabel = value
	}
	if value, ok := tc.mutation.Unit(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: timing.FieldUnit,
		})
		t.Unit = value
	}
	if value, ok := tc.mutation.Variable(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: timing.FieldVariable,
		})
		t.Variable = value
	}
	if value, ok := tc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: timing.FieldValue,
		})
		t.Value = value
	}
	if nodes := tc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   timing.EventsTable,
			Columns: []string{timing.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return t, _spec
}

// TimingCreateBulk is the builder for creating a bulk of Timing entities.
type TimingCreateBulk struct {
	config
	builders []*TimingCreate
}

// Save creates the Timing entities in the database.
func (tcb *TimingCreateBulk) Save(ctx context.Context) ([]*Timing, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Timing, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*TimingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (tcb *TimingCreateBulk) SaveX(ctx context.Context) []*Timing {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

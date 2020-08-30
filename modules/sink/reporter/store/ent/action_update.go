// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/action"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// ActionUpdate is the builder for updating Action entities.
type ActionUpdate struct {
	config
	hooks      []Hook
	mutation   *ActionMutation
	predicates []predicate.Action
}

// Where adds a new predicate for the builder.
func (au *ActionUpdate) Where(ps ...predicate.Action) *ActionUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetAction sets the action field.
func (au *ActionUpdate) SetAction(s string) *ActionUpdate {
	au.mutation.SetAction(s)
	return au
}

// SetActionLabel sets the action_label field.
func (au *ActionUpdate) SetActionLabel(s string) *ActionUpdate {
	au.mutation.SetActionLabel(s)
	return au
}

// SetProperty sets the property field.
func (au *ActionUpdate) SetProperty(s string) *ActionUpdate {
	au.mutation.SetProperty(s)
	return au
}

// SetValue sets the value field.
func (au *ActionUpdate) SetValue(b []byte) *ActionUpdate {
	au.mutation.SetValue(b)
	return au
}

// SetEventID sets the event edge to Event by id.
func (au *ActionUpdate) SetEventID(id uuid.UUID) *ActionUpdate {
	au.mutation.SetEventID(id)
	return au
}

// SetNillableEventID sets the event edge to Event by id if the given value is not nil.
func (au *ActionUpdate) SetNillableEventID(id *uuid.UUID) *ActionUpdate {
	if id != nil {
		au = au.SetEventID(*id)
	}
	return au
}

// SetEvent sets the event edge to Event.
func (au *ActionUpdate) SetEvent(e *Event) *ActionUpdate {
	return au.SetEventID(e.ID)
}

// Mutation returns the ActionMutation object of the builder.
func (au *ActionUpdate) Mutation() *ActionMutation {
	return au.mutation
}

// ClearEvent clears the event edge to Event.
func (au *ActionUpdate) ClearEvent() *ActionUpdate {
	au.mutation.ClearEvent()
	return au
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *ActionUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *ActionUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ActionUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ActionUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ActionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   action.Table,
			Columns: action.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: action.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Action(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldAction,
		})
	}
	if value, ok := au.mutation.ActionLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldActionLabel,
		})
	}
	if value, ok := au.mutation.Property(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldProperty,
		})
	}
	if value, ok := au.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: action.FieldValue,
		})
	}
	if au.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   action.EventTable,
			Columns: []string{action.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   action.EventTable,
			Columns: []string{action.EventColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{action.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ActionUpdateOne is the builder for updating a single Action entity.
type ActionUpdateOne struct {
	config
	hooks    []Hook
	mutation *ActionMutation
}

// SetAction sets the action field.
func (auo *ActionUpdateOne) SetAction(s string) *ActionUpdateOne {
	auo.mutation.SetAction(s)
	return auo
}

// SetActionLabel sets the action_label field.
func (auo *ActionUpdateOne) SetActionLabel(s string) *ActionUpdateOne {
	auo.mutation.SetActionLabel(s)
	return auo
}

// SetProperty sets the property field.
func (auo *ActionUpdateOne) SetProperty(s string) *ActionUpdateOne {
	auo.mutation.SetProperty(s)
	return auo
}

// SetValue sets the value field.
func (auo *ActionUpdateOne) SetValue(b []byte) *ActionUpdateOne {
	auo.mutation.SetValue(b)
	return auo
}

// SetEventID sets the event edge to Event by id.
func (auo *ActionUpdateOne) SetEventID(id uuid.UUID) *ActionUpdateOne {
	auo.mutation.SetEventID(id)
	return auo
}

// SetNillableEventID sets the event edge to Event by id if the given value is not nil.
func (auo *ActionUpdateOne) SetNillableEventID(id *uuid.UUID) *ActionUpdateOne {
	if id != nil {
		auo = auo.SetEventID(*id)
	}
	return auo
}

// SetEvent sets the event edge to Event.
func (auo *ActionUpdateOne) SetEvent(e *Event) *ActionUpdateOne {
	return auo.SetEventID(e.ID)
}

// Mutation returns the ActionMutation object of the builder.
func (auo *ActionUpdateOne) Mutation() *ActionMutation {
	return auo.mutation
}

// ClearEvent clears the event edge to Event.
func (auo *ActionUpdateOne) ClearEvent() *ActionUpdateOne {
	auo.mutation.ClearEvent()
	return auo
}

// Save executes the query and returns the updated entity.
func (auo *ActionUpdateOne) Save(ctx context.Context) (*Action, error) {

	var (
		err  error
		node *Action
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ActionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ActionUpdateOne) SaveX(ctx context.Context) *Action {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *ActionUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ActionUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ActionUpdateOne) sqlSave(ctx context.Context) (a *Action, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   action.Table,
			Columns: action.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: action.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Action.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Action(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldAction,
		})
	}
	if value, ok := auo.mutation.ActionLabel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldActionLabel,
		})
	}
	if value, ok := auo.mutation.Property(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: action.FieldProperty,
		})
	}
	if value, ok := auo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: action.FieldValue,
		})
	}
	if auo.mutation.EventCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   action.EventTable,
			Columns: []string{action.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   action.EventTable,
			Columns: []string{action.EventColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	a = &Action{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{action.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}
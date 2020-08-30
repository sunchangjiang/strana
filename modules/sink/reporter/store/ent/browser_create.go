// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/browser"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// BrowserCreate is the builder for creating a Browser entity.
type BrowserCreate struct {
	config
	mutation *BrowserMutation
	hooks    []Hook
}

// SetName sets the name field.
func (bc *BrowserCreate) SetName(s string) *BrowserCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetVersion sets the version field.
func (bc *BrowserCreate) SetVersion(s string) *BrowserCreate {
	bc.mutation.SetVersion(s)
	return bc
}

// SetUseragent sets the useragent field.
func (bc *BrowserCreate) SetUseragent(s string) *BrowserCreate {
	bc.mutation.SetUseragent(s)
	return bc
}

// SetNillableUseragent sets the useragent field if the given value is not nil.
func (bc *BrowserCreate) SetNillableUseragent(s *string) *BrowserCreate {
	if s != nil {
		bc.SetUseragent(*s)
	}
	return bc
}

// AddEventIDs adds the events edge to Event by ids.
func (bc *BrowserCreate) AddEventIDs(ids ...uuid.UUID) *BrowserCreate {
	bc.mutation.AddEventIDs(ids...)
	return bc
}

// AddEvents adds the events edges to Event.
func (bc *BrowserCreate) AddEvents(e ...*Event) *BrowserCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return bc.AddEventIDs(ids...)
}

// Mutation returns the BrowserMutation object of the builder.
func (bc *BrowserCreate) Mutation() *BrowserMutation {
	return bc.mutation
}

// Save creates the Browser in the database.
func (bc *BrowserCreate) Save(ctx context.Context) (*Browser, error) {
	if err := bc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Browser
	)
	if len(bc.hooks) == 0 {
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BrowserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BrowserCreate) SaveX(ctx context.Context) *Browser {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BrowserCreate) preSave() error {
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := bc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New("ent: missing required field \"version\"")}
	}
	return nil
}

func (bc *BrowserCreate) sqlSave(ctx context.Context) (*Browser, error) {
	b, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	b.ID = int(id)
	return b, nil
}

func (bc *BrowserCreate) createSpec() (*Browser, *sqlgraph.CreateSpec) {
	var (
		b     = &Browser{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: browser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: browser.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: browser.FieldName,
		})
		b.Name = value
	}
	if value, ok := bc.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: browser.FieldVersion,
		})
		b.Version = value
	}
	if value, ok := bc.mutation.Useragent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: browser.FieldUseragent,
		})
		b.Useragent = value
	}
	if nodes := bc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   browser.EventsTable,
			Columns: []string{browser.EventsColumn},
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
	return b, _spec
}

// BrowserCreateBulk is the builder for creating a bulk of Browser entities.
type BrowserCreateBulk struct {
	config
	builders []*BrowserCreate
}

// Save creates the Browser entities in the database.
func (bcb *BrowserCreateBulk) Save(ctx context.Context) ([]*Browser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Browser, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*BrowserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (bcb *BrowserCreateBulk) SaveX(ctx context.Context) []*Browser {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
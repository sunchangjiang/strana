// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/page"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// PageCreate is the builder for creating a Page entity.
type PageCreate struct {
	config
	mutation *PageMutation
	hooks    []Hook
}

// SetHostname sets the hostname field.
func (pc *PageCreate) SetHostname(s string) *PageCreate {
	pc.mutation.SetHostname(s)
	return pc
}

// SetPath sets the path field.
func (pc *PageCreate) SetPath(s string) *PageCreate {
	pc.mutation.SetPath(s)
	return pc
}

// SetReferrer sets the referrer field.
func (pc *PageCreate) SetReferrer(s string) *PageCreate {
	pc.mutation.SetReferrer(s)
	return pc
}

// SetNillableReferrer sets the referrer field if the given value is not nil.
func (pc *PageCreate) SetNillableReferrer(s *string) *PageCreate {
	if s != nil {
		pc.SetReferrer(*s)
	}
	return pc
}

// SetSearch sets the search field.
func (pc *PageCreate) SetSearch(s string) *PageCreate {
	pc.mutation.SetSearch(s)
	return pc
}

// SetNillableSearch sets the search field if the given value is not nil.
func (pc *PageCreate) SetNillableSearch(s *string) *PageCreate {
	if s != nil {
		pc.SetSearch(*s)
	}
	return pc
}

// SetTitle sets the title field.
func (pc *PageCreate) SetTitle(s string) *PageCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetNillableTitle sets the title field if the given value is not nil.
func (pc *PageCreate) SetNillableTitle(s *string) *PageCreate {
	if s != nil {
		pc.SetTitle(*s)
	}
	return pc
}

// SetHash sets the hash field.
func (pc *PageCreate) SetHash(s string) *PageCreate {
	pc.mutation.SetHash(s)
	return pc
}

// SetNillableHash sets the hash field if the given value is not nil.
func (pc *PageCreate) SetNillableHash(s *string) *PageCreate {
	if s != nil {
		pc.SetHash(*s)
	}
	return pc
}

// AddEventIDs adds the events edge to Event by ids.
func (pc *PageCreate) AddEventIDs(ids ...uuid.UUID) *PageCreate {
	pc.mutation.AddEventIDs(ids...)
	return pc
}

// AddEvents adds the events edges to Event.
func (pc *PageCreate) AddEvents(e ...*Event) *PageCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pc.AddEventIDs(ids...)
}

// Mutation returns the PageMutation object of the builder.
func (pc *PageCreate) Mutation() *PageMutation {
	return pc.mutation
}

// Save creates the Page in the database.
func (pc *PageCreate) Save(ctx context.Context) (*Page, error) {
	if err := pc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Page
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PageCreate) SaveX(ctx context.Context) *Page {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PageCreate) preSave() error {
	if _, ok := pc.mutation.Hostname(); !ok {
		return &ValidationError{Name: "hostname", err: errors.New("ent: missing required field \"hostname\"")}
	}
	if _, ok := pc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New("ent: missing required field \"path\"")}
	}
	return nil
}

func (pc *PageCreate) sqlSave(ctx context.Context) (*Page, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PageCreate) createSpec() (*Page, *sqlgraph.CreateSpec) {
	var (
		pa    = &Page{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: page.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: page.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Hostname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: page.FieldHostname,
		})
		pa.Hostname = value
	}
	if value, ok := pc.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: page.FieldPath,
		})
		pa.Path = value
	}
	if value, ok := pc.mutation.Referrer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: page.FieldReferrer,
		})
		pa.Referrer = value
	}
	if value, ok := pc.mutation.Search(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: page.FieldSearch,
		})
		pa.Search = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: page.FieldTitle,
		})
		pa.Title = value
	}
	if value, ok := pc.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: page.FieldHash,
		})
		pa.Hash = value
	}
	if nodes := pc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   page.EventsTable,
			Columns: []string{page.EventsColumn},
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
	return pa, _spec
}

// PageCreateBulk is the builder for creating a bulk of Page entities.
type PageCreateBulk struct {
	config
	builders []*PageCreate
}

// Save creates the Page entities in the database.
func (pcb *PageCreateBulk) Save(ctx context.Context) ([]*Page, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Page, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*PageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (pcb *PageCreateBulk) SaveX(ctx context.Context) []*Page {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/alias"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/predicate"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/user"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// AliasQuery is the builder for querying Alias entities.
type AliasQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Alias
	// eager-loading edges.
	withEvent *EventQuery
	withUser  *UserQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (aq *AliasQuery) Where(ps ...predicate.Alias) *AliasQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AliasQuery) Limit(limit int) *AliasQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AliasQuery) Offset(offset int) *AliasQuery {
	aq.offset = &offset
	return aq
}

// Order adds an order step to the query.
func (aq *AliasQuery) Order(o ...OrderFunc) *AliasQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryEvent chains the current query on the event edge.
func (aq *AliasQuery) QueryEvent() *EventQuery {
	query := &EventQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(alias.Table, alias.FieldID, aq.sqlQuery()),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, alias.EventTable, alias.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the user edge.
func (aq *AliasQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(alias.Table, alias.FieldID, aq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, alias.UserTable, alias.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Alias entity in the query. Returns *NotFoundError when no alias was found.
func (aq *AliasQuery) First(ctx context.Context) (*Alias, error) {
	as, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(as) == 0 {
		return nil, &NotFoundError{alias.Label}
	}
	return as[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AliasQuery) FirstX(ctx context.Context) *Alias {
	a, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return a
}

// FirstID returns the first Alias id in the query. Returns *NotFoundError when no id was found.
func (aq *AliasQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{alias.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (aq *AliasQuery) FirstXID(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Alias entity in the query, returns an error if not exactly one entity was returned.
func (aq *AliasQuery) Only(ctx context.Context) (*Alias, error) {
	as, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(as) {
	case 1:
		return as[0], nil
	case 0:
		return nil, &NotFoundError{alias.Label}
	default:
		return nil, &NotSingularError{alias.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AliasQuery) OnlyX(ctx context.Context) *Alias {
	a, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// OnlyID returns the only Alias id in the query, returns an error if not exactly one id was returned.
func (aq *AliasQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{alias.Label}
	default:
		err = &NotSingularError{alias.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AliasQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AliasSlice.
func (aq *AliasQuery) All(ctx context.Context) ([]*Alias, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AliasQuery) AllX(ctx context.Context) []*Alias {
	as, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return as
}

// IDs executes the query and returns a list of Alias ids.
func (aq *AliasQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aq.Select(alias.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AliasQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AliasQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AliasQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AliasQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AliasQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AliasQuery) Clone() *AliasQuery {
	return &AliasQuery{
		config:     aq.config,
		limit:      aq.limit,
		offset:     aq.offset,
		order:      append([]OrderFunc{}, aq.order...),
		unique:     append([]string{}, aq.unique...),
		predicates: append([]predicate.Alias{}, aq.predicates...),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

//  WithEvent tells the query-builder to eager-loads the nodes that are connected to
// the "event" edge. The optional arguments used to configure the query builder of the edge.
func (aq *AliasQuery) WithEvent(opts ...func(*EventQuery)) *AliasQuery {
	query := &EventQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withEvent = query
	return aq
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "user" edge. The optional arguments used to configure the query builder of the edge.
func (aq *AliasQuery) WithUser(opts ...func(*UserQuery)) *AliasQuery {
	query := &UserQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withUser = query
	return aq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		From string `json:"from,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Alias.Query().
//		GroupBy(alias.FieldFrom).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aq *AliasQuery) GroupBy(field string, fields ...string) *AliasGroupBy {
	group := &AliasGroupBy{config: aq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		From string `json:"from,omitempty"`
//	}
//
//	client.Alias.Query().
//		Select(alias.FieldFrom).
//		Scan(ctx, &v)
//
func (aq *AliasQuery) Select(field string, fields ...string) *AliasSelect {
	selector := &AliasSelect{config: aq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(), nil
	}
	return selector
}

func (aq *AliasQuery) prepareQuery(ctx context.Context) error {
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AliasQuery) sqlAll(ctx context.Context) ([]*Alias, error) {
	var (
		nodes       = []*Alias{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withEvent != nil,
			aq.withUser != nil,
		}
	)
	if aq.withEvent != nil || aq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, alias.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Alias{config: aq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aq.withEvent; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Alias)
		for i := range nodes {
			if fk := nodes[i].event_alias; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(event.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "event_alias" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Event = n
			}
		}
	}

	if query := aq.withUser; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*Alias)
		for i := range nodes {
			if fk := nodes[i].user_aliases; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_aliases" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (aq *AliasQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AliasQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (aq *AliasQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   alias.Table,
			Columns: alias.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: alias.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AliasQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(alias.Table)
	selector := builder.Select(t1.Columns(alias.Columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(alias.Columns...)...)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AliasGroupBy is the builder for group-by Alias entities.
type AliasGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AliasGroupBy) Aggregate(fns ...AggregateFunc) *AliasGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scan the result into the given value.
func (agb *AliasGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (agb *AliasGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := agb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (agb *AliasGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AliasGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (agb *AliasGroupBy) StringsX(ctx context.Context) []string {
	v, err := agb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (agb *AliasGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = agb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{alias.Label}
	default:
		err = fmt.Errorf("ent: AliasGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (agb *AliasGroupBy) StringX(ctx context.Context) string {
	v, err := agb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (agb *AliasGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AliasGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (agb *AliasGroupBy) IntsX(ctx context.Context) []int {
	v, err := agb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (agb *AliasGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = agb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{alias.Label}
	default:
		err = fmt.Errorf("ent: AliasGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (agb *AliasGroupBy) IntX(ctx context.Context) int {
	v, err := agb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (agb *AliasGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AliasGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (agb *AliasGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := agb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (agb *AliasGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = agb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{alias.Label}
	default:
		err = fmt.Errorf("ent: AliasGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (agb *AliasGroupBy) Float64X(ctx context.Context) float64 {
	v, err := agb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (agb *AliasGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: AliasGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (agb *AliasGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := agb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (agb *AliasGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = agb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{alias.Label}
	default:
		err = fmt.Errorf("ent: AliasGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (agb *AliasGroupBy) BoolX(ctx context.Context) bool {
	v, err := agb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (agb *AliasGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := agb.sqlQuery().Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AliasGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql
	columns := make([]string, 0, len(agb.fields)+len(agb.fns))
	columns = append(columns, agb.fields...)
	for _, fn := range agb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(agb.fields...)
}

// AliasSelect is the builder for select fields of Alias entities.
type AliasSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (as *AliasSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := as.path(ctx)
	if err != nil {
		return err
	}
	as.sql = query
	return as.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (as *AliasSelect) ScanX(ctx context.Context, v interface{}) {
	if err := as.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (as *AliasSelect) Strings(ctx context.Context) ([]string, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AliasSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (as *AliasSelect) StringsX(ctx context.Context) []string {
	v, err := as.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (as *AliasSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = as.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{alias.Label}
	default:
		err = fmt.Errorf("ent: AliasSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (as *AliasSelect) StringX(ctx context.Context) string {
	v, err := as.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (as *AliasSelect) Ints(ctx context.Context) ([]int, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AliasSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (as *AliasSelect) IntsX(ctx context.Context) []int {
	v, err := as.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (as *AliasSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = as.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{alias.Label}
	default:
		err = fmt.Errorf("ent: AliasSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (as *AliasSelect) IntX(ctx context.Context) int {
	v, err := as.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (as *AliasSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AliasSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (as *AliasSelect) Float64sX(ctx context.Context) []float64 {
	v, err := as.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (as *AliasSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = as.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{alias.Label}
	default:
		err = fmt.Errorf("ent: AliasSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (as *AliasSelect) Float64X(ctx context.Context) float64 {
	v, err := as.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (as *AliasSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: AliasSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (as *AliasSelect) BoolsX(ctx context.Context) []bool {
	v, err := as.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (as *AliasSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = as.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{alias.Label}
	default:
		err = fmt.Errorf("ent: AliasSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (as *AliasSelect) BoolX(ctx context.Context) bool {
	v, err := as.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (as *AliasSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sqlQuery().Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (as *AliasSelect) sqlQuery() sql.Querier {
	selector := as.sql
	selector.Select(selector.Columns(as.fields...)...)
	return selector
}

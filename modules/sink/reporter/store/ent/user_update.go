// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/blushft/strana/modules/sink/reporter/store/ent/alias"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/event"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/group"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/predicate"
	"github.com/blushft/strana/modules/sink/reporter/store/ent/user"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks      []Hook
	mutation   *UserMutation
	predicates []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetIsAnonymous sets the is_anonymous field.
func (uu *UserUpdate) SetIsAnonymous(b bool) *UserUpdate {
	uu.mutation.SetIsAnonymous(b)
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the name field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// ClearName clears the value of name.
func (uu *UserUpdate) ClearName() *UserUpdate {
	uu.mutation.ClearName()
	return uu
}

// SetTitle sets the title field.
func (uu *UserUpdate) SetTitle(s string) *UserUpdate {
	uu.mutation.SetTitle(s)
	return uu
}

// SetNillableTitle sets the title field if the given value is not nil.
func (uu *UserUpdate) SetNillableTitle(s *string) *UserUpdate {
	if s != nil {
		uu.SetTitle(*s)
	}
	return uu
}

// ClearTitle clears the value of title.
func (uu *UserUpdate) ClearTitle() *UserUpdate {
	uu.mutation.ClearTitle()
	return uu
}

// SetFirstName sets the first_name field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetNillableFirstName sets the first_name field if the given value is not nil.
func (uu *UserUpdate) SetNillableFirstName(s *string) *UserUpdate {
	if s != nil {
		uu.SetFirstName(*s)
	}
	return uu
}

// ClearFirstName clears the value of first_name.
func (uu *UserUpdate) ClearFirstName() *UserUpdate {
	uu.mutation.ClearFirstName()
	return uu
}

// SetLastName sets the last_name field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetNillableLastName sets the last_name field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastName(s *string) *UserUpdate {
	if s != nil {
		uu.SetLastName(*s)
	}
	return uu
}

// ClearLastName clears the value of last_name.
func (uu *UserUpdate) ClearLastName() *UserUpdate {
	uu.mutation.ClearLastName()
	return uu
}

// SetEmail sets the email field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the email field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of email.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetUsername sets the username field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the username field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// ClearUsername clears the value of username.
func (uu *UserUpdate) ClearUsername() *UserUpdate {
	uu.mutation.ClearUsername()
	return uu
}

// SetAge sets the age field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.mutation.ResetAge()
	uu.mutation.SetAge(i)
	return uu
}

// SetNillableAge sets the age field if the given value is not nil.
func (uu *UserUpdate) SetNillableAge(i *int) *UserUpdate {
	if i != nil {
		uu.SetAge(*i)
	}
	return uu
}

// AddAge adds i to age.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	uu.mutation.AddAge(i)
	return uu
}

// ClearAge clears the value of age.
func (uu *UserUpdate) ClearAge() *UserUpdate {
	uu.mutation.ClearAge()
	return uu
}

// SetBirthday sets the birthday field.
func (uu *UserUpdate) SetBirthday(t time.Time) *UserUpdate {
	uu.mutation.SetBirthday(t)
	return uu
}

// SetNillableBirthday sets the birthday field if the given value is not nil.
func (uu *UserUpdate) SetNillableBirthday(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetBirthday(*t)
	}
	return uu
}

// ClearBirthday clears the value of birthday.
func (uu *UserUpdate) ClearBirthday() *UserUpdate {
	uu.mutation.ClearBirthday()
	return uu
}

// SetGender sets the gender field.
func (uu *UserUpdate) SetGender(u user.Gender) *UserUpdate {
	uu.mutation.SetGender(u)
	return uu
}

// SetNillableGender sets the gender field if the given value is not nil.
func (uu *UserUpdate) SetNillableGender(u *user.Gender) *UserUpdate {
	if u != nil {
		uu.SetGender(*u)
	}
	return uu
}

// ClearGender clears the value of gender.
func (uu *UserUpdate) ClearGender() *UserUpdate {
	uu.mutation.ClearGender()
	return uu
}

// SetPhone sets the phone field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// ClearPhone clears the value of phone.
func (uu *UserUpdate) ClearPhone() *UserUpdate {
	uu.mutation.ClearPhone()
	return uu
}

// SetWebsite sets the website field.
func (uu *UserUpdate) SetWebsite(s string) *UserUpdate {
	uu.mutation.SetWebsite(s)
	return uu
}

// SetNillableWebsite sets the website field if the given value is not nil.
func (uu *UserUpdate) SetNillableWebsite(s *string) *UserUpdate {
	if s != nil {
		uu.SetWebsite(*s)
	}
	return uu
}

// ClearWebsite clears the value of website.
func (uu *UserUpdate) ClearWebsite() *UserUpdate {
	uu.mutation.ClearWebsite()
	return uu
}

// SetExtra sets the extra field.
func (uu *UserUpdate) SetExtra(m map[string]interface{}) *UserUpdate {
	uu.mutation.SetExtra(m)
	return uu
}

// ClearExtra clears the value of extra.
func (uu *UserUpdate) ClearExtra() *UserUpdate {
	uu.mutation.ClearExtra()
	return uu
}

// AddAliasIDs adds the aliases edge to Alias by ids.
func (uu *UserUpdate) AddAliasIDs(ids ...int) *UserUpdate {
	uu.mutation.AddAliasIDs(ids...)
	return uu
}

// AddAliases adds the aliases edges to Alias.
func (uu *UserUpdate) AddAliases(a ...*Alias) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddAliasIDs(ids...)
}

// AddEventIDs adds the events edge to Event by ids.
func (uu *UserUpdate) AddEventIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddEventIDs(ids...)
	return uu
}

// AddEvents adds the events edges to Event.
func (uu *UserUpdate) AddEvents(e ...*Event) *UserUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.AddEventIDs(ids...)
}

// AddGroupIDs adds the groups edge to Group by ids.
func (uu *UserUpdate) AddGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.AddGroupIDs(ids...)
	return uu
}

// AddGroups adds the groups edges to Group.
func (uu *UserUpdate) AddGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGroupIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// RemoveAliasIDs removes the aliases edge to Alias by ids.
func (uu *UserUpdate) RemoveAliasIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveAliasIDs(ids...)
	return uu
}

// RemoveAliases removes aliases edges to Alias.
func (uu *UserUpdate) RemoveAliases(a ...*Alias) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveAliasIDs(ids...)
}

// RemoveEventIDs removes the events edge to Event by ids.
func (uu *UserUpdate) RemoveEventIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveEventIDs(ids...)
	return uu
}

// RemoveEvents removes events edges to Event.
func (uu *UserUpdate) RemoveEvents(e ...*Event) *UserUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.RemoveEventIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (uu *UserUpdate) RemoveGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveGroupIDs(ids...)
	return uu
}

// RemoveGroups removes groups edges to Group.
func (uu *UserUpdate) RemoveGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := uu.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return 0, &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.IsAnonymous(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldIsAnonymous,
		})
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if uu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTitle,
		})
	}
	if uu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldTitle,
		})
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFirstName,
		})
	}
	if uu.mutation.FirstNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldFirstName,
		})
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastName,
		})
	}
	if uu.mutation.LastNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldLastName,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if uu.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if uu.mutation.UsernameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if uu.mutation.AgeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldAge,
		})
	}
	if value, ok := uu.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldBirthday,
		})
	}
	if uu.mutation.BirthdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldBirthday,
		})
	}
	if value, ok := uu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldGender,
		})
	}
	if uu.mutation.GenderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldGender,
		})
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
	}
	if uu.mutation.PhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPhone,
		})
	}
	if value, ok := uu.mutation.Website(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldWebsite,
		})
	}
	if uu.mutation.WebsiteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldWebsite,
		})
	}
	if value, ok := uu.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldExtra,
		})
	}
	if uu.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldExtra,
		})
	}
	if nodes := uu.mutation.RemovedAliasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AliasesTable,
			Columns: []string{user.AliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AliasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AliasesTable,
			Columns: []string{user.AliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uu.mutation.RemovedEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.EventsTable,
			Columns: []string{user.EventsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.EventsTable,
			Columns: []string{user.EventsColumn},
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
	if nodes := uu.mutation.RemovedGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// SetIsAnonymous sets the is_anonymous field.
func (uuo *UserUpdateOne) SetIsAnonymous(b bool) *UserUpdateOne {
	uuo.mutation.SetIsAnonymous(b)
	return uuo
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the name field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// ClearName clears the value of name.
func (uuo *UserUpdateOne) ClearName() *UserUpdateOne {
	uuo.mutation.ClearName()
	return uuo
}

// SetTitle sets the title field.
func (uuo *UserUpdateOne) SetTitle(s string) *UserUpdateOne {
	uuo.mutation.SetTitle(s)
	return uuo
}

// SetNillableTitle sets the title field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTitle(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTitle(*s)
	}
	return uuo
}

// ClearTitle clears the value of title.
func (uuo *UserUpdateOne) ClearTitle() *UserUpdateOne {
	uuo.mutation.ClearTitle()
	return uuo
}

// SetFirstName sets the first_name field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetNillableFirstName sets the first_name field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFirstName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFirstName(*s)
	}
	return uuo
}

// ClearFirstName clears the value of first_name.
func (uuo *UserUpdateOne) ClearFirstName() *UserUpdateOne {
	uuo.mutation.ClearFirstName()
	return uuo
}

// SetLastName sets the last_name field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetNillableLastName sets the last_name field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLastName(*s)
	}
	return uuo
}

// ClearLastName clears the value of last_name.
func (uuo *UserUpdateOne) ClearLastName() *UserUpdateOne {
	uuo.mutation.ClearLastName()
	return uuo
}

// SetEmail sets the email field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the email field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of email.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetUsername sets the username field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the username field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// ClearUsername clears the value of username.
func (uuo *UserUpdateOne) ClearUsername() *UserUpdateOne {
	uuo.mutation.ClearUsername()
	return uuo
}

// SetAge sets the age field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.mutation.ResetAge()
	uuo.mutation.SetAge(i)
	return uuo
}

// SetNillableAge sets the age field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAge(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetAge(*i)
	}
	return uuo
}

// AddAge adds i to age.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	uuo.mutation.AddAge(i)
	return uuo
}

// ClearAge clears the value of age.
func (uuo *UserUpdateOne) ClearAge() *UserUpdateOne {
	uuo.mutation.ClearAge()
	return uuo
}

// SetBirthday sets the birthday field.
func (uuo *UserUpdateOne) SetBirthday(t time.Time) *UserUpdateOne {
	uuo.mutation.SetBirthday(t)
	return uuo
}

// SetNillableBirthday sets the birthday field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBirthday(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetBirthday(*t)
	}
	return uuo
}

// ClearBirthday clears the value of birthday.
func (uuo *UserUpdateOne) ClearBirthday() *UserUpdateOne {
	uuo.mutation.ClearBirthday()
	return uuo
}

// SetGender sets the gender field.
func (uuo *UserUpdateOne) SetGender(u user.Gender) *UserUpdateOne {
	uuo.mutation.SetGender(u)
	return uuo
}

// SetNillableGender sets the gender field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGender(u *user.Gender) *UserUpdateOne {
	if u != nil {
		uuo.SetGender(*u)
	}
	return uuo
}

// ClearGender clears the value of gender.
func (uuo *UserUpdateOne) ClearGender() *UserUpdateOne {
	uuo.mutation.ClearGender()
	return uuo
}

// SetPhone sets the phone field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// ClearPhone clears the value of phone.
func (uuo *UserUpdateOne) ClearPhone() *UserUpdateOne {
	uuo.mutation.ClearPhone()
	return uuo
}

// SetWebsite sets the website field.
func (uuo *UserUpdateOne) SetWebsite(s string) *UserUpdateOne {
	uuo.mutation.SetWebsite(s)
	return uuo
}

// SetNillableWebsite sets the website field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWebsite(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetWebsite(*s)
	}
	return uuo
}

// ClearWebsite clears the value of website.
func (uuo *UserUpdateOne) ClearWebsite() *UserUpdateOne {
	uuo.mutation.ClearWebsite()
	return uuo
}

// SetExtra sets the extra field.
func (uuo *UserUpdateOne) SetExtra(m map[string]interface{}) *UserUpdateOne {
	uuo.mutation.SetExtra(m)
	return uuo
}

// ClearExtra clears the value of extra.
func (uuo *UserUpdateOne) ClearExtra() *UserUpdateOne {
	uuo.mutation.ClearExtra()
	return uuo
}

// AddAliasIDs adds the aliases edge to Alias by ids.
func (uuo *UserUpdateOne) AddAliasIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddAliasIDs(ids...)
	return uuo
}

// AddAliases adds the aliases edges to Alias.
func (uuo *UserUpdateOne) AddAliases(a ...*Alias) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddAliasIDs(ids...)
}

// AddEventIDs adds the events edge to Event by ids.
func (uuo *UserUpdateOne) AddEventIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddEventIDs(ids...)
	return uuo
}

// AddEvents adds the events edges to Event.
func (uuo *UserUpdateOne) AddEvents(e ...*Event) *UserUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.AddEventIDs(ids...)
}

// AddGroupIDs adds the groups edge to Group by ids.
func (uuo *UserUpdateOne) AddGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddGroupIDs(ids...)
	return uuo
}

// AddGroups adds the groups edges to Group.
func (uuo *UserUpdateOne) AddGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGroupIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// RemoveAliasIDs removes the aliases edge to Alias by ids.
func (uuo *UserUpdateOne) RemoveAliasIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveAliasIDs(ids...)
	return uuo
}

// RemoveAliases removes aliases edges to Alias.
func (uuo *UserUpdateOne) RemoveAliases(a ...*Alias) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveAliasIDs(ids...)
}

// RemoveEventIDs removes the events edge to Event by ids.
func (uuo *UserUpdateOne) RemoveEventIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveEventIDs(ids...)
	return uuo
}

// RemoveEvents removes events edges to Event.
func (uuo *UserUpdateOne) RemoveEvents(e ...*Event) *UserUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.RemoveEventIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (uuo *UserUpdateOne) RemoveGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveGroupIDs(ids...)
	return uuo
}

// RemoveGroups removes groups edges to Group.
func (uuo *UserUpdateOne) RemoveGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if v, ok := uuo.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return nil, &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}

	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uuo.mutation.IsAnonymous(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldIsAnonymous,
		})
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if uuo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTitle,
		})
	}
	if uuo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldTitle,
		})
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFirstName,
		})
	}
	if uuo.mutation.FirstNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldFirstName,
		})
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastName,
		})
	}
	if uuo.mutation.LastNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldLastName,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if uuo.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if uuo.mutation.UsernameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if value, ok := uuo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
	}
	if uuo.mutation.AgeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: user.FieldAge,
		})
	}
	if value, ok := uuo.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldBirthday,
		})
	}
	if uuo.mutation.BirthdayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: user.FieldBirthday,
		})
	}
	if value, ok := uuo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldGender,
		})
	}
	if uuo.mutation.GenderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: user.FieldGender,
		})
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPhone,
		})
	}
	if uuo.mutation.PhoneCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldPhone,
		})
	}
	if value, ok := uuo.mutation.Website(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldWebsite,
		})
	}
	if uuo.mutation.WebsiteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldWebsite,
		})
	}
	if value, ok := uuo.mutation.Extra(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: user.FieldExtra,
		})
	}
	if uuo.mutation.ExtraCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: user.FieldExtra,
		})
	}
	if nodes := uuo.mutation.RemovedAliasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AliasesTable,
			Columns: []string{user.AliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AliasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AliasesTable,
			Columns: []string{user.AliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: alias.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := uuo.mutation.RemovedEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.EventsTable,
			Columns: []string{user.EventsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.EventsTable,
			Columns: []string{user.EventsColumn},
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
	if nodes := uuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	u = &User{config: uuo.config}
	_spec.Assign = u.assignValues
	_spec.ScanValues = u.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return u, nil
}

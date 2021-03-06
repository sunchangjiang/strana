package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.Bool("new_user"),
		field.Bool("is_unique"),
		field.Bool("is_bounce"),
		field.Bool("is_finished"),
		field.Int("duration").Optional(),
		field.Time("started_at"),
		field.Time("finished_at").Nillable().Optional(),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).Ref("session"),
	}
}

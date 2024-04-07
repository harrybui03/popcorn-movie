package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Room schema
type Room struct {
	ent.Schema
}

// Field of Room
func (Room) Field() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Unique(),
		field.Int("room_number"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of Room
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("theater", Theater.Type).Ref("rooms").Unique(),
	}
}

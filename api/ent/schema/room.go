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

// Fields of Room
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Unique(),
		field.Int("room_number").Positive(),
		field.UUID("theater_id", uuid.UUID{}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of Room
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("theater", Theater.Type).Field("theater_id").Ref("rooms").Required().Unique(),
		edge.To("seats", Seat.Type),
		edge.To("showTimes", ShowTime.Type),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Seat schema
type Seat struct {
	ent.Schema
}

// Fields of Seat
func (Seat) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Unique(),
		field.String("seat_number"),
		field.UUID("room_id", uuid.UUID{}),
		field.Enum("category").Values("STANDARD", "DOUBLE"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of Seat
func (Seat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).Field("room_id").Ref("seats").Required().Unique(),
		edge.To("tickets", Ticket.Type),
	}
}

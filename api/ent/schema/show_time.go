package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type ShowTime struct {
	ent.Schema
}

func (ShowTime) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Unique(),
		field.Time("start_at"),
		field.Time("end_at"),
		field.UUID("movie_id", uuid.UUID{}),
		field.UUID("room_id", uuid.UUID{}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (ShowTime) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).Field("room_id").Ref("showTimes").Required().Unique(),
		edge.From("movie", Movie.Type).Field("movie_id").Ref("showTimes").Required().Unique(),
		edge.To("tickets", Ticket.Type),
	}
}

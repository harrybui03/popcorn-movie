package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	ent.Schema
}

func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Unique(),
		field.Bool("is_booked").Default(false),
		field.Float("price").Positive(),
		field.UUID("transaction_id", uuid.UUID{}),
		field.UUID("seat_id", uuid.UUID{}),
		field.UUID("show_time_id", uuid.UUID{}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("transaction", Transaction.Type).Field("transaction_id").Ref("tickets").Required().Unique(),
		edge.From("seat", Seat.Type).Field("seat_id").Ref("tickets").Required().Unique(),
		edge.From("show_time", ShowTime.Type).Field("show_time_id").Ref("tickets").Required().Unique(),
	}
}

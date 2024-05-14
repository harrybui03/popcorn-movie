package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ent.Schema
}

func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Unique(),
		field.Float("total"),
		field.UUID("user_id", uuid.UUID{}),
		field.Int("code").Unique().Optional(),
		field.Enum("status").Values("PENDING", "PAID", "CANCEL").Default("PENDING"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Field("user_id").Ref("transactions").Required().Unique(),
		edge.To("tickets", Ticket.Type),
		edge.To("food_order_lines", FoodOrderLine.Type),
	}
}

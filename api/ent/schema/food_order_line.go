package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type FoodOrderLine struct {
	ent.Schema
}

func (FoodOrderLine) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Unique(),
		field.Int("quantity"),
		field.UUID("food_id", uuid.UUID{}),
		field.UUID("transaction_id", uuid.UUID{}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (FoodOrderLine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("food", Food.Type).Field("food_id").Ref("food_order_lines").Required().Unique(),
		edge.From("transaction", Transaction.Type).Field("transaction_id").Ref("food_order_lines").Required().Unique(),
	}
}

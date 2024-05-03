package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Food struct {
	ent.Schema
}

func (Food) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Unique(),
		field.String("name"),
		field.Float("price"),
		field.String("image"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (Food) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("food_order_lines", FoodOrderLine.Type),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Movie struct {
	ent.Schema
}

func (Movie) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().Unique(),
		field.String("title"),
		field.String("genre"),
		field.Enum("status").Values("UPCOMING", "ONGOING", "OVER"),
		field.String("language"),
		field.String("director"),
		field.String("cast"),
		field.String("poster"),
		field.String("rated"),
		field.Int("duration"),
		field.String("trailer"),
		field.Time("opening_day"),
		field.String("story"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Movie) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("showTimes", ShowTime.Type),
		edge.To("comments", Comment.Type),
	}
}

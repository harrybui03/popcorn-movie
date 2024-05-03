package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ent.Schema
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Unique(),
		field.UUID("movie_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.Float("rating").Positive(),
		field.String("description"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Field("user_id").Ref("comments").Required().Unique(),
		edge.From("movie", Movie.Type).Field("movie_id").Ref("comments").Required().Unique(),
	}
}

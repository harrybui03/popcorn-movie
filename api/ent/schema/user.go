package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Unique(),
		field.String("displayname").MaxLen(255),
		field.String("email").MaxLen(255),
		field.String("password").MaxLen(255),
		field.Bool("is_locked").Default(false),
		field.Enum("role").Values("CUSTOMER", "STAFF", "TICKET_MANAGER", "ADMIN").Default("CUSTOMER"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transactions", Transaction.Type),
		edge.To("comments", Comment.Type),
		edge.To("reset_password", ResetPassword.Type),
	}
}

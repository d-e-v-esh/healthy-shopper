package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("Username").NotEmpty().Unique(),
		field.String("EmailAddress").NotEmpty().Unique(),
		field.String("Password").NotEmpty(),
		field.String("FirstName").NotEmpty(),
		field.String("LastName").NotEmpty(),
		field.Time("CreatedAt").Default(time.Now).Immutable(),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

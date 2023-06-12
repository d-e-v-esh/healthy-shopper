package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserAddress struct {
	ent.Schema
}

func (UserAddress) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Unique(),
		field.Int("address_id").Unique(),
		field.Bool("is_default").Default(false),
	}
}

// Edges of the User.
func (UserAddress) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("user", User.Type).Ref("user_address").Unique().Field("user_id").Required(),
		edge.To("address", Address.Type).Field("address_id").Unique().Required(),
	}
}

func (UserAddress) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

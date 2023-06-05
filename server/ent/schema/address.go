package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.Int("address_id").Unique(),
		field.String("phone_number").MaxLen(20).NotEmpty(),
		field.String("address_line1").MaxLen(100).NotEmpty(),
		field.String("address_line2").MaxLen(100).Optional(),
		field.String("city").MaxLen(50).NotEmpty(),
		field.String("state").MaxLen(50).NotEmpty(),
		field.String("country").MaxLen(50).NotEmpty(),
		field.String("postal_code").MaxLen(20).NotEmpty(),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_address", UserAddress.Type).Ref("address"),
	}
}

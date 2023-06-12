package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderStatus holds the schema definition for the Address entity.
type OrderStatus struct {
	ent.Schema
}

// Fields of the OrderStatus.
func (OrderStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("status").MaxLen(20).NotEmpty(),
	}
}

// Edges of the OrderStatus.
func (OrderStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("shop_order", ShopOrder.Type),
	}
}

func (OrderStatus) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

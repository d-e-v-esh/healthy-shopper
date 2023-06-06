package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ShippingMethod holds the schema definition for the Address entity.
type ShippingMethod struct {
	ent.Schema
}

// Fields of the ShippingMethod.
func (ShippingMethod) Fields() []ent.Field {
	return []ent.Field{
		field.String("shipping_method").MaxLen(25).NotEmpty(),
		field.Float("shipping_cost").Positive(),
	}
}

// Edges of the ShippingMethod.
func (ShippingMethod) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("shop_order", ShopOrder.Type),
	}
}

func (ShippingMethod) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

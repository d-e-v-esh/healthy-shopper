package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ShoppingCart holds the schema definition for the ShoppingCart entity.
type ShoppingCart struct {
	ent.Schema
}

// Fields of the ShoppingCart.
func (ShoppingCart) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Unique(),
	}
}

// Edges of the ShoppingCart.
func (ShoppingCart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("shopping_cart").Unique().Field("user_id").Required(),
		edge.To("shopping_cart_item", ShoppingCartItem.Type).Unique(), // not required as shopping cart can be empty
	}
}

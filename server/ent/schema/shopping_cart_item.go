package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ShoppingCartItem holds the schema definition for the ShoppingCartItem entity.
type ShoppingCartItem struct {
	ent.Schema
}

// Fields of the ShoppingCartItem.
func (ShoppingCartItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("shopping_cart_id").Unique(),
		field.Int("product_item_id").Unique(),
		field.Int("quantity").Positive(),
	}
}

// Edges of the ShoppingCartItem.
func (ShoppingCartItem) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("shopping_cart", ShoppingCart.Type).Ref("shopping_cart_item").Unique().Field("shopping_cart_id").Required(),
		edge.From("product_item", ProductItem.Type).Ref("shopping_cart_item").Unique().Field("product_item_id").Required(),
	}
}

func (ShoppingCartItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

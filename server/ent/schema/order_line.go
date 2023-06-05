package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OrderLine holds the schema definition for the Address entity.
type OrderLine struct {
	ent.Schema
}

// Fields of the OrderLine.
func (OrderLine) Fields() []ent.Field {
	return []ent.Field{
		field.Int("product_item_id").Unique(),
		field.Int("shop_order_id").Unique(),
		field.Int("quantity").Positive(),
		field.Float("price").Positive(),
	}
}

// Edges of the OrderLine.
func (OrderLine) Edges() []ent.Edge {
	return []ent.Edge{}
}

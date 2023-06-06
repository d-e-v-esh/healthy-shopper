package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
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
	return []ent.Edge{
		edge.From("product_item", ProductItem.Type).Ref("order_line").Unique().Field("product_item_id").Required(),

		edge.To("user_review", UserReview.Type),
	}
}

func (OrderLine) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

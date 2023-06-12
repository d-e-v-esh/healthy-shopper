package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductItem holds the schema definition for the Product entity.
type ProductItem struct {
	ent.Schema
}

// Fields of the ProductItem.
func (ProductItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("product_id").Unique(),
		field.String("stock_keeping_unit").MaxLen(500).NotEmpty(),
		field.Int("quantity_in_stock").Positive(),
		field.String("product_image").MaxLen(500).NotEmpty(),
		field.Float32("price").Positive(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ProductItem.
func (ProductItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("product_item").Unique().Field("product_id").Required(),

		edge.To("order_line", OrderLine.Type),
		edge.To("shopping_cart_item", ShoppingCartItem.Type),
	}
}

func (ProductItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

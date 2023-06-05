package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).NotEmpty(),
		field.String("description").MaxLen(500).NotEmpty(),
		field.String("product_image").MaxLen(500).NotEmpty(),
		field.Int("ingredients_list_id").Unique().Optional(),
		field.Int("product_category_id").Unique().Optional(),
		field.Int("nutritional_information_id").Unique().Optional(),
		field.Int("promotion_id").Unique().Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product_item", ProductItem.Type).Unique(),
		edge.To("promotion", Promotion.Type).Field("promotion_id").Unique(),
	}
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}

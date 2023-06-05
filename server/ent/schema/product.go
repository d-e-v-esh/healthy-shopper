package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("product_id").Unique(),
		field.String("name").MaxLen(50).NotEmpty(),
		field.String("description").MaxLen(500).NotEmpty(),
		field.String("product_image").MaxLen(500).NotEmpty(),
		field.Int("product_category_id").Positive(),
		field.Int("ingredients_list_id").Positive(),
		field.Int("nutritional_information_id").Positive(),
		field.Int("promotion_id").Positive(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}

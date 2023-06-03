package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").NotEmpty(),
		field.String("Description").NotEmpty(),
		field.String("ProductImage").NotEmpty(),
		field.Int("ProductCategoryID").Positive(),
		field.Int("IngredientsListID").Positive(),
		field.Int("NutritionalInformationID").Positive(),
		field.Int("PromotionID").Positive(),

		field.Time("CreatedAt").Default(time.Now).Immutable(),
		field.Time("UpdatedAt").Default(time.Now).UpdateDefault(time.Now)}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}

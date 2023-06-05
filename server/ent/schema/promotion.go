package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Promotion holds the schema definition for the Promotion entity.
type Promotion struct {
	ent.Schema
}

// Fields of the Promotion.
func (Promotion) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).NotEmpty(),
		field.String("description").MaxLen(500).NotEmpty(),
		field.Int("discount_percentage").Range(0, 100),
		field.Time("start_date"),
		field.Time("end_date"),
	}
}

// Edges of the Promotion.
func (Promotion) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("product", Product.Type).Ref("promotion").Required(), // This can't be unique because there no field associated with this edge
	}
}

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IngredientsTable holds the schema definition for the Promotion entity.
type IngredientsTable struct {
	ent.Schema
}

// Fields of the IngredientsTable.
func (IngredientsTable) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).NotEmpty(),
		field.String("description").MaxLen(500).NotEmpty(),
	}
}

// Edges of the IngredientsTable.
func (IngredientsTable) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("product", Product.Type).Ref("ingredients_table").Required(), // This can't be unique because there no field associated with this edge
	}
}

func (IngredientsTable) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NutritionalInformation holds the schema definition for the NutritionalInformation entity.
type NutritionalInformation struct {
	ent.Schema
}

// Fields of the NutritionalInformation.
func (NutritionalInformation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("nutritional_information_table_id").Unique(),
		field.Float("n_value").Positive(),
		field.String("n_measurement_unit").MaxLen(6).NotEmpty(),
	}
}

// Edges of the NutritionalInformation.
func (NutritionalInformation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("nutritional_information").Required(),

		edge.To("nutritional_information_table", NutritionalInformationTable.Type).Unique().Field("nutritional_information_table_id").Required(), // Required here means that the nutritional information must have a nutritional information table
	}
}

func (NutritionalInformation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

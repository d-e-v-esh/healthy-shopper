package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NutritionalInformationTable holds the schema definition for the NutritionalInformation entity.
type NutritionalInformationTable struct {
	ent.Schema
}

// Fields of the NutritionalInformationTable.
func (NutritionalInformationTable) Fields() []ent.Field {
	return []ent.Field{
		field.String("parameter").MaxLen(30).NotEmpty(),
		field.Float("value").Positive(),
		field.String("measurement_unit").MaxLen(6).NotEmpty(),
	}
}

// Edges of the NutritionalInformationTable.
func (NutritionalInformationTable) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nutritional_information", NutritionalInformation.Type).Ref("nutritional_information_table"),
	}
}

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserReview holds the schema definition for the UserReview entity.
type UserReview struct {
	ent.Schema
}

// Fields of the UserReview.
func (UserReview) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_review_id").Unique(),
		field.Int("user_id").Unique(),
		field.Int("ordered_product_id").Unique(),
		field.Int("rating").Range(1, 5),
		field.String("review").MaxLen(500).NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserReview.
func (UserReview) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_review").Unique().Field("user_id").Required(),
	}
}

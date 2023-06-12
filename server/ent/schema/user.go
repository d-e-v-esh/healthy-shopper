package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").MaxLen(20).NotEmpty().Unique(),
		field.String("email_address").MaxLen(60).NotEmpty().Unique(),
		field.String("password").MaxLen(100).NotEmpty(),
		field.String("first_name").MaxLen(50).NotEmpty(),
		field.String("last_name").MaxLen(50).NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_address", UserAddress.Type),
		edge.To("user_review", UserReview.Type),
		edge.To("shopping_cart", ShoppingCart.Type),
		edge.To("shop_order", ShopOrder.Type),
	}
}

package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ShopOrder holds the schema definition for the ShopOrder entity.
type ShopOrder struct {
	ent.Schema
}

// Fields of the ShopOrder.
func (ShopOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Time("order_date_and_time").Default(time.Now),
		field.String("payment_method").MaxLen(20).NotEmpty(),
		field.Float("total_price").Positive(),
		field.Int("user_id").Unique(),
		field.Int("shipping_address_id").Unique(),
		field.Int("shipping_method_id").Unique(),
		field.Int("order_status_id").Unique(),
	}
}

// Edges of the ShopOrder.
func (ShopOrder) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("user", User.Type).Ref("shop_order").Field("user_id").Unique().Required(),

		edge.From("shipping_method", ShippingMethod.Type).Ref("shop_order").Field("shipping_method_id").Unique().Required(),

		edge.From("order_status", OrderStatus.Type).Ref("shop_order").Field("order_status_id").Unique().Required(),

		edge.To("shipping_address", ShippingAddress.Type).Field("shipping_address_id").Unique().Required(),
	}
}

func (ShopOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

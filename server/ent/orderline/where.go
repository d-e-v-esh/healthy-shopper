// Code generated by ent, DO NOT EDIT.

package orderline

import (
	"healthyshopper/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldLTE(FieldID, id))
}

// ProductItemID applies equality check predicate on the "product_item_id" field. It's identical to ProductItemIDEQ.
func ProductItemID(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldEQ(FieldProductItemID, v))
}

// ShopOrderID applies equality check predicate on the "shop_order_id" field. It's identical to ShopOrderIDEQ.
func ShopOrderID(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldEQ(FieldShopOrderID, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldEQ(FieldQuantity, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldEQ(FieldPrice, v))
}

// ProductItemIDEQ applies the EQ predicate on the "product_item_id" field.
func ProductItemIDEQ(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldEQ(FieldProductItemID, v))
}

// ProductItemIDNEQ applies the NEQ predicate on the "product_item_id" field.
func ProductItemIDNEQ(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldNEQ(FieldProductItemID, v))
}

// ProductItemIDIn applies the In predicate on the "product_item_id" field.
func ProductItemIDIn(vs ...int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldIn(FieldProductItemID, vs...))
}

// ProductItemIDNotIn applies the NotIn predicate on the "product_item_id" field.
func ProductItemIDNotIn(vs ...int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldNotIn(FieldProductItemID, vs...))
}

// ShopOrderIDEQ applies the EQ predicate on the "shop_order_id" field.
func ShopOrderIDEQ(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldEQ(FieldShopOrderID, v))
}

// ShopOrderIDNEQ applies the NEQ predicate on the "shop_order_id" field.
func ShopOrderIDNEQ(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldNEQ(FieldShopOrderID, v))
}

// ShopOrderIDIn applies the In predicate on the "shop_order_id" field.
func ShopOrderIDIn(vs ...int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldIn(FieldShopOrderID, vs...))
}

// ShopOrderIDNotIn applies the NotIn predicate on the "shop_order_id" field.
func ShopOrderIDNotIn(vs ...int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldNotIn(FieldShopOrderID, vs...))
}

// ShopOrderIDGT applies the GT predicate on the "shop_order_id" field.
func ShopOrderIDGT(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldGT(FieldShopOrderID, v))
}

// ShopOrderIDGTE applies the GTE predicate on the "shop_order_id" field.
func ShopOrderIDGTE(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldGTE(FieldShopOrderID, v))
}

// ShopOrderIDLT applies the LT predicate on the "shop_order_id" field.
func ShopOrderIDLT(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldLT(FieldShopOrderID, v))
}

// ShopOrderIDLTE applies the LTE predicate on the "shop_order_id" field.
func ShopOrderIDLTE(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldLTE(FieldShopOrderID, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldLTE(FieldQuantity, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.OrderLine {
	return predicate.OrderLine(sql.FieldLTE(FieldPrice, v))
}

// HasProductItem applies the HasEdge predicate on the "product_item" edge.
func HasProductItem() predicate.OrderLine {
	return predicate.OrderLine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductItemTable, ProductItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductItemWith applies the HasEdge predicate on the "product_item" edge with a given conditions (other predicates).
func HasProductItemWith(preds ...predicate.ProductItem) predicate.OrderLine {
	return predicate.OrderLine(func(s *sql.Selector) {
		step := newProductItemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserReview applies the HasEdge predicate on the "user_review" edge.
func HasUserReview() predicate.OrderLine {
	return predicate.OrderLine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserReviewTable, UserReviewColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserReviewWith applies the HasEdge predicate on the "user_review" edge with a given conditions (other predicates).
func HasUserReviewWith(preds ...predicate.UserReview) predicate.OrderLine {
	return predicate.OrderLine(func(s *sql.Selector) {
		step := newUserReviewStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderLine) predicate.OrderLine {
	return predicate.OrderLine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderLine) predicate.OrderLine {
	return predicate.OrderLine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderLine) predicate.OrderLine {
	return predicate.OrderLine(func(s *sql.Selector) {
		p(s.Not())
	})
}

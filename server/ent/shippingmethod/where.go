// Code generated by ent, DO NOT EDIT.

package shippingmethod

import (
	"healthyshopper/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldLTE(FieldID, id))
}

// ShippingMethod applies equality check predicate on the "shipping_method" field. It's identical to ShippingMethodEQ.
func ShippingMethod(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldEQ(FieldShippingMethod, v))
}

// ShippingCost applies equality check predicate on the "shipping_cost" field. It's identical to ShippingCostEQ.
func ShippingCost(v float64) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldEQ(FieldShippingCost, v))
}

// ShippingMethodEQ applies the EQ predicate on the "shipping_method" field.
func ShippingMethodEQ(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldEQ(FieldShippingMethod, v))
}

// ShippingMethodNEQ applies the NEQ predicate on the "shipping_method" field.
func ShippingMethodNEQ(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldNEQ(FieldShippingMethod, v))
}

// ShippingMethodIn applies the In predicate on the "shipping_method" field.
func ShippingMethodIn(vs ...string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldIn(FieldShippingMethod, vs...))
}

// ShippingMethodNotIn applies the NotIn predicate on the "shipping_method" field.
func ShippingMethodNotIn(vs ...string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldNotIn(FieldShippingMethod, vs...))
}

// ShippingMethodGT applies the GT predicate on the "shipping_method" field.
func ShippingMethodGT(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldGT(FieldShippingMethod, v))
}

// ShippingMethodGTE applies the GTE predicate on the "shipping_method" field.
func ShippingMethodGTE(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldGTE(FieldShippingMethod, v))
}

// ShippingMethodLT applies the LT predicate on the "shipping_method" field.
func ShippingMethodLT(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldLT(FieldShippingMethod, v))
}

// ShippingMethodLTE applies the LTE predicate on the "shipping_method" field.
func ShippingMethodLTE(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldLTE(FieldShippingMethod, v))
}

// ShippingMethodContains applies the Contains predicate on the "shipping_method" field.
func ShippingMethodContains(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldContains(FieldShippingMethod, v))
}

// ShippingMethodHasPrefix applies the HasPrefix predicate on the "shipping_method" field.
func ShippingMethodHasPrefix(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldHasPrefix(FieldShippingMethod, v))
}

// ShippingMethodHasSuffix applies the HasSuffix predicate on the "shipping_method" field.
func ShippingMethodHasSuffix(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldHasSuffix(FieldShippingMethod, v))
}

// ShippingMethodEqualFold applies the EqualFold predicate on the "shipping_method" field.
func ShippingMethodEqualFold(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldEqualFold(FieldShippingMethod, v))
}

// ShippingMethodContainsFold applies the ContainsFold predicate on the "shipping_method" field.
func ShippingMethodContainsFold(v string) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldContainsFold(FieldShippingMethod, v))
}

// ShippingCostEQ applies the EQ predicate on the "shipping_cost" field.
func ShippingCostEQ(v float64) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldEQ(FieldShippingCost, v))
}

// ShippingCostNEQ applies the NEQ predicate on the "shipping_cost" field.
func ShippingCostNEQ(v float64) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldNEQ(FieldShippingCost, v))
}

// ShippingCostIn applies the In predicate on the "shipping_cost" field.
func ShippingCostIn(vs ...float64) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldIn(FieldShippingCost, vs...))
}

// ShippingCostNotIn applies the NotIn predicate on the "shipping_cost" field.
func ShippingCostNotIn(vs ...float64) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldNotIn(FieldShippingCost, vs...))
}

// ShippingCostGT applies the GT predicate on the "shipping_cost" field.
func ShippingCostGT(v float64) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldGT(FieldShippingCost, v))
}

// ShippingCostGTE applies the GTE predicate on the "shipping_cost" field.
func ShippingCostGTE(v float64) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldGTE(FieldShippingCost, v))
}

// ShippingCostLT applies the LT predicate on the "shipping_cost" field.
func ShippingCostLT(v float64) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldLT(FieldShippingCost, v))
}

// ShippingCostLTE applies the LTE predicate on the "shipping_cost" field.
func ShippingCostLTE(v float64) predicate.ShippingMethod {
	return predicate.ShippingMethod(sql.FieldLTE(FieldShippingCost, v))
}

// HasShopOrder applies the HasEdge predicate on the "shop_order" edge.
func HasShopOrder() predicate.ShippingMethod {
	return predicate.ShippingMethod(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ShopOrderTable, ShopOrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShopOrderWith applies the HasEdge predicate on the "shop_order" edge with a given conditions (other predicates).
func HasShopOrderWith(preds ...predicate.ShopOrder) predicate.ShippingMethod {
	return predicate.ShippingMethod(func(s *sql.Selector) {
		step := newShopOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ShippingMethod) predicate.ShippingMethod {
	return predicate.ShippingMethod(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ShippingMethod) predicate.ShippingMethod {
	return predicate.ShippingMethod(func(s *sql.Selector) {
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
func Not(p predicate.ShippingMethod) predicate.ShippingMethod {
	return predicate.ShippingMethod(func(s *sql.Selector) {
		p(s.Not())
	})
}

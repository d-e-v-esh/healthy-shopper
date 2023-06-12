// Code generated by ent, DO NOT EDIT.

package orderstatus

import (
	"healthyshopper/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldLTE(FieldID, id))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldEQ(FieldStatus, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.OrderStatus {
	return predicate.OrderStatus(sql.FieldContainsFold(FieldStatus, v))
}

// HasShopOrder applies the HasEdge predicate on the "shop_order" edge.
func HasShopOrder() predicate.OrderStatus {
	return predicate.OrderStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ShopOrderTable, ShopOrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShopOrderWith applies the HasEdge predicate on the "shop_order" edge with a given conditions (other predicates).
func HasShopOrderWith(preds ...predicate.ShopOrder) predicate.OrderStatus {
	return predicate.OrderStatus(func(s *sql.Selector) {
		step := newShopOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderStatus) predicate.OrderStatus {
	return predicate.OrderStatus(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderStatus) predicate.OrderStatus {
	return predicate.OrderStatus(func(s *sql.Selector) {
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
func Not(p predicate.OrderStatus) predicate.OrderStatus {
	return predicate.OrderStatus(func(s *sql.Selector) {
		p(s.Not())
	})
}

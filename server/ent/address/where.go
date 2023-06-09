// Code generated by ent, DO NOT EDIT.

package address

import (
	"healthyshopper/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldID, id))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldPhoneNumber, v))
}

// AddressLine1 applies equality check predicate on the "address_line1" field. It's identical to AddressLine1EQ.
func AddressLine1(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldAddressLine1, v))
}

// AddressLine2 applies equality check predicate on the "address_line2" field. It's identical to AddressLine2EQ.
func AddressLine2(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldAddressLine2, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldState, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// PostalCode applies equality check predicate on the "postal_code" field. It's identical to PostalCodeEQ.
func PostalCode(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldPostalCode, v))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// AddressLine1EQ applies the EQ predicate on the "address_line1" field.
func AddressLine1EQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldAddressLine1, v))
}

// AddressLine1NEQ applies the NEQ predicate on the "address_line1" field.
func AddressLine1NEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldAddressLine1, v))
}

// AddressLine1In applies the In predicate on the "address_line1" field.
func AddressLine1In(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldAddressLine1, vs...))
}

// AddressLine1NotIn applies the NotIn predicate on the "address_line1" field.
func AddressLine1NotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldAddressLine1, vs...))
}

// AddressLine1GT applies the GT predicate on the "address_line1" field.
func AddressLine1GT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldAddressLine1, v))
}

// AddressLine1GTE applies the GTE predicate on the "address_line1" field.
func AddressLine1GTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldAddressLine1, v))
}

// AddressLine1LT applies the LT predicate on the "address_line1" field.
func AddressLine1LT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldAddressLine1, v))
}

// AddressLine1LTE applies the LTE predicate on the "address_line1" field.
func AddressLine1LTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldAddressLine1, v))
}

// AddressLine1Contains applies the Contains predicate on the "address_line1" field.
func AddressLine1Contains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldAddressLine1, v))
}

// AddressLine1HasPrefix applies the HasPrefix predicate on the "address_line1" field.
func AddressLine1HasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldAddressLine1, v))
}

// AddressLine1HasSuffix applies the HasSuffix predicate on the "address_line1" field.
func AddressLine1HasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldAddressLine1, v))
}

// AddressLine1EqualFold applies the EqualFold predicate on the "address_line1" field.
func AddressLine1EqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldAddressLine1, v))
}

// AddressLine1ContainsFold applies the ContainsFold predicate on the "address_line1" field.
func AddressLine1ContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldAddressLine1, v))
}

// AddressLine2EQ applies the EQ predicate on the "address_line2" field.
func AddressLine2EQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldAddressLine2, v))
}

// AddressLine2NEQ applies the NEQ predicate on the "address_line2" field.
func AddressLine2NEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldAddressLine2, v))
}

// AddressLine2In applies the In predicate on the "address_line2" field.
func AddressLine2In(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldAddressLine2, vs...))
}

// AddressLine2NotIn applies the NotIn predicate on the "address_line2" field.
func AddressLine2NotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldAddressLine2, vs...))
}

// AddressLine2GT applies the GT predicate on the "address_line2" field.
func AddressLine2GT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldAddressLine2, v))
}

// AddressLine2GTE applies the GTE predicate on the "address_line2" field.
func AddressLine2GTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldAddressLine2, v))
}

// AddressLine2LT applies the LT predicate on the "address_line2" field.
func AddressLine2LT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldAddressLine2, v))
}

// AddressLine2LTE applies the LTE predicate on the "address_line2" field.
func AddressLine2LTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldAddressLine2, v))
}

// AddressLine2Contains applies the Contains predicate on the "address_line2" field.
func AddressLine2Contains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldAddressLine2, v))
}

// AddressLine2HasPrefix applies the HasPrefix predicate on the "address_line2" field.
func AddressLine2HasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldAddressLine2, v))
}

// AddressLine2HasSuffix applies the HasSuffix predicate on the "address_line2" field.
func AddressLine2HasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldAddressLine2, v))
}

// AddressLine2IsNil applies the IsNil predicate on the "address_line2" field.
func AddressLine2IsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldAddressLine2))
}

// AddressLine2NotNil applies the NotNil predicate on the "address_line2" field.
func AddressLine2NotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldAddressLine2))
}

// AddressLine2EqualFold applies the EqualFold predicate on the "address_line2" field.
func AddressLine2EqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldAddressLine2, v))
}

// AddressLine2ContainsFold applies the ContainsFold predicate on the "address_line2" field.
func AddressLine2ContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldAddressLine2, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCity, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldState, v))
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldState, v))
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldState, v))
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldState, v))
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldState, v))
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldState, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCountry, v))
}

// PostalCodeEQ applies the EQ predicate on the "postal_code" field.
func PostalCodeEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldPostalCode, v))
}

// PostalCodeNEQ applies the NEQ predicate on the "postal_code" field.
func PostalCodeNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldPostalCode, v))
}

// PostalCodeIn applies the In predicate on the "postal_code" field.
func PostalCodeIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldPostalCode, vs...))
}

// PostalCodeNotIn applies the NotIn predicate on the "postal_code" field.
func PostalCodeNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldPostalCode, vs...))
}

// PostalCodeGT applies the GT predicate on the "postal_code" field.
func PostalCodeGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldPostalCode, v))
}

// PostalCodeGTE applies the GTE predicate on the "postal_code" field.
func PostalCodeGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldPostalCode, v))
}

// PostalCodeLT applies the LT predicate on the "postal_code" field.
func PostalCodeLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldPostalCode, v))
}

// PostalCodeLTE applies the LTE predicate on the "postal_code" field.
func PostalCodeLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldPostalCode, v))
}

// PostalCodeContains applies the Contains predicate on the "postal_code" field.
func PostalCodeContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldPostalCode, v))
}

// PostalCodeHasPrefix applies the HasPrefix predicate on the "postal_code" field.
func PostalCodeHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldPostalCode, v))
}

// PostalCodeHasSuffix applies the HasSuffix predicate on the "postal_code" field.
func PostalCodeHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldPostalCode, v))
}

// PostalCodeEqualFold applies the EqualFold predicate on the "postal_code" field.
func PostalCodeEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldPostalCode, v))
}

// PostalCodeContainsFold applies the ContainsFold predicate on the "postal_code" field.
func PostalCodeContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldPostalCode, v))
}

// HasUserAddress applies the HasEdge predicate on the "user_address" edge.
func HasUserAddress() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserAddressTable, UserAddressColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserAddressWith applies the HasEdge predicate on the "user_address" edge with a given conditions (other predicates).
func HasUserAddressWith(preds ...predicate.UserAddress) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := newUserAddressStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
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
func Not(p predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		p(s.Not())
	})
}

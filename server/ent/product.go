// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"healthyshopper/ent/product"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Description holds the value of the "Description" field.
	Description string `json:"Description,omitempty"`
	// ProductImage holds the value of the "ProductImage" field.
	ProductImage string `json:"ProductImage,omitempty"`
	// ProductCategoryID holds the value of the "ProductCategoryID" field.
	ProductCategoryID int `json:"ProductCategoryID,omitempty"`
	// IngredientsListID holds the value of the "IngredientsListID" field.
	IngredientsListID int `json:"IngredientsListID,omitempty"`
	// NutritionalInformationID holds the value of the "NutritionalInformationID" field.
	NutritionalInformationID int `json:"NutritionalInformationID,omitempty"`
	// PromotionID holds the value of the "PromotionID" field.
	PromotionID int `json:"PromotionID,omitempty"`
	// CreatedAt holds the value of the "CreatedAt" field.
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
	// UpdatedAt holds the value of the "UpdatedAt" field.
	UpdatedAt    time.Time `json:"UpdatedAt,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldID, product.FieldProductCategoryID, product.FieldIngredientsListID, product.FieldNutritionalInformationID, product.FieldPromotionID:
			values[i] = new(sql.NullInt64)
		case product.FieldName, product.FieldDescription, product.FieldProductImage:
			values[i] = new(sql.NullString)
		case product.FieldCreatedAt, product.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case product.FieldProductImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ProductImage", values[i])
			} else if value.Valid {
				pr.ProductImage = value.String
			}
		case product.FieldProductCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ProductCategoryID", values[i])
			} else if value.Valid {
				pr.ProductCategoryID = int(value.Int64)
			}
		case product.FieldIngredientsListID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field IngredientsListID", values[i])
			} else if value.Valid {
				pr.IngredientsListID = int(value.Int64)
			}
		case product.FieldNutritionalInformationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field NutritionalInformationID", values[i])
			} else if value.Valid {
				pr.NutritionalInformationID = int(value.Int64)
			}
		case product.FieldPromotionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field PromotionID", values[i])
			} else if value.Valid {
				pr.PromotionID = int(value.Int64)
			}
		case product.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreatedAt", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case product.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdatedAt", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Product.
// This includes values selected through modifiers, order, etc.
func (pr *Product) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return NewProductClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("Name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("Description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("ProductImage=")
	builder.WriteString(pr.ProductImage)
	builder.WriteString(", ")
	builder.WriteString("ProductCategoryID=")
	builder.WriteString(fmt.Sprintf("%v", pr.ProductCategoryID))
	builder.WriteString(", ")
	builder.WriteString("IngredientsListID=")
	builder.WriteString(fmt.Sprintf("%v", pr.IngredientsListID))
	builder.WriteString(", ")
	builder.WriteString("NutritionalInformationID=")
	builder.WriteString(fmt.Sprintf("%v", pr.NutritionalInformationID))
	builder.WriteString(", ")
	builder.WriteString("PromotionID=")
	builder.WriteString(fmt.Sprintf("%v", pr.PromotionID))
	builder.WriteString(", ")
	builder.WriteString("CreatedAt=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("UpdatedAt=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

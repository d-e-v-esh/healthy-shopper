// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"healthyshopper/ent/ingredientstable"
	"healthyshopper/ent/nutritionalinformation"
	"healthyshopper/ent/product"
	"healthyshopper/ent/productitem"
	"healthyshopper/ent/promotion"
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
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// ProductImage holds the value of the "product_image" field.
	ProductImage string `json:"product_image,omitempty"`
	// IngredientsTableID holds the value of the "ingredients_table_id" field.
	IngredientsTableID int `json:"ingredients_table_id,omitempty"`
	// ProductCategoryID holds the value of the "product_category_id" field.
	ProductCategoryID int `json:"product_category_id,omitempty"`
	// NutritionalInformationID holds the value of the "nutritional_information_id" field.
	NutritionalInformationID int `json:"nutritional_information_id,omitempty"`
	// PromotionID holds the value of the "promotion_id" field.
	PromotionID int `json:"promotion_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges        ProductEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// ProductItem holds the value of the product_item edge.
	ProductItem *ProductItem `json:"product_item,omitempty"`
	// Promotion holds the value of the promotion edge.
	Promotion *Promotion `json:"promotion,omitempty"`
	// IngredientsTable holds the value of the ingredients_table edge.
	IngredientsTable *IngredientsTable `json:"ingredients_table,omitempty"`
	// NutritionalInformation holds the value of the nutritional_information edge.
	NutritionalInformation *NutritionalInformation `json:"nutritional_information,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int
}

// ProductItemOrErr returns the ProductItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) ProductItemOrErr() (*ProductItem, error) {
	if e.loadedTypes[0] {
		if e.ProductItem == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: productitem.Label}
		}
		return e.ProductItem, nil
	}
	return nil, &NotLoadedError{edge: "product_item"}
}

// PromotionOrErr returns the Promotion value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) PromotionOrErr() (*Promotion, error) {
	if e.loadedTypes[1] {
		if e.Promotion == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: promotion.Label}
		}
		return e.Promotion, nil
	}
	return nil, &NotLoadedError{edge: "promotion"}
}

// IngredientsTableOrErr returns the IngredientsTable value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) IngredientsTableOrErr() (*IngredientsTable, error) {
	if e.loadedTypes[2] {
		if e.IngredientsTable == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: ingredientstable.Label}
		}
		return e.IngredientsTable, nil
	}
	return nil, &NotLoadedError{edge: "ingredients_table"}
}

// NutritionalInformationOrErr returns the NutritionalInformation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) NutritionalInformationOrErr() (*NutritionalInformation, error) {
	if e.loadedTypes[3] {
		if e.NutritionalInformation == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: nutritionalinformation.Label}
		}
		return e.NutritionalInformation, nil
	}
	return nil, &NotLoadedError{edge: "nutritional_information"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldID, product.FieldIngredientsTableID, product.FieldProductCategoryID, product.FieldNutritionalInformationID, product.FieldPromotionID:
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
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case product.FieldProductImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field product_image", values[i])
			} else if value.Valid {
				pr.ProductImage = value.String
			}
		case product.FieldIngredientsTableID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ingredients_table_id", values[i])
			} else if value.Valid {
				pr.IngredientsTableID = int(value.Int64)
			}
		case product.FieldProductCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_category_id", values[i])
			} else if value.Valid {
				pr.ProductCategoryID = int(value.Int64)
			}
		case product.FieldNutritionalInformationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nutritional_information_id", values[i])
			} else if value.Valid {
				pr.NutritionalInformationID = int(value.Int64)
			}
		case product.FieldPromotionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field promotion_id", values[i])
			} else if value.Valid {
				pr.PromotionID = int(value.Int64)
			}
		case product.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case product.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
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

// QueryProductItem queries the "product_item" edge of the Product entity.
func (pr *Product) QueryProductItem() *ProductItemQuery {
	return NewProductClient(pr.config).QueryProductItem(pr)
}

// QueryPromotion queries the "promotion" edge of the Product entity.
func (pr *Product) QueryPromotion() *PromotionQuery {
	return NewProductClient(pr.config).QueryPromotion(pr)
}

// QueryIngredientsTable queries the "ingredients_table" edge of the Product entity.
func (pr *Product) QueryIngredientsTable() *IngredientsTableQuery {
	return NewProductClient(pr.config).QueryIngredientsTable(pr)
}

// QueryNutritionalInformation queries the "nutritional_information" edge of the Product entity.
func (pr *Product) QueryNutritionalInformation() *NutritionalInformationQuery {
	return NewProductClient(pr.config).QueryNutritionalInformation(pr)
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
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("product_image=")
	builder.WriteString(pr.ProductImage)
	builder.WriteString(", ")
	builder.WriteString("ingredients_table_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.IngredientsTableID))
	builder.WriteString(", ")
	builder.WriteString("product_category_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.ProductCategoryID))
	builder.WriteString(", ")
	builder.WriteString("nutritional_information_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.NutritionalInformationID))
	builder.WriteString(", ")
	builder.WriteString("promotion_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.PromotionID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

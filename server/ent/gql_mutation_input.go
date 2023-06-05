// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateProductInput represents a mutation input for creating products.
type CreateProductInput struct {
	Name                     string
	Description              string
	ProductImage             string
	IngredientsListID        *int
	ProductCategoryID        *int
	CreatedAt                *time.Time
	UpdatedAt                *time.Time
	ProductItemID            *int
	PromotionID              *int
	NutritionalInformationID *int
}

// Mutate applies the CreateProductInput on the ProductMutation builder.
func (i *CreateProductInput) Mutate(m *ProductMutation) {
	m.SetName(i.Name)
	m.SetDescription(i.Description)
	m.SetProductImage(i.ProductImage)
	if v := i.IngredientsListID; v != nil {
		m.SetIngredientsListID(*v)
	}
	if v := i.ProductCategoryID; v != nil {
		m.SetProductCategoryID(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.ProductItemID; v != nil {
		m.SetProductItemID(*v)
	}
	if v := i.PromotionID; v != nil {
		m.SetPromotionID(*v)
	}
	if v := i.NutritionalInformationID; v != nil {
		m.SetNutritionalInformationID(*v)
	}
}

// SetInput applies the change-set in the CreateProductInput on the ProductCreate builder.
func (c *ProductCreate) SetInput(i CreateProductInput) *ProductCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateProductItemInput represents a mutation input for creating productitems.
type CreateProductItemInput struct {
	StockKeepingUnit string
	QuantityInStock  int
	ProductImage     string
	Price            float32
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	ProductID        int
}

// Mutate applies the CreateProductItemInput on the ProductItemMutation builder.
func (i *CreateProductItemInput) Mutate(m *ProductItemMutation) {
	m.SetStockKeepingUnit(i.StockKeepingUnit)
	m.SetQuantityInStock(i.QuantityInStock)
	m.SetProductImage(i.ProductImage)
	m.SetPrice(i.Price)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetProductID(i.ProductID)
}

// SetInput applies the change-set in the CreateProductItemInput on the ProductItemCreate builder.
func (c *ProductItemCreate) SetInput(i CreateProductItemInput) *ProductItemCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Username        string
	EmailAddress    string
	Password        string
	FirstName       string
	LastName        string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	UserAddresIDs   []int
	UserReviewIDs   []int
	ShoppingCartIDs []int
	ShopOrderIDs    []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetUsername(i.Username)
	m.SetEmailAddress(i.EmailAddress)
	m.SetPassword(i.Password)
	m.SetFirstName(i.FirstName)
	m.SetLastName(i.LastName)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.UserAddresIDs; len(v) > 0 {
		m.AddUserAddresIDs(v...)
	}
	if v := i.UserReviewIDs; len(v) > 0 {
		m.AddUserReviewIDs(v...)
	}
	if v := i.ShoppingCartIDs; len(v) > 0 {
		m.AddShoppingCartIDs(v...)
	}
	if v := i.ShopOrderIDs; len(v) > 0 {
		m.AddShopOrderIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserAddressInput represents a mutation input for creating useraddresses.
type CreateUserAddressInput struct {
	IsDefault *bool
	UserID    int
	AddressID int
}

// Mutate applies the CreateUserAddressInput on the UserAddressMutation builder.
func (i *CreateUserAddressInput) Mutate(m *UserAddressMutation) {
	if v := i.IsDefault; v != nil {
		m.SetIsDefault(*v)
	}
	m.SetUserID(i.UserID)
	m.SetAddressID(i.AddressID)
}

// SetInput applies the change-set in the CreateUserAddressInput on the UserAddressCreate builder.
func (c *UserAddressCreate) SetInput(i CreateUserAddressInput) *UserAddressCreate {
	i.Mutate(c.Mutation())
	return c
}

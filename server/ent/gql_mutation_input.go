// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateProductInput represents a mutation input for creating products.
type CreateProductInput struct {
	ProductID                int
	Name                     string
	Description              string
	ProductImage             string
	ProductCategoryID        int
	IngredientsListID        int
	NutritionalInformationID int
	PromotionID              int
	CreatedAt                *time.Time
	UpdatedAt                *time.Time
}

// Mutate applies the CreateProductInput on the ProductMutation builder.
func (i *CreateProductInput) Mutate(m *ProductMutation) {
	m.SetProductID(i.ProductID)
	m.SetName(i.Name)
	m.SetDescription(i.Description)
	m.SetProductImage(i.ProductImage)
	m.SetProductCategoryID(i.ProductCategoryID)
	m.SetIngredientsListID(i.IngredientsListID)
	m.SetNutritionalInformationID(i.NutritionalInformationID)
	m.SetPromotionID(i.PromotionID)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
}

// SetInput applies the change-set in the CreateProductInput on the ProductCreate builder.
func (c *ProductCreate) SetInput(i CreateProductInput) *ProductCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	UserID        int
	Username      string
	EmailAddress  string
	Password      string
	FirstName     string
	LastName      string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	UserAddresIDs []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetUserID(i.UserID)
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

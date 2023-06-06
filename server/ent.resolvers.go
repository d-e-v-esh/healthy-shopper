package healthyshopper

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"healthyshopper/ent"
)

// Price is the resolver for the price field.
func (r *productItemResolver) Price(ctx context.Context, obj *ent.ProductItem) (float64, error) {
	panic(fmt.Errorf("not implemented: Price - price"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Addresses is the resolver for the addresses field.
func (r *queryResolver) Addresses(ctx context.Context) ([]*ent.Address, error) {
	panic(fmt.Errorf("not implemented: Addresses - addresses"))
}

// IngredientsTables is the resolver for the ingredientsTables field.
func (r *queryResolver) IngredientsTables(ctx context.Context) ([]*ent.IngredientsTable, error) {
	panic(fmt.Errorf("not implemented: IngredientsTables - ingredientsTables"))
}

// NutritionalInformations is the resolver for the nutritionalInformations field.
func (r *queryResolver) NutritionalInformations(ctx context.Context) ([]*ent.NutritionalInformation, error) {
	panic(fmt.Errorf("not implemented: NutritionalInformations - nutritionalInformations"))
}

// NutritionalInformationTables is the resolver for the nutritionalInformationTables field.
func (r *queryResolver) NutritionalInformationTables(ctx context.Context) ([]*ent.NutritionalInformationTable, error) {
	panic(fmt.Errorf("not implemented: NutritionalInformationTables - nutritionalInformationTables"))
}

// OrderLines is the resolver for the orderLines field.
func (r *queryResolver) OrderLines(ctx context.Context) ([]*ent.OrderLine, error) {
	panic(fmt.Errorf("not implemented: OrderLines - orderLines"))
}

// OrderStatusSlice is the resolver for the orderStatusSlice field.
func (r *queryResolver) OrderStatusSlice(ctx context.Context) ([]*ent.OrderStatus, error) {
	panic(fmt.Errorf("not implemented: OrderStatusSlice - orderStatusSlice"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*ent.Product, error) {
	panic(fmt.Errorf("not implemented: Products - products"))
}

// ProductItems is the resolver for the productItems field.
func (r *queryResolver) ProductItems(ctx context.Context) ([]*ent.ProductItem, error) {
	panic(fmt.Errorf("not implemented: ProductItems - productItems"))
}

// Promotions is the resolver for the promotions field.
func (r *queryResolver) Promotions(ctx context.Context) ([]*ent.Promotion, error) {
	panic(fmt.Errorf("not implemented: Promotions - promotions"))
}

// ShippingAddresses is the resolver for the shippingAddresses field.
func (r *queryResolver) ShippingAddresses(ctx context.Context) ([]*ent.ShippingAddress, error) {
	panic(fmt.Errorf("not implemented: ShippingAddresses - shippingAddresses"))
}

// ShippingMethods is the resolver for the shippingMethods field.
func (r *queryResolver) ShippingMethods(ctx context.Context) ([]*ent.ShippingMethod, error) {
	panic(fmt.Errorf("not implemented: ShippingMethods - shippingMethods"))
}

// ShopOrders is the resolver for the shopOrders field.
func (r *queryResolver) ShopOrders(ctx context.Context) ([]*ent.ShopOrder, error) {
	panic(fmt.Errorf("not implemented: ShopOrders - shopOrders"))
}

// ShoppingCarts is the resolver for the shoppingCarts field.
func (r *queryResolver) ShoppingCarts(ctx context.Context) ([]*ent.ShoppingCart, error) {
	panic(fmt.Errorf("not implemented: ShoppingCarts - shoppingCarts"))
}

// ShoppingCartItems is the resolver for the shoppingCartItems field.
func (r *queryResolver) ShoppingCartItems(ctx context.Context) ([]*ent.ShoppingCartItem, error) {
	panic(fmt.Errorf("not implemented: ShoppingCartItems - shoppingCartItems"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// UserAddresses is the resolver for the userAddresses field.
func (r *queryResolver) UserAddresses(ctx context.Context) ([]*ent.UserAddress, error) {
	panic(fmt.Errorf("not implemented: UserAddresses - userAddresses"))
}

// UserReviews is the resolver for the userReviews field.
func (r *queryResolver) UserReviews(ctx context.Context) ([]*ent.UserReview, error) {
	panic(fmt.Errorf("not implemented: UserReviews - userReviews"))
}

// Price is the resolver for the price field.
func (r *createProductItemInputResolver) Price(ctx context.Context, obj *ent.CreateProductItemInput, data float64) error {
	panic(fmt.Errorf("not implemented: Price - price"))
}

// Price is the resolver for the price field.
func (r *updateProductItemInputResolver) Price(ctx context.Context, obj *ent.UpdateProductItemInput, data *float64) error {
	panic(fmt.Errorf("not implemented: Price - price"))
}

// ProductItem returns ProductItemResolver implementation.
func (r *Resolver) ProductItem() ProductItemResolver { return &productItemResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// CreateProductItemInput returns CreateProductItemInputResolver implementation.
func (r *Resolver) CreateProductItemInput() CreateProductItemInputResolver {
	return &createProductItemInputResolver{r}
}

// UpdateProductItemInput returns UpdateProductItemInputResolver implementation.
func (r *Resolver) UpdateProductItemInput() UpdateProductItemInputResolver {
	return &updateProductItemInputResolver{r}
}

type productItemResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type createProductItemInputResolver struct{ *Resolver }
type updateProductItemInputResolver struct{ *Resolver }

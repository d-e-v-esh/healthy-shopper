// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"healthyshopper/ent/address"
	"healthyshopper/ent/product"
	"healthyshopper/ent/productitem"
	"healthyshopper/ent/promotion"
	"healthyshopper/ent/shoppingcart"
	"healthyshopper/ent/shoppingcartitem"
	"healthyshopper/ent/user"
	"healthyshopper/ent/useraddress"
	"healthyshopper/ent/userreview"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AddressQuery) CollectFields(ctx context.Context, satisfies ...string) (*AddressQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AddressQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(address.Columns))
		selectedFields = []string{address.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "userAddress":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserAddressClient{config: a.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, useraddressImplementors)...); err != nil {
				return err
			}
			a.WithNamedUserAddress(alias, func(wq *UserAddressQuery) {
				*wq = *query
			})
		case "phoneNumber":
			if _, ok := fieldSeen[address.FieldPhoneNumber]; !ok {
				selectedFields = append(selectedFields, address.FieldPhoneNumber)
				fieldSeen[address.FieldPhoneNumber] = struct{}{}
			}
		case "addressLine1":
			if _, ok := fieldSeen[address.FieldAddressLine1]; !ok {
				selectedFields = append(selectedFields, address.FieldAddressLine1)
				fieldSeen[address.FieldAddressLine1] = struct{}{}
			}
		case "addressLine2":
			if _, ok := fieldSeen[address.FieldAddressLine2]; !ok {
				selectedFields = append(selectedFields, address.FieldAddressLine2)
				fieldSeen[address.FieldAddressLine2] = struct{}{}
			}
		case "city":
			if _, ok := fieldSeen[address.FieldCity]; !ok {
				selectedFields = append(selectedFields, address.FieldCity)
				fieldSeen[address.FieldCity] = struct{}{}
			}
		case "state":
			if _, ok := fieldSeen[address.FieldState]; !ok {
				selectedFields = append(selectedFields, address.FieldState)
				fieldSeen[address.FieldState] = struct{}{}
			}
		case "country":
			if _, ok := fieldSeen[address.FieldCountry]; !ok {
				selectedFields = append(selectedFields, address.FieldCountry)
				fieldSeen[address.FieldCountry] = struct{}{}
			}
		case "postalCode":
			if _, ok := fieldSeen[address.FieldPostalCode]; !ok {
				selectedFields = append(selectedFields, address.FieldPostalCode)
				fieldSeen[address.FieldPostalCode] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		a.Select(selectedFields...)
	}
	return nil
}

type addressPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AddressPaginateOption
}

func newAddressPaginateArgs(rv map[string]any) *addressPaginateArgs {
	args := &addressPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pr *ProductQuery) CollectFields(ctx context.Context, satisfies ...string) (*ProductQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pr, nil
	}
	if err := pr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pr, nil
}

func (pr *ProductQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(product.Columns))
		selectedFields = []string{product.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "productItem":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ProductItemClient{config: pr.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, productitemImplementors)...); err != nil {
				return err
			}
			pr.withProductItem = query
		case "promotion":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PromotionClient{config: pr.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, promotionImplementors)...); err != nil {
				return err
			}
			pr.withPromotion = query
			if _, ok := fieldSeen[product.FieldPromotionID]; !ok {
				selectedFields = append(selectedFields, product.FieldPromotionID)
				fieldSeen[product.FieldPromotionID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[product.FieldName]; !ok {
				selectedFields = append(selectedFields, product.FieldName)
				fieldSeen[product.FieldName] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[product.FieldDescription]; !ok {
				selectedFields = append(selectedFields, product.FieldDescription)
				fieldSeen[product.FieldDescription] = struct{}{}
			}
		case "productImage":
			if _, ok := fieldSeen[product.FieldProductImage]; !ok {
				selectedFields = append(selectedFields, product.FieldProductImage)
				fieldSeen[product.FieldProductImage] = struct{}{}
			}
		case "ingredientsListID":
			if _, ok := fieldSeen[product.FieldIngredientsListID]; !ok {
				selectedFields = append(selectedFields, product.FieldIngredientsListID)
				fieldSeen[product.FieldIngredientsListID] = struct{}{}
			}
		case "productCategoryID":
			if _, ok := fieldSeen[product.FieldProductCategoryID]; !ok {
				selectedFields = append(selectedFields, product.FieldProductCategoryID)
				fieldSeen[product.FieldProductCategoryID] = struct{}{}
			}
		case "nutritionalInformationID":
			if _, ok := fieldSeen[product.FieldNutritionalInformationID]; !ok {
				selectedFields = append(selectedFields, product.FieldNutritionalInformationID)
				fieldSeen[product.FieldNutritionalInformationID] = struct{}{}
			}
		case "promotionID":
			if _, ok := fieldSeen[product.FieldPromotionID]; !ok {
				selectedFields = append(selectedFields, product.FieldPromotionID)
				fieldSeen[product.FieldPromotionID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[product.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, product.FieldCreatedAt)
				fieldSeen[product.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[product.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, product.FieldUpdatedAt)
				fieldSeen[product.FieldUpdatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pr.Select(selectedFields...)
	}
	return nil
}

type productPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ProductPaginateOption
}

func newProductPaginateArgs(rv map[string]any) *productPaginateArgs {
	args := &productPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pi *ProductItemQuery) CollectFields(ctx context.Context, satisfies ...string) (*ProductItemQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pi, nil
	}
	if err := pi.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pi, nil
}

func (pi *ProductItemQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(productitem.Columns))
		selectedFields = []string{productitem.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "product":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ProductClient{config: pi.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, productImplementors)...); err != nil {
				return err
			}
			pi.withProduct = query
			if _, ok := fieldSeen[productitem.FieldProductID]; !ok {
				selectedFields = append(selectedFields, productitem.FieldProductID)
				fieldSeen[productitem.FieldProductID] = struct{}{}
			}
		case "productID":
			if _, ok := fieldSeen[productitem.FieldProductID]; !ok {
				selectedFields = append(selectedFields, productitem.FieldProductID)
				fieldSeen[productitem.FieldProductID] = struct{}{}
			}
		case "stockKeepingUnit":
			if _, ok := fieldSeen[productitem.FieldStockKeepingUnit]; !ok {
				selectedFields = append(selectedFields, productitem.FieldStockKeepingUnit)
				fieldSeen[productitem.FieldStockKeepingUnit] = struct{}{}
			}
		case "quantityInStock":
			if _, ok := fieldSeen[productitem.FieldQuantityInStock]; !ok {
				selectedFields = append(selectedFields, productitem.FieldQuantityInStock)
				fieldSeen[productitem.FieldQuantityInStock] = struct{}{}
			}
		case "productImage":
			if _, ok := fieldSeen[productitem.FieldProductImage]; !ok {
				selectedFields = append(selectedFields, productitem.FieldProductImage)
				fieldSeen[productitem.FieldProductImage] = struct{}{}
			}
		case "price":
			if _, ok := fieldSeen[productitem.FieldPrice]; !ok {
				selectedFields = append(selectedFields, productitem.FieldPrice)
				fieldSeen[productitem.FieldPrice] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[productitem.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, productitem.FieldCreatedAt)
				fieldSeen[productitem.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[productitem.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, productitem.FieldUpdatedAt)
				fieldSeen[productitem.FieldUpdatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pi.Select(selectedFields...)
	}
	return nil
}

type productitemPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ProductItemPaginateOption
}

func newProductItemPaginateArgs(rv map[string]any) *productitemPaginateArgs {
	args := &productitemPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pr *PromotionQuery) CollectFields(ctx context.Context, satisfies ...string) (*PromotionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pr, nil
	}
	if err := pr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pr, nil
}

func (pr *PromotionQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(promotion.Columns))
		selectedFields = []string{promotion.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "product":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ProductClient{config: pr.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, productImplementors)...); err != nil {
				return err
			}
			pr.WithNamedProduct(alias, func(wq *ProductQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[promotion.FieldName]; !ok {
				selectedFields = append(selectedFields, promotion.FieldName)
				fieldSeen[promotion.FieldName] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[promotion.FieldDescription]; !ok {
				selectedFields = append(selectedFields, promotion.FieldDescription)
				fieldSeen[promotion.FieldDescription] = struct{}{}
			}
		case "discountPercentage":
			if _, ok := fieldSeen[promotion.FieldDiscountPercentage]; !ok {
				selectedFields = append(selectedFields, promotion.FieldDiscountPercentage)
				fieldSeen[promotion.FieldDiscountPercentage] = struct{}{}
			}
		case "startDate":
			if _, ok := fieldSeen[promotion.FieldStartDate]; !ok {
				selectedFields = append(selectedFields, promotion.FieldStartDate)
				fieldSeen[promotion.FieldStartDate] = struct{}{}
			}
		case "endDate":
			if _, ok := fieldSeen[promotion.FieldEndDate]; !ok {
				selectedFields = append(selectedFields, promotion.FieldEndDate)
				fieldSeen[promotion.FieldEndDate] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pr.Select(selectedFields...)
	}
	return nil
}

type promotionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PromotionPaginateOption
}

func newPromotionPaginateArgs(rv map[string]any) *promotionPaginateArgs {
	args := &promotionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sc *ShoppingCartQuery) CollectFields(ctx context.Context, satisfies ...string) (*ShoppingCartQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sc, nil
	}
	if err := sc.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sc, nil
}

func (sc *ShoppingCartQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(shoppingcart.Columns))
		selectedFields = []string{shoppingcart.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: sc.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			sc.withUser = query
			if _, ok := fieldSeen[shoppingcart.FieldUserID]; !ok {
				selectedFields = append(selectedFields, shoppingcart.FieldUserID)
				fieldSeen[shoppingcart.FieldUserID] = struct{}{}
			}
		case "shoppingCartItem":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ShoppingCartItemClient{config: sc.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, shoppingcartitemImplementors)...); err != nil {
				return err
			}
			sc.withShoppingCartItem = query
		case "userID":
			if _, ok := fieldSeen[shoppingcart.FieldUserID]; !ok {
				selectedFields = append(selectedFields, shoppingcart.FieldUserID)
				fieldSeen[shoppingcart.FieldUserID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		sc.Select(selectedFields...)
	}
	return nil
}

type shoppingcartPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ShoppingCartPaginateOption
}

func newShoppingCartPaginateArgs(rv map[string]any) *shoppingcartPaginateArgs {
	args := &shoppingcartPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (sci *ShoppingCartItemQuery) CollectFields(ctx context.Context, satisfies ...string) (*ShoppingCartItemQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return sci, nil
	}
	if err := sci.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return sci, nil
}

func (sci *ShoppingCartItemQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(shoppingcartitem.Columns))
		selectedFields = []string{shoppingcartitem.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "shoppingCart":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ShoppingCartClient{config: sci.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, shoppingcartImplementors)...); err != nil {
				return err
			}
			sci.withShoppingCart = query
			if _, ok := fieldSeen[shoppingcartitem.FieldShoppingCartID]; !ok {
				selectedFields = append(selectedFields, shoppingcartitem.FieldShoppingCartID)
				fieldSeen[shoppingcartitem.FieldShoppingCartID] = struct{}{}
			}
		case "shoppingCartID":
			if _, ok := fieldSeen[shoppingcartitem.FieldShoppingCartID]; !ok {
				selectedFields = append(selectedFields, shoppingcartitem.FieldShoppingCartID)
				fieldSeen[shoppingcartitem.FieldShoppingCartID] = struct{}{}
			}
		case "productItemID":
			if _, ok := fieldSeen[shoppingcartitem.FieldProductItemID]; !ok {
				selectedFields = append(selectedFields, shoppingcartitem.FieldProductItemID)
				fieldSeen[shoppingcartitem.FieldProductItemID] = struct{}{}
			}
		case "quantity":
			if _, ok := fieldSeen[shoppingcartitem.FieldQuantity]; !ok {
				selectedFields = append(selectedFields, shoppingcartitem.FieldQuantity)
				fieldSeen[shoppingcartitem.FieldQuantity] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		sci.Select(selectedFields...)
	}
	return nil
}

type shoppingcartitemPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ShoppingCartItemPaginateOption
}

func newShoppingCartItemPaginateArgs(rv map[string]any) *shoppingcartitemPaginateArgs {
	args := &shoppingcartitemPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "userAddress":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserAddressClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, useraddressImplementors)...); err != nil {
				return err
			}
			u.WithNamedUserAddress(alias, func(wq *UserAddressQuery) {
				*wq = *query
			})
		case "userReview":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserReviewClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, userreviewImplementors)...); err != nil {
				return err
			}
			u.WithNamedUserReview(alias, func(wq *UserReviewQuery) {
				*wq = *query
			})
		case "shoppingCart":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ShoppingCartClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, shoppingcartImplementors)...); err != nil {
				return err
			}
			u.WithNamedShoppingCart(alias, func(wq *ShoppingCartQuery) {
				*wq = *query
			})
		case "username":
			if _, ok := fieldSeen[user.FieldUsername]; !ok {
				selectedFields = append(selectedFields, user.FieldUsername)
				fieldSeen[user.FieldUsername] = struct{}{}
			}
		case "emailAddress":
			if _, ok := fieldSeen[user.FieldEmailAddress]; !ok {
				selectedFields = append(selectedFields, user.FieldEmailAddress)
				fieldSeen[user.FieldEmailAddress] = struct{}{}
			}
		case "password":
			if _, ok := fieldSeen[user.FieldPassword]; !ok {
				selectedFields = append(selectedFields, user.FieldPassword)
				fieldSeen[user.FieldPassword] = struct{}{}
			}
		case "firstName":
			if _, ok := fieldSeen[user.FieldFirstName]; !ok {
				selectedFields = append(selectedFields, user.FieldFirstName)
				fieldSeen[user.FieldFirstName] = struct{}{}
			}
		case "lastName":
			if _, ok := fieldSeen[user.FieldLastName]; !ok {
				selectedFields = append(selectedFields, user.FieldLastName)
				fieldSeen[user.FieldLastName] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[user.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldCreatedAt)
				fieldSeen[user.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[user.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldUpdatedAt)
				fieldSeen[user.FieldUpdatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ua *UserAddressQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserAddressQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ua, nil
	}
	if err := ua.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ua, nil
}

func (ua *UserAddressQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(useraddress.Columns))
		selectedFields = []string{useraddress.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: ua.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			ua.withUser = query
			if _, ok := fieldSeen[useraddress.FieldUserID]; !ok {
				selectedFields = append(selectedFields, useraddress.FieldUserID)
				fieldSeen[useraddress.FieldUserID] = struct{}{}
			}
		case "address":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AddressClient{config: ua.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, addressImplementors)...); err != nil {
				return err
			}
			ua.withAddress = query
			if _, ok := fieldSeen[useraddress.FieldAddressID]; !ok {
				selectedFields = append(selectedFields, useraddress.FieldAddressID)
				fieldSeen[useraddress.FieldAddressID] = struct{}{}
			}
		case "userID":
			if _, ok := fieldSeen[useraddress.FieldUserID]; !ok {
				selectedFields = append(selectedFields, useraddress.FieldUserID)
				fieldSeen[useraddress.FieldUserID] = struct{}{}
			}
		case "addressID":
			if _, ok := fieldSeen[useraddress.FieldAddressID]; !ok {
				selectedFields = append(selectedFields, useraddress.FieldAddressID)
				fieldSeen[useraddress.FieldAddressID] = struct{}{}
			}
		case "isDefault":
			if _, ok := fieldSeen[useraddress.FieldIsDefault]; !ok {
				selectedFields = append(selectedFields, useraddress.FieldIsDefault)
				fieldSeen[useraddress.FieldIsDefault] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ua.Select(selectedFields...)
	}
	return nil
}

type useraddressPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserAddressPaginateOption
}

func newUserAddressPaginateArgs(rv map[string]any) *useraddressPaginateArgs {
	args := &useraddressPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ur *UserReviewQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserReviewQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ur, nil
	}
	if err := ur.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ur, nil
}

func (ur *UserReviewQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(userreview.Columns))
		selectedFields = []string{userreview.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: ur.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			ur.withUser = query
			if _, ok := fieldSeen[userreview.FieldUserID]; !ok {
				selectedFields = append(selectedFields, userreview.FieldUserID)
				fieldSeen[userreview.FieldUserID] = struct{}{}
			}
		case "userID":
			if _, ok := fieldSeen[userreview.FieldUserID]; !ok {
				selectedFields = append(selectedFields, userreview.FieldUserID)
				fieldSeen[userreview.FieldUserID] = struct{}{}
			}
		case "orderedProductID":
			if _, ok := fieldSeen[userreview.FieldOrderedProductID]; !ok {
				selectedFields = append(selectedFields, userreview.FieldOrderedProductID)
				fieldSeen[userreview.FieldOrderedProductID] = struct{}{}
			}
		case "rating":
			if _, ok := fieldSeen[userreview.FieldRating]; !ok {
				selectedFields = append(selectedFields, userreview.FieldRating)
				fieldSeen[userreview.FieldRating] = struct{}{}
			}
		case "review":
			if _, ok := fieldSeen[userreview.FieldReview]; !ok {
				selectedFields = append(selectedFields, userreview.FieldReview)
				fieldSeen[userreview.FieldReview] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[userreview.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, userreview.FieldCreatedAt)
				fieldSeen[userreview.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[userreview.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, userreview.FieldUpdatedAt)
				fieldSeen[userreview.FieldUpdatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ur.Select(selectedFields...)
	}
	return nil
}

type userreviewPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserReviewPaginateOption
}

func newUserReviewPaginateArgs(rv map[string]any) *userreviewPaginateArgs {
	args := &userreviewPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}

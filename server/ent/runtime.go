// Code generated by ent, DO NOT EDIT.

package ent

import (
	"healthyshopper/ent/address"
	"healthyshopper/ent/ingredientstable"
	"healthyshopper/ent/nutritionalinformation"
	"healthyshopper/ent/nutritionalinformationtable"
	"healthyshopper/ent/orderline"
	"healthyshopper/ent/orderstatus"
	"healthyshopper/ent/product"
	"healthyshopper/ent/productitem"
	"healthyshopper/ent/promotion"
	"healthyshopper/ent/schema"
	"healthyshopper/ent/shippingaddress"
	"healthyshopper/ent/shippingmethod"
	"healthyshopper/ent/shoporder"
	"healthyshopper/ent/shoppingcartitem"
	"healthyshopper/ent/user"
	"healthyshopper/ent/useraddress"
	"healthyshopper/ent/userreview"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	addressFields := schema.Address{}.Fields()
	_ = addressFields
	// addressDescPhoneNumber is the schema descriptor for phone_number field.
	addressDescPhoneNumber := addressFields[0].Descriptor()
	// address.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	address.PhoneNumberValidator = func() func(string) error {
		validators := addressDescPhoneNumber.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(phone_number string) error {
			for _, fn := range fns {
				if err := fn(phone_number); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// addressDescAddressLine1 is the schema descriptor for address_line1 field.
	addressDescAddressLine1 := addressFields[1].Descriptor()
	// address.AddressLine1Validator is a validator for the "address_line1" field. It is called by the builders before save.
	address.AddressLine1Validator = func() func(string) error {
		validators := addressDescAddressLine1.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(address_line1 string) error {
			for _, fn := range fns {
				if err := fn(address_line1); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// addressDescAddressLine2 is the schema descriptor for address_line2 field.
	addressDescAddressLine2 := addressFields[2].Descriptor()
	// address.AddressLine2Validator is a validator for the "address_line2" field. It is called by the builders before save.
	address.AddressLine2Validator = addressDescAddressLine2.Validators[0].(func(string) error)
	// addressDescCity is the schema descriptor for city field.
	addressDescCity := addressFields[3].Descriptor()
	// address.CityValidator is a validator for the "city" field. It is called by the builders before save.
	address.CityValidator = func() func(string) error {
		validators := addressDescCity.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(city string) error {
			for _, fn := range fns {
				if err := fn(city); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// addressDescState is the schema descriptor for state field.
	addressDescState := addressFields[4].Descriptor()
	// address.StateValidator is a validator for the "state" field. It is called by the builders before save.
	address.StateValidator = func() func(string) error {
		validators := addressDescState.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(state string) error {
			for _, fn := range fns {
				if err := fn(state); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// addressDescCountry is the schema descriptor for country field.
	addressDescCountry := addressFields[5].Descriptor()
	// address.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	address.CountryValidator = func() func(string) error {
		validators := addressDescCountry.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(country string) error {
			for _, fn := range fns {
				if err := fn(country); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// addressDescPostalCode is the schema descriptor for postal_code field.
	addressDescPostalCode := addressFields[6].Descriptor()
	// address.PostalCodeValidator is a validator for the "postal_code" field. It is called by the builders before save.
	address.PostalCodeValidator = func() func(string) error {
		validators := addressDescPostalCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(postal_code string) error {
			for _, fn := range fns {
				if err := fn(postal_code); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	ingredientstableFields := schema.IngredientsTable{}.Fields()
	_ = ingredientstableFields
	// ingredientstableDescName is the schema descriptor for name field.
	ingredientstableDescName := ingredientstableFields[0].Descriptor()
	// ingredientstable.NameValidator is a validator for the "name" field. It is called by the builders before save.
	ingredientstable.NameValidator = func() func(string) error {
		validators := ingredientstableDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// ingredientstableDescDescription is the schema descriptor for description field.
	ingredientstableDescDescription := ingredientstableFields[1].Descriptor()
	// ingredientstable.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	ingredientstable.DescriptionValidator = func() func(string) error {
		validators := ingredientstableDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	nutritionalinformationFields := schema.NutritionalInformation{}.Fields()
	_ = nutritionalinformationFields
	// nutritionalinformationDescNValue is the schema descriptor for n_value field.
	nutritionalinformationDescNValue := nutritionalinformationFields[1].Descriptor()
	// nutritionalinformation.NValueValidator is a validator for the "n_value" field. It is called by the builders before save.
	nutritionalinformation.NValueValidator = nutritionalinformationDescNValue.Validators[0].(func(float64) error)
	// nutritionalinformationDescNMeasurementUnit is the schema descriptor for n_measurement_unit field.
	nutritionalinformationDescNMeasurementUnit := nutritionalinformationFields[2].Descriptor()
	// nutritionalinformation.NMeasurementUnitValidator is a validator for the "n_measurement_unit" field. It is called by the builders before save.
	nutritionalinformation.NMeasurementUnitValidator = func() func(string) error {
		validators := nutritionalinformationDescNMeasurementUnit.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(n_measurement_unit string) error {
			for _, fn := range fns {
				if err := fn(n_measurement_unit); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	nutritionalinformationtableFields := schema.NutritionalInformationTable{}.Fields()
	_ = nutritionalinformationtableFields
	// nutritionalinformationtableDescParameter is the schema descriptor for parameter field.
	nutritionalinformationtableDescParameter := nutritionalinformationtableFields[0].Descriptor()
	// nutritionalinformationtable.ParameterValidator is a validator for the "parameter" field. It is called by the builders before save.
	nutritionalinformationtable.ParameterValidator = func() func(string) error {
		validators := nutritionalinformationtableDescParameter.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(parameter string) error {
			for _, fn := range fns {
				if err := fn(parameter); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// nutritionalinformationtableDescValue is the schema descriptor for value field.
	nutritionalinformationtableDescValue := nutritionalinformationtableFields[1].Descriptor()
	// nutritionalinformationtable.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	nutritionalinformationtable.ValueValidator = nutritionalinformationtableDescValue.Validators[0].(func(float64) error)
	// nutritionalinformationtableDescMeasurementUnit is the schema descriptor for measurement_unit field.
	nutritionalinformationtableDescMeasurementUnit := nutritionalinformationtableFields[2].Descriptor()
	// nutritionalinformationtable.MeasurementUnitValidator is a validator for the "measurement_unit" field. It is called by the builders before save.
	nutritionalinformationtable.MeasurementUnitValidator = func() func(string) error {
		validators := nutritionalinformationtableDescMeasurementUnit.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(measurement_unit string) error {
			for _, fn := range fns {
				if err := fn(measurement_unit); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	orderlineFields := schema.OrderLine{}.Fields()
	_ = orderlineFields
	// orderlineDescQuantity is the schema descriptor for quantity field.
	orderlineDescQuantity := orderlineFields[2].Descriptor()
	// orderline.QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	orderline.QuantityValidator = orderlineDescQuantity.Validators[0].(func(int) error)
	// orderlineDescPrice is the schema descriptor for price field.
	orderlineDescPrice := orderlineFields[3].Descriptor()
	// orderline.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	orderline.PriceValidator = orderlineDescPrice.Validators[0].(func(float64) error)
	orderstatusFields := schema.OrderStatus{}.Fields()
	_ = orderstatusFields
	// orderstatusDescStatus is the schema descriptor for status field.
	orderstatusDescStatus := orderstatusFields[0].Descriptor()
	// orderstatus.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	orderstatus.StatusValidator = func() func(string) error {
		validators := orderstatusDescStatus.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(status string) error {
			for _, fn := range fns {
				if err := fn(status); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[0].Descriptor()
	// product.NameValidator is a validator for the "name" field. It is called by the builders before save.
	product.NameValidator = func() func(string) error {
		validators := productDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// productDescDescription is the schema descriptor for description field.
	productDescDescription := productFields[1].Descriptor()
	// product.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	product.DescriptionValidator = func() func(string) error {
		validators := productDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// productDescProductImage is the schema descriptor for product_image field.
	productDescProductImage := productFields[2].Descriptor()
	// product.ProductImageValidator is a validator for the "product_image" field. It is called by the builders before save.
	product.ProductImageValidator = func() func(string) error {
		validators := productDescProductImage.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(product_image string) error {
			for _, fn := range fns {
				if err := fn(product_image); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// productDescCreatedAt is the schema descriptor for created_at field.
	productDescCreatedAt := productFields[7].Descriptor()
	// product.DefaultCreatedAt holds the default value on creation for the created_at field.
	product.DefaultCreatedAt = productDescCreatedAt.Default.(func() time.Time)
	productitemFields := schema.ProductItem{}.Fields()
	_ = productitemFields
	// productitemDescStockKeepingUnit is the schema descriptor for stock_keeping_unit field.
	productitemDescStockKeepingUnit := productitemFields[1].Descriptor()
	// productitem.StockKeepingUnitValidator is a validator for the "stock_keeping_unit" field. It is called by the builders before save.
	productitem.StockKeepingUnitValidator = func() func(string) error {
		validators := productitemDescStockKeepingUnit.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(stock_keeping_unit string) error {
			for _, fn := range fns {
				if err := fn(stock_keeping_unit); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// productitemDescQuantityInStock is the schema descriptor for quantity_in_stock field.
	productitemDescQuantityInStock := productitemFields[2].Descriptor()
	// productitem.QuantityInStockValidator is a validator for the "quantity_in_stock" field. It is called by the builders before save.
	productitem.QuantityInStockValidator = productitemDescQuantityInStock.Validators[0].(func(int) error)
	// productitemDescProductImage is the schema descriptor for product_image field.
	productitemDescProductImage := productitemFields[3].Descriptor()
	// productitem.ProductImageValidator is a validator for the "product_image" field. It is called by the builders before save.
	productitem.ProductImageValidator = func() func(string) error {
		validators := productitemDescProductImage.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(product_image string) error {
			for _, fn := range fns {
				if err := fn(product_image); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// productitemDescPrice is the schema descriptor for price field.
	productitemDescPrice := productitemFields[4].Descriptor()
	// productitem.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	productitem.PriceValidator = productitemDescPrice.Validators[0].(func(float32) error)
	// productitemDescCreatedAt is the schema descriptor for created_at field.
	productitemDescCreatedAt := productitemFields[5].Descriptor()
	// productitem.DefaultCreatedAt holds the default value on creation for the created_at field.
	productitem.DefaultCreatedAt = productitemDescCreatedAt.Default.(func() time.Time)
	promotionFields := schema.Promotion{}.Fields()
	_ = promotionFields
	// promotionDescName is the schema descriptor for name field.
	promotionDescName := promotionFields[0].Descriptor()
	// promotion.NameValidator is a validator for the "name" field. It is called by the builders before save.
	promotion.NameValidator = func() func(string) error {
		validators := promotionDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// promotionDescDescription is the schema descriptor for description field.
	promotionDescDescription := promotionFields[1].Descriptor()
	// promotion.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	promotion.DescriptionValidator = func() func(string) error {
		validators := promotionDescDescription.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(description string) error {
			for _, fn := range fns {
				if err := fn(description); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// promotionDescDiscountPercentage is the schema descriptor for discount_percentage field.
	promotionDescDiscountPercentage := promotionFields[2].Descriptor()
	// promotion.DiscountPercentageValidator is a validator for the "discount_percentage" field. It is called by the builders before save.
	promotion.DiscountPercentageValidator = promotionDescDiscountPercentage.Validators[0].(func(int) error)
	shippingaddressFields := schema.ShippingAddress{}.Fields()
	_ = shippingaddressFields
	// shippingaddressDescPhoneNumber is the schema descriptor for phone_number field.
	shippingaddressDescPhoneNumber := shippingaddressFields[0].Descriptor()
	// shippingaddress.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	shippingaddress.PhoneNumberValidator = func() func(string) error {
		validators := shippingaddressDescPhoneNumber.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(phone_number string) error {
			for _, fn := range fns {
				if err := fn(phone_number); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// shippingaddressDescAddressLine1 is the schema descriptor for address_line1 field.
	shippingaddressDescAddressLine1 := shippingaddressFields[1].Descriptor()
	// shippingaddress.AddressLine1Validator is a validator for the "address_line1" field. It is called by the builders before save.
	shippingaddress.AddressLine1Validator = func() func(string) error {
		validators := shippingaddressDescAddressLine1.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(address_line1 string) error {
			for _, fn := range fns {
				if err := fn(address_line1); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// shippingaddressDescAddressLine2 is the schema descriptor for address_line2 field.
	shippingaddressDescAddressLine2 := shippingaddressFields[2].Descriptor()
	// shippingaddress.AddressLine2Validator is a validator for the "address_line2" field. It is called by the builders before save.
	shippingaddress.AddressLine2Validator = shippingaddressDescAddressLine2.Validators[0].(func(string) error)
	// shippingaddressDescCity is the schema descriptor for city field.
	shippingaddressDescCity := shippingaddressFields[3].Descriptor()
	// shippingaddress.CityValidator is a validator for the "city" field. It is called by the builders before save.
	shippingaddress.CityValidator = func() func(string) error {
		validators := shippingaddressDescCity.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(city string) error {
			for _, fn := range fns {
				if err := fn(city); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// shippingaddressDescState is the schema descriptor for state field.
	shippingaddressDescState := shippingaddressFields[4].Descriptor()
	// shippingaddress.StateValidator is a validator for the "state" field. It is called by the builders before save.
	shippingaddress.StateValidator = func() func(string) error {
		validators := shippingaddressDescState.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(state string) error {
			for _, fn := range fns {
				if err := fn(state); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// shippingaddressDescCountry is the schema descriptor for country field.
	shippingaddressDescCountry := shippingaddressFields[5].Descriptor()
	// shippingaddress.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	shippingaddress.CountryValidator = func() func(string) error {
		validators := shippingaddressDescCountry.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(country string) error {
			for _, fn := range fns {
				if err := fn(country); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// shippingaddressDescPostalCode is the schema descriptor for postal_code field.
	shippingaddressDescPostalCode := shippingaddressFields[6].Descriptor()
	// shippingaddress.PostalCodeValidator is a validator for the "postal_code" field. It is called by the builders before save.
	shippingaddress.PostalCodeValidator = func() func(string) error {
		validators := shippingaddressDescPostalCode.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(postal_code string) error {
			for _, fn := range fns {
				if err := fn(postal_code); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	shippingmethodFields := schema.ShippingMethod{}.Fields()
	_ = shippingmethodFields
	// shippingmethodDescShippingMethod is the schema descriptor for shipping_method field.
	shippingmethodDescShippingMethod := shippingmethodFields[0].Descriptor()
	// shippingmethod.ShippingMethodValidator is a validator for the "shipping_method" field. It is called by the builders before save.
	shippingmethod.ShippingMethodValidator = func() func(string) error {
		validators := shippingmethodDescShippingMethod.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(shipping_method string) error {
			for _, fn := range fns {
				if err := fn(shipping_method); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// shippingmethodDescShippingCost is the schema descriptor for shipping_cost field.
	shippingmethodDescShippingCost := shippingmethodFields[1].Descriptor()
	// shippingmethod.ShippingCostValidator is a validator for the "shipping_cost" field. It is called by the builders before save.
	shippingmethod.ShippingCostValidator = shippingmethodDescShippingCost.Validators[0].(func(float64) error)
	shoporderFields := schema.ShopOrder{}.Fields()
	_ = shoporderFields
	// shoporderDescOrderDateAndTime is the schema descriptor for order_date_and_time field.
	shoporderDescOrderDateAndTime := shoporderFields[0].Descriptor()
	// shoporder.DefaultOrderDateAndTime holds the default value on creation for the order_date_and_time field.
	shoporder.DefaultOrderDateAndTime = shoporderDescOrderDateAndTime.Default.(func() time.Time)
	// shoporderDescPaymentMethod is the schema descriptor for payment_method field.
	shoporderDescPaymentMethod := shoporderFields[1].Descriptor()
	// shoporder.PaymentMethodValidator is a validator for the "payment_method" field. It is called by the builders before save.
	shoporder.PaymentMethodValidator = func() func(string) error {
		validators := shoporderDescPaymentMethod.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(payment_method string) error {
			for _, fn := range fns {
				if err := fn(payment_method); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// shoporderDescTotalPrice is the schema descriptor for total_price field.
	shoporderDescTotalPrice := shoporderFields[2].Descriptor()
	// shoporder.TotalPriceValidator is a validator for the "total_price" field. It is called by the builders before save.
	shoporder.TotalPriceValidator = shoporderDescTotalPrice.Validators[0].(func(float64) error)
	shoppingcartitemFields := schema.ShoppingCartItem{}.Fields()
	_ = shoppingcartitemFields
	// shoppingcartitemDescQuantity is the schema descriptor for quantity field.
	shoppingcartitemDescQuantity := shoppingcartitemFields[2].Descriptor()
	// shoppingcartitem.QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	shoppingcartitem.QuantityValidator = shoppingcartitemDescQuantity.Validators[0].(func(int) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmailAddress is the schema descriptor for email_address field.
	userDescEmailAddress := userFields[1].Descriptor()
	// user.EmailAddressValidator is a validator for the "email_address" field. It is called by the builders before save.
	user.EmailAddressValidator = func() func(string) error {
		validators := userDescEmailAddress.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email_address string) error {
			for _, fn := range fns {
				if err := fn(email_address); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[3].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = func() func(string) error {
		validators := userDescFirstName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(first_name string) error {
			for _, fn := range fns {
				if err := fn(first_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[4].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = func() func(string) error {
		validators := userDescLastName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(last_name string) error {
			for _, fn := range fns {
				if err := fn(last_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	useraddressFields := schema.UserAddress{}.Fields()
	_ = useraddressFields
	// useraddressDescIsDefault is the schema descriptor for is_default field.
	useraddressDescIsDefault := useraddressFields[2].Descriptor()
	// useraddress.DefaultIsDefault holds the default value on creation for the is_default field.
	useraddress.DefaultIsDefault = useraddressDescIsDefault.Default.(bool)
	userreviewFields := schema.UserReview{}.Fields()
	_ = userreviewFields
	// userreviewDescRating is the schema descriptor for rating field.
	userreviewDescRating := userreviewFields[2].Descriptor()
	// userreview.RatingValidator is a validator for the "rating" field. It is called by the builders before save.
	userreview.RatingValidator = userreviewDescRating.Validators[0].(func(int) error)
	// userreviewDescReview is the schema descriptor for review field.
	userreviewDescReview := userreviewFields[3].Descriptor()
	// userreview.ReviewValidator is a validator for the "review" field. It is called by the builders before save.
	userreview.ReviewValidator = func() func(string) error {
		validators := userreviewDescReview.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(review string) error {
			for _, fn := range fns {
				if err := fn(review); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userreviewDescCreatedAt is the schema descriptor for created_at field.
	userreviewDescCreatedAt := userreviewFields[4].Descriptor()
	// userreview.DefaultCreatedAt holds the default value on creation for the created_at field.
	userreview.DefaultCreatedAt = userreviewDescCreatedAt.Default.(func() time.Time)
	// userreviewDescUpdatedAt is the schema descriptor for updated_at field.
	userreviewDescUpdatedAt := userreviewFields[5].Descriptor()
	// userreview.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	userreview.DefaultUpdatedAt = userreviewDescUpdatedAt.Default.(func() time.Time)
	// userreview.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	userreview.UpdateDefaultUpdatedAt = userreviewDescUpdatedAt.UpdateDefault.(func() time.Time)
}

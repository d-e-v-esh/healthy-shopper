// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AddressesColumns holds the columns for the "addresses" table.
	AddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "phone_number", Type: field.TypeString, Size: 20},
		{Name: "address_line1", Type: field.TypeString, Size: 100},
		{Name: "address_line2", Type: field.TypeString, Nullable: true, Size: 100},
		{Name: "city", Type: field.TypeString, Size: 50},
		{Name: "state", Type: field.TypeString, Size: 50},
		{Name: "country", Type: field.TypeString, Size: 50},
		{Name: "postal_code", Type: field.TypeString, Size: 20},
	}
	// AddressesTable holds the schema information for the "addresses" table.
	AddressesTable = &schema.Table{
		Name:       "addresses",
		Columns:    AddressesColumns,
		PrimaryKey: []*schema.Column{AddressesColumns[0]},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "description", Type: field.TypeString, Size: 500},
		{Name: "product_image", Type: field.TypeString, Size: 500},
		{Name: "ingredients_list_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "product_category_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "nutritional_information_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "promotion_id", Type: field.TypeInt, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_promotions_promotion",
				Columns:    []*schema.Column{ProductsColumns[9]},
				RefColumns: []*schema.Column{PromotionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProductItemsColumns holds the columns for the "product_items" table.
	ProductItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "stock_keeping_unit", Type: field.TypeString, Size: 500},
		{Name: "quantity_in_stock", Type: field.TypeInt},
		{Name: "product_image", Type: field.TypeString, Size: 500},
		{Name: "price", Type: field.TypeFloat32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "product_id", Type: field.TypeInt, Unique: true},
	}
	// ProductItemsTable holds the schema information for the "product_items" table.
	ProductItemsTable = &schema.Table{
		Name:       "product_items",
		Columns:    ProductItemsColumns,
		PrimaryKey: []*schema.Column{ProductItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "product_items_products_product_item",
				Columns:    []*schema.Column{ProductItemsColumns[7]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PromotionsColumns holds the columns for the "promotions" table.
	PromotionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "description", Type: field.TypeString, Size: 500},
		{Name: "discount_percentage", Type: field.TypeInt},
		{Name: "start_date", Type: field.TypeTime},
		{Name: "end_date", Type: field.TypeTime},
	}
	// PromotionsTable holds the schema information for the "promotions" table.
	PromotionsTable = &schema.Table{
		Name:       "promotions",
		Columns:    PromotionsColumns,
		PrimaryKey: []*schema.Column{PromotionsColumns[0]},
	}
	// ShoppingCartsColumns holds the columns for the "shopping_carts" table.
	ShoppingCartsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ShoppingCartsTable holds the schema information for the "shopping_carts" table.
	ShoppingCartsTable = &schema.Table{
		Name:       "shopping_carts",
		Columns:    ShoppingCartsColumns,
		PrimaryKey: []*schema.Column{ShoppingCartsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "shopping_carts_users_shopping_cart",
				Columns:    []*schema.Column{ShoppingCartsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ShoppingCartItemsColumns holds the columns for the "shopping_cart_items" table.
	ShoppingCartItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "product_item_id", Type: field.TypeInt, Unique: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "shopping_cart_id", Type: field.TypeInt, Unique: true},
	}
	// ShoppingCartItemsTable holds the schema information for the "shopping_cart_items" table.
	ShoppingCartItemsTable = &schema.Table{
		Name:       "shopping_cart_items",
		Columns:    ShoppingCartItemsColumns,
		PrimaryKey: []*schema.Column{ShoppingCartItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "shopping_cart_items_shopping_carts_shopping_cart_item",
				Columns:    []*schema.Column{ShoppingCartItemsColumns[3]},
				RefColumns: []*schema.Column{ShoppingCartsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "email_address", Type: field.TypeString, Unique: true, Size: 60},
		{Name: "password", Type: field.TypeString, Size: 100},
		{Name: "first_name", Type: field.TypeString, Size: 50},
		{Name: "last_name", Type: field.TypeString, Size: 50},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserAddressesColumns holds the columns for the "user_addresses" table.
	UserAddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "is_default", Type: field.TypeBool, Default: false},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "address_id", Type: field.TypeInt},
	}
	// UserAddressesTable holds the schema information for the "user_addresses" table.
	UserAddressesTable = &schema.Table{
		Name:       "user_addresses",
		Columns:    UserAddressesColumns,
		PrimaryKey: []*schema.Column{UserAddressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_addresses_users_user_address",
				Columns:    []*schema.Column{UserAddressesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_addresses_addresses_address",
				Columns:    []*schema.Column{UserAddressesColumns[3]},
				RefColumns: []*schema.Column{AddressesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserReviewsColumns holds the columns for the "user_reviews" table.
	UserReviewsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ordered_product_id", Type: field.TypeInt, Unique: true},
		{Name: "rating", Type: field.TypeInt},
		{Name: "review", Type: field.TypeString, Size: 500},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
	}
	// UserReviewsTable holds the schema information for the "user_reviews" table.
	UserReviewsTable = &schema.Table{
		Name:       "user_reviews",
		Columns:    UserReviewsColumns,
		PrimaryKey: []*schema.Column{UserReviewsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_reviews_users_user_review",
				Columns:    []*schema.Column{UserReviewsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AddressesTable,
		ProductsTable,
		ProductItemsTable,
		PromotionsTable,
		ShoppingCartsTable,
		ShoppingCartItemsTable,
		UsersTable,
		UserAddressesTable,
		UserReviewsTable,
	}
)

func init() {
	ProductsTable.ForeignKeys[0].RefTable = PromotionsTable
	ProductItemsTable.ForeignKeys[0].RefTable = ProductsTable
	ShoppingCartsTable.ForeignKeys[0].RefTable = UsersTable
	ShoppingCartItemsTable.ForeignKeys[0].RefTable = ShoppingCartsTable
	UserAddressesTable.ForeignKeys[0].RefTable = UsersTable
	UserAddressesTable.ForeignKeys[1].RefTable = AddressesTable
	UserReviewsTable.ForeignKeys[0].RefTable = UsersTable
}

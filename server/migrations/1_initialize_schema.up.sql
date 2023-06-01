CREATE TABLE "User" (
  "UserID" int PRIMARY KEY,
  "Username" varchar(20) UNIQUE NOT NULL,
  "EmailAddress" varchar(60) UNIQUE NOT NULL,
  "Password" varchar(100) NOT NULL,
  "FirstName" varchar(50) NOT NULL,
  "LastName" varchar(50) NOT NULL,
  "CreatedAt" timestamptz DEFAULT (now()),
  "UpdatedAt" timestamptz
);

CREATE TABLE "UserAddress" (
  "UserID" int,
  "AddressID" int,
  "IsDefault" boolean
);

CREATE TABLE "Address" (
  "AddressID" int PRIMARY KEY,
  "PhoneNumber" int NOT NULL,
  "AddressLine1" varchar(100) NOT NULL,
  "AddressLine2" varchar(100) NOT NULL,
  "City" varchar(35) NOT NULL,
  "State" varchar(20) NOT NULL,
  "PINCode" varchar(10) NOT NULL
);

CREATE TABLE "Product" (
  "ProductID" int PRIMARY KEY,
  "ProductCategoryID" int NOT NULL,
  "Name" varchar(50) NOT NULL,
  "Description" varchar(500) NOT NULL,
  "ProductImage" varchar(500) NOT NULL,
  "IngredientsListID" int UNIQUE NOT NULL,
  "NutritionalInformationID" int UNIQUE NOT NULL,
  "PromotionID" int UNIQUE NOT NULL,
  "CreatedAt" timestamptz DEFAULT (now()),
  "UpdatedAt" timestamptz
);

CREATE TABLE "NutritionalInformation" (
  "NutritionalInformationID" int PRIMARY KEY,
  "NutritionalInformationListID" int UNIQUE NOT NULL,
  "NAmount" int,
  "NMeasuringUnit" varchar(6)
);

CREATE TABLE "NutritionalInformationList" (
  "NutritionalInformationListID" int PRIMARY KEY,
  "Parameter" varchar(20),
  "PerNAmount" float4,
  "MeasuringUnit" varchar(6)
);

CREATE TABLE "IngredientsList" (
  "IngredientsListID" int PRIMARY KEY,
  "Name" varchar(100),
  "Description" varchar(500)
);

CREATE TABLE "ProductItem" (
  "ProductItemID" int PRIMARY KEY,
  "ProductID" int UNIQUE NOT NULL,
  "StockKeepingUnit" varchar(100),
  "QuantityInStock" int,
  "ProductImage" varchar(500),
  "Price" float4
);

CREATE TABLE "ProductCategory" (
  "ProductCategoryID" int PRIMARY KEY,
  "ParentCategoryID" int UNIQUE NOT NULL,
  "CategoryName" varchar(20)
);

CREATE TABLE "ProductConfiguration" (
  "ProductItemID" int UNIQUE NOT NULL,
  "VariationOptionID" int UNIQUE NOT NULL
);

CREATE TABLE "Variation" (
  "VariationID" int PRIMARY KEY,
  "ProductCategoryID" int UNIQUE NOT NULL,
  "Name" varchar(20)
);

CREATE TABLE "VariationOption" (
  "VariationOptionID" int PRIMARY KEY,
  "VariationID" int UNIQUE NOT NULL,
  "Value" varchar(50)
);

CREATE TABLE "Promotion" (
  "PromotionID" int PRIMARY KEY,
  "Name" varchar(100),
  "Description" varchar(500),
  "DiscountRate" int,
  "StartDate" timestamptz,
  "EndDate" timestamptz
);

CREATE TABLE "ShoppingCart" (
  "ShoppingCartID" int PRIMARY KEY,
  "UserID" int UNIQUE NOT NULL
);

CREATE TABLE "ShoppingCartItem" (
  "ShoppingCartItemID" int PRIMARY KEY,
  "ShoppingCartID" int,
  "ProductItemID" int,
  "Quantity" int
);

CREATE TABLE "ShopOrder" (
  "ShopOrderID" int PRIMARY KEY,
  "UserID" int UNIQUE NOT NULL,
  "OrderDateAndTime" timestamptz DEFAULT (now()),
  "PaymentMethod" varchar(20),
  "ShippingAddressID" int UNIQUE NOT NULL,
  "ShippingMethodID" int UNIQUE NOT NULL,
  "OrderTotal" float4,
  "OrderStatusID" int UNIQUE NOT NULL
);

CREATE TABLE "ShippingAddress" (
  "ShippingAddressID" int PRIMARY KEY,
  "PhoneNumber" int,
  "AddressLine1" varchar(100),
  "AddressLine2" varchar(100),
  "City" varchar(35),
  "State" varchar(20),
  "PINCode" varchar(10)
);

CREATE TABLE "ShippingMethod" (
  "ShippingMethodID" int PRIMARY KEY,
  "Name" varchar(25),
  "Price" float4
);

CREATE TABLE "OrderStatus" (
  "OrderStatusID" int PRIMARY KEY,
  "Status" varchar(20)
);

CREATE TABLE "OrderLine" (
  "OrderLineID" int PRIMARY KEY,
  "ProductItemID" int UNIQUE NOT NULL,
  "ShopOrderID" int UNIQUE NOT NULL,
  "Quantity" int,
  "Price" float4
);

CREATE TABLE "UserReview" (
  "UserReviewID" int PRIMARY KEY,
  "UserID" int UNIQUE NOT NULL,
  "OrderedProductID" int UNIQUE NOT NULL,
  "RatingValue" int,
  "Comment" varchar(500)
);

ALTER TABLE "UserAddress" ADD FOREIGN KEY ("UserID") REFERENCES "User" ("UserID");

ALTER TABLE "UserAddress" ADD FOREIGN KEY ("AddressID") REFERENCES "Address" ("AddressID");

ALTER TABLE "ProductCategory" ADD FOREIGN KEY ("ParentCategoryID") REFERENCES "ProductCategory" ("ProductCategoryID");

ALTER TABLE "Product" ADD FOREIGN KEY ("ProductCategoryID") REFERENCES "ProductCategory" ("ProductCategoryID");

ALTER TABLE "Variation" ADD FOREIGN KEY ("ProductCategoryID") REFERENCES "ProductCategory" ("ProductCategoryID");

ALTER TABLE "ProductItem" ADD FOREIGN KEY ("ProductID") REFERENCES "Product" ("ProductID");

ALTER TABLE "VariationOption" ADD FOREIGN KEY ("VariationID") REFERENCES "Variation" ("VariationID");

ALTER TABLE "ProductConfiguration" ADD FOREIGN KEY ("VariationOptionID") REFERENCES "VariationOption" ("VariationOptionID");

ALTER TABLE "ProductConfiguration" ADD FOREIGN KEY ("ProductItemID") REFERENCES "ProductItem" ("ProductItemID");

ALTER TABLE "UserReview" ADD FOREIGN KEY ("UserID") REFERENCES "User" ("UserID");

ALTER TABLE "ShoppingCart" ADD FOREIGN KEY ("UserID") REFERENCES "User" ("UserID");

ALTER TABLE "ShoppingCartItem" ADD FOREIGN KEY ("ShoppingCartID") REFERENCES "ShoppingCart" ("ShoppingCartID");

ALTER TABLE "ShoppingCartItem" ADD FOREIGN KEY ("ProductItemID") REFERENCES "ProductItem" ("ProductItemID");

ALTER TABLE "OrderLine" ADD FOREIGN KEY ("ProductItemID") REFERENCES "ProductItem" ("ProductItemID");

ALTER TABLE "OrderLine" ADD FOREIGN KEY ("ShopOrderID") REFERENCES "ShopOrder" ("ShopOrderID");

ALTER TABLE "ShopOrder" ADD FOREIGN KEY ("ShippingMethodID") REFERENCES "ShippingMethod" ("ShippingMethodID");

ALTER TABLE "ShopOrder" ADD FOREIGN KEY ("OrderStatusID") REFERENCES "OrderStatus" ("OrderStatusID");

ALTER TABLE "UserReview" ADD FOREIGN KEY ("OrderedProductID") REFERENCES "OrderLine" ("OrderLineID");

ALTER TABLE "IngredientsList" ADD FOREIGN KEY ("IngredientsListID") REFERENCES "Product" ("IngredientsListID");

ALTER TABLE "NutritionalInformation" ADD FOREIGN KEY ("NutritionalInformationID") REFERENCES "Product" ("NutritionalInformationID");

ALTER TABLE "NutritionalInformationList" ADD FOREIGN KEY ("NutritionalInformationListID") REFERENCES "NutritionalInformation" ("NutritionalInformationListID");

ALTER TABLE "Promotion" ADD FOREIGN KEY ("PromotionID") REFERENCES "Product" ("PromotionID");

ALTER TABLE "ShippingAddress" ADD FOREIGN KEY ("ShippingAddressID") REFERENCES "ShopOrder" ("ShippingAddressID");

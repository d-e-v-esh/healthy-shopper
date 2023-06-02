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
  "UserID" int NOT NULL,
  "AddressID" int NOT NULL,
  "IsDefault" boolean NOT NULL
);

CREATE TABLE "Address" (
  "AddressID" int PRIMARY KEY,
  "PhoneNumber" varchar(20) NOT NULL,
  "AddressLine1" varchar(100) NOT NULL,
  "AddressLine2" varchar(100) NOT NULL,
  "City" varchar(35) NOT NULL,
  "State" varchar(20) NOT NULL,
  "PINCode" varchar(10) NOT NULL
);

CREATE TABLE "Product" (
  "ProductID" int PRIMARY KEY,
  "ProductCategoryID" int UNIQUE,
  "Name" varchar(50) NOT NULL,
  "Description" varchar(500) NOT NULL,
  "ProductImage" varchar(500) NOT NULL,
  "IngredientsListID" int UNIQUE,
  "NutritionalInformationID" int UNIQUE,
  "PromotionID" int UNIQUE,
  "CreatedAt" timestamptz DEFAULT (now()),
  "UpdatedAt" timestamptz
);

CREATE TABLE "NutritionalInformation" (
  "NutritionalInformationID" int PRIMARY KEY,
  "NutritionalInformationListID" int UNIQUE NOT NULL,
  "NAmount" int NOT NULL,
  "NMeasuringUnit" varchar(6) NOT NULL
);

CREATE TABLE "NutritionalInformationList" (
  "NutritionalInformationListID" int PRIMARY KEY,
  "Parameter" varchar(20) NOT NULL,
  "PerNAmount" float4 NOT NULL,
  "MeasuringUnit" varchar(6) NOT NULL
);

CREATE TABLE "IngredientsList" (
  "IngredientsListID" int PRIMARY KEY,
  "Name" varchar(100) NOT NULL,
  "Description" varchar(500) NOT NULL
);

CREATE TABLE "ProductItem" (
  "ProductItemID" int PRIMARY KEY,
  "ProductID" int UNIQUE NOT NULL,
  "StockKeepingUnit" varchar(100) NOT NULL,
  "QuantityInStock" int NOT NULL,
  "ProductImage" varchar(500) NOT NULL,
  "Price" float4 NOT NULL
);

CREATE TABLE "ProductCategory" (
  "ProductCategoryID" int PRIMARY KEY,
  "ParentCategoryID" int UNIQUE NOT NULL,
  "CategoryName" varchar(20) NOT NULL
);

CREATE TABLE "ProductConfiguration" (
  "ProductItemID" int UNIQUE NOT NULL,
  "VariationOptionID" int UNIQUE NOT NULL
);

CREATE TABLE "Variation" (
  "VariationID" int PRIMARY KEY,
  "ProductCategoryID" int UNIQUE,
  "Name" varchar(20) NOT NULL
);

CREATE TABLE "VariationOption" (
  "VariationOptionID" int PRIMARY KEY,
  "VariationID" int UNIQUE NOT NULL,
  "Value" varchar(50) NOT NULL
);

CREATE TABLE "Promotion" (
  "PromotionID" int PRIMARY KEY,
  "Name" varchar(100) NOT NULL,
  "Description" varchar(500) NOT NULL,
  "DiscountRate" int NOT NULL,
  "StartDate" timestamptz NOT NULL,
  "EndDate" timestamptz NOT NULL
);

CREATE TABLE "ShoppingCart" (
  "ShoppingCartID" int PRIMARY KEY,
  "UserID" int UNIQUE NOT NULL
);

CREATE TABLE "ShoppingCartItem" (
  "ShoppingCartItemID" int PRIMARY KEY,
  "ShoppingCartID" int UNIQUE NOT NULL,
  "ProductItemID" int UNIQUE NOT NULL,
  "Quantity" int NOT NULL
);

CREATE TABLE "ShopOrder" (
  "ShopOrderID" int PRIMARY KEY,
  "UserID" int UNIQUE NOT NULL,
  "OrderDateAndTime" timestamptz NOT NULL DEFAULT (now()),
  "PaymentMethod" varchar(20) NOT NULL,
  "ShippingAddressID" int UNIQUE NOT NULL,
  "ShippingMethodID" int UNIQUE NOT NULL,
  "OrderTotal" float4 NOT NULL,
  "OrderStatusID" int UNIQUE NOT NULL
);

CREATE TABLE "ShippingAddress" (
  "ShippingAddressID" int PRIMARY KEY,
  "PhoneNumber" varchar(20) NOT NULL,
  "AddressLine1" varchar(100) NOT NULL,
  "AddressLine2" varchar(100) NOT NULL,
  "City" varchar(35) NOT NULL,
  "State" varchar(20) NOT NULL,
  "PINCode" varchar(10) NOT NULL
);

CREATE TABLE "ShippingMethod" (
  "ShippingMethodID" int PRIMARY KEY,
  "Name" varchar(25) NOT NULL,
  "Price" float4 NOT NULL
);

CREATE TABLE "OrderStatus" (
  "OrderStatusID" int PRIMARY KEY,
  "Status" varchar(20) NOT NULL
);

CREATE TABLE "OrderLine" (
  "OrderLineID" int PRIMARY KEY,
  "ProductItemID" int UNIQUE,
  "ShopOrderID" int UNIQUE,
  "Quantity" int NOT NULL,
  "Price" float4 NOT NULL
);

CREATE TABLE "UserReview" (
  "UserReviewID" int PRIMARY KEY,
  "UserID" int UNIQUE NOT NULL,
  "OrderedProductID" int UNIQUE NOT NULL,
  "RatingValue" int NOT NULL,
  "Comment" varchar(500) NOT NULL
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

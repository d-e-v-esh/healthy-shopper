CREATE TABLE "User" (
  "UserID" int PRIMARY KEY,
  "Username" varchar(20),
  "EmailAddress" varchar(60),
  "PhoneNumber" int,
  "Password" varchar(100),
  "FirstName" varchar(50),
  "LastName" varchar(50),
  "CreatedAt" timestamp
);

CREATE TABLE "UserAddress" (
  "UserID" int,
  "AddressID" int,
  "IsDefault" boolean
);

CREATE TABLE "Address" (
  "AddressID" int PRIMARY KEY,
  "AddressLine1" varchar(100),
  "AddressLine2" varchar(100),
  "City" varchar(35),
  "State" varchar(20),
  "PINCode" varchar(10)
);

CREATE TABLE "Product" (
  "ProductID" int PRIMARY KEY,
  "ProductCategoryID" int,
  "Name" varchar(50),
  "Description" varchar(500),
  "ProductImage" varchar(500),
  "IngredientsListID" int,
  "NutritionalInformationID" int,
  "PromotionID" int
);

CREATE TABLE "NutritionalInformation" (
  "NutritionalInformationID" int PRIMARY KEY,
  "NutritionalInformationListID" int,
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
  "ProductID" int,
  "StockKeepingUnit" varchar(100),
  "QuantityInStock" int,
  "ProductImage" varchar(500),
  "Price" float4
);

CREATE TABLE "ProductCategory" (
  "ProductCategoryID" int PRIMARY KEY,
  "ParentCategoryID" int,
  "CategoryName" varchar(20)
);

CREATE TABLE "ProductConfiguration" (
  "ProductItemID" int,
  "VariationOptionID" int
);

CREATE TABLE "Variation" (
  "VariationID" int PRIMARY KEY,
  "ProductCategoryID" int,
  "Name" varchar(20)
);

CREATE TABLE "VariationOption" (
  "VariationOptionID" int PRIMARY KEY,
  "VariationID" int,
  "Value" varchar(50)
);

CREATE TABLE "Promotion" (
  "PromotionID" int PRIMARY KEY,
  "Name" varchar(100),
  "Description" varchar(500),
  "DiscountRate" int,
  "StartDate" timestamp,
  "EndDate" timestamp
);

CREATE TABLE "ShoppingCart" (
  "ShoppingCartID" int PRIMARY KEY,
  "UserID" int
);

CREATE TABLE "ShoppingCartItem" (
  "ShoppingCartItemID" int PRIMARY KEY,
  "ShoppingCartID" int,
  "ProductItemID" int,
  "Quantity" int
);

CREATE TABLE "ShopOrder" (
  "ShopOrderID" int PRIMARY KEY,
  "UserID" int,
  "OrderDate" timestamp,
  "PaymentMethod" varchar(20),
  "ShippingAddressID" int,
  "ShippingMethodID" int,
  "OrderTotal" float4,
  "OrderStatusID" int
);

CREATE TABLE "ShippingAddress" (
  "ShippingAddressID" int PRIMARY KEY,
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
  "ProductItemID" int,
  "ShopOrderID" int,
  "Quantity" int,
  "Price" float4
);

CREATE TABLE "UserReview" (
  "UserReviewID" int PRIMARY KEY,
  "UserID" int,
  "OrderedProductID" int,
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

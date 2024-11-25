DROP TABLE IF EXISTS SuppliedProduct;
DROP TABLE IF EXISTS Supply;
DROP TABLE IF EXISTS Supplier;
DROP TABLE IF EXISTS WarehouseProduct;
DROP TABLE IF EXISTS OrderStep;
DROP TABLE IF EXISTS ProductInOrder;
DROP TABLE IF EXISTS "Order";
DROP TABLE IF EXISTS ShoppingCartProduct;
DROP TABLE IF EXISTS ShoppingCart;
DROP TABLE IF EXISTS Warehouse;
DROP TABLE IF EXISTS Favorites;
DROP TABLE IF EXISTS Employee;
DROP TABLE IF EXISTS Attachment;
DROP TABLE IF EXISTS "Comment";
DROP TABLE IF EXISTS Review;
DROP TABLE IF EXISTS ProductCost;
DROP TABLE IF EXISTS ProductCharacteristic;
DROP TABLE IF EXISTS Characteristic;
DROP TABLE IF EXISTS ProductCategory;
DROP TABLE IF EXISTS Category;
DROP TABLE IF EXISTS Product;
DROP TABLE IF EXISTS Manufacturer;
DROP TABLE IF EXISTS ClientAddress;
DROP TABLE IF EXISTS AuthorizedClient;
DROP TABLE IF EXISTS "User";
DROP TABLE IF EXISTS "Role";

DROP DOMAIN IF EXISTS StatusDomain;
CREATE DOMAIN StatusDomain AS VARCHAR(10)
CHECK (VALUE IN ('ACTIVE', 'DELETED'));

-- Создание таблицы ролей (Role)
CREATE TABLE "Role" (
    RoleId SERIAL PRIMARY KEY,
    RoleName VARCHAR(50) NOT NULL,
    AccessLevel INT NOT NULL
);

-- Создание таблицы пользователей (User)
CREATE TABLE "User" (
    UserId SERIAL PRIMARY KEY,
    Login VARCHAR(50) NOT NULL,
    PassHash VARCHAR(255) NOT NULL,
    FirstName VARCHAR(50) NOT NULL,
    SecondName VARCHAR(50),
    SurName VARCHAR(50),
    IconLink VARCHAR(255),
    Phone VARCHAR(20) UNIQUE,
    Email VARCHAR(100) UNIQUE NOT NULL,
    BirthdayDate DATE,
	Status StatusDomain DEFAULT 'ACTIVE',
    RoleId INT NOT NULL REFERENCES "Role"(RoleId)
);

-- Создание таблицы клиентов
CREATE TABLE AuthorizedClient(
	UserId SERIAL PRIMARY KEY,
	Description TEXT,
	Town VARCHAR(50),
	Rating NUMERIC(2, 1),
	FOREIGN KEY (UserId) REFERENCES "User"(UserId)
);

-- Создание таблицы адресов клиентов
CREATE TABLE ClientAddress (
    AddressId SERIAL PRIMARY KEY,
    UserId INT NOT NULL REFERENCES AuthorizedClient,
    "Name" VARCHAR(100),
    Town VARCHAR(100),
    Street VARCHAR(100),
    Building VARCHAR(20),
    "Floor" INT,
    Apartment INT,
    IntercomCode VARCHAR(20),
    "Comment" TEXT
);

-- Создание таблицы производителей (Manufacturer)
CREATE TABLE Manufacturer (
    ManufacturerId SERIAL PRIMARY KEY,
    "Name" VARCHAR(100) NOT NULL,
    Description TEXT
);

-- Создание таблицы продуктов (Product)
CREATE TABLE Product (
    ProductId SERIAL PRIMARY KEY,
    "Name" VARCHAR(100) NOT NULL,
    Status StatusDomain NOT NULL DEFAULT 'ACTIVE',
    Description TEXT,
    ManufacturerId INT REFERENCES Manufacturer(ManufacturerId)
);

-- Создание таблицы категорий продуктов (Category)
CREATE TABLE Category (
    CategoryId SERIAL PRIMARY KEY,
    "Name" VARCHAR(50) NOT NULL,
    Description TEXT
);

-- Создание таблицы связей продукт-категория 
CREATE TABLE ProductCategory (
    ProductId INT REFERENCES Product(ProductId) ON DELETE CASCADE,
    CategoryId INT REFERENCES Category(CategoryId) ON DELETE CASCADE,
    PRIMARY KEY (ProductId, CategoryId)
);

-- Создание таблицы характеристик продуктов (Characteristic)
CREATE TABLE Characteristic (
    CharacteristicId SERIAL PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Category VARCHAR(50)
);

-- Создание таблицы связей продукт-характеристика (ProductCharacteristic)
CREATE TABLE ProductCharacteristic (
    ProductId INT REFERENCES Product(ProductId) ON DELETE CASCADE,
    CharacteristicId INT REFERENCES Characteristic(CharacteristicId) ON DELETE CASCADE,
    Value VARCHAR(255),
    PRIMARY KEY (ProductId, CharacteristicId)
);

-- Создание таблицы цены продукта (ProductCost)
CREATE TABLE ProductCost (
    CostId SERIAL PRIMARY KEY,
    ProductId INT REFERENCES Product(ProductId) ON DELETE CASCADE,
    Value NUMERIC(10, 2) NOT NULL,
    Currency VARCHAR(10) NOT NULL,
    StartTimeStamp TIMESTAMP DEFAULT NOW()
);

-- Создание таблицы отзывов (Review)
CREATE TABLE Review (
    ReviewId SERIAL PRIMARY KEY,
    UserId INT REFERENCES "User"(UserId),
    Score INT CHECK (Score >= 1 AND Score <= 5),
    Description TEXT,
    TimeStamp TIMESTAMP DEFAULT NOW(),
    ProductId INT REFERENCES Product(ProductId)
);

-- Создание таблицы комментариев (Comment)
CREATE TABLE "Comment" (
    CommentId SERIAL PRIMARY KEY,
    UserId INT REFERENCES AuthorizedClient(UserId),
    ReviewId INT REFERENCES Review(ReviewId),
    "Text" TEXT,
    TimeStamp TIMESTAMP DEFAULT NOW()
);

-- Создание таблицы вложений к отзывам (Attachment)
CREATE TABLE Attachment (
    Attachment_Link VARCHAR(255) PRIMARY KEY,
    ReviewId INT REFERENCES Review(ReviewId),
	Description TEXT
);

CREATE TABLE Favorites(
	ProductId INT REFERENCES Product(ProductId) ON DELETE CASCADE,
    UserId INT REFERENCES AuthorizedClient(UserId) ON DELETE CASCADE,
    PRIMARY KEY (ProductId, UserId)
);

-- Создание таблицы корзин покупок (ShoppingCart)
CREATE TABLE ShoppingCart (
    CartId SERIAL PRIMARY KEY,
    ClientId INT REFERENCES AuthorizedClient(UserId)
);

-- Создание таблицы продуктов в корзине (ShoppingCartProduct)
CREATE TABLE ShoppingCartProduct (
    CartId INT REFERENCES ShoppingCart(CartId) ON DELETE CASCADE,
    ProductId INT REFERENCES Product(ProductId) ON DELETE CASCADE,
    Count INT NOT NULL,
    IsSelected BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (CartId, ProductId)
);

-- Создание таблицы заказов (Order)
CREATE TABLE "Order" (
    OrderId SERIAL PRIMARY KEY,
    Status VARCHAR(50) NOT NULL,
    TimeStamp TIMESTAMP DEFAULT NOW(),
    CartId INT REFERENCES ShoppingCart(CartId),
    AddressId INT REFERENCES ClientAddress(AddressId)
);

-- Создание таблицы сотрудников (Employee)
CREATE TABLE Employee (
    EmployeeId SERIAL PRIMARY KEY,
    Department VARCHAR(100),
    Vacancy VARCHAR(100),
    Grade VARCHAR(50),
    Salary NUMERIC(10, 2),
	FOREIGN KEY (EmployeeId) REFERENCES "User"(UserId)
);

-- Создание таблицы шагов заказа (OrderStep)
CREATE TABLE OrderStep (
    StepId SERIAL PRIMARY KEY,
    StepName VARCHAR(100) NOT NULL,
    OrderId INT REFERENCES "Order"(OrderId) ON DELETE CASCADE,
    TimeStampStart TIMESTAMP,
    TimeStampEnd TIMESTAMP,
    Description TEXT,
    Status VARCHAR(50),
    ResponsibleId INT REFERENCES Employee(EmployeeId)
);

-- Создание таблицы складов (Warehouse)
CREATE TABLE Warehouse (
    WarehouseId SERIAL PRIMARY KEY,
    "Name" VARCHAR(100) NOT NULL,
    Address TEXT,
    "Value" NUMERIC(10, 2)
);

-- Создание таблицы продуктов в корзине (ShoppingCartProduct)
CREATE TABLE ProductInOrder (
    OrderId INT REFERENCES "Order"(OrderId) ON DELETE CASCADE,
    ProductId INT REFERENCES Product(ProductId) ON DELETE CASCADE,
    "Count" INT NOT NULL,
    WarehouseId INT NOT NULL REFERENCES Warehouse(WarehouseId),
    PRIMARY KEY(OrderId, ProductId)
);

-- Создание таблицы продуктов на складе (WarehouseProduct)
CREATE TABLE WarehouseProduct (
    ProductId INT REFERENCES Product(ProductId) ON DELETE CASCADE,
    WarehouseId INT REFERENCES Warehouse(WarehouseId) ON DELETE CASCADE,
    "Count" INT NOT NULL,
    "Section" VARCHAR(50),
    Shelf VARCHAR(50),
    PRIMARY KEY (ProductId, WarehouseId)
);

-- Создание таблицы поставщиков (Supplier)
CREATE TABLE Supplier (
    SupplierId SERIAL PRIMARY KEY,
    "Name" VARCHAR(100) NOT NULL,
    Phone VARCHAR(20),
    Address TEXT,
    ContactName VARCHAR(100),
    DirectorName VARCHAR(100)
);

-- Создание таблицы поставок (Supply)
CREATE TABLE Supply (
    SupplyId SERIAL PRIMARY KEY,
    SupplierId INT REFERENCES Supplier(SupplierId) ON DELETE CASCADE,
    TimeStamp TIMESTAMP DEFAULT NOW(),
    "Count" INT NOT NULL,
    WarehouseId INT REFERENCES Warehouse(WarehouseId),
    ResponsibleId INT REFERENCES Employee(EmployeeId)
);

-- Создание таблицы продуктов в поставке (SuppliedProduct)
CREATE TABLE SuppliedProduct (
    ProductId INT REFERENCES Product(ProductId) ON DELETE CASCADE,
    SupplyId INT REFERENCES Supply(SupplyId) ON DELETE CASCADE,
    "Count" INT NOT NULL,
    Price NUMERIC(10, 2),
    PRIMARY KEY (ProductId, SupplyId)
);








-- Заполнение таблицы ролей
INSERT INTO "Role" (RoleName, AccessLevel) VALUES 
('Admin', 10),
('User', 1),
('Manager', 5);

-- Заполнение таблицы пользователей
INSERT INTO "User" (Login, PassHash, FirstName, SecondName, SurName, Phone, Email, BirthdayDate, Status, RoleId) VALUES
('admin', 'hashedpassword123', 'John', 'Admin', 'Doe', '1234567890', 'admin@example.com', '1990-01-01', 'ACTIVE', 1),
('user1', 'hashedpassword456', 'Jane', NULL, 'Smith', '0987654321', 'jane.smith@example.com', '1995-05-15', 'ACTIVE', 2),
('logistics_manager', 'hashedpassword789', 'Alice', 'Logistics', 'Brown', '0987654342', 'alice.brown@example.com', '1985-06-10', 'ACTIVE', 3),
('sales_rep', 'hashedpassword012', 'Bob', 'Sales', 'Green', '0987542321', 'bob.green@example.com', '1992-11-20', 'ACTIVE', 3);

INSERT INTO Employee (EmployeeId, Department, Vacancy, Grade, Salary) VALUES
(1, 'Logistics', 'Warehouse Manager', 'Senior', 5000.00),
(2, 'Sales', 'Sales Representative', 'Junior', 3000.00);

-- Заполнение таблицы авторизованных клиентов
INSERT INTO AuthorizedClient (UserId, Description, Town, Rating) VALUES
(2, 'Regular customer', 'New York', 4.5);

-- Заполнение таблицы адресов клиентов
INSERT INTO ClientAddress (UserId, "Name", Town, Street, Building, "Floor", Apartment, IntercomCode, "Comment") VALUES
(2, 'Home', 'New York', '5th Avenue', '10', 2, 101, '1234', 'Delivery instructions');

-- Заполнение таблицы производителей
INSERT INTO Manufacturer ("Name", Description) VALUES
('TechCorp', 'Leading tech manufacturer'),
('HomeGoods', 'Household products');

-- Заполнение таблицы продуктов
INSERT INTO Product ("Name", Status, Description, ManufacturerId) VALUES
('Laptop', 'ACTIVE', 'High-performance laptop', 1),
('Vacuum Cleaner', 'ACTIVE', 'Powerful vacuum cleaner', 2);

-- Заполнение таблицы категорий
INSERT INTO Category ("Name", Description) VALUES
('Electronics', 'Electronic devices and gadgets'),
('Home Appliances', 'Appliances for household use');

-- Заполнение связей продукт-категория
INSERT INTO ProductCategory (ProductId, CategoryId) VALUES
(1, 1),
(2, 2);

-- Заполнение таблицы характеристик
INSERT INTO Characteristic (Name, Category) VALUES
('Processor', 'Electronics'),
('Power', 'Home Appliances');

-- Заполнение характеристик продукта
INSERT INTO ProductCharacteristic (ProductId, CharacteristicId, Value) VALUES
(1, 1, 'Intel i7'),
(2, 2, '1200W');

-- Заполнение таблицы цен продуктов
INSERT INTO ProductCost (ProductId, Value, Currency) VALUES
(1, 1200.00, 'USD'),
(2, 200.00, 'USD');

-- Заполнение таблицы складов
INSERT INTO Warehouse ("Name", Address, "Value") VALUES
('Central Warehouse', '123 Warehouse St, NY', 50000.00),
('Secondary Warehouse', '456 Storage Ln, LA', 30000.00);

-- Заполнение продуктов на складе
INSERT INTO WarehouseProduct (ProductId, WarehouseId, "Count", "Section", Shelf) VALUES
(1, 1, 10, 'A', '1'),
(2, 2, 15, 'B', '3');

-- Заполнение таблицы поставщиков
INSERT INTO Supplier ("Name", Phone, Address, ContactName, DirectorName) VALUES
('Tech Supplies Inc.', '1112223333', '789 Supplier Rd, NY', 'Alice Brown', 'Bob Green');

-- Заполнение таблицы поставок
INSERT INTO Supply (SupplierId, TimeStamp, "Count", WarehouseId, ResponsibleId) VALUES
(1, NOW(), 20, 1, 1);

-- Заполнение продуктов в поставке
INSERT INTO SuppliedProduct (ProductId, SupplyId, "Count", Price) VALUES
(1, 1, 10, 1100.00),
(2, 1, 10, 190.00);

-- Заполнение таблицы корзин покупок
INSERT INTO ShoppingCart (ClientId) VALUES
(2);

-- Заполнение продуктов в корзине
INSERT INTO ShoppingCartProduct (CartId, ProductId, "Count", IsSelected) VALUES
(1, 1, 1, TRUE),
(1, 2, 2, FALSE);

-- Заполнение таблицы заказов
INSERT INTO "Order" (Status, CartId, AddressId) VALUES
('Pending', 1, 1);

-- Заполнение шагов заказа
INSERT INTO OrderStep (StepName, OrderId, TimeStampStart, Description, Status, ResponsibleId) VALUES
('Order Received', 1, NOW(), 'Order has been received', 'In Progress', 1);

-- Заполнение отзывов
INSERT INTO Review (UserId, Score, Description, ProductId) VALUES
(2, 5, 'Great product!', 1);

-- Заполнение комментариев
INSERT INTO "Comment" (UserId, ReviewId, "Text") VALUES
(2, 1, 'I completely agree!');

-- Заполнение избранного
INSERT INTO Favorites (ProductId, UserId) VALUES
(1, 2);

INSERT INTO Attachment (Attachment_Link, ReviewId, Description) VALUES
('https://example.com/review1/image1.jpg', 1, 'Фото товара с первого отзыва'),
('https://example.com/review2/image2.jpg', 1, 'Фото товара со второго отзыва'),
('https://example.com/review3/image3.jpg', 1, 'Скриншот ошибки товара'),
('https://example.com/review4/image4.jpg', 1, 'Фотография упаковки');

-- Заполнение таблицы ProductInOrder
INSERT INTO ProductInOrder (OrderId, ProductId, "Count", WarehouseId) VALUES
(1, 1, 2, 1), -- Заказ №1, продукт №101, количество 2, со склада №1
(1, 2, 1, 2); -- Заказ №1, продукт №102, количество 1, со склада №2
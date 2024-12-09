-- Заполнение таблицы ролей
INSERT INTO "role" (role_name, access_level) VALUES 
('Admin', 10),
('User', 1),
('Manager', 5);

-- Заполнение таблицы пользователей
INSERT INTO "user" (login, pass_hash, first_name, second_name, sur_name, phone, email, birthday_date, status, role_id) VALUES
('admin', 'hashedpassword123', 'John', 'Admin', 'Doe', '1234567890', 'admin@example.com', '1990-01-01', 'ACTIVE', 1),
('user1', 'hashedpassword456', 'Jane', NULL, 'Smith', '0987654321', 'jane.smith@example.com', '1995-05-15', 'ACTIVE', 2),
('logistics_manager', 'hashedpassword789', 'Alice', 'Logistics', 'Brown', '0987654342', 'alice.brown@example.com', '1985-06-10', 'ACTIVE', 3),
('sales_rep', 'hashedpassword012', 'Bob', 'Sales', 'Green', '0987542321', 'bob.green@example.com', '1992-11-20', 'ACTIVE', 3);

-- Заполнение таблицы сотрудников
INSERT INTO employee (employee_id, department, vacancy, grade, salary) VALUES
(1, 'Logistics', 'Warehouse Manager', 'Senior', 5000.00),
(2, 'Sales', 'Sales Representative', 'Junior', 3000.00);

-- Заполнение таблицы авторизованных клиентов
INSERT INTO authorized_client (user_id, description, town, rating) VALUES
(2, 'Regular customer', 'New York', 4.5);

-- Заполнение таблицы адресов клиентов
INSERT INTO client_address (user_id, name, town, street, building, floor, apartment, intercom_code, comment) VALUES
(2, 'Home', 'New York', '5th Avenue', '10', 2, 101, '1234', 'Delivery instructions');

-- Заполнение таблицы производителей
INSERT INTO manufacturer (name, description) VALUES
('TechCorp', 'Leading tech manufacturer'),
('HomeGoods', 'Household products');

-- Заполнение таблицы продуктов
INSERT INTO product (name, status, description,image_link, manufacturer_id) VALUES
('Laptop', 'ACTIVE', 'High-performance laptop', '/image_0001.png', 1),
('Vacuum Cleaner', 'ACTIVE', 'Powerful vacuum cleaner', '/image_0002.png', 2);

-- Заполнение таблицы категорий
INSERT INTO category (name, description) VALUES
('Electronics', 'Electronic devices and gadgets'),
('Home Appliances', 'Appliances for household use');

-- Заполнение связей продукт-категория
INSERT INTO product_category (product_id, category_id) VALUES
(1, 1),
(2, 2);

-- Заполнение таблицы характеристик
INSERT INTO characteristic (name, category) VALUES
('Processor', 'Electronics'),
('Power', 'Home Appliances');

-- Заполнение характеристик продукта
INSERT INTO product_characteristic (product_id, characteristic_id, value) VALUES
(1, 1, 'Intel i7'),
(2, 2, '1200W');

-- Заполнение таблицы цен продуктов
INSERT INTO product_cost (product_id, value, currency) VALUES
(1, 1200.00, 'USD'),
(2, 200.00, 'USD');

-- Заполнение таблицы складов
INSERT INTO warehouse (name, address, value) VALUES
('Central Warehouse', '123 Warehouse St, NY', 50000.00),
('Secondary Warehouse', '456 Storage Ln, LA', 30000.00);

-- Заполнение продуктов на складе
INSERT INTO warehouse_product (product_id, warehouse_id, count, section, shelf) VALUES
(1, 1, 10, 'A', '1'),
(2, 2, 15, 'B', '3');

-- Заполнение таблицы поставщиков
INSERT INTO supplier (name, phone, address, contact_name, director_name) VALUES
('Tech Supplies Inc.', '1112223333', '789 Supplier Rd, NY', 'Alice Brown', 'Bob Green');

-- Заполнение таблицы поставок
INSERT INTO supply (supplier_id, time_stamp, count, warehouse_id, responsible_id) VALUES
(1, NOW(), 20, 1, 1);

-- Заполнение продуктов в поставке
INSERT INTO supplied_product (product_id, supply_id, count, price) VALUES
(1, 1, 10, 1100.00),
(2, 1, 10, 190.00);

-- Заполнение таблицы корзин покупок
INSERT INTO shopping_cart (client_id) VALUES
(2);

-- Заполнение продуктов в корзине
INSERT INTO shopping_cart_product (cart_id, product_id, count, is_selected) VALUES
(1, 1, 1, TRUE),
(1, 2, 2, FALSE);

-- Заполнение таблицы заказов
INSERT INTO "order" (status, cart_id, address_id) VALUES
('Pending', 1, 1);

-- Заполнение шагов заказа
INSERT INTO order_step (step_name, order_id, time_stamp_start, description, status, responsible_id) VALUES
('Order Received', 1, NOW(), 'Order has been received', 'In Progress', 1);

-- Заполнение отзывов
INSERT INTO review (user_id, score, description, product_id) VALUES
(2, 5, 'Great product!', 1);

-- Заполнение комментариев
INSERT INTO "comment" (user_id, review_id, text) VALUES
(2, 1, 'I completely agree!');

-- Заполнение избранного
INSERT INTO favorites (product_id, user_id) VALUES
(1, 2);

-- Заполнение таблицы вложений
INSERT INTO attachment (attachment_link, review_id, description) VALUES
('https://example.com/review1/image1.jpg', 1, 'Фото товара с первого отзыва'),
('https://example.com/review2/image2.jpg', 1, 'Фото товара со второго отзыва'),
('https://example.com/review3/image3.jpg', 1, 'Скриншот ошибки товара'),
('https://example.com/review4/image4.jpg', 1, 'Фотография упаковки');

-- Заполнение таблицы product_in_order
INSERT INTO product_in_order (order_id, product_id, count, warehouse_id) VALUES
(1, 1, 2, 1),
(1, 2, 1, 2);

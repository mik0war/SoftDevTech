DROP TABLE IF EXISTS supplied_product;
DROP TABLE IF EXISTS supply;
DROP TABLE IF EXISTS supplier;
DROP TABLE IF EXISTS warehouse_product;
DROP TABLE IF EXISTS order_step;
DROP TABLE IF EXISTS product_in_order;
DROP TABLE IF EXISTS "order";
DROP TABLE IF EXISTS shopping_cart_product;
DROP TABLE IF EXISTS shopping_cart;
DROP TABLE IF EXISTS warehouse;
DROP TABLE IF EXISTS favorites;
DROP TABLE IF EXISTS employee;
DROP TABLE IF EXISTS attachment;
DROP TABLE IF EXISTS "comment";
DROP TABLE IF EXISTS review;
DROP TABLE IF EXISTS product_cost;
DROP TABLE IF EXISTS product_characteristic;
DROP TABLE IF EXISTS characteristic;
DROP TABLE IF EXISTS product_category;
DROP TABLE IF EXISTS category;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS manufacturer;
DROP TABLE IF EXISTS client_address;
DROP TABLE IF EXISTS authorized_client;
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "role";

DROP DOMAIN IF EXISTS status_domain;
CREATE DOMAIN status_domain AS VARCHAR(10)
CHECK (VALUE IN ('ACTIVE', 'DELETED'));

-- Создание таблицы ролей (role)
CREATE TABLE "role" (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL,
    access_level INT NOT NULL
);

-- Создание таблицы пользователей (user)
CREATE TABLE "user" (
    user_id SERIAL PRIMARY KEY,
    login VARCHAR(50) NOT NULL,
    pass_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    second_name VARCHAR(50),
    sur_name VARCHAR(50),
    icon_link VARCHAR(255),
    phone VARCHAR(20) UNIQUE,
    email VARCHAR(100) UNIQUE NOT NULL,
    birthday_date DATE,
    status status_domain DEFAULT 'ACTIVE',
    role_id INT NOT NULL REFERENCES "role"(role_id)
);

-- Создание таблицы клиентов
CREATE TABLE authorized_client(
    user_id SERIAL PRIMARY KEY,
    description TEXT,
    town VARCHAR(50),
    rating NUMERIC(2, 1),
    FOREIGN KEY (user_id) REFERENCES "user"(user_id)
);

-- Создание таблицы адресов клиентов
CREATE TABLE client_address (
    address_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES authorized_client(user_id),
    name VARCHAR(100),
    town VARCHAR(100),
    street VARCHAR(100),
    building VARCHAR(20),
    floor INT,
    apartment INT,
    intercom_code VARCHAR(20),
    comment TEXT
);

-- Создание таблицы производителей (manufacturer)
CREATE TABLE manufacturer (
    manufacturer_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

-- Создание таблицы продуктов (product)
CREATE TABLE product (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    status status_domain NOT NULL DEFAULT 'ACTIVE',
    description TEXT,
	image_link VARCHAR(255),
    manufacturer_id INT REFERENCES manufacturer(manufacturer_id)
);

-- Создание таблицы категорий продуктов (category)
CREATE TABLE category (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT
);

-- Создание таблицы связей продукт-категория
CREATE TABLE product_category (
    product_id INT REFERENCES product(product_id) ON DELETE CASCADE,
    category_id INT REFERENCES category(category_id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, category_id)
);

-- Создание таблицы характеристик продуктов (characteristic)
CREATE TABLE characteristic (
    characteristic_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(50)
);

-- Создание таблицы связей продукт-характеристика (product_characteristic)
CREATE TABLE product_characteristic (
    product_id INT REFERENCES product(product_id) ON DELETE CASCADE,
    characteristic_id INT REFERENCES characteristic(characteristic_id) ON DELETE CASCADE,
    value VARCHAR(255),
    PRIMARY KEY (product_id, characteristic_id)
);

-- Создание таблицы цены продукта (product_cost)
CREATE TABLE product_cost (
    cost_id SERIAL PRIMARY KEY,
    product_id INT REFERENCES product(product_id) ON DELETE CASCADE,
    value NUMERIC(10, 2) NOT NULL,
    currency VARCHAR(10) NOT NULL,
    start_time_stamp TIMESTAMP DEFAULT NOW()
);

-- Создание таблицы отзывов (review)
CREATE TABLE review (
    review_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES "user"(user_id),
    score INT CHECK (score >= 1 AND score <= 5),
    description TEXT,
    time_stamp TIMESTAMP DEFAULT NOW(),
    product_id INT REFERENCES product(product_id)
);

-- Создание таблицы комментариев (comment)
CREATE TABLE "comment" (
    comment_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES authorized_client(user_id),
    review_id INT REFERENCES review(review_id),
    text TEXT,
    time_stamp TIMESTAMP DEFAULT NOW()
);

-- Создание таблицы вложений к отзывам (attachment)
CREATE TABLE attachment (
    attachment_link VARCHAR(255) PRIMARY KEY,
    review_id INT REFERENCES review(review_id),
    description TEXT
);

CREATE TABLE favorites(
    product_id INT REFERENCES product(product_id) ON DELETE CASCADE,
    user_id INT REFERENCES authorized_client(user_id) ON DELETE CASCADE,
    PRIMARY KEY (product_id, user_id)
);

-- Создание таблицы корзин покупок (shopping_cart)
CREATE TABLE shopping_cart (
    cart_id SERIAL PRIMARY KEY,
    client_id INT REFERENCES authorized_client(user_id)
);

-- Создание таблицы продуктов в корзине (shopping_cart_product)
CREATE TABLE shopping_cart_product (
    cart_id INT REFERENCES shopping_cart(cart_id) ON DELETE CASCADE,
    product_id INT REFERENCES product(product_id) ON DELETE CASCADE,
    count INT NOT NULL,
    is_selected BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (cart_id, product_id)
);

-- Создание таблицы заказов (order)
CREATE TABLE "order" (
    order_id SERIAL PRIMARY KEY,
    status VARCHAR(50) NOT NULL,
    time_stamp TIMESTAMP DEFAULT NOW(),
    cart_id INT REFERENCES shopping_cart(cart_id),
    address_id INT REFERENCES client_address(address_id)
);

-- Создание таблицы сотрудников (employee)
CREATE TABLE employee (
    employee_id SERIAL PRIMARY KEY,
    department VARCHAR(100),
    vacancy VARCHAR(100),
    grade VARCHAR(50),
    salary NUMERIC(10, 2),
    FOREIGN KEY (employee_id) REFERENCES "user"(user_id)
);

-- Создание таблицы шагов заказа (order_step)
CREATE TABLE order_step (
    step_id SERIAL PRIMARY KEY,
    step_name VARCHAR(100) NOT NULL,
    order_id INT REFERENCES "order"(order_id) ON DELETE CASCADE,
    time_stamp_start TIMESTAMP,
    time_stamp_end TIMESTAMP,
    description TEXT,
    status VARCHAR(50),
    responsible_id INT REFERENCES employee(employee_id)
);

-- Создание таблицы складов (warehouse)
CREATE TABLE warehouse (
    warehouse_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address TEXT,
    value NUMERIC(10, 2)
);

-- Создание таблицы продуктов на складе (warehouse_product)
CREATE TABLE warehouse_product (
    product_id INT REFERENCES product(product_id) ON DELETE CASCADE,
    warehouse_id INT REFERENCES warehouse(warehouse_id) ON DELETE CASCADE,
    count INT NOT NULL,
    section VARCHAR(50),
    shelf VARCHAR(50),
    PRIMARY KEY (product_id, warehouse_id)
);

-- Создание таблицы поставщиков (supplier)
CREATE TABLE supplier (
    supplier_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    contact_name VARCHAR(100),
    director_name VARCHAR(100)
);

-- Создание таблицы поставок (supply)
CREATE TABLE supply (
    supply_id SERIAL PRIMARY KEY,
    supplier_id INT REFERENCES supplier(supplier_id) ON DELETE CASCADE,
    time_stamp TIMESTAMP DEFAULT NOW(),
    count INT NOT NULL,
    warehouse_id INT REFERENCES warehouse(warehouse_id),
    responsible_id INT REFERENCES employee(employee_id)
);

-- Создание таблицы продуктов в поставке (supplied_product)
CREATE TABLE supplied_product (
    product_id INT REFERENCES product(product_id) ON DELETE CASCADE,
    supply_id INT REFERENCES supply(supply_id) ON DELETE CASCADE,
    count INT NOT NULL,
    price NUMERIC(10, 2),
    PRIMARY KEY (product_id, supply_id)
);

-- Создание таблицы продуктов в заказе (product_in_order)
CREATE TABLE product_in_order (
    order_id INT REFERENCES "order"(order_id) ON DELETE CASCADE,
    product_id INT REFERENCES product(product_id) ON DELETE CASCADE,
    count INT NOT NULL,
    warehouse_id INT NOT NULL REFERENCES warehouse(warehouse_id),
    PRIMARY KEY (order_id, product_id)
);

# Технологии разработки программного обеспечения
### Корнеев А.Д. ПИМО-01-24

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/ERD.jpg)

# Описание сущностей базы данных онлайн-магазина

| Название сущности         | Описание                                                                                 |
|---------------------------|------------------------------------------------------------------------------------------|
| **User**                  | Пользователь системы.                                  |
| **Role**                  | Роли пользователей, определяющие их права доступа.                                      |
| **Warehouse**             | Склад, где хранятся товары.                                                             |
| **Employee**              | Описание сотрудников компании.                                 |
| **Order**                 | Заказ, сделанный пользователем.                                                          |
| **OrderStep**             | Этапы обработки заказа.                                                                  |
| **Supply**                | Поставка товаров на склад.                                                                |
| **ShoppingCart**          | Корзина покупок пользователя.                                                            |
| **ProductCost**           | Стоимости товара. Позволяет хранить историю цен                                          |
| **Manufacturer**          | Производитель товара.                                                                    |
| **Category**              | Категория товаров.                                                                       |
| **Characteristic**        | Характеристики товара (например, размер, цвет и т.д.).                                   |
| **Supplier**              | Поставщики товара.                           |
| **Product**               | Товар, который может быть куплен в магазине.                                            |
| **Favorites**             | Избранные товары пользователя.                                                           |
| **AuthorizedClient**      | Авторизованный клиент, который может совершать заказы.                  |
| **ClientAddress**         | Адреса клиента для доставки товаров.                                                    |
| **Comment**               | Комментарии пользователей о товарах, процессе покупки и обслуживании.                  |
| **Review**                | Отзывы пользователей о товарах или услугах.                                            |
| **Attachment**            | Вложения, связанные с товарами, заказами или отзывами (например, фотографии).          |

## Сущности для связи M:M

| Название сущности            | Описание                                                                  |
|------------------------------|---------------------------------------------------------------------------|
| **ShoppingCartProduct**      | Множественная связь между корзиной покупок и товарами.                  |
| **ProductCategory**          | Множественная связь между товарами и категориями.                        |
| **ProductCharacteristic**    | Множественная связь между товарами и их характеристиками.               |
| **WarehouseProduct**         | Множественная связь между складами и товарами, хранящимися на них.      |

# Скрипты для создания и заполнения БД

[Скрипт создания БД](Create_database.sql)

[Скрипт заполнения БД](Insert_data.sql)

# Проверка заполнения таблиц

**Таблица Role**
```sql
SELECT * FROM "Role";
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/1.png)
**Таблица User**
```sql
SELECT * FROM "User";
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/2.png)

**Таблица AuthorizedClient**
```sql
SELECT * FROM AuthorizedClient;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/3.png)

**Таблица ClientAddress**
```sql
SELECT * FROM ClientAddress;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/4.png)

**Таблица Manufacturer**
```sql
SELECT * FROM Manufacturer;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/5.png)

**Таблица Product**
```sql
SELECT * FROM Product;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/6.png)

**Таблица Category**
```sql
SELECT * FROM Category;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/7.png)

**Таблица ProductCategory**
```sql
SELECT * FROM ProductCategory;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/8.png)

**Таблица Characteristic**
```sql
SELECT * FROM Characteristic;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/9.png)

**Таблица ProductCharacteristic**
```sql
SELECT * FROM ProductCharacteristic;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/10.png)

**Таблица ProductCost**
```sql
SELECT * FROM ProductCost;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images//11.png)

**Таблица Review**
```sql
SELECT * FROM Review;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/12.png)

**Таблица Comment**
```sql
SELECT * FROM "Comment";
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/13.png)

**Таблица Attachment**
```sql
SELECT * FROM Attachment;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/14.png)

**Таблица Favorites**
```sql
SELECT * FROM Favorites;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/15.png)

**Таблица ShoppingCart**
```sql
SELECT * FROM ShoppingCart;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/16.png)

**Таблица ShoppingCartProduct**
```sql
SELECT * FROM ShoppingCartProduct;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/17.png)

**Таблица Order**
```sql
SELECT * FROM "Order";
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/18.png)

**Таблица OrderStep**
```sql
SELECT * FROM OrderStep;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/19.png)

**Таблица Employee**
```sql
SELECT * FROM Employee;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/20.png)

**Таблица Warehouse**
```sql
SELECT * FROM Warehouse;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/21.png)

**Таблица ProductInOrder**
```sql
SELECT * FROM ProductInOrder;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/22.png)

**Таблица WarehouseProduct**
```sql
SELECT * FROM WarehouseProduct;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/23.png)

**Таблица Supplier**
```sql
SELECT * FROM Supplier;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/24.png)

**Таблица Supply**
```sql
SELECT * FROM Supply;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/25.png)

**Таблица SuppliedProduct**
```sql
SELECT * FROM SuppliedProduct;
```

![alt text](https://github.com/mik0war/SoftDevTech/blob/online_shop_db/images/26.png)


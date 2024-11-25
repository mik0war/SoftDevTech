# Технологии разработки программного обеспечения
### Корнеев А.Д. ПИМО-01-24

# API онлайн-магазина

С использованием библиотеки GIN был создан API онлайн-магазина (CRUD-приложение)

## Обрабатываемые эндпоинты
### Тестирование с помощью postman

1. **Получение данных о всех товарах**

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API/images/postman_get_all.PNG)

2. **Получение данных об одном товаре**

- *Успешный ответ*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API/images/postman_get_id.PNG)

- *Ошибка (отсутствует товар с таким id)*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API/images/postman_get_id_er.PNG)

3. **Создание товара**

- *Успешный ответ*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API/images/postman_post.PNG)

- *Ошибка (пустое тело запроса)*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API/images/postman_post_err.PNG)

4. **Обновление товара**

- *Успешный ответ*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API/images/postman_put.PNG)

- *Успешный ответ (товар не существовал, создан новый)*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API/images/postman_put_new.PNG)

- *Ошибка (пустое тело запроса)*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API/images/postman_put_err.PNG)

5. **Удаление товара**

- *Успешный ответ*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API/images/postman_delete.PNG)

- *Ошибка (отсутствует товар с таким id)*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API/images/postman_delete_err.PNG)

### Тестирование с помощью swagger

Список задокументированных эндпоинтов

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API/images/swagger_all.PNG)

Документация доступна по ссылке, если проект запущен на локальном хосте: http://localhost:8080/swagger/index.html

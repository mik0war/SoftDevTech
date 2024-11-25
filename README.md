# Технологии разработки программного обеспечения
### Корнеев А.Д. ПИМО-01-24

# API онлайн-магазина

С использованием библиотеки GIN был создан API онлайн-магазина (CRUD-приложение). К API был добавлен механизм JWT токенов. 
Доступ к ресурсам теперь осуществляется только с использованием токенов.
В токенах содержится информация об уровне доступа, если он недостаточен - доступ к ресурсу запрещён.

## Обрабатываемые эндпоинты
### Тестирование с помощью postman

1. **Регистрация нового пользователя**

При регистрации всем пользователям выдаётся роль User с минимальным уровнем доступа. В коде программно создана одна учётная запись с данными для входа
```JSON
{
  "username": "admin",
  "password": "root"
}
```

- *Успешный ответ*
  
![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/register.PNG)

- *Ошибка (неверное тело запроса)*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/register_err.PNG)

- *Ошибка (пользователь уже существует)*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/register_err2.PNG)

2. **Логин**

- *Успешный ответ*
  
![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/login.PNG)

- *Ошибка (неверное тело запроса)*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/login_err2.PNG)

- *Ошибка (неверный логин/пароль)*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/login_err.PNG)

3. **Refresh token**

- *Успешный ответ*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/refresh.PNG)

- *Ошибка (просроченный токен)*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/refresh_err.PNG)

- *Ошибка (неправильный токен)*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/refresh_wrong_token.PNG)

## Использование API с токенами

Теперь для доступа к любому ресурсу надо иметь токен. Для эндпоинтов получения товаров (всех / одного по id) необходимо передать в запросе токен с уровнем доступа ```User```. 
Для доступа к эндпоинтам создания, обновления и удаления товаров необходимо передать токен уровня доступа ```Admin```.

- *Попытка обращения к ресурсу без токена*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/get_empty_token.PNG)

- *Попытка обращения к ресурсу с токеном*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/get.PNG)

Для доступа к эндпоинту создания товара необходимо передать токен с ролью Admin.

- *Попытка обращения к ресурсу с токеном обычного пользователя*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/create_by_user.PNG)

- *Попытка обращения к ресурсу с токеном админа*

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/create_by_admin.PNG)

### Тестирование с помощью swagger

Список задокументированных эндпоинтов

![alt-text](https://github.com/mik0war/SoftDevTech/blob/go_API_jwt-tokens/images/swagger_all.PNG)

Документация доступна по ссылке, если проект запущен на локальном хосте: http://localhost:8080/swagger/index.html

package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"online-shop-API/internal/types"
	"strconv"
	"time"
)

// @Summary Создать новый заказ
// @Description Создает заказ с указанным списком продуктов
// @Tags orders
// @Accept json
// @Produce json
// @Param Authorization header string true "Токен"
// @Param order body types.OrderData true "Информация о заказе"
// @Success 201 {object} types.SuccessResponse "Успешное создание заказа"
// @Failure 400 {object} types.ErrorResponse "Неверный запрос (например, пустой список продуктов)"
// @Failure 500 {object} types.ErrorResponse "Ошибка на стороне сервера (например, не удалось создать заказ или добавить продукт)"
// @Router /orders [post]
func (handler *Handler) CreateOrder(c *gin.Context) {
	var orderRequest types.Order

	// Парсим тело запроса
	if err := c.BindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	// Проверяем, что список продуктов не пуст
	if len(orderRequest.Products) == 0 {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "order must contain at least one product"})
		return
	}

	// Вставка заказа в базу данных
	newOrder := types.Order{
		Status:    "IN PROGRESS",
		TimeStamp: time.Now(),
	}

	orderID, err := handler.ProductRepo.CreateOrder(&newOrder) // Метод для создания заказа в репозитории
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "failed to create order"})
		return
	}

	// Добавление продуктов в заказ
	for _, product := range orderRequest.Products {
		productOrder := types.ProductOrder{
			OrderID:   orderID,
			ProductID: product.ProductID,
			Count:     product.Count,
		}

		if err := handler.ProductRepo.AddProductToOrder(&productOrder); err != nil {
			c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "failed to add product to order"})
			return
		}
	}

	c.JSON(http.StatusCreated, types.SuccessResponse{Message: strconv.Itoa(orderID)})
}

// @Summary Получить детали заказа по ID
// @Description Извлекает детали конкретного заказа, включая продукты, связанные с ним.
// @Tags orders
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Токен"
// @Param id path int true "ID заказа"
// @Success 200 {object} types.Order "Детали заказа"
// @Failure 400 {object} types.ErrorResponse "Неверный ID заказа"
// @Failure 404 {object} types.ErrorResponse "Заказ не найден"
// @Failure 500 {object} types.ErrorResponse "Ошибка сервера"
// @Router /orders/{id} [get]
func (handler *Handler) GetOrder(c *gin.Context) {
	// Получаем order_id из параметров маршрута
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid order ID"})
		return
	}

	// Извлекаем данные заказа из репозитория
	order, err := handler.ProductRepo.GetOrderWithProducts(uint(orderID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, types.ErrorResponse{Error: "order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, types.ErrorResponse{Error: "failed to retrieve order"})
		}
		return
	}

	// Возвращаем заказ клиенту
	c.JSON(http.StatusOK, order)
}

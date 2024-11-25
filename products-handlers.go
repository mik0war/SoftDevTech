package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-shop-API/data"
	"online-shop-API/types"
	"online-shop-API/utils"
)

// @Summary Получить все товары
// @Description Возвращает список всех товаров
// @Tags products
// @Produce json
// @Param limit query int false "Максимальное количество товаров"
// @Param offset query int false "Сдвиг"
// @Param category_id query int false "Id категории"
// @Param Authorization header string true "Access token"
// @Success 200 {array} types.Product
// @Failure 401 {object} types.ErrorResponse
// @Router /products [get]
func getProducts(c *gin.Context) {
	params := c.Request.URL.Query()

	limit := utils.GetQueryParam(params,
		"limit",
		data.GetDataLength(),
		func(value int, insteadValue int) bool {
			return value < 0 || value > insteadValue
		},
		utils.GetInt,
	)

	offset := utils.GetQueryParam(params,
		"offset",
		0,
		func(value int, insteadValue int) bool {
			return value < insteadValue
		},
		utils.GetInt)

	categoryId := utils.GetQueryParam(params,
		"category_id",
		"all",
		func(value string, insteadValue string) bool {
			return value == ""
		},
		utils.GetString)

	returnedProducts := data.GetProductsData(limit, offset, categoryId)

	c.JSON(http.StatusOK, returnedProducts)
}

// @Summary Получить товар по ID
// @Description Возвращает информацию о товаре по его ID
// @Tags products
// @Produce json
// @Param id path string true "ID товара"
// @Param Authorization header string true "Access token"
// @Success 200 {object} types.Product
// @Failure 404 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Router /products/{id} [get]
func getProduct(c *gin.Context) {
	id := c.Param("id")
	for _, product := range data.Products {
		if product.ID == id {
			c.JSON(http.StatusOK, product)
			return
		}
	}

	c.JSON(http.StatusNotFound, types.ErrorResponse{Error: "Product not found"})
}

// @Summary Создать новый товар
// @Description Добавляет новый товар в список
// @Tags products
// @Accept json
// @Produce json
// @Param product body types.Product true "Информация о товаре"
// @Param Authorization header string true "Access token"
// @Success 201 {object} types.Product
// @Failure 400 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Router /products [post]
func createProduct(c *gin.Context) {
	var product types.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	c.JSON(http.StatusCreated, data.CreateNewProduct(product))
}

// @Summary Удалить товар
// @Description Удаляет товар по ID
// @Tags products
// @Produce json
// @Param id path string true "ID товара"
// @Param Authorization header string true "Access token"
// @Success 200 {object} types.SuccessResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Router /products/{id} [delete]
func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	_, err := data.DeleteProductData(id)

	if err != nil {
		c.JSON(http.StatusNotFound, types.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.SuccessResponse{Message: "ok"})
}

// @Summary Обновить существующий товар
// @Description Обновляет данные товара по ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "ID товара"
// @Param product body types.Product true "Новые данные товара"
// @Param Authorization header string true "Access token"
// @Success 201 {object} types.SuccessResponse
// @Success 202 {object} types.SuccessResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 401 {object} types.ErrorResponse
// @Router /products/{id} [put]
func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var product types.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	index, err := data.UpdateProductData(id, product)

	if err != nil {
		c.JSON(http.StatusCreated,
			types.ErrorResponse{Error: err.Error() + "new one " + string(rune(index))})
	}
	c.JSON(http.StatusAccepted, types.SuccessResponse{Message: string(rune(index))})
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-shop-API/internal/data"
	"online-shop-API/internal/types"
	"online-shop-API/internal/utils"
	"strconv"
)

type ProductHandler struct {
	productRepo data.ProductRepository
}

// @Summary Получить все товары
// @Description Возвращает список всех товаров
// @Tags products
// @Produce json
// @Success 200 {array} types.Product
// @Router /products [get]
func (handler *ProductHandler) getProducts(c *gin.Context) {
	params := c.Request.URL.Query()

	page := utils.GetQueryParam(params,
		"page",
		1,
		func(value int, insteadValue int) bool {
			return value < 0 || value > insteadValue
		},
		utils.GetInt,
	)

	pageSize := utils.GetQueryParam(params,
		"pageSize",
		100,
		func(value int, insteadValue int) bool {
			return value < insteadValue
		},
		utils.GetInt)

	_ = utils.GetQueryParam(params,
		"category_id",
		"all",
		func(value string, insteadValue string) bool {
			return value == ""
		},
		utils.GetString)

	filters := map[string]interface{}{}

	returnedProducts, _, _ := handler.productRepo.GetProducts(filters, page, pageSize)

	c.JSON(http.StatusOK, returnedProducts)
}

// @Summary Получить товар по ID
// @Description Возвращает информацию о товаре по его ID
// @Tags products
// @Produce json
// @Param id path string true "ID товара"
// @Success 200 {object} types.Product
// @Failure 404 {object} types.ErrorResponse
// @Router /products/{id} [get]
func (handler *ProductHandler) getProduct(c *gin.Context) {
	id := c.Param("id")
	filters := map[string]interface{}{
		"product.product_id": id,
	}
	product, totalSize, err := handler.productRepo.GetProducts(filters, 1, 1)

	if err == nil && totalSize > 0 {
		c.JSON(http.StatusOK, product)
	} else {
		c.JSON(http.StatusNotFound, types.ErrorResponse{Error: "Product not found"})
	}
}

// @Summary Создать новый товар
// @Description Добавляет новый товар в список
// @Tags products
// @Accept json
// @Produce json
// @Param product body types.Product true "Информация о товаре"
// @Success 201 {object} types.Product
// @Failure 400 {object} types.ErrorResponse
// @Router /products [post]
func (handler *ProductHandler) createProduct(c *gin.Context) {
	var product types.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	answer, err := handler.productRepo.AddProduct(&product)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, types.ErrorResponse{Error: err.Error()})
	}
	c.JSON(http.StatusCreated, answer)
}

// @Summary Удалить товар
// @Description Удаляет товар по ID
// @Tags products
// @Produce json
// @Param id path string true "ID товара"
// @Success 200 {object} types.Product
// @Failure 404 {object} types.ErrorResponse
// @Router /products/{id} [delete]
func (handler *ProductHandler) deleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := handler.productRepo.DeleteProduct(uint(id))

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
// @Success 201 {object} types.SuccessResponse
// @Success 202 {object} types.SuccessResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /products/{id} [put]
func (handler *ProductHandler) updateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product types.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	newProduct, err := handler.productRepo.UpdateProduct(uint(id), &product)

	if err != nil {
		c.JSON(http.StatusCreated,
			types.ErrorResponse{Error: err.Error() + "new one"})
	}
	c.JSON(http.StatusAccepted, newProduct)
}

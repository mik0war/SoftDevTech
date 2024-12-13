package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-shop-API/internal/types"
	"online-shop-API/internal/utils"
	"strconv"
)

// @Summary Получить все товары
// @Description Возвращает список всех товаров
// @Tags products
// @Produce json
// @Param Authorization header string true "Токен"
// @Success 200 {array} types.Product
// @Router /products [get]
func (handler *Handler) GetProducts(c *gin.Context) {
	params := c.Request.URL.Query()

	page := utils.GetQueryParam(params,
		"page",
		1,
		func(value int, insteadValue int) bool {
			return value < insteadValue
		},
		utils.GetInt,
	)

	pageSize := utils.GetQueryParam(params,
		"pageSize",
		100,
		func(value int, insteadValue int) bool {
			return value < 0 || value > insteadValue
		},
		utils.GetInt)

	category := utils.GetQueryParam(params,
		"category",
		"ALL",
		func(value string, insteadValue string) bool {
			return value == ""
		},
		utils.GetString)

	filters := map[string]interface{}{
		"category.name": category,
	}

	returnedProducts, _, err := handler.ProductRepo.GetProducts(filters, page, pageSize)
	if err == nil {
		c.JSON(http.StatusOK, returnedProducts)
	} else {
		c.JSON(http.StatusInternalServerError, err)
	}
}

// @Summary Получить товар по ID
// @Description Возвращает информацию о товаре по его ID
// @Tags products
// @Produce json
// @Param Authorization header string true "Токен"
// @Param id path string true "ID товара"
// @Success 200 {object} types.Product
// @Failure 404 {object} types.ErrorResponse
// @Router /products/{id} [get]
func (handler *Handler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	filters := map[string]interface{}{
		"product.product_id": id,
	}
	product, totalSize, err := handler.ProductRepo.GetProducts(filters, 1, 1)

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
// @Param Authorization header string true "Токен"
// @Produce json
// @Param product body types.ProductData true "Информация о товаре"
// @Success 201 {object} types.Product
// @Failure 400 {object} types.ErrorResponse
// @Router /products [post]
func (handler *Handler) CreateProduct(c *gin.Context) {
	var product types.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	answer, err := handler.ProductRepo.AddProduct(&product)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, types.ErrorResponse{Error: err.Error()})
	} else {
		c.JSON(http.StatusCreated, answer)
	}
}

// @Summary Удалить товар
// @Description Удаляет товар по ID
// @Tags products
// @Produce json
// @Param Authorization header string true "Токен"
// @Param id path string true "ID товара"
// @Success 200 {object} types.ProductData
// @Failure 404 {object} types.ErrorResponse
// @Router /products/{id} [delete]
func (handler *Handler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := handler.ProductRepo.DeleteProduct(uint(id))

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
// @Param Authorization header string true "Токен"
// @Param id path string true "ID товара"
// @Param product body types.Product true "Новые данные товара"
// @Success 201 {object} types.SuccessResponse
// @Success 202 {object} types.SuccessResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /products/{id} [put]
func (handler *Handler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product types.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	newProduct, err := handler.ProductRepo.UpdateProduct(uint(id), &product)

	if err != nil {
		c.JSON(http.StatusCreated,
			types.ErrorResponse{Error: err.Error() + "new one"})
	}
	c.JSON(http.StatusAccepted, newProduct)
}

// @Summary Добавить стоимость товара
// @Description Добавляет новую стоимость для существующего товара
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "ID товара"
// @Param Authorization header string true "Токен"
// @Param cost body types.ProductCost true "Информация о стоимости"
// @Success 201 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 503 {object} types.ErrorResponse
// @Router /products/{id}/cost [post]
func (handler *Handler) AddCost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var cost types.ProductCost
	if err := c.BindJSON(&cost); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	err := handler.ProductRepo.AddCostToProduct(id, &cost)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, types.ErrorResponse{Error: err.Error()})
	} else {
		c.JSON(http.StatusCreated, types.SuccessResponse{Message: "ok"})
	}
}

// @Summary Добавить категорию товара
// @Description Добавляет новую категорию для существующего товара
// @Tags products
// @Produce json
// @Param id path string true "ID товара"
// @Param Authorization header string true "Токен"
// @Param category query string true "Название категории"
// @Success 201 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse
// @Failure 503 {object} types.ErrorResponse
// @Router /products/{id}/category [post]
func (handler *Handler) AddCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	params := c.Request.URL.Query()

	categoryName := utils.GetQueryParam(params,
		"category",
		"ALL",
		func(value string, insteadValue string) bool {
			return value == ""
		},
		utils.GetString)

	err := handler.ProductRepo.AddCategoryToProduct(id, categoryName)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, types.ErrorResponse{Error: err.Error()})
	} else {
		c.JSON(http.StatusCreated, types.SuccessResponse{Message: "ok"})
	}
}

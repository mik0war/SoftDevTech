package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "online-shop-API/docs"
	"online-shop-API/types"
	"strconv"
)

// @Summary Получить все товары
// @Description Возвращает список всех товаров
// @Tags products
// @Produce json
// @Param limit query int false "Максимальное количество товаров"
// @Param offset query int false "Сдвиг"
// @Success 200 {array} types.Product
// @Router /products [get]
func getProducts(c *gin.Context) {
	params := c.Request.URL.Query()

	limit := getQueryParam(params,
		"limit",
		getDataLength(),
		func(value int, insteadValue int) bool {
			return value < 0 || value > insteadValue
		},
		getInt,
	)

	offset := getQueryParam(params,
		"offset",
		0,
		func(value int, insteadValue int) bool {
			return value < insteadValue
		},
		getInt)

	categoryId := getQueryParam(params,
		"category_id",
		"all",
		func(value string, insteadValue string) bool {
			return value == ""
		},
		getString)

	returnedProducts := getProductsData(limit, offset, categoryId)

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
func getProduct(c *gin.Context) {
	id := c.Param("id")
	for _, product := range products {
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
// @Success 201 {object} types.Product
// @Failure 400 {object} types.ErrorResponse
// @Router /products [post]
func createProduct(c *gin.Context) {
	var product types.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	c.JSON(http.StatusCreated, createNewProduct(product))
}

// @Summary Удалить товар
// @Description Удаляет товар по ID
// @Tags products
// @Produce json
// @Param id path string true "ID товара"
// @Success 200 {object} types.SuccessResponse
// @Failure 404 {object} types.ErrorResponse
// @Router /products/{id} [delete]
func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	_, err := deleteProductData(id)

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
func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var product types.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{Error: "invalid request"})
		return
	}

	index, err := updateProductData(id, product)

	if err != nil {
		c.JSON(http.StatusCreated,
			types.SuccessResponse{Message: err.Error() + ", new one " + strconv.Itoa(index)})
	} else {
		c.JSON(http.StatusAccepted, types.SuccessResponse{Message: strconv.Itoa(index)})
	}
}

// @title           Online shop API Swagger
// @version         1.0
// @description     This is a sample online-shop server

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      127.0.0.1:8080

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Создаем новый роутер Gin
	router := gin.Default()

	router.Use(cors.Default())

	// Определяем маршруты для продуктов
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)
	router.POST("/products/", createProduct)
	router.DELETE("/products/:id", deleteProduct)
	router.PUT("/products/:id", updateProduct)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запускаем сервер на порту 8080
	err := router.Run(":8080")
	if err != nil {
		return
	}

}

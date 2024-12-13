package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "online-shop-API/docs"
	"online-shop-API/internal/data"
)

// @title           Online shop API Swagger
// @version         1.0
// @description     This is a sample online-shop server
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Создаем новый роутер Gin
	data.InitDB()
	repository := data.NewProductRepository(data.Db)
	handler := Handler{*repository}

	router := gin.Default()

	router.POST("/auth/login", handler.login)
	router.POST("/auth/refresh", refresh)
	router.POST("/auth/register", handler.registration)

	adminGroup := router.Group("/")
	adminRole, err := repository.GetRole("Admin")
	if err != nil {
		return
	}
	adminGroup.Use(authMiddleware(*adminRole))
	{
		adminGroup.POST("/products/", handler.createProduct)
		adminGroup.POST("/products/:id/cost", handler.addCost)
		adminGroup.POST("/products/:id/category", handler.addCategory)
		adminGroup.DELETE("/products/:id", handler.deleteProduct)
		adminGroup.PUT("/products/:id", handler.updateProduct)
	}

	userGroup := router.Group("/")
	userRole, err := repository.GetRole("User")
	if err != nil {
		return
	}
	userGroup.Use(authMiddleware(*userRole))
	{
		userGroup.GET("/products", handler.getProducts)
		userGroup.GET("/products/:id", handler.getProduct)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запускаем сервер на порту 8080
	err = router.Run(":8080")
	if err != nil {
		return
	}

}

type Handler struct {
	productRepo data.Repository
}

package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "online-shop-API/docs"
	"online-shop-API/types"
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
	router := gin.Default()

	router.POST("/login", login)
	router.POST("/refresh", refresh)
	router.POST("/registration", registration)

	adminGroup := router.Group("/")
	adminRole := types.Role{Name: "admin"}
	adminGroup.Use(authMiddleware(adminRole))
	{
		adminGroup.POST("/products/", createProduct)
		adminGroup.DELETE("/products/:id", deleteProduct)
		adminGroup.PUT("/products/:id", updateProduct)
	}

	userGroup := router.Group("/")
	userRole := types.Role{Name: "user"}
	userGroup.Use(authMiddleware(userRole))
	{
		userGroup.GET("/products", getProducts)
		userGroup.GET("/products/:id", getProduct)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запускаем сервер на порту 8080
	err := router.Run(":8080")
	if err != nil {
		return
	}

}

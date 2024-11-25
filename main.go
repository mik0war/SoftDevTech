package main

import (
	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:8080", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Accept", "Origin"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
	}))

	router.POST("/login", login)
	router.POST("/refresh", refresh)
	router.POST("/register", registration)

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

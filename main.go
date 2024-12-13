package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "online-shop-API/docs"
	"online-shop-API/internal/data"
	"online-shop-API/internal/handlers"
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
	repository := data.NewRepository(data.Db)
	handler := handlers.Handler{ProductRepo: *repository}

	router := gin.Default()

	router.POST("/auth/login", handler.Login)
	router.POST("/auth/refresh", handlers.Refresh)
	router.POST("/auth/register", handler.Registration)

	adminGroup := router.Group("/")
	adminRole, err := repository.GetRole("Admin")
	if err != nil {
		return
	}

	adminGroup.Use(handlers.AuthMiddleware(*adminRole))
	{
		adminGroup.GET("/orders/:id", handler.GetOrder)

		adminGroup.POST("/products/", handler.CreateProduct)
		adminGroup.POST("/products/:id/cost", handler.AddCost)
		adminGroup.POST("/products/:id/category", handler.AddCategory)
		adminGroup.DELETE("/products/:id", handler.DeleteProduct)
		adminGroup.PUT("/products/:id", handler.UpdateProduct)
	}

	userGroup := router.Group("/")
	userRole, err := repository.GetRole("User")
	if err != nil {
		return
	}
	userGroup.Use(handlers.AuthMiddleware(*userRole))
	{
		userGroup.GET("/products", handler.GetProducts)
		userGroup.GET("/products/:id", handler.GetProduct)
		userGroup.POST("/orders", handler.CreateOrder)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запускаем сервер на порту 8080
	err = router.Run(":8080")
	if err != nil {
		return
	}

}

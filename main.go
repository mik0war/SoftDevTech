package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getProducts() {

}

func main() {
	// Создаем новый роутер Gin
	router := gin.Default()

	// Определяем маршруты
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("Start server on ", router.BasePath())
	// Запускаем сервер на порту 8080
	router.Run(":8080")

}

package main

import (
	"github.com/gin-gonic/gin"

	"primaryserver/handler"
)

func main() {
	router := gin.Default()
	router.GET("/health", handler.HealthCheck)
	router.GET("/productNames/:wordToSearch", handler.GetSubStringMatch)
	router.GET("/productResults/:productName", handler.GetProductResults)
	router.Run("localhost:8080")
}

package main

import (
	"primaryserver/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/health", handler.HealthCheck)
	router.GET("/search/:word", handler.GetSubStringMatch)
	router.GET("/products/:name", handler.GetProductResults)
	router.Run("localhost:8080")
}

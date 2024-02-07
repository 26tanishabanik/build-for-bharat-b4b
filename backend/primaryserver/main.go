package main

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var productMap map[string]ProductDetails = map[string]ProductDetails{}
var productMapMutex sync.RWMutex
var wg sync.WaitGroup

func main() {
	router := gin.Default()
	router.GET("/health", HealthCheck)
	router.GET("/productNames/:wordToSearch", GetSubStringMatch)
	router.GET("/productResults/:productName", GetProductResults)
	router.Run("localhost:8080")
}

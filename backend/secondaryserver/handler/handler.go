package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"secondaryserver/utils"
)

// handler for validating whether API is running or not
func HealthCheck(c *gin.Context) {
	now := time.Now()
	health := make(map[string]string)
	health["now"] = now.Format(time.ANSIC)
	c.JSON(http.StatusOK, health)
}

func HandleSearch(c *gin.Context) {
	productName := c.Param("productName")
	results, err := utils.SearchProducts(productName)
	if err != nil {
		fmt.Println("error in getting the product list from elasticsearch: ", err)
		return
	}
	c.JSON(http.StatusOK, results)
}
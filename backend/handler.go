package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// handler for validating whether API is running or not
func HealthCheck(c *gin.Context) {
	now := time.Now()
	health := make(map[string]string)
	health["now"] = now.Format(time.ANSIC)
	c.JSON(http.StatusOK, health)
}

func GetSubStringMatch(c *gin.Context) {
	wordToSearch := c.Param("wordToSearch")
	payload := Payload{}
	for _, name := range ListOfProductNames {
		if strings.Contains(strings.ToLower(name), strings.ToLower(wordToSearch)) {
			payload.MatchedWords = append(payload.MatchedWords, name)
		}
	}
	if len(payload.MatchedWords) > 0 {
		c.JSON(http.StatusOK, payload)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "word not found"})
	}
}

func GetProductResults(c *gin.Context) {
	productName := c.Param("productName")
	UUID := c.PostForm("uuid")

	if UUID == "" {
		newUUID := uuid.New().String()
		c.JSON(http.StatusOK, gin.H{
			"uuid": newUUID,
		})
		productMap[newUUID] = ProductDetails{
			Products:          []map[string]interface{}{},
			IsProcessComplete: false,
			Errors:            []error{},
		}
		go searchProductsInSecondaryServer(productName, newUUID)
	} else {
		var productDetails ProductDetails
		productMapMutex.RLock()
		productDetails = productMap[UUID]
		productMapMutex.RUnlock()

		if productDetails.IsProcessComplete {
			c.JSON(http.StatusOK, gin.H{UUID: productDetails.Products})
			delete(productMap, UUID)
		} else {
			c.JSON(http.StatusOK, gin.H{UUID: []map[string]interface{}{}})
		}

	}
}
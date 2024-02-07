package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"primaryserver/model"
	"primaryserver/utils"
)

// handler for validating whether API is running or not
func HealthCheck(c *gin.Context) {
	now := time.Now()
	health := make(map[string]string)
	health["now"] = now.Format(time.ANSIC)
	c.JSON(http.StatusOK, health)
}

func GetSubStringMatch(c *gin.Context) {
	wordToSearch := c.Param("word")
	payload := model.Payload{}
	for _, name := range model.ListOfProductNames {
		if strings.Contains(strings.ToLower(name), strings.ToLower(wordToSearch)) {
			payload.MatchedWords = append(payload.MatchedWords, name)
		}
	}
	c.JSON(http.StatusOK, payload)
}

func GetProductResults(c *gin.Context) {
	productName := c.Param("name")
	clientUUID := c.Query("uuid")

	if clientUUID == "" {
		newUUID := uuid.New().String()
		c.JSON(http.StatusOK, gin.H{
			"uuid": newUUID,
		})
		model.ProductMapMutex.Lock()
		model.ProductMap[newUUID] = model.ProductDetails{
			Products:          []model.Product{},
			IsProcessComplete: false,
			Errors:            []error{},
		}
		model.ProductMapMutex.Unlock()
		go utils.SearchProductsInSecondaryServer(productName, newUUID)
	} else {
		var productDetails model.ProductDetails
		var response model.ResponseToClient
		model.ProductMapMutex.RLock()
		productDetails = model.ProductMap[clientUUID]
		model.ProductMapMutex.RUnlock()

		if productDetails.IsProcessComplete {
			response.Products = productDetails.Products
			response.UUID = clientUUID
			fmt.Printf("Received products for uuid %s: %v\n", clientUUID, productDetails.Products)
			c.JSON(http.StatusOK, response)
			delete(model.ProductMap, clientUUID)
		} else {
			response.Products = []model.Product{}
			response.UUID = clientUUID
			c.JSON(http.StatusOK, response)
		}
	}
}

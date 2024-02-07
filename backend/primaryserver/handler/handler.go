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
	wordToSearch := c.Param("wordToSearch")
	payload := model.Payload{}
	for _, name := range model.ListOfProductNames {
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
	// UUID := c.PostForm("uuid")
	var request model.RequestFromClient
	if err := c.BindJSON(&request); err != nil {
		fmt.Println("error in binding request body from client: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if request.UUID == "" {
		newUUID := uuid.New().String()
		c.JSON(http.StatusOK, gin.H{
			"uuid": newUUID,
		})
		model.ProductMap[newUUID] = model.ProductDetails{
			Products:          []model.Product{},
			IsProcessComplete: false,
			Errors:            []error{},
		}
		go utils.SearchProductsInSecondaryServer(productName, newUUID)
	} else {
		var productDetails model.ProductDetails
		var response model.ResponseToClient
		model.ProductMapMutex.RLock()
		productDetails = model.ProductMap[request.UUID]
		model.ProductMapMutex.RUnlock()

		if productDetails.IsProcessComplete {
			response.Products = productDetails.Products
			response.UUID = request.UUID
			fmt.Printf("Received products for uuid %s: %v\n", request.UUID, productDetails.Products)
			c.JSON(http.StatusOK, response)
			delete(model.ProductMap, request.UUID)
		} else {
			response.Products = []model.Product{}
			response.UUID = request.UUID
			c.JSON(http.StatusOK, response)
		}
	}
}
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gin-gonic/gin"
)

var index_name = "dmart2"

func decodeResponse(resp *esapi.Response, target interface{}) error {
	if err := decode(resp.Body, target); err != nil {
		return fmt.Errorf("error decoding response body: %w", err)
	}
	return nil
}

func decode(reader io.Reader, target interface{}) error {
	if err := json.NewDecoder(reader).Decode(target); err != nil {
		return fmt.Errorf("error decoding JSON: %w", err)
	}
	return nil
}

func searchProducts(productName string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf(`
		{
			"query": { 
				"match": {
					"productName": "%s"
			  	}
			}
		}
	`, productName)
	resp, err := es.Search(
		es.Search.WithIndex(index_name),
		es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		fmt.Println("error in searching a index: ", err)
		return nil, err
	}
	var result map[string]interface{}
	if err := decodeResponse(resp, &result); err != nil {
		fmt.Printf("error decoding response: %s\n", err)
		return nil, err
	}
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	var results []map[string]interface{}
	for _, hit := range hits {
		results = append(results, hit.(map[string]interface{})["_source"].(map[string]interface{}))
	}
	return results, nil
}

func HandleSearch(c *gin.Context) {
	productName := c.Param("productName")
	results, err := searchProducts(productName)
	if err != nil {
		fmt.Println("error in getting the product list from elasticsearch: ", err)
		return
	}
	c.JSON(http.StatusOK, results)
}

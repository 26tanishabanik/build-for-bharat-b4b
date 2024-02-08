package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"

	"secondaryserver/model"
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

func SearchProducts(productName string) ([]model.Product, error) {
	var results []model.Product
	var response map[string]interface{}
	query := fmt.Sprintf(`
		{
			"query": { 
				"match": {
					"productName": "%s"
			  	}
			}
		}
	`, productName)
	resp, err := model.ES.Search(
		model.ES.Search.WithIndex(index_name),
		model.ES.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		fmt.Println("error in searching a index: ", err)
		return nil, err
	}
	if err := decodeResponse(resp, &response); err != nil {
		fmt.Printf("error decoding response: %s\n", err)
		return nil, err
	}
	fmt.Println("Results: ", results)
	hits := response["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		result := model.Product{}
		// productInterface := hit.(map[string]interface{})["_source"].(map[string]interface{})
		productInterface := hit.(map[string]interface{})
		log.Println("_id: ", productInterface["_id"])
		updatedData, err := json.Marshal(productInterface)
		if err != nil {
			fmt.Println("error in marshalling raw product: ", err)
			return nil, err
		}
		err = json.Unmarshal(updatedData, &result)
		if err != nil {
			fmt.Println("error in unmarshalling raw product to Product struct: ", err)
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

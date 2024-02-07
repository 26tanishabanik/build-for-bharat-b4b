package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"primaryserver/pkg/server"
	"primaryserver/model"
)

const (
	SecondaryServerSearchEndpoint = "%s/search/%s"
)

func decodeResponse(body io.Reader, target interface{}) error {
	if err := decode(body, target); err != nil {
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

func SearchProductsInSecondaryServer(productName, uuid string) {
	secondaryServers := server.GetSecondaryServers()

	for _, server := range secondaryServers {
		serverURL := fmt.Sprintf(SecondaryServerSearchEndpoint, server, productName)
		model.Wg.Add(1)
		go func(productName, serverURL, uuid string) {
			defer model.Wg.Done()
			products, err := getSecondaryServerResponse(productName, serverURL)
			if err != nil {
				model.ProductMapMutex.Lock()
				productDetails := model.ProductMap[uuid]
				productDetails.Errors = append(productDetails.Errors, err)
				model.ProductMap[uuid] = productDetails
				model.ProductMapMutex.Unlock()

				return
			}
			model.ProductMapMutex.Lock()
			productDetails := model.ProductMap[uuid]
			productDetails.Products = append(productDetails.Products, products...)
			model.ProductMap[uuid] = productDetails
			model.ProductMapMutex.Unlock()
		}(productName, serverURL, uuid)
	}

	model.Wg.Wait()

	model.ProductMapMutex.Lock()
	productDetails := model.ProductMap[uuid]
	productDetails.IsProcessComplete = true
	model.ProductMap[uuid] = productDetails
	model.ProductMapMutex.Unlock()
}

func getSecondaryServerResponse(productName, url string) ([]model.Product, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error in sending get request to secondary server: ", err)
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Printf("expected status code to be 200 but got %v\n", response.StatusCode)
		return nil, fmt.Errorf("status code is not 200")
	}
	var products []model.Product
	err = decodeResponse(response.Body, &products)
	if err != nil {
		fmt.Printf("error decoding response: %s\n", err)
		return nil, err
	}
	return products, nil
}
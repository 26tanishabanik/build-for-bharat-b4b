package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"primaryserver/pkg/server"
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

func searchProductsInSecondaryServer(productName, uuid string) {
	secondaryServers := server.GetSecondaryServers()

	for _, server := range secondaryServers {
		serverURL := fmt.Sprintf(SecondaryServerSearchEndpoint, server, productName)
		wg.Add(1)
		go func(productName, serverURL, uuid string) {
			defer wg.Done()
			products, err := getSecondaryServerResponse(productName, serverURL)
			if err != nil {
				productMapMutex.Lock()
				productDetails := productMap[uuid]
				productDetails.Errors = append(productDetails.Errors, err)
				productMap[uuid] = productDetails
				productMapMutex.Unlock()

				return
			}
			productMapMutex.Lock()
			productDetails := productMap[uuid]
			productDetails.Products = append(productDetails.Products, products...)
			productMap[uuid] = productDetails
			productMapMutex.Unlock()
		}(productName, serverURL, uuid)
	}

	wg.Wait()

	productMapMutex.Lock()
	productDetails := productMap[uuid]
	productDetails.IsProcessComplete = true
	productMap[uuid] = productDetails
	productMapMutex.Unlock()
}

func getSecondaryServerResponse(productName, url string) ([]map[string]interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("error in sending get request to secondary server: ", err)
		return nil, err
	}
	fmt.Println("Hi, 75")
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Printf("expected status code to be 200 but got %v\n", response.StatusCode)
		return nil, fmt.Errorf("status code is not 200")
	}
	fmt.Println("Hi, 83")
	var products []map[string]interface{}
	err = decodeResponse(response.Body, &products)
	fmt.Println("Hi, 86")
	if err != nil {
		fmt.Printf("error decoding response: %s\n", err)
		return nil, err
	}
	return products, nil
}
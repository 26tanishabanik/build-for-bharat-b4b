package model

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8"
)

type Product struct {
	Price       string `json:"price"`
	ProductID   string `json:"productID"`
	Quantity    string `json:"quantity"`
	ProductName string `json:"productName"`
	Seller      string `json:"seller"`
	Image       string `json:"image"`
}

var ES *elasticsearch.Client

func (p *Product) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	p.Price = raw["_source"].(map[string]interface{})["price"].(string)
	p.ProductID = raw["_id"].(string)
	p.Quantity = raw["_source"].(map[string]interface{})["quantity"].(string)
	p.ProductName = raw["_source"].(map[string]interface{})["productName"].(string)
	p.Seller = raw["_source"].(map[string]interface{})["seller"].(string)
	p.Image = raw["_source"].(map[string]interface{})["image"].(string)
	return nil
}

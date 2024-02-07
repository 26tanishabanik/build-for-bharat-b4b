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
	p.Price = raw["price"].(string)
	p.ProductID = raw["productID"].(string)
	p.Quantity = raw["quantity"].(string)
	p.ProductName = raw["productName"].(string)
	p.Seller = raw["seller"].(string)
	p.Image = raw["image"].(string)
	return nil
}
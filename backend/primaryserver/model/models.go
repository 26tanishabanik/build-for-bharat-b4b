package model

import (
	"sync"
)

var ListOfProductNames = []string{
	"Apple",
	"Apricot",
	"Green Apple",
	"Red Apple",
	"Fuji Apple",
	"HoneyCrisp Apple",
	"Orange",
	"Kinnow",
	"Green Grapes",
}

type Payload struct {
	MatchedWords []string `json:"matchingWords"`
}

type ProductDetails struct {
	Products          []Product
	IsProcessComplete bool
	Errors            []error
}

type ResponseToClient struct {
	UUID     string    `json:"uuid"`
	Products []Product `json:"products"`
}

type RequestFromClient struct {
	UUID     string    `json:"uuid"`
}

type Product struct {
	Price       string `json:"price"`
	ProductID   string `json:"productID"`
	Quantity    string `json:"quantity"`
	ProductName string `json:"productName"`
	Seller      string `json:"seller"`
	Image       string `json:"image"`
}

var ProductMap map[string]ProductDetails = map[string]ProductDetails{}
var ProductMapMutex sync.RWMutex
var Wg sync.WaitGroup

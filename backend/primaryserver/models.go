package main

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
	Products          []map[string]interface{}
	IsProcessComplete bool
	Errors            []error
}
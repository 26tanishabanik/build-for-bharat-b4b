package model

import (
	"sync"
)

var ListOfProductNames = []string{
	"Apple", "Apricot", "Avocado", "Banana", "Bilberry", "Blackberry", "Blueberry",
	"Boysenberry", "Cactus", "Cantaloupe", "Carambola", "Cherry", "Clementine", "Coconut",
	"Cranberry", "Currant", "Date", "Dragonfruit", "Durian", "Elderberry", "Feijoa", "Fig",
	"Grape", "Grapefruit", "Guava", "Huckleberry", "Jackfruit", "Jujube", "Kiwi", "Kumquat",
	"Lemon", "Lime", "Lingonberry", "Loganberry", "Lychee", "Mango", "Mulberry", "Nectarine",
	"Orange", "Papaya", "Passionfruit", "Peach", "Pear", "Persimmon", "Plum", "Pomegranate",
	"Raspberry", "Redcurrant", "Soursop", "Starfruit", "Strawberry", "Tamarillo", "Tangerine",
	"Uvaia", "Ugli", "Watermelon", "Ackee", "Acerola", "Açaí", "Ambarella", "Babaco", "Barbados",
	"Bilimbi", "Breadfruit", "Caja", "Calamondin", "Canistel", "Capulin", "Carob", "Casaba",
	"Chaunsa", "Chayote", "Cherimoya", "Chico", "Cloudberry", "Cocoplum", "Cupuaçu", "Damson",
	"Duku", "Emu Apple", "Endive", "Feijoa", "Fuyu", "Gac", "Goumi", "Grumichama", "Hala", "Imbe",
	"Jabuticaba", "Jicama", "Kaffir", "Kepel", "Langsat", "Lemonade", "Longan", "Mamey", "Mamoncillo",
	"Mangaba", "Mange", "Mankai", "Maracuja", "Medlar", "Miracle", "Monstera", "Morinda", "Nashi",
	"Nectar", "Olive", "Papino", "Pepino", "Peruvian", "Pili", "Pitanga", "Pitaya", "Plumcot",
	"Pomelo", "Quince", "Raisin", "Rambutan", "Red", "Salak", "Santol", "Sapote", "Sarsi",
	"Saskatoon", "Seville", "Sloe", "Sorrel", "Soursop", "Sweet", "Tamarind", "Tangelo",
	"Tangerillo", "Thorn", "Toona", "Turmeric", "Ugni", "Vanilla", "Velvet", "Wampee", "Wood",
	"Yangmei", "Yantok", "Yumberry", "Zalacca", "Zapote", "Zinfandel", "Ziziphus", "Zucchero",
	"Zwetschge", "Akebia", "Zabola", "Ugni", "Waru", "Zokor", "Tkeika", "Illawarra", "Bombax",
	"Zeylanica", "Cudrania", "Ant", "Tucuman", "Zambola", "Plinia", "Agarita", "Tutuma",
	"Mirabelle", "Beli", "Tucuman", "Aratiles", "Goldenberry", "Kumis", "Aji", "Tampoi",
	"Falsa", "Tepin", "Lamantan", "Yunnan", "Rollinia", "Jujuba", "Marolo", "Carunda",
	"Nance", "Mamoncillo", "Jostaberry", "Tungmoui", "Ananas", "Duku", "Jabuticaba",
	"Tamarillo", "Bilimbi", "Abiu", "Achiote", "Uvalha", "Yacón", "Uvilla", "Willughbeia",
	"Yunnan", "Umeboshi", "Bilva",
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
	IsResult bool      `json:"isResult"`
	Products []Product `json:"products"`
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

package main

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"

	"secondaryserver/handler"
	"secondaryserver/model"
)

var cloud_id = "My_deployment:YXNpYS1zb3V0aDEuZ2NwLmVsYXN0aWMtY2xvdWQuY29tJGZhNzM5ZmQ4YzI1NzQ3MDA4NWQ3ZTg5MGQyMDdkZjJmJGVlYTFmOTJkMzJjODRiMjQ4YjJiNTNjMWU2ZDBlOGJj"
var api_key = "MU5XQ1hvMEI5NWN3MVdkYzBuSTE6VExwUTRDdUpUMTZ2REg0c1NWRjdIQQ=="

// https://fa739fd8c257470085d7e890d207df2f.asia-south1.gcp.elastic-cloud.com:443

func initElasticsearch() {
	var err error
	cfg := elasticsearch.Config{
        CloudID: cloud_id,
        APIKey: api_key,
	}
	model.ES, err = elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("error in creating a new client: ", err)
		return
	}
	fmt.Println(model.ES.Info())
}

func main() {
	initElasticsearch()
	router := gin.Default()
	router.GET("/health", handler.HealthCheck)
	router.GET("/search/:productName", handler.HandleSearch)
	router.Run("localhost:9090")
}


package main

import (
	"context"

	queryengine "canto-api/query"
	redisclient "canto-api/redisclient"
	"canto-api/serve"

	"github.com/gin-gonic/gin"
)

func main() {
	redisclient.NewClient()
	var ctx = context.Background()
	go queryengine.Run(ctx) // run query engine in background

	router := gin.Default()
	router.GET("/get", serve.GetSmartContractData)

	router.Run()
}

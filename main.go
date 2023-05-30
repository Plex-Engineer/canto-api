package main

import (
	"context"

	"canto-api/config"
	queryengine "canto-api/query"
	"canto-api/serve"

	"github.com/gin-gonic/gin"
)

func main() {
	config.NewConfig()
	ctx := context.Background()
	go queryengine.Run(ctx) // run query engine in background

	router := gin.Default()
	router.GET("/get", serve.GetSmartContractData)

	router.Run()
}

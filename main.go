package main

import (
	"context"

	queryengine "canto-api/query"
	"canto-api/serve"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func main() {
	go queryengine.Run(ctx) // run query engine in background

	router := gin.Default()
	router.GET("/get", serve.GetSmartContractData)

	router.Run()
}

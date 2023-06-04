package main

import (
	"context"

	"canto-api/config"
	queryengine "canto-api/query"
	requestengine "canto-api/serve"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.NewConfig()
	ctx := context.Background()
	go queryengine.Run(ctx) // run query engine in background

	// router := gin.Default()
	// router.GET("/get", requestengine.GetSmartContractData)
	// router.Run()

	app := fiber.New()

	app.Get("/", requestengine.GetSmartContractDataFiber)

	app.Listen(":3000")
}

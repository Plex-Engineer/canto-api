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

	server := "fiber"

	if server == "fiber" {
		app := fiber.New()
		app.Get("/", requestengine.GetSmartContractDataFiber)
		app.Listen(":6009")
	}
}

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
	go queryengine.RunNative(ctx)

	server := "fiber"

	if server == "fiber" {
		app := fiber.New()
		app.Get("/", requestengine.GetSmartContractDataFiber)
		app.Get("/apr", requestengine.QueryStakingAPR)
		app.Get("/validators", requestengine.QueryValidators)
		app.Get("/validators/:address", requestengine.QueryValidatorByAddress)
		app.Get("/csrs", requestengine.QueryCSRs)
		app.Get("/csrs/:id", requestengine.QueryCSRByID)
		app.Get("/gov/proposals", requestengine.QueryProposals)
		app.Get("/gov/proposals/:id", requestengine.QueryProposalByID)
		app.Listen(":3000")
	}
}

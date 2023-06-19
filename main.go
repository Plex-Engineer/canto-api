package main

import (
	"context"

	"canto-api/config"
	cqe "canto-api/query/contracts"
	nqe "canto-api/query/native"
	re "canto-api/requests"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.NewConfig()
	ctx := context.Background()
	go cqe.Run(ctx) // run contract query engine
	go nqe.Run(ctx) // run native query engine

	server := "fiber"

	if server == "fiber" {
		app := fiber.New()
		app.Get("/", re.GetSmartContractDataFiber)
		app.Get("/apr", re.QueryStakingAPR)
		app.Get("/validators", re.QueryValidators)
		app.Get("/validators/:address", re.QueryValidatorByAddress)
		app.Get("/csrs", re.QueryCSRs)
		app.Get("/csrs/:id", re.QueryCSRByID)
		app.Get("/gov/proposals", re.QueryProposals)
		app.Get("/gov/proposals/:id", re.QueryProposalByID)
		app.Listen(":3000")
	}
}

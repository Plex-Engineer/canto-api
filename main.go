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

	app := fiber.New(fiber.Config{
		AppName:      "Canto API",
		ServerHeader: "Fiber",
	})
	app.Get("/", re.GetSmartContractDataFiber)

	routerValidator(app)
	routerCSR(app)
	routerGovernance(app)
	routerStaking(app)
	routerLending(app)
	routerLiquidityPool(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
func routerLiquidityPool(app *fiber.App) {
	liquidity := app.Group("/lp")
	liquidity.Get("/", re.QueryLP)
	liquidity.Get("/:address", re.QueryLpByAddress)
}

func routerLending(app *fiber.App) {
	lending := app.Group("/lending")
	lending.Get("/", re.QueryLending)
	lending.Get("/:address", re.QueryLendingByAddress)
}

func routerValidator(app *fiber.App) {
	validators := app.Group("/validators")
	validators.Get("/", re.QueryValidators)
	validators.Get("/:address", re.QueryValidatorByAddress)

}

func routerCSR(app *fiber.App) {
	csr := app.Group("/csr")
	csr.Get("/", re.QueryCSRs)
	csr.Get("/:id", re.QueryCSRByID)
}

func routerGovernance(app *fiber.App) {
	gov := app.Group("/gov")
	gov.Get("/proposals", re.QueryProposals)
	gov.Get("/proposals/:id", re.QueryProposalByID)
}

func routerStaking(app *fiber.App) {
	staking := app.Group("/staking")
	staking.Get("/apr", re.QueryStakingAPR)
}

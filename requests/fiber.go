package requests

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func Run(ctx context.Context) {
	app := fiber.New(
		fiber.Config{
			AppName:      "Canto API",
			ServerHeader: "Fiber",
		})

	// get all general contract routes
	routes := GetGeneralContractRoutes()
	for _, route := range routes {
		app.Get(route, GetGeneralContractDataFiber)
	}

	routerCSR(app)
	routerGovernance(app)
	routerStaking(app)
	routerLiquidityPool(app)
	routerCtokens(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func routerCtokens(app *fiber.App) {
	lending := app.Group("/v1/lending")
	lending.Get("/ctokens", QueryCTokens)
	lending.Get("/ctoken/:address", QueryCTokenByAddress)
}

func routerLiquidityPool(app *fiber.App) {
	liquidity := app.Group("/v1/dex")
	liquidity.Get("/pairs", QueryPairs)
	liquidity.Get("/pair/:address", QueryPairsByAddress)
}

func routerCSR(app *fiber.App) {
	csr := app.Group("/v1/csr")
	csr.Get("/", QueryCSRs)
	csr.Get("/:id", QueryCSRByID)
}

func routerGovernance(app *fiber.App) {
	gov := app.Group("/v1/gov")
	gov.Get("/proposals", QueryProposals)
	gov.Get("/proposals/:id", QueryProposalByID)
}

func routerStaking(app *fiber.App) {
	staking := app.Group("/v1/staking")
	staking.Get("/apr", QueryStakingAPR)
	staking.Get("/validators", QueryValidators)
	staking.Get("/validators/:address", QueryValidatorByAddress)
}

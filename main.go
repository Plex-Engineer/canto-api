package main

import (
	"context"
	"fmt"
	"strings"

	"canto-api/config"
	cqe "canto-api/query/contracts"
	nqe "canto-api/query/native"
	re "canto-api/requests"

	"github.com/gofiber/fiber/v2"
)

// getGeneralContractRoutes returns a slice of routes for general contracts
func getGeneralContractRoutes() []string {
	routes := []string{}
	for _, contract := range config.ContractCalls {
		for index, method := range contract.Methods {
			// check if the contract has keys
			if len(contract.Keys) == 0 {
				// generate route from name, method and argument of contracts
				route := contract.Name + "/" + strings.Split(method, "(")[0]

				if len(contract.Args[index]) != 0 {
					route += "/" + fmt.Sprintf("%v", contract.Args[index][0])
				}

				routes = append(routes, route)
			}
		}
	}

	return routes
}

func main() {

	config.NewConfig()
	ctx := context.Background()
	go cqe.Run(ctx) // run contract query engine
	go nqe.Run(ctx) // run native query engine

	app := fiber.New(fiber.Config{
		AppName:      "Canto API",
		ServerHeader: "Fiber",
	})
	app.Get("/", re.GetGeneralContractDataFiber)

	// get all general contract routes
	routes := getGeneralContractRoutes()

	for _, route := range routes {
		app.Get(route, re.GetGeneralContractDataFiber)
	}

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
	liquidity.Get("/", re.QueryPairs)
	liquidity.Get("/:address", re.QueryPairsByAddress)
}

func routerLending(app *fiber.App) {
	lending := app.Group("/ctokens")
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

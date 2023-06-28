package requestengine

import (
	"context"
	"fmt"
	"strings"

	"canto-api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// GetGeneralContractRoutes returns a slice of routes for general contracts
func GetGeneralContractRoutes() []string {
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
		log.Fatal().Err(err).Msg("Error fiber server")
	}
}

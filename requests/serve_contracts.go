package requests

import (
	"fmt"
	"strings"

	"canto-api/config"
	"canto-api/rediskeys"
	"context"

	"github.com/gofiber/fiber/v2"
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

// Processed Pairs
func QueryPairs(ctx *fiber.Ctx) error {
	val, err := getStoreValueFromKey(rediskeys.ProcessedPairs)
	if err != nil {
		return redisKeyNotFound(ctx, rediskeys.ProcessedPairs)
	}
	return ctx.Status(StatusOkay).SendString(val)
}

func QueryPairsByAddress(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), rediskeys.ProcessedPairsMap, ctx.Params("address")).Result()
	if err != nil {
		val = "Pair not found"
	}
	return ctx.SendString(val)
}

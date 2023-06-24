package requests

import (
	"fmt"
	"strings"

	"canto-api/config"
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

func GetGeneralContractDataFiber(ctx *fiber.Ctx) error {

	// assemble key from route
	var key string
	route := strings.Split(ctx.Route().Path, `/`)

	for index, part := range route {
		if index > 1 {
			key += ":" + part
		} else if index == 1 {
			key += part
		}
	}

	rdb := config.RDB
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		panic(err)
	}
	return ctx.SendString(val)
}

// Processed Pairs
func QueryPairs(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(config.ProcessedPairs)
	if err != nil {
		return RedisKeyNotFound(ctx, config.ProcessedPairs)
	}
	return ctx.Status(StatusOkay).SendString(val)
}

func QueryPairsByAddress(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), config.ProcessedPairsMap, ctx.Params("address")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, config.ProcessedPairsMap)
	}
	return ctx.SendString(val)
}

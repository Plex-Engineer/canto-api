package requests

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"

	contracts "canto-api/query/contracts"

	"canto-api/config"
	"canto-api/rediskeys"
)

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

func GetSmartContractDataFiber(ctx *fiber.Ctx) error {

	rdb := config.RDB

	val, err := rdb.Get(context.Background(), "supplyspeeds:ccanto:0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488").Result()
	if err != nil {
		panic(err)
	}
	return ctx.SendString(val)
}

func getStoreValueFromKey(key string) string {
	rdb := config.RDB
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func QueryLP(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("pairs"))
}

func QueryLpByAddress(ctx *fiber.Ctx) error {
	allPairs := new([]config.Pair)
	pairsJson := getStoreValueFromKey("pairs")
	err := json.Unmarshal([]byte(pairsJson), &allPairs)
	if err != nil {
		return err
	}
	for _, pair := range *allPairs {
		if pair.Address == ctx.Params("address") {
			resp := contracts.GeneralResultToString(pair)
			return ctx.SendString(resp)
		}
	}
	return ctx.SendString("address not found")
}

func QueryLending(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("lending"))
}

func QueryLendingByAddress(ctx *fiber.Ctx) error {
	allCTokens := new([]config.Token)
	cTokensJson := getStoreValueFromKey("lending")
	err := json.Unmarshal([]byte(cTokensJson), &allCTokens)
	if err != nil {
		return err
	}
	for _, cToken := range *allCTokens {
		if cToken.Address == ctx.Params("address") {
			resp := contracts.GeneralResultToString(cToken)
			return ctx.SendString(resp)
		}
	}
	return ctx.SendString("address not found")
}

func QueryStakingAPR(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey(rediskeys.StakingAPR))
}

func QueryValidators(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey(rediskeys.AllValidators))
}
func QueryValidatorByAddress(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), rediskeys.ValidatorMap, ctx.Params("address")).Result()
	if err != nil {
		val = "Validator not found"
	}
	return ctx.SendString(val)
}

// CSR
func QueryCSRs(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey(rediskeys.AllCSRs))
}
func QueryCSRByID(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), rediskeys.CSRMap, ctx.Params("id")).Result()
	if err != nil {
		val = "CSR not found"
	}
	return ctx.SendString(val)
}

// GOVSHUTTLE
func QueryProposals(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey(rediskeys.AllProposals))
}
func QueryProposalByID(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), rediskeys.ProposalMap, ctx.Params("id")).Result()
	if err != nil {
		return ctx.SendString("id not found")
	}
	return ctx.SendString(val)
}

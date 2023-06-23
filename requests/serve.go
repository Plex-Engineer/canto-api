package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"


	"github.com/gofiber/fiber/v2"

	contracts "canto-api/query/contracts"

	"canto-api/config"
	"canto-api/rediskeys"
)


var (
	StatusBadRequest          = fiber.ErrBadRequest
	StatusInternalServerError = fiber.ErrInternalServerError
	StatusOkay                = fiber.StatusOK
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

func getStoreValueFromKey(key string) (string, error) {
	rdb := config.RDB
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func redisKeyNotFound(ctx *fiber.Ctx, key string) error {
	return ctx.Status(StatusInternalServerError.Code).SendString(fmt.Sprintf("%s not found", key))
}
func idNotFound(ctx *fiber.Ctx, key string, id string) error {
	return ctx.Status(StatusBadRequest.Code).SendString(fmt.Sprintf("%s: %s not found", key, id))
}

func QueryLP(ctx *fiber.Ctx) error {
	val, err := getStoreValueFromKey("pairs")
	if err != nil {
		return redisKeyNotFound(ctx, "pairs")
	}
	return ctx.Status(StatusOkay).SendString(val)
}

func QueryLpByAddress(ctx *fiber.Ctx) error {
	allPairs := new([]config.Pair)
	pairsJson, err := getStoreValueFromKey("pairs")
	if err != nil {
		return redisKeyNotFound(ctx, "pairs")
	}
	err = json.Unmarshal([]byte(pairsJson), &allPairs)
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
	return ctx.SendString(getStoreValueFromKey("ctokens"))
}

func QueryLendingByAddress(ctx *fiber.Ctx) error {
	allCTokens := new([]config.Token)
	cTokensJson := getStoreValueFromKey("ctokens")
	err := json.Unmarshal([]byte(cTokensJson), &allCTokens)

	if err != nil {
		return redisKeyNotFound(ctx, "lending")
	}
	err = json.Unmarshal([]byte(cTokensJson), &allCTokens)
	if err != nil {
		return ctx.Status(StatusInternalServerError.Code).SendString(err.Error())
	}
	for _, cToken := range *allCTokens {
		if cToken.Address == ctx.Params("address") {
			resp := contracts.GeneralResultToString(cToken)
			return ctx.SendString(resp)
		}
	}
	return ctx.SendString("address not found")
}

// STAKING
func QueryStakingAPR(ctx *fiber.Ctx) error {
	val, err := getStoreValueFromKey(rediskeys.StakingAPR)
	if err != nil {
		return redisKeyNotFound(ctx, rediskeys.StakingAPR)
	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryValidators(ctx *fiber.Ctx) error {
	val, err := getStoreValueFromKey(rediskeys.AllValidators)
	if err != nil {
		return redisKeyNotFound(ctx, rediskeys.AllValidators)

	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryValidatorByAddress(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), rediskeys.ValidatorMap, ctx.Params("address")).Result()
	if err != nil {
		return idNotFound(ctx, rediskeys.ValidatorMap, ctx.Params("address"))
	}
	return ctx.Status(StatusOkay).SendString(val)
}

// CSR
func QueryCSRs(ctx *fiber.Ctx) error {
	val, err := getStoreValueFromKey(rediskeys.AllCSRs)
	if err != nil {
		return redisKeyNotFound(ctx, rediskeys.AllCSRs)
	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryCSRByID(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), rediskeys.CSRMap, ctx.Params("id")).Result()
	if err != nil {
		return idNotFound(ctx, rediskeys.CSRMap, ctx.Params("id"))
	}
	return ctx.Status(StatusOkay).SendString(val)
}

// GOVSHUTTLE
func QueryProposals(ctx *fiber.Ctx) error {
	val, err := getStoreValueFromKey(rediskeys.AllProposals)
	if err != nil {
		return redisKeyNotFound(ctx, rediskeys.AllProposals)
	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryProposalByID(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), rediskeys.ProposalMap, ctx.Params("id")).Result()
	if err != nil {
		return idNotFound(ctx, rediskeys.ProposalMap, ctx.Params("id"))
	}
	return ctx.Status(StatusOkay).SendString(val)
}

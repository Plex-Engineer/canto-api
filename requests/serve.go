package requests

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"canto-api/config"
	"canto-api/rediskeys"
)

func GetSmartContractDataFiber(ctx *fiber.Ctx) error {
	rdb := config.RDB

	val, err := rdb.Get(context.Background(), "cTokens:cUSDC:getCash").Result()
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

// STAKING
func QueryStakingAPR(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey(redisKeys.StakingAPR))
}

func QueryValidators(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey(redisKeys.AllValidators))
}
func QueryValidatorByAddress(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), redisKeys.ValidatorMap, ctx.Params("address")).Result()
	if err != nil {
		val = "Validator not found"
	}
	return ctx.SendString(val)
}

// CSR
func QueryCSRs(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey(redisKeys.AllCSRs))
}
func QueryCSRByID(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), redisKeys.CSRMap, ctx.Params("id")).Result()
	if err != nil {
		val = "CSR not found"
	}
	return ctx.SendString(val)
}

// GOVSHUTTLE
func QueryProposals(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey(redisKeys.AllProposals))
}
func QueryProposalByID(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), redisKeys.ProposalMap, ctx.Params("id")).Result()
	if err != nil {
		return ctx.SendString("id not found")
	}
	return ctx.SendString(val)
}

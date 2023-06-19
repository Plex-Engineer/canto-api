package requests

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"canto-api/config"
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
	return ctx.SendString(getStoreValueFromKey("stakingApr"))
}

func QueryValidators(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("validators"))
}
func QueryValidatorByAddress(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), "validatorMap", ctx.Params("address")).Result()
	if err != nil {
		val = "Validator not found"
	}
	return ctx.SendString(val)
}

// CSR
func QueryCSRs(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("csrs"))
}
func QueryCSRByID(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), "csrMap", ctx.Params("id")).Result()
	if err != nil {
		val = "CSR not found"
	}
	return ctx.SendString(val)
}

// GOVSHUTTLE
func QueryProposals(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("proposals"))
}
func QueryProposalByID(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), "proposalMap", ctx.Params("id")).Result()
	if err != nil {
		return ctx.SendString("id not found")
	}
	return ctx.SendString(val)
}

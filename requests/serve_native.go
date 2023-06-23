package requests

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"canto-api/config"
	"canto-api/rediskeys"
)

// STAKING
func QueryStakingAPR(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(rediskeys.StakingAPR)
	if err != nil {
		return RedisKeyNotFound(ctx, rediskeys.StakingAPR)
	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryValidators(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(rediskeys.AllValidators)
	if err != nil {
		return RedisKeyNotFound(ctx, rediskeys.AllValidators)

	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryValidatorByAddress(ctx *fiber.Ctx) error {
	err := CheckValidatorAddress(ctx.Params("address"))
	if err != nil {
		return InvalidParameters(ctx, err)
	}
	val, err := config.RDB.HGet(context.Background(), rediskeys.ValidatorMap, ctx.Params("address")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, fmt.Sprintf("validator address: %s ", ctx.Params("address")))
	}
	return ctx.Status(StatusOkay).SendString(val)
}

// CSR
func QueryCSRs(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(rediskeys.AllCSRs)
	if err != nil {
		return RedisKeyNotFound(ctx, rediskeys.AllCSRs)
	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryCSRByID(ctx *fiber.Ctx) error {
	err := CheckIdString(ctx.Params("id"))
	if err != nil {
		return InvalidParameters(ctx, err)
	}
	val, err := config.RDB.HGet(context.Background(), rediskeys.CSRMap, ctx.Params("id")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, fmt.Sprintf("csr nft id: %s ", ctx.Params("id")))
	}
	return ctx.Status(StatusOkay).SendString(val)
}

// GOVSHUTTLE
func QueryProposals(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(rediskeys.AllProposals)
	if err != nil {
		return RedisKeyNotFound(ctx, rediskeys.AllProposals)
	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryProposalByID(ctx *fiber.Ctx) error {
	err := CheckIdString(ctx.Params("id"))
	if err != nil {
		return InvalidParameters(ctx, err)
	}
	val, err := config.RDB.HGet(context.Background(), rediskeys.ProposalMap, ctx.Params("id")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, fmt.Sprintf("proposal id: %s ", ctx.Params("id")))
	}
	return ctx.Status(StatusOkay).SendString(val)
}

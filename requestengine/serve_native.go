package requestengine

import (
	"context"
	"fmt"

	"canto-api/config"

	"github.com/gofiber/fiber/v2"
)

// STAKING
func QueryStakingAPR(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(config.StakingAPR)
	if err != nil {
		return RedisKeyNotFound(ctx, config.StakingAPR)
	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryValidators(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(config.AllValidators)
	if err != nil {
		return RedisKeyNotFound(ctx, config.AllValidators)

	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryValidatorByAddress(ctx *fiber.Ctx) error {
	err := CheckValidatorAddress(ctx.Params("address"))
	if err != nil {
		return InvalidParameters(ctx, err)
	}
	val, err := config.RDB.HGet(context.Background(), config.ValidatorMap, ctx.Params("address")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, fmt.Sprintf("validator address: %s ", ctx.Params("address")))
	}
	return ctx.Status(StatusOkay).SendString(val)
}

// CSR
func QueryCSRs(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(config.AllCSRs)
	if err != nil {
		return RedisKeyNotFound(ctx, config.AllCSRs)
	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryCSRByID(ctx *fiber.Ctx) error {
	err := CheckIdString(ctx.Params("id"))
	if err != nil {
		return InvalidParameters(ctx, err)
	}
	val, err := config.RDB.HGet(context.Background(), config.CSRMap, ctx.Params("id")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, fmt.Sprintf("csr nft id: %s ", ctx.Params("id")))
	}
	return ctx.Status(StatusOkay).SendString(val)
}

// GOVSHUTTLE
func QueryProposals(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(config.AllProposals)
	if err != nil {
		return RedisKeyNotFound(ctx, config.AllProposals)
	}
	return ctx.Status(StatusOkay).SendString(val)
}
func QueryProposalByID(ctx *fiber.Ctx) error {
	err := CheckIdString(ctx.Params("id"))
	if err != nil {
		return InvalidParameters(ctx, err)
	}
	val, err := config.RDB.HGet(context.Background(), config.ProposalMap, ctx.Params("id")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, fmt.Sprintf("proposal id: %s ", ctx.Params("id")))
	}
	return ctx.Status(StatusOkay).SendString(val)
}

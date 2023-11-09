package requestengine

import (
	"context"
	"fmt"

	"canto-api/config"
	nativequeryengine "canto-api/queryengine/native"

	"github.com/gofiber/fiber/v2"
)

// QueryStakingAPR godoc
// @Summary      Query current staking APR
// @Description  return string of current staking APR
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /staking/apr [get]
func QueryStakingAPR(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(config.StakingAPR)
	if err != nil {
		return RedisKeyNotFound(ctx, config.StakingAPR)
	}
	return ctx.Status(StatusOkay).SendString(val)
}

// QueryValidators godoc
// @Summary      Query validator list
// @Description  return json list of validators
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /staking/validators [get]
func QueryValidators(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(config.AllValidators)
	if err != nil {
		return RedisKeyNotFound(ctx, config.AllValidators)

	}
	return ctx.Status(StatusOkay).SendString(val)
}

// QueryValidatorByAddress godoc
// @Summary      Query validator by address
// @Description  return json object of validator
// @Accept       json
// @Produce      json
// @Param        address path string true "validator address"
// @Success      200  {object}  string
// @Router       /staking/validators/{address} [get]
func QueryValidatorByAddress(ctx *fiber.Ctx) error {
	err := CheckValidatorAddress(ctx.Params("address"))
	if err != nil {
		return InvalidParameters(ctx, err)
	}
	val, err := config.RDB.HGet(context.Background(), config.ValidatorMap, ctx.Params("address")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, fmt.Sprintf("validator address: %s ", ctx.Params("address")))
	}
	// generate json result string
	result := nativequeryengine.GeneralResultToString(map[string]interface{}{
		"results": val,
	})
	return ctx.Status(StatusOkay).SendString(result)
}

// QueryCSRs godoc
// @Summary      Query CSR list
// @Description  return json list of CSRs
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /csr [get]
func QueryCSRs(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(config.AllCSRs)
	if err != nil {
		return RedisKeyNotFound(ctx, config.AllCSRs)
	}
	return ctx.Status(StatusOkay).SendString(val)
}

// QueryCSRByID godoc
// @Summary      Query CSR by id
// @Description  return json object of CSR
// @Accept       json
// @Produce      json
// @Param        id path string true "CSR nft id"
// @Success      200  {object}  string
// @Router       /csr/{id} [get]
func QueryCSRByID(ctx *fiber.Ctx) error {
	err := CheckIdString(ctx.Params("id"))
	if err != nil {
		return InvalidParameters(ctx, err)
	}
	val, err := config.RDB.HGet(context.Background(), config.CSRMap, ctx.Params("id")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, fmt.Sprintf("csr nft id: %s ", ctx.Params("id")))
	}
	// generate json result string
	result := nativequeryengine.GeneralResultToString(map[string]interface{}{
		"results": val,
	})
	return ctx.Status(StatusOkay).SendString(result)
}

// QueryProposals godoc
// @Summary      Query proposal list
// @Description  return json list of proposals
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /gov/proposals [get]
func QueryProposals(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(config.AllProposals)
	if err != nil {
		return RedisKeyNotFound(ctx, config.AllProposals)
	}
	return ctx.Status(StatusOkay).SendString(val)
}

// QueryProposals godoc
// @Summary      Query proposal by id
// @Description  return json object of proposal
// @Accept       json
// @Produce      json
// @Param        id path string true "proposal id"
// @Success      200  {object}  string
// @Router       /gov/proposals/{id} [get]
func QueryProposalByID(ctx *fiber.Ctx) error {
	err := CheckIdString(ctx.Params("id"))
	if err != nil {
		return InvalidParameters(ctx, err)
	}
	val, err := config.RDB.HGet(context.Background(), config.ProposalMap, ctx.Params("id")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, fmt.Sprintf("proposal id: %s ", ctx.Params("id")))
	}
	// generate json result string
	result := nativequeryengine.GeneralResultToString(map[string]interface{}{
		"results": val,
	})
	return ctx.Status(StatusOkay).SendString(result)
}

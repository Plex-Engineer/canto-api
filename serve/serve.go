package serve

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"canto-api/config"
	"canto-api/query"

	csr "github.com/Canto-Network/Canto/v6/x/csr/types"
)

func GetSmartContractDataFiber(ctx *fiber.Ctx) error {
	rdb := config.RDB

	val, err := rdb.Get(context.Background(), "key").Result()
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

func QueryStakingAPR(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("stakingApr"))
}

func QueryValidators(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("validators"))
}
func QueryValidatorByAddress(ctx *fiber.Ctx) error {
	allValidators := new([]query.GetValidatorsResponse)
	validatorJson := getStoreValueFromKey("validators")
	json.Unmarshal([]byte(validatorJson), &allValidators)
	for _, validator := range *allValidators {
		if validator.OperatorAddress == ctx.Params("address") {
			resp := query.GeneralResultToString(validator)
			return ctx.SendString(resp)
		}
	}
	return ctx.SendString("address not found")
}

func QueryCSRs(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("csrs"))
}
func QueryCSRByID(ctx *fiber.Ctx) error {
	numIdQuery, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.SendString("id not found")
	}

	allCSRS := new([]csr.CSR)
	csrJson := getStoreValueFromKey("csrs")
	json.Unmarshal([]byte(csrJson), &allCSRS)
	for _, csr := range *allCSRS {
		if int(csr.Id) == numIdQuery {
			return ctx.SendString(csr.String())
		}
	}
	return ctx.SendString("id not found")
}

func QueryProposals(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("proposals"))
}
func QueryProposalByID(ctx *fiber.Ctx) error {
	numIdQuery, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.SendString("id not found")
	}

	allProposals := new([]query.GetProposalsResponse)
	proposalJson := getStoreValueFromKey("proposals")
	json.Unmarshal([]byte(proposalJson), &allProposals)

	for _, proposal := range *allProposals {
		if int(proposal.ProposalId) == numIdQuery {
			resp := query.GeneralResultToString(proposal)
			return ctx.SendString(resp)
		}
	}
	return ctx.SendString("id not found")
}

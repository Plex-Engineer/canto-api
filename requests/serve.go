package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"canto-api/config"
	contracts "canto-api/query/contracts"
	native "canto-api/query/native"

	csr "github.com/Canto-Network/Canto/v6/x/csr/types"
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
	return ctx.SendString(getStoreValueFromKey("stakingApr"))
}

func QueryValidators(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("validators"))
}
func QueryValidatorByAddress(ctx *fiber.Ctx) error {
	allValidators := new([]native.GetValidatorsResponse)
	validatorJson := getStoreValueFromKey("validators")
	err := json.Unmarshal([]byte(validatorJson), &allValidators)
	if err != nil {
		return err
	}
	for _, validator := range *allValidators {
		if validator.OperatorAddress == ctx.Params("address") {
			resp := contracts.GeneralResultToString(validator)
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

	cErr := json.Unmarshal([]byte(csrJson), &allCSRS)
	if cErr != nil {
		return err
	}
	for _, csrItem := range *allCSRS {
		if int(csrItem.Id) == numIdQuery {
			return ctx.SendString(csrItem.String())
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

	allProposals := new([]native.GetProposalsResponse)
	proposalJson := getStoreValueFromKey("proposals")
	pErr := json.Unmarshal([]byte(proposalJson), &allProposals)
	if pErr != nil {
		return pErr
	}

	for _, proposal := range *allProposals {
		if int(proposal.ProposalId) == numIdQuery {
			resp := contracts.GeneralResultToString(proposal)
			return ctx.SendString(resp)
		}
	}
	return ctx.SendString("id not found")
}

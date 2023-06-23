package query

import (
	"context"

	query "github.com/cosmos/cosmos-sdk/types/query"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type Validator struct {
	// operator_address defines the address of the validator's operator; bech encoded in JSON.
	OperatorAddress string `json:"operator_address"`
	// jailed defined whether the validator has been jailed from bonded status or not.
	Jailed bool `json:"jailed"`
	// status defines the validator's status (bonded(3)/unbonding(2)/unbonded(1)).
	Status string `json:"status"`
	// tokens defines the amount of staking tokens delegated to the validator.
	Tokens string `json:"tokens"`
	// description of validator includes moniker, identity, website, security contact, and details.
	Description staking.Description `json:"description"`
	// commission defines the commission rate.
	Commission string `json:"commission"`
}

// get all Validators for staking
// will return full response string and mapping of operator address to response string
func GetValidators(ctx context.Context, queryClient staking.QueryClient) ([]Validator, map[string]string) {
	respValidators, err := queryClient.Validators(ctx, &staking.QueryValidatorsRequest{
		Pagination: &query.PageRequest{
			Limit: 500,
		},
	})
	CheckError(err)
	allValidators := new([]Validator)
	validatorMap := make(map[string]string)
	for _, validator := range respValidators.Validators {
		valResponse := Validator{
			OperatorAddress: validator.OperatorAddress,
			Jailed:          validator.Jailed,
			Status:          validator.Status.String(),
			Tokens:          validator.Tokens.String(),
			Description:     validator.Description,
			Commission:      validator.Commission.CommissionRates.Rate.String(),
		}
		*allValidators = append(*allValidators, valResponse)
		validatorMap[validator.OperatorAddress] = GeneralResultToString(valResponse)
	}
	return *allValidators, validatorMap
}
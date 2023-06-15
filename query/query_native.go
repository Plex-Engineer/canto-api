package query

import (
	"context"
	"errors"
	"fmt"
	"time"

	"canto-api/config"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"

	csr "github.com/Canto-Network/Canto/v6/x/csr/types"
	inflation "github.com/Canto-Network/Canto/v6/x/inflation/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	query "github.com/cosmos/cosmos-sdk/types/query"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type NativeQueryEngine struct {
	redisclient *redis.Client
	interval    time.Duration
	//query handlers
	CSRQueryHandler       csr.QueryClient
	StakingQueryHandler   staking.QueryClient
	InflationQueryHandler inflation.QueryClient
}

// Returns a NativeQueryEngine instance
func NewNativeQueryEngine() *NativeQueryEngine {
	return &NativeQueryEngine{
		redisclient:           config.RDB,
		interval:              time.Duration(config.QueryInterval),
		CSRQueryHandler:       csr.NewQueryClient(config.GrpcClient),
		StakingQueryHandler:   staking.NewQueryClient(config.GrpcClient),
		InflationQueryHandler: inflation.NewQueryClient(config.GrpcClient),
	}
}

func (nqe *NativeQueryEngine) SetCacheWithResult(ctx context.Context, key string, result interface{}) error {
	// set key in redis
	ret := GeneralResultToString(&result)
	err := nqe.redisclient.Set(ctx, key, ret, 0).Err()
	if err != nil {
		return errors.New("NativeQueryEngine::SetCacheWithResult - " + err.Error())
	}
	return nil
}

// get all CSRS
// TODO: finish this when csr storage is ready
func getCSRS(ctx context.Context, queryClient csr.QueryClient) {
	resp, err := queryClient.CSRs(ctx, &csr.QueryCSRsRequest{})
	checkError(err)
	fmt.Println(resp)
}

// make type for what will be returned from getValidatrs
type GetValidatorsResponse struct {
	// operator_address defines the address of the validator's operator; bech encoded in JSON.
	OperatorAddress string
	// jailed defined whether the validator has been jailed from bonded status or not.
	Jailed bool
	// status defines the validator's status (bonded(3)/unbonding(2)/unbonded(1)).
	Status string
	// tokens defines the amount of staking tokens delegated to the validator.
	Tokens string
	// description of validator includes moniker, identity, website, security contact, and details.
	Description staking.Description
	// commission defines the commission rate.
	Commission string
}

// get all Validators for staking
func getValidators(ctx context.Context, queryClient staking.QueryClient) []GetValidatorsResponse {
	validators, err := queryClient.Validators(ctx, &staking.QueryValidatorsRequest{
		Pagination: &query.PageRequest{
			Limit: 500,
		},
	})
	checkError(err)
	modifiedValidators := new([]GetValidatorsResponse)
	for _, validator := range validators.Validators {
		*modifiedValidators = append(*modifiedValidators, GetValidatorsResponse{
			OperatorAddress: validator.OperatorAddress,
			Jailed:          validator.Jailed,
			Status:          validator.Status.String(),
			Tokens:          validator.Tokens.String(),
			Description:     validator.Description,
			Commission:      validator.Commission.CommissionRates.Rate.String(),
		})
	}
	return *modifiedValidators
}

// StartNativeQueryEngine starts the query engine and runs the ticker
// on the interval specified in config
func (nqe *NativeQueryEngine) StartQueryEngine(ctx context.Context) {
	ticker := time.NewTicker(nqe.interval * time.Second)
	for range ticker.C {
		//get pool
		pool, err := nqe.StakingQueryHandler.Pool(ctx, &staking.QueryPoolRequest{})
		checkError(err)

		//get mint provision
		mintProvision, err := nqe.InflationQueryHandler.EpochMintProvision(ctx, &inflation.QueryEpochMintProvisionRequest{}, &grpc.EmptyCallOption{})
		checkError(err)

		stakingApr := GetStakingAPR(*pool, *mintProvision)

		err = nqe.SetCacheWithResult(ctx, "stakingApr", stakingApr)
		checkError(err)

		//VALIDATORS
		validators := getValidators(ctx, nqe.StakingQueryHandler)
		err = nqe.SetCacheWithResult(ctx, "validators", validators)
		checkError(err)
	}
}

// RunNative initializes a NativeQueryEngine and starts it
func RunNative(ctx context.Context) {
	nqe := NewNativeQueryEngine()
	nqe.StartQueryEngine(ctx)
}

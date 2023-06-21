package query

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"canto-api/config"
	"canto-api/rediskeys"

	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"

	csr "github.com/Canto-Network/Canto/v6/x/csr/types"
	inflation "github.com/Canto-Network/Canto/v6/x/inflation/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	query "github.com/cosmos/cosmos-sdk/types/query"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type NativeQueryEngine struct {
	redisclient *redis.Client
	interval    time.Duration
	//query handlers
	CSRQueryHandler       csr.QueryClient
	GovQueryHandler       gov.QueryClient
	InflationQueryHandler inflation.QueryClient
	StakingQueryHandler   staking.QueryClient
}

// Returns a NativeQueryEngine instance
func NewNativeQueryEngine() *NativeQueryEngine {
	return &NativeQueryEngine{
		redisclient:           config.RDB,
		interval:              time.Duration(config.QueryInterval),
		CSRQueryHandler:       csr.NewQueryClient(config.GrpcClient),
		GovQueryHandler:       gov.NewQueryClient(config.GrpcClient),
		InflationQueryHandler: inflation.NewQueryClient(config.GrpcClient),
		StakingQueryHandler:   staking.NewQueryClient(config.GrpcClient),
	}
}

// set json to to cache (will be list of structs, or single strings)
func (nqe *NativeQueryEngine) SetJsonToCache(ctx context.Context, key string, result interface{}) error {
	// set key in redis
	ret := GeneralResultToString(result)
	err := nqe.redisclient.Set(ctx, key, ret, 0).Err()
	if err != nil {
		return errors.New("NativeQueryEngine::SetJsonToCache - " + err.Error())
	}
	return nil
}

// set mapping to cache (to easy lookup by id in queries)
func (nqe *NativeQueryEngine) SetMappingToCache(ctx context.Context, key string, result map[string]string) error {
	//set key in redis
	err := nqe.redisclient.HSet(ctx, key, result).Err()
	if err != nil {
		return errors.New("NativeQueryEngine::SetMappingToCache - " + err.Error())
	}
	return nil
}

// CSR
type CSR struct {
	// ID of the CSR
	Id uint64 `json:"id"`
	// all contracts under this csr id
	Contracts []string `json:"contracts"`
	// total number of transactions under this csr id
	Txs uint64 `json:"txs"`
	// The cumulative revenue for this CSR NFT -> represented as a big.Int
	Revenue string `json:"revenue"`
}

// get all CSRS
// will return full response string and mapping of nft id to response string
func getCSRS(ctx context.Context, queryClient csr.QueryClient) ([]CSR, map[string]string) {
	resp, err := queryClient.CSRs(ctx, &csr.QueryCSRsRequest{Pagination: &query.PageRequest{
		Limit: 500,
	}})
	CheckError(err)
	allCsrs := new([]CSR)
	csrMap := make(map[string]string)
	for _, csr := range resp.GetCsrs() {
		csrResponse := CSR{
			Id:        csr.GetId(),
			Contracts: csr.GetContracts(),
			Txs:       csr.GetTxs(),
			Revenue:   csr.Revenue.String(),
		}
		*allCsrs = append(*allCsrs, csrResponse)
		csrMap[strconv.Itoa(int(csr.GetId()))] = GeneralResultToString(csrResponse)
	}
	return *allCsrs, csrMap
}

// STAKING

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
func getValidators(ctx context.Context, queryClient staking.QueryClient) ([]Validator, map[string]string) {
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

// StartNativeQueryEngine starts the query engine and runs the ticker
// on the interval specified in config
func (nqe *NativeQueryEngine) StartQueryEngine(ctx context.Context) {
	ticker := time.NewTicker(nqe.interval * time.Second)
	for range ticker.C {
		//
		// STAKING
		//
		// get pool
		pool, err := nqe.StakingQueryHandler.Pool(ctx, &staking.QueryPoolRequest{})
		CheckError(err)

		// get mint provision
		mintProvision, err := nqe.InflationQueryHandler.EpochMintProvision(ctx, &inflation.QueryEpochMintProvisionRequest{}, &grpc.EmptyCallOption{})
		CheckError(err)

		// get global staking apr
		stakingApr := GetStakingAPR(*pool, *mintProvision)

		// save to cache
		err = nqe.SetJsonToCache(ctx, rediskeys.StakingAPR, stakingApr)
		CheckError(err)

		// get and save all validators to cache
		validators, validatorMap := getValidators(ctx, nqe.StakingQueryHandler)
		err = nqe.SetJsonToCache(ctx, rediskeys.AllValidators, validators)
		CheckError(err)
		err = nqe.SetMappingToCache(ctx, rediskeys.ValidatorMap, validatorMap)
		CheckError(err)

		//
		// CSR
		//
		csrs, csrMap := getCSRS(ctx, nqe.CSRQueryHandler)
		err = nqe.SetJsonToCache(ctx, rediskeys.AllCSRs, csrs)
		CheckError(err)
		err = nqe.SetMappingToCache(ctx, rediskeys.CSRMap, csrMap)
		CheckError(err)

		//
		// GOVSHUTTLE
		//
		proposals, proposalMap := GetAllProposals(ctx, nqe.GovQueryHandler)
		err = nqe.SetJsonToCache(ctx, rediskeys.AllProposals, proposals)
		CheckError(err)
		err = nqe.SetMappingToCache(ctx, rediskeys.ProposalMap, proposalMap)
		CheckError(err)

		userVote, err := GetUserVote(ctx, nqe.GovQueryHandler, 89, "0x4BBd79Bb8293c85b216F2ef337aaCf1D0F401FC7")
		fmt.Println(userVote)
		fmt.Println(err)
	}
}

// RunNative initializes a NativeQueryEngine and starts it
func Run(ctx context.Context) {
	nqe := NewNativeQueryEngine()
	nqe.StartQueryEngine(ctx)
}

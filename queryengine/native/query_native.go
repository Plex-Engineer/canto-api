package queryengine

import (
	"context"
	"errors"
	"time"

	"canto-api/config"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	csr "github.com/Canto-Network/Canto/v6/x/csr/types"
	inflation "github.com/Canto-Network/Canto/v6/x/inflation/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func printError(err error) {
	if err != nil {
		log.Fatal().Err(err).Msg("NativeQueryEngine:" + err.Error())
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
		return errors.New("SetJsonToCache:" + err.Error())
	}
	return nil
}

// set mapping to cache (to easy lookup by id in queries)
func (nqe *NativeQueryEngine) SetMappingToCache(ctx context.Context, key string, result map[string]string) error {
	//set key in redis
	err := nqe.redisclient.HSet(ctx, key, result).Err()
	if err != nil {
		return errors.New("SetMappingToCache:" + err.Error())
	}
	return nil
}

// StartNativeQueryEngine starts the query engine and runs the ticker
// on the interval specified in config
func (nqe *NativeQueryEngine) StartQueryEngine(ctx context.Context) {
	ticker := time.NewTicker(nqe.interval * time.Second)
	for range ticker.C {
		//
		// STAKING
		//
		stakingApr, err := GetStakingAPR(ctx, nqe.StakingQueryHandler, nqe.InflationQueryHandler)
		printError(err)
		// save to cache
		err = nqe.SetJsonToCache(ctx, config.StakingAPR, stakingApr)
		printError(err)

		// get and save all validators to cache
		validators, validatorMap, err := GetValidators(ctx, nqe.StakingQueryHandler)
		printError(err)
		err = nqe.SetJsonToCache(ctx, config.AllValidators, validators)
		printError(err)
		err = nqe.SetMappingToCache(ctx, config.ValidatorMap, validatorMap)
		printError(err)

		//
		// CSR
		//
		csrs, csrMap, err := GetCSRS(ctx, nqe.CSRQueryHandler)
		printError(err)
		err = nqe.SetJsonToCache(ctx, config.AllCSRs, csrs)
		printError(err)
		err = nqe.SetMappingToCache(ctx, config.CSRMap, csrMap)
		printError(err)

		//
		// GOVSHUTTLE
		//
		proposals, proposalMap, err := GetAllProposals(ctx, nqe.GovQueryHandler)
		printError(err)
		err = nqe.SetJsonToCache(ctx, config.AllProposals, proposals)
		printError(err)
		err = nqe.SetMappingToCache(ctx, config.ProposalMap, proposalMap)
		printError(err)
	}
}

// RunNative initializes a NativeQueryEngine and starts it
func Run(ctx context.Context) {
	nqe := NewNativeQueryEngine()
	nqe.StartQueryEngine(ctx)
}

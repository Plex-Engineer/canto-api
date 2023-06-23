package query

import (
	"context"
	"errors"
	"time"

	"canto-api/config"
	"canto-api/rediskeys"

	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"

	csr "github.com/Canto-Network/Canto/v6/x/csr/types"
	inflation "github.com/Canto-Network/Canto/v6/x/inflation/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

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
		validators, validatorMap := GetValidators(ctx, nqe.StakingQueryHandler)
		err = nqe.SetJsonToCache(ctx, rediskeys.AllValidators, validators)
		CheckError(err)
		err = nqe.SetMappingToCache(ctx, rediskeys.ValidatorMap, validatorMap)
		CheckError(err)

		//
		// CSR
		//
		csrs, csrMap := GetCSRS(ctx, nqe.CSRQueryHandler)
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

	}
}

// RunNative initializes a NativeQueryEngine and starts it
func Run(ctx context.Context) {
	nqe := NewNativeQueryEngine()
	nqe.StartQueryEngine(ctx)
}

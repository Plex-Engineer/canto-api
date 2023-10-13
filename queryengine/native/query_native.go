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
	// generate json result string
	result = GeneralResultToString(map[string]interface{}{
		"results": ret,
	})
	err := nqe.redisclient.Set(ctx, key, result, 0).Err()
	if err != nil {
		return errors.New("SetJsonToCache: " + err.Error())
	}
	return nil
}

// set mapping to cache (to easy lookup by id in queries)
func (nqe *NativeQueryEngine) SetMapToCache(ctx context.Context, key string, result map[string]string) error {
	//set key in redis
	err := nqe.redisclient.HSet(ctx, key, result).Err()
	if err != nil {
		return errors.New("SetMappingToCache: " + err.Error())
	}
	return nil
}

func nativeQueryEngineFatalLog(err error, function string, msg string) {
	log.Fatal().
		Err(err).
		Str("func", function).
		Msg(msg)
}

// StartNativeQueryEngine starts the query engine and runs the ticker
// on the interval specified in config
func (nqe *NativeQueryEngine) StartNativeQueryEngine(ctx context.Context) {
	ticker := time.NewTicker(nqe.interval * time.Second)
	for range ticker.C {
		//
		// STAKING
		//
		stakingApr, err := GetStakingAPR(ctx, nqe.StakingQueryHandler, nqe.InflationQueryHandler)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to get staking APR")
		}
		// save to cache
		err = nqe.SetJsonToCache(ctx, config.StakingAPR, stakingApr)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to set staking APR")
		}
		// get and save all validators to cache
		validators, validatorMap, err := GetValidators(ctx, nqe.StakingQueryHandler)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to get validators")
		}
		err = nqe.SetJsonToCache(ctx, config.AllValidators, validators)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to set validators")
		}
		err = nqe.SetMapToCache(ctx, config.ValidatorMap, validatorMap)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to set validator map")
		}

		//
		// CSR
		//
		csrs, csrMap, err := GetCSRS(ctx, nqe.CSRQueryHandler)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to get CSRs")
		}
		err = nqe.SetJsonToCache(ctx, config.AllCSRs, csrs)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to set CSRs")
		}
		err = nqe.SetMapToCache(ctx, config.CSRMap, csrMap)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to set CSR map")
		}

		//
		// GOVSHUTTLE
		//
		proposals, proposalMap, httpString, err := GetAllProposals(ctx, nqe.GovQueryHandler)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to get proposals")
		}
		err = nqe.SetJsonToCache(ctx, config.AllProposals, proposals)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to set proposals")
		}
		err = nqe.SetMapToCache(ctx, config.ProposalMap, proposalMap)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to set proposal map")
		}
		err = nqe.SetJsonToCache(ctx, config.AllProposalsHttp, httpString)
		if err != nil {
			nativeQueryEngineFatalLog(err, "StartNativeQueryEngine", "failed to set proposals http string")
		}
	}
}

// RunNative initializes a NativeQueryEngine and starts it
func Run(ctx context.Context) {
	nqe := NewNativeQueryEngine()
	nqe.StartNativeQueryEngine(ctx)
}

package query

import (
	"context"
	"fmt"
	"time"

	"canto-api/config"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"

	csr "github.com/Canto-Network/Canto/v6/x/csr/types"
	inflation "github.com/Canto-Network/Canto/v6/x/inflation/types"
	"github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	// "google.golang.org/grpc"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

/**
*OTHER WAY OF QUERYING GRPC
 */
// out := new(inflation.QueryEpochMintProvisionResponse)
// err11 := config.GrpcClient.Invoke(ctx, "/canto.inflation.v1.Query/EpochMintProvision", &inflation.QueryEpochMintProvisionRequest{}, out)
// checkError(err11)
// fmt.Println(out)
// fmt.Println(out.GetEpochMintProvision().Amount)

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

// get all CSRS
func getCSRS(ctx context.Context, queryClient csr.QueryClient) {
	resp, err := queryClient.CSRs(ctx, &csr.QueryCSRsRequest{})
	if err != nil {
		fmt.Println("Error: ", resp)
	}
	fmt.Println(resp)
}

// get staking apr
func (nqe *NativeQueryEngine) getStakingAPR(ctx context.Context) {
	//get bonded tokens
	pool, err := nqe.StakingQueryHandler.Pool(ctx, &staking.QueryPoolRequest{})
	checkError(err)
	bondedTokens := pool.GetPool().BondedTokens

	//get mint provision from epoch
	epoch, err := nqe.InflationQueryHandler.EpochMintProvision(ctx, &inflation.QueryEpochMintProvisionRequest{}, &grpc.EmptyCallOption{})
	checkError(err)
	
	//get amount (will be in acanto)
	mintProvision := epoch.GetEpochMintProvision().Amount

	//calculate apr (mint provision / bonded tokens) * 365 (days) * 100%
	apr := mintProvision.Mul(types.NewDec(36500)).QuoInt(bondedTokens)

	//print apr
	fmt.Println(apr)
}

// StartNativeQueryEngine starts the query engine and runs the ticker
// on the interval specified in config
func (nqe *NativeQueryEngine) StartQueryEngine(ctx context.Context) {
	ticker := time.NewTicker(nqe.interval * time.Second)
	for range ticker.C {
		// getCSRS(ctx, nqe.CSRQueryHandler)
		nqe.getStakingAPR(ctx)

		resp1, err1 := nqe.StakingQueryHandler.Validators(ctx, &staking.QueryValidatorsRequest{})
		if err1 != nil {
			fmt.Println("Error: ", err1)
		}
		fmt.Println("Error: ", resp1)
	}
}

// RunNative initializes a NativeQueryEngine and starts it
func RunNative(ctx context.Context) {
	nqe := NewNativeQueryEngine()
	nqe.StartQueryEngine(ctx)
}

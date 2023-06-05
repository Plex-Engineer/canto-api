package query

import (
	"context"
	"log"
	"time"

	"canto-api/config"
	"canto-api/multicall"

	"github.com/redis/go-redis/v9"
)

// QueryEngine queries smart contracts directly from a node
// and stores the data in a Redis database on a regular interval.
type QueryEngine struct {
	redisclient *redis.Client
	interval    time.Duration
	mcinstance  *multicall.Multicall
	viewcalls   multicall.ViewCalls
}

// Returns a QueryEngine instance with all necessary objects for
// query engine to run.
func NewQueryEngine() *QueryEngine {

	mc, err := multicall.NewMulticall(config.MulticallAddress, config.EthClient)
	if err != nil {
		log.Fatal(err)
	}

	vcs, err := ProcessContractCalls(config.ContractCalls)
	if err != nil {
		log.Fatal(err)
	}

	return &QueryEngine{
		redisclient: config.RDB,
		interval:    time.Duration(config.QueryInterval),
		mcinstance:  mc,
		viewcalls:   vcs,
	}
}

func (qe *QueryEngine) StartQueryEngine(ctx context.Context) {
	calldata := GetCallData(qe.viewcalls)

	ticker := time.NewTicker(qe.interval * time.Second)
	for range ticker.C {
		// call functions in multicall contract
		res, err := qe.mcinstance.Aggregate(nil, calldata)
		if err != nil {
			log.Fatal(err)
		}

		ret, err := qe.viewcalls.Decode(res)
		if err != nil {
			log.Fatal(err)
		}

		SetCacheWithResult(ctx, qe.redisclient, ret)
	}
}

func Run(ctx context.Context) {
	qe := NewQueryEngine()
	qe.StartQueryEngine(ctx)
}

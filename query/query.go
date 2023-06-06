package query

import (
	"context"
	"fmt"
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

// SetCacheWithResult sets the result of a multicall query in Redis
// and returns an error if any occur.
func (qe *QueryEngine) SetCacheWithResult(ctx context.Context, redisclient *redis.Client, results *multicall.Result) error {
	// convert result slice to string
	ret := ResultToString(results)
	fmt.Println("Reponse Data------------------------------------------------------------------------------------------------------", ret)
	// set key in redis
	// err := redisclient.Set(ctx, "key", string(ret), 0).Err()
	// if err != nil {
	// 	return errors.New("QueryEngine::SetCacheWithResult - " + err.Error())
	// }
	return nil
}

// StartQueryEngine starts the query engine and runs the ticker
// on the interval specified in config .
func (qe *QueryEngine) StartQueryEngine(ctx context.Context) {
	calldata, err := GetCallData(qe.viewcalls)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Payload-------------", payload)
	ticker := time.NewTicker(qe.interval * time.Second)
	for range ticker.C {
		// call functions in multicall contract
		res, err := qe.mcinstance.Aggregate(nil, calldata)
		if err != nil {
			log.Fatal(err)
		}

		// decode results
		ret, err := qe.viewcalls.Decode(res)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println("Response data is--------------", ret)
		// set results to redis cache
		err = qe.SetCacheWithResult(ctx, qe.redisclient, ret)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Run initializes a QueryEngine instance and starts it.
func Run(ctx context.Context) {
	qe := NewQueryEngine()
	qe.StartQueryEngine(ctx)
}

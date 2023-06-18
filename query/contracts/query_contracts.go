package query

import (
	"context"
	"errors"
	"log"
	"time"

	"canto-api/config"
	"canto-api/multicall"

	"github.com/redis/go-redis/v9"
)

// gets viewcalls from the contracts
func ProcessContractCalls(contracts []config.Contract) (multicall.ViewCalls, error) {

	vcs := multicall.ViewCalls{}

	for _, contract := range contracts {
		for index, method := range contract.Methods {
			// validate address
			if err := validateAddress(contract.Address); err != nil {
				return nil, err
			}
			vc := multicall.NewViewCall(
				contract.Keys[index],
				contract.Address,
				method,
				contract.Args[index],
			)

			if err := vc.Validate(); err != nil {
				return nil, errors.New("QueryEngine::ProcessContractCalls - " + err.Error())
			}

			vcs = append(vcs, vc)
		}
	}

	return vcs, nil
}

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

	for key, value := range results.Calls {
		// convert result slice to string
		ret := ResultToString(value)
		// set key in redis
		err := redisclient.Set(ctx, key, string(ret), 0).Err()
		if err != nil {
			return errors.New("QueryEngine::SetCacheWithResult - " + err.Error())
		}
	}
	return nil
}

// StartQueryEngine starts the query engine and runs the ticker
// on the interval specified in config .
func (qe *QueryEngine) StartQueryEngine(ctx context.Context) {
	calldata, err := GetCallData(qe.viewcalls)
	if err != nil {
		log.Fatal(err)
	}

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

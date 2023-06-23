package query

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"canto-api/config"
	"canto-api/multicall"

	"github.com/redis/go-redis/v9"
)

type TokensMap map[string]map[string]interface{}
type PairsMap map[string]map[string]interface{}

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

// gets viewcalls from the contracts
func ProcessContractCalls(contracts []config.Contract) (multicall.ViewCalls, error) {
	vcs := multicall.ViewCalls{}

	for _, contract := range contracts {
		for index, method := range contract.Methods {
			// validate address
			if err := validateAddress(contract.Address); err != nil {
				return nil, err
			}
			var key string
			// check if the contract has keys
			if len(contract.Keys) == 0 {
				// generate key from name, method and argument of contracts
				key = contract.Name + ":" + strings.Split(method, "(")[0]
				if len(contract.Args[index]) != 0 {
					key += ":" + fmt.Sprintf("%v", contract.Args[index][0])
				}
			} else {
				key = contract.Keys[index]
			}
			vc := multicall.NewViewCall(
				key,
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

func (qe *QueryEngine) ProcessMulticallResults(ctx context.Context, results *multicall.Result) (TokensMap, PairsMap, map[string]interface{}) {
	// Declare and initialize maps for ctokens, pairs and others
	ctokens := make(TokensMap)
	pairs := make(PairsMap)
	others := make(map[string]interface{})

	// Iterate the results to separate them into ctokens, pairs and other according to their keys
	for key, value := range results.Calls {
		// split the keys at ':'
		keys := strings.Split(key, ":")
		if keys[0] == "cTokens" {
			// Check if the keys[1] map(ex: cCanto) is already initialized
			if ctokens[keys[1]] == nil {
				// Initialize the keys[1] map(ex: cCanto)
				ctokens[keys[1]] = make(map[string]interface{})
			}

			// store the key, value for keys[1](ex: cCanto)
			ctokens[keys[1]][keys[2]] = value

		} else if keys[0] == "lpPairs" {
			// Check if the keys[1] map(ex: CantoNoteLP) is already initialized
			if pairs[keys[1]] == nil {
				// Initialize the keys[1] map(ex: CantoNoteLP)
				pairs[keys[1]] = make(map[string]interface{})
			}

			// store the key, value for keys[1](ex: CantoNoteLP)
			pairs[keys[1]][keys[2]] = value

		} else {
			others[key] = value
		}

	}
	return ctokens, pairs, others
}

// SetCacheWithFpi sets the result of ctokens and pairs in Redis
// and returns an error if any occur.
func (qe *QueryEngine) SetCacheWithFpi(ctx context.Context, redisclient *redis.Client, ctokens TokensMap, pairs PairsMap) error {

	var err error

	// convert ctokens to json string
	ctokensJsonString := ResultToString(ctokens)

	// convert pairs to json string
	pairsJsonString := ResultToString(pairs)

	// set ctokens in redis
	err = redisclient.Set(ctx, "ctokens", ctokensJsonString, 0).Err()
	if err != nil {
		return errors.New("QueryEngine::SetCacheWithFpi - " + err.Error())
	}

	// set pairs in redis
	err = redisclient.Set(ctx, "pairs", pairsJsonString, 0).Err()
	if err != nil {
		return errors.New("QueryEngine::SetCacheWithFpi - " + err.Error())
	}

	return nil
}

// SetCacheWithResult sets the result of a multicall query in Redis
// and returns an error if any occur.
func (qe *QueryEngine) SetCacheWithResult(ctx context.Context, redisclient *redis.Client, results map[string]interface{}) error {

	// iterate others map and set keys in redis

	for key, value := range results {
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

		// get ctokens, pairs and others from multicall results
		ctokens, pairs, others := qe.ProcessMulticallResults(ctx, ret)

		// set ctokens, pairs to redis cache with fpi
		err = qe.SetCacheWithFpi(ctx, qe.redisclient, ctokens, pairs)
		if err != nil {
			log.Fatal(err)
		}

		// set results to redis cache
		err = qe.SetCacheWithResult(ctx, qe.redisclient, others)
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

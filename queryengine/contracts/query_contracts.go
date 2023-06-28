package queryengine

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"canto-api/config"
	"canto-api/multicall"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type TokensMap map[string]map[string][]interface{}
type PairsMap map[string]map[string][]interface{}

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
		contractQueryEngineFatalLog(err, "NewQueryEngine", "failed to create multicall instance")
	}

	vcs, err := ProcessContractCalls(config.ContractCalls)
	if err != nil {
		contractQueryEngineFatalLog(err, "NewQueryEngine", "failed to process contract calls")
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
				return nil, errors.New("ProcessContractCalls: " + err.Error())
			}

			vcs = append(vcs, vc)
		}
	}

	return vcs, nil
}

func (qe *QueryEngine) ProcessMulticallResults(ctx context.Context, results *multicall.Result) (TokensMap, PairsMap, map[string][]interface{}, error) {
	// Declare and initialize maps for ctokens, pairs and others
	ctokens := make(TokensMap)
	pairs := make(PairsMap)
	others := make(map[string][]interface{})

	// Iterate the results to separate them into ctokens, pairs and other according to their keys
	for key, value := range results.Calls {
		// split the keys at ':'
		keys := strings.Split(key, ":")
		if keys[0] == "cTokens" {
			if len(keys) < 3 {
				return nil, nil, nil, errors.New("ProcessMulticallResults: invalid key for cTokens")
			}
			// Check if the keys[1] map(ex: address of cCanto) is already initialized
			if ctokens[keys[1]] == nil {
				// Initialize the keys[1] map(ex: address of cCanto)
				ctokens[keys[1]] = make(map[string][]interface{})
			}

			// store the key, value for keys[1](ex: address of cCanto)
			ctokens[keys[1]][keys[2]] = value

		} else if keys[0] == "lpPairs" {
			if len(keys) < 3 {
				return nil, nil, nil, errors.New("ProcessMulticallResults: invalid key for lpPairs")
			}
			// Check if the keys[1] map(ex: address of CantoNoteLP) is already initialized
			if pairs[keys[1]] == nil {
				// Initialize the keys[1] map(ex: CantoNoteLP)
				pairs[keys[1]] = make(map[string][]interface{})
			}

			// store the key, value for keys[1](ex: address of CantoNoteLP)
			pairs[keys[1]][keys[2]] = value

		} else {
			others[key] = value
		}

	}
	return ctokens, pairs, others, nil
}

func contractQueryEngineFatalLog(err error, function string, msg string) {
	log.Fatal().
		Err(err).
		Str("func", function).
		Msg(msg)
}

// StartQueryEngine starts the query engine and runs the ticker
// on the interval specified in config .
func (qe *QueryEngine) StartContractQueryEngine(ctx context.Context) {
	log.Info().Msg("starting query engine")
	// get calldata from multicall contract
	calldata, err := GetCallData(qe.viewcalls)
	if err != nil {
		contractQueryEngineFatalLog(err, "StartContractQueryEngine", "failed to get calldata")
	}

	ticker := time.NewTicker(qe.interval * time.Second)
	for range ticker.C {
		log.Info().Msg("querying contracts...")
		// call functions in multicall contract
		res, err := qe.mcinstance.Aggregate(nil, calldata)
		if err != nil {
			contractQueryEngineFatalLog(err, "StartContractQueryEngine", "failed to aggregate")
		}

		// decode results
		ret, err := qe.viewcalls.Decode(res)
		if err != nil {
			contractQueryEngineFatalLog(err, "StartContractQueryEngine", "failed to decode results")
		}

		// get ctokens, pairs and others from multicall results
		ctokens, pairs, others, err := qe.ProcessMulticallResults(ctx, ret)
		if err != nil {
			contractQueryEngineFatalLog(err, "StartContractQueryEngine", "failed to process multicall results")
		}

		// set general contracts to redis cache
		err = qe.SetCacheWithGeneral(ctx, others)
		if err != nil {
			contractQueryEngineFatalLog(err, "StartContractQueryEngine", "failed to set general contracts to redis cache")
		}

		// process pairs data and set to redis
		err = qe.SetCacheWithProcessedPairs(ctx, pairs)
		if err != nil {
			contractQueryEngineFatalLog(err, "StartContractQueryEngine", "failed to set processed pairs to redis cache")
		}

		// process ctokens data and set to redis
		err = qe.SetCacheWithProcessedCTokens(ctx, ctokens)
		if err != nil {
			contractQueryEngineFatalLog(err, "StartContractQueryEngine", "failed to set processed ctokens to redis cache")
		}
		log.Info().Msg("successfully queried contracts...")
	}
}

// Run initializes a QueryEngine instance and starts it.
func Run(ctx context.Context) {
	qe := NewQueryEngine()
	qe.StartContractQueryEngine(ctx)
}

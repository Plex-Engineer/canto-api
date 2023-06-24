package query

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"canto-api/config"
	"canto-api/multicall"

	"github.com/redis/go-redis/v9"
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

func (qe *QueryEngine) ProcessMulticallResults(ctx context.Context, results *multicall.Result) (TokensMap, PairsMap, map[string][]interface{}) {
	// Declare and initialize maps for ctokens, pairs and others
	ctokens := make(TokensMap)
	pairs := make(PairsMap)
	others := make(map[string][]interface{})

	// Iterate the results to separate them into ctokens, pairs and other according to their keys
	for key, value := range results.Calls {
		// split the keys at ':'
		keys := strings.Split(key, ":")
		if keys[0] == "cTokens" {
			// Check if the keys[1] map(ex: address of cCanto) is already initialized
			if ctokens[keys[1]] == nil {
				// Initialize the keys[1] map(ex: address of cCanto)
				ctokens[keys[1]] = make(map[string][]interface{})
			}

			// store the key, value for keys[1](ex: address of cCanto)
			ctokens[keys[1]][keys[2]] = value

		} else if keys[0] == "lpPairs" {
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
	return ctokens, pairs, others
}

// SetCacheWithFpi sets the result of ctokens and pairs in Redis
// and returns an error if any occurs.
func (qe *QueryEngine) SetCacheWithFpi(ctx context.Context, redisclient *redis.Client, ctokens TokensMap, pairs PairsMap) error {
	// set ctokens as a json string in redis
	err := qe.SetJsonToCache(ctx, redisclient, config.CTokens, ctokens)

	if err != nil {
		return errors.New("QueryEngine::SetCacheWithFpi - " + err.Error())
	}

	// set pairs as a json string in redis
	err = qe.SetJsonToCache(ctx, redisclient, config.Pairs, pairs)

	if err != nil {
		return errors.New("QueryEngine::SetCacheWithFpi - " + err.Error())
	}

	return nil
}

// SetCacheWithResult sets the result of a multicall query in Redis
// and returns an error if any occur.
func (qe *QueryEngine) SetCacheWithResult(ctx context.Context, redisclient *redis.Client, results map[string][]interface{}) error {

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

type ProcessedPair struct {
	Address     string       `json:"address"`
	Symbol      string       `json:"symbol"`
	Decimals    int64        `json:"decimals"`
	Token1      config.Token `json:"token1"`
	Token2      config.Token `json:"token2"`
	Stable      bool         `json:"stable"`
	CDecimal    int64        `json:"cDecimals"`
	CLpAddress  string       `json:"cLpAddress"`
	TotalSupply string       `json:"totalSupply"`
	Tvl         string       `json:"tvl"`
	Ratio       string       `json:"ratio"`
	AToB        bool         `json:"aTob"`
	Price1      string       `json:"price1"`
	Price2      string       `json:"price2"`
	LpPrice     string       `json:"lpPrice"`
	Reserve1    string       `json:"reserve1"`
	Reserve2    string       `json:"reserve2"`
}

// This function takes symbol and data of a pair(ex:CantoAtomLP), calculates and adds additional required data and returns the processed pair data
func (qe *QueryEngine) GetProcessedPairs(ctx context.Context, pairs PairsMap) ([]ProcessedPair, map[string]string) {
	processedPairs := []ProcessedPair{}
	processedPairsMap := make(map[string]string)

	// loop over all pairs
	// key is address of lp pair and value is a map of pair data
	for key, value := range pairs {
		// get all the data and process
		reserve1, _ := InterfaceToBigInt(value["reserves"][0])
		reserve2, _ := InterfaceToBigInt(value["reserves"][1])
		totalSupply, _ := InterfaceToBigInt(value["totalSupply"][0])
		price1, _ := InterfaceToString(value["underlyingPriceTokenA"][0])
		price2, _ := InterfaceToString(value["underlyingPriceTokenB"][0])
		lpPrice, _ := InterfaceToBigInt(value["underlyingPriceLp"][0])

		// calculate total value locked by multiplying lp price and total supply using Mul() method
		var tvl = new(big.Int).Mul(lpPrice, totalSupply)

		// divide tvl with 1e18 using Div() method
		tvl.Div(tvl, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

		// get ratio of reserve and reserve2
		ratio, aTob := GetLpPairRatio(reserve1, reserve2)

		// get lp pair data
		symbol, decimals, token1, token2, stable, cDecimals, cLpAddress := config.GetLpPairData(key)
		processedPair := ProcessedPair{
			Address:     key,
			Symbol:      symbol,
			Decimals:    decimals,
			Token1:      token1,
			Token2:      token2,
			Stable:      stable,
			CDecimal:    cDecimals,
			CLpAddress:  cLpAddress,
			TotalSupply: totalSupply.String(),
			Tvl:         tvl.String(),
			Ratio:       ratio.String(),
			AToB:        aTob,
			Price1:      price1,
			Price2:      price2,
			LpPrice:     lpPrice.String(),
			Reserve1:    reserve1.String(),
			Reserve2:    reserve2.String(),
		}

		processedPairs = append(processedPairs, processedPair)
		processedPairsMap[key] = ResultToString(processedPair)
	}
	return processedPairs, processedPairsMap
}

// SetJsonToCache will take key, result and sets the resulte as a json string to redis
func (qe *QueryEngine) SetJsonToCache(ctx context.Context, redisclient *redis.Client, key string, result interface{}) error {
	// convert result to json string
	ret := ResultToString(result)
	err := redisclient.Set(ctx, key, ret, 0).Err()
	if err != nil {
		return errors.New("QueryEngine::SetJsonToCache - " + err.Error())
	}
	return nil
}

// SetMapToCache will take key, result map and sets to redis using HSet()
func (qe *QueryEngine) SetMapToCache(ctx context.Context, redisclient *redis.Client, key string, result map[string]string) error {
	//set key in redis
	err := qe.redisclient.HSet(ctx, key, result).Err()
	if err != nil {
		return errors.New("QueryEngine::SetMapToCache - " + err.Error())
	}
	return nil
}

// This function gets the pairs data from redis, processes it and sets the processed pairs data to redis
func (qe *QueryEngine) SetCacheWithProcessedPairs(ctx context.Context, redisclient *redis.Client) error {
	//get pairs data
	val, err := redisclient.Get(context.Background(), config.Pairs).Result()
	if err != nil {
		return errors.New("QueryEngine::SetCacheWithProcessedPairs - " + err.Error())
	}

	// decalare a pairs variable of type PairsMap
	var pairs PairsMap

	//unmarshal json data
	err = json.Unmarshal([]byte(val), &pairs)

	if err != nil {
		return errors.New("QueryEngine::SetCacheWithProcessedPairs - " + err.Error())
	}

	// get processed pairs data
	processedPairs, processedPairsMap := qe.GetProcessedPairs(ctx, pairs)

	//  set processed pairs as a json string to redis
	err = qe.SetJsonToCache(ctx, redisclient, config.ProcessedPairs, processedPairs)

	if err != nil {
		return errors.New("QueryEngine::SetCacheWithProcessedPairs - " + err.Error())
	}

	//  set processed pairs map as a json string to redis
	err = qe.SetMapToCache(ctx, redisclient, config.ProcessedPairsMap, processedPairsMap)

	if err != nil {
		return errors.New("QueryEngine::SetCacheWithProcessedPairs - " + err.Error())
	}

	// fmt.Println(processedPairs)

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

		// process pairs data and set to redis
		err = qe.SetCacheWithProcessedPairs(ctx, qe.redisclient)
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

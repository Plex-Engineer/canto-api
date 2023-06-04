package query

import (
	"context"
	"log"
	"time"

	"canto-api/config"
	"canto-api/multicall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
)

// QueryEngine queries smart contracts directly from a node
// and stores the data in a Redis database on a regular interval.
type QueryEngine struct {
	redisclient *redis.Client
	ethclient   *ethclient.Client
	interval    time.Duration
	mcaddress   common.Address // multicall contract address
	vcs         multicall.ViewCalls
}

// Returns a QueryEngine instance with all necessary objects for
// query engine to run.
func NewQueryEngine() *QueryEngine {

	vcs, err := ProcessContractCalls(config.ContractCalls)
	if err != nil {
		log.Fatal(err)
	}

	return &QueryEngine{
		redisclient: config.RDB,
		ethclient:   config.EthClient,
		interval:    time.Duration(config.QueryInterval),
		mcaddress:   config.MulticallAddress,
		vcs:         vcs,
	}
}

func (qe *QueryEngine) StartQueryEngine(ctx context.Context) {

	multicallInstance, err := multicall.NewMulticall(qe.mcaddress, qe.ethclient)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := qe.vcs.GetCallData()
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(qe.interval * time.Second)
	for range ticker.C {
		// call functions in multicall contract
		res, err := multicallInstance.Aggregate(nil, payload)
		if err != nil {
			log.Fatal(err)
		}

		ret, err := qe.vcs.Decode(res)
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

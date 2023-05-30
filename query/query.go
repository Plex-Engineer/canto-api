package query

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pair "canto-api/contracts"
	redisclient "canto-api/redisclient"

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
}

// Returns a QueryEngine instance with redis client, eth client, and interval.
func NewQueryEngine() *QueryEngine {
	ethclient, err := ethclient.Dial("https://mainnode.plexnode.org:8545")
	if err != nil {
		log.Fatal(err)
	}

	return &QueryEngine{
		redisclient: redisclient.RDB,
		ethclient:   ethclient,
		interval:    5,
	}
}

func (qe *QueryEngine) StartQueryEngine(ctx context.Context) {

	// canto note pair address
	address := common.HexToAddress("0x1D20635535307208919f0b67c3B2065965A85aA9")

	pairInstance, err := pair.NewPair(address, qe.ethclient)
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(qe.interval * time.Second)
	for range ticker.C {
		// get reserves from pair contract
		reserves, err := pairInstance.GetReserves(nil)
		if err != nil {
			log.Fatal(err)
		}

		// must marshal reserves to json in order to store in redis
		b, err := json.Marshal(reserves)
		if err != nil {
			log.Fatal(err)
		}

		// set key in redis
		err = qe.redisclient.Set(ctx, "key", b, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func Run(ctx context.Context) {
	qe := NewQueryEngine()
	qe.StartQueryEngine(ctx)
}

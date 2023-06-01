package query

import (
	"context"
	"encoding/json"
	"fmt"
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
}

// Returns a QueryEngine instance with redis client, eth client, and interval.
func NewQueryEngine() *QueryEngine {
	return &QueryEngine{
		redisclient: config.RDB,
		ethclient:   config.EthClient,
		interval:    5,
		mcaddress:   common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
	}
}

func (qe *QueryEngine) StartQueryEngine(ctx context.Context) {

	multicallInstance, err := multicall.NewMulticall(qe.mcaddress, qe.ethclient)
	if err != nil {
		log.Fatal(err)
	}

	// vc := multicall.NewViewCall(
	// 	"getReserves()",
	// 	[]interface{}{},
	// )

	vc := multicall.NewViewCall(
		"balanceOf(address)(uint256)",
		[]interface{}{"0x66945A3A0f7D3D85A5d1A309AA321436945A329A"},
	)

	payload, err := vc.CallData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("total payload: ", payload)

	// canto note pair address
	// pairAddress := common.HexToAddress("0x1D20635535307208919f0b67c3B2065965A85aA9")

	// usdc address
	erc20Address := common.HexToAddress("0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd")

	// multicall struct
	reserves := multicall.Multicall3Call{
		Target:   erc20Address,
		CallData: payload,
	}

	multicallArray := []multicall.Multicall3Call{reserves}

	ticker := time.NewTicker(qe.interval * time.Second)
	for range ticker.C {
		// call functions in multicall contract
		res, err := multicallInstance.Aggregate(nil, multicallArray)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res)

		// must marshal reserves to json in order to store in redis
		b, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}

		// ret := abi.Unpack([]string{"uint112", "uint112", "uint32"}, res.ReturnData[0])

		// fmt.Println(b)

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

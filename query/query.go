package query

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"canto-api/config"
	multicall "canto-api/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	return &QueryEngine{
		redisclient: config.RDB,
		ethclient:   config.EthClient,
		interval:    5,
	}
}

func (qe *QueryEngine) StartQueryEngine(ctx context.Context) {

	// multicall contract address
	mcaddress := common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")
	multicallInstance, err := multicall.NewMulticall(mcaddress, qe.ethclient)
	if err != nil {
		log.Fatal(err)
	}

	// canto note pair address
	pairAddress := common.HexToAddress("0x1D20635535307208919f0b67c3B2065965A85aA9")

	methodSuffix := "getReserves()"
	hash := crypto.Keccak256([]byte(methodSuffix))
	methodBytes := hash[:4]

	payload := make([]byte, 0)
	payload = append(payload, methodBytes...)

	// multicall struct
	reserves := multicall.Multicall3Call{
		Target:   pairAddress,
		CallData: payload,
	}

	multicallArray := []multicall.Multicall3Call{reserves}

	fmt.Println(multicallArray)

	// pairInstance, err := multicall.NewPair(pairAddress, qe.ethclient)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := qe.ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID := 7700
	bigID := big.NewInt(int64(chainID))

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, bigID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(0))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	ticker := time.NewTicker(qe.interval * time.Second)
	for range ticker.C {
		// call functions in multicall contract
		fmt.Println("calling multicall contract")
		res, err := multicallInstance.Aggregate(nil, multicallArray)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("called multicall contract")

		// get reserves from pair contract
		// reserves, err := pairInstance.GetReserves(nil)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// must marshal reserves to json in order to store in redis
		b, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}

		// ret := abi.Unpack([]string{"uint112", "uint112", "uint32"}, res.ReturnData[0])

		fmt.Println(b)

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

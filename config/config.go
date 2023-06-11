package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
)

var (
	RDB              *redis.Client
	EthClient        *ethclient.Client
	ContractCalls    []Contract // list of calls to make
	MulticallAddress common.Address
	QueryInterval    uint
)

func NewConfig() {
	// RDB = redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// EthClient, _ = ethclient.Dial("https://canto-testnet.plexnode.wtf")
	// MulticallAddress = common.HexToAddress("0x75c8e3eFA6f1B797f75141c0aecfDc83b83e4bF6")

	EthClient, _ = ethclient.Dial("https://mainnode.plexnode.org:8545")
	MulticallAddress = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")
	ContractCalls = calls
	// ContractCalls = getAllContractCalls()
	// fmt.Println("Contract calls------------------------", ContractCalls)

	QueryInterval = 5
}

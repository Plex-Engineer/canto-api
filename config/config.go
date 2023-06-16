package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

var (
	RDB              *redis.Client
	EthClient        *ethclient.Client
	GrpcClient       *grpc.ClientConn
	ContractCalls    []Contract // list of calls to make
	MulticallAddress common.Address
	QueryInterval    uint
)

func NewConfig() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	EthClient, _ = ethclient.Dial("https://mainnode.plexnode.org:8545")
	GrpcClient, _ = grpc.Dial("143.198.228.162:9090", grpc.WithInsecure())
	calls, err := getContractsFromJson("./config/jsons/contracts.json")
	if err != nil {
		fmt.Println("Error getting contracts from json:", err)
	}
	ContractCalls = calls
	MulticallAddress = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")
	QueryInterval = 5
}

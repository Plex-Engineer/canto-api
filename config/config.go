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
	FPIConfig        TokensInfo
)

/*
 * @brief: NewConfig
 * @param: none
 * @return: none
 * @desc: initialize config variables (acts as a constructor)
 */
func NewConfig() {

	// Initialize redis client
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Initialize eth client using mainnet rpc
	EthClient, _ = ethclient.Dial("https://mainnode.plexnode.org:8545")

	// Initialize grpc client using mainnet rpc
	GrpcClient, _ = grpc.Dial("143.198.228.162:9090", grpc.WithInsecure())

	// get tokens data from tokens.json
	FPIConfig = getAllTokensFromJson("./config/jsons/fpi_mainnet.json")

	// set multicall address
	MulticallAddress = common.HexToAddress(FPIConfig.MulticallV3)

	// set query interval per block (5 seconds)
	QueryInterval = 5

	// get general contracts from contracts.json
	generalCalls, err := getContractsFromJson("./config/jsons/contracts.json")
	if err != nil {
		panic(fmt.Sprintf("Error getting general contracts: %v", err))
	}

	// get FPI contracts from tokens.json
	fpiCalls := getAllFPI()
	calls := append(fpiCalls, generalCalls...)
	ContractCalls = calls
}

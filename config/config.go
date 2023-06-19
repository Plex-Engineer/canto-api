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
	TokensConfig     TokensInfo
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

	// set multicall address
	MulticallAddress = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")

	// set query interval per block (5 seconds)
	QueryInterval = 5

	// get general contracts from contracts.json
	generalCalls, err := getContractsFromJson("./config/jsons/contracts.json")
	if err != nil {
		fmt.Println("Error getting contracts from json:", err)
	}

	// get FPI contracts from tokens.json
	fpiCalls := getAllFPI("./config/jsons/tokens.json")

	// append calls to get all contract calls
	calls := append(fpiCalls, generalCalls...)
	ContractCalls = calls

	fmt.Println("contract calls: ", ContractCalls)
}

package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
)

var (
	RDB                   *redis.Client
	EthClient             *ethclient.Client
	ContractCalls         []Contract // list of calls to make
	MulticallAddress      common.Address
	QueryInterval         uint
	ParsedTokens          Tokens
	ParsedContractsConfig ContractsConfig
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

	// get parsed tokens from json
	ParsedTokens = getAllTokensFromJson(false)

	// get parsed contracts from json
	ParsedContractsConfig = getAllContractsConfigFromJson()

	// get extra all contract calls from json
	ContractCalls = getAllContractsFromJson()

	// set multicall address
	MulticallAddress = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")

	// set query interval per block (5 seconds)
	QueryInterval = 5
}

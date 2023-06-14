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

func NewConfig() {

	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	EthClient, _ = ethclient.Dial("https://mainnode.plexnode.org:8545")
	ParsedTokens = getAllTokensFromJson(false)
	ParsedContractsConfig = getAllContractsConfigFromJson()
	ContractCalls = getAllContractsFromJson()
	MulticallAddress = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")
	QueryInterval = 5
}

package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	BlockNumber         = "BLOCK_NUMBER"
	StakingAPR          = "STAKING_APR"
	AllValidators       = "ALL_VALIDATORS"
	ValidatorMap        = "VALIDATOR_MAP"
	AllCSRs             = "ALL_CSRS"
	CSRMap              = "CSR_MAP"
	AllProposals        = "ALL_PROPOSALS"
	ProposalMap         = "PROPOSAL_MAP"
	Pairs               = "PAIRS"
	ProcessedPairs      = "PROCESSED_PAIRS"
	ProcessedPairsMap   = "PROCESSED_PAIRS_MAP"
	CTokens             = "CTOKENS"
	ProcessedCTokens    = "PROCESSED_CTOKENS"
	ProcessedCTokensMap = "PROCESSED_CTOKENS_MAP"
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
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// Initialize redis client
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		Password: "",
		DB:       0,
	})

	// Initialize eth client using mainnet rpc
	rpcUrl := os.Getenv("CANTO_MAINNET_RPC_URL")
	ethclient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal().Msgf("Error initializing eth client: %v", err)
	}
	EthClient = ethclient

	// Initialize grpc client using mainnet rpc
	grpcUrl := os.Getenv("CANTO_MAINNET_GRPC_URL")
	GrpcClient, err = grpc.Dial(grpcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("Error initializing grpc client: %v", err)
	}
	// is testnet
	isTestnet := os.Getenv("TESTNET")
	var fpiFile string
	var contractsFile string
	if isTestnet == "true" {
		fpiFile = "./config/jsons/fpi_testnet.json"
		contractsFile = "./config/jsons/contracts_testnet.json"
	} else {
		fpiFile = "./config/jsons/fpi_mainnet.json"
		contractsFile = "./config/jsons/contracts.json"
	}
	// get tokens data from tokens.json
	FPIConfig, err = getFPIFromJson(fpiFile)
	if err != nil {
		log.Fatal().Msgf("Error getting tokens data from json: %v", err)
	}
	// get general contracts from contracts.json
	generalCalls, err := getContractsFromJson(contractsFile)
	if err != nil {
		log.Fatal().Msgf("Error getting general contracts from json: %v", err)
	}

	// set multicall address
	mcAddress := os.Getenv("MULTICALL_ADDRESS")
	MulticallAddress = common.HexToAddress(mcAddress)

	// set query interval per block (5 seconds)
	queryInterval, err := strconv.Atoi(os.Getenv("QUERY_INTERVAL"))
	if err != nil {
		log.Fatal().Msgf("Error converting query interval to int: %v", err)
	}
	QueryInterval = uint(queryInterval)

	// get FPI contracts from tokens.json
	fpiCalls := getAllFPI()
	calls := append(fpiCalls, generalCalls...)
	ContractCalls = calls
}

func SetBackupRPC() {
	// Initialize eth client using backup rpc
	rpcUrl := os.Getenv("CANTO_BACKUP_RPC_URL")
	ethclient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal().Msgf("Error initializing eth client: %v", err)
	}
	EthClient = ethclient
}

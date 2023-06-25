package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
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
	ProcessedCTokens    = "PROCESSED_CTokens"
	ProcessedCTokensMap = "PROCESSED_CTokens_MAP"
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
	ethclient, err := ethclient.Dial("https://mainnode.plexnode.org:8545")
	if err != nil {
		log.Fatal().Msgf("Error initializing eth client: %v", err)
	}
	EthClient = ethclient

	// Initialize grpc client using mainnet rpc
	GrpcClient, err = grpc.Dial("143.198.228.162:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("Error initializing grpc client: %v", err)
	}

	// get tokens data from tokens.json
	FPIConfig, err = getFPIFromJson("./config/jsons/fpi_mainnet.json")
	if err != nil {
		log.Fatal().Msgf("Error getting tokens data from json: %v", err)
	}

	// set multicall address
	MulticallAddress = common.HexToAddress(FPIConfig.MulticallV3)

	// set query interval per block (5 seconds)
	QueryInterval = 5

	// get general contracts from contracts.json
	generalCalls, err := getContractsFromJson("./config/jsons/contracts.json")
	if err != nil {
		log.Fatal().Msgf("Error getting general contracts from json: %v", err)
	}

	// get FPI contracts from tokens.json
	fpiCalls := getAllFPI()
	calls := append(fpiCalls, generalCalls...)
	ContractCalls = calls
}

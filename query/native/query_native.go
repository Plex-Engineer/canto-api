package query

import (
	"context"
	"errors"
	"time"

	"canto-api/config"

	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc"

	csr "github.com/Canto-Network/Canto/v6/x/csr/types"
	inflation "github.com/Canto-Network/Canto/v6/x/inflation/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type NativeQueryEngine struct {
	redisclient *redis.Client
	interval    time.Duration
	//query handlers
	CSRQueryHandler       csr.QueryClient
	GovQueryHandler       gov.QueryClient
	InflationQueryHandler inflation.QueryClient
	StakingQueryHandler   staking.QueryClient
}

// Returns a NativeQueryEngine instance
func NewNativeQueryEngine() *NativeQueryEngine {
	return &NativeQueryEngine{
		redisclient:           config.RDB,
		interval:              time.Duration(config.QueryInterval),
		CSRQueryHandler:       csr.NewQueryClient(config.GrpcClient),
		GovQueryHandler:       gov.NewQueryClient(config.GrpcClient),
		InflationQueryHandler: inflation.NewQueryClient(config.GrpcClient),
		StakingQueryHandler:   staking.NewQueryClient(config.GrpcClient),
	}
}

func (nqe *NativeQueryEngine) SetCacheWithResult(ctx context.Context, key string, result interface{}) error {
	// set key in redis
	ret := GeneralResultToString(result)
	err := nqe.redisclient.Set(ctx, key, ret, 0).Err()
	if err != nil {
		return errors.New("NativeQueryEngine::SetCacheWithResult - " + err.Error())
	}
	return nil
}

// CSR

// get all CSRS
func getCSRS(ctx context.Context, queryClient csr.QueryClient) []csr.CSR {
	resp, err := queryClient.CSRs(ctx, &csr.QueryCSRsRequest{})
	checkError(err)
	return resp.GetCsrs()
}

// STAKING

// make type for what will be returned from getValidatrs
type GetValidatorsResponse struct {
	// operator_address defines the address of the validator's operator; bech encoded in JSON.
	OperatorAddress string
	// jailed defined whether the validator has been jailed from bonded status or not.
	Jailed bool
	// status defines the validator's status (bonded(3)/unbonding(2)/unbonded(1)).
	Status string
	// tokens defines the amount of staking tokens delegated to the validator.
	Tokens string
	// description of validator includes moniker, identity, website, security contact, and details.
	Description staking.Description
	// commission defines the commission rate.
	Commission string
}

// get all Validators for staking
func getValidators(ctx context.Context, queryClient staking.QueryClient) []GetValidatorsResponse {
	validators, err := queryClient.Validators(ctx, &staking.QueryValidatorsRequest{
		Pagination: &query.PageRequest{
			Limit: 500,
		},
	})
	checkError(err)
	modifiedValidators := new([]GetValidatorsResponse)
	for _, validator := range validators.Validators {
		*modifiedValidators = append(*modifiedValidators, GetValidatorsResponse{
			OperatorAddress: validator.OperatorAddress,
			Jailed:          validator.Jailed,
			Status:          validator.Status.String(),
			Tokens:          validator.Tokens.String(),
			Description:     validator.Description,
			Commission:      validator.Commission.CommissionRates.Rate.String(),
		})
	}
	return *modifiedValidators
}

// GOVSHUTTLE
type GetProposalsResponse struct {
	// proposalId defines the unique id of the proposal.
	ProposalId uint64
	// typeUrl indentifies the type of the proposal by a serialized protocol buffer message
	TypeUrl string
	// status defines the current status of the proposal.
	Status string
	// finalVote defined the result of the proposal
	FinalVote gov.TallyResult
	// submitTime defines the block time the proposal was submitted.
	SubmitTime time.Time
	// depositEndTime defines the time when the proposal deposit period will end.
	DepositEndTime time.Time
	// totalDeposit defines the total amount of coins deposited on this proposal
	TotalDeposit sdk.Coins
	// votingStartTime defines the time when the proposal voting period will start
	VotingStartTime time.Time
	// votingEndTime defines the time when the proposal voting period will end
	VotingEndTime time.Time
}

func getAllProposals(ctx context.Context, queryClient gov.QueryClient) []GetProposalsResponse {
	resp, err := queryClient.Proposals(ctx, &gov.QueryProposalsRequest{})
	checkError(err)
	allProposals := new([]GetProposalsResponse)
	for _, proposal := range resp.GetProposals() {
		*allProposals = append(*allProposals, GetProposalsResponse{
			ProposalId:      proposal.ProposalId,
			TypeUrl:         proposal.Content.TypeUrl,
			Status:          proposal.Status.String(),
			FinalVote:       proposal.FinalTallyResult,
			SubmitTime:      proposal.SubmitTime,
			DepositEndTime:  proposal.DepositEndTime,
			TotalDeposit:    proposal.TotalDeposit,
			VotingStartTime: proposal.VotingStartTime,
			VotingEndTime:   proposal.VotingEndTime,
		})
	}
	return *allProposals
}

// StartNativeQueryEngine starts the query engine and runs the ticker
// on the interval specified in config
func (nqe *NativeQueryEngine) StartQueryEngine(ctx context.Context) {
	ticker := time.NewTicker(nqe.interval * time.Second)
	for range ticker.C {
		//
		// STAKING
		//
		// get pool
		pool, err := nqe.StakingQueryHandler.Pool(ctx, &staking.QueryPoolRequest{})
		checkError(err)

		// get mint provision
		mintProvision, err := nqe.InflationQueryHandler.EpochMintProvision(ctx, &inflation.QueryEpochMintProvisionRequest{}, &grpc.EmptyCallOption{})
		checkError(err)

		// get global staking apr
		stakingApr := GetStakingAPR(*pool, *mintProvision)

		// save to cache
		err = nqe.SetCacheWithResult(ctx, "stakingApr", stakingApr)
		checkError(err)

		// get and save all validators to cache
		validators := getValidators(ctx, nqe.StakingQueryHandler)
		err = nqe.SetCacheWithResult(ctx, "validators", validators)
		checkError(err)

		//
		// CSR
		//
		csrs := getCSRS(ctx, nqe.CSRQueryHandler)
		err = nqe.SetCacheWithResult(ctx, "csrs", csrs)
		checkError(err)

		//
		// GOVSHUTTLE
		//
		proposals := getAllProposals(ctx, nqe.GovQueryHandler)
		err = nqe.SetCacheWithResult(ctx, "proposals", proposals)
		checkError(err)
	}
}

// RunNative initializes a NativeQueryEngine and starts it
func Run(ctx context.Context) {
	nqe := NewNativeQueryEngine()
	nqe.StartQueryEngine(ctx)
}

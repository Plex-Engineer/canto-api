package queryengine

import (
	"context"
	"strconv"
	"time"

	csr "github.com/Canto-Network/Canto/v6/x/csr/types"
	inflation "github.com/Canto-Network/Canto/v6/x/inflation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/rs/zerolog/log"
)

// STAKING
type Validator struct {
	// operator_address defines the address of the validator's operator; bech encoded in JSON.
	OperatorAddress string `json:"operator_address"`
	// jailed defined whether the validator has been jailed from bonded status or not.
	Jailed bool `json:"jailed"`
	// status defines the validator's status (bonded(3)/unbonding(2)/unbonded(1)).
	Status string `json:"status"`
	// tokens defines the amount of staking tokens delegated to the validator.
	Tokens string `json:"tokens"`
	// description of validator includes moniker, identity, website, security contact, and details.
	Description staking.Description `json:"description"`
	// commission defines the commission rate.
	Commission string `json:"commission"`
}

// get all Validators for staking
// will return full response string and mapping of operator address to response string
func GetValidators(ctx context.Context, queryClient staking.QueryClient) ([]Validator, map[string]string, error) {
	respValidators, err := queryClient.Validators(ctx, &staking.QueryValidatorsRequest{
		Pagination: &query.PageRequest{
			Limit: 1000,
		},
	})
	if err != nil {
		return nil, nil, err
	}
	allValidators := new([]Validator)
	validatorMap := make(map[string]string)
	for _, validator := range respValidators.Validators {
		valResponse := Validator{
			OperatorAddress: validator.OperatorAddress,
			Jailed:          validator.Jailed,
			Status:          validator.Status.String(),
			Tokens:          validator.Tokens.String(),
			Description:     validator.Description,
			Commission:      validator.Commission.CommissionRates.Rate.String(),
		}
		*allValidators = append(*allValidators, valResponse)
		validatorMap[validator.OperatorAddress] = GeneralResultToString(valResponse)
	}
	return *allValidators, validatorMap, nil
}
func GetStakingAPR(ctx context.Context, stakingQueryClient staking.QueryClient, inflationQueryClient inflation.QueryClient) (string, error) {
	//get pool
	pool, err := stakingQueryClient.Pool(ctx, &staking.QueryPoolRequest{})
	if err != nil {
		return "", err
	}
	// get mint provision
	mintProvision, err := inflationQueryClient.EpochMintProvision(ctx, &inflation.QueryEpochMintProvisionRequest{})
	if err != nil {
		return "", err
	}
	// get global staking apr
	stakingApr := CalculateStakingAPR(*pool, *mintProvision)
	return stakingApr.String(), nil
}

// GOVSHUTTLE

type Proposal struct {
	// proposalId defines the unique id of the proposal.
	ProposalId uint64 `json:"proposal_id"`
	// typeUrl indentifies the type of the proposal by a serialized protocol buffer message
	TypeUrl string `json:"type_url"`
	// title of the proposal
	Title string `json:"title"`
	// description of the proposal
	Description string `json:"description"`
	// status defines the current status of the proposal.
	Status string `json:"status"`
	// finalVote defined the result of the proposal
	FinalVote gov.TallyResult `json:"final_vote"`
	// submitTime defines the block time the proposal was submitted.
	SubmitTime time.Time `json:"submit_time"`
	// depositEndTime defines the time when the proposal deposit period will end.
	DepositEndTime time.Time `json:"deposit_end_time"`
	// totalDeposit defines the total amount of coins deposited on this proposal
	TotalDeposit sdk.Coins `json:"total_deposit"`
	// votingStartTime defines the time when the proposal voting period will start
	VotingStartTime time.Time `json:"voting_start_time"`
	// votingEndTime defines the time when the proposal voting period will end
	VotingEndTime time.Time `json:"voting_end_time"`
}

// get all proposals from gov shuttle
// will return full response string and mapping of proposal id to response string
func GetAllProposals(ctx context.Context, queryClient gov.QueryClient) ([]Proposal, map[string]string, error) {
	resp, err := queryClient.Proposals(ctx, &gov.QueryProposalsRequest{
		Pagination: &query.PageRequest{
			Limit: 1000,
		},
	})
	if err != nil {
		return nil, nil, err
	}
	allProposals := new([]Proposal)
	proposalMap := make(map[string]string)
	for _, proposal := range resp.GetProposals() {
		// deal with votes
		var votes gov.TallyResult
		// if vote is still ongoing, query the current tally
		if proposal.Status == 2 {
			resp, err := queryClient.TallyResult(ctx, &gov.QueryTallyResultRequest{
				ProposalId: proposal.ProposalId,
			})
			if err == nil {
				votes = resp.Tally
			}
		} else {
			votes = proposal.FinalTallyResult
		}

		// get proposal metadata
		title := ""
		description := ""
		metadata, err := GetProposalMetadata(proposal.Content)
		if err != nil {
			log.Log().Msgf("Error getting proposal metadata: %v", err)
		} else {
			title = metadata.Title
			description = metadata.Description
		}

		proposalResponse := Proposal{
			ProposalId:      proposal.ProposalId,
			TypeUrl:         proposal.Content.TypeUrl,
			Title:           title,
			Description:     description,
			Status:          proposal.Status.String(),
			FinalVote:       votes,
			SubmitTime:      proposal.SubmitTime,
			DepositEndTime:  proposal.DepositEndTime,
			TotalDeposit:    proposal.TotalDeposit,
			VotingStartTime: proposal.VotingStartTime,
			VotingEndTime:   proposal.VotingEndTime,
		}
		*allProposals = append(*allProposals, proposalResponse)
		proposalMap[strconv.Itoa(int(proposal.ProposalId))] = GeneralResultToString(proposalResponse)
	}
	return *allProposals, proposalMap, nil
}

// CSR
type CSR struct {
	// ID of the CSR
	Id uint64 `json:"id"`
	// all contracts under this csr id
	Contracts []string `json:"contracts"`
	// total number of transactions under this csr id
	Txs uint64 `json:"txs"`
	// The cumulative revenue for this CSR NFT -> represented as a big.Int
	Revenue string `json:"revenue"`
}

// get all CSRS
// will return full response string and mapping of nft id to response string
func GetCSRS(ctx context.Context, queryClient csr.QueryClient) ([]CSR, map[string]string, error) {
	resp, err := queryClient.CSRs(ctx, &csr.QueryCSRsRequest{Pagination: &query.PageRequest{
		Limit: 1000,
	}})
	if err != nil {
		return nil, nil, err
	}
	allCsrs := new([]CSR)
	csrMap := make(map[string]string)
	for _, csr := range resp.GetCsrs() {
		csrResponse := CSR{
			Id:        csr.GetId(),
			Contracts: csr.GetContracts(),
			Txs:       csr.GetTxs(),
			Revenue:   csr.Revenue.String(),
		}
		*allCsrs = append(*allCsrs, csrResponse)
		csrMap[strconv.Itoa(int(csr.GetId()))] = GeneralResultToString(csrResponse)
	}
	return *allCsrs, csrMap, nil
}

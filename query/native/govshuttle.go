package query

import (
	"context"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type Proposal struct {
	// proposalId defines the unique id of the proposal.
	ProposalId uint64 `json:"proposal_id"`
	// typeUrl indentifies the type of the proposal by a serialized protocol buffer message
	TypeUrl string `json:"type_url"`
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
func GetAllProposals(ctx context.Context, queryClient gov.QueryClient) ([]Proposal, map[string]string) {
	resp, err := queryClient.Proposals(ctx, &gov.QueryProposalsRequest{})
	CheckError(err)
	allProposals := new([]Proposal)
	proposalMap := make(map[string]string)
	for _, proposal := range resp.GetProposals() {
		proposalResponse := Proposal{
			ProposalId:      proposal.ProposalId,
			TypeUrl:         proposal.Content.TypeUrl,
			Status:          proposal.Status.String(),
			FinalVote:       proposal.FinalTallyResult,
			SubmitTime:      proposal.SubmitTime,
			DepositEndTime:  proposal.DepositEndTime,
			TotalDeposit:    proposal.TotalDeposit,
			VotingStartTime: proposal.VotingStartTime,
			VotingEndTime:   proposal.VotingEndTime,
		}
		*allProposals = append(*allProposals, proposalResponse)
		proposalMap[strconv.Itoa(int(proposal.ProposalId))] = GeneralResultToString(proposalResponse)
	}
	return *allProposals, proposalMap
}
// get user vote on active proposal
// will not work if proposal is not active
func GetUserVote(ctx context.Context, queryClient gov.QueryClient, proposalId uint64, ethAddress string) (gov.WeightedVoteOptions, error) {
	cantoAddress, err := EthToCantoAddress(ethAddress)
	if err != nil {
		return gov.WeightedVoteOptions{}, err
	}
	resp, err := queryClient.Vote(ctx, &gov.QueryVoteRequest{ProposalId: proposalId, Voter: cantoAddress})
	if err != nil {
		return gov.WeightedVoteOptions{}, err
	}
	return resp.GetVote().Options, nil
}

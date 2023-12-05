package queryengine

import (
	"fmt"

	erc20 "github.com/Canto-Network/Canto/v6/x/erc20/types"
	govshuttle "github.com/Canto-Network/Canto/v6/x/govshuttle/types"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	upgrade "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

type BasicMetadata struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetProposalMetadata(content *types1.Any) (BasicMetadata, error) {
	typeUrl := content.TypeUrl
	switch typeUrl {
	case "/canto.govshuttle.v1.LendingMarketProposal":
		var metadata govshuttle.LendingMarketProposal
		metadata.Unmarshal(content.Value)
		return BasicMetadata{
			Title:       metadata.GetTitle(),
			Description: metadata.GetDescription(),
		}, nil
	case "/canto.erc20.v1.RegisterCoinProposal":
		var metadata erc20.RegisterCoinProposal
		metadata.Unmarshal(content.Value)
		return BasicMetadata{
			Title:       metadata.GetTitle(),
			Description: metadata.GetDescription(),
		}, nil
	case "/cosmos.distribution.v1beta1.CommunityPoolSpendProposal":
		var metadata distribution.CommunityPoolSpendProposal
		metadata.Unmarshal(content.Value)
		return BasicMetadata{
			Title:       metadata.GetTitle(),
			Description: metadata.GetDescription(),
		}, nil
	case "/cosmos.gov.v1beta1.TextProposal":
		var metadata gov.TextProposal
		metadata.Unmarshal(content.Value)
		return BasicMetadata{
			Title:       metadata.GetTitle(),
			Description: metadata.GetDescription(),
		}, nil
	case "/cosmos.params.v1beta1.ParameterChangeProposal":
		var metadata params.ParameterChangeProposal
		metadata.Unmarshal(content.Value)
		return BasicMetadata{
			Title:       metadata.GetTitle(),
			Description: metadata.GetDescription(),
		}, nil
	case "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal":
		var metadata upgrade.SoftwareUpgradeProposal
		metadata.Unmarshal(content.Value)
		return BasicMetadata{
			Title:       metadata.GetTitle(),
			Description: metadata.GetDescription(),
		}, nil
	default:
		return BasicMetadata{}, fmt.Errorf("Proposal type: %s not found", typeUrl)
	}
}

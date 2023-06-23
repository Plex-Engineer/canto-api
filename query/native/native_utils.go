package query

import (
	"canto-api/config"
	"context"
	"encoding/json"

	evmos "github.com/evmos/ethermint/x/evm/types"

	inflation "github.com/Canto-Network/Canto/v6/x/inflation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// CalculateStakingAPR returns the APR for a all bonded tokens and mint provision for current epoch
func CalculateStakingAPR(pool staking.QueryPoolResponse, mintProvision inflation.QueryEpochMintProvisionResponse) sdk.Dec {
	//get bonded tokens from pool
	bondedTokens := pool.GetPool().BondedTokens
	//get mint provision amount from epoch (in acanto)
	mintProvisionAmount := mintProvision.GetEpochMintProvision().Amount

	//check if bonded tokens are zero so we don't divide by zero
	if bondedTokens.IsZero() {
		return sdk.NewDec(0)
	}

	//calculate apr (mint provision / bonded tokens) * 365 (days) * 100%
	return mintProvisionAmount.Mul(sdk.NewDec(36500)).QuoInt(bondedTokens)
}

func GeneralResultToString(results interface{}) string {
	ret, err := json.Marshal(results)
	if err != nil {
		return "QueryEngine::GeneralResultToString - " + err.Error()
	}
	return string(ret)
}

func EthToCantoAddress(ethAddress string) (string, error) {
	emvosQuery := evmos.NewQueryClient(config.GrpcClient)
	resp, err := emvosQuery.CosmosAccount(context.Background(), &evmos.QueryCosmosAccountRequest{Address: ethAddress})
	if err != nil {
		return "", err
	}
	// cantoAddress1 := new(evmos.QueryCosmosAccountResponse)
	// fmt.Println(cantoAddress1)
	// config.GrpcClient.Invoke(context.Background(), "/ethermint/evm/v1/cosmos_account", &evmos.QueryCosmosAccountRequest{Address: "0x8915da99B69e84DE6C97928d378D9887482C671c"}, cantoAddress1)

	// fmt.Println(cantoAddress1.GetCosmosAddress())

	// accountquery := accounts.NewQueryClient(config.GrpcClient)
	// account, err := accountquery.Account(context.Background(), &accounts.QueryAccountRequest{Address: "canto13y2a4xdkn6zdumyhj2xn0rvcsayzcecuhmwl8d"})
	// fmt.Print(account.Account.TypeUrl)

	// fmt.Println(sdk.GetConfig().GetBech32AccountAddrPrefix())
	// addressBytes := eth.HexToAddress(ethAddress).Bytes()
	// fmt.Println(addressBytes)
	// bech32Acc := sdk.AccAddress(addressBytes)
	// fmt.Println(bech32Acc.String())
	// bech32Prefix := strings.SplitN(bech32Acc.String(), "1", 2)[0]
	// fmt.Println(bech32Prefix)

	// //get canto address from eth address
	// cantoAddress, err := cantoUtils.GetcantoAddressFromBech32("canto13y2a4xdkn6zdumyhj2xn0rvcsayzcecuhmwl8d")
	// fmt.Println(cantoAddress.String())
	// addressss, err := sdk.GetFromBech32("canto13y2a4xdkn6zdumyhj2xn0rvcsayzcecuhmwl8d", "canto")
	// fmt.Println(addressss)
	// if err != nil {
	// 	return "", err
	// }
	return resp.GetCosmosAddress(), nil
}

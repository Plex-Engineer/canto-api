package config

import (
	"fmt"
	"io"
	"os"
)

func getCTokenContractCalls() []Contract {

	calls := []Contract{}

	//get cTokens
	for _, token := range TokensConfig.CTokens {
		tokenKey := token.Symbol
		calls = append(calls, Contract{
			Name:    token.Name,
			Address: token.Address,
			Keys: []string{
				"cTokens:" + tokenKey + ":getCash",
				"cTokens:" + tokenKey + ":exchangeRateStored",
				"cTokens:" + tokenKey + ":supplyRatePerBlock",
				"cTokens:" + tokenKey + ":borrowRatePerBlock",
			},
			Methods: []string{
				"getCash()(uint256)",
				"exchangeRateStored()(uint256)",
				"supplyRatePerBlock()(uint256)",
				"borrowRatePerBlock()(uint256)",
			},
			Args: [][]interface{}{
				{},
				{},
				{},
				{},
			},
		})
	}
	return calls
}

func getPairsContractsCalls() []Contract {
	calls := []Contract{}

	for _, pair := range TokensConfig.Pairs {
		pairKey := pair.Symbol
		calls = append(calls, Contract{
			Name:    pair.Name,
			Address: pair.Address,
			Keys: []string{
				"lpPairs:" + pairKey + ":reserves",
				"lpPairs:" + pairKey + ":tokens",
				"lpPairs:" + pairKey + ":stable",
			},
			Methods: []string{
				"getReserves()(uint256,uint256,uint256)",
				"tokens()(address,address)",
				"stable()(bool)",
			},
			Args: [][]interface{}{
				{},
				{},
				{},
			},
		})
	}
	return calls

}

// read jsoon to get all tokens and set them to TokensConfig
func getAllTokensFromJson(path string) TokensInfo {
	var TokensInfo TokensInfo

	tokensFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer tokensFile.Close()

	tokensByteValue, _ := io.ReadAll(tokensFile)
	tokens, err := UnmarshalTokens(tokensByteValue)

	if err != nil {
		fmt.Println(err)
	}

	TokensInfo.CTokens = tokens.CTokens
	TokensInfo.Pairs = tokens.Pairs
	return TokensInfo
}

func getAllFPI(path string) []Contract {
	calls := []Contract{}
	TokensConfig = getAllTokensFromJson(path)

	calls = append(calls, getCTokenContractCalls()...)
	calls = append(calls, getPairsContractsCalls()...)

	return calls
}

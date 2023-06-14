package config

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type Contract struct {
	Name    string
	Address string
	Keys    []string
	Methods []string
	Args    [][]interface{}
}

func getCTokenFromTokenAddress(cTokens []Token, keyName string, underlying string) (Token, error) {
	for _, token := range cTokens {
		// fmt.Println("token underlying: ", *token.Underlying, "pair underlying: ", underlying, "keyName: ", keyName)

		if *token.Underlying == underlying {
			return token, nil
		}
	}

	notFound := Token{}
	return notFound, errors.New(underlying + " token :  not found : " + keyName)
}

func getCTokensFromJson(tokens []Token, raw RawContracts) []Contract {

	calls := []Contract{}

	//get cTokens
	for _, token := range tokens {
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

		calls = append(calls, Contract{
			Name:    "Router",
			Address: raw["7700"].Router.Address,
			Keys: []string{
				"cTokens:" + tokenKey + ":price",
			},
			Methods: []string{
				"getUnderlyingPrice(address)(uint256)",
			},
			Args: [][]interface{}{
				{token.Address},
			},
		})

		calls = append(calls, Contract{
			Name:    "Comptroller",
			Address: raw["7700"].Comptroller.Address,
			Keys: []string{
				"cTokens:" + tokenKey + ":markets",
				"cTokens:" + tokenKey + ":supplySpeeds",
				"cTokens:" + tokenKey + ":borrowCaps",
			},
			Methods: []string{
				"markets(address)(bool, uint256, bool)",
				"compSupplySpeeds(address)(uint256)",
				"borrowCaps(address)(uint256)",
			},
			Args: [][]interface{}{
				{token.Address},
				{token.Address},
				{token.Address},
			},
		})
	}

	return calls

}

func getLPPairsFromJson(tokenData Tokens, contractData RawContracts) []Contract {
	calls := []Contract{}

	for _, pair := range tokenData.Pairs {
		pairKey := pair.Symbol

		// cTokenA, err := getCTokenFromTokenAddress(tokenData.CTokens, "tokenA", pair.TokenA)

		// if err != nil {
		// 	fmt.Println(err)
		// }

		// cTokenB, err := getCTokenFromTokenAddress(tokenData.CTokens, "tokenB", pair.TokenB)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		cPair, err := getCTokenFromTokenAddress(tokenData.CTokens, "cPair", pair.Address)

		if err != nil {
			fmt.Println(err)
		}

		calls = append(calls, Contract{
			Name:    "Router",
			Address: contractData[tokenData.ChainID].Router.Address,
			Keys: []string{
				"lpPairs:" + pairKey + ":reserves",
				"lpPairs:" + pairKey + ":price",
			},
			Methods: []string{
				"getReserves(address,address,bool)(uint256, uint256)",
				"getUnderlyingPrice(address)(uint256)",
			},
			Args: [][]interface{}{
				{pair.TokenA, pair.TokenB, pair.Stable},
				{cPair.Address},
			},
		})
	}
	return calls

}

func getAllContractsFromJson(isTestnet bool) []Contract {

	fileName := "tokens.json"

	if isTestnet {
		fileName = "testnet_tokens.json"
	}

	calls := []Contract{}

	tokensFile, err := os.Open("./config/jsons/" + fileName)

	if err != nil {
		fmt.Println(err)
	}

	contractsFile, err := os.Open("./config/jsons/contracts.json")

	if err != nil {
		fmt.Println(err)
	}
	defer tokensFile.Close()
	defer contractsFile.Close()

	tokensByteValue, _ := io.ReadAll(tokensFile)
	tokens, err := UnmarshalTokens(tokensByteValue)

	if err != nil {
		fmt.Println(err)
	}

	contractsByteValue, _ := io.ReadAll(contractsFile)
	contracts, err := UnmarshalRawContracts(contractsByteValue)

	if err != nil {
		fmt.Println(err)
	}

	calls = append(calls, getCTokensFromJson(tokens.CTokens, contracts)...)
	calls = append(calls, getLPPairsFromJson(tokens, contracts)...)

	return calls
}

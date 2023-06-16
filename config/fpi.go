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

		calls = append(calls, Contract{
			Name:    "Router",
			Address: ContractsConfig[TokensConfig.ChainID].Router.Address,
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
			Address: ContractsConfig[TokensConfig.ChainID].Comptroller.Address,
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

func getPairsContractsCalls() []Contract {
	calls := []Contract{}

	for _, pair := range TokensConfig.Pairs {
		pairKey := pair.Symbol

		// cTokenA, err := getCTokenFromTokenAddress(tokenData.CTokens, "tokenA", pair.TokenA)

		// if err != nil {
		// 	fmt.Println(err)
		// }

		// cTokenB, err := getCTokenFromTokenAddress(tokenData.CTokens, "tokenB", pair.TokenB)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		cPair, err := getCTokenFromTokenAddress(TokensConfig.CTokens, "cPair", pair.Address)

		if err != nil {
			fmt.Println(err)
		}

		calls = append(calls, Contract{
			Name:    "Router",
			Address: ContractsConfig[TokensConfig.ChainID].Router.Address,
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

func getAllTokensFromJson(isTestnet bool) TokensInfo {
	fileName := "tokens.json"

	if isTestnet {
		fileName = "testnet_tokens.json"
	}

	tokensFile, err := os.Open("./config/jsons/" + fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer tokensFile.Close()

	tokensByteValue, _ := io.ReadAll(tokensFile)
	tokens, err := UnmarshalTokens(tokensByteValue)

	if err != nil {
		fmt.Println(err)
	}

	return tokens
}

func getContractsDataFromJson() ContractsInfo {

	contractsFile, err := os.Open("./config/jsons/contracts.json")

	if err != nil {
		fmt.Println(err)
	}

	defer contractsFile.Close()

	contractsByteValue, _ := io.ReadAll(contractsFile)
	contracts, err := UnmarshalContractsConfig(contractsByteValue)

	if err != nil {
		fmt.Println(err)
	}
	return contracts
}

func getAllContracts() []Contract {
	calls := []Contract{}

	calls = append(calls, getCTokenContractCalls()...)
	calls = append(calls, getPairsContractsCalls()...)

	return calls
}

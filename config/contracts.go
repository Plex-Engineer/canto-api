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

func getCTokensFromJson(tokens []Token) []Contract {

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
			Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
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
			Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
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

func getLPPairsFromJson(data Tokens) []Contract {
	calls := []Contract{}

	for _, pair := range data.LpPairs {
		pairKey := pair.Symbol

		cTokenA, err := getCTokenFromTokenAddress(data.CTokens, "tokenA", pair.TokenA)

		if err != nil {
			fmt.Println(err)
		}

		cTokenB, err := getCTokenFromTokenAddress(data.CTokens, "tokenB", pair.TokenB)
		if err != nil {
			fmt.Println(err)
		}

		cPair, err := getCTokenFromTokenAddress(data.CTokens, "cPair", pair.Address)

		if err != nil {
			fmt.Println(err)
		}

		calls = append(calls, Contract{
			Name:    pair.Name,
			Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
			Keys: []string{
				"lpPairs:" + pairKey + ":reserves",
				"lpPairs:" + pairKey + ":tokenA",
				"lpPairs:" + pairKey + ":tokenB",
				"lpPairs:" + pairKey + ":pair",
			},
			Methods: []string{
				"getReserves(address,address,bool)(uint256, uint256)",
				"getUnderlyingPrice(address)(uint256)",
				"getUnderlyingPrice(address)(uint256)",
				"getUnderlyingPrice(address)(uint256)",
			},
			Args: [][]interface{}{
				{pair.TokenA, pair.TokenB, pair.Stable},
				{cTokenA.Address},
				{cTokenB.Address},
				{cPair.Address},
			},
		})
	}
	fmt.Println("calls: ", calls)
	return calls

}

func getAllContractsFromJson() []Contract {
	jsonFile, err := os.Open("./config/jsons/mainnet_tokens.json")

	calls := []Contract{}

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	// var result map[string]map[string]string interface{}
	// json.Unmarshal([]byte(byteValue), &result)
	res, err := UnmarshalTokens(byteValue)

	if err != nil {
		fmt.Println(err)
	}
	calls = append(calls, getCTokensFromJson(res.CTokens)...)
	calls = append(calls, getLPPairsFromJson(res)...)

	return calls
}

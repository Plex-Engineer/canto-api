package config

import (
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

func getContractsFromJsonFile() []Contract {

	jsonFile, err := os.Open("./config/jsons/mainnet_tokens.json")

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

	calls := []Contract{}

	//get cTokens
	for _, token := range res.CTokens {
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

	// //get tokens
	// for _, token := range res.Tokens {
	// 	tokenKey := token.Symbol
	// 	calls = append(calls, Contract{
	// 		Name:    token.Name,
	// 		Address: token.Address,
	// 		Keys: []string{
	// 			"tokens:" + tokenKey + ":balanceOf",
	// 			"tokens:" + tokenKey + ":totalSupply",
	// 		},
	// 		Methods: []string{
	// 			"balanceOf(address)(uint256)",
	// 			"totalSupply()(uint256)",
	// 		},
	// 		Args: [][]interface{}{
	// 			{},
	// 			{},
	// 		},
	// 	})
	// }

	return calls

}

package config

import (
	"fmt"
	"io/ioutil"
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
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// var result map[string]map[string]string interface{}
	// json.Unmarshal([]byte(byteValue), &result)
	res, err := UnmarshalTokens(byteValue)

	if err != nil {
		fmt.Println(err)
	}

	calls := []Contract{}

	for _, token := range res.CTokens {
		calls = append(calls, Contract{
			Name:    token.Name,
			Address: token.Address,
			Keys: []string{
				"cTokens:" + token.Address + ":getCash",
				"cTokens:" + token.Address + ":exchangeRateStored",
				"cTokens:" + token.Address + ":supplyRatePerBlock",
				"cTokens:" + token.Address + ":borrowRatePerBlock",
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

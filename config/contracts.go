package config

import (
	"encoding/json"
	"errors"
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

//returns a list of contract calls from the contracts.json

func getContractsFromJson(pathToContractsJsonFile string) ([]Contract, error) {
	var calls []Contract

	// Open the JSON file
	file, err := os.Open(pathToContractsJsonFile)
	if err != nil {
		return nil, errors.New("Config::getContractsFromJson - " + err.Error())
	}
	defer file.Close()

	// Read the file content
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("Config::getContractsFromJson - " + err.Error())
	}

	//Umarshall the data and store in calls
	unmarshallErr := json.Unmarshal(fileContent, &calls)

	if unmarshallErr != nil {
		fmt.Println("Error:", unmarshallErr)
		return nil, errors.New("Config::getContractsFromJson - " + err.Error())
	}
	return calls, nil

}

// var calls []Contract = []Contract{
// 	{
// 		Name:    "canto/note basev1pair",
// 		Address: "0x1D20635535307208919f0b67c3B2065965A85aA9",
// 		Methods: []string{
// 			"getReserves()(int256, int256, uint32)",
// 			"decimals()(uint8)",
// 		},
// 		Args: [][]interface{}{
// 			{},
// 			{},
// 		},
// 	},
// 	{
// 		Name:    "usdc erc20",
// 		Address: "0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd",
// 		Methods: []string{
// 			"balanceOf(address)(uint256)",
// 		},
// 		Args: [][]interface{}{
// 			{
// 				"0x7f1A3B16969DecE24d383980efba7cF5929464F8",
// 			},
// 		},
// 	},
// }

package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ContractAddresses struct {
	Mainnet ContractList `json:"mainnet"`
	Testnet ContractList `json:"testnet"`
}
type ContractList struct {
	Comptroller string `json:"comptroller"`
	Router      string `json:"router"`
	Reservoir   string `json:"reservoir"`
	MulticallV1 string `json:"multicallV1"`
	MulticallV2 string `json:"multicallV2"`
	MulticallV3 string `json:"multicallV3"`
}

// parses contract_addresses.json and returns contract addresses for all chains
func getContractAddressesFromJson(path string) ContractAddresses {
	var contractAddresses ContractAddresses

	// Open the JSON file
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	// Read the file content
	fileContent, _ := io.ReadAll(file)

	//Umarshall the data and store in contractAddresses
	err = json.Unmarshal(fileContent, &contractAddresses)

	if err != nil {
		fmt.Println(err)
	}

	return contractAddresses
}

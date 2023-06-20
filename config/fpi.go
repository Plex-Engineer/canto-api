package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// this function generates and returns required contract calls for all ctokens in TokensConfig
func getCTokenContractCalls() []Contract {
	// Declare and initialize a list of Contracts
	calls := []Contract{}

	// iterate over all ctokens in config and generate contract calls to query required ctoken data
	for _, token := range TokensConfig.CTokens {
		// get required ctoken data from its contract
		calls = append(calls, Contract{
			Name:    token.Symbol,
			Address: token.Address,
			Keys: []string{
				"cTokens:" + token.Symbol + ":cash",
				"cTokens:" + token.Symbol + ":exchangeRateStored",
				"cTokens:" + token.Symbol + ":supplyRatePerBlock",
				"cTokens:" + token.Symbol + ":borrowRatePerBlock",
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

		// getUnderlyingPrice data of ctoken from router contract
		calls = append(calls, Contract{
			Name:    token.Symbol + "pricefeed",
			Address: ContractAddressesConfig.Mainnet.Router,
			Keys: []string{
				"cTokens:" + token.Symbol + ":underlyingPrice",
			},
			Methods: []string{
				"getUnderlyingPrice(address)(uint256)",
			},
			Args: [][]interface{}{
				{token.Address},
			},
		})

		// get markets, compSupplySpeeds and borrowCaps data of ctoken from comptroller contract
		calls = append(calls, Contract{
			Name:    token.Symbol + "comptroller",
			Address: ContractAddressesConfig.Mainnet.Comptroller,
			Keys: []string{
				"cTokens:" + token.Symbol + ":markets",
				"cTokens:" + token.Symbol + ":compSupplySpeeds",
				"cTokens:" + token.Symbol + ":borrowCaps",
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

// this function returns the ctoken address of the token having address equal to underlyingAddress
func getCTokenAddress(underlyingAddress string) string {
	// Declare cNoteAddress which wll be returned when ctoken is not found for underlyingAddress
	var cNoteAddress string

	// iterate through ctokens config to get the ctoken address of the token with underlyingAddress
	for _, token := range TokensConfig.CTokens {
		//check if the current ctoken is cnote and store its address
		if token.Symbol == "cNOTE" {
			cNoteAddress = token.Address
		}

		// check if the current ctoken has given underlying address and return ctoken address if true
		if token.Underlying == underlyingAddress {
			return token.Address
		}
	}

	// return cnote address when token with underlyingAddress is not found in ctokens list
	return cNoteAddress
}

// this function generates and returns required contract calls for all pairs in TokensConfig
func getPairsContractsCalls() []Contract {
	// Declare and initialize a list of Contracts
	calls := []Contract{}

	// iterare over all the pairs in config and generate contract calls to query required pair data
	for _, pair := range TokensConfig.Pairs {
		// get required pair data from its contract
		calls = append(calls, Contract{
			Name:    pair.Symbol,
			Address: pair.Address,
			Keys: []string{
				"lpPairs:" + pair.Symbol + ":totalSupply",
			},
			Methods: []string{
				"totalSupply()(uint256)",
			},
			Args: [][]interface{}{
				{},
			},
		})

		// get reserves, underlying prices of tokenA, tokenB and Lp from router contract
		calls = append(calls, Contract{
			Name:    pair.Symbol + "pricefeed",
			Address: ContractAddressesConfig.Mainnet.Router,
			Keys: []string{
				"lpPairs:" + pair.Symbol + ":reserves",
				"lpPairs:" + pair.Symbol + ":underlyingPriceTokenA",
				"lpPairs:" + pair.Symbol + ":underlyingPriceTokenB",
				"lpPairs:" + pair.Symbol + ":underlyingPriceLp",
			},
			Methods: []string{
				"getReserves(address,address,bool)(uint256, uint256)",
				"getUnderlyingPrice(address)(uint256)",
				"getUnderlyingPrice(address)(uint256)",
				"getUnderlyingPrice(address)(uint256)",
			},
			Args: [][]interface{}{
				{pair.TokenA, pair.TokenB, pair.Stable},
				{getCTokenAddress(pair.TokenA)},
				{getCTokenAddress(pair.TokenB)},
				{getCTokenAddress(pair.Address)},
			},
		})
	}
	return calls

}

// parses tokens.json and returns tokens data
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

func getAllFPI(path string) []Contract {
	// Declare and initialize a list of Contracts
	calls := []Contract{}

	// add required contract calls for all cTokens
	calls = append(calls, getCTokenContractCalls()...)

	// add required contract calls for all pairs
	calls = append(calls, getPairsContractsCalls()...)

	return calls
}

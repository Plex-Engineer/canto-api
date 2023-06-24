// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    tokens, err := UnmarshalTokens(bytes)
//    bytes, err = tokens.Marshal()

package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type TokensInfo struct {
	Name        string   `json:"name"`
	Version     float64  `json:"version"`
	Keywords    []string `json:"keywords"`
	ChainID     string   `json:"chainid"`
	Comptroller string   `json:"comptroller"`
	Router      string   `json:"router"`
	Reservoir   string   `json:"reservoir"`
	MulticallV3 string   `json:"multicallV3"`
	CTokens     []Token  `json:"ctokens"`
	Tokens      []Token  `json:"tokens"`
	Pairs       []Pair   `json:"pairs"`
}

type Token struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Symbol     string `json:"symbol"`
	Decimals   int64  `json:"decimals"`
	Underlying string `json:"underlying,omitempty"`
	ChainID    string `json:"chainId"`
	LogoURI    string `json:"logoURI,omitempty"`
}

type Pair struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int64  `json:"decimals"`
	Address  string `json:"address"`
	Stable   bool   `json:"stable"`
	TokenA   string `json:"tokenA"`
	TokenB   string `json:"tokenB"`
	ChainID  string `json:"chainId"`
}

// parses tokens.json and returns tokens data
func getFPIFromJson(path string) TokensInfo {
	var TokensInfo TokensInfo

	tokensFile, err := os.Open(path)

	if err != nil {
		panic(fmt.Sprintf("Error opening tokens.json: %v", err))
	}

	defer tokensFile.Close()

	tokensByteValue, _ := io.ReadAll(tokensFile)
	err = json.Unmarshal(tokensByteValue, &TokensInfo)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling tokens.json: %v", err))
	}

	return TokensInfo
}

// this function returns the ctoken address of the token having address equal to underlyingAddress
func GetCTokenAddress(underlyingAddress string) (cTokenAddress string) {
	// iterate through ctokens config to get the ctoken address of the token with underlyingAddress
	for _, token := range FPIConfig.CTokens {
		// check if the current ctoken has given underlying address and return ctoken address if true
		if token.Underlying == underlyingAddress {
			cTokenAddress = token.Address
			return
		}
	}
	return
}

// this function returns the ctoken decimals of the token having address equal to underlyingAddress
func GetCTokenDecimals(underlyingAddress string) (decimals int64) {

	// iterate through ctokens config to get the ctoken decimals of the token with underlyingAddress
	for _, token := range FPIConfig.CTokens {
		// check if the current ctoken has given underlying address and return ctoken address if true
		if token.Underlying == underlyingAddress {
			decimals = token.Decimals
			return
		}
	}
	return
}

// get token data from tokens config using token address and return
func GetTokenData(address string) (result Token) {
	for _, token := range FPIConfig.Tokens {
		if token.Address == address {
			result = token
			return
		}
	}
	return
}

// get lp pair data (Address, Decimals, Token1, Token2, Stable, CDecimal, cLPaddress) from tokens config using pair symbol and return
func GetLpPairData(address string) (symbol string, decimals int64, token1 Token, token2 Token, stable bool, cDecimals int64, cLpAddress string) {
	for _, pair := range FPIConfig.Pairs {
		if pair.Address == address {
			symbol = pair.Symbol
			decimals = pair.Decimals
			token1 = GetTokenData(pair.TokenA)
			token2 = GetTokenData(pair.TokenB)
			stable = pair.Stable
			cDecimals = GetCTokenDecimals(address)
			cLpAddress = GetCTokenAddress(pair.Address)
			return
		}
	}
	return
}

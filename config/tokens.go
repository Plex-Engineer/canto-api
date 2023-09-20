// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    tokens, err := UnmarshalTokens(bytes)
//    bytes, err = tokens.Marshal()

package config

import (
	"encoding/json"
	"errors"
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

// this stores data of underlying token/pair of ctoken
type Underlying struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals int64  `json:"decimals"`
	LogoURI  string `json:"logoURI,omitempty"`
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
	LogoURI  string `json:"logoURI,omitempty"`
}

// parses tokens.json and returns tokens data
func getFPIFromJson(path string) (TokensInfo, error) {
	var TokensInfo TokensInfo

	tokensFile, err := os.Open(path)

	if err != nil {
		return TokensInfo, errors.New("Config::getFPIFromJson - " + err.Error())
	}
	defer tokensFile.Close()

	tokensByteValue, _ := io.ReadAll(tokensFile)
	err = json.Unmarshal(tokensByteValue, &TokensInfo)
	if err != nil {
		return TokensInfo, errors.New("Config::getFPIFromJson - " + err.Error())
	}

	return TokensInfo, nil
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

// this function returns the ctoken address of the token having given sumbol
func GetCTokenAddressBySymbol(symbol string) (cTokenAddress string) {
	// iterate through ctokens config to get the ctoken address of the token with given symbol
	for _, token := range FPIConfig.CTokens {
		// check if the current ctoken has given symbol
		if token.Symbol == symbol {
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

// get underying data of ctoken from tokens config using token/pair address
func GetUnderlyingData(address string) (result Underlying) {
	// iterate over all tokens and return data if token found
	for _, token := range FPIConfig.Tokens {
		if token.Address == address {
			result = Underlying{
				Name:     token.Name,
				Address:  token.Address,
				Decimals: token.Decimals,
				Symbol:   token.Symbol,
				LogoURI:  token.LogoURI,
			}
			return
		}
	}
	// iterate over all pairs and return data if pair found
	for _, pair := range FPIConfig.Pairs {
		if pair.Address == address {
			result = Underlying{
				Name:     pair.Name,
				Address:  pair.Address,
				Decimals: pair.Decimals,
				Symbol:   pair.Symbol,
				LogoURI:  pair.LogoURI,
			}
			return
		}
	}
	return
}

// get lp pair data (Address, Decimals, Token1, Token2, Stable, CDecimal, cLPaddress) from tokens config using pair symbol and return
func GetLpPairData(address string) (symbol string, decimals int64, token1 Token, token2 Token, stable bool, cDecimals int64, cLpAddress string, logoUri string) {
	for _, pair := range FPIConfig.Pairs {
		if pair.Address == address {
			symbol = pair.Symbol
			decimals = pair.Decimals
			token1 = GetTokenData(pair.TokenA)
			token2 = GetTokenData(pair.TokenB)
			stable = pair.Stable
			cDecimals = GetCTokenDecimals(address)
			cLpAddress = GetCTokenAddress(pair.Address)
			logoUri = pair.LogoURI
			return
		}
	}
	return
}

// get ctoken data (Symbol, Name, Decimals, Underlying) from tokens config using cToken address
func GetCTokenData(address string) (symbol string, name string, decimals int64, underlying Underlying) {
	for _, cToken := range FPIConfig.CTokens {
		if cToken.Address == address {
			symbol = cToken.Symbol
			name = cToken.Name
			decimals = cToken.Decimals
			underlying = GetUnderlyingData(cToken.Underlying)
			return
		}
	}
	return
}

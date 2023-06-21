// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    tokens, err := UnmarshalTokens(bytes)
//    bytes, err = tokens.Marshal()

package config

import "encoding/json"

func UnmarshalTokens(data []byte) (TokensInfo, error) {
	var r TokensInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TokensInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TokensInfo struct {
	Name     string   `json:"name"`
	Version  float64  `json:"version"`
	Keywords []string `json:"keywords"`
	ChainID  string   `json:"chainId"`
	CTokens  []Token  `json:"cTokens"`
	Tokens   []Token  `json:"tokens"`
	Pairs    []Pair   `json:"pairs"`
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

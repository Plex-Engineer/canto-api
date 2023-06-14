// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    tokens, err := UnmarshalTokens(bytes)
//    bytes, err = tokens.Marshal()

package config

import "encoding/json"

func UnmarshalTokens(data []byte) (Tokens, error) {
	var r Tokens
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Tokens) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Tokens struct {
	Name     string   `json:"name"`
	Version  float64  `json:"version"`
	Keywords []string `json:"keywords"`
	ChainID  string   `json:"chainId"`
	CTokens  []Token  `json:"cTokens"`
	Tokens   []Token  `json:"tokens"`
	Pairs    []Pair   `json:"pairs"`
}

type Token struct {
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	Symbol     string  `json:"symbol"`
	Decimals   int64   `json:"decimals"`
	Underlying *string `json:"underlying,omitempty"`
	ChainID    string  `json:"chainId"`
	LogoURI    *string `json:"logoURI,omitempty"`
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

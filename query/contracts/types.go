package query

import "canto-api/config"

type ProcessedPair struct {
	Address     string       `json:"address"`
	Symbol      string       `json:"symbol"`
	Decimals    int64        `json:"decimals"`
	Token1      config.Token `json:"token1"`
	Token2      config.Token `json:"token2"`
	Stable      bool         `json:"stable"`
	CDecimal    int64        `json:"cDecimals"`
	CLpAddress  string       `json:"cLpAddress"`
	TotalSupply string       `json:"totalSupply"`
	Tvl         string       `json:"tvl"`
	Ratio       string       `json:"ratio"`
	AToB        bool         `json:"aTob"`
	Price1      string       `json:"price1"`
	Price2      string       `json:"price2"`
	LpPrice     string       `json:"lpPrice"`
	Reserve1    string       `json:"reserve1"`
	Reserve2    string       `json:"reserve2"`
}

type ProcessedCToken struct {
	Address       string       `json:"address"`
	Symbol        string       `json:"symbol"`
	Decimals      int64        `json:"decimals"`
	Underlying    config.Token `json:"underlying"`
	Price         string       `json:"price"`
	TotalSupply   string       `json:"totalSupply"`
	ExchangeRate  string       `json:"exchangeRate"`
	CTokenAddress string       `json:"cTokenAddress"`
}

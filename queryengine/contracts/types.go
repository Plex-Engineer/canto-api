package queryengine

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
	LogoURI     string       `json:"logoURI,omitempty"`
}
type ProcessedCToken struct {
	Address               string            `json:"address"`
	Symbol                string            `json:"symbol"`
	Name                  string            `json:"name"`
	Decimals              int64             `json:"decimals"`
	Underlying            config.Underlying `json:"underlying"`
	Cash                  string            `json:"cash"`
	ExchangeRate          string            `json:"exchangeRate"`
	CollateralFactor      string            `json:"collateralFactor"`
	Price                 string            `json:"price"`
	BorrowCap             string            `json:"borrowCap"`
	IsListed              bool              `json:"isListed"`
	Liquidity             string            `json:"liquidity"`
	SupplyApy             string            `json:"supplyApy"`
	SupplyApr             string            `json:"supplyApr"`
	BorrowApy             string            `json:"borrowApy"`
	BorrowApr             string            `json:"borrowApr"`
	DistApy               string            `json:"distApy"`
	DistApr               string            `json:"distApr"`
	CompSupplyState       string            `json:"compSupplyState"`
	UnderlyingTotalSupply string            `json:"underlyingTotalSupply"`
}

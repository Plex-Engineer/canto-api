package requestengine

type Pairs struct {
	BlockNumber string
	PairsStruct []ProcessedPair
}

type ProcessedPair struct {
	Address     string
	Symbol      string
	Decimals    uint64
	Token1      Token
	Token0      Token
	Stable      bool
	CDecimals   uint64
	CLPAddress  string
	TotalSupply string
	Tvl         string
	Ratio       string
	ATob        bool
	Price1      string
	Price2      string
	LpPrice     string
	Reserve1    string
	Reserve2    string
}

type Token struct {
	Name     string
	Address  string
	Symbol   string
	Decimals uint64
	ChainId  uint64
	LogoURI  string
}

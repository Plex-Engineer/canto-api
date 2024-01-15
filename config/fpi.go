package config

func getAllFPI() []Contract {
	// Declare and initialize a list of Contracts
	calls := []Contract{}

	// add required contract calls for all cTokens
	calls = append(calls, getCTokenContractCalls()...)

	// add required contract calls for all pairs
	calls = append(calls, getPairsContractsCalls()...)

	return calls
}

// this function generates and returns required contract calls for all ctokens in TokensConfig
func getCTokenContractCalls() []Contract {
	// Declare and initialize a list of Contracts
	calls := []Contract{}

	// iterate over all ctokens in config and generate contract calls to query required ctoken data
	for _, token := range FPIConfig.CTokens {
		// get required ctoken data from its contract
		calls = append(calls, Contract{
			Name:    token.Symbol,
			Address: token.Address,
			Keys: []string{
				"cTokens:" + token.Address + ":cash",
				"cTokens:" + token.Address + ":exchangeRateStored",
				"cTokens:" + token.Address + ":supplyRatePerBlock",
				"cTokens:" + token.Address + ":borrowRatePerBlock",
				"cTokens:" + token.Address + ":totalSupply",
			},
			Methods: []string{
				"getCash()(uint256)",
				"exchangeRateStored()(uint256)",
				"supplyRatePerBlock()(uint256)",
				"borrowRatePerBlock()(uint256)",
				"totalSupply()(uint256)",
			},
			Args: [][]interface{}{
				{},
				{},
				{},
				{},
				{},
			},
		})

		// getUnderlyingPrice data of ctoken from price oracle contract
		calls = append(calls, Contract{
			Name:    token.Symbol + "pricefeed",
			Address: FPIConfig.PriceOracle,
			Keys: []string{
				"cTokens:" + token.Address + ":underlyingPrice",
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
			Address: FPIConfig.Comptroller,
			Keys: []string{
				"cTokens:" + token.Address + ":markets",
				"cTokens:" + token.Address + ":compSupplySpeeds",
				"cTokens:" + token.Address + ":compBorrowSpeeds",
				"cTokens:" + token.Address + ":borrowCaps",
				"cTokens:" + token.Address + ":compSupplyState",
			},
			Methods: []string{
				"markets(address)(bool, uint256, bool)",
				"compSupplySpeeds(address)(uint256)",
				"compBorrowSpeeds(address)(uint256)",
				"borrowCaps(address)(uint256)",
				"compSupplyState(address)(uint256)",
			},
			Args: [][]interface{}{
				{token.Address},
				{token.Address},
				{token.Address},
				{token.Address},
				{token.Address},
			},
		})

		// get total supply of underlying too, will help with calculating liquidity
		calls = append(calls, Contract{
			Name:    token.Symbol + "underlyingSupply",
			Address: token.Underlying,
			Keys: []string{
				"cTokens:" + token.Address + ":underlyingSupply",
			},
			Methods: []string{
				"totalSupply()(uint256)",
			},
			Args: [][]interface{}{
				{},
			},
		})

		// check token tags and add calls accordingly
		if token.Tags != nil {
			for _, tag := range token.Tags {
				if tag == "hashnote" {
					calls = append(calls, Contract{
						Name:    token.Symbol + "latestRoundDetails",
						Address: "0x1d18c02bC80b1921255E71cF2939C03258d75470",
						Keys: []string{
							"cTokens:" + token.Address + ":latestRoundDetails",
						},
						Methods: []string{
							"latestRoundDetails()(uint80,uint256,uint256,uint256,uint256)",
						},
						Args: [][]interface{}{
							{},
						},
					})
				}
			}
		}
	}

	return calls
}

// this function generates and returns required contract calls for all pairs in TokensConfig
func getPairsContractsCalls() []Contract {
	// Declare and initialize a list of Contracts
	calls := []Contract{}

	// iterare over all the pairs in config and generate contract calls to query required pair data
	for _, pair := range FPIConfig.Pairs {
		// get required pair data from its contract
		calls = append(calls, Contract{
			Name:    pair.Symbol,
			Address: pair.Address,
			Keys: []string{
				"lpPairs:" + pair.Address + ":reserves",
				"lpPairs:" + pair.Address + ":tokens",
				"lpPairs:" + pair.Address + ":stable",
				"lpPairs:" + pair.Address + ":totalSupply",
			},
			Methods: []string{
				"getReserves()(uint256,uint256,uint256)",
				"tokens()(address,address)",
				"stable()(bool)",
				"totalSupply()(uint256)",
			},
			Args: [][]interface{}{
				{},
				{},
				{},
				{},
			},
		})
		// get reserves
		calls = append(calls, Contract{
			Name:    pair.Symbol + "reserves",
			Address: FPIConfig.Router,
			Keys: []string{
				"lpPairs:" + pair.Address + ":reserves",
			},
			Methods: []string{
				"getReserves(address,address,bool)(uint256, uint256)",
			},
			Args: [][]interface{}{
				{pair.TokenA, pair.TokenB, pair.Stable},
			},
		})

		// get underlying prices of tokenA, tokenB and Lp from price oracle contract
		calls = append(calls, Contract{
			Name:    pair.Symbol + "pricefeed",
			Address: FPIConfig.PriceOracle,
			Keys: []string{
				"lpPairs:" + pair.Address + ":underlyingPriceTokenA",
				"lpPairs:" + pair.Address + ":underlyingPriceTokenB",
				"lpPairs:" + pair.Address + ":underlyingPriceLp",
			},
			Methods: []string{
				"getUnderlyingPrice(address)(uint256)",
				"getUnderlyingPrice(address)(uint256)",
				"getUnderlyingPrice(address)(uint256)",
			},
			Args: [][]interface{}{
				{GetCTokenAddress(pair.TokenA)},
				{GetCTokenAddress(pair.TokenB)},
				{GetCTokenAddress(pair.Address)},
			},
		})
	}
	return calls

}

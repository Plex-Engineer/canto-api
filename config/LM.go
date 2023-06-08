package config

import (
	"fmt"
	"strings"
)

var TEST_C_TOKENS = []CToken{
	CTOKENS["CantoTestnet"]["CCANTO"],
	CTOKENS["CantoTestnet"]["CNOTE"],
	CTOKENS["CantoTestnet"]["CETH"],
	CTOKENS["CantoTestnet"]["CATOM"],
	CTOKENS["CantoTestnet"]["CUSDC"],
	CTOKENS["CantoTestnet"]["CUSDT"],
	CTOKENS["CantoTestnet"]["CCantoNote"],
	CTOKENS["CantoTestnet"]["CCantoAtom"],
	CTOKENS["CantoTestnet"]["CNoteUSDC"],
	CTOKENS["CantoTestnet"]["CNoteUSDT"],
	CTOKENS["CantoTestnet"]["CCantoETH"],
}

var MAIN_C_TOKENS = []CToken{
	CTOKENS["CantoMainnet"]["CCANTO"],
	CTOKENS["CantoMainnet"]["CNOTE"],
	CTOKENS["CantoMainnet"]["CETH"],
	CTOKENS["CantoMainnet"]["CATOM"],
	CTOKENS["CantoMainnet"]["CUSDC"],
	CTOKENS["CantoMainnet"]["CUSDT"],
	CTOKENS["CantoMainnet"]["CCantoNote"],
	CTOKENS["CantoMainnet"]["CCantoAtom"],
	CTOKENS["CantoMainnet"]["CNoteUSDC"],
	CTOKENS["CantoMainnet"]["CNoteUSDT"],
	CTOKENS["CantoMainnet"]["CCantoETH"],
}

// LM_CALLS_MAIN_C_TOKENS

func getMainnetLendingMarketCalls() []Contract {

	calls := make([]Contract, 0)
	for _, token := range MAIN_C_TOKENS {

		calls = append(calls, Contract{
			// Name:    "ccanto token",
			Name:    fmt.Sprintf("%s token", strings.ToLower(token.Name)),
			Address: token.Address,
			Keys: []string{
				fmt.Sprintf("cash:%s", strings.ToLower(token.Name)),
				fmt.Sprintf("exchangerate:%s", strings.ToLower(token.Name)),
				fmt.Sprintf("supplyrate:%s", strings.ToLower(token.Name)),
				fmt.Sprintf("borrowrate:%s", strings.ToLower(token.Name)),
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

		calls = append(calls, Contract{
			// Name:    "ccanto pricefeed",
			Name:    fmt.Sprintf("%s pricefeed", strings.ToLower(token.Name)),
			Address: ADDRESSES.CantoMainnet.PriceFeed,
			Keys: []string{
				fmt.Sprintf("pricefeed:%s", strings.ToLower(token.Name)),
			},
			Methods: []string{
				"getUnderlyingPrice(address)(uint256)",
			},
			Args: [][]interface{}{
				{token.Address},
			},
		})

		calls = append(calls, Contract{
			// Name:    "ccanto comptroller",
			Name:    fmt.Sprintf("%s comptroller", strings.ToLower(token.Name)),
			Address: ADDRESSES.CantoMainnet.Comptroller,
			Keys: []string{
				fmt.Sprintf("markets:%s", strings.ToLower(token.Name)),
				fmt.Sprintf("supplyspeeds:%s", strings.ToLower(token.Name)),
				fmt.Sprintf("borrowcaps:%s", strings.ToLower(token.Name)),
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

// LM_CALLS_TEST_C_TOKENS

// func getTestnetLendingMarketCalls() []Contract {

// 	calls := make([]Contract, 0)
// 	for _, token := range TEST_C_TOKENS {

// 		calls = append(calls, Contract{
// 			Name:    "Cash",
// 			Address: token.Address,
// 			Methods: []string{
// 				"getCash()(uint256)",
// 			},
// 			Args: [][]interface{}{
// 				{},
// 			},
// 		})
// 		calls = append(calls, Contract{
// 			Name:    "ExchangeRate",
// 			Address: token.Address,
// 			Methods: []string{
// 				"exchangeRateStored()(uint256)",
// 			},
// 			Args: [][]interface{}{
// 				{},
// 			},
// 		})
// 		calls = append(calls, Contract{
// 			Name:    "SupplyRate",
// 			Address: token.Address,
// 			Methods: []string{
// 				"supplyRatePerBlock()(uint256)",
// 			},
// 			Args: [][]interface{}{
// 				{},
// 			},
// 		})
// 		calls = append(calls, Contract{
// 			Name:    "Borrow Rate",
// 			Address: token.Address,
// 			Methods: []string{
// 				"borrowRatePerBlock()(uint256)",
// 			},
// 			Args: [][]interface{}{
// 				{},
// 			},
// 		})
// 		calls = append(calls, Contract{
// 			Name:    "Markets",
// 			Address: ADDRESSES.Testnet.Comptroller,
// 			Methods: []string{
// 				"markets(address)(bool, uint256, bool)",
// 			},
// 			Args: [][]interface{}{
// 				{token.Address},
// 			},
// 		})
// 		calls = append(calls, Contract{
// 			Name:    "UnderlyingPrice",
// 			Address: ADDRESSES.Testnet.PriceFeed,
// 			Methods: []string{
// 				"getUnderlyingPrice(address)(uint256)",
// 			},
// 			Args: [][]interface{}{
// 				{token.Address},
// 			},
// 		})
// 		calls = append(calls, Contract{
// 			Name:    "SupplySpeed",
// 			Address: ADDRESSES.Testnet.Comptroller,
// 			Methods: []string{
// 				"compSupplySpeeds(address)(uint256)",
// 			},
// 			Args: [][]interface{}{
// 				{token.Address},
// 			},
// 		})
// 		calls = append(calls, Contract{
// 			Name:    "BorrowCaps",
// 			Address: ADDRESSES.Testnet.Comptroller,
// 			Methods: []string{
// 				"borrowCaps(address)(uint256)",
// 			},
// 			Args: [][]interface{}{
// 				{token.Address},
// 			},
// 		})
// 	}
// 	return calls
// }
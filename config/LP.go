package config

import (
	"fmt"
	"strings"
)

type Pair struct {
	address    string
	cLPaddress string
	token1     Token
	token2     Token
	decimals   int
	cDecimals  int
	stable     bool
}

var TEST_PAIRS = []Pair{
	{
		address:    ADDRESSES.Testnet.CantoNoteLP,
		cLPaddress: ADDRESSES.Testnet.CCantoNoteLP,
		token1: Token{
			Symbol:   "CANTO",
			Name:     "WCanto",
			Decimals: DECIMALS["CANTO"],
			Address:  ADDRESSES.Testnet.WCANTO,
			IsERC20:  true,
			IsLP:     false,
		},
		token2: Token{
			Symbol:   "NOTE",
			Name:     "Note",
			Decimals: DECIMALS["Note"],
			Address:  ADDRESSES.Testnet.Note,
			IsERC20:  true,
			IsLP:     false,
		},
		decimals:  TOKENS["CantoTestnet"]["CantoNote"].Decimals,
		cDecimals: CTOKENS["CantoTestnet"]["CCantoNote"].Decimals,
		stable:    false,
	},
	{
		address:    ADDRESSES.Testnet.CantoETHLP,
		cLPaddress: ADDRESSES.Testnet.CCantoETHLP,
		token1: Token{
			Symbol:   "CANTO",
			Name:     "WCanto",
			Decimals: DECIMALS["CANTO"],
			Address:  ADDRESSES.Testnet.WCANTO,
			IsERC20:  true,
			IsLP:     false,
		},
		token2: Token{
			Symbol:   "ETH",
			Name:     "ETH",
			Decimals: DECIMALS["ETH"],
			Address:  ADDRESSES.Testnet.ETH,
			IsERC20:  true,
			IsLP:     false,
		},
		decimals:  TOKENS["CantoTestnet"]["CantoETH"].Decimals,
		cDecimals: CTOKENS["CantoTestnet"]["CCantoETH"].Decimals,
		stable:    false,
	},
	{
		address:    ADDRESSES.Testnet.CantoAtomLP,
		cLPaddress: ADDRESSES.Testnet.CCantoAtomLP,
		token1: Token{
			Symbol:   "CANTO",
			Name:     "WCanto",
			Decimals: DECIMALS["CANTO"],
			Address:  ADDRESSES.Testnet.WCANTO,
			IsERC20:  true,
			IsLP:     false,
		},
		token2: Token{
			Symbol:   "ATOM",
			Name:     "ATOM",
			Decimals: DECIMALS["ATOM"],
			Address:  ADDRESSES.Testnet.ATOM,
			IsERC20:  true,
			IsLP:     false,
		},
		decimals:  TOKENS["CantoTestnet"]["CantoAtom"].Decimals,
		cDecimals: CTOKENS["CantoTestnet"]["CCantoAtom"].Decimals,
		stable:    false,
	},
	{
		address:    ADDRESSES.Testnet.NoteUSDCLP,
		cLPaddress: ADDRESSES.Testnet.CNoteUSDCLP,
		token1: Token{
			Symbol:   "NOTE",
			Name:     "Note",
			Decimals: DECIMALS["Note"],
			Address:  ADDRESSES.Testnet.Note,
			IsERC20:  true,
			IsLP:     false,
		},
		token2: Token{
			Symbol:   "USDC",
			Name:     "USDC",
			Decimals: DECIMALS["USDC"],
			Address:  ADDRESSES.Testnet.USDC,
			IsERC20:  true,
			IsLP:     false,
		},
		decimals:  TOKENS["CantoTestnet"]["NoteUSDC"].Decimals,
		cDecimals: CTOKENS["CantoTestnet"]["CNoteUSDC"].Decimals,
		stable:    true,
	},
	{
		address:    ADDRESSES.Testnet.NoteUSDTLP,
		cLPaddress: ADDRESSES.Testnet.CNoteUSDTLP,
		token1: Token{
			Symbol:   "NOTE",
			Name:     "Note",
			Decimals: DECIMALS["Note"],
			Address:  ADDRESSES.Testnet.Note,
			IsERC20:  true,
			IsLP:     false,
		},
		token2: Token{
			Symbol:   "USDT",
			Name:     "USDT",
			Decimals: DECIMALS["USDT"],
			Address:  ADDRESSES.Testnet.USDT,
			IsERC20:  true,
			IsLP:     false,
		},
		decimals:  TOKENS["CantoTestnet"]["NoteUSDT"].Decimals,
		cDecimals: CTOKENS["CantoTestnet"]["CNoteUSDT"].Decimals,
		stable:    true,
	},
}

var MAIN_PAIRS = []Pair{
	{
		address:    ADDRESSES.CantoMainnet.CantoNoteLP,
		cLPaddress: ADDRESSES.CantoMainnet.CCantoNoteLP,
		token1: Token{
			Symbol:   "CANTO",
			Name:     "WCanto",
			Decimals: DECIMALS["CANTO"],
			Address:  ADDRESSES.CantoMainnet.WCANTO,
			IsERC20:  true,
			IsLP:     false,
		},
		token2: Token{
			Symbol:   "NOTE",
			Name:     "Note",
			Decimals: DECIMALS["Note"],
			Address:  ADDRESSES.CantoMainnet.Note,
			IsERC20:  true,
			IsLP:     false,
		},
		decimals:  TOKENS["CantoMainnet"]["CantoNote"].Decimals,
		cDecimals: CTOKENS["CantoMainnet"]["CCantoNote"].Decimals,
		stable:    false,
	},
	{
		address:    ADDRESSES.CantoMainnet.CantoETHLP,
		cLPaddress: ADDRESSES.CantoMainnet.CCantoETHLP,
		token1: Token{
			Symbol:   "CANTO",
			Name:     "WCanto",
			Decimals: DECIMALS["CANTO"],
			Address:  ADDRESSES.CantoMainnet.WCANTO,
			IsERC20:  true,
			IsLP:     false,
		},
		token2: Token{
			Symbol:   "ETH",
			Name:     "ETH",
			Decimals: DECIMALS["ETH"],
			Address:  ADDRESSES.CantoMainnet.ETH,
			IsERC20:  true,
			IsLP:     false,
		},
		decimals:  TOKENS["CantoMainnet"]["CantoETH"].Decimals,
		cDecimals: CTOKENS["CantoMainnet"]["CCantoETH"].Decimals,
		stable:    false,
	},
	{
		address:    ADDRESSES.CantoMainnet.CantoAtomLP,
		cLPaddress: ADDRESSES.CantoMainnet.CCantoAtomLP,
		token1: Token{
			Symbol:   "CANTO",
			Name:     "WCanto",
			Decimals: DECIMALS["CANTO"],
			Address:  ADDRESSES.CantoMainnet.WCANTO,
			IsERC20:  true,
			IsLP:     false,
		},
		token2: Token{
			Symbol:   "ATOM",
			Name:     "ATOM",
			Decimals: DECIMALS["ATOM"],
			Address:  ADDRESSES.CantoMainnet.ATOM,
			IsERC20:  true,
			IsLP:     false,
		},
		decimals:  TOKENS["CantoMainnet"]["CantoAtom"].Decimals,
		cDecimals: CTOKENS["CantoMainnet"]["CCantoAtom"].Decimals,
		stable:    false,
	},
	{
		address:    ADDRESSES.CantoMainnet.NoteUSDCLP,
		cLPaddress: ADDRESSES.CantoMainnet.CNoteUSDCLP,
		token1: Token{
			Symbol:   "NOTE",
			Name:     "Note",
			Decimals: DECIMALS["Note"],
			Address:  ADDRESSES.CantoMainnet.Note,
			IsERC20:  true,
			IsLP:     false,
		},
		token2: Token{
			Symbol:   "USDC",
			Name:     "USDC",
			Decimals: DECIMALS["USDC"],
			Address:  ADDRESSES.CantoMainnet.USDC,
			IsERC20:  true,
			IsLP:     false,
		},
		decimals:  TOKENS["CantoMainnet"]["NoteUSDC"].Decimals,
		cDecimals: CTOKENS["CantoMainnet"]["CNoteUSDC"].Decimals,
		stable:    true,
	},
	{
		address:    ADDRESSES.CantoMainnet.NoteUSDTLP,
		cLPaddress: ADDRESSES.CantoMainnet.CNoteUSDTLP,
		token1: Token{
			Symbol:   "NOTE",
			Name:     "Note",
			Decimals: DECIMALS["Note"],
			Address:  ADDRESSES.CantoMainnet.Note,
			IsERC20:  true,
			IsLP:     false,
		},
		token2: Token{
			Symbol:   "USDT",
			Name:     "USDT",
			Decimals: DECIMALS["USDT"],
			Address:  ADDRESSES.CantoMainnet.USDT,
			IsERC20:  true,
			IsLP:     false,
		},
		decimals:  TOKENS["CantoMainnet"]["NoteUSDT"].Decimals,
		cDecimals: CTOKENS["CantoMainnet"]["CNoteUSDT"].Decimals,
		stable:    true,
	},
}

// LP_CALLS_MAIN_PAIRS

func getMainnetLiquidityPoolCalls() []Contract {

	cTokenAddress := func(underlyingAddress string) string {
		for _, token := range CTOKENS["CantoMainnet"] {
			if token.Underlying.Address == underlyingAddress {
				return token.Address
			}
		}
		return CTOKENS["CantoMainnet"]["NOTE"].Address
	}

	calls := make([]Contract, 0)
	for _, pair := range MAIN_PAIRS {

		calls = append(calls, Contract{
			// Name:    "usdc/note pair",
			Name:    fmt.Sprintf("%s/%s pair", strings.ToLower(pair.token1.Name), strings.ToLower(pair.token2.Name)),
			Address: pair.address,
			Keys: []string{
				fmt.Sprintf("lptotalsupply:%s/%s", strings.ToLower(pair.token1.Name), strings.ToLower(pair.token2.Name)),
			},
			Methods: []string{
				"totalSupply()(uint256)",
			},
			Args: [][]interface{}{
				{},
			},
		})

		calls = append(calls, Contract{
			// Name:    "usdc/note pricefeed",
			Name:    fmt.Sprintf("%s/%s pricefeed", strings.ToLower(pair.token1.Name), strings.ToLower(pair.token2.Name)),
			Address: ADDRESSES.CantoMainnet.PriceFeed,
			Keys: []string{
				fmt.Sprintf("reserves:%s/%s", strings.ToLower(pair.token1.Name), strings.ToLower(pair.token2.Name)),
				fmt.Sprintf("underlyingprice:%s", strings.ToLower(pair.token1.Name)),
				fmt.Sprintf("underlyingprice:%s", strings.ToLower(pair.token2.Name)),
				fmt.Sprintf("underlyingprice:%s:%s", strings.ToLower(pair.token1.Name), strings.ToLower(pair.token2.Name)),
			},
			Methods: []string{
				"getReserves(address,address,bool)(uint256)",
				"getUnderlyingPrice(address)(uint256)",
				"getUnderlyingPrice(address)(uint256)",
				"getUnderlyingPrice(address)(uint256)",
			},
			Args: [][]interface{}{
				{pair.token1.Address, pair.token2.Address, pair.stable},
				{cTokenAddress(pair.token1.Address)},
				{cTokenAddress(pair.token2.Address)},
				{pair.cLPaddress},
			},
		})

	}
	return calls
}

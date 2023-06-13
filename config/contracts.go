package config

type Contract struct {
	Name    string
	Address string
	Keys    []string
	Methods []string
	Args    [][]interface{}
}

var calls []Contract = []Contract{
	{
		Name:    "ccanto token",
		Address: "0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
		Keys: []string{
			"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488:cash",
			"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488:exchangerate",
			"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488:supplyrate",
			"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488:borrowrate",
		},
		Methods: []string{
			"getReserves()(int256, int256, uint32)",
			"decimals()(uint8)",
		},
		Args: [][]interface{}{
			{},
			{},
			{},
		},
	},
	{
		Name:    "note/usdc pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getReserves(address,address,bool)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
				"0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd",
				true,
			},
			{
				"0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			},
			{
				"0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
			},
			{
				"0xD6a97e43FC885A83E97d599796458A331E580800",
			},
		},
	},
	{
		Name:    "note/usdt pair",
		Address: "0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833",
		Methods: []string{
			"totalSupply()(uint256)",
		},
		Args: [][]interface{}{
			{},
		},
	},
	{
		Name:    "usdc erc20",
		Address: "0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd",
		Methods: []string{
			"balanceOf(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
			},
		},
	},
}

// var calls []Contract = []Contract{
// 	{
// 		Name:    "canto/note basev1pair",
// 		Address: "0x1D20635535307208919f0b67c3B2065965A85aA9",
// 		Methods: []string{
// 			"getReserves()(int256, int256, uint32)",
// 			"decimals()(uint8)",
// 		},
// 		Args: [][]interface{}{
// 			{},
// 			{},
// 		},
// 	},
// 	{
// 		Name:    "usdc erc20",
// 		Address: "0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd",
// 		Methods: []string{
// 			"balanceOf(address)(uint256)",
// 		},
// 		Args: [][]interface{}{
// 			{
// 				"0x7f1A3B16969DecE24d383980efba7cF5929464F8",
// 			},
// 		},
// 	},
// }

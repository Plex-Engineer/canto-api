package config

type Contract struct {
	Name    string
	Address string
	Methods []string
	Args    [][]interface{}
}

// type Contract struct {
// 	Name    string
// 	Address string
// 	Keys    []string
// 	Methods []string
// 	Args    [][]interface{}
// }

func getAllContractCalls() []Contract {
	calls := make([]Contract, 0)
	mainnetLmCalls := getMainnetLendingMarketCalls()
	mainnetLPCalls := getMainnetLiquidityPoolCalls()

	calls = append(calls, mainnetLmCalls...)
	calls = append(calls, mainnetLPCalls...)

	// fmt.Println("Contracts-----------", calls)

	// file, err := os.Create("output.json")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return calls
	// }
	// defer file.Close()

	// encoder := json.NewEncoder(file)
	// encoder.SetIndent("", "  ")

	// err = encoder.Encode(calls)
	// if err != nil {
	// 	fmt.Println("Error encoding JSON:", err)
	// 	return calls
	// }

	// fmt.Println("Data written to output.json")

	return calls
}

var calls = []Contract{
	{
		Name:    "ccanto token",
		Address: "0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
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
	},
	{
		Name:    "ccanto pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			},
		},
	},
	{
		Name:    "ccanto comptroller",
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Methods: []string{
			"markets(address)(bool, uint256, bool)",
			"compSupplySpeeds(address)(uint256)",
			"borrowCaps(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			},
			{
				"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			},
			{
				"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			},
		},
	},
	{
		Name:    "cnote token",
		Address: "0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
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
	},
	{
		Name:    "cnote pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			},
		},
	},
	{
		Name:    "cnote comptroller",
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Methods: []string{
			"markets(address)(bool, uint256, bool)",
			"compSupplySpeeds(address)(uint256)",
			"borrowCaps(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			},
			{
				"0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			},
			{
				"0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			},
		},
	},
	{
		Name:    "ceth token",
		Address: "0x830b9849e7d79b92408a86a557e7baaacbec6030",
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
	},
	{
		Name:    "ceth pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x830b9849e7d79b92408a86a557e7baaacbec6030",
			},
		},
	},
	{
		Name:    "ceth comptroller",
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Methods: []string{
			"markets(address)(bool, uint256, bool)",
			"compSupplySpeeds(address)(uint256)",
			"borrowCaps(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x830b9849e7d79b92408a86a557e7baaacbec6030",
			},
			{
				"0x830b9849e7d79b92408a86a557e7baaacbec6030",
			},
			{
				"0x830b9849e7d79b92408a86a557e7baaacbec6030",
			},
		},
	},
	{
		Name:    "catom token",
		Address: "0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
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
	},
	{
		Name:    "catom pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
			},
		},
	},
	{
		Name:    "catom comptroller",
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Methods: []string{
			"markets(address)(bool, uint256, bool)",
			"compSupplySpeeds(address)(uint256)",
			"borrowCaps(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
			},
			{
				"0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
			},
			{
				"0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
			},
		},
	},
	{
		Name:    "cusdc token",
		Address: "0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
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
	},
	{
		Name:    "cusdc pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
			},
		},
	},
	{
		Name:    "cusdc comptroller",
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Methods: []string{
			"markets(address)(bool, uint256, bool)",
			"compSupplySpeeds(address)(uint256)",
			"borrowCaps(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
			},
			{
				"0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
			},
			{
				"0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
			},
		},
	},
	{
		Name:    "cusdt token",
		Address: "0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
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
	},
	{
		Name:    "cusdt pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
			},
		},
	},
	{
		Name:    "cusdt comptroller",
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Methods: []string{
			"markets(address)(bool, uint256, bool)",
			"compSupplySpeeds(address)(uint256)",
			"borrowCaps(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
			},
			{
				"0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
			},
			{
				"0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
			},
		},
	},
	{
		Name:    "ccantonotelp token",
		Address: "0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
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
	},
	{
		Name:    "ccantonotelp pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
			},
		},
	},
	{
		Name:    "ccantonotelp comptroller",
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Methods: []string{
			"markets(address)(bool, uint256, bool)",
			"compSupplySpeeds(address)(uint256)",
			"borrowCaps(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
			},
			{
				"0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
			},
			{
				"0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
			},
		},
	},
	{
		Name:    "ccantoatomlp token",
		Address: "0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
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
	},
	{
		Name:    "ccantoatomlp pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
			},
		},
	},
	{
		Name:    "ccantoatomlp comptroller",
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Methods: []string{
			"markets(address)(bool, uint256, bool)",
			"compSupplySpeeds(address)(uint256)",
			"borrowCaps(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
			},
			{
				"0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
			},
			{
				"0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
			},
		},
	},
	{
		Name:    "cnoteusdclp token",
		Address: "0xD6a97e43FC885A83E97d599796458A331E580800",
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
	},
	{
		Name:    "cnoteusdclp pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xD6a97e43FC885A83E97d599796458A331E580800",
			},
		},
	},
	{
		Name:    "cnoteusdclp comptroller",
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Methods: []string{
			"markets(address)(bool, uint256, bool)",
			"compSupplySpeeds(address)(uint256)",
			"borrowCaps(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xD6a97e43FC885A83E97d599796458A331E580800",
			},
			{
				"0xD6a97e43FC885A83E97d599796458A331E580800",
			},
			{
				"0xD6a97e43FC885A83E97d599796458A331E580800",
			},
		},
	},
	{
		Name:    "cnoteusdtlp token",
		Address: "0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
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
	},
	{
		Name:    "cnoteusdtlp pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
			},
		},
	},
	{
		Name:    "cnoteusdtlp comptroller",
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Methods: []string{
			"markets(address)(bool, uint256, bool)",
			"compSupplySpeeds(address)(uint256)",
			"borrowCaps(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
			},
			{
				"0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
			},
			{
				"0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
			},
		},
	},
	{
		Name:    "ccantoethlp token",
		Address: "0xb49A395B39A0b410675406bEE7bD06330CB503E3",
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
	},
	{
		Name:    "ccantoethlp pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xb49A395B39A0b410675406bEE7bD06330CB503E3",
			},
		},
	},
	{
		Name:    "ccantoethlp comptroller",
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Methods: []string{
			"markets(address)(bool, uint256, bool)",
			"compSupplySpeeds(address)(uint256)",
			"borrowCaps(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0xb49A395B39A0b410675406bEE7bD06330CB503E3",
			},
			{
				"0xb49A395B39A0b410675406bEE7bD06330CB503E3",
			},
			{
				"0xb49A395B39A0b410675406bEE7bD06330CB503E3",
			},
		},
	},
	{
		Name:    "wcanto/note pair",
		Address: "0x1D20635535307208919f0b67c3B2065965A85aA9",
		Methods: []string{
			"totalSupply()(uint256)",
		},
		Args: [][]interface{}{
			{},
		},
	},
	{
		Name:    "wcanto/note pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getReserves(address,address,bool)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x826551890Dc65655a0Aceca109aB11AbDbD7a07B",
				"0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
				false,
			},
			{
				"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			},
			{
				"0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			},
			{
				"0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
			},
		},
	},
	{
		Name:    "wcanto/eth pair",
		Address: "0x216400ba362d8FCE640085755e47075109718C8B",
		Methods: []string{
			"totalSupply()(uint256)",
		},
		Args: [][]interface{}{
			{},
		},
	},
	{
		Name:    "wcanto/eth pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getReserves(address,address,bool)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x826551890Dc65655a0Aceca109aB11AbDbD7a07B",
				"0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687",
				false,
			},
			{
				"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			},
			{
				"0x830b9849e7d79b92408a86a557e7baaacbec6030",
			},
			{
				"0xb49A395B39A0b410675406bEE7bD06330CB503E3",
			},
		},
	},
	{
		Name:    "wcanto/atom pair",
		Address: "0x30838619C55B787BafC3A4cD9aEa851C1cfB7b19",
		Methods: []string{
			"totalSupply()(uint256)",
		},
		Args: [][]interface{}{
			{},
		},
	},
	{
		Name:    "wcanto/atom pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Methods: []string{
			"getReserves(address,address,bool)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
			"getUnderlyingPrice(address)(uint256)",
		},
		Args: [][]interface{}{
			{
				"0x826551890Dc65655a0Aceca109aB11AbDbD7a07B",
				"0xecEEEfCEE421D8062EF8d6b4D814efe4dc898265",
				false,
			},
			{
				"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			},
			{
				"0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
			},
			{
				"0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
			},
		},
	},
	{
		Name:    "note/usdc pair",
		Address: "0x9571997a66D63958e1B3De9647C22bD6b9e7228c",
		Methods: []string{
			"totalSupply()(uint256)",
		},
		Args: [][]interface{}{
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
		Name:    "note/usdt pricefeed",
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
				"0xd567B3d7B8FE3C79a1AD8dA978812cfC4Fa05e75",
				true,
			},
			{
				"0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			},
			{
				"0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
			},
			{
				"0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
			},
		},
	},
}

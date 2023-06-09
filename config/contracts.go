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
			"cash:ccanto",
			"exchangerate:ccanto",
			"supplyrate:ccanto",
			"borrowrate:ccanto",
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
	},
	{
		Name:    "ccanto pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Keys: []string{
			"pricefeed:ccanto",
		},
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
		Keys: []string{
			"markets:ccanto",
			"supplyspeeds:ccanto",
			"borrowcaps:ccanto",
		},
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
		Keys: []string{
			"cash:cnote",
			"exchangerate:cnote",
			"supplyrate:cnote",
			"borrowrate:cnote",
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
	},
	{
		Name:    "cnote pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Keys: []string{
			"pricefeed:cnote",
		},
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
		Keys: []string{
			"markets:cnote",
			"supplyspeeds:cnote",
			"borrowcaps:cnote",
		},
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
		Keys: []string{
			"cash:ceth",
			"exchangerate:ceth",
			"supplyrate:ceth",
			"borrowrate:ceth",
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
	},
	{
		Name:    "ceth pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Keys: []string{
			"pricefeed:ceth",
		},
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
		Keys: []string{
			"markets:ceth",
			"supplyspeeds:ceth",
			"borrowcaps:ceth",
		},
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
		Keys: []string{
			"cash:catom",
			"exchangerate:catom",
			"supplyrate:catom",
			"borrowrate:catom",
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
	},
	{
		Name:    "catom pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Keys: []string{
			"pricefeed:catom",
		},
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
		Keys: []string{
			"markets:catom",
			"supplyspeeds:catom",
			"borrowcaps:catom",
		},
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
		Keys: []string{
			"cash:cusdc",
			"exchangerate:cusdc",
			"supplyrate:cusdc",
			"borrowrate:cusdc",
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
	},
	{
		Name:    "cusdc pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Keys: []string{
			"pricefeed:cusdc",
		},
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
		Keys: []string{
			"markets:cusdc",
			"supplyspeeds:cusdc",
			"borrowcaps:cusdc",
		},
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
		Keys: []string{
			"cash:cusdt",
			"exchangerate:cusdt",
			"supplyrate:cusdt",
			"borrowrate:cusdt",
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
	},
	{
		Name:    "cusdt pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Keys: []string{
			"pricefeed:cusdt",
		},
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
		Keys: []string{
			"markets:cusdt",
			"supplyspeeds:cusdt",
			"borrowcaps:cusdt",
		},
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
		Keys: []string{
			"cash:ccantonotelp",
			"exchangerate:ccantonotelp",
			"supplyrate:ccantonotelp",
			"borrowrate:ccantonotelp",
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
	},
	{
		Name:    "ccantonotelp pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Keys: []string{
			"pricefeed:ccantonotelp",
		},
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
		Keys: []string{
			"markets:ccantonotelp",
			"supplyspeeds:ccantonotelp",
			"borrowcaps:ccantonotelp",
		},
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
		Keys: []string{
			"cash:ccantoatomlp",
			"exchangerate:ccantoatomlp",
			"supplyrate:ccantoatomlp",
			"borrowrate:ccantoatomlp",
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
	},
	{
		Name:    "ccantoatomlp pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Keys: []string{
			"pricefeed:ccantoatomlp",
		},
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
		Keys: []string{
			"markets:ccantoatomlp",
			"supplyspeeds:ccantoatomlp",
			"borrowcaps:ccantoatomlp",
		},
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
		Keys: []string{
			"cash:cnoteusdclp",
			"exchangerate:cnoteusdclp",
			"supplyrate:cnoteusdclp",
			"borrowrate:cnoteusdclp",
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
	},
	{
		Name:    "cnoteusdclp pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Keys: []string{
			"pricefeed:cnoteusdclp",
		},
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
		Keys: []string{
			"markets:cnoteusdclp",
			"supplyspeeds:cnoteusdclp",
			"borrowcaps:cnoteusdclp",
		},
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
		Keys: []string{
			"cash:cnoteusdtlp",
			"exchangerate:cnoteusdtlp",
			"supplyrate:cnoteusdtlp",
			"borrowrate:cnoteusdtlp",
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
	},
	{
		Name:    "cnoteusdtlp pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Keys: []string{
			"pricefeed:cnoteusdtlp",
		},
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
		Keys: []string{
			"markets:cnoteusdtlp",
			"supplyspeeds:cnoteusdtlp",
			"borrowcaps:cnoteusdtlp",
		},
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
		Keys: []string{
			"cash:ccantoethlp",
			"exchangerate:ccantoethlp",
			"supplyrate:ccantoethlp",
			"borrowrate:ccantoethlp",
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
	},
	{
		Name:    "ccantoethlp pricefeed",
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Keys: []string{
			"pricefeed:ccantoethlp",
		},
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
		Keys: []string{
			"markets:ccantoethlp",
			"supplyspeeds:ccantoethlp",
			"borrowcaps:ccantoethlp",
		},
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
		Keys: []string{
			"lptotalsupply:wcanto/note",
		},
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
		Keys: []string{
			"reserves:wcanto/note",
			"underlyingprice:wcanto",
			"underlyingprice:note",
			"underlyingprice:wcanto:note",
		},
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
		Keys: []string{
			"lptotalsupply:wcanto/eth",
		},
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
		Keys: []string{
			"reserves:wcanto/eth",
			"underlyingprice:wcanto",
			"underlyingprice:eth",
			"underlyingprice:wcanto:eth",
		},
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
		Keys: []string{
			"lptotalsupply:wcanto/atom",
		},
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
		Keys: []string{
			"reserves:wcanto/atom",
			"underlyingprice:wcanto",
			"underlyingprice:atom",
			"underlyingprice:wcanto:atom",
		},
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
		Keys: []string{
			"lptotalsupply:note/usdc",
		},
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
		Keys: []string{
			"reserves:note/usdc",
			"underlyingprice:note",
			"underlyingprice:usdc",
			"underlyingprice:note:usdc",
		},
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
		Keys: []string{
			"lptotalsupply:note/usdt",
		},
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
		Keys: []string{
			"reserves:note/usdt",
			"underlyingprice:note",
			"underlyingprice:usdt",
			"underlyingprice:note:usdt",
		},
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

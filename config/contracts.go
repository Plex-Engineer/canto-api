package config

type Contract struct {
	Address string
	Names   []string
	Methods []string
	Args    [][]interface{}
}

var calls []Contract = []Contract{
	{
		Address: "0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
		Names: []string{
			"Cash/0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			"ExchangeRate/0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			"SupplyRate/0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			"BorrowRate/0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
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
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Names: []string{
			"Markets/0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			"SupplySpeed/0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
			"BorrowCaps/0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
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
		Address: "0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
		Names: []string{
			"Cash/0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			"ExchangeRate/0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			"SupplyRate/0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			"BorrowRate/0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
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
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Names: []string{
			"Markets/0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			"SupplySpeed/0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			"BorrowCaps/0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
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
		Address: "0x830b9849e7d79b92408a86a557e7baaacbec6030",
		Names: []string{
			"Cash/0x830b9849e7d79b92408a86a557e7baaacbec6030",
			"ExchangeRate/0x830b9849e7d79b92408a86a557e7baaacbec6030",
			"SupplyRate/0x830b9849e7d79b92408a86a557e7baaacbec6030",
			"BorrowRate/0x830b9849e7d79b92408a86a557e7baaacbec6030",
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
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Names: []string{
			"Markets/0x830b9849e7d79b92408a86a557e7baaacbec6030",
			"SupplySpeed/0x830b9849e7d79b92408a86a557e7baaacbec6030",
			"BorrowCaps/0x830b9849e7d79b92408a86a557e7baaacbec6030",
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
		Address: "0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
		Names: []string{
			"Cash/0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
			"ExchangeRate/0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
			"SupplyRate/0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
			"BorrowRate/0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
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
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Names: []string{
			"Markets/0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
			"SupplySpeed/0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
			"BorrowCaps/0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
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
		Address: "0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
		Names: []string{
			"Cash/0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
			"ExchangeRate/0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
			"SupplyRate/0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
			"BorrowRate/0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
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
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Names: []string{
			"Markets/0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
			"SupplySpeed/0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
			"BorrowCaps/0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
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
		Address: "0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
		Names: []string{
			"Cash/0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
			"ExchangeRate/0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
			"SupplyRate/0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
			"BorrowRate/0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
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
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Names: []string{
			"Markets/0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
			"SupplySpeed/0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
			"BorrowCaps/0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
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
		Address: "0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
		Names: []string{
			"Cash/0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
			"ExchangeRate/0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
			"SupplyRate/0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
			"BorrowRate/0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
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
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Names: []string{
			"Markets/0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
			"SupplySpeed/0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
			"BorrowCaps/0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
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
		Address: "0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
		Names: []string{
			"Cash/0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
			"ExchangeRate/0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
			"SupplyRate/0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
			"BorrowRate/0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
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
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Names: []string{
			"Markets/0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
			"SupplySpeed/0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
			"BorrowCaps/0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
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
		Address: "0xD6a97e43FC885A83E97d599796458A331E580800",
		Names: []string{
			"Cash/0xD6a97e43FC885A83E97d599796458A331E580800",
			"ExchangeRate/0xD6a97e43FC885A83E97d599796458A331E580800",
			"SupplyRate/0xD6a97e43FC885A83E97d599796458A331E580800",
			"BorrowRate/0xD6a97e43FC885A83E97d599796458A331E580800",
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
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Names: []string{
			"Markets/0xD6a97e43FC885A83E97d599796458A331E580800",
			"SupplySpeed/0xD6a97e43FC885A83E97d599796458A331E580800",
			"BorrowCaps/0xD6a97e43FC885A83E97d599796458A331E580800",
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
		Address: "0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
		Names: []string{
			"Cash/0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
			"ExchangeRate/0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
			"SupplyRate/0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
			"BorrowRate/0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
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
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Names: []string{
			"Markets/0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
			"SupplySpeed/0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
			"BorrowCaps/0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
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
		Address: "0xb49A395B39A0b410675406bEE7bD06330CB503E3",
		Names: []string{
			"Cash/0xb49A395B39A0b410675406bEE7bD06330CB503E3",
			"ExchangeRate/0xb49A395B39A0b410675406bEE7bD06330CB503E3",
			"SupplyRate/0xb49A395B39A0b410675406bEE7bD06330CB503E3",
			"BorrowRate/0xb49A395B39A0b410675406bEE7bD06330CB503E3",
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
		Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Names: []string{
			"Markets/0xb49A395B39A0b410675406bEE7bD06330CB503E3",
			"SupplySpeed/0xb49A395B39A0b410675406bEE7bD06330CB503E3",
			"BorrowCaps/0xb49A395B39A0b410675406bEE7bD06330CB503E3",
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
		Address: "0x1D20635535307208919f0b67c3B2065965A85aA9",
		Names: []string{
			"TotalSupply/0x1D20635535307208919f0b67c3B2065965A85aA9",
		},
		Methods: []string{
			"totalSupply()(uint256)",
		},
		Args: [][]interface{}{
			{},
		},
	},
	{
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Names: []string{
			"Reserves/0x1D20635535307208919f0b67c3B2065965A85aA9",
			"TokenOneUnderlyingPrice/0x1D20635535307208919f0b67c3B2065965A85aA9",
			"TokenOneUnderlyingPrice/0x1D20635535307208919f0b67c3B2065965A85aA9",
			"LPUnderlyingPrice/0x1D20635535307208919f0b67c3B2065965A85aA9",
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
		Address: "0x216400ba362d8FCE640085755e47075109718C8B",
		Names: []string{
			"TotalSupply/0x216400ba362d8FCE640085755e47075109718C8B",
		},
		Methods: []string{
			"totalSupply()(uint256)",
		},
		Args: [][]interface{}{
			{},
		},
	},
	{
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Names: []string{
			"Reserves/0x216400ba362d8FCE640085755e47075109718C8B",
			"TokenOneUnderlyingPrice/0x216400ba362d8FCE640085755e47075109718C8B",
			"TokenOneUnderlyingPrice/0x216400ba362d8FCE640085755e47075109718C8B",
			"LPUnderlyingPrice/0x216400ba362d8FCE640085755e47075109718C8B",
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
		Address: "0x30838619C55B787BafC3A4cD9aEa851C1cfB7b19",
		Names: []string{
			"TotalSupply/0x30838619C55B787BafC3A4cD9aEa851C1cfB7b19",
		},
		Methods: []string{
			"totalSupply()(uint256)",
		},
		Args: [][]interface{}{
			{},
		},
	},
	{
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Names: []string{
			"Reserves/0x30838619C55B787BafC3A4cD9aEa851C1cfB7b19",
			"TokenOneUnderlyingPrice/0x30838619C55B787BafC3A4cD9aEa851C1cfB7b19",
			"TokenOneUnderlyingPrice/0x30838619C55B787BafC3A4cD9aEa851C1cfB7b19",
			"LPUnderlyingPrice/0x30838619C55B787BafC3A4cD9aEa851C1cfB7b19",
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
		Address: "0x9571997a66D63958e1B3De9647C22bD6b9e7228c",
		Names: []string{
			"TotalSupply/0x9571997a66D63958e1B3De9647C22bD6b9e7228c",
		},
		Methods: []string{
			"totalSupply()(uint256)",
		},
		Args: [][]interface{}{
			{},
		},
	},
	{
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Names: []string{
			"Reserves/0x9571997a66D63958e1B3De9647C22bD6b9e7228c",
			"TokenOneUnderlyingPrice/0x9571997a66D63958e1B3De9647C22bD6b9e7228c",
			"TokenOneUnderlyingPrice/0x9571997a66D63958e1B3De9647C22bD6b9e7228c",
			"LPUnderlyingPrice/0x9571997a66D63958e1B3De9647C22bD6b9e7228c",
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
		Address: "0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833",
		Names: []string{
			"TotalSupply/0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833",
		},
		Methods: []string{
			"totalSupply()(uint256)",
		},
		Args: [][]interface{}{
			{},
		},
	},
	{
		Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
		Names: []string{
			"Reserves/0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833",
			"TokenOneUnderlyingPrice/0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833",
			"TokenOneUnderlyingPrice/0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833",
			"LPUnderlyingPrice/0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833",
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

package config

type Contract struct {
	Name    string
	Address string
	Methods []string
	Args    [][]interface{}
}

var calls []Contract = []Contract{
	{
		Name:    "canto/note basev1pair",
		Address: "0x1D20635535307208919f0b67c3B2065965A85aA9",
		Methods: []string{
			"getReserves()(int256, int256, uint32)",
			"decimals()(uint8)",
		},
		Args: [][]interface{}{
			{},
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

package config

type Addresses struct {
	CantoMainnet      CantoMainnetAddresses
	Testnet           TestnetAddresses
	GravityBridgeTest GravityBridgeTestAddresses
	ETHMainnet        ETHMainnetAddresses
}

type CantoMainnetAddresses struct {
	CantoNoteLP  string
	CCantoNoteLP string
	CantoAtomLP  string
	CCantoAtomLP string
	NoteUSDTLP   string
	CNoteUSDCLP  string
	NoteUSDCLP   string
	CNoteUSDTLP  string
	CantoETHLP   string
	CCantoETHLP  string
	Comptroller  string
	Note         string
	USDC         string
	USDT         string
	ATOM         string
	ETH          string
	CNote        string
	CUSDC        string
	CUSDT        string
	CATOM        string
	CETH         string
	CCanto       string
	WCANTO       string
	PriceFeed    string
}

type TestnetAddresses struct {
	CantoNoteLP  string
	CCantoNoteLP string
	CantoAtomLP  string
	CCantoAtomLP string
	NoteUSDCLP   string
	CNoteUSDCLP  string
	NoteUSDTLP   string
	CNoteUSDTLP  string
	CantoETHLP   string
	CCantoETHLP  string
	Comptroller  string
	Note         string
	USDC         string
	USDT         string
	ATOM         string
	ETH          string
	CNote        string
	CUSDC        string
	CUSDT        string
	CATOM        string
	CETH         string
	CCanto       string
	WCANTO       string
	PriceFeed    string
}

type GravityBridgeTestAddresses struct {
	E2H           string
	BYE           string
	MAX           string
	GravityBridge string
}

type ETHMainnetAddresses struct {
	USDT          string
	USDC          string
	WETH          string
	GravityBridge string
}

var ADDRESSES = Addresses{
	CantoMainnet: CantoMainnetAddresses{
		CantoNoteLP:  "0x1D20635535307208919f0b67c3B2065965A85aA9",
		CCantoNoteLP: "0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
		CantoAtomLP:  "0x30838619C55B787BafC3A4cD9aEa851C1cfB7b19",
		CCantoAtomLP: "0xC0D6574b2fe71eED8Cd305df0DA2323237322557",
		NoteUSDTLP:   "0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833",
		CNoteUSDCLP:  "0xD6a97e43FC885A83E97d599796458A331E580800",
		NoteUSDCLP:   "0x9571997a66D63958e1B3De9647C22bD6b9e7228c",
		CNoteUSDTLP:  "0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
		CantoETHLP:   "0x216400ba362d8FCE640085755e47075109718C8B",
		CCantoETHLP:  "0xb49A395B39A0b410675406bEE7bD06330CB503E3",
		Comptroller:  "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
		Note:         "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
		USDC:         "0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd",
		USDT:         "0xd567B3d7B8FE3C79a1AD8dA978812cfC4Fa05e75",
		ATOM:         "0xecEEEfCEE421D8062EF8d6b4D814efe4dc898265",
		ETH:          "0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687",
		CNote:        "0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
		CUSDC:        "0xdE59F060D7ee2b612E7360E6C1B97c4d8289Ca2e",
		CUSDT:        "0x6b46ba92d7e94FfA658698764f5b8dfD537315A9",
		CATOM:        "0x617383F201076e7cE0f6E625D1a983b3D1bd277A",
		CETH:         "0x830b9849e7d79b92408a86a557e7baaacbec6030",
		CCanto:       "0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
		WCANTO:       "0x826551890Dc65655a0Aceca109aB11AbDbD7a07B",
		PriceFeed:    "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
	},
	Testnet: TestnetAddresses{
		CantoNoteLP:  "0x395E6ce7891f32278375Ff551B8ed61dF5579fE3",
		CCantoNoteLP: "0x2fd02CDB9Be9428d4eC2Ae969e52710601E219C6",
		CantoAtomLP:  "0x2bDF6c1302efc3c03D9C95f6fb5a4826A6bD964b",
		CCantoAtomLP: "0x4777Dc2b41f1f2Bd878205A61c1eA2609749928C",
		NoteUSDCLP:   "0x2db30A39Ec88247da8906506DB8E9dd933A5C775",
		CNoteUSDCLP:  "0xB2C5512a8A70835Cb9aBe830C9e61FBDdcd1dC81",
		NoteUSDTLP:   "0x252631e22e1ECc2fc0E811562605ed624B7E31d5",
		CNoteUSDTLP:  "0xBeD263484AEDFD449eE1ed8f0b4799192026E190",
		CantoETHLP:   "0x905D3d7F4C892d535160f1E2BA55f23Cd306718b",
		CCantoETHLP:  "0xf301c9d5804Fab3dd207ef75f78509db6393f37F",
		Comptroller:  "0x9514c07bC6e80B652e4264E64f589C59065C231f",
		Note:         "0x03F734Bd9847575fDbE9bEaDDf9C166F880B5E5f",
		USDC:         "0xc51534568489f47949A828C8e3BF68463bdF3566",
		USDT:         "0x4fC30060226c45D8948718C95a78dFB237e88b40",
		ATOM:         "0x40E41DC5845619E7Ba73957449b31DFbfB9678b2",
		ETH:          "0xCa03230E7FB13456326a234443aAd111AC96410A",
		CNote:        "0x04E52476d318CdF739C38BD41A922787D441900c",
		CUSDC:        "0x9160c5760a540cAfA24F90102cAA14C50497d5b7",
		CUSDT:        "0x3BEe0A8209e6F8c5c743F21e0cA99F2cb7806d1F",
		CATOM:        "0xaC8eF6d224438A2a69830596a73d1F2Ae26A304a",
		CETH:         "0x47b2e93B96E41175A013a0E3b0bE8b243FCD1B8C",
		CCanto:       "0x6F087493d6BcAE27C15766d8E657db6e98Bb27A9",
		WCANTO:       "0x7dBe1B86D5F48A54a1a86860EEFd2e2f332AA3b1",
		PriceFeed:    "0x35Bb487aEF22B92cCD53Bf16a1f83F3Dd7685DB0",
	},
	GravityBridgeTest: GravityBridgeTestAddresses{
		E2H:           "0x0AcA8e2C7b824c5a543eC400D95299A576827685",
		BYE:           "0x3DF223f25A8BddAEcA6b302fb4eB1F1712dA5E66",
		MAX:           "0x01900d3CCd634b8135D7F07A68e88090dF006A6b",
		GravityBridge: "0x5C5EdA728Ea22C415c63b60D9c717d8eB10e0C67",
	},
	ETHMainnet: ETHMainnetAddresses{
		USDT:          "0xdac17f958d2ee523a2206206994597c13d831ec7",
		USDC:          "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		WETH:          "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		GravityBridge: "0xCd7c84678994A44d9B96547b30CE1bF465D068ea",
	},
}

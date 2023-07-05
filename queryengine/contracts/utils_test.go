package queryengine

import (
	"canto-api/multicall"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestGetCallData(t *testing.T) {
	type args struct {
		vcs multicall.ViewCalls
	}
	tests := []struct {
		name    string
		args    args
		want    []multicall.Multicall3Call
		wantErr bool
	}{
		{
			name: "function with no argument",
			args: args{
				vcs: multicall.ViewCalls{multicall.NewViewCall(
					"decimals:0x0000",
					"0x0000",
					"decimals()",
					[]interface{}{},
				)},
			},
			want: []multicall.Multicall3Call{
				{
					Target:   common.HexToAddress("0x0000"),
					CallData: decodeHelper("313ce567"),
				},
			},
			wantErr: false,
		},
		{
			name: "function with one argument",
			args: args{
				vcs: multicall.ViewCalls{multicall.NewViewCall(
					"balanceof:0x0000",
					"0x0000",
					"balanceOf(address)(uint256)",
					[]interface{}{
						"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
					},
				)},
			},
			want: []multicall.Multicall3Call{
				{
					Target:   common.HexToAddress("0x0000"),
					CallData: decodeHelper("70a0823100000000000000000000000071C7656EC7ab88b098defB751B7401B5f6d8976F"),
				},
			},
			wantErr: false,
		},
		{
			name: "function with multiple arguments",
			args: args{
				vcs: multicall.ViewCalls{multicall.NewViewCall(
					"dosomething:0x0000",
					"0x0000",
					"doSomething(address,address,uint256)(uint256)",
					[]interface{}{
						"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
						"0xDBC05B1ECB4FDAEF943819C0B04E9EF6DF4BABD6",
						"18",
					},
				)},
			},
			want: []multicall.Multicall3Call{
				{
					Target:   common.HexToAddress("0x0000"),
					CallData: decodeHelper("c82fa99300000000000000000000000071C7656EC7ab88b098defB751B7401B5f6d8976F000000000000000000000000DBC05B1ECB4FDAEF943819C0B04E9EF6DF4BABD60000000000000000000000000000000000000000000000000000000000000012"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCallData(tt.args.vcs)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCallData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCallData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func decodeHelper(s string) []byte {
	ret, err := hex.DecodeString(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("return value: ", ret)
	return ret
}

func TestResultToString(t *testing.T) {
	type args struct {
		results interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "slice of integers input",
			args: args{
				results: []int{1, 2, 3, 4, 5},
			},
			want:    "[1,2,3,4,5]",
			wantErr: false,
		},
		{
			name: "slice of string input",
			args: args{
				results: []string{"This", "is", "test", "input"},
			},
			want:    `["This","is","test","input"]`,
			wantErr: false,
		},
		{
			name: "map of string, int input",
			args: args{
				results: map[string]int{
					"key1": 1,
					"key2": 2,
				},
			},
			want:    `{"key1":1,"key2":2}`,
			wantErr: false,
		},
		{
			name: "empty input",
			args: args{
				results: nil,
			},
			want:    "null",
			wantErr: false,
		},
		{
			name: "invalid input",
			args: args{
				results: make(chan string),
			},
			want:    "QueryEngine::ResultToString - json: unsupported type: chan string",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ResultToString(tt.args.results)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResultToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid address",
			args: args{
				address: "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
			},
			wantErr: false,
		},
		{
			name: "short length address",
			args: args{
				address: "0x4e71A2E537B7f9D",
			},
			wantErr: true,
		},
		{
			name: "address not starting with 0x",
			args: args{
				address: "4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
			},
			wantErr: true,
		},
		{
			name: "address with invalid characters %, @ etc",
			args: args{
				address: "0x#$e71A2E537B7f9D9413D3991D37958c0b5e1e%@&",
			},
			wantErr: true,
		},
		{
			name: "empty address",
			args: args{
				address: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInterfaceToString(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "string input",
			args: args{
				value: interface{}("This is a test string"),
			},
			want:    "This is a test string",
			wantErr: false,
		},
		{
			name: "empty string input",
			args: args{
				value: interface{}(""),
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "nil input",
			args: args{
				value: nil,
			},
			wantErr: true,
		},
		{
			name: "integer input",
			args: args{
				value: interface{}(256),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InterfaceToString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("InterfaceToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InterfaceToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceToBool(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "true boolean input",
			args: args{
				value: interface{}(true),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "false boolean input",
			args: args{
				value: interface{}(false),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "nil input",
			args: args{
				value: nil,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "integer input",
			args: args{
				value: interface{}(256),
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InterfaceToBool(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("InterfaceToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InterfaceToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceToBigInt(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "string input",
			args: args{
				value: interface{}("256"),
			},
			want:    big.NewInt(256),
			wantErr: false,
		},
		{
			name: "empty string input",
			args: args{
				value: interface{}(""),
			},
			want:    big.NewInt(0),
			wantErr: false,
		},
		{
			name: "nil input",
			args: args{
				value: nil,
			},
			want:    big.NewInt(0),
			wantErr: true,
		},
		{
			name: "integer input",
			args: args{
				value: interface{}(256),
			},
			want:    big.NewInt(0),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InterfaceToBigInt(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("InterfaceToBigInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterfaceToBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatUnits(t *testing.T) {
	type args struct {
		value    *big.Int
		decimals int64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "large input",
			args: args{
				value:    big.NewInt(123456789123456789),
				decimals: 10,
			},
			want: 12345678.9123456789,
		},
		{
			name: "small input",
			args: args{
				value:    big.NewInt(123456),
				decimals: 6,
			},
			want: 0.123456,
		},
		{
			name: "zero input",
			args: args{
				value:    big.NewInt(0),
				decimals: 10,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatUnits(tt.args.value, tt.args.decimals); got != tt.want {
				t.Errorf("FormatUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLpPairRatio(t *testing.T) {
	type args struct {
		reserve1 *big.Int
		reserve2 *big.Int
	}
	tests := []struct {
		name  string
		args  args
		want  *big.Int
		want1 bool
	}{
		{
			name: "reserve1 > reserve2",
			args: args{
				reserve1: big.NewInt(1000),
				reserve2: big.NewInt(100),
			},
			want:  new(big.Int).Mul(big.NewInt(10), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)),
			want1: true,
		},
		{
			name: "reserve1 < reserve2",
			args: args{
				reserve1: big.NewInt(100),
				reserve2: big.NewInt(1000),
			},
			want:  new(big.Int).Mul(big.NewInt(10), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)),
			want1: false,
		},
		{
			name: "reserve1 = reserve2",
			args: args{
				reserve1: big.NewInt(100),
				reserve2: big.NewInt(100),
			},
			want:  new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil),
			want1: true,
		},
		{
			name: "reserve1 is zero",
			args: args{
				reserve1: big.NewInt(0),
				reserve2: big.NewInt(1000),
			},
			want:  new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil),
			want1: true,
		},
		{
			name: "reserve2 is zero",
			args: args{
				reserve1: big.NewInt(1000),
				reserve2: big.NewInt(0),
			},
			want:  new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil),
			want1: true,
		},
		{
			name: "reserve1 is nil",
			args: args{
				reserve1: nil,
				reserve2: big.NewInt(1000),
			},
			want:  new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil),
			want1: true,
		},
		{
			name: "reserve2 is nil",
			args: args{
				reserve1: big.NewInt(1000),
				reserve2: nil,
			},
			want:  new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil),
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetLpPairRatio(tt.args.reserve1, tt.args.reserve2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLpPairRatio() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetLpPairRatio() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAPY(t *testing.T) {
	type args struct {
		blockRate *big.Int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "zero blockrate",
			args: args{
				blockRate: big.NewInt(0),
			},
			want: 0,
		},
		{
			name: "1e9 blockrate",
			args: args{
				blockRate: big.NewInt(1000000000),
			},
			want: 0.5452009284922177,
		},
		{
			name: "500000 blockrate",
			args: args{
				blockRate: big.NewInt(500000),
			},
			want: 0.0002718624332986863,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := APY(tt.args.blockRate); got != tt.want {
				t.Errorf("APY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_distributionAPY(t *testing.T) {
	type args struct {
		compSupplySpeed float64
		tokenSupply     float64
		tokenPrice      float64
		cantoPrice      float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "zero token price",
			args: args{
				compSupplySpeed: 10,
				tokenSupply:     100000,
				tokenPrice:      0,
				cantoPrice:      0.5,
			},
			want: 0,
		},
		{
			name: "zero token supply",
			args: args{
				compSupplySpeed: 10,
				tokenSupply:     0,
				tokenPrice:      1,
				cantoPrice:      0.5,
			},
			want: 0,
		},
		{
			name: "non zero inputs",
			args: args{
				compSupplySpeed: 10,
				tokenSupply:     100000,
				tokenPrice:      1,
				cantoPrice:      0.5,
			},
			want: 27186.20689655173,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributionAPY(tt.args.compSupplySpeed, tt.args.tokenSupply, tt.args.tokenPrice, tt.args.cantoPrice); got != tt.want {
				t.Errorf("distributionAPY() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestGetProcessedPairs(t *testing.T) {
// 	fpiJsonFile := "../../config/jsons/fpi_mainnet.json"
// 	contractsJsonFile := "../../config/jsons/contracts.json"
// 	config.NewConfig(fpiJsonFile, contractsJsonFile)
// 	type args struct {
// 		ctx   context.Context
// 		pairs PairsMap
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  []ProcessedPair
// 		want1 map[string]string
// 	}{
// 		{
// 			name: "empty pairs input",
// 			args: args{
// 				ctx:   context.Background(),
// 				pairs: nil,
// 			},
// 			want:  []ProcessedPair{},
// 			want1: map[string]string{},
// 		},
// 		{
// 			name: "nil pairs input",
// 			args: args{
// 				ctx:   context.Background(),
// 				pairs: nil,
// 			},
// 			want:  []ProcessedPair{},
// 			want1: map[string]string{},
// 		},
// 		{
// 			name: "map with one pair",
// 			args: args{
// 				ctx: context.Background(),
// 				pairs: map[string]map[string][]interface{}{
// 					"0x216400ba362d8FCE640085755e47075109718C8B": {
// 						"underlyingPriceTokenA": []interface{}{"102054337703150806"},
// 						"underlyingPriceLp":     []interface{}{"27518245740008256705"},
// 						"stable":                []interface{}{false},
// 						"totalSupply":           []interface{}{"158447005043477829048735"},
// 						"tokens":                []interface{}{"0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687", "0x826551890Dc65655a0Aceca109aB11AbDbD7a07B"},
// 						"reserves":              []interface{}{"21339635928010035060601074", "1176501303977446962544"},
// 						"underlyingPriceTokenB": []interface{}{"1855039815280839814183"},
// 					},
// 				},
// 			},
// 			want: []ProcessedPair{
// 				{
// 					Address:  "0x216400ba362d8FCE640085755e47075109718C8B",
// 					Symbol:   "CantoETHLP",
// 					Decimals: 18,
// 					Token1: config.Token{
// 						Name:       "Wrapped Canto",
// 						Address:    "0x826551890Dc65655a0Aceca109aB11AbDbD7a07B",
// 						Symbol:     "wCANTO",
// 						Decimals:   18,
// 						Underlying: "",
// 						ChainID:    "7700",
// 						LogoURI:    "https://canto.io/tokens/Canto.svg",
// 					},
// 					Token2: config.Token{
// 						Name:       "Ethereum",
// 						Address:    "0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687",
// 						Symbol:     "ETH",
// 						Decimals:   18,
// 						Underlying: "",
// 						ChainID:    "7700",
// 						LogoURI:    "https://canto.io/tokens/ETH.svg",
// 					},
// 					Stable:      false,
// 					CDecimal:    18,
// 					CLpAddress:  "0xb49A395B39A0b410675406bEE7bD06330CB503E3",
// 					TotalSupply: "158447005043477829048735",
// 					Tvl:         "4360183621554750534183584.00",
// 					Ratio:       "18138216979332057712297",
// 					AToB:        true,
// 					Price1:      "102054337703150806",
// 					Price2:      "1855039815280839814183",
// 					LpPrice:     "27518245740008256705",
// 					Reserve1:    "21339635928010035060601074",
// 					Reserve2:    "1176501303977446962544",
// 				},
// 			},
// 			want1: map[string]string{
// 				"0x216400ba362d8FCE640085755e47075109718C8B": "{\"address\":\"0x216400ba362d8FCE640085755e47075109718C8B\",\"symbol\":\"CantoETHLP\",\"decimals\":18,\"token1\":{\"name\":\"Wrapped Canto\",\"address\":\"0x826551890Dc65655a0Aceca109aB11AbDbD7a07B\",\"symbol\":\"wCANTO\",\"decimals\":18,\"chainId\":\"7700\",\"logoURI\":\"https://canto.io/tokens/Canto.svg\"},\"token2\":{\"name\":\"Ethereum\",\"address\":\"0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687\",\"symbol\":\"ETH\",\"decimals\":18,\"chainId\":\"7700\",\"logoURI\":\"https://canto.io/tokens/ETH.svg\"},\"stable\":false,\"cDecimals\":18,\"cLpAddress\":\"0xb49A395B39A0b410675406bEE7bD06330CB503E3\",\"totalSupply\":\"158447005043477829048735\",\"tvl\":\"4360183621554750534183584.00\",\"ratio\":\"18138216979332057712297\",\"aTob\":true,\"price1\":\"102054337703150806\",\"price2\":\"1855039815280839814183\",\"lpPrice\":\"27518245740008256705\",\"reserve1\":\"21339635928010035060601074\",\"reserve2\":\"1176501303977446962544\"}",
// 			},
// 		},
// 		{
// 			name: "map with two pairs",
// 			args: args{
// 				ctx: context.Background(),
// 				pairs: map[string]map[string][]interface{}{
// 					"0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833": {
// 						"underlyingPriceTokenA": []interface{}{"1000000000000000000"},
// 						"underlyingPriceLp":     []interface{}{"1996635704103199593949505"},
// 						"stable":                []interface{}{true},
// 						"totalSupply":           []interface{}{"3097284938113219610"},
// 						"tokens":                []interface{}{"0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503", "0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd"},
// 						"reserves":              []interface{}{"2740031970260676545709047", "3454842755268"},
// 						"underlyingPriceTokenB": []interface{}{"996895679424841423000000000000"},
// 					},
// 					"0x216400ba362d8FCE640085755e47075109718C8B": {
// 						"underlyingPriceTokenA": []interface{}{"102054337703150806"},
// 						"underlyingPriceLp":     []interface{}{"27518245740008256705"},
// 						"stable":                []interface{}{false},
// 						"totalSupply":           []interface{}{"158447005043477829048735"},
// 						"tokens":                []interface{}{"0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687", "0x826551890Dc65655a0Aceca109aB11AbDbD7a07B"},
// 						"reserves":              []interface{}{"21339635928010035060601074", "1176501303977446962544"},
// 						"underlyingPriceTokenB": []interface{}{"1855039815280839814183"},
// 					},
// 				},
// 			},
// 			want: []ProcessedPair{
// 				{
// 					Address:  "0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833",
// 					Symbol:   "NoteUSDTLP",
// 					Decimals: 18,
// 					Token1: config.Token{
// 						Name:       "Note",
// 						Address:    "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
// 						Symbol:     "NOTE",
// 						Decimals:   18,
// 						Underlying: "",
// 						ChainID:    "7700",
// 						LogoURI:    "https://canto.io/tokens/Note.svg",
// 					},
// 					Token2: config.Token{
// 						Name:       "USD Tether",
// 						Address:    "0xd567B3d7B8FE3C79a1AD8dA978812cfC4Fa05e75",
// 						Symbol:     "USDT",
// 						Decimals:   6,
// 						Underlying: "",
// 						ChainID:    "7700",
// 						LogoURI:    "https://canto.io/tokens/USDT.svg",
// 					},
// 					Stable:      true,
// 					CDecimal:    18,
// 					CLpAddress:  "0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F",
// 					TotalSupply: "3097284938113219610",
// 					Tvl:         "6184149693217923215678160.00",
// 					Ratio:       "793098894611811657388465274975",
// 					AToB:        true,
// 					Price1:      "1000000000000000000",
// 					Price2:      "996895679424841423000000000000",
// 					LpPrice:     "1996635704103199593949505",
// 					Reserve1:    "2740031970260676545709047",
// 					Reserve2:    "3454842755268",
// 				},
// 				{
// 					Address:  "0x216400ba362d8FCE640085755e47075109718C8B",
// 					Symbol:   "CantoETHLP",
// 					Decimals: 18,
// 					Token1: config.Token{
// 						Name:       "Wrapped Canto",
// 						Address:    "0x826551890Dc65655a0Aceca109aB11AbDbD7a07B",
// 						Symbol:     "wCANTO",
// 						Decimals:   18,
// 						Underlying: "",
// 						ChainID:    "7700",
// 						LogoURI:    "https://canto.io/tokens/Canto.svg",
// 					},
// 					Token2: config.Token{
// 						Name:       "Ethereum",
// 						Address:    "0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687",
// 						Symbol:     "ETH",
// 						Decimals:   18,
// 						Underlying: "",
// 						ChainID:    "7700",
// 						LogoURI:    "https://canto.io/tokens/ETH.svg",
// 					},
// 					Stable:      false,
// 					CDecimal:    18,
// 					CLpAddress:  "0xb49A395B39A0b410675406bEE7bD06330CB503E3",
// 					TotalSupply: "158447005043477829048735",
// 					Tvl:         "4360183621554750534183584.00",
// 					Ratio:       "18138216979332057712297",
// 					AToB:        true,
// 					Price1:      "102054337703150806",
// 					Price2:      "1855039815280839814183",
// 					LpPrice:     "27518245740008256705",
// 					Reserve1:    "21339635928010035060601074",
// 					Reserve2:    "1176501303977446962544",
// 				},
// 			},
// 			want1: map[string]string{
// 				"0x216400ba362d8FCE640085755e47075109718C8B": "{\"address\":\"0x216400ba362d8FCE640085755e47075109718C8B\",\"symbol\":\"CantoETHLP\",\"decimals\":18,\"token1\":{\"name\":\"Wrapped Canto\",\"address\":\"0x826551890Dc65655a0Aceca109aB11AbDbD7a07B\",\"symbol\":\"wCANTO\",\"decimals\":18,\"chainId\":\"7700\",\"logoURI\":\"https://canto.io/tokens/Canto.svg\"},\"token2\":{\"name\":\"Ethereum\",\"address\":\"0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687\",\"symbol\":\"ETH\",\"decimals\":18,\"chainId\":\"7700\",\"logoURI\":\"https://canto.io/tokens/ETH.svg\"},\"stable\":false,\"cDecimals\":18,\"cLpAddress\":\"0xb49A395B39A0b410675406bEE7bD06330CB503E3\",\"totalSupply\":\"158447005043477829048735\",\"tvl\":\"4360183621554750534183584.00\",\"ratio\":\"18138216979332057712297\",\"aTob\":true,\"price1\":\"102054337703150806\",\"price2\":\"1855039815280839814183\",\"lpPrice\":\"27518245740008256705\",\"reserve1\":\"21339635928010035060601074\",\"reserve2\":\"1176501303977446962544\"}",
// 				"0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833": "{\"address\":\"0x35DB1f3a6A6F07f82C76fCC415dB6cFB1a7df833\",\"symbol\":\"NoteUSDTLP\",\"decimals\":18,\"token1\":{\"name\":\"Note\",\"address\":\"0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503\",\"symbol\":\"NOTE\",\"decimals\":18,\"chainId\":\"7700\",\"logoURI\":\"https://canto.io/tokens/Note.svg\"},\"token2\":{\"name\":\"USD Tether\",\"address\":\"0xd567B3d7B8FE3C79a1AD8dA978812cfC4Fa05e75\",\"symbol\":\"USDT\",\"decimals\":6,\"chainId\":\"7700\",\"logoURI\":\"https://canto.io/tokens/USDT.svg\"},\"stable\":true,\"cDecimals\":18,\"cLpAddress\":\"0xf0cd6b5cE8A01D1B81F1d8B76643866c5816b49F\",\"totalSupply\":\"3097284938113219610\",\"tvl\":\"6184149693217923215678160.00\",\"ratio\":\"793098894611811657388465274975\",\"aTob\":true,\"price1\":\"1000000000000000000\",\"price2\":\"996895679424841423000000000000\",\"lpPrice\":\"1996635704103199593949505\",\"reserve1\":\"2740031970260676545709047\",\"reserve2\":\"3454842755268\"}"},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1 := GetProcessedPairs(tt.args.ctx, tt.args.pairs)
// 			sort.Slice(got, func(i, j int) bool {
// 				return got[i].Address < got[j].Address
// 			})
// 			sort.Slice(tt.want, func(i, j int) bool {
// 				return tt.want[i].Address < tt.want[j].Address
// 			})
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetProcessedPairs() got = %v, want %v", got, tt.want)
// 			}
// 			if !reflect.DeepEqual(got1, tt.want1) {
// 				t.Errorf("GetProcessedPairs() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

// func TestGetProcessedCTokens(t *testing.T) {
// 	fpiJsonFile := "../../config/jsons/fpi_mainnet.json"
// 	contractsJsonFile := "../../config/jsons/contracts.json"
// 	config.NewConfig(fpiJsonFile, contractsJsonFile)
// 	type args struct {
// 		ctx     context.Context
// 		cTokens TokensMap
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  []ProcessedCToken
// 		want1 map[string]string
// 	}{
// 		{
// 			name: "map with one ctokens",
// 			args: args{
// 				ctx: context.Background(),
// 				cTokens: map[string]map[string][]interface{}{
// 					"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488": {
// 						"totalSupply":        []interface{}{"46051451200466458262227"},
// 						"supplyRatePerBlock": []interface{}{"0"},
// 						"borrowRatePerBlock": []interface{}{"0"},
// 						"underlyingPrice":    []interface{}{"100631493381492554"},
// 						"cash":               []interface{}{"46051451200466458262227"},
// 						"exchangeRateStored": []interface{}{"1000000000000000000"},
// 						"compBorrowSpeeds":   []interface{}{"0"},
// 						"borrowCaps":         []interface{}{"1"},
// 						"compSupplySpeeds":   []interface{}{"0"},
// 						"markets":            []interface{}{true, "0", false},
// 					},
// 				},
// 			},
// 			want: []ProcessedCToken{
// 				{
// 					Address:  "0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
// 					Symbol:   "cCANTO",
// 					Name:     "Collateral Canto",
// 					Decimals: 18,
// 					Underlying: config.Underlying{
// 						Address:  "0x826551890Dc65655a0Aceca109aB11AbDbD7a07B",
// 						Symbol:   "wCANTO",
// 						Name:     "Wrapped Canto",
// 						Decimals: 18,
// 					},
// 					Cash:             "46051451200466458262227",
// 					ExchangeRate:     "1000000000000000000",
// 					CollateralFactor: "0",
// 					Price:            "100631493381492554",
// 					BorrowCap:        "1",
// 					IsListed:         true,
// 					Liquidity:        "4634.23",
// 					SupplyApy:        "0.00",
// 					BorrowApy:        "0.00",
// 					DistApy:          "0.00",
// 				},
// 			},
// 			want1: map[string]string{
// 				"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488": "{\"address\":\"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488\",\"symbol\":\"cCANTO\",\"name\":\"Collateral Canto\",\"decimals\":18,\"underlying\":{\"address\":\"0x826551890Dc65655a0Aceca109aB11AbDbD7a07B\",\"symbol\":\"wCANTO\",\"name\":\"Wrapped Canto\",\"decimals\":18},\"cash\":\"46051451200466458262227\",\"exchangeRate\":\"1000000000000000000\",\"collateralFactor\":\"0\",\"price\":\"100631493381492554\",\"borrowCap\":\"1\",\"isListed\":true,\"liquidity\":\"4634.23\",\"supplyApy\":\"0.00\",\"borrowApy\":\"0.00\",\"distApy\":\"0.00\"}",
// 			},
// 		},
// 		{
// 			name: "map with two ctokens",
// 			args: args{
// 				ctx: context.Background(),
// 				cTokens: map[string]map[string][]interface{}{
// 					"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488": {
// 						"totalSupply":        []interface{}{"46051451200466458262227"},
// 						"supplyRatePerBlock": []interface{}{"0"},
// 						"borrowRatePerBlock": []interface{}{"0"},
// 						"underlyingPrice":    []interface{}{"100631493381492554"},
// 						"cash":               []interface{}{"46051451200466458262227"},
// 						"exchangeRateStored": []interface{}{"1000000000000000000"},
// 						"compBorrowSpeeds":   []interface{}{"0"},
// 						"borrowCaps":         []interface{}{"1"},
// 						"compSupplySpeeds":   []interface{}{"0"},
// 						"markets":            []interface{}{true, "0", false},
// 					},
// 					"0x830b9849e7d79b92408a86a557e7baaacbec6030": {
// 						"totalSupply":        []interface{}{"1192404869333399889"},
// 						"supplyRatePerBlock": []interface{}{"0"},
// 						"borrowRatePerBlock": []interface{}{"0"},
// 						"underlyingPrice":    []interface{}{"1875172785699970219034"},
// 						"cash":               []interface{}{"1192404869333399889"},
// 						"exchangeRateStored": []interface{}{"1000000000000000000"},
// 						"compBorrowSpeeds":   []interface{}{"0"},
// 						"borrowCaps":         []interface{}{"1"},
// 						"compSupplySpeeds":   []interface{}{"0"},
// 						"markets":            []interface{}{true, "0", false},
// 					},
// 				},
// 			},
// 			want: []ProcessedCToken{
// 				{
// 					Address:  "0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
// 					Symbol:   "cCANTO",
// 					Name:     "Collateral Canto",
// 					Decimals: 18,
// 					Underlying: config.Underlying{
// 						Address:  "0x826551890Dc65655a0Aceca109aB11AbDbD7a07B",
// 						Symbol:   "wCANTO",
// 						Name:     "Wrapped Canto",
// 						Decimals: 18,
// 					},
// 					Cash:             "46051451200466458262227",
// 					ExchangeRate:     "1000000000000000000",
// 					CollateralFactor: "0",
// 					Price:            "100631493381492554",
// 					BorrowCap:        "1",
// 					IsListed:         true,
// 					Liquidity:        "4634.23",
// 					SupplyApy:        "0.00",
// 					BorrowApy:        "0.00",
// 					DistApy:          "0.00",
// 				},
// 				{
// 					Address:  "0x830b9849e7d79b92408a86a557e7baaacbec6030",
// 					Symbol:   "cETH",
// 					Name:     "Collateral Ethereum",
// 					Decimals: 18,
// 					Underlying: config.Underlying{
// 						Address:  "0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687",
// 						Symbol:   "ETH",
// 						Name:     "Ethereum",
// 						Decimals: 18,
// 					},
// 					Cash:             "1192404869333399889",
// 					ExchangeRate:     "1000000000000000000",
// 					CollateralFactor: "0",
// 					Price:            "1875172785699970219034",
// 					BorrowCap:        "1",
// 					IsListed:         true,
// 					Liquidity:        "2235.97",
// 					SupplyApy:        "0.00",
// 					BorrowApy:        "0.00",
// 					DistApy:          "0.00",
// 				},
// 			},
// 			want1: map[string]string{
// 				"0x830b9849e7d79b92408a86a557e7baaacbec6030": "{\"address\":\"0x830b9849e7d79b92408a86a557e7baaacbec6030\",\"symbol\":\"cETH\",\"name\":\"Collateral Ethereum\",\"decimals\":18,\"underlying\":{\"address\":\"0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687\",\"symbol\":\"ETH\",\"name\":\"Ethereum\",\"decimals\":18},\"cash\":\"1192404869333399889\",\"exchangeRate\":\"1000000000000000000\",\"collateralFactor\":\"0\",\"price\":\"1875172785699970219034\",\"borrowCap\":\"1\",\"isListed\":true,\"liquidity\":\"2235.97\",\"supplyApy\":\"0.00\",\"borrowApy\":\"0.00\",\"distApy\":\"0.00\"}",
// 				"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488": "{\"address\":\"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488\",\"symbol\":\"cCANTO\",\"name\":\"Collateral Canto\",\"decimals\":18,\"underlying\":{\"address\":\"0x826551890Dc65655a0Aceca109aB11AbDbD7a07B\",\"symbol\":\"wCANTO\",\"name\":\"Wrapped Canto\",\"decimals\":18},\"cash\":\"46051451200466458262227\",\"exchangeRate\":\"1000000000000000000\",\"collateralFactor\":\"0\",\"price\":\"100631493381492554\",\"borrowCap\":\"1\",\"isListed\":true,\"liquidity\":\"4634.23\",\"supplyApy\":\"0.00\",\"borrowApy\":\"0.00\",\"distApy\":\"0.00\"}",
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1 := GetProcessedCTokens(tt.args.ctx, tt.args.cTokens)
// 			sort.Slice(got, func(i, j int) bool {
// 				return got[i].Address < got[j].Address
// 			})
// 			sort.Slice(tt.want, func(i, j int) bool {
// 				return tt.want[i].Address < tt.want[j].Address
// 			})
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetProcessedCTokens() got = %v, want %v", got, tt.want)
// 			}
// 			if !reflect.DeepEqual(got1, tt.want1) {
// 				t.Errorf("GetProcessedCTokens() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    tokens, err := UnmarshalTokens(bytes)
//    bytes, err = tokens.Marshal()

package config

import (
	"reflect"
	"testing"
)

func Test_getFPIFromJson(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    TokensInfo
		wantErr bool
	}{
		{
			name: "only one ctoken",
			args: args{
				path: "./jsons/tests/fpi_mainnet_test_01.json",
			},
			want: TokensInfo{
				Name:        "Canto Mainnet Token List",
				Version:     0.1,
				Keywords:    []string{"canto", "mainnet"},
				ChainID:     "7700",
				Comptroller: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
				Router:      "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
				Reservoir:   "0x07C50Bf0804A06860AeACAcFaf029F9a1c014F91",
				MulticallV3: "0xcA11bde05977b3631167028862bE2a173976CA11",
				CTokens: []Token{
					{
						Name:       "Collateral Note",
						Address:    "0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
						Symbol:     "cNOTE",
						Decimals:   18,
						Underlying: "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
						ChainID:    "7700",
					},
				},
				Tokens: []Token{},
				Pairs:  []Pair{},
			},
			wantErr: false,
		},
		{
			name: "only one token",
			args: args{
				path: "./jsons/tests/fpi_mainnet_test_02.json",
			},
			want: TokensInfo{
				Name:        "Canto Mainnet Token List",
				Version:     0.1,
				Keywords:    []string{"canto", "mainnet"},
				ChainID:     "7700",
				Comptroller: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
				Router:      "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
				Reservoir:   "0x07C50Bf0804A06860AeACAcFaf029F9a1c014F91",
				MulticallV3: "0xcA11bde05977b3631167028862bE2a173976CA11",
				CTokens:     []Token{},
				Tokens: []Token{
					{
						Name:     "Note",
						Address:  "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
						Symbol:   "NOTE",
						Decimals: 18,
						LogoURI:  "https://canto.io/tokens/Note.svg",
						ChainID:  "7700",
					},
				},
				Pairs: []Pair{},
			},
			wantErr: false,
		},
		{
			name: "only one pair",
			args: args{
				path: "./jsons/tests/fpi_mainnet_test_03.json",
			},
			want: TokensInfo{
				Name:        "Canto Mainnet Token List",
				Version:     0.1,
				Keywords:    []string{"canto", "mainnet"},
				ChainID:     "7700",
				Comptroller: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
				Router:      "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
				Reservoir:   "0x07C50Bf0804A06860AeACAcFaf029F9a1c014F91",
				MulticallV3: "0xcA11bde05977b3631167028862bE2a173976CA11",
				CTokens:     []Token{},
				Tokens:      []Token{},
				Pairs: []Pair{
					{
						Name:     "Canto Note Liquidity Provider",
						Symbol:   "CantoNoteLP",
						Decimals: 18,
						Address:  "0x1D20635535307208919f0b67c3B2065965A85aA9",
						Stable:   false,
						TokenA:   "0x826551890Dc65655a0Aceca109aB11AbDbD7a07B",
						TokenB:   "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
						ChainID:  "7700",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ctoken, token and pair",
			args: args{
				path: "./jsons/tests/fpi_mainnet_test_04.json",
			},
			want: TokensInfo{
				Name:        "Canto Mainnet Token List",
				Version:     0.1,
				Keywords:    []string{"canto", "mainnet"},
				ChainID:     "7700",
				Comptroller: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
				Router:      "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
				Reservoir:   "0x07C50Bf0804A06860AeACAcFaf029F9a1c014F91",
				MulticallV3: "0xcA11bde05977b3631167028862bE2a173976CA11",
				CTokens: []Token{
					{
						Name:       "Collateral Note",
						Address:    "0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
						Symbol:     "cNOTE",
						Decimals:   18,
						Underlying: "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
						ChainID:    "7700",
					},
				},
				Tokens: []Token{
					{
						Name:     "Note",
						Address:  "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
						Symbol:   "NOTE",
						Decimals: 18,
						LogoURI:  "https://canto.io/tokens/Note.svg",
						ChainID:  "7700",
					},
				},
				Pairs: []Pair{
					{
						Name:     "Canto Note Liquidity Provider",
						Symbol:   "CantoNoteLP",
						Decimals: 18,
						Address:  "0x1D20635535307208919f0b67c3B2065965A85aA9",
						Stable:   false,
						TokenA:   "0x826551890Dc65655a0Aceca109aB11AbDbD7a07B",
						TokenB:   "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
						ChainID:  "7700",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid path to json",
			args: args{
				path: "./jsons/tests/invalid_path.json",
			},
			want:    TokensInfo{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFPIFromJson(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFPIFromJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFPIFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCTokenAddress(t *testing.T) {
	fpiJsonFile := "./jsons/fpi_mainnet.json"
	contractsJsonFile := "./jsons/contracts.json"
	NewConfig(fpiJsonFile, contractsJsonFile)
	type args struct {
		underlyingAddress string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty adress",
			args: args{
				underlyingAddress: "",
			},
			want: "",
		},
		{
			name: "invalid adress",
			args: args{
				underlyingAddress: "0x2222222222222222",
			},
			want: "",
		},
		{
			name: "valid adress",
			args: args{
				underlyingAddress: "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
			},
			want: "0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCTokenAddress := GetCTokenAddress(tt.args.underlyingAddress); gotCTokenAddress != tt.want {
				t.Errorf("GetCTokenAddress() = %v, want %v", gotCTokenAddress, tt.want)
			}
		})
	}
}

func TestGetCTokenAddressBySymbol(t *testing.T) {
	fpiJsonFile := "./jsons/fpi_mainnet.json"
	contractsJsonFile := "./jsons/contracts.json"
	NewConfig(fpiJsonFile, contractsJsonFile)
	type args struct {
		symbol string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty symbol",
			args: args{
				symbol: "",
			},
			want: "",
		},
		{
			name: "invalid symbol",
			args: args{
				symbol: "$$$$$",
			},
			want: "",
		},
		{
			name: "valid symbol",
			args: args{
				symbol: "cNOTE",
			},
			want: "0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCTokenAddress := GetCTokenAddressBySymbol(tt.args.symbol); gotCTokenAddress != tt.want {
				t.Errorf("GetCTokenAddressBySymbol() = %v, want %v", gotCTokenAddress, tt.want)
			}
		})
	}
}

func TestGetCTokenDecimals(t *testing.T) {
	fpiJsonFile := "./jsons/fpi_mainnet.json"
	contractsJsonFile := "./jsons/contracts.json"
	NewConfig(fpiJsonFile, contractsJsonFile)
	type args struct {
		underlyingAddress string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "empty address",
			args: args{
				underlyingAddress: "",
			},
			want: 0,
		},
		{
			name: "invalid address",
			args: args{
				underlyingAddress: "0x2222222222222222",
			},
			want: 0,
		},
		{
			name: "valid address",
			args: args{
				underlyingAddress: "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCTokenDecimals(tt.args.underlyingAddress); got != tt.want {
				t.Errorf("GetCTokenDecimals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTokenData(t *testing.T) {
	fpiJsonFile := "./jsons/fpi_mainnet.json"
	contractsJsonFile := "./jsons/contracts.json"
	NewConfig(fpiJsonFile, contractsJsonFile)
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want Token
	}{
		{
			name: "empty address",
			args: args{
				address: "",
			},
			want: Token{},
		},
		{
			name: "invalid address",
			args: args{
				address: "0x2222222222222222",
			},
			want: Token{},
		},
		{
			name: "valid address",
			args: args{
				address: "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
			},
			want: Token{
				Name:     "Note",
				Address:  "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
				Symbol:   "NOTE",
				Decimals: 18,
				LogoURI:  "https://canto.io/tokens/Note.svg",
				ChainID:  "7700",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTokenData(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTokenData() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetUnderlyingData(t *testing.T) {
	fpiJsonFile := "./jsons/fpi_mainnet.json"
	contractsJsonFile := "./jsons/contracts.json"
	NewConfig(fpiJsonFile, contractsJsonFile)
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want Underlying
	}{
		{
			name: "empty address",
			args: args{
				address: "",
			},
			want: Underlying{},
		},
		{
			name: "invalid address",
			args: args{
				address: "0x2222222222222222",
			},
			want: Underlying{},
		},
		{
			name: "valid token address",
			args: args{
				address: "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
			},
			want: Underlying{
				Name:     "Note",
				Address:  "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
				Symbol:   "NOTE",
				Decimals: 18,
			},
		},
		{
			name: "valid pair address",
			args: args{
				address: "0x1D20635535307208919f0b67c3B2065965A85aA9",
			},
			want: Underlying{
				Name:     "Canto Note Liquidity Provider",
				Symbol:   "CantoNoteLP",
				Decimals: 18,
				Address:  "0x1D20635535307208919f0b67c3B2065965A85aA9",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUnderlyingData(tt.args.address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnderlyingData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLpPairData(t *testing.T) {
	fpiJsonFile := "./jsons/fpi_mainnet.json"
	contractsJsonFile := "./jsons/contracts.json"
	NewConfig(fpiJsonFile, contractsJsonFile)
	type args struct {
		address string
	}
	tests := []struct {
		name           string
		args           args
		wantSymbol     string
		wantDecimals   int64
		wantToken1     Token
		wantToken2     Token
		wantStable     bool
		wantCDecimals  int64
		wantCLpAddress string
	}{
		{
			name: "empty address",
			args: args{
				address: "",
			},
			wantSymbol:     "",
			wantDecimals:   0,
			wantToken1:     Token{},
			wantToken2:     Token{},
			wantStable:     false,
			wantCDecimals:  0,
			wantCLpAddress: "",
		},
		{
			name: "invalid address",
			args: args{
				address: "0x2222222222222222",
			},
			wantSymbol:     "",
			wantDecimals:   0,
			wantToken1:     Token{},
			wantToken2:     Token{},
			wantStable:     false,
			wantCDecimals:  0,
			wantCLpAddress: "",
		},
		{
			name: "valid address",
			args: args{
				address: "0x1D20635535307208919f0b67c3B2065965A85aA9",
			},
			wantSymbol:   "CantoNoteLP",
			wantDecimals: 18,
			wantToken1: Token{
				Name:     "Wrapped Canto",
				Address:  "0x826551890Dc65655a0Aceca109aB11AbDbD7a07B",
				Symbol:   "wCANTO",
				Decimals: 18,
				LogoURI:  "https://canto.io/tokens/Canto.svg",
				ChainID:  "7700",
			},
			wantToken2: Token{
				Name:     "Note",
				Address:  "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
				Symbol:   "NOTE",
				Decimals: 18,
				LogoURI:  "https://canto.io/tokens/Note.svg",
				ChainID:  "7700",
			},
			wantStable:     false,
			wantCDecimals:  18,
			wantCLpAddress: "0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSymbol, gotDecimals, gotToken1, gotToken2, gotStable, gotCDecimals, gotCLpAddress := GetLpPairData(tt.args.address)
			if gotSymbol != tt.wantSymbol {
				t.Errorf("GetLpPairData() gotSymbol = %v, want %v", gotSymbol, tt.wantSymbol)
			}
			if gotDecimals != tt.wantDecimals {
				t.Errorf("GetLpPairData() gotDecimals = %v, want %v", gotDecimals, tt.wantDecimals)
			}
			if !reflect.DeepEqual(gotToken1, tt.wantToken1) {
				t.Errorf("GetLpPairData() gotToken1 = %v, want %v", gotToken1, tt.wantToken1)
			}
			if !reflect.DeepEqual(gotToken2, tt.wantToken2) {
				t.Errorf("GetLpPairData() gotToken2 = %v, want %v", gotToken2, tt.wantToken2)
			}
			if gotStable != tt.wantStable {
				t.Errorf("GetLpPairData() gotStable = %v, want %v", gotStable, tt.wantStable)
			}
			if gotCDecimals != tt.wantCDecimals {
				t.Errorf("GetLpPairData() gotCDecimals = %v, want %v", gotCDecimals, tt.wantCDecimals)
			}
			if gotCLpAddress != tt.wantCLpAddress {
				t.Errorf("GetLpPairData() gotCLpAddress = %v, want %v", gotCLpAddress, tt.wantCLpAddress)
			}
		})
	}
}

func TestGetCTokenData(t *testing.T) {
	fpiJsonFile := "./jsons/fpi_mainnet.json"
	contractsJsonFile := "./jsons/contracts.json"
	NewConfig(fpiJsonFile, contractsJsonFile)
	type args struct {
		address string
	}
	tests := []struct {
		name           string
		args           args
		wantSymbol     string
		wantName       string
		wantDecimals   int64
		wantUnderlying Underlying
	}{
		{
			name: "empty address",
			args: args{
				address: "",
			},
			wantSymbol:     "",
			wantName:       "",
			wantDecimals:   0,
			wantUnderlying: Underlying{},
		},
		{
			name: "invalid address",
			args: args{
				address: "0x2222222222222222",
			},
			wantSymbol:     "",
			wantName:       "",
			wantDecimals:   0,
			wantUnderlying: Underlying{},
		},
		{
			name: "valid token address",
			args: args{
				address: "0xEe602429Ef7eCe0a13e4FfE8dBC16e101049504C",
			},
			wantSymbol:   "cNOTE",
			wantName:     "Collateral Note",
			wantDecimals: 18,
			wantUnderlying: Underlying{
				Name:     "Note",
				Address:  "0x4e71A2E537B7f9D9413D3991D37958c0b5e1e503",
				Symbol:   "NOTE",
				Decimals: 18,
			},
		},
		{
			name: "valid pair address",
			args: args{
				address: "0x3C96dCfd875253A37acB3D2B102b6f328349b16B",
			},
			wantSymbol:   "cCantoNoteLP",
			wantName:     "Collateral Canto Note Liquidity Provider",
			wantDecimals: 18,
			wantUnderlying: Underlying{
				Name:     "Canto Note Liquidity Provider",
				Address:  "0x1D20635535307208919f0b67c3B2065965A85aA9",
				Symbol:   "CantoNoteLP",
				Decimals: 18,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSymbol, gotName, gotDecimals, gotUnderlying := GetCTokenData(tt.args.address)
			if gotSymbol != tt.wantSymbol {
				t.Errorf("GetCTokenData() gotSymbol = %v, want %v", gotSymbol, tt.wantSymbol)
			}
			if gotName != tt.wantName {
				t.Errorf("GetCTokenData() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotDecimals != tt.wantDecimals {
				t.Errorf("GetCTokenData() gotDecimals = %v, want %v", gotDecimals, tt.wantDecimals)
			}
			if !reflect.DeepEqual(gotUnderlying, tt.wantUnderlying) {
				t.Errorf("GetCTokenData() gotUnderlying = %v, want %v", gotUnderlying, tt.wantUnderlying)
			}
		})
	}
}

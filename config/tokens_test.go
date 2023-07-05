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

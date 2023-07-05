package queryengine

import (
	"canto-api/config"
	"canto-api/multicall"
	"context"
	"reflect"
	"testing"
)

func TestProcessContractCalls(t *testing.T) {
	type args struct {
		contracts []config.Contract
	}

	tests := []struct {
		name    string
		args    args
		want    multicall.ViewCalls
		wantErr bool
	}{
		{
			name: "incorrect target address",
			args: args{
				contracts: []config.Contract{
					{
						Name:    "base contract",
						Address: "0x00",
						Keys: []string{
							"decimals:0x00",
						},
						Methods: []string{
							"decimals()",
						},
						Args: [][]interface{}{
							{},
						},
					},
				},
			},
			want: multicall.ViewCalls{multicall.NewViewCall(
				"decimals:0x00",
				"0x00",
				"decimals()",
				[]interface{}{},
			)},
			wantErr: true,
		},
		{
			name: "one contract, one function, with no arguments",
			args: args{
				contracts: []config.Contract{
					{
						Name:    "base contract",
						Address: "0x0000000000000000000000000000000000000000",
						Keys: []string{
							"decimals:0x0000000000000000000000000000000000000000",
						},
						Methods: []string{
							"decimals()",
						},
						Args: [][]interface{}{
							{},
						},
					},
				},
			},
			want: multicall.ViewCalls{multicall.NewViewCall(
				"decimals:0x0000000000000000000000000000000000000000",
				"0x0000000000000000000000000000000000000000",
				"decimals()",
				[]interface{}{},
			)},
			wantErr: false,
		},
		{
			name: "one contract, one function, with one argument",
			args: args{
				contracts: []config.Contract{
					{
						Name:    "sample erc20",
						Address: "0x0000000000000000000000000000000000000000",
						Keys: []string{
							"balanceof:0x0000000000000000000000000000000000000000",
						},
						Methods: []string{
							"balanceOf(address)(uint256)",
						},
						Args: [][]interface{}{
							{
								"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
							},
						},
					},
				},
			},
			want: multicall.ViewCalls{multicall.NewViewCall(
				"balanceof:0x0000000000000000000000000000000000000000",
				"0x0000000000000000000000000000000000000000",
				"balanceOf(address)(uint256)",
				[]interface{}{"0x71C7656EC7ab88b098defB751B7401B5f6d8976F"},
			)},
			wantErr: false,
		},
		{
			name: "one contract, multiple functions, with multiple arguments",
			args: args{
				contracts: []config.Contract{
					{
						Name:    "sample erc20",
						Address: "0x0000000000000000000000000000000000000000",
						Keys: []string{
							"balanceof:0x0000000000000000000000000000000000000000",
							"allowance:0x0000000000000000000000000000000000000000",
						},
						Methods: []string{
							"balanceOf(address)(uint256)",
							"allowance(address,address)(uint256)",
						},
						Args: [][]interface{}{
							{
								"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
							},
							{
								"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
								"0xDBC05B1ECB4FDAEF943819C0B04E9EF6DF4BABD6",
							},
						},
					},
				},
			},
			want: multicall.ViewCalls{
				multicall.NewViewCall(
					"balanceof:0x0000000000000000000000000000000000000000",
					"0x0000000000000000000000000000000000000000",
					"balanceOf(address)(uint256)",
					[]interface{}{"0x71C7656EC7ab88b098defB751B7401B5f6d8976F"},
				),
				multicall.NewViewCall(
					"allowance:0x0000000000000000000000000000000000000000",
					"0x0000000000000000000000000000000000000000",
					"allowance(address,address)(uint256)",
					[]interface{}{
						"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
						"0xDBC05B1ECB4FDAEF943819C0B04E9EF6DF4BABD6",
					},
				),
			},
			wantErr: false,
		},
		{
			name: "multiple contracts, multiple functions, with multiple arguments",
			args: args{
				contracts: []config.Contract{
					{
						Name:    "sample erc20",
						Address: "0x0000000000000000000000000000000000000000",
						Keys: []string{
							"balanceof:0x0000000000000000000000000000000000000000",
							"allowance:0x0000000000000000000000000000000000000000",
						},
						Methods: []string{
							"balanceOf(address)(uint256)",
							"allowance(address,address)(uint256)",
						},
						Args: [][]interface{}{
							{
								"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
							},
							{
								"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
								"0xDBC05B1ECB4FDAEF943819C0B04E9EF6DF4BABD6",
							},
						},
					},
					{
						Name:    "sample pair",
						Address: "0x0000000000000000000000000000000000000001",
						Keys: []string{
							"balanceof:0x0000000000000000000000000000000000000001",
						},
						Methods: []string{
							"balanceOf(address)(uint256)",
						},
						Args: [][]interface{}{
							{
								"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
							},
						},
					},
				},
			},
			want: multicall.ViewCalls{
				multicall.NewViewCall(
					"balanceof:0x0000000000000000000000000000000000000000",
					"0x0000000000000000000000000000000000000000",
					"balanceOf(address)(uint256)",
					[]interface{}{"0x71C7656EC7ab88b098defB751B7401B5f6d8976F"},
				),
				multicall.NewViewCall(
					"allowance:0x0000000000000000000000000000000000000000",
					"0x0000000000000000000000000000000000000000",
					"allowance(address,address)(uint256)",
					[]interface{}{
						"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
						"0xDBC05B1ECB4FDAEF943819C0B04E9EF6DF4BABD6",
					},
				),
				multicall.NewViewCall(
					"balanceof:0x0000000000000000000000000000000000000001",
					"0x0000000000000000000000000000000000000001",
					"balanceOf(address)(uint256)",
					[]interface{}{
						"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
					},
				),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessContractCalls(tt.args.contracts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessContractCalls() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("ProcessContractCalls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessMulticallResults(t *testing.T) {
	type args struct {
		ctx     context.Context
		results *multicall.Result
	}
	tests := []struct {
		name    string
		args    args
		want    TokensMap
		want1   PairsMap
		want2   map[string][]interface{}
		wantErr bool
	}{
		{
			name: "nil input",
			args: args{
				ctx:     context.Background(),
				results: nil,
			},
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name: "nil calls input",
			args: args{
				ctx: context.Background(),
				results: &multicall.Result{
					BlockNumber: 4885700,
					Calls:       nil,
				},
			},
			want:    TokensMap{},
			want1:   PairsMap{},
			want2:   map[string][]interface{}{},
			wantErr: false,
		},
		{
			name: "only ctoken input",
			args: args{
				ctx: context.Background(),
				results: &multicall.Result{
					BlockNumber: 4885700,
					Calls: map[string][]interface{}{
						"cTokens:0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488:borrowCaps":         {1},
						"cTokens:0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488:borrowRatePerBlock": {0},
					},
				},
			},
			want: TokensMap{
				"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488": {
					"borrowCaps":         {1},
					"borrowRatePerBlock": {0},
				},
			},
			want1:   PairsMap{},
			want2:   map[string][]interface{}{},
			wantErr: false,
		},
		{
			name: "only pair input",
			args: args{
				ctx: context.Background(),
				results: &multicall.Result{
					BlockNumber: 4885700,
					Calls: map[string][]interface{}{
						"lpPairs:0x1D20635535307208919f0b67c3B2065965A85aA9:totalSupply": {1143452325947156014},
						"lpPairs:0x1D20635535307208919f0b67c3B2065965A85aA9:reserves":    {3583523874919165340, 364880926971712949},
					},
				},
			},
			want: TokensMap{},
			want1: PairsMap{
				"0x1D20635535307208919f0b67c3B2065965A85aA9": map[string][]interface{}{
					"totalSupply": {1143452325947156014},
					"reserves":    {3583523874919165340, 364880926971712949},
				},
			},
			want2:   map[string][]interface{}{},
			wantErr: false,
		},
		{
			name: "only other input",
			args: args{
				ctx: context.Background(),
				results: &multicall.Result{
					BlockNumber: 4885700,
					Calls: map[string][]interface{}{
						"pricefeed:getUnderlyingPrice:0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488": {1143452325947156014},
					},
				},
			},
			want:  TokensMap{},
			want1: PairsMap{},
			want2: map[string][]interface{}{
				"pricefeed:getUnderlyingPrice:0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488": {1143452325947156014},
			},
			wantErr: false,
		},
		{
			name: "invalid pair input",
			args: args{
				ctx: context.Background(),
				results: &multicall.Result{
					BlockNumber: 4885700,
					Calls: map[string][]interface{}{
						"lpPairs:totalSupply": {1143452325947156014},
						"lpPairs:0x1D20635535307208919f0b67c3B2065965A85aA9:reserves": {3583523874919165340, 364880926971712949},
					},
				},
			},
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name: "invalid ctoken input",
			args: args{
				ctx: context.Background(),
				results: &multicall.Result{
					BlockNumber: 4885700,
					Calls: map[string][]interface{}{
						"cTokens:0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488:borrowCaps": {1},
						"cTokens:0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488":            {0},
					},
				},
			},
			want:    nil,
			want1:   nil,
			want2:   nil,
			wantErr: true,
		},
		{
			name: "ctoken, pair and other input",
			args: args{
				ctx: context.Background(),
				results: &multicall.Result{
					BlockNumber: 4885700,
					Calls: map[string][]interface{}{
						"cTokens:0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488:borrowCaps":           {1},
						"cTokens:0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488:borrowRatePerBlock":   {0},
						"lpPairs:0x1D20635535307208919f0b67c3B2065965A85aA9:totalSupply":          {1143452325947156014},
						"lpPairs:0x1D20635535307208919f0b67c3B2065965A85aA9:reserves":             {3583523874919165340, 364880926971712949},
						"pricefeed:getUnderlyingPrice:0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488": {1143452325947156014},
					},
				},
			},
			want: TokensMap{
				"0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488": {
					"borrowCaps":         {1},
					"borrowRatePerBlock": {0},
				},
			},
			want1: PairsMap{
				"0x1D20635535307208919f0b67c3B2065965A85aA9": map[string][]interface{}{
					"totalSupply": {1143452325947156014},
					"reserves":    {3583523874919165340, 364880926971712949},
				},
			},
			want2: map[string][]interface{}{
				"pricefeed:getUnderlyingPrice:0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488": {1143452325947156014},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got, got1, got2, err := ProcessMulticallResults(tt.args.ctx, tt.args.results)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessMulticallResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessMulticallResults() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProcessMulticallResults() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("ProcessMulticallResults() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

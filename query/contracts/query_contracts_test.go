package query

import (
	"canto-api/config"
	"canto-api/multicall"
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

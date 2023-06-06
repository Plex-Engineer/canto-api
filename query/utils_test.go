package query

import (
	"canto-api/config"
	"canto-api/multicall"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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
						Address: "0x00",
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
				"Balance/0x00",
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
						// Name:    "base contract",
						Address: "0x0000000000000000000000000000000000000000",
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
				"Decimals/0x0000000000000000000000000000000000000000",
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
						// Name:    "sample erc20",
						Address: "0x0000000000000000000000000000000000000000",
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
				"Balance/0x0000000000000000000000000000000000000000",
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
						// Name:    "sample erc20",
						Address: "0x0000000000000000000000000000000000000000",
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
					"Balance/0x0000000000000000000000000000000000000000",
					"0x0000000000000000000000000000000000000000",
					"balanceOf(address)(uint256)",
					[]interface{}{"0x71C7656EC7ab88b098defB751B7401B5f6d8976F"},
				),
				multicall.NewViewCall(
					"Allowance/0x0000000000000000000000000000000000000000",
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
						// Name:    "sample erc20",
						Address: "0x0000000000000000000000000000000000000000",
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
						// Name:    "sample pair",
						Address: "0x0000000000000000000000000000000000000001",
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
					"Balance/0x0000000000000000000000000000000000000000",
					"0x0000000000000000000000000000000000000000",
					"balanceOf(address)(uint256)",
					[]interface{}{"0x71C7656EC7ab88b098defB751B7401B5f6d8976F"},
				),
				multicall.NewViewCall(
					"Allowance/0x0000000000000000000000000000000000000000",
					"0x0000000000000000000000000000000000000000",
					"allowance(address,address)(uint256)",
					[]interface{}{
						"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
						"0xDBC05B1ECB4FDAEF943819C0B04E9EF6DF4BABD6",
					},
				),
				multicall.NewViewCall(
					"Balance/0x0000000000000000000000000000000000000001",
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
					"Decimals/0x0000",
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
					"Balance/0x0000",
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
					"DoSomething/0x0000",
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

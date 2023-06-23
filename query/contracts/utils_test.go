package query

import (
	"canto-api/multicall"
	"encoding/hex"
	"fmt"
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
					"name",
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
					"name",
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
					"name",
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

package config

import (
	"reflect"
	"testing"
)

func TestGetContractsFromJson(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		want    []Contract
		wantErr bool
	}{
		{
			name: "one contract, one function, with no arguments",
			args: "./jsons/tests/contract_test_01.json",
			want: []Contract{
				{
					Name:    "ccanto token",
					Address: "0xB65Ec550ff356EcA6150F733bA9B954b2e0Ca488",
					Keys: []string{
						"cash",
					},
					Methods: []string{
						"getCash()(uint256)",
					},
					Args: [][]interface{}{
						{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "multiple contract, multiple function, with multiple arguments",
			args: "./jsons/tests/contract_test_02.json",
			want: []Contract{
				{
					Name:    "pricefeed",
					Address: "0xa252eEE9BDe830Ca4793F054B506587027825a8e",
					Keys: []string{
						"price",
					},
					Methods: []string{
						"getUnderlyingPrice(address)(uint256)",
					},
					Args: [][]interface{}{
						{
							"0x830b9849e7d79b92408a86a557e7baaacbec6030",
						},
					},
				},
				{
					Name:    "comptroller",
					Address: "0x5E23dC409Fc2F832f83CEc191E245A191a4bCc5C",
					Keys: []string{
						"markets",
						"supplyspeeds",
						"borrowcaps",
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getContractsFromJson(tt.args)
			if (err != nil) && tt.wantErr {
				t.Errorf("getContractsFromJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) != tt.wantErr {
				t.Errorf("getContractsFromJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

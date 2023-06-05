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
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcessContractCalls(tt.args.contracts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessContractCalls() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessContractCalls() = %v, want %v", got, tt.want)
			}
		})
	}
}

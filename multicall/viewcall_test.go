package multicall

import (
	"reflect"
	"testing"
)

func TestViewCall_argumentTypes(t *testing.T) {
	type fields struct {
		target    string
		method    string
		arguments []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "no arguments",
			fields: fields{
				target:    "0x0000",
				method:    "decimals()",
				arguments: []interface{}{},
			},
			want: []string{},
		},
		{
			name: "one argument",
			fields: fields{
				target:    "0x0000",
				method:    "balanceOf(address)",
				arguments: []interface{}{"0x71C7656EC7ab88b098defB751B7401B5f6d8976F"},
			},
			want: []string{"address"},
		},
		{
			name: "multiple arguments",
			fields: fields{
				target:    "0x0000",
				method:    "doSomething(address,address,uint256,string)",
				arguments: []interface{}{"0x71C7656EC7ab88b098defB751B7401B5f6d8976F"},
			},
			want: []string{
				"address",
				"address",
				"uint256",
				"string",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := ViewCall{
				target:    tt.fields.target,
				method:    tt.fields.method,
				arguments: tt.fields.arguments,
			}
			if got := call.argumentTypes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ViewCall.argumentTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestViewCall_methodCallData(t *testing.T) {
	type fields struct {
		target    string
		method    string
		arguments []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "no arguments",
			fields: fields{
				target:    "0x0000",
				method:    "decimals()",
				arguments: []interface{}{},
			},
			want:    []byte{0x31, 0x3c, 0xe5, 0x67},
			wantErr: false,
		},
		{
			name: "multiple arguments",
			fields: fields{
				target:    "0x0000",
				method:    "doSomething(address,address,uint256,string)",
				arguments: []interface{}{},
			},
			want:    []byte{0x40, 0xb3, 0xc2, 0x1e},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := ViewCall{
				target:    tt.fields.target,
				method:    tt.fields.method,
				arguments: tt.fields.arguments,
			}
			got, err := call.methodCallData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ViewCall.methodCallData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ViewCall.methodCallData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestViewCall_argsCallData(t *testing.T) {
	type fields struct {
		target    string
		method    string
		arguments []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "no arguments",
			fields: fields{
				target:    "0x0000",
				method:    "decimals()",
				arguments: []interface{}{},
			},
			want:    []byte(nil),
			wantErr: false,
		},
		{
			name: "one argument",
			fields: fields{
				target: "0x0000",
				method: "balanceOf(address)",
				arguments: []interface{}{
					"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
				},
			},
			want: []byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x71, 0xc7, 0x65, 0x6e, 0xc7, 0xab, 0x88, 0xb0, 0x98, 0xde, 0xfb, 0x75, 0x1b, 0x74, 0x01, 0xb5, 0xf6, 0xd8, 0x97, 0x6f,
			},
			wantErr: false,
		},
		{
			name: "multiples arguments",
			fields: fields{
				target: "0x0000",
				method: "doSomething(address,int256)",
				arguments: []interface{}{
					"0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
					"18",
				},
			},
			want: []byte{
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x71, 0xc7, 0x65, 0x6e, 0xc7, 0xab, 0x88, 0xb0, 0x98, 0xde, 0xfb, 0x75, 0x1b, 0x74, 0x01, 0xb5, 0xf6, 0xd8, 0x97, 0x6f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call := ViewCall{
				target:    tt.fields.target,
				method:    tt.fields.method,
				arguments: tt.fields.arguments,
			}
			got, err := call.argsCallData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ViewCall.argsCallData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ViewCall.argsCallData() = %v, want %v", got, tt.want)
			}
		})
	}
}

package bic

import (
	"testing"
)

func TestBusinessIdentifierCode_String(t *testing.T) {
	tests := []struct {
		name string
		b    BankIdentifierCode
		want string
	}{
		{
			name: "ok",
			b:    "123",
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("String() = %v, wantErr %v", got, tt.want)
			}
		})
	}
}

func TestBusinessIdentifierCode_Validate(t *testing.T) {
	tests := []struct {
		code    BankIdentifierCode
		wantErr bool
	}{
		{code: "GEBABEBB", wantErr: false},
		{code: "BKAUATWW", wantErr: false},
		{code: "UCJAES2MXXX", wantErr: false},
		{code: "PBNKDEFF100", wantErr: false},
		{code: "", wantErr: true},
		{code: "   ", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.code.String(), func(t *testing.T) {
			if got := tt.code.Validate(); (got != nil) != tt.wantErr {
				t.Errorf("\ngot  :%v\nwantErr :%v", got, tt.wantErr)
			}
		})
	}
}

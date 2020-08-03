package bic

import "testing"

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
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBusinessIdentifierCode_Validate(t *testing.T) {
	tests := []struct {
		code BankIdentifierCode
		want bool
	}{
		{code: "GBDSC", want: false},
		{code: "MONZ", want: false},
		{code: "SRLG", want: false},
		{code: "ATMB", want: false},
		{code: "", want: true},
		{code: "111111", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.code.String(), func(t *testing.T) {
			if got := tt.code.Validate(); (got != nil) != tt.want {
				t.Errorf("\ngot  :%v\nwant :%v", got, tt.want)
			}
		})
	}
}

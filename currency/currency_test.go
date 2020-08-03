package currency

import (
	"reflect"
	"testing"
)

func TestCurrency_Validate(t *testing.T) {
	tests := []struct {
		name    string
		c       Currency
		wantErr bool
	}{
		{
			name:    "GBP",
			c:       PoundSterling,
			wantErr: false,
		},
		{
			name:    "Non-Enrolled Currency",
			c:       Currency("SOMETHING"),
			wantErr: true,
		},
		{
			name:    "nil Currency",
			c:       Currency(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("\ngot  :%v\nwant :%v", err, tt.wantErr)
			}
		})
	}
}

func TestCurrency_String(t *testing.T) {
	tests := []struct {
		name string
		c    Currency
		want string
	}{
		{
			name: "OK",
			c:    PoundSterling,
			want: "GBP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("\ngot  :%v\nwant :%v", got, tt.want)
			}
		})
	}
}

func TestCurrency_getAllCurrencies(t *testing.T) {
	t.Run("allCurrencies", func(t *testing.T) {
		expected := Currencies{PoundSterling}
		if got := getAllCurrencies(); !reflect.DeepEqual(got, expected) {
			t.Errorf("\ngot  :%v\nwant :%v", got, expected)
		}
	})
}

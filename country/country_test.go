package country

import (
	"testing"
)

func TestCountry_String(t *testing.T) {
	tests := []struct {
		name string
		c    Country
		want string
	}{
		{
			name: "ok uk",
			c:    UnitedKingdomofGreatBritainandNorthernIrelandthe,
			want: "GB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountry_Validate(t *testing.T) {
	tests := []struct {
		name    string
		c       Country
		wantErr bool
	}{
		{
			name:    "ok uk",
			c:       UnitedKingdomofGreatBritainandNorthernIrelandthe,
			wantErr: false,
		},
		{
			name:    "invalidCountry",
			c:       Country("CountryNotInTheList"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getAllCountries(t *testing.T) {
	wantLen := 249
	if got := getAllCountries(); len(got) != wantLen {
		t.Errorf("getAllCountries() len = %v, wantLen %v", len(got), wantLen)
	}
}

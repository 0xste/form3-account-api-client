package account_type

import "testing"

func TestAccountType_String(t *testing.T) {
	tests := []struct {
		name string
		a    AccountType
		want string
	}{
		{
			name: "ok",
			a:    TypeAccount,
			want: "accounts",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountType_Validate(t *testing.T) {
	tests := []struct {
		name    string
		a       AccountType
		wantErr bool
	}{
		{
			name:    "ok",
			a:       TypeAccount,
			wantErr: false,
		},
		{
			name:    "invalid account type",
			a:       "SomethingElse",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

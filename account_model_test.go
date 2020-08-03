package accountapi_client

import (
	"form3-accountapi-client/account_type"
	"form3-accountapi-client/uuid"
	"testing"
	"time"
)

func TestAccount_Validate(t *testing.T) {
	type fields struct {
		AccountType      account_type.AccountType
		Id               uuid.UUID
		CreatedOn        time.Time
		ModifiedOn       time.Time
		OrganisationUnit uuid.UUID
		Version          uint64
		Attributes       AccountAttributes
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "invalid",
			fields: fields{
				AccountType:      "",
				Id:               "",
				CreatedOn:        time.Time{},
				ModifiedOn:       time.Time{},
				OrganisationUnit: "",
				Version:          0,
				Attributes: AccountAttributes{
					alternativeBankAccountNames: nil,
					bankId:                      0,
					bankIdCode:                  "",
					baseCurrency:                "",
					country:                     "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				AccountType:      tt.fields.AccountType,
				Id:               tt.fields.Id,
				CreatedOn:        tt.fields.CreatedOn,
				ModifiedOn:       tt.fields.ModifiedOn,
				OrganisationUnit: tt.fields.OrganisationUnit,
				Version:          tt.fields.Version,
				Attributes:       tt.fields.Attributes,
			}
			if err := a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
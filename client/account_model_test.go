package client

import (
	"form3-accountapi-client/account_type"
	"form3-accountapi-client/bic"
	"form3-accountapi-client/country"
	"form3-accountapi-client/currency"
	"form3-accountapi-client/uuid"
	"reflect"
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
		Version          int64
		Attributes       AccountAttributes
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid account",
			fields: fields{
				AccountType:      account_type.TypeAccount,
				Id:               uuid.MustUUID(uuid.NewV4()),
				CreatedOn:        time.Now(),
				ModifiedOn:       time.Now(),
				OrganisationUnit: uuid.MustUUID(uuid.NewV4()),
				Version:          1,
				Attributes: AccountAttributes{
					AlternativeBankAccountNames: []string{"one", "two"},
					BankId:                      "123455",
					BankIdCode:                  bic.BankIdentifierCode("BARC"),
					BaseCurrency:                currency.PoundSterling,
					Country:                     country.UnitedKingdomofGreatBritainandNorthernIrelandthe,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid account type",
			fields: fields{
				AccountType:      account_type.AccountType("business"),
				Id:               uuid.MustUUID(uuid.NewV4()),
				CreatedOn:        time.Now(),
				ModifiedOn:       time.Now(),
				OrganisationUnit: "123456",
				Version:          1,
				Attributes: AccountAttributes{
					AlternativeBankAccountNames: []string{"one", "two"},
					BankId:                      "123455",
					BankIdCode:                  bic.BankIdentifierCode("BARC"),
					BaseCurrency:                currency.PoundSterling,
					Country:                     country.UnitedKingdomofGreatBritainandNorthernIrelandthe,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid id",
			fields: fields{
				AccountType:      account_type.TypeAccount,
				Id:               "123456",
				CreatedOn:        time.Now(),
				ModifiedOn:       time.Now(),
				OrganisationUnit: "123456",
				Version:          1,
				Attributes: AccountAttributes{
					AlternativeBankAccountNames: []string{"one", "two"},
					BankId:                      "123455",
					BankIdCode:                  bic.BankIdentifierCode("BARC"),
					BaseCurrency:                currency.PoundSterling,
					Country:                     country.UnitedKingdomofGreatBritainandNorthernIrelandthe,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid BankId",
			fields: fields{
				AccountType:      account_type.TypeAccount,
				Id:               uuid.MustUUID(uuid.NewV4()),
				CreatedOn:        time.Now(),
				ModifiedOn:       time.Now(),
				OrganisationUnit: uuid.MustUUID(uuid.NewV4()),
				Version:          1,
				Attributes: AccountAttributes{
					AlternativeBankAccountNames: []string{"some account name"},
					BankId:                      "0",
					BankIdCode:                  bic.BankIdentifierCode("BARC"),
					BaseCurrency:                currency.PoundSterling,
					Country:                     country.UnitedKingdomofGreatBritainandNorthernIrelandthe,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid BaseCurrency",
			fields: fields{
				AccountType:      account_type.TypeAccount,
				Id:               uuid.MustUUID(uuid.NewV4()),
				CreatedOn:        time.Now(),
				ModifiedOn:       time.Now(),
				OrganisationUnit: uuid.MustUUID(uuid.NewV4()),
				Version:          1,
				Attributes: AccountAttributes{
					AlternativeBankAccountNames: []string{"some account name"},
					BankId:                      "12345",
					BankIdCode:                  bic.BankIdentifierCode("BARC"),
					BaseCurrency:                currency.Currency("NOTENROLLED"),
					Country:                     country.UnitedKingdomofGreatBritainandNorthernIrelandthe,
				},
			},
			wantErr: true,
		},
		{
			name: "nil BaseCurrency",
			fields: fields{
				AccountType:      account_type.TypeAccount,
				Id:               uuid.MustUUID(uuid.NewV4()),
				CreatedOn:        time.Now(),
				ModifiedOn:       time.Now(),
				OrganisationUnit: uuid.MustUUID(uuid.NewV4()),
				Version:          1,
				Attributes: AccountAttributes{
					AlternativeBankAccountNames: []string{"some account name"},
					BankId:                      "12345",
					BankIdCode:                  bic.BankIdentifierCode("BARC"),
					BaseCurrency:                currency.Currency(""),
					Country:                     country.UnitedKingdomofGreatBritainandNorthernIrelandthe,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid Country",
			fields: fields{
				AccountType:      account_type.TypeAccount,
				Id:               uuid.MustUUID(uuid.NewV4()),
				CreatedOn:        time.Now(),
				ModifiedOn:       time.Now(),
				OrganisationUnit: uuid.MustUUID(uuid.NewV4()),
				Version:          1,
				Attributes: AccountAttributes{
					AlternativeBankAccountNames: []string{"some account name"},
					BankId:                      "12345",
					BankIdCode:                  bic.BankIdentifierCode("BARC"),
					BaseCurrency:                currency.Currency("GBP"),
					Country:                     country.Country("SOMECOUNTRY"),
				},
			},
			wantErr: true,
		},
		{
			name: "invalid Country",
			fields: fields{
				AccountType:      account_type.TypeAccount,
				Id:               uuid.MustUUID(uuid.NewV4()),
				CreatedOn:        time.Now(),
				ModifiedOn:       time.Now(),
				OrganisationUnit: uuid.MustUUID(uuid.NewV4()),
				Version:          1,
				Attributes: AccountAttributes{
					AlternativeBankAccountNames: []string{"some account name"},
					BankId:                      "12345",
					BankIdCode:                  bic.BankIdentifierCode("BARC"),
					BaseCurrency:                currency.Currency("GBP"),
					Country:                     country.Country("nil"),
				},
			},
			wantErr: true,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				AccountType:    tt.fields.AccountType,
				Id:             tt.fields.Id,
				CreatedOn:      tt.fields.CreatedOn,
				ModifiedOn:     tt.fields.ModifiedOn,
				OrganisationId: tt.fields.OrganisationUnit,
				Version:        tt.fields.Version,
				Attributes:     tt.fields.Attributes,
			}
			if err := a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccount_Builder(t *testing.T) {
	someUuid := uuid.MustUUID(uuid.NewV4())
	now := time.Now()

	attr := AccountAttributes{}
	attr.
		WithAlternativeBankAccountNames([]string{"some bank"}).
		WithCountry(country.UnitedKingdomofGreatBritainandNorthernIrelandthe).
		WithBankId("1234").
		WithBic("BARC").
		WithBaseCurrency(currency.PoundSterling).
		WithCountry(country.UnitedKingdomofGreatBritainandNorthernIrelandthe)

	account := Account{}
	actual := account.
		WithAccountType(account_type.TypeAccount).
		WithId(someUuid).
		WithCreatedOn(now.AddDate(-1, 0, 0)).
		WithModifiedOn(now).
		WithOrganisationId(someUuid).
		WithVersion(1).
		WithAttributes(attr)

	expected := &Account{
		AccountType:    account_type.TypeAccount,
		Id:             someUuid,
		CreatedOn:      now.AddDate(-1, 0, 0),
		ModifiedOn:     now,
		OrganisationId: someUuid,
		Version:        1,
		Attributes:       AccountAttributes{
			AlternativeBankAccountNames: []string{"some bank"},
			BankId:                      "1234",
			BankIdCode:                  "BARC",
			BaseCurrency:                currency.PoundSterling,
			Country:                     country.UnitedKingdomofGreatBritainandNorthernIrelandthe,
		},
	}
	if ! reflect.DeepEqual(actual, expected){
		t.Errorf("\nGot  : %v\nwant :%v", actual, expected)
	}
}
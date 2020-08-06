package main

import (
	"form3-accountapi-client/account_type"
	"form3-accountapi-client/country"
	"form3-accountapi-client/currency"
	"form3-accountapi-client/uuid"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func Test_limitAndOffset(t *testing.T) {
	testAccounts := generateTestAccounts(100)
	type args struct {
		data []Account
		skip int
		size int
	}
	tests := []struct {
		name string
		args args
		want []uuid.UUID
	}{
		{
			name: "simple skip",
			args: args{
				data: testAccounts[0:9],
				skip: 0,
				size: 9,
			},
			want: getIdsForRange(testAccounts[0:9], 0, 9),
		},
		{
			name: "simple size",
			args: args{
				data:   testAccounts[0:9],
				skip:  0,
				size: 1,
			},
			want: getIdsForRange(testAccounts[0:9], 0, 1),
		},
		{
			name: "breach upper bounds entirely",
			args: args{
				data:   testAccounts[0:1],
				skip:  10,
				size: 1,
			},
			want: getIdsForRange(testAccounts[0:0], 0, 0),
		},
		{
			name: "breach upper bounds partially",
			args: args{
				data:   testAccounts[0:10],
				skip:  5,
				size: 10,
			},
			want: getIdsForRange(testAccounts[0:10], 5, 10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accounts := paginate(tt.args.data, tt.args.size, tt.args.skip)
			if len(accounts) != len(tt.want){
				t.Errorf("\ngotLen  :%d\nwantLen :%d", len(accounts), len(tt.want))
			}
			if ! reflect.DeepEqual(getIds(accounts), tt.want){
				t.Errorf("\ngot  :%v\nwant :%v", getIds(accounts), tt.want)
			}
		})
	}
}



func getIdsForRange(accounts []Account, from, to int) []uuid.UUID {
	return getIds(accounts[from:to])
}

func getIds(r []Account) []uuid.UUID {
	var ids []uuid.UUID
	for _, account := range r {
		ids = append(ids, account.Id)
	}
	return ids
}

func generateTestAccounts(nAccounts int) []Account {
	var accounts []Account
	for i := 0 ; i < nAccounts ; i ++ {
		attr := AccountAttributes{}
		attr.
			WithAlternativeBankAccountNames([]string{"some bank"}).
			WithCountry(country.UnitedKingdomofGreatBritainandNorthernIrelandthe).
			WithBankId(strconv.Itoa(rand.Int())).
			WithBic("BARC").
			WithBaseCurrency(currency.PoundSterling).
			WithCountry(country.UnitedKingdomofGreatBritainandNorthernIrelandthe)

		a := Account{}
		account := a.
			WithAccountType(account_type.TypeAccount).
			WithId(uuid.MustUUID(uuid.NewV4())).
			WithCreatedOn(time.Now().AddDate(0,0,-10)).
			WithModifiedOn(time.Now()).
			WithOrganisationId(uuid.MustUUID(uuid.NewV4())).
			WithVersion(1).
			WithAttributes(attr)

		accounts = append(accounts, *account)
	}
	return accounts
}

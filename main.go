package main

import (
	"context"
	"form3-accountapi-client/account_type"
	"form3-accountapi-client/client"
	"form3-accountapi-client/country"
	"form3-accountapi-client/currency"
	"form3-accountapi-client/uuid"
	"log"
	"time"
)

// Exemplary code demonstrating the client library for the AccountClient
func main(){
	// Initialise an instance of the account account_client
	accountClient, err := client.NewAccountClient("http", "127.0.0.1", 8080, 1)
	if err != nil{
		panic(err)
	}

	// context passed for trace id headers, and timeouts etc TODO
	ctx := context.Background()

	// health check
	err = accountClient.Check(ctx)
	if err != nil{
		panic(err)
	}

	// create account setup data
	attr := client.AccountAttributes{}
	attr.
		WithAlternativeBankAccountNames([]string{"some bank"}).
		WithCountry(country.UnitedKingdomofGreatBritainandNorthernIrelandthe).
		WithBankId("1234").
		WithBic("BARC").
		WithBaseCurrency(currency.PoundSterling).
		WithCountry(country.UnitedKingdomofGreatBritainandNorthernIrelandthe)

	a := client.Account{}
	account := a.
		WithAccountType(account_type.TypeAccount).
		WithId(uuid.MustUUID(uuid.NewV4())).
		WithCreatedOn(time.Now().AddDate(0,0,-10)).
		WithModifiedOn(time.Now()).
		WithOrganisationId(uuid.MustUUID(uuid.NewV4())).
		WithVersion(1).
		WithAttributes(attr)

	// successful create account
	accountCreated, err := accountClient.Create(ctx, account)
	if err != nil{
		panic(err)
	}
	log.Printf("account created with id %s as expected \n", accountCreated.Id)

	// duplicate account created
	_, err = accountClient.Create(ctx, account)
	if _, ok := err.(*client.ErrDuplicateAccount) ; ok {
		log.Printf("duplicate account not created with id %s as expected\n", account.Id)
	}

	// invalid account passed
	invalidAcct := client.Account{}
	invalidAcct.WithId(uuid.MustUUID(uuid.NewV4()))
	_, err = accountClient.Create(ctx, &invalidAcct)
	if _, ok := err.(*client.ErrInvalidRequest) ; ok {
		log.Printf("invalid account not created with id %s as expected\n", invalidAcct.Id)
	}

	// Get Account for non-existent
	nonExistentAccount, err := uuid.FromStringV4("ad27e265-9605-4b4b-a0e5-3003ea9cc4de")
	if err != nil{
		panic("account-id is not valid")
	}
	_, err = accountClient.Fetch(ctx, nonExistentAccount)
	if notFoundErr, ok := err.(*client.ErrAccountNotFound) ; ok {
		log.Println(notFoundErr, "as expected")
	}

	// Get Account for existent
	accountRetrieved, err := accountClient.Fetch(ctx, accountCreated.Id)
	if err != nil{
		panic(err)
	}
	log.Printf("account retrieved with id %s as expected \n", accountRetrieved.Id)

	// List all Accounts
	accountsRetrieved, err := accountClient.List(ctx, 100, 0)
	if err != nil{
		panic(err)
	}
	log.Printf("list retrieved %d accounts \n %v\n", len(accountsRetrieved), getIds(accountsRetrieved))

	// paginated accounts
	paginated, err := accountClient.List(ctx, 1, 0)
	if err != nil{
		panic(err)
	}
	if len(paginated) == 1 {
		log.Printf("paginated accounts to single record successfully %v", getIds(paginated))
	}

	// delete account that exists
	err = accountClient.Delete(ctx, accountsRetrieved[0].Id, accountsRetrieved[0].Version)
	if err != nil{
		panic(err)
	}
	log.Printf("account deleted successfully %s", accountRetrieved.Id.String())

	// delete account that doesn't exist
	err = accountClient.Delete(ctx, uuid.MustUUID(uuid.NewV4()), 1)
	if _, ok := err.(*client.ErrAccountNotFound) ; ok {
		log.Printf("account not deleted successfully %s", err)
	}
}

// helper pulls ids from a list of accounts
func getIds(r []client.Account) []uuid.UUID {
	var ids []uuid.UUID
	for _, account := range r {
		ids = append(ids, account.Id)
	}
	return ids
}

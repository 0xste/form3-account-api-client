package main

import (
	"context"
	"form3-accountapi-client/account_type"
	"form3-accountapi-client/country"
	"form3-accountapi-client/currency"
	"form3-accountapi-client/uuid"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	envProtcol    string = "API_PROTOCOL"
	envHost       string = "API_HOST"
	envPort       string = "API_PORT"
	envVersion    string = "API_VERSION"
	fmtLineBreak     string = "-------------------------------------------------------------"
)

func main() {

	log.Println("Stefano Mantini 06/08/2020 Form3tech-oss/interview-accountapi")
	log.Println(fmtLineBreak)
	for {
		runTest()
		log.Print("\n\n\n")
		time.Sleep(10*time.Second)
	}
}

func getEnv(key string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		log.Fatalf("key %s not found", key)
	}
	return value
}
func getEnvInt(key string) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		log.Fatalf("key %s not found", key)
	}
	i, err := strconv.Atoi(value)
	if err != nil{
		log.Fatalf("key %s not integer", key)
	}
	return i
}

func ping(path string) {
	_, err := http.Get(path)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// Exemplary code demonstrating the client library for the AccountClient
func runTest() {

	// Initialise an instance of the account account_client
	accountClient, err := NewAccountClient(getEnv(envProtcol), getEnv(envHost), getEnvInt(envPort), getEnvInt(envVersion))
	if err != nil {
		log.Fatal(err)
	}

	// context passed for trace id headers, and timeouts etc TODO
	ctx := context.Background()

	// health check
	log.Println("attempting health check")
	err = accountClient.Health(ctx)
	if err != nil {
		log.Fatalf("have you started the backing services? %v", err)
	}
	log.Println("Healthcheck completed successfully")
	log.Println(fmtLineBreak)

	// create account setup data
	attr := AccountAttributes{}
	attr.
		WithAlternativeBankAccountNames([]string{"some bank"}).
		WithCountry(country.UnitedKingdomofGreatBritainandNorthernIrelandthe).
		WithBankId("1234").
		WithBic("GEBABEBB").
		WithBankIdCode("BARC").
		WithBaseCurrency(currency.PoundSterling).
		WithCountry(country.UnitedKingdomofGreatBritainandNorthernIrelandthe)

	a := Account{}
	account := a.
		WithAccountType(account_type.TypeAccount).
		WithId(uuid.MustUUID(uuid.NewV4())).
		WithCreatedOn(time.Now().AddDate(0, 0, -10)).
		WithModifiedOn(time.Now()).
		WithOrganisationId(uuid.MustUUID(uuid.NewV4())).
		WithVersion(1).
		WithAttributes(attr)

	// create accounts
	log.Printf("attempting to create account %s", account.Id)
	accountCreated, err := accountClient.Create(ctx, account)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("account created with id %s as expected \n", accountCreated.Id)
	log.Println(fmtLineBreak)
	account.Id = uuid.MustUUID(uuid.NewV4())
	log.Printf("attempting to create account 2 %s", account.Id)
	acct2, err := accountClient.Create(ctx, account)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("account created with id %s as expected \n", acct2.Id)
	log.Println(fmtLineBreak)

	// duplicate account created
	log.Printf("attempting to create duplicate account")
	_, err = accountClient.Create(ctx, account)
	if _, ok := err.(*ErrDuplicateAccount); ok {
		log.Printf("duplicate account not created with id %s as expected\n", account.Id)
		log.Println(fmtLineBreak)
	}

	// invalid/nil account passed
	invalidAcct := Account{}
	invalidAcct.WithId(uuid.MustUUID(uuid.NewV4()))
	invalidAcct.WithAttributes(AccountAttributes{Bic: "123"})
	log.Printf("attempting to create with an invalid account payload")
	_, err = accountClient.Create(ctx, &invalidAcct)
	if _, ok := err.(*ErrInvalidRequest); ok {
		log.Printf("invalid account not created with id %s as expected\n", invalidAcct.Id)
		log.Println(fmtLineBreak)
	}

	// fetch Account for non-existent
	nonExistentAccount, err := uuid.FromStringV4("ad27e265-9605-4b4b-a0e5-3003ea9cc4de")
	log.Printf("attempting to fetch with an inexistent account %s", nonExistentAccount)
	if err != nil {
		log.Fatalf("account-id is not valid")
	}
	_, err = accountClient.Fetch(ctx, nonExistentAccount)
	if notFoundErr, ok := err.(*ErrAccountNotFound); ok {
		log.Println(notFoundErr, "as expected")
		log.Println(fmtLineBreak)
	}

	// Get Account for existent
	log.Printf("attempting to fetch with an existing account %s", accountCreated.Id)
	accountRetrieved, err := accountClient.Fetch(ctx, accountCreated.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("account retrieved with id %s as expected \n", accountRetrieved.Id)
	log.Println(fmtLineBreak)

	// List all Accounts
	log.Printf("attempting to list all accounts")
	accountsRetrieved, err := accountClient.List(ctx, 100, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("list retrieved %d accounts \n", len(accountsRetrieved))
	log.Println(fmtLineBreak)

	// simple limit/offset for paginated accounts
	log.Printf("attempting to list account subset of all %d accounts, with limit %d and offset %d", len(accountsRetrieved), 1, 1)
	paginated, err := accountClient.List(ctx, 1, 1)
	if err != nil {
		log.Fatal(err)
	}
	if len(paginated) == 1 {
		log.Printf("paginated %d accounts to single record successfully", len(accountsRetrieved))
		log.Println(fmtLineBreak)
	}

	// delete account that doesn't exist
	err = accountClient.Delete(ctx, uuid.MustUUID(uuid.NewV4()), 1)
	if _, ok := err.(*ErrAccountNotFound); ok {
		log.Printf("account not deleted successfully %s", err)
		log.Println(fmtLineBreak)
	}

	// delete found accounts
	log.Printf("Deleting all returned accounts %d", len(accountsRetrieved))
	for _, acct := range accountsRetrieved {
		// delete account that exists
		err = accountClient.Delete(ctx, acct.Id, acct.Version)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("account deleted successfully %s", acct.Id)
		log.Println(fmtLineBreak)
	}
}

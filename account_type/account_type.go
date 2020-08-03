package account_type

type AccountTypes []AccountType
type AccountType string

const TypeAccount AccountType = "accounts"

func (a *AccountType) String() string {
	return string(*a)
}

func (a *AccountType) Validate() error {
	for _, accountType := range getAccountTypes() {
		if accountType == *a {
			return nil
		}
	}
	return &ErrInvalidAccountType{*a}
}

func getAccountTypes() AccountTypes {
	return AccountTypes{TypeAccount}
}

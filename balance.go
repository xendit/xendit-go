package xendit

// BalanceAccountTypeEnum constants are the available balance account type
type BalanceAccountTypeEnum string

// This consists the values that BalanceAccountType can take
const (
	BalanceAccountTypeCash    BalanceAccountTypeEnum = "CASH"
	BalanceAccountTypeHolding BalanceAccountTypeEnum = "HOLDING"
	BalanceAccountTypeTax     BalanceAccountTypeEnum = "TAX"
)

// Balance contains data from Xendit's API response of balance related requests.
// For more details see https://xendit.github.io/apireference/?bash#balances.
type Balance struct {
	Balance float64 `json:"balance"`
}

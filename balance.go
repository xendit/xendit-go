package xendit

// BalanceAccountTypeEnum constants are the available e-wallet type
type BalanceAccountTypeEnum string

// This consists the values that EWalletTypeEnum can take
const (
	BalanceAccountTypeCash    BalanceAccountTypeEnum = "CASH"
	BalanceAccountTypeHolding BalanceAccountTypeEnum = "HOLDING"
	BalanceAccountTypeTax     BalanceAccountTypeEnum = "TAX"
)

// String returns the BalanceAccountTypeEnum in type string
func (b *BalanceAccountTypeEnum) String() string {
	return string(*b)
}

// Balance contains data from Xendit's API response of balance related request.
// For more details see https://xendit.github.io/apireference/?bash#balances.
type Balance struct {
	Balance float64 `json:"balance"`
}

package xendit

// Balance contains data from Xendit's API response of balance related request.
// For more details see https://xendit.github.io/apireference/?bash#balances.
type Balance struct {
	Balance float64 `json:"balance"`
}

package disbursement

type DisbursementCallback struct {
	ID                      string  `json:"id"`
	Created                 string  `json:"created"`
	Updated                 string  `json:"updated"`
	ExternalID              string  `json:"external_id"`
	UserID                  string  `json:"user_id"`
	Amount                  float64 `json:"amount"`
	BankCode                string  `json:"bank_code"`
	AccountHolderName       string  `json:"account_holder_name"`
	DisbursementDescription string  `json:"disbursement_description"`
	Status                  string  `json:"status"`
	FailureCode             string  `json:"failure_code,omitempty"`
	IsInstant               bool    `json:"is_instant,omitempty"`
}

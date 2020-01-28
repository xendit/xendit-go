package xendit

// Disbursement contains data from Xendit's API response of disbursement related requests.
// For more details see https://xendit.github.io/apireference/?bash#disbursement.
type Disbursement struct {
	ID                      string   `json:"id"`
	UserID                  string   `json:"user_id"`
	ExternalID              string   `json:"external_id"`
	Amount                  float64  `json:"amount"`
	BankCode                string   `json:"bank_code"`
	AccountHolderName       string   `json:"account_holder_name"`
	DisbursementDescription string   `json:"disbursement_description"`
	Status                  string   `json:"status"`
	EmailTo                 []string `json:"email_to,omitempty"`
	EmailCC                 []string `json:"email_cc,omitempty"`
	EmailBCC                []string `json:"email_bcc,omitempty"`
	IsInstant               *bool    `json:"is_instant,omitempty"`
	FailureCode             string   `json:"failure_code,omitempty"`
}

// DisbursementBank contains data from Xendit's API response of Get Disbursement Banks.
type DisbursementBank struct {
	Name            string `json:"name"`
	Code            string `json:"code"`
	CanDisburse     *bool  `json:"can_disburse"`
	CanNameValidate *bool  `json:"can_name_validate"`
}

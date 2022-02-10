package xendit

import "time"

// Disbursement contains data from Xendit's API response of disbursement related requests.
// For more details see https://xendit.github.io/apireference/?bash#disbursement.
// For documentation of subpackage disbursement, checkout https://pkg.go.dev/github.com/xendit/xendit-go/disbursement
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
	IsInstant               bool     `json:"is_instant,omitempty"`
	FailureCode             string   `json:"failure_code,omitempty"`
}

// DisbursementBank contains data from Xendit's API response of Get Disbursement Banks.
type DisbursementBank struct {
	Name            string `json:"name"`
	Code            string `json:"code"`
	CanDisburse     bool   `json:"can_disburse"`
	CanNameValidate bool   `json:"can_name_validate"`
}

// BatchDisbursement contains data from Xendit's API response of batch disbursement.
// For more details see https://xendit.github.io/apireference/?bash#batch-disbursement.
type BatchDisbursement struct {
	Created             *time.Time `json:"created"`
	Reference           string     `json:"reference"`
	TotalUploadedAmount float64    `json:"total_uploaded_amount"`
	TotalUploadedCount  int        `json:"total_uploaded_count"`
	Status              string     `json:"status"`
	ID                  string     `json:"id"`
}

// Disbursement contains data from Xendit's API response of disbursement related requests.
// For more details see https://docs.google.com/document/d/1eK7rt6AwMZHcAN1wxgqhsesQ_im35t2Q9iuM7bOXI04/edit#.
// For documentation of subpackage disbursement, checkout https://pkg.go.dev/github.com/xendit/xendit-go/disbursement
type DisbursementPh struct {
	ID                  string                 `json:"id"`
	ReferenceID         string                 `json:"reference_id"`
	Currency            string                 `json:"currency"`
	Amount              float64                `json:"amount"`
	ChannelCode         string                 `json:"channel_code"`
	Description         string                 `json:"description"`
	Status              string                 `json:"status"`
	Created             *time.Time             `json:"created,omitempty"`
	Updated             *time.Time             `json:"updated,omitempty"`
	ReceiptNotification ReceiptNotification    `json:"receipt_notification,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
	FailureCode         string                 `json:"failure_code,omitempty"`
}

// Beneficiary is data that contained in Create at Beneficiary
type Beneficiary struct {
	Type         string `json:"type"`
	GivenNames   string `json:"given_names"`
	MiddleName   string `json:"middle_name"`
	Surname      string `json:"surname"`
	BusinessName string `json:"business_name"`
	StreetLine1  string `json:"street_line1"`
	StreetLine2  string `json:"street_line2"`
	City         string `json:"city"`
	Province     string `json:"province"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
	MobileNumber string `json:"mobile_number"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
}

// ReceiptNotification is data that contained in Create at ReceiptNotification
type ReceiptNotification struct {
	EmailTo  []string `json:"email_to,omitempty"`
	EmailCC  []string `json:"email_cc,omitempty"`
	EmailBCC []string `json:"email_bcc,omitempty"`
}

// DisbursementChannel contains data from Xendit's API response of Get Disbursement Channels.
type DisbursementChannel struct {
	ChannelCode     string       `json:"channel_code"`
	Name            string       `json:"name"`
	ChannelCategory string       `json:"channel_category"`
	AmountLimits    AmountLimits `json:"amount_limits"`
	Currency        string       `json:"currency"`
}

// AmountLimits is data that contained in DisbursementChannel at Amount Limit.
type AmountLimits struct {
	Minimum          float64 `json:"minimum"`
	Maximum          float64 `json:"maximum"`
	MinimumIncrement float64 `json:"minimum_increment"`
}

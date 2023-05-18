package invoice

import "github.com/xendit/xendit-go"

type InvoiceCallback struct {
	ID                 string                       `json:"id"`
	Fees               []xendit.InvoiceFee          `json:"fees"`
	Amount             int                          `json:"amount"`
	Status             string                       `json:"status"`
	Created            string                       `json:"created"`
	IsHigh             bool                         `json:"is_high"`
	PaidAt             string                       `json:"paid_at"`
	Updated            string                       `json:"updated"`
	UserID             string                       `json:"user_id"`
	Currency           string                       `json:"currency"`
	PaymentID          string                       `json:"payment_id"`
	ExternalID         string                       `json:"external_id"`
	PaidAmount         int                          `json:"paid_amount"`
	MerchantName       string                       `json:"merchant_name"`
	PaymentMethod      string                       `json:"payment_method"`
	PaymentChannel     string                       `json:"payment_channel"`
	PaymentDetails     *xendit.InvoicePaymentDetail `json:"payment_details"`
	PaymentMethodID    string                       `json:"payment_method_id"`
	BankCode           string                       `json:"bank_code,omitempty"`
	PayerEmail         string                       `json:"payer_email,omitempty"`
	Description        string                       `json:"description,omitempty"`
	PaymentDestination string                       `json:"payment_destination,omitempty"`
	SuccessRedirectURL string                       `json:"success_redirect_url,omitempty"`
	FailureRedirectURL string                       `json:"failure_redirect_url,omitempty"`
}

package invoice

import "github.com/xendit/xendit-go"

type InvoiceCallback struct {
	ID                           string                       `json:"id"`
	UserID                       string                       `json:"user_id"`
	ExternalID                   string                       `json:"external_id"`
	IsHigh                       bool                         `json:"is_high"`
	Status                       string                       `json:"status"`
	MerchantName                 string                       `json:"merchant_name"`
	Amount                       float64                      `json:"amount"`
	Created                      string                       `json:"created"`
	Updated                      string                       `json:"updated"`
	PayerEmail                   string                       `json:"payer_email,omitempty"`
	Description                  string                       `json:"description,omitempty"`
	PaymentID                    string                       `json:"payment_id,omitempty"`
	PaidAmount                   float64                      `json:"paid_amount,omitempty"`
	PaymentMethod                string                       `json:"payment_method,omitempty"`
	BankCode                     string                       `json:"bank_code,omitempty"`
	EwalletType                  string                       `json:"ewallet_type,omitempty"`
	CallbackAuthenticationToken  string                       `json:"callback_authentication_token,omitempty"`
	CallbackURL                  string                       `json:"callback_url,omitempty"`
	BusinessID                   string                       `json:"business_id,omitempty"`
	BusinessName                 string                       `json:"business_name,omitempty"`
	OnDemandLink                 string                       `json:"on_demand_link,omitempty"`
	RecurringPaymentID           string                       `json:"recurring_payment_id,omitempty"`
	CreditCardChargeID           string                       `json:"credit_card_charge_id,omitempty"`
	Currency                     string                       `json:"currency,omitempty"`
	InitialCurrency              string                       `json:"initial_currency,omitempty"`
	InitialAmount                float64                      `json:"initial_amount,omitempty"`
	PaidAt                       *string                      `json:"paid_at,omitempty"`
	MidLabel                     string                       `json:"mid_label,omitempty"`
	PaymentChannel               string                       `json:"payment_channel,omitempty"`
	PaymentDestination           string                       `json:"payment_destination,omitempty"`
	SuccessRedirectURL           string                       `json:"success_redirect_url,omitempty"`
	FailureRedirectURL           string                       `json:"failure_redirect_url,omitempty"`
	CreditCardToken              string                       `json:"credit_card_token,omitempty"`
	PaymentMethodID              string                       `json:"payment_method_id,omitempty"`
	ShouldAuthenticateCreditCard bool                         `json:"should_authenticate_credit_card,omitempty"`
	PaymentDetails               *xendit.InvoicePaymentDetail `json:"payment_details,omitempty"`
	Fees                         *[]xendit.InvoiceFee         `json:"fees,omitempty"`
	Items                        *[]xendit.InvoiceItem        `json:"items,omitempty"`
}

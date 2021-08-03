package directdebitpayment

import (
	"net/url"

	"github.com/xendit/xendit-go"
)

// CreateDirectDebitPaymentParams contains parameters for CreateDirectDebitPayment
type CreateDirectDebitPaymentParams struct {
	IdempotencyKey     string                         `json:"-"`
	ForUserID          string                         `json:"-"`
	ReferenceID        string                         `json:"reference_id" validate:"required"`
	PaymentMethodID    string                         `json:"payment_method_id" validate:"required"`
	Currency           string                         `json:"currency" validate:"required"`
	Amount             float64                        `json:"amount" validate:"required"`
	CallbackURL        string                         `json:"callback_url,omitempty"`
	EnableOTP          bool                           `json:"enable_otp,omitempty"`
	Description        string                         `json:"description,omitempty"`
	Device             xendit.DirectDebitDevice       `json:"device,omitempty"`
	Basket             []xendit.DirectDebitBasketItem `json:"basket,omitempty"`
	SuccessRedirectURL string                         `json:"success_redirect_url,omitempty"`
	FailureRedirectURL string                         `json:"failure_redirect_url,omitempty"`
	Metadata           map[string]interface{}         `json:"metadata,omitempty"`
}

// ValidateOTPForDirectDebitPaymentParams contains parameters for ValidateOTPForDirectDebitPayment
type ValidateOTPForDirectDebitPaymentParams struct {
	ForUserID     string `json:"-"`
	DirectDebitID string `json:"direct_debit_id" validate:"required"`
	OTPCode       string `json:"otp_code" validate:"required"`
}

// GetDirectDebitPaymentStatusByIDParams contains parameters for GetDirectDebitPaymentStatusByID
type GetDirectDebitPaymentStatusByIDParams struct {
	ForUserID string `json:"-"`
	ID        string `json:"id" validate:"required"`
}

// GetDirectDebitPaymentStatusByReferenceIDParams contains parameters for GetDirectDebitPaymentStatusByReferenceID
type GetDirectDebitPaymentStatusByReferenceIDParams struct {
	ForUserID   string `json:"-"`
	ReferenceID string `json:"reference_id" validate:"required"`
}

// QueryString creates query string from GetDirectDebitPaymentStatusByReferenceIDParams, ignores nil values
func (p *GetDirectDebitPaymentStatusByReferenceIDParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("reference_id", p.ReferenceID)

	return urlValues.Encode()
}

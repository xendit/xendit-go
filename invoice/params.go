package invoice

import (
	"net/url"
	"strconv"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/urlvalues"
)

// CreateParams contains parameters for Create
type CreateParams struct {
	ForUserID                      string                                       `json:"-"`
	WithFeeRule                    string                                       `json:"-"`
	ExternalID                     string                                       `json:"external_id" validate:"required"`
	Amount                         float64                                      `json:"amount" validate:"required"`
	Description                    string                                       `json:"description,omitempty"`
	PayerEmail                     string                                       `json:"payer_email,omitempty"`
	ShouldSendEmail                *bool                                        `json:"should_send_email,omitempty"`
	Customer                       xendit.InvoiceCustomer                       `json:"customer,omitempty"`
	CustomerNotificationPreference xendit.InvoiceCustomerNotificationPreference `json:"customer_notification_preference,omitempty"`
	InvoiceDuration                int                                          `json:"invoice_duration,omitempty"`
	SuccessRedirectURL             string                                       `json:"success_redirect_url,omitempty"`
	FailureRedirectURL             string                                       `json:"failure_redirect_url,omitempty"`
	PaymentMethods                 []string                                     `json:"payment_methods,omitempty"`
	Currency                       string                                       `json:"currency,omitempty"`
	FixedVA                        *bool                                        `json:"fixed_va,omitempty"`
	CallbackVirtualAccountID       string                                       `json:"callback_virtual_account_id,omitempty"`
	MidLabel                       string                                       `json:"mid_label,omitempty"`
	ReminderTimeUnit               string                                       `json:"reminder_time_unit,omitempty"`
	ReminderTime                   int                                          `json:"reminder_time,omitempty"`
	Locale                         string                                       `json:"locale,omitempty"`
	Items                          []xendit.InvoiceItem                         `json:"items,omitempty"`
	Fees                           []xendit.InvoiceFee                          `json:"fees,omitempty"`
	ShouldAuthenticateCreditCard   bool                                         `json:"should_authenticate_credit_card,omitempty"`
}

// GetParams contains parameters for Get
type GetParams struct {
	ForUserID string `json:"-"`
	ID        string `json:"id" validate:"required"`
}

// GetAllParams contains parameters for GetAll
type GetAllParams struct {
	ForUserID          string    `json:"-"`
	Statuses           []string  `json:"statuses,omitempty"`
	Limit              int       `json:"limit,omitempty"`
	CreatedAfter       time.Time `json:"created_after,omitempty"`
	CreatedBefore      time.Time `json:"created_before,omitempty"`
	PaidAfter          time.Time `json:"paid_after,omitempty"`
	PaidBefore         time.Time `json:"paid_before,omitempty"`
	ExpiredAfter       time.Time `json:"expired_after,omitempty"`
	ExpiredBefore      time.Time `json:"expired_before,omitempty"`
	LastInvoiceID      string    `json:"last_invoice_id,omitempty"`
	ClientTypes        []string  `json:"client_types,omitempty"`
	PaymentChannels    []string  `json:"payment_channels,omitempty"`
	OnDemandLink       string    `json:"on_demand_link,omitempty"`
	RecurringPaymentID string    `json:"recurring_payment_id,omitempty"`
}

// QueryString creates query string from GetAllParams, ignores nil values
func (p *GetAllParams) QueryString() string {
	urlValues := &url.Values{}

	urlvalues.AddStringSliceToURLValues(urlValues, p.Statuses, "statuses")
	if p.Limit > 0 {
		urlValues.Add("limit", strconv.Itoa(p.Limit))
	}
	urlvalues.AddTimeToURLValues(urlValues, p.CreatedAfter, "created_after")
	urlvalues.AddTimeToURLValues(urlValues, p.CreatedBefore, "created_before")
	urlvalues.AddTimeToURLValues(urlValues, p.PaidAfter, "paid_after")
	urlvalues.AddTimeToURLValues(urlValues, p.PaidBefore, "paid_before")
	urlvalues.AddTimeToURLValues(urlValues, p.PaidBefore, "paid_before")
	urlvalues.AddTimeToURLValues(urlValues, p.ExpiredAfter, "expired_after")
	urlvalues.AddTimeToURLValues(urlValues, p.ExpiredBefore, "expired_before")
	urlvalues.AddStringSliceToURLValues(urlValues, p.ClientTypes, "client_types")
	urlvalues.AddStringSliceToURLValues(urlValues, p.PaymentChannels, "payment_channels")
	if p.OnDemandLink != "" {
		urlValues.Add("on_demand", p.OnDemandLink)
	}
	if p.RecurringPaymentID != "" {
		urlValues.Add("recurring_payment_id", p.RecurringPaymentID)
	}

	return urlValues.Encode()
}

// ExpireParams contains parameters for Expire
type ExpireParams struct {
	ID        string `json:"id" validate:"required"`
	ForUserID string `json:"-"`
}

package form

// CreateInvoiceData is data that needed to CreateInvoice
type CreateInvoiceData struct {
	ExternalID               string   `json:"external_id"`
	PayerEmail               string   `json:"payer_email"`
	Description              string   `json:"description"`
	Amount                   int64    `json:"amount"`
	ShouldSendEmail          bool     `json:"should_send_email,omitempty"`
	CallbackVirtualAccountID string   `json:"callback_virtual_account_id,omitempty"`
	InvoiceDuration          int      `json:"invoice_duration,omitempty"`
	SuccessRedirectURL       string   `json:"success_redirect_url,omitempty"`
	FailureRedirectURL       string   `json:"failure_redirect_url,omitempty"`
	PaymentMethods           []string `json:"payment_methods,omitempty"`
	MidLabel                 string   `json:"mid_label,omitempty"`
	Currency                 string   `json:"currency,omitempty"`
	FixedVA                  bool     `json:"fixed_va,omitempty"`
}

// GetAllInvoicesData is data that needed to GetAllInvoices
type GetAllInvoicesData struct {
	Statuses           []string `json:"statuses,omitempty"`
	Limit              int      `json:"limit,omitempty"`
	CreatedAfter       string   `json:"created_after,omitempty"`
	CreatedBefore      string   `json:"created_before,omitempty"`
	PaidAfter          string   `json:"paid_after,omitempty"`
	PaidBefore         string   `json:"paid_before,omitempty"`
	ExpiredAfter       string   `json:"expired_after,omitempty"`
	ExpiredBefore      string   `json:"expired_before,omitempty"`
	LastInvoiceID      string   `json:"last_invoice_id,omitempty"`
	ClientTypes        []string `json:"client_types,omitempty"`
	PaymentChannels    []string `json:"payment_channels,omitempty"`
	OnDemandLink       string   `json:"on_demand_link,omitempty"`
	RecurringPaymentID string   `json:"recurring_payment_id,omitempty"`
}

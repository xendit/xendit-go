package xendit

import (
	"net/url"
	"strconv"
	"time"

	"github.com/xendit/xendit-go/utils"
)

// Invoice is data that contained in API response of invoice related request
type Invoice struct {
	ID                        string                `json:"id"`
	ExternalID                string                `json:"external_id"`
	UserID                    string                `json:"user_id"`
	PayerEmail                string                `json:"payer_email"`
	Description               string                `json:"description"`
	Amount                    float64               `json:"amount"`
	MerchantName              string                `json:"merchant_name"`
	MerchantProfilePictureURL string                `json:"merchant_profile_picture_url"`
	InvoiceURL                string                `json:"invoice_url"`
	ExpiryDate                string                `json:"expiry_date"`
	AvailableBanks            []InvoiceBank         `json:"available_banks,omitempty"`
	AvailableEWallets         []InvoiceEWallet      `json:"available_ewallets,omitempty"`
	AvailableRetailOutlets    []InvoiceRetailOutlet `json:"available_retail_outlets,omitempty"`
	ShouldExcludeCreditCard   bool                  `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool                  `json:"should_send_email"`
	Created                   string                `json:"created"`
	Updated                   string                `json:"updated"`
	BankCode                  string                `json:"bank_code,omitempty"`
	PaidAmount                string                `json:"paid_amount,omitempty"`
	AdjustedReceivedAmount    string                `json:"adjusted_received_amount,omitempty"`
	RecurringPaymentID        string                `json:"recurring_payment_id,omitempty"`
	CreditCardChargeID        string                `json:"credit_card_charge_id,omitempty"`
	Currency                  string                `json:"currency,omitempty"`
	InitialCurrency           string                `json:"initial_currency,omitempty"`
	InitialAmount             string                `json:"initial_amount,omitempty"`
	PaidAt                    string                `json:"paid_at,omitempty"`
	MidLabel                  string                `json:"mid_label,omitempty"`
	PaymentChannel            string                `json:"payment_channel,omitempty"`
	PaymentMethod             string                `json:"payment_method,omitempty"`
	PaymentDestination        string                `json:"payment_destination,omitempty"`
	SuccessRedirectURL        string                `json:"success_redirect_url,omitempty"`
	FailureRedirectURL        string                `json:"failure_redirect_url,omitempty"`
	Items                     string                `json:"items,omitempty"`
	FixedVA                   string                `json:"fixed_va,omitempty"`
	ErrorCode                 string                `json:"error_code,omitempty"`
	ErrorMessage              string                `json:"message,omitempty"`
	Errors                    interface{}           `json:"errors,omitempty"`
}

// InvoiceBank is data that contained in API response of invoice in available_banks
type InvoiceBank struct {
	BankCode          string  `json:"bank_code"`
	CollectionType    string  `json:"collection_type"`
	BankAccountNumber string  `json:"bank_account_number"`
	TransferAmount    float64 `json:"transfer_amount"`
	BankBranch        string  `json:"bank_branch"`
	AccountHolderName string  `json:"account_holder_name"`
	IdentityAmount    int     `json:"identity_amount"`
}

// InvoiceEWallet is data that contained in response at availableEWallets
type InvoiceEWallet struct {
	EWalletType string `json:"ewallet_type"`
}

// InvoiceRetailOutlet is data that contained in response at availableEWallets
type InvoiceRetailOutlet struct {
	RetailOutletName string  `json:"retail_outlet_name"`
	PaymentCode      string  `json:"payment_code"`
	TransferAmount   float64 `json:"transfer_amount"`
	MerchantName     string  `json:"merchant_name,omitempty"`
}

// InvoiceItem is data that contained in response at items
type InvoiceItem struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// CreateInvoiceParams is parameters that are needed to CreateInvoice
type CreateInvoiceParams struct {
	ExternalID               string   `json:"external_id" validate:"required"`
	PayerEmail               string   `json:"payer_email" validate:"required"`
	Description              string   `json:"description" validate:"required"`
	Amount                   float64  `json:"amount" validate:"required"`
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

// GetAllInvoicesParams is parameters that are needed to GetAllInvoices
type GetAllInvoicesParams struct {
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

// QueryString create query string from GetAllInvoicesParams, ignore nil values
func (p *GetAllInvoicesParams) QueryString() string {
	urlValues := &url.Values{}

	utils.AddStringSliceToURLValues(urlValues, p.Statuses, "statuses")
	if p.Limit > 0 {
		urlValues.Add("limit", strconv.Itoa(p.Limit))
	}
	utils.AddTimeToURLValues(urlValues, p.CreatedAfter, "created_after")
	utils.AddTimeToURLValues(urlValues, p.CreatedBefore, "created_before")
	utils.AddTimeToURLValues(urlValues, p.PaidAfter, "paid_after")
	utils.AddTimeToURLValues(urlValues, p.PaidBefore, "paid_before")
	utils.AddTimeToURLValues(urlValues, p.PaidBefore, "paid_before")
	utils.AddTimeToURLValues(urlValues, p.ExpiredAfter, "expired_after")
	utils.AddTimeToURLValues(urlValues, p.ExpiredBefore, "expired_before")
	utils.AddStringSliceToURLValues(urlValues, p.ClientTypes, "client_types")
	utils.AddStringSliceToURLValues(urlValues, p.PaymentChannels, "payment_channels")
	if p.OnDemandLink != "" {
		urlValues.Add("on_demand", p.OnDemandLink)
	}
	if p.RecurringPaymentID != "" {
		urlValues.Add("recurring_payment_id", p.RecurringPaymentID)
	}

	return urlValues.Encode()
}

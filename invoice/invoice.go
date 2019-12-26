package invoice

import (
	form "github.com/xendit/xendit-go-client/form"
	option "github.com/xendit/xendit-go-client/option"
)

// Invoice is ...
type Invoice struct {
	Opt *option.Option
}

// bank is data that contained in API response of invoice in available_banks
type bank struct {
	BankCode          string `json:"bank_code"`
	CollectionType    string `json:"collection_type"`
	BankAccountNumber string `json:"bank_account_number"`
	TransferAmount    int64  `json:"transfer_amount"`
	BankBranch        string `json:"bank_branch"`
	AccountHolderName string `json:"account_holder_name"`
	IdentityAmount    int    `json:"identity_amount"`
}

// eWallet is data that contained in response at availableEWallets
type eWallet struct {
	EWalletType string `json:"ewallet_type"`
}

// retailOutlet is data that contained in response at availableEWallets
type retailOutlet struct {
	RetailOutletName string `json:"retail_outlet_name"`
	PaymentCode      string `json:"payment_code"`
	TransferAmount   int64  `json:"transfer_amount"`
	MerchantName     string `json:"merchant_name,omitempty"`
}

//
type item struct {
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int    `json:"quantity"`
}

// Response is data that contained in API response of invoice related request
type Response struct {
	ID                        string         `json:"id"`
	ExternalID                string         `json:"external_id"`
	UserID                    string         `json:"user_id"`
	PayerEmail                string         `json:"payer_email"`
	Description               string         `json:"description"`
	Amount                    int64          `json:"amount"`
	MerchantName              string         `json:"merchant_name"`
	MerchantProfilePictureURL string         `json:"merchant_profile_picture_url"`
	InvoiceURL                string         `json:"invoice_url"`
	ExpiryDate                string         `json:"expiry_date"`
	AvailableBanks            []bank         `json:"available_banks,omitempty"`
	AvailableEWallets         []eWallet      `json:"available_ewallets,omitempty"`
	AvailableRetailOutlets    []retailOutlet `json:"available_retail_outlets,omitempty"`
	ShouldExcludeCreditCard   bool           `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool           `json:"should_send_email"`
	Created                   string         `json:"created"`
	Updated                   string         `json:"updated"`
	BankCode                  string         `json:"bank_code,omitempty"`
	PaidAmount                string         `json:"paid_amount,omitempty"`
	AdjustedReceivedAmount    string         `json:"adjusted_received_amount,omitempty"`
	RecurringPaymentID        string         `json:"recurring_payment_id,omitempty"`
	CreditCardChargeID        string         `json:"credit_card_charge_id,omitempty"`
	Currency                  string         `json:"currency,omitempty"`
	InitialCurrency           string         `json:"initial_currency,omitempty"`
	InitialAmount             string         `json:"initial_amount,omitempty"`
	PaidAt                    string         `json:"paid_at,omitempty"`
	MidLabel                  string         `json:"mid_label,omitempty"`
	PaymentChannel            string         `json:"payment_channel,omitempty"`
	PaymentMethod             string         `json:"payment_method,omitempty"`
	PaymentDestination        string         `json:"payment_destination,omitempty"`
	SuccessRedirectURL        string         `json:"success_redirect_url,omitempty"`
	FailureRedirectURL        string         `json:"failure_redirect_url,omitempty"`
	Items                     string         `json:"items,omitempty"`
	FixedVA                   string         `json:"fixed_va,omitempty"`
	ErrorCode                 string         `json:"error_code"`
}

// CreateInvoice creates new invoice
func (i Invoice) CreateInvoice(data form.CreateInvoiceData) (Response, error) {
	// TODO: validate the input data

	var response Response
	return response, nil
}

// GetInvoice gets one invoice
func (i Invoice) GetInvoice(invoiceID string) (Response, error) {
	var response Response
	return response, nil
}

// ExpireInvoice expire the created invoice
func (i Invoice) ExpireInvoice(invoiceID string) (Response, error) {
	var response Response
	return response, nil
}

// GetAllInvoices gets all invoices with conditions
func (i Invoice) GetAllInvoices(data form.GetAllInvoicesData) ([]Response, error) {
	var responses []Response
	return responses, nil
}

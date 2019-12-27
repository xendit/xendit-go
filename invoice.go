package xendit

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
	AvailableBanks            []invoiceBank         `json:"available_banks,omitempty"`
	AvailableEWallets         []invoiceEWallet      `json:"available_ewallets,omitempty"`
	AvailableRetailOutlets    []invoiceRetailOutlet `json:"available_retail_outlets,omitempty"`
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
}

// invoiceBank is data that contained in API response of invoice in available_banks
type invoiceBank struct {
	BankCode          string  `json:"bank_code"`
	CollectionType    string  `json:"collection_type"`
	BankAccountNumber string  `json:"bank_account_number"`
	TransferAmount    float64 `json:"transfer_amount"`
	BankBranch        string  `json:"bank_branch"`
	AccountHolderName string  `json:"account_holder_name"`
	IdentityAmount    int     `json:"identity_amount"`
}

// invoiceEWallet is data that contained in response at availableEWallets
type invoiceEWallet struct {
	EWalletType string `json:"ewallet_type"`
}

// invoiceRetailOutlet is data that contained in response at availableEWallets
type invoiceRetailOutlet struct {
	RetailOutletName string  `json:"retail_outlet_name"`
	PaymentCode      string  `json:"payment_code"`
	TransferAmount   float64 `json:"transfer_amount"`
	MerchantName     string  `json:"merchant_name,omitempty"`
}

// invoiceItem is data that contained in response at items
type invoiceItem struct {
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

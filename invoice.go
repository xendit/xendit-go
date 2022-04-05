package xendit

import "time"

// Invoice contains data from Xendit's API response of invoice related requests.
// For more API details see https://xendit.github.io/apireference/?bash#invoices.
// For documentation of subpackage invoice, checkout https://pkg.go.dev/github.com/xendit/xendit-go/invoice
type Invoice struct {
	ID                             string                                `json:"id"`
	InvoiceURL                     string                                `json:"invoice_url"`
	UserID                         string                                `json:"user_id,omitempty"`
	ExternalID                     string                                `json:"external_id"`
	Status                         string                                `json:"status"`
	MerchantName                   string                                `json:"merchant_name"`
	MerchantProfilePictureURL      string                                `json:"merchant_profile_picture_url,omitempty"`
	Amount                         float64                               `json:"amount"`
	Locale                         string                                `json:"locale,omitempty"`
	Items                          []InvoiceItem                         `json:"items,omitempty"`
	Fees                           []InvoiceFee                          `json:"fees,omitempty"`
	PayerEmail                     string                                `json:"payer_email,omitempty"`
	Description                    string                                `json:"description,omitempty"`
	ExpiryDate                     *time.Time                            `json:"expiry_date"`
	Customer                       InvoiceCustomer                       `json:"customer,omitempty"`
	CustomerNotificationPreference InvoiceCustomerNotificationPreference `json:"customer_notification_preference,omitempty"`
	AvailableBanks                 []InvoiceBank                         `json:"available_banks,omitempty"`
	AvailableEWallets              []InvoiceEWallet                      `json:"available_ewallets,omitempty"`
	AvailableRetailOutlets         []InvoiceRetailOutlet                 `json:"available_retail_outlets,omitempty"`
	ReminderDate                   *time.Time                            `json:"reminder_date,omitempty"`
	FixedVA                        bool                                  `json:"fixed_va,omitempty"`
	MidLabel                       string                                `json:"mid_label,omitempty"`
	ShouldExcludeCreditCard        bool                                  `json:"should_exclude_credit_card,omitempty"`
	ShouldAuthenticateCreditCard   bool                                  `json:"should_authenticate_credit_card,omitempty"`
	ShouldSendEmail                bool                                  `json:"should_send_email"`
	Created                        *time.Time                            `json:"created"`
	Updated                        *time.Time                            `json:"updated"`
	Currency                       string                                `json:"currency,omitempty"`
	PaidAt                         *time.Time                            `json:"paid_at,omitempty"`
	PaymentMethod                  string                                `json:"payment_method,omitempty"`
	PaymentChannel                 string                                `json:"payment_channel,omitempty"`
	PaymentDestination             string                                `json:"payment_destination,omitempty"`
	PaymentDetail                  InvoicePaymentDetail                  `json:"payment_details,omitempty"`
	SuccessRedirectURL             string                                `json:"success_redirect_url,omitempty"`
	FailureRedirectURL             string                                `json:"failure_redirect_url,omitempty"`
	RecurringPaymentID             string                                `json:"recurring_payment_id,omitempty"`
	CreditCardChargeID             string                                `json:"credit_card_charge_id,omitempty"`
	AdjustedReceivedAmount         float64                               `json:"adjusted_received_amount,omitempty"`
	//deprecated
	BankCode   string  `json:"bank_code,omitempty"`
	PaidAmount float64 `json:"paid_amount,omitempty"`
}

// InvoiceBank is data that contained in `Invoice` at AvailableBanks
type InvoiceBank struct {
	BankCode          string  `json:"bank_code"`
	CollectionType    string  `json:"collection_type"`
	BankAccountNumber string  `json:"bank_account_number"`
	TransferAmount    float64 `json:"transfer_amount"`
	BankBranch        string  `json:"bank_branch"`
	AccountHolderName string  `json:"account_holder_name"`
	IdentityAmount    int     `json:"identity_amount"`
}

// InvoiceEWallet is data that contained in `Invoice` at AvailableEWallets
type InvoiceEWallet struct {
	EWalletType string `json:"ewallet_type"`
}

// InvoiceRetailOutlet is data that contained in `Invoice` at AvailableRetailOutlets
type InvoiceRetailOutlet struct {
	RetailOutletName string `json:"retail_outlet_name"`
	//WILL BE DEPRECATED SOON
	PaymentCode string `json:"payment_code,omitempty"`
	//WILL BE DEPRECATED SOON
	TransferAmount float64 `json:"transfer_amount,omitempty"`
	MerchantName   string  `json:"merchant_name,omitempty"`
}

// InvoiceItem is data that contained in `Invoice` at Items
type InvoiceItem struct {
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required"`
	Quantity int     `json:"quantity" validate:"required"`
	Category string  `json:"category,omitempty"`
	Url      string  `json:"url,omitempty"`
}

// InvoiceCustomer is data that contained in `Invoice` at Customer
type InvoiceCustomer struct {
	GivenNames   string `json:"given_names,omitempty"`
	Email        string `json:"email,omitempty"`
	MobileNumber string `json:"mobile_number,omitempty"`
	Address      string `json:"address,omitempty"`
}

// InvoiceCustomerNotificationPreference is data that contained in `Invoice` at CustomerNotificationPreference
type InvoiceCustomerNotificationPreference struct {
	InvoiceCreated  []string `json:"invoice_created,omitempty"`
	InvoiceReminder []string `json:"invoice_reminder,omitempty"`
	InvoicePaid     []string `json:"invoice_paid,omitempty"`
	InvoiceExpired  []string `json:"invoice_expired,omitempty"`
}

// InvoiceFee is data that contained in `Invoice` at Fees
type InvoiceFee struct {
	Type  string  `json:"type" validate:"required"`
	Value float64 `json:"value" validate:"required"`
}

// InvoiceFInvoicePaymentDetaile is data that contained in `Invoice` at PaymentDetail
type InvoicePaymentDetail struct {
	ReceiptId string `json:"receipt_id,omitempty"`
	Source    string `json:"source,omitempty"`
}

/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.9.0
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
	"time"
)

// checks if the Invoice type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Invoice{}

// Invoice An object representing details for an invoice.
type Invoice struct {
	// The unique identifier for the invoice.
	Id *string `json:"id,omitempty"`
	// The external identifier for the invoice.
	ExternalId string `json:"external_id"`
	// The user ID associated with the invoice.
	UserId string `json:"user_id"`
	// The email address of the payer.
	PayerEmail *string `json:"payer_email,omitempty"`
	// A description of the invoice.
	Description *string `json:"description,omitempty"`
	PaymentMethod *InvoicePaymentMethod `json:"payment_method,omitempty"`
	Status InvoiceStatus `json:"status"`
	// The name of the merchant.
	MerchantName string `json:"merchant_name"`
	// The URL of the merchant's profile picture.
	MerchantProfilePictureUrl string `json:"merchant_profile_picture_url"`
	// The locale or language used for the invoice.
	Locale *string `json:"locale,omitempty"`
	// The total amount of the invoice.
	Amount float64 `json:"amount"`
	// Representing a date and time in ISO 8601 format.
	ExpiryDate time.Time `json:"expiry_date"`
	// The URL to view the invoice.
	InvoiceUrl string `json:"invoice_url"`
	// An array of available banks for payment.
	AvailableBanks []Bank `json:"available_banks"`
	// An array of available retail outlets for payment.
	AvailableRetailOutlets []RetailOutlet `json:"available_retail_outlets"`
	// An array of available e-wallets for payment.
	AvailableEwallets []Ewallet `json:"available_ewallets"`
	// An array of available QR codes for payment.
	AvailableQrCodes []QrCode `json:"available_qr_codes"`
	// An array of available direct debit options for payment.
	AvailableDirectDebits []DirectDebit `json:"available_direct_debits"`
	// An array of available pay-later options for payment.
	AvailablePaylaters []Paylater `json:"available_paylaters"`
	// Indicates whether credit card payments should be excluded.
	ShouldExcludeCreditCard *bool `json:"should_exclude_credit_card,omitempty"`
	// Indicates whether email notifications should be sent.
	ShouldSendEmail bool `json:"should_send_email"`
	// Representing a date and time in ISO 8601 format.
	Created time.Time `json:"created"`
	// Representing a date and time in ISO 8601 format.
	Updated time.Time `json:"updated"`
	// The URL to redirect to on successful payment.
	SuccessRedirectUrl *string `json:"success_redirect_url,omitempty"`
	// The URL to redirect to on payment failure.
	FailureRedirectUrl *string `json:"failure_redirect_url,omitempty"`
	// Indicates whether credit card authentication is required.
	ShouldAuthenticateCreditCard *bool `json:"should_authenticate_credit_card,omitempty"`
	Currency *InvoiceCurrency `json:"currency,omitempty"`
	// An array of items included in the invoice.
	Items []InvoiceItem `json:"items,omitempty"`
	// Indicates whether the virtual account is fixed.
	FixedVa *bool `json:"fixed_va,omitempty"`
	// Representing a date and time in ISO 8601 format.
	ReminderDate *time.Time `json:"reminder_date,omitempty"`
	Customer *CustomerObject `json:"customer,omitempty"`
	CustomerNotificationPreference *NotificationPreference `json:"customer_notification_preference,omitempty"`
	// An array of fees associated with the invoice.
	Fees []InvoiceFee `json:"fees,omitempty"`
	ChannelProperties *ChannelProperties `json:"channel_properties,omitempty"`
	// A free-format JSON for additional information that you may use. Object can be up to 50 keys, with key names up to 40 characters long and values up to 500 characters long.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewInvoice instantiates a new Invoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoice(externalId string, userId string, status InvoiceStatus, merchantName string, merchantProfilePictureUrl string, amount float64, expiryDate time.Time, invoiceUrl string, availableBanks []Bank, availableRetailOutlets []RetailOutlet, availableEwallets []Ewallet, availableQrCodes []QrCode, availableDirectDebits []DirectDebit, availablePaylaters []Paylater, shouldSendEmail bool, created time.Time, updated time.Time) *Invoice {
	this := Invoice{}
	this.ExternalId = externalId
	this.UserId = userId
	this.Status = status
	this.MerchantName = merchantName
	this.MerchantProfilePictureUrl = merchantProfilePictureUrl
	this.Amount = amount
	this.ExpiryDate = expiryDate
	this.InvoiceUrl = invoiceUrl
	this.AvailableBanks = availableBanks
	this.AvailableRetailOutlets = availableRetailOutlets
	this.AvailableEwallets = availableEwallets
	this.AvailableQrCodes = availableQrCodes
	this.AvailableDirectDebits = availableDirectDebits
	this.AvailablePaylaters = availablePaylaters
	this.ShouldSendEmail = shouldSendEmail
	this.Created = created
	this.Updated = updated
	return &this
}

// NewInvoiceWithDefaults instantiates a new Invoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceWithDefaults() *Invoice {
	this := Invoice{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Invoice) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Invoice) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Invoice) SetId(v string) {
	o.Id = &v
}

// GetExternalId returns the ExternalId field value
func (o *Invoice) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *Invoice) SetExternalId(v string) {
	o.ExternalId = v
}

// GetUserId returns the UserId field value
func (o *Invoice) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *Invoice) SetUserId(v string) {
	o.UserId = v
}

// GetPayerEmail returns the PayerEmail field value if set, zero value otherwise.
func (o *Invoice) GetPayerEmail() string {
	if o == nil || utils.IsNil(o.PayerEmail) {
		var ret string
		return ret
	}
	return *o.PayerEmail
}

// GetPayerEmailOk returns a tuple with the PayerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetPayerEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PayerEmail) {
		return nil, false
	}
	return o.PayerEmail, true
}

// HasPayerEmail returns a boolean if a field has been set.
func (o *Invoice) HasPayerEmail() bool {
	if o != nil && !utils.IsNil(o.PayerEmail) {
		return true
	}

	return false
}

// SetPayerEmail gets a reference to the given string and assigns it to the PayerEmail field.
func (o *Invoice) SetPayerEmail(v string) {
	o.PayerEmail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Invoice) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Invoice) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Invoice) SetDescription(v string) {
	o.Description = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *Invoice) GetPaymentMethod() InvoicePaymentMethod {
	if o == nil || utils.IsNil(o.PaymentMethod) {
		var ret InvoicePaymentMethod
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetPaymentMethodOk() (*InvoicePaymentMethod, bool) {
	if o == nil || utils.IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *Invoice) HasPaymentMethod() bool {
	if o != nil && !utils.IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given InvoicePaymentMethod and assigns it to the PaymentMethod field.
func (o *Invoice) SetPaymentMethod(v InvoicePaymentMethod) {
	o.PaymentMethod = &v
}

// GetStatus returns the Status field value
func (o *Invoice) GetStatus() InvoiceStatus {
	if o == nil {
		var ret InvoiceStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetStatusOk() (*InvoiceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Invoice) SetStatus(v InvoiceStatus) {
	o.Status = v
}

// GetMerchantName returns the MerchantName field value
func (o *Invoice) GetMerchantName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetMerchantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantName, true
}

// SetMerchantName sets field value
func (o *Invoice) SetMerchantName(v string) {
	o.MerchantName = v
}

// GetMerchantProfilePictureUrl returns the MerchantProfilePictureUrl field value
func (o *Invoice) GetMerchantProfilePictureUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantProfilePictureUrl
}

// GetMerchantProfilePictureUrlOk returns a tuple with the MerchantProfilePictureUrl field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetMerchantProfilePictureUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantProfilePictureUrl, true
}

// SetMerchantProfilePictureUrl sets field value
func (o *Invoice) SetMerchantProfilePictureUrl(v string) {
	o.MerchantProfilePictureUrl = v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *Invoice) GetLocale() string {
	if o == nil || utils.IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetLocaleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *Invoice) HasLocale() bool {
	if o != nil && !utils.IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *Invoice) SetLocale(v string) {
	o.Locale = &v
}

// GetAmount returns the Amount field value
func (o *Invoice) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Invoice) SetAmount(v float64) {
	o.Amount = v
}

// GetExpiryDate returns the ExpiryDate field value
func (o *Invoice) GetExpiryDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiryDate, true
}

// SetExpiryDate sets field value
func (o *Invoice) SetExpiryDate(v time.Time) {
	o.ExpiryDate = v
}

// GetInvoiceUrl returns the InvoiceUrl field value
func (o *Invoice) GetInvoiceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceUrl
}

// GetInvoiceUrlOk returns a tuple with the InvoiceUrl field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetInvoiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceUrl, true
}

// SetInvoiceUrl sets field value
func (o *Invoice) SetInvoiceUrl(v string) {
	o.InvoiceUrl = v
}

// GetAvailableBanks returns the AvailableBanks field value
func (o *Invoice) GetAvailableBanks() []Bank {
	if o == nil {
		var ret []Bank
		return ret
	}

	return o.AvailableBanks
}

// GetAvailableBanksOk returns a tuple with the AvailableBanks field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetAvailableBanksOk() ([]Bank, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableBanks, true
}

// SetAvailableBanks sets field value
func (o *Invoice) SetAvailableBanks(v []Bank) {
	o.AvailableBanks = v
}

// GetAvailableRetailOutlets returns the AvailableRetailOutlets field value
func (o *Invoice) GetAvailableRetailOutlets() []RetailOutlet {
	if o == nil {
		var ret []RetailOutlet
		return ret
	}

	return o.AvailableRetailOutlets
}

// GetAvailableRetailOutletsOk returns a tuple with the AvailableRetailOutlets field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetAvailableRetailOutletsOk() ([]RetailOutlet, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableRetailOutlets, true
}

// SetAvailableRetailOutlets sets field value
func (o *Invoice) SetAvailableRetailOutlets(v []RetailOutlet) {
	o.AvailableRetailOutlets = v
}

// GetAvailableEwallets returns the AvailableEwallets field value
func (o *Invoice) GetAvailableEwallets() []Ewallet {
	if o == nil {
		var ret []Ewallet
		return ret
	}

	return o.AvailableEwallets
}

// GetAvailableEwalletsOk returns a tuple with the AvailableEwallets field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetAvailableEwalletsOk() ([]Ewallet, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableEwallets, true
}

// SetAvailableEwallets sets field value
func (o *Invoice) SetAvailableEwallets(v []Ewallet) {
	o.AvailableEwallets = v
}

// GetAvailableQrCodes returns the AvailableQrCodes field value
func (o *Invoice) GetAvailableQrCodes() []QrCode {
	if o == nil {
		var ret []QrCode
		return ret
	}

	return o.AvailableQrCodes
}

// GetAvailableQrCodesOk returns a tuple with the AvailableQrCodes field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetAvailableQrCodesOk() ([]QrCode, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableQrCodes, true
}

// SetAvailableQrCodes sets field value
func (o *Invoice) SetAvailableQrCodes(v []QrCode) {
	o.AvailableQrCodes = v
}

// GetAvailableDirectDebits returns the AvailableDirectDebits field value
func (o *Invoice) GetAvailableDirectDebits() []DirectDebit {
	if o == nil {
		var ret []DirectDebit
		return ret
	}

	return o.AvailableDirectDebits
}

// GetAvailableDirectDebitsOk returns a tuple with the AvailableDirectDebits field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetAvailableDirectDebitsOk() ([]DirectDebit, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableDirectDebits, true
}

// SetAvailableDirectDebits sets field value
func (o *Invoice) SetAvailableDirectDebits(v []DirectDebit) {
	o.AvailableDirectDebits = v
}

// GetAvailablePaylaters returns the AvailablePaylaters field value
func (o *Invoice) GetAvailablePaylaters() []Paylater {
	if o == nil {
		var ret []Paylater
		return ret
	}

	return o.AvailablePaylaters
}

// GetAvailablePaylatersOk returns a tuple with the AvailablePaylaters field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetAvailablePaylatersOk() ([]Paylater, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailablePaylaters, true
}

// SetAvailablePaylaters sets field value
func (o *Invoice) SetAvailablePaylaters(v []Paylater) {
	o.AvailablePaylaters = v
}

// GetShouldExcludeCreditCard returns the ShouldExcludeCreditCard field value if set, zero value otherwise.
func (o *Invoice) GetShouldExcludeCreditCard() bool {
	if o == nil || utils.IsNil(o.ShouldExcludeCreditCard) {
		var ret bool
		return ret
	}
	return *o.ShouldExcludeCreditCard
}

// GetShouldExcludeCreditCardOk returns a tuple with the ShouldExcludeCreditCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetShouldExcludeCreditCardOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ShouldExcludeCreditCard) {
		return nil, false
	}
	return o.ShouldExcludeCreditCard, true
}

// HasShouldExcludeCreditCard returns a boolean if a field has been set.
func (o *Invoice) HasShouldExcludeCreditCard() bool {
	if o != nil && !utils.IsNil(o.ShouldExcludeCreditCard) {
		return true
	}

	return false
}

// SetShouldExcludeCreditCard gets a reference to the given bool and assigns it to the ShouldExcludeCreditCard field.
func (o *Invoice) SetShouldExcludeCreditCard(v bool) {
	o.ShouldExcludeCreditCard = &v
}

// GetShouldSendEmail returns the ShouldSendEmail field value
func (o *Invoice) GetShouldSendEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShouldSendEmail
}

// GetShouldSendEmailOk returns a tuple with the ShouldSendEmail field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetShouldSendEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShouldSendEmail, true
}

// SetShouldSendEmail sets field value
func (o *Invoice) SetShouldSendEmail(v bool) {
	o.ShouldSendEmail = v
}

// GetCreated returns the Created field value
func (o *Invoice) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Invoice) SetCreated(v time.Time) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *Invoice) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *Invoice) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetSuccessRedirectUrl returns the SuccessRedirectUrl field value if set, zero value otherwise.
func (o *Invoice) GetSuccessRedirectUrl() string {
	if o == nil || utils.IsNil(o.SuccessRedirectUrl) {
		var ret string
		return ret
	}
	return *o.SuccessRedirectUrl
}

// GetSuccessRedirectUrlOk returns a tuple with the SuccessRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetSuccessRedirectUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SuccessRedirectUrl) {
		return nil, false
	}
	return o.SuccessRedirectUrl, true
}

// HasSuccessRedirectUrl returns a boolean if a field has been set.
func (o *Invoice) HasSuccessRedirectUrl() bool {
	if o != nil && !utils.IsNil(o.SuccessRedirectUrl) {
		return true
	}

	return false
}

// SetSuccessRedirectUrl gets a reference to the given string and assigns it to the SuccessRedirectUrl field.
func (o *Invoice) SetSuccessRedirectUrl(v string) {
	o.SuccessRedirectUrl = &v
}

// GetFailureRedirectUrl returns the FailureRedirectUrl field value if set, zero value otherwise.
func (o *Invoice) GetFailureRedirectUrl() string {
	if o == nil || utils.IsNil(o.FailureRedirectUrl) {
		var ret string
		return ret
	}
	return *o.FailureRedirectUrl
}

// GetFailureRedirectUrlOk returns a tuple with the FailureRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetFailureRedirectUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FailureRedirectUrl) {
		return nil, false
	}
	return o.FailureRedirectUrl, true
}

// HasFailureRedirectUrl returns a boolean if a field has been set.
func (o *Invoice) HasFailureRedirectUrl() bool {
	if o != nil && !utils.IsNil(o.FailureRedirectUrl) {
		return true
	}

	return false
}

// SetFailureRedirectUrl gets a reference to the given string and assigns it to the FailureRedirectUrl field.
func (o *Invoice) SetFailureRedirectUrl(v string) {
	o.FailureRedirectUrl = &v
}

// GetShouldAuthenticateCreditCard returns the ShouldAuthenticateCreditCard field value if set, zero value otherwise.
func (o *Invoice) GetShouldAuthenticateCreditCard() bool {
	if o == nil || utils.IsNil(o.ShouldAuthenticateCreditCard) {
		var ret bool
		return ret
	}
	return *o.ShouldAuthenticateCreditCard
}

// GetShouldAuthenticateCreditCardOk returns a tuple with the ShouldAuthenticateCreditCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetShouldAuthenticateCreditCardOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ShouldAuthenticateCreditCard) {
		return nil, false
	}
	return o.ShouldAuthenticateCreditCard, true
}

// HasShouldAuthenticateCreditCard returns a boolean if a field has been set.
func (o *Invoice) HasShouldAuthenticateCreditCard() bool {
	if o != nil && !utils.IsNil(o.ShouldAuthenticateCreditCard) {
		return true
	}

	return false
}

// SetShouldAuthenticateCreditCard gets a reference to the given bool and assigns it to the ShouldAuthenticateCreditCard field.
func (o *Invoice) SetShouldAuthenticateCreditCard(v bool) {
	o.ShouldAuthenticateCreditCard = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Invoice) GetCurrency() InvoiceCurrency {
	if o == nil || utils.IsNil(o.Currency) {
		var ret InvoiceCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCurrencyOk() (*InvoiceCurrency, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Invoice) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given InvoiceCurrency and assigns it to the Currency field.
func (o *Invoice) SetCurrency(v InvoiceCurrency) {
	o.Currency = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Invoice) GetItems() []InvoiceItem {
	if o == nil || utils.IsNil(o.Items) {
		var ret []InvoiceItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetItemsOk() ([]InvoiceItem, bool) {
	if o == nil || utils.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Invoice) HasItems() bool {
	if o != nil && !utils.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InvoiceItem and assigns it to the Items field.
func (o *Invoice) SetItems(v []InvoiceItem) {
	o.Items = v
}

// GetFixedVa returns the FixedVa field value if set, zero value otherwise.
func (o *Invoice) GetFixedVa() bool {
	if o == nil || utils.IsNil(o.FixedVa) {
		var ret bool
		return ret
	}
	return *o.FixedVa
}

// GetFixedVaOk returns a tuple with the FixedVa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetFixedVaOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.FixedVa) {
		return nil, false
	}
	return o.FixedVa, true
}

// HasFixedVa returns a boolean if a field has been set.
func (o *Invoice) HasFixedVa() bool {
	if o != nil && !utils.IsNil(o.FixedVa) {
		return true
	}

	return false
}

// SetFixedVa gets a reference to the given bool and assigns it to the FixedVa field.
func (o *Invoice) SetFixedVa(v bool) {
	o.FixedVa = &v
}

// GetReminderDate returns the ReminderDate field value if set, zero value otherwise.
func (o *Invoice) GetReminderDate() time.Time {
	if o == nil || utils.IsNil(o.ReminderDate) {
		var ret time.Time
		return ret
	}
	return *o.ReminderDate
}

// GetReminderDateOk returns a tuple with the ReminderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetReminderDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ReminderDate) {
		return nil, false
	}
	return o.ReminderDate, true
}

// HasReminderDate returns a boolean if a field has been set.
func (o *Invoice) HasReminderDate() bool {
	if o != nil && !utils.IsNil(o.ReminderDate) {
		return true
	}

	return false
}

// SetReminderDate gets a reference to the given time.Time and assigns it to the ReminderDate field.
func (o *Invoice) SetReminderDate(v time.Time) {
	o.ReminderDate = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *Invoice) GetCustomer() CustomerObject {
	if o == nil || utils.IsNil(o.Customer) {
		var ret CustomerObject
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomerOk() (*CustomerObject, bool) {
	if o == nil || utils.IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *Invoice) HasCustomer() bool {
	if o != nil && !utils.IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CustomerObject and assigns it to the Customer field.
func (o *Invoice) SetCustomer(v CustomerObject) {
	o.Customer = &v
}

// GetCustomerNotificationPreference returns the CustomerNotificationPreference field value if set, zero value otherwise.
func (o *Invoice) GetCustomerNotificationPreference() NotificationPreference {
	if o == nil || utils.IsNil(o.CustomerNotificationPreference) {
		var ret NotificationPreference
		return ret
	}
	return *o.CustomerNotificationPreference
}

// GetCustomerNotificationPreferenceOk returns a tuple with the CustomerNotificationPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomerNotificationPreferenceOk() (*NotificationPreference, bool) {
	if o == nil || utils.IsNil(o.CustomerNotificationPreference) {
		return nil, false
	}
	return o.CustomerNotificationPreference, true
}

// HasCustomerNotificationPreference returns a boolean if a field has been set.
func (o *Invoice) HasCustomerNotificationPreference() bool {
	if o != nil && !utils.IsNil(o.CustomerNotificationPreference) {
		return true
	}

	return false
}

// SetCustomerNotificationPreference gets a reference to the given NotificationPreference and assigns it to the CustomerNotificationPreference field.
func (o *Invoice) SetCustomerNotificationPreference(v NotificationPreference) {
	o.CustomerNotificationPreference = &v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *Invoice) GetFees() []InvoiceFee {
	if o == nil || utils.IsNil(o.Fees) {
		var ret []InvoiceFee
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetFeesOk() ([]InvoiceFee, bool) {
	if o == nil || utils.IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *Invoice) HasFees() bool {
	if o != nil && !utils.IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []InvoiceFee and assigns it to the Fees field.
func (o *Invoice) SetFees(v []InvoiceFee) {
	o.Fees = v
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *Invoice) GetChannelProperties() ChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret ChannelProperties
		return ret
	}
	return *o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetChannelPropertiesOk() (*ChannelProperties, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *Invoice) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given ChannelProperties and assigns it to the ChannelProperties field.
func (o *Invoice) SetChannelProperties(v ChannelProperties) {
	o.ChannelProperties = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Invoice) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Invoice) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Invoice) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o Invoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["external_id"] = o.ExternalId
	toSerialize["user_id"] = o.UserId
	if !utils.IsNil(o.PayerEmail) {
		toSerialize["payer_email"] = o.PayerEmail
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	toSerialize["status"] = o.Status
	toSerialize["merchant_name"] = o.MerchantName
	toSerialize["merchant_profile_picture_url"] = o.MerchantProfilePictureUrl
	if !utils.IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	toSerialize["amount"] = o.Amount
	toSerialize["expiry_date"] = o.ExpiryDate
	toSerialize["invoice_url"] = o.InvoiceUrl
	toSerialize["available_banks"] = o.AvailableBanks
	toSerialize["available_retail_outlets"] = o.AvailableRetailOutlets
	toSerialize["available_ewallets"] = o.AvailableEwallets
	toSerialize["available_qr_codes"] = o.AvailableQrCodes
	toSerialize["available_direct_debits"] = o.AvailableDirectDebits
	toSerialize["available_paylaters"] = o.AvailablePaylaters
	if !utils.IsNil(o.ShouldExcludeCreditCard) {
		toSerialize["should_exclude_credit_card"] = o.ShouldExcludeCreditCard
	}
	toSerialize["should_send_email"] = o.ShouldSendEmail
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	if !utils.IsNil(o.SuccessRedirectUrl) {
		toSerialize["success_redirect_url"] = o.SuccessRedirectUrl
	}
	if !utils.IsNil(o.FailureRedirectUrl) {
		toSerialize["failure_redirect_url"] = o.FailureRedirectUrl
	}
	if !utils.IsNil(o.ShouldAuthenticateCreditCard) {
		toSerialize["should_authenticate_credit_card"] = o.ShouldAuthenticateCreditCard
	}
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !utils.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !utils.IsNil(o.FixedVa) {
		toSerialize["fixed_va"] = o.FixedVa
	}
	if !utils.IsNil(o.ReminderDate) {
		toSerialize["reminder_date"] = o.ReminderDate
	}
	if !utils.IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !utils.IsNil(o.CustomerNotificationPreference) {
		toSerialize["customer_notification_preference"] = o.CustomerNotificationPreference
	}
	if !utils.IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableInvoice struct {
	value *Invoice
	isSet bool
}

func (v NullableInvoice) Get() *Invoice {
	return v.value
}

func (v *NullableInvoice) Set(val *Invoice) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoice(val *Invoice) *NullableInvoice {
	return &NullableInvoice{value: val, isSet: true}
}

func (v NullableInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



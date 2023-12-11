/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.6.0
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the InvoiceCallback type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InvoiceCallback{}

// InvoiceCallback Invoice Callback Object
type InvoiceCallback struct {
	// An invoice ID generated by Xendit
	Id string `json:"id"`
	// ID of your choice (typically the unique identifier of an invoice in your system)
	ExternalId string `json:"external_id"`
	// Xendit Business ID
	UserId string `json:"user_id"`
	// The status of the invoice.
	Status string `json:"status"`
	// The name of company or website
	MerchantName string `json:"merchant_name"`
	// Nominal amount for the invoice
	Amount float64 `json:"amount"`
	// Email of the payer
	PayerEmail *string `json:"payer_email,omitempty"`
	// Description for the invoice
	Description *string `json:"description,omitempty"`
	// Total amount paid for the invoice
	PaidAmount *float64 `json:"paid_amount,omitempty"`
	// The date and time when the invoice was created.
	Created string `json:"created"`
	// The date and time when the invoice was last updated.
	Updated string `json:"updated"`
	// The currency of the invoice.
	Currency string `json:"currency"`
	// The date and time when the invoice was paid.
	PaidAt *string `json:"paid_at,omitempty"`
	// The payment method used for the invoice.
	PaymentMethod *string `json:"payment_method,omitempty"`
	// The payment channel.
	PaymentChannel *string `json:"payment_channel,omitempty"`
	// The payment destination.
	PaymentDestination *string `json:"payment_destination,omitempty"`
	PaymentDetails *PaymentDetails `json:"payment_details,omitempty"`
	// The ID of the payment.
	PaymentId *string `json:"payment_id,omitempty"`
	// The URL to redirect to on successful payment.
	SuccessRedirectUrl *string `json:"success_redirect_url,omitempty"`
	// The URL to redirect to on payment failure.
	FailureRedirectUrl *string `json:"failure_redirect_url,omitempty"`
	// The ID associated with a credit card charge (if applicable).
	CreditCardChargeId *string `json:"credit_card_charge_id,omitempty"`
	Items []InvoiceCallbackItem `json:"items,omitempty"`
	// An array of fees associated with the invoice.
	Fees []InvoiceFee `json:"fees,omitempty"`
	// Indicates whether credit card authentication is required.
	ShouldAuthenticateCreditCard *bool `json:"should_authenticate_credit_card,omitempty"`
	// The bank code for the bank details.
	BankCode *string `json:"bank_code,omitempty"`
	// The type of eWallet.
	EwalletType *string `json:"ewallet_type,omitempty"`
	// The on-demand link.
	OnDemandLink *string `json:"on_demand_link,omitempty"`
	// The ID of the recurring payment.
	RecurringPaymentId *string `json:"recurring_payment_id,omitempty"`
}

// NewInvoiceCallback instantiates a new InvoiceCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceCallback(id string, externalId string, userId string, status string, merchantName string, amount float64, created string, updated string, currency string) *InvoiceCallback {
	this := InvoiceCallback{}
	this.Id = id
	this.ExternalId = externalId
	this.UserId = userId
	this.Status = status
	this.MerchantName = merchantName
	this.Amount = amount
	this.Created = created
	this.Updated = updated
	this.Currency = currency
	return &this
}

// NewInvoiceCallbackWithDefaults instantiates a new InvoiceCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceCallbackWithDefaults() *InvoiceCallback {
	this := InvoiceCallback{}
	return &this
}

// GetId returns the Id field value
func (o *InvoiceCallback) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InvoiceCallback) SetId(v string) {
	o.Id = v
}

// GetExternalId returns the ExternalId field value
func (o *InvoiceCallback) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *InvoiceCallback) SetExternalId(v string) {
	o.ExternalId = v
}

// GetUserId returns the UserId field value
func (o *InvoiceCallback) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *InvoiceCallback) SetUserId(v string) {
	o.UserId = v
}

// GetStatus returns the Status field value
func (o *InvoiceCallback) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InvoiceCallback) SetStatus(v string) {
	o.Status = v
}

// GetMerchantName returns the MerchantName field value
func (o *InvoiceCallback) GetMerchantName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetMerchantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantName, true
}

// SetMerchantName sets field value
func (o *InvoiceCallback) SetMerchantName(v string) {
	o.MerchantName = v
}

// GetAmount returns the Amount field value
func (o *InvoiceCallback) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InvoiceCallback) SetAmount(v float64) {
	o.Amount = v
}

// GetPayerEmail returns the PayerEmail field value if set, zero value otherwise.
func (o *InvoiceCallback) GetPayerEmail() string {
	if o == nil || utils.IsNil(o.PayerEmail) {
		var ret string
		return ret
	}
	return *o.PayerEmail
}

// GetPayerEmailOk returns a tuple with the PayerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetPayerEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PayerEmail) {
		return nil, false
	}
	return o.PayerEmail, true
}

// HasPayerEmail returns a boolean if a field has been set.
func (o *InvoiceCallback) HasPayerEmail() bool {
	if o != nil && !utils.IsNil(o.PayerEmail) {
		return true
	}

	return false
}

// SetPayerEmail gets a reference to the given string and assigns it to the PayerEmail field.
func (o *InvoiceCallback) SetPayerEmail(v string) {
	o.PayerEmail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InvoiceCallback) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InvoiceCallback) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InvoiceCallback) SetDescription(v string) {
	o.Description = &v
}

// GetPaidAmount returns the PaidAmount field value if set, zero value otherwise.
func (o *InvoiceCallback) GetPaidAmount() float64 {
	if o == nil || utils.IsNil(o.PaidAmount) {
		var ret float64
		return ret
	}
	return *o.PaidAmount
}

// GetPaidAmountOk returns a tuple with the PaidAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetPaidAmountOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.PaidAmount) {
		return nil, false
	}
	return o.PaidAmount, true
}

// HasPaidAmount returns a boolean if a field has been set.
func (o *InvoiceCallback) HasPaidAmount() bool {
	if o != nil && !utils.IsNil(o.PaidAmount) {
		return true
	}

	return false
}

// SetPaidAmount gets a reference to the given float64 and assigns it to the PaidAmount field.
func (o *InvoiceCallback) SetPaidAmount(v float64) {
	o.PaidAmount = &v
}

// GetCreated returns the Created field value
func (o *InvoiceCallback) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *InvoiceCallback) SetCreated(v string) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *InvoiceCallback) GetUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *InvoiceCallback) SetUpdated(v string) {
	o.Updated = v
}

// GetCurrency returns the Currency field value
func (o *InvoiceCallback) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *InvoiceCallback) SetCurrency(v string) {
	o.Currency = v
}

// GetPaidAt returns the PaidAt field value if set, zero value otherwise.
func (o *InvoiceCallback) GetPaidAt() string {
	if o == nil || utils.IsNil(o.PaidAt) {
		var ret string
		return ret
	}
	return *o.PaidAt
}

// GetPaidAtOk returns a tuple with the PaidAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetPaidAtOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaidAt) {
		return nil, false
	}
	return o.PaidAt, true
}

// HasPaidAt returns a boolean if a field has been set.
func (o *InvoiceCallback) HasPaidAt() bool {
	if o != nil && !utils.IsNil(o.PaidAt) {
		return true
	}

	return false
}

// SetPaidAt gets a reference to the given string and assigns it to the PaidAt field.
func (o *InvoiceCallback) SetPaidAt(v string) {
	o.PaidAt = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *InvoiceCallback) GetPaymentMethod() string {
	if o == nil || utils.IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetPaymentMethodOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *InvoiceCallback) HasPaymentMethod() bool {
	if o != nil && !utils.IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *InvoiceCallback) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetPaymentChannel returns the PaymentChannel field value if set, zero value otherwise.
func (o *InvoiceCallback) GetPaymentChannel() string {
	if o == nil || utils.IsNil(o.PaymentChannel) {
		var ret string
		return ret
	}
	return *o.PaymentChannel
}

// GetPaymentChannelOk returns a tuple with the PaymentChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetPaymentChannelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentChannel) {
		return nil, false
	}
	return o.PaymentChannel, true
}

// HasPaymentChannel returns a boolean if a field has been set.
func (o *InvoiceCallback) HasPaymentChannel() bool {
	if o != nil && !utils.IsNil(o.PaymentChannel) {
		return true
	}

	return false
}

// SetPaymentChannel gets a reference to the given string and assigns it to the PaymentChannel field.
func (o *InvoiceCallback) SetPaymentChannel(v string) {
	o.PaymentChannel = &v
}

// GetPaymentDestination returns the PaymentDestination field value if set, zero value otherwise.
func (o *InvoiceCallback) GetPaymentDestination() string {
	if o == nil || utils.IsNil(o.PaymentDestination) {
		var ret string
		return ret
	}
	return *o.PaymentDestination
}

// GetPaymentDestinationOk returns a tuple with the PaymentDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetPaymentDestinationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentDestination) {
		return nil, false
	}
	return o.PaymentDestination, true
}

// HasPaymentDestination returns a boolean if a field has been set.
func (o *InvoiceCallback) HasPaymentDestination() bool {
	if o != nil && !utils.IsNil(o.PaymentDestination) {
		return true
	}

	return false
}

// SetPaymentDestination gets a reference to the given string and assigns it to the PaymentDestination field.
func (o *InvoiceCallback) SetPaymentDestination(v string) {
	o.PaymentDestination = &v
}

// GetPaymentDetails returns the PaymentDetails field value if set, zero value otherwise.
func (o *InvoiceCallback) GetPaymentDetails() PaymentDetails {
	if o == nil || utils.IsNil(o.PaymentDetails) {
		var ret PaymentDetails
		return ret
	}
	return *o.PaymentDetails
}

// GetPaymentDetailsOk returns a tuple with the PaymentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetPaymentDetailsOk() (*PaymentDetails, bool) {
	if o == nil || utils.IsNil(o.PaymentDetails) {
		return nil, false
	}
	return o.PaymentDetails, true
}

// HasPaymentDetails returns a boolean if a field has been set.
func (o *InvoiceCallback) HasPaymentDetails() bool {
	if o != nil && !utils.IsNil(o.PaymentDetails) {
		return true
	}

	return false
}

// SetPaymentDetails gets a reference to the given PaymentDetails and assigns it to the PaymentDetails field.
func (o *InvoiceCallback) SetPaymentDetails(v PaymentDetails) {
	o.PaymentDetails = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *InvoiceCallback) GetPaymentId() string {
	if o == nil || utils.IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetPaymentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *InvoiceCallback) HasPaymentId() bool {
	if o != nil && !utils.IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *InvoiceCallback) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetSuccessRedirectUrl returns the SuccessRedirectUrl field value if set, zero value otherwise.
func (o *InvoiceCallback) GetSuccessRedirectUrl() string {
	if o == nil || utils.IsNil(o.SuccessRedirectUrl) {
		var ret string
		return ret
	}
	return *o.SuccessRedirectUrl
}

// GetSuccessRedirectUrlOk returns a tuple with the SuccessRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetSuccessRedirectUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SuccessRedirectUrl) {
		return nil, false
	}
	return o.SuccessRedirectUrl, true
}

// HasSuccessRedirectUrl returns a boolean if a field has been set.
func (o *InvoiceCallback) HasSuccessRedirectUrl() bool {
	if o != nil && !utils.IsNil(o.SuccessRedirectUrl) {
		return true
	}

	return false
}

// SetSuccessRedirectUrl gets a reference to the given string and assigns it to the SuccessRedirectUrl field.
func (o *InvoiceCallback) SetSuccessRedirectUrl(v string) {
	o.SuccessRedirectUrl = &v
}

// GetFailureRedirectUrl returns the FailureRedirectUrl field value if set, zero value otherwise.
func (o *InvoiceCallback) GetFailureRedirectUrl() string {
	if o == nil || utils.IsNil(o.FailureRedirectUrl) {
		var ret string
		return ret
	}
	return *o.FailureRedirectUrl
}

// GetFailureRedirectUrlOk returns a tuple with the FailureRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetFailureRedirectUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FailureRedirectUrl) {
		return nil, false
	}
	return o.FailureRedirectUrl, true
}

// HasFailureRedirectUrl returns a boolean if a field has been set.
func (o *InvoiceCallback) HasFailureRedirectUrl() bool {
	if o != nil && !utils.IsNil(o.FailureRedirectUrl) {
		return true
	}

	return false
}

// SetFailureRedirectUrl gets a reference to the given string and assigns it to the FailureRedirectUrl field.
func (o *InvoiceCallback) SetFailureRedirectUrl(v string) {
	o.FailureRedirectUrl = &v
}

// GetCreditCardChargeId returns the CreditCardChargeId field value if set, zero value otherwise.
func (o *InvoiceCallback) GetCreditCardChargeId() string {
	if o == nil || utils.IsNil(o.CreditCardChargeId) {
		var ret string
		return ret
	}
	return *o.CreditCardChargeId
}

// GetCreditCardChargeIdOk returns a tuple with the CreditCardChargeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetCreditCardChargeIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreditCardChargeId) {
		return nil, false
	}
	return o.CreditCardChargeId, true
}

// HasCreditCardChargeId returns a boolean if a field has been set.
func (o *InvoiceCallback) HasCreditCardChargeId() bool {
	if o != nil && !utils.IsNil(o.CreditCardChargeId) {
		return true
	}

	return false
}

// SetCreditCardChargeId gets a reference to the given string and assigns it to the CreditCardChargeId field.
func (o *InvoiceCallback) SetCreditCardChargeId(v string) {
	o.CreditCardChargeId = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InvoiceCallback) GetItems() []InvoiceCallbackItem {
	if o == nil || utils.IsNil(o.Items) {
		var ret []InvoiceCallbackItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetItemsOk() ([]InvoiceCallbackItem, bool) {
	if o == nil || utils.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InvoiceCallback) HasItems() bool {
	if o != nil && !utils.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InvoiceCallbackItem and assigns it to the Items field.
func (o *InvoiceCallback) SetItems(v []InvoiceCallbackItem) {
	o.Items = v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *InvoiceCallback) GetFees() []InvoiceFee {
	if o == nil || utils.IsNil(o.Fees) {
		var ret []InvoiceFee
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetFeesOk() ([]InvoiceFee, bool) {
	if o == nil || utils.IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *InvoiceCallback) HasFees() bool {
	if o != nil && !utils.IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []InvoiceFee and assigns it to the Fees field.
func (o *InvoiceCallback) SetFees(v []InvoiceFee) {
	o.Fees = v
}

// GetShouldAuthenticateCreditCard returns the ShouldAuthenticateCreditCard field value if set, zero value otherwise.
func (o *InvoiceCallback) GetShouldAuthenticateCreditCard() bool {
	if o == nil || utils.IsNil(o.ShouldAuthenticateCreditCard) {
		var ret bool
		return ret
	}
	return *o.ShouldAuthenticateCreditCard
}

// GetShouldAuthenticateCreditCardOk returns a tuple with the ShouldAuthenticateCreditCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetShouldAuthenticateCreditCardOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ShouldAuthenticateCreditCard) {
		return nil, false
	}
	return o.ShouldAuthenticateCreditCard, true
}

// HasShouldAuthenticateCreditCard returns a boolean if a field has been set.
func (o *InvoiceCallback) HasShouldAuthenticateCreditCard() bool {
	if o != nil && !utils.IsNil(o.ShouldAuthenticateCreditCard) {
		return true
	}

	return false
}

// SetShouldAuthenticateCreditCard gets a reference to the given bool and assigns it to the ShouldAuthenticateCreditCard field.
func (o *InvoiceCallback) SetShouldAuthenticateCreditCard(v bool) {
	o.ShouldAuthenticateCreditCard = &v
}

// GetBankCode returns the BankCode field value if set, zero value otherwise.
func (o *InvoiceCallback) GetBankCode() string {
	if o == nil || utils.IsNil(o.BankCode) {
		var ret string
		return ret
	}
	return *o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetBankCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BankCode) {
		return nil, false
	}
	return o.BankCode, true
}

// HasBankCode returns a boolean if a field has been set.
func (o *InvoiceCallback) HasBankCode() bool {
	if o != nil && !utils.IsNil(o.BankCode) {
		return true
	}

	return false
}

// SetBankCode gets a reference to the given string and assigns it to the BankCode field.
func (o *InvoiceCallback) SetBankCode(v string) {
	o.BankCode = &v
}

// GetEwalletType returns the EwalletType field value if set, zero value otherwise.
func (o *InvoiceCallback) GetEwalletType() string {
	if o == nil || utils.IsNil(o.EwalletType) {
		var ret string
		return ret
	}
	return *o.EwalletType
}

// GetEwalletTypeOk returns a tuple with the EwalletType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetEwalletTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EwalletType) {
		return nil, false
	}
	return o.EwalletType, true
}

// HasEwalletType returns a boolean if a field has been set.
func (o *InvoiceCallback) HasEwalletType() bool {
	if o != nil && !utils.IsNil(o.EwalletType) {
		return true
	}

	return false
}

// SetEwalletType gets a reference to the given string and assigns it to the EwalletType field.
func (o *InvoiceCallback) SetEwalletType(v string) {
	o.EwalletType = &v
}

// GetOnDemandLink returns the OnDemandLink field value if set, zero value otherwise.
func (o *InvoiceCallback) GetOnDemandLink() string {
	if o == nil || utils.IsNil(o.OnDemandLink) {
		var ret string
		return ret
	}
	return *o.OnDemandLink
}

// GetOnDemandLinkOk returns a tuple with the OnDemandLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetOnDemandLinkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OnDemandLink) {
		return nil, false
	}
	return o.OnDemandLink, true
}

// HasOnDemandLink returns a boolean if a field has been set.
func (o *InvoiceCallback) HasOnDemandLink() bool {
	if o != nil && !utils.IsNil(o.OnDemandLink) {
		return true
	}

	return false
}

// SetOnDemandLink gets a reference to the given string and assigns it to the OnDemandLink field.
func (o *InvoiceCallback) SetOnDemandLink(v string) {
	o.OnDemandLink = &v
}

// GetRecurringPaymentId returns the RecurringPaymentId field value if set, zero value otherwise.
func (o *InvoiceCallback) GetRecurringPaymentId() string {
	if o == nil || utils.IsNil(o.RecurringPaymentId) {
		var ret string
		return ret
	}
	return *o.RecurringPaymentId
}

// GetRecurringPaymentIdOk returns a tuple with the RecurringPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallback) GetRecurringPaymentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RecurringPaymentId) {
		return nil, false
	}
	return o.RecurringPaymentId, true
}

// HasRecurringPaymentId returns a boolean if a field has been set.
func (o *InvoiceCallback) HasRecurringPaymentId() bool {
	if o != nil && !utils.IsNil(o.RecurringPaymentId) {
		return true
	}

	return false
}

// SetRecurringPaymentId gets a reference to the given string and assigns it to the RecurringPaymentId field.
func (o *InvoiceCallback) SetRecurringPaymentId(v string) {
	o.RecurringPaymentId = &v
}

func (o InvoiceCallback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceCallback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["external_id"] = o.ExternalId
	toSerialize["user_id"] = o.UserId
	toSerialize["status"] = o.Status
	toSerialize["merchant_name"] = o.MerchantName
	toSerialize["amount"] = o.Amount
	if !utils.IsNil(o.PayerEmail) {
		toSerialize["payer_email"] = o.PayerEmail
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.PaidAmount) {
		toSerialize["paid_amount"] = o.PaidAmount
	}
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	toSerialize["currency"] = o.Currency
	if !utils.IsNil(o.PaidAt) {
		toSerialize["paid_at"] = o.PaidAt
	}
	if !utils.IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !utils.IsNil(o.PaymentChannel) {
		toSerialize["payment_channel"] = o.PaymentChannel
	}
	if !utils.IsNil(o.PaymentDestination) {
		toSerialize["payment_destination"] = o.PaymentDestination
	}
	if !utils.IsNil(o.PaymentDetails) {
		toSerialize["payment_details"] = o.PaymentDetails
	}
	if !utils.IsNil(o.PaymentId) {
		toSerialize["payment_id"] = o.PaymentId
	}
	if !utils.IsNil(o.SuccessRedirectUrl) {
		toSerialize["success_redirect_url"] = o.SuccessRedirectUrl
	}
	if !utils.IsNil(o.FailureRedirectUrl) {
		toSerialize["failure_redirect_url"] = o.FailureRedirectUrl
	}
	if !utils.IsNil(o.CreditCardChargeId) {
		toSerialize["credit_card_charge_id"] = o.CreditCardChargeId
	}
	if !utils.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !utils.IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !utils.IsNil(o.ShouldAuthenticateCreditCard) {
		toSerialize["should_authenticate_credit_card"] = o.ShouldAuthenticateCreditCard
	}
	if !utils.IsNil(o.BankCode) {
		toSerialize["bank_code"] = o.BankCode
	}
	if !utils.IsNil(o.EwalletType) {
		toSerialize["ewallet_type"] = o.EwalletType
	}
	if !utils.IsNil(o.OnDemandLink) {
		toSerialize["on_demand_link"] = o.OnDemandLink
	}
	if !utils.IsNil(o.RecurringPaymentId) {
		toSerialize["recurring_payment_id"] = o.RecurringPaymentId
	}
	return toSerialize, nil
}

type NullableInvoiceCallback struct {
	value *InvoiceCallback
	isSet bool
}

func (v NullableInvoiceCallback) Get() *InvoiceCallback {
	return v.value
}

func (v *NullableInvoiceCallback) Set(val *InvoiceCallback) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceCallback(val *InvoiceCallback) *NullableInvoiceCallback {
	return &NullableInvoiceCallback{value: val, isSet: true}
}

func (v NullableInvoiceCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



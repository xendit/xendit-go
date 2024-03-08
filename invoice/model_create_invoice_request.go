/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.7.6
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the CreateInvoiceRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateInvoiceRequest{}

// CreateInvoiceRequest An object representing for an invoice creation request.
type CreateInvoiceRequest struct {
	// The external ID of the invoice.
	ExternalId string `json:"external_id"`
	// The invoice amount.
	Amount float64 `json:"amount"`
	// The email address of the payer.
	PayerEmail *string `json:"payer_email,omitempty"`
	// A description of the payment.
	Description *string `json:"description,omitempty"`
	// The duration of the invoice.
	InvoiceDuration *string `json:"invoice_duration,omitempty"`
	// The ID of the callback virtual account.
	CallbackVirtualAccountId *string `json:"callback_virtual_account_id,omitempty"`
	// Indicates whether email notifications should be sent.
	ShouldSendEmail *bool `json:"should_send_email,omitempty"`
	Customer *CustomerObject `json:"customer,omitempty"`
	CustomerNotificationPreference *NotificationPreference `json:"customer_notification_preference,omitempty"`
	// The URL to redirect to on successful payment.
	SuccessRedirectUrl *string `json:"success_redirect_url,omitempty"`
	// The URL to redirect to on payment failure.
	FailureRedirectUrl *string `json:"failure_redirect_url,omitempty"`
	// An array of available payment methods.
	PaymentMethods []string `json:"payment_methods,omitempty"`
	// The middle label.
	MidLabel *string `json:"mid_label,omitempty"`
	// Indicates whether credit card authentication is required.
	ShouldAuthenticateCreditCard *bool `json:"should_authenticate_credit_card,omitempty"`
	// The currency of the invoice.
	Currency *string `json:"currency,omitempty"`
	// The reminder time.
	ReminderTime *float32 `json:"reminder_time,omitempty"`
	// The local.
	Local *string `json:"local,omitempty"`
	// The unit of the reminder time.
	ReminderTimeUnit *string `json:"reminder_time_unit,omitempty"`
	// An array of items included in the invoice.
	Items []InvoiceItem `json:"items,omitempty"`
	// An array of fees associated with the invoice.
	Fees []InvoiceFee `json:"fees,omitempty"`
	ChannelProperties *ChannelProperties `json:"channel_properties,omitempty"`
}

// NewCreateInvoiceRequest instantiates a new CreateInvoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInvoiceRequest(externalId string, amount float64) *CreateInvoiceRequest {
	this := CreateInvoiceRequest{}
	this.ExternalId = externalId
	this.Amount = amount
	return &this
}

// NewCreateInvoiceRequestWithDefaults instantiates a new CreateInvoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInvoiceRequestWithDefaults() *CreateInvoiceRequest {
	this := CreateInvoiceRequest{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *CreateInvoiceRequest) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *CreateInvoiceRequest) SetExternalId(v string) {
	o.ExternalId = v
}

// GetAmount returns the Amount field value
func (o *CreateInvoiceRequest) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateInvoiceRequest) SetAmount(v float64) {
	o.Amount = v
}

// GetPayerEmail returns the PayerEmail field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetPayerEmail() string {
	if o == nil || utils.IsNil(o.PayerEmail) {
		var ret string
		return ret
	}
	return *o.PayerEmail
}

// GetPayerEmailOk returns a tuple with the PayerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetPayerEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PayerEmail) {
		return nil, false
	}
	return o.PayerEmail, true
}

// HasPayerEmail returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasPayerEmail() bool {
	if o != nil && !utils.IsNil(o.PayerEmail) {
		return true
	}

	return false
}

// SetPayerEmail gets a reference to the given string and assigns it to the PayerEmail field.
func (o *CreateInvoiceRequest) SetPayerEmail(v string) {
	o.PayerEmail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateInvoiceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetInvoiceDuration returns the InvoiceDuration field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetInvoiceDuration() string {
	if o == nil || utils.IsNil(o.InvoiceDuration) {
		var ret string
		return ret
	}
	return *o.InvoiceDuration
}

// GetInvoiceDurationOk returns a tuple with the InvoiceDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetInvoiceDurationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InvoiceDuration) {
		return nil, false
	}
	return o.InvoiceDuration, true
}

// HasInvoiceDuration returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasInvoiceDuration() bool {
	if o != nil && !utils.IsNil(o.InvoiceDuration) {
		return true
	}

	return false
}

// SetInvoiceDuration gets a reference to the given string and assigns it to the InvoiceDuration field.
func (o *CreateInvoiceRequest) SetInvoiceDuration(v string) {
	o.InvoiceDuration = &v
}

// GetCallbackVirtualAccountId returns the CallbackVirtualAccountId field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetCallbackVirtualAccountId() string {
	if o == nil || utils.IsNil(o.CallbackVirtualAccountId) {
		var ret string
		return ret
	}
	return *o.CallbackVirtualAccountId
}

// GetCallbackVirtualAccountIdOk returns a tuple with the CallbackVirtualAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetCallbackVirtualAccountIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CallbackVirtualAccountId) {
		return nil, false
	}
	return o.CallbackVirtualAccountId, true
}

// HasCallbackVirtualAccountId returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasCallbackVirtualAccountId() bool {
	if o != nil && !utils.IsNil(o.CallbackVirtualAccountId) {
		return true
	}

	return false
}

// SetCallbackVirtualAccountId gets a reference to the given string and assigns it to the CallbackVirtualAccountId field.
func (o *CreateInvoiceRequest) SetCallbackVirtualAccountId(v string) {
	o.CallbackVirtualAccountId = &v
}

// GetShouldSendEmail returns the ShouldSendEmail field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetShouldSendEmail() bool {
	if o == nil || utils.IsNil(o.ShouldSendEmail) {
		var ret bool
		return ret
	}
	return *o.ShouldSendEmail
}

// GetShouldSendEmailOk returns a tuple with the ShouldSendEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetShouldSendEmailOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ShouldSendEmail) {
		return nil, false
	}
	return o.ShouldSendEmail, true
}

// HasShouldSendEmail returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasShouldSendEmail() bool {
	if o != nil && !utils.IsNil(o.ShouldSendEmail) {
		return true
	}

	return false
}

// SetShouldSendEmail gets a reference to the given bool and assigns it to the ShouldSendEmail field.
func (o *CreateInvoiceRequest) SetShouldSendEmail(v bool) {
	o.ShouldSendEmail = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetCustomer() CustomerObject {
	if o == nil || utils.IsNil(o.Customer) {
		var ret CustomerObject
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetCustomerOk() (*CustomerObject, bool) {
	if o == nil || utils.IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasCustomer() bool {
	if o != nil && !utils.IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CustomerObject and assigns it to the Customer field.
func (o *CreateInvoiceRequest) SetCustomer(v CustomerObject) {
	o.Customer = &v
}

// GetCustomerNotificationPreference returns the CustomerNotificationPreference field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetCustomerNotificationPreference() NotificationPreference {
	if o == nil || utils.IsNil(o.CustomerNotificationPreference) {
		var ret NotificationPreference
		return ret
	}
	return *o.CustomerNotificationPreference
}

// GetCustomerNotificationPreferenceOk returns a tuple with the CustomerNotificationPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetCustomerNotificationPreferenceOk() (*NotificationPreference, bool) {
	if o == nil || utils.IsNil(o.CustomerNotificationPreference) {
		return nil, false
	}
	return o.CustomerNotificationPreference, true
}

// HasCustomerNotificationPreference returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasCustomerNotificationPreference() bool {
	if o != nil && !utils.IsNil(o.CustomerNotificationPreference) {
		return true
	}

	return false
}

// SetCustomerNotificationPreference gets a reference to the given NotificationPreference and assigns it to the CustomerNotificationPreference field.
func (o *CreateInvoiceRequest) SetCustomerNotificationPreference(v NotificationPreference) {
	o.CustomerNotificationPreference = &v
}

// GetSuccessRedirectUrl returns the SuccessRedirectUrl field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetSuccessRedirectUrl() string {
	if o == nil || utils.IsNil(o.SuccessRedirectUrl) {
		var ret string
		return ret
	}
	return *o.SuccessRedirectUrl
}

// GetSuccessRedirectUrlOk returns a tuple with the SuccessRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetSuccessRedirectUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SuccessRedirectUrl) {
		return nil, false
	}
	return o.SuccessRedirectUrl, true
}

// HasSuccessRedirectUrl returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasSuccessRedirectUrl() bool {
	if o != nil && !utils.IsNil(o.SuccessRedirectUrl) {
		return true
	}

	return false
}

// SetSuccessRedirectUrl gets a reference to the given string and assigns it to the SuccessRedirectUrl field.
func (o *CreateInvoiceRequest) SetSuccessRedirectUrl(v string) {
	o.SuccessRedirectUrl = &v
}

// GetFailureRedirectUrl returns the FailureRedirectUrl field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetFailureRedirectUrl() string {
	if o == nil || utils.IsNil(o.FailureRedirectUrl) {
		var ret string
		return ret
	}
	return *o.FailureRedirectUrl
}

// GetFailureRedirectUrlOk returns a tuple with the FailureRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetFailureRedirectUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FailureRedirectUrl) {
		return nil, false
	}
	return o.FailureRedirectUrl, true
}

// HasFailureRedirectUrl returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasFailureRedirectUrl() bool {
	if o != nil && !utils.IsNil(o.FailureRedirectUrl) {
		return true
	}

	return false
}

// SetFailureRedirectUrl gets a reference to the given string and assigns it to the FailureRedirectUrl field.
func (o *CreateInvoiceRequest) SetFailureRedirectUrl(v string) {
	o.FailureRedirectUrl = &v
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetPaymentMethods() []string {
	if o == nil || utils.IsNil(o.PaymentMethods) {
		var ret []string
		return ret
	}
	return o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetPaymentMethodsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasPaymentMethods() bool {
	if o != nil && !utils.IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given []string and assigns it to the PaymentMethods field.
func (o *CreateInvoiceRequest) SetPaymentMethods(v []string) {
	o.PaymentMethods = v
}

// GetMidLabel returns the MidLabel field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetMidLabel() string {
	if o == nil || utils.IsNil(o.MidLabel) {
		var ret string
		return ret
	}
	return *o.MidLabel
}

// GetMidLabelOk returns a tuple with the MidLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetMidLabelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MidLabel) {
		return nil, false
	}
	return o.MidLabel, true
}

// HasMidLabel returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasMidLabel() bool {
	if o != nil && !utils.IsNil(o.MidLabel) {
		return true
	}

	return false
}

// SetMidLabel gets a reference to the given string and assigns it to the MidLabel field.
func (o *CreateInvoiceRequest) SetMidLabel(v string) {
	o.MidLabel = &v
}

// GetShouldAuthenticateCreditCard returns the ShouldAuthenticateCreditCard field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetShouldAuthenticateCreditCard() bool {
	if o == nil || utils.IsNil(o.ShouldAuthenticateCreditCard) {
		var ret bool
		return ret
	}
	return *o.ShouldAuthenticateCreditCard
}

// GetShouldAuthenticateCreditCardOk returns a tuple with the ShouldAuthenticateCreditCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetShouldAuthenticateCreditCardOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ShouldAuthenticateCreditCard) {
		return nil, false
	}
	return o.ShouldAuthenticateCreditCard, true
}

// HasShouldAuthenticateCreditCard returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasShouldAuthenticateCreditCard() bool {
	if o != nil && !utils.IsNil(o.ShouldAuthenticateCreditCard) {
		return true
	}

	return false
}

// SetShouldAuthenticateCreditCard gets a reference to the given bool and assigns it to the ShouldAuthenticateCreditCard field.
func (o *CreateInvoiceRequest) SetShouldAuthenticateCreditCard(v bool) {
	o.ShouldAuthenticateCreditCard = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CreateInvoiceRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetReminderTime returns the ReminderTime field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetReminderTime() float32 {
	if o == nil || utils.IsNil(o.ReminderTime) {
		var ret float32
		return ret
	}
	return *o.ReminderTime
}

// GetReminderTimeOk returns a tuple with the ReminderTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetReminderTimeOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.ReminderTime) {
		return nil, false
	}
	return o.ReminderTime, true
}

// HasReminderTime returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasReminderTime() bool {
	if o != nil && !utils.IsNil(o.ReminderTime) {
		return true
	}

	return false
}

// SetReminderTime gets a reference to the given float32 and assigns it to the ReminderTime field.
func (o *CreateInvoiceRequest) SetReminderTime(v float32) {
	o.ReminderTime = &v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetLocal() string {
	if o == nil || utils.IsNil(o.Local) {
		var ret string
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetLocalOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Local) {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasLocal() bool {
	if o != nil && !utils.IsNil(o.Local) {
		return true
	}

	return false
}

// SetLocal gets a reference to the given string and assigns it to the Local field.
func (o *CreateInvoiceRequest) SetLocal(v string) {
	o.Local = &v
}

// GetReminderTimeUnit returns the ReminderTimeUnit field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetReminderTimeUnit() string {
	if o == nil || utils.IsNil(o.ReminderTimeUnit) {
		var ret string
		return ret
	}
	return *o.ReminderTimeUnit
}

// GetReminderTimeUnitOk returns a tuple with the ReminderTimeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetReminderTimeUnitOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReminderTimeUnit) {
		return nil, false
	}
	return o.ReminderTimeUnit, true
}

// HasReminderTimeUnit returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasReminderTimeUnit() bool {
	if o != nil && !utils.IsNil(o.ReminderTimeUnit) {
		return true
	}

	return false
}

// SetReminderTimeUnit gets a reference to the given string and assigns it to the ReminderTimeUnit field.
func (o *CreateInvoiceRequest) SetReminderTimeUnit(v string) {
	o.ReminderTimeUnit = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetItems() []InvoiceItem {
	if o == nil || utils.IsNil(o.Items) {
		var ret []InvoiceItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetItemsOk() ([]InvoiceItem, bool) {
	if o == nil || utils.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasItems() bool {
	if o != nil && !utils.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InvoiceItem and assigns it to the Items field.
func (o *CreateInvoiceRequest) SetItems(v []InvoiceItem) {
	o.Items = v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetFees() []InvoiceFee {
	if o == nil || utils.IsNil(o.Fees) {
		var ret []InvoiceFee
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetFeesOk() ([]InvoiceFee, bool) {
	if o == nil || utils.IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasFees() bool {
	if o != nil && !utils.IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []InvoiceFee and assigns it to the Fees field.
func (o *CreateInvoiceRequest) SetFees(v []InvoiceFee) {
	o.Fees = v
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *CreateInvoiceRequest) GetChannelProperties() ChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret ChannelProperties
		return ret
	}
	return *o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInvoiceRequest) GetChannelPropertiesOk() (*ChannelProperties, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *CreateInvoiceRequest) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given ChannelProperties and assigns it to the ChannelProperties field.
func (o *CreateInvoiceRequest) SetChannelProperties(v ChannelProperties) {
	o.ChannelProperties = &v
}

func (o CreateInvoiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInvoiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["external_id"] = o.ExternalId
	toSerialize["amount"] = o.Amount
	if !utils.IsNil(o.PayerEmail) {
		toSerialize["payer_email"] = o.PayerEmail
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.InvoiceDuration) {
		toSerialize["invoice_duration"] = o.InvoiceDuration
	}
	if !utils.IsNil(o.CallbackVirtualAccountId) {
		toSerialize["callback_virtual_account_id"] = o.CallbackVirtualAccountId
	}
	if !utils.IsNil(o.ShouldSendEmail) {
		toSerialize["should_send_email"] = o.ShouldSendEmail
	}
	if !utils.IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !utils.IsNil(o.CustomerNotificationPreference) {
		toSerialize["customer_notification_preference"] = o.CustomerNotificationPreference
	}
	if !utils.IsNil(o.SuccessRedirectUrl) {
		toSerialize["success_redirect_url"] = o.SuccessRedirectUrl
	}
	if !utils.IsNil(o.FailureRedirectUrl) {
		toSerialize["failure_redirect_url"] = o.FailureRedirectUrl
	}
	if !utils.IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !utils.IsNil(o.MidLabel) {
		toSerialize["mid_label"] = o.MidLabel
	}
	if !utils.IsNil(o.ShouldAuthenticateCreditCard) {
		toSerialize["should_authenticate_credit_card"] = o.ShouldAuthenticateCreditCard
	}
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !utils.IsNil(o.ReminderTime) {
		toSerialize["reminder_time"] = o.ReminderTime
	}
	if !utils.IsNil(o.Local) {
		toSerialize["local"] = o.Local
	}
	if !utils.IsNil(o.ReminderTimeUnit) {
		toSerialize["reminder_time_unit"] = o.ReminderTimeUnit
	}
	if !utils.IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !utils.IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	return toSerialize, nil
}

type NullableCreateInvoiceRequest struct {
	value *CreateInvoiceRequest
	isSet bool
}

func (v NullableCreateInvoiceRequest) Get() *CreateInvoiceRequest {
	return v.value
}

func (v *NullableCreateInvoiceRequest) Set(val *CreateInvoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInvoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInvoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInvoiceRequest(val *CreateInvoiceRequest) *NullableCreateInvoiceRequest {
	return &NullableCreateInvoiceRequest{value: val, isSet: true}
}

func (v NullableCreateInvoiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInvoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



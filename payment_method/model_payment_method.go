/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
	"time"
)

// checks if the PaymentMethod type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentMethod{}

// PaymentMethod struct for PaymentMethod
type PaymentMethod struct {
	Id string `json:"id"`
	BusinessId *string `json:"business_id,omitempty"`
	Type *PaymentMethodType `json:"type,omitempty"`
	Country *PaymentMethodCountry `json:"country,omitempty"`
	CustomerId NullableString `json:"customer_id,omitempty"`
	Customer map[string]interface{} `json:"customer,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Status *PaymentMethodStatus `json:"status,omitempty"`
	Reusability *PaymentMethodReusability `json:"reusability,omitempty"`
	Actions []PaymentMethodAction `json:"actions,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	BillingInformation NullableBillingInformation `json:"billing_information,omitempty"`
	FailureCode NullableString `json:"failure_code,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
	Ewallet NullableEWallet `json:"ewallet,omitempty"`
	DirectDebit NullableDirectDebit `json:"direct_debit,omitempty"`
	OverTheCounter NullableOverTheCounter `json:"over_the_counter,omitempty"`
	Card NullableCard `json:"card,omitempty"`
	QrCode NullableQRCode `json:"qr_code,omitempty"`
	VirtualAccount NullableVirtualAccount `json:"virtual_account,omitempty"`
}

// NewPaymentMethod instantiates a new PaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethod(id string) *PaymentMethod {
	this := PaymentMethod{}
	this.Id = id
	return &this
}

// NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodWithDefaults() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// GetId returns the Id field value
func (o *PaymentMethod) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentMethod) SetId(v string) {
	o.Id = v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *PaymentMethod) GetBusinessId() string {
	if o == nil || utils.IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetBusinessIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *PaymentMethod) HasBusinessId() bool {
	if o != nil && !utils.IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *PaymentMethod) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethod) GetType() PaymentMethodType {
	if o == nil || utils.IsNil(o.Type) {
		var ret PaymentMethodType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetTypeOk() (*PaymentMethodType, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethod) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PaymentMethodType and assigns it to the Type field.
func (o *PaymentMethod) SetType(v PaymentMethodType) {
	o.Type = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PaymentMethod) GetCountry() PaymentMethodCountry {
	if o == nil || utils.IsNil(o.Country) {
		var ret PaymentMethodCountry
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCountryOk() (*PaymentMethodCountry, bool) {
	if o == nil || utils.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PaymentMethod) HasCountry() bool {
	if o != nil && !utils.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given PaymentMethodCountry and assigns it to the Country field.
func (o *PaymentMethod) SetCountry(v PaymentMethodCountry) {
	o.Country = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetCustomerId() string {
	if o == nil || utils.IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *PaymentMethod) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *PaymentMethod) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *PaymentMethod) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *PaymentMethod) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetCustomer returns the Customer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetCustomer() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetCustomerOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Customer) {
		return map[string]interface{}{}, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *PaymentMethod) HasCustomer() bool {
	if o != nil && utils.IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given map[string]interface{} and assigns it to the Customer field.
func (o *PaymentMethod) SetCustomer(v map[string]interface{}) {
	o.Customer = v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *PaymentMethod) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetReferenceIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *PaymentMethod) HasReferenceId() bool {
	if o != nil && !utils.IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *PaymentMethod) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentMethod) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PaymentMethod) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PaymentMethod) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PaymentMethod) UnsetDescription() {
	o.Description.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentMethod) GetStatus() PaymentMethodStatus {
	if o == nil || utils.IsNil(o.Status) {
		var ret PaymentMethodStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentMethod) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PaymentMethodStatus and assigns it to the Status field.
func (o *PaymentMethod) SetStatus(v PaymentMethodStatus) {
	o.Status = &v
}

// GetReusability returns the Reusability field value if set, zero value otherwise.
func (o *PaymentMethod) GetReusability() PaymentMethodReusability {
	if o == nil || utils.IsNil(o.Reusability) {
		var ret PaymentMethodReusability
		return ret
	}
	return *o.Reusability
}

// GetReusabilityOk returns a tuple with the Reusability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetReusabilityOk() (*PaymentMethodReusability, bool) {
	if o == nil || utils.IsNil(o.Reusability) {
		return nil, false
	}
	return o.Reusability, true
}

// HasReusability returns a boolean if a field has been set.
func (o *PaymentMethod) HasReusability() bool {
	if o != nil && !utils.IsNil(o.Reusability) {
		return true
	}

	return false
}

// SetReusability gets a reference to the given PaymentMethodReusability and assigns it to the Reusability field.
func (o *PaymentMethod) SetReusability(v PaymentMethodReusability) {
	o.Reusability = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *PaymentMethod) GetActions() []PaymentMethodAction {
	if o == nil || utils.IsNil(o.Actions) {
		var ret []PaymentMethodAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetActionsOk() ([]PaymentMethodAction, bool) {
	if o == nil || utils.IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *PaymentMethod) HasActions() bool {
	if o != nil && !utils.IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []PaymentMethodAction and assigns it to the Actions field.
func (o *PaymentMethod) SetActions(v []PaymentMethodAction) {
	o.Actions = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentMethod) HasMetadata() bool {
	if o != nil && utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PaymentMethod) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetBillingInformation returns the BillingInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetBillingInformation() BillingInformation {
	if o == nil || utils.IsNil(o.BillingInformation.Get()) {
		var ret BillingInformation
		return ret
	}
	return *o.BillingInformation.Get()
}

// GetBillingInformationOk returns a tuple with the BillingInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetBillingInformationOk() (*BillingInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingInformation.Get(), o.BillingInformation.IsSet()
}

// HasBillingInformation returns a boolean if a field has been set.
func (o *PaymentMethod) HasBillingInformation() bool {
	if o != nil && o.BillingInformation.IsSet() {
		return true
	}

	return false
}

// SetBillingInformation gets a reference to the given NullableBillingInformation and assigns it to the BillingInformation field.
func (o *PaymentMethod) SetBillingInformation(v BillingInformation) {
	o.BillingInformation.Set(&v)
}
// SetBillingInformationNil sets the value for BillingInformation to be an explicit nil
func (o *PaymentMethod) SetBillingInformationNil() {
	o.BillingInformation.Set(nil)
}

// UnsetBillingInformation ensures that no value is present for BillingInformation, not even an explicit nil
func (o *PaymentMethod) UnsetBillingInformation() {
	o.BillingInformation.Unset()
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetFailureCode() string {
	if o == nil || utils.IsNil(o.FailureCode.Get()) {
		var ret string
		return ret
	}
	return *o.FailureCode.Get()
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetFailureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureCode.Get(), o.FailureCode.IsSet()
}

// HasFailureCode returns a boolean if a field has been set.
func (o *PaymentMethod) HasFailureCode() bool {
	if o != nil && o.FailureCode.IsSet() {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given NullableString and assigns it to the FailureCode field.
func (o *PaymentMethod) SetFailureCode(v string) {
	o.FailureCode.Set(&v)
}
// SetFailureCodeNil sets the value for FailureCode to be an explicit nil
func (o *PaymentMethod) SetFailureCodeNil() {
	o.FailureCode.Set(nil)
}

// UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
func (o *PaymentMethod) UnsetFailureCode() {
	o.FailureCode.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PaymentMethod) GetCreated() time.Time {
	if o == nil || utils.IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCreatedOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PaymentMethod) HasCreated() bool {
	if o != nil && !utils.IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *PaymentMethod) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *PaymentMethod) GetUpdated() time.Time {
	if o == nil || utils.IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *PaymentMethod) HasUpdated() bool {
	if o != nil && !utils.IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *PaymentMethod) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetEwallet returns the Ewallet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetEwallet() EWallet {
	if o == nil || utils.IsNil(o.Ewallet.Get()) {
		var ret EWallet
		return ret
	}
	return *o.Ewallet.Get()
}

// GetEwalletOk returns a tuple with the Ewallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetEwalletOk() (*EWallet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ewallet.Get(), o.Ewallet.IsSet()
}

// HasEwallet returns a boolean if a field has been set.
func (o *PaymentMethod) HasEwallet() bool {
	if o != nil && o.Ewallet.IsSet() {
		return true
	}

	return false
}

// SetEwallet gets a reference to the given NullableEWallet and assigns it to the Ewallet field.
func (o *PaymentMethod) SetEwallet(v EWallet) {
	o.Ewallet.Set(&v)
}
// SetEwalletNil sets the value for Ewallet to be an explicit nil
func (o *PaymentMethod) SetEwalletNil() {
	o.Ewallet.Set(nil)
}

// UnsetEwallet ensures that no value is present for Ewallet, not even an explicit nil
func (o *PaymentMethod) UnsetEwallet() {
	o.Ewallet.Unset()
}

// GetDirectDebit returns the DirectDebit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetDirectDebit() DirectDebit {
	if o == nil || utils.IsNil(o.DirectDebit.Get()) {
		var ret DirectDebit
		return ret
	}
	return *o.DirectDebit.Get()
}

// GetDirectDebitOk returns a tuple with the DirectDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetDirectDebitOk() (*DirectDebit, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectDebit.Get(), o.DirectDebit.IsSet()
}

// HasDirectDebit returns a boolean if a field has been set.
func (o *PaymentMethod) HasDirectDebit() bool {
	if o != nil && o.DirectDebit.IsSet() {
		return true
	}

	return false
}

// SetDirectDebit gets a reference to the given NullableDirectDebit and assigns it to the DirectDebit field.
func (o *PaymentMethod) SetDirectDebit(v DirectDebit) {
	o.DirectDebit.Set(&v)
}
// SetDirectDebitNil sets the value for DirectDebit to be an explicit nil
func (o *PaymentMethod) SetDirectDebitNil() {
	o.DirectDebit.Set(nil)
}

// UnsetDirectDebit ensures that no value is present for DirectDebit, not even an explicit nil
func (o *PaymentMethod) UnsetDirectDebit() {
	o.DirectDebit.Unset()
}

// GetOverTheCounter returns the OverTheCounter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetOverTheCounter() OverTheCounter {
	if o == nil || utils.IsNil(o.OverTheCounter.Get()) {
		var ret OverTheCounter
		return ret
	}
	return *o.OverTheCounter.Get()
}

// GetOverTheCounterOk returns a tuple with the OverTheCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetOverTheCounterOk() (*OverTheCounter, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverTheCounter.Get(), o.OverTheCounter.IsSet()
}

// HasOverTheCounter returns a boolean if a field has been set.
func (o *PaymentMethod) HasOverTheCounter() bool {
	if o != nil && o.OverTheCounter.IsSet() {
		return true
	}

	return false
}

// SetOverTheCounter gets a reference to the given NullableOverTheCounter and assigns it to the OverTheCounter field.
func (o *PaymentMethod) SetOverTheCounter(v OverTheCounter) {
	o.OverTheCounter.Set(&v)
}
// SetOverTheCounterNil sets the value for OverTheCounter to be an explicit nil
func (o *PaymentMethod) SetOverTheCounterNil() {
	o.OverTheCounter.Set(nil)
}

// UnsetOverTheCounter ensures that no value is present for OverTheCounter, not even an explicit nil
func (o *PaymentMethod) UnsetOverTheCounter() {
	o.OverTheCounter.Unset()
}

// GetCard returns the Card field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetCard() Card {
	if o == nil || utils.IsNil(o.Card.Get()) {
		var ret Card
		return ret
	}
	return *o.Card.Get()
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetCardOk() (*Card, bool) {
	if o == nil {
		return nil, false
	}
	return o.Card.Get(), o.Card.IsSet()
}

// HasCard returns a boolean if a field has been set.
func (o *PaymentMethod) HasCard() bool {
	if o != nil && o.Card.IsSet() {
		return true
	}

	return false
}

// SetCard gets a reference to the given NullableCard and assigns it to the Card field.
func (o *PaymentMethod) SetCard(v Card) {
	o.Card.Set(&v)
}
// SetCardNil sets the value for Card to be an explicit nil
func (o *PaymentMethod) SetCardNil() {
	o.Card.Set(nil)
}

// UnsetCard ensures that no value is present for Card, not even an explicit nil
func (o *PaymentMethod) UnsetCard() {
	o.Card.Unset()
}

// GetQrCode returns the QrCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetQrCode() QRCode {
	if o == nil || utils.IsNil(o.QrCode.Get()) {
		var ret QRCode
		return ret
	}
	return *o.QrCode.Get()
}

// GetQrCodeOk returns a tuple with the QrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetQrCodeOk() (*QRCode, bool) {
	if o == nil {
		return nil, false
	}
	return o.QrCode.Get(), o.QrCode.IsSet()
}

// HasQrCode returns a boolean if a field has been set.
func (o *PaymentMethod) HasQrCode() bool {
	if o != nil && o.QrCode.IsSet() {
		return true
	}

	return false
}

// SetQrCode gets a reference to the given NullableQRCode and assigns it to the QrCode field.
func (o *PaymentMethod) SetQrCode(v QRCode) {
	o.QrCode.Set(&v)
}
// SetQrCodeNil sets the value for QrCode to be an explicit nil
func (o *PaymentMethod) SetQrCodeNil() {
	o.QrCode.Set(nil)
}

// UnsetQrCode ensures that no value is present for QrCode, not even an explicit nil
func (o *PaymentMethod) UnsetQrCode() {
	o.QrCode.Unset()
}

// GetVirtualAccount returns the VirtualAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetVirtualAccount() VirtualAccount {
	if o == nil || utils.IsNil(o.VirtualAccount.Get()) {
		var ret VirtualAccount
		return ret
	}
	return *o.VirtualAccount.Get()
}

// GetVirtualAccountOk returns a tuple with the VirtualAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetVirtualAccountOk() (*VirtualAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualAccount.Get(), o.VirtualAccount.IsSet()
}

// HasVirtualAccount returns a boolean if a field has been set.
func (o *PaymentMethod) HasVirtualAccount() bool {
	if o != nil && o.VirtualAccount.IsSet() {
		return true
	}

	return false
}

// SetVirtualAccount gets a reference to the given NullableVirtualAccount and assigns it to the VirtualAccount field.
func (o *PaymentMethod) SetVirtualAccount(v VirtualAccount) {
	o.VirtualAccount.Set(&v)
}
// SetVirtualAccountNil sets the value for VirtualAccount to be an explicit nil
func (o *PaymentMethod) SetVirtualAccountNil() {
	o.VirtualAccount.Set(nil)
}

// UnsetVirtualAccount ensures that no value is present for VirtualAccount, not even an explicit nil
func (o *PaymentMethod) UnsetVirtualAccount() {
	o.VirtualAccount.Unset()
}

func (o PaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !utils.IsNil(o.BusinessId) {
		toSerialize["business_id"] = o.BusinessId
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
    }
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
    }
	if !utils.IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
    }
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.Reusability) {
		toSerialize["reusability"] = o.Reusability
	}
	if !utils.IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
    }
	if o.BillingInformation.IsSet() {
		toSerialize["billing_information"] = o.BillingInformation.Get()
    }
	if o.FailureCode.IsSet() {
		toSerialize["failure_code"] = o.FailureCode.Get()
    }
	if !utils.IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !utils.IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if o.Ewallet.IsSet() {
		toSerialize["ewallet"] = o.Ewallet.Get()
    }
	if o.DirectDebit.IsSet() {
		toSerialize["direct_debit"] = o.DirectDebit.Get()
    }
	if o.OverTheCounter.IsSet() {
		toSerialize["over_the_counter"] = o.OverTheCounter.Get()
    }
	if o.Card.IsSet() {
		toSerialize["card"] = o.Card.Get()
    }
	if o.QrCode.IsSet() {
		toSerialize["qr_code"] = o.QrCode.Get()
    }
	if o.VirtualAccount.IsSet() {
		toSerialize["virtual_account"] = o.VirtualAccount.Get()
    }
	return toSerialize, nil
}

type NullablePaymentMethod struct {
	value *PaymentMethod
	isSet bool
}

func (v NullablePaymentMethod) Get() *PaymentMethod {
	return v.value
}

func (v *NullablePaymentMethod) Set(val *PaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethod(val *PaymentMethod) *NullablePaymentMethod {
	return &NullablePaymentMethod{value: val, isSet: true}
}

func (v NullablePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



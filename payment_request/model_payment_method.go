/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.1
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the PaymentMethod type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentMethod{}

// PaymentMethod struct for PaymentMethod
type PaymentMethod struct {
	Id string `json:"id"`
	Type PaymentMethodType `json:"type"`
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Description NullableString `json:"description,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
	Card NullableCard `json:"card,omitempty"`
	DirectDebit NullableDirectDebit `json:"direct_debit,omitempty"`
	Ewallet NullableEWallet `json:"ewallet,omitempty"`
	OverTheCounter NullableOverTheCounter `json:"over_the_counter,omitempty"`
	VirtualAccount NullableVirtualAccount `json:"virtual_account,omitempty"`
	QrCode NullableQRCode `json:"qr_code,omitempty"`
	Reusability PaymentMethodReusability `json:"reusability"`
	Status PaymentMethodStatus `json:"status"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPaymentMethod instantiates a new PaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethod(id string, type_ PaymentMethodType, reusability PaymentMethodReusability, status PaymentMethodStatus) *PaymentMethod {
	this := PaymentMethod{}
	this.Id = id
	this.Type = type_
	this.Reusability = reusability
	this.Status = status
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

// GetType returns the Type field value
func (o *PaymentMethod) GetType() PaymentMethodType {
	if o == nil {
		var ret PaymentMethodType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetTypeOk() (*PaymentMethodType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethod) SetType(v PaymentMethodType) {
	o.Type = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PaymentMethod) GetCreated() string {
	if o == nil || utils.IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCreatedOk() (*string, bool) {
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

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *PaymentMethod) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *PaymentMethod) GetUpdated() string {
	if o == nil || utils.IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetUpdatedOk() (*string, bool) {
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

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *PaymentMethod) SetUpdated(v string) {
	o.Updated = &v
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

// GetReusability returns the Reusability field value
func (o *PaymentMethod) GetReusability() PaymentMethodReusability {
	if o == nil {
		var ret PaymentMethodReusability
		return ret
	}

	return o.Reusability
}

// GetReusabilityOk returns a tuple with the Reusability field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetReusabilityOk() (*PaymentMethodReusability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reusability, true
}

// SetReusability sets field value
func (o *PaymentMethod) SetReusability(v PaymentMethodReusability) {
	o.Reusability = v
}

// GetStatus returns the Status field value
func (o *PaymentMethod) GetStatus() PaymentMethodStatus {
	if o == nil {
		var ret PaymentMethodStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentMethod) SetStatus(v PaymentMethodStatus) {
	o.Status = v
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
	toSerialize["type"] = o.Type
	if !utils.IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !utils.IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
    }
	if !utils.IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if o.Card.IsSet() {
		toSerialize["card"] = o.Card.Get()
    }
	if o.DirectDebit.IsSet() {
		toSerialize["direct_debit"] = o.DirectDebit.Get()
    }
	if o.Ewallet.IsSet() {
		toSerialize["ewallet"] = o.Ewallet.Get()
    }
	if o.OverTheCounter.IsSet() {
		toSerialize["over_the_counter"] = o.OverTheCounter.Get()
    }
	if o.VirtualAccount.IsSet() {
		toSerialize["virtual_account"] = o.VirtualAccount.Get()
    }
	if o.QrCode.IsSet() {
		toSerialize["qr_code"] = o.QrCode.Get()
    }
	toSerialize["reusability"] = o.Reusability
	toSerialize["status"] = o.Status
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
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



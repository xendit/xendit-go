/*
Payment Requests

This API is used for Payment Requests

API version: 1.44.1
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the PaymentMethodParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentMethodParameters{}

// PaymentMethodParameters struct for PaymentMethodParameters
type PaymentMethodParameters struct {
	Type PaymentMethodType `json:"type"`
	Reusability PaymentMethodReusability `json:"reusability"`
	Description NullableString `json:"description,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
	DirectDebit NullableDirectDebitParameters `json:"direct_debit,omitempty"`
	Ewallet NullableEWalletParameters `json:"ewallet,omitempty"`
	OverTheCounter NullableOverTheCounterParameters `json:"over_the_counter,omitempty"`
	VirtualAccount NullableVirtualAccountParameters `json:"virtual_account,omitempty"`
	QrCode NullableQRCodeParameters `json:"qr_code,omitempty"`
}

// NewPaymentMethodParameters instantiates a new PaymentMethodParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodParameters(type_ PaymentMethodType, reusability PaymentMethodReusability) *PaymentMethodParameters {
	this := PaymentMethodParameters{}
	this.Type = type_
	this.Reusability = reusability
	return &this
}

// NewPaymentMethodParametersWithDefaults instantiates a new PaymentMethodParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodParametersWithDefaults() *PaymentMethodParameters {
	this := PaymentMethodParameters{}
	return &this
}

// GetType returns the Type field value
func (o *PaymentMethodParameters) GetType() PaymentMethodType {
	if o == nil {
		var ret PaymentMethodType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodParameters) GetTypeOk() (*PaymentMethodType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethodParameters) SetType(v PaymentMethodType) {
	o.Type = v
}

// GetReusability returns the Reusability field value
func (o *PaymentMethodParameters) GetReusability() PaymentMethodReusability {
	if o == nil {
		var ret PaymentMethodReusability
		return ret
	}

	return o.Reusability
}

// GetReusabilityOk returns a tuple with the Reusability field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodParameters) GetReusabilityOk() (*PaymentMethodReusability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reusability, true
}

// SetReusability sets field value
func (o *PaymentMethodParameters) SetReusability(v PaymentMethodReusability) {
	o.Reusability = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodParameters) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodParameters) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PaymentMethodParameters) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PaymentMethodParameters) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PaymentMethodParameters) UnsetDescription() {
	o.Description.Unset()
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *PaymentMethodParameters) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodParameters) GetReferenceIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasReferenceId() bool {
	if o != nil && !utils.IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *PaymentMethodParameters) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetDirectDebit returns the DirectDebit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodParameters) GetDirectDebit() DirectDebitParameters {
	if o == nil || utils.IsNil(o.DirectDebit.Get()) {
		var ret DirectDebitParameters
		return ret
	}
	return *o.DirectDebit.Get()
}

// GetDirectDebitOk returns a tuple with the DirectDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodParameters) GetDirectDebitOk() (*DirectDebitParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectDebit.Get(), o.DirectDebit.IsSet()
}

// HasDirectDebit returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasDirectDebit() bool {
	if o != nil && o.DirectDebit.IsSet() {
		return true
	}

	return false
}

// SetDirectDebit gets a reference to the given NullableDirectDebitParameters and assigns it to the DirectDebit field.
func (o *PaymentMethodParameters) SetDirectDebit(v DirectDebitParameters) {
	o.DirectDebit.Set(&v)
}
// SetDirectDebitNil sets the value for DirectDebit to be an explicit nil
func (o *PaymentMethodParameters) SetDirectDebitNil() {
	o.DirectDebit.Set(nil)
}

// UnsetDirectDebit ensures that no value is present for DirectDebit, not even an explicit nil
func (o *PaymentMethodParameters) UnsetDirectDebit() {
	o.DirectDebit.Unset()
}

// GetEwallet returns the Ewallet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodParameters) GetEwallet() EWalletParameters {
	if o == nil || utils.IsNil(o.Ewallet.Get()) {
		var ret EWalletParameters
		return ret
	}
	return *o.Ewallet.Get()
}

// GetEwalletOk returns a tuple with the Ewallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodParameters) GetEwalletOk() (*EWalletParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ewallet.Get(), o.Ewallet.IsSet()
}

// HasEwallet returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasEwallet() bool {
	if o != nil && o.Ewallet.IsSet() {
		return true
	}

	return false
}

// SetEwallet gets a reference to the given NullableEWalletParameters and assigns it to the Ewallet field.
func (o *PaymentMethodParameters) SetEwallet(v EWalletParameters) {
	o.Ewallet.Set(&v)
}
// SetEwalletNil sets the value for Ewallet to be an explicit nil
func (o *PaymentMethodParameters) SetEwalletNil() {
	o.Ewallet.Set(nil)
}

// UnsetEwallet ensures that no value is present for Ewallet, not even an explicit nil
func (o *PaymentMethodParameters) UnsetEwallet() {
	o.Ewallet.Unset()
}

// GetOverTheCounter returns the OverTheCounter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodParameters) GetOverTheCounter() OverTheCounterParameters {
	if o == nil || utils.IsNil(o.OverTheCounter.Get()) {
		var ret OverTheCounterParameters
		return ret
	}
	return *o.OverTheCounter.Get()
}

// GetOverTheCounterOk returns a tuple with the OverTheCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodParameters) GetOverTheCounterOk() (*OverTheCounterParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.OverTheCounter.Get(), o.OverTheCounter.IsSet()
}

// HasOverTheCounter returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasOverTheCounter() bool {
	if o != nil && o.OverTheCounter.IsSet() {
		return true
	}

	return false
}

// SetOverTheCounter gets a reference to the given NullableOverTheCounterParameters and assigns it to the OverTheCounter field.
func (o *PaymentMethodParameters) SetOverTheCounter(v OverTheCounterParameters) {
	o.OverTheCounter.Set(&v)
}
// SetOverTheCounterNil sets the value for OverTheCounter to be an explicit nil
func (o *PaymentMethodParameters) SetOverTheCounterNil() {
	o.OverTheCounter.Set(nil)
}

// UnsetOverTheCounter ensures that no value is present for OverTheCounter, not even an explicit nil
func (o *PaymentMethodParameters) UnsetOverTheCounter() {
	o.OverTheCounter.Unset()
}

// GetVirtualAccount returns the VirtualAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodParameters) GetVirtualAccount() VirtualAccountParameters {
	if o == nil || utils.IsNil(o.VirtualAccount.Get()) {
		var ret VirtualAccountParameters
		return ret
	}
	return *o.VirtualAccount.Get()
}

// GetVirtualAccountOk returns a tuple with the VirtualAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodParameters) GetVirtualAccountOk() (*VirtualAccountParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualAccount.Get(), o.VirtualAccount.IsSet()
}

// HasVirtualAccount returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasVirtualAccount() bool {
	if o != nil && o.VirtualAccount.IsSet() {
		return true
	}

	return false
}

// SetVirtualAccount gets a reference to the given NullableVirtualAccountParameters and assigns it to the VirtualAccount field.
func (o *PaymentMethodParameters) SetVirtualAccount(v VirtualAccountParameters) {
	o.VirtualAccount.Set(&v)
}
// SetVirtualAccountNil sets the value for VirtualAccount to be an explicit nil
func (o *PaymentMethodParameters) SetVirtualAccountNil() {
	o.VirtualAccount.Set(nil)
}

// UnsetVirtualAccount ensures that no value is present for VirtualAccount, not even an explicit nil
func (o *PaymentMethodParameters) UnsetVirtualAccount() {
	o.VirtualAccount.Unset()
}

// GetQrCode returns the QrCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodParameters) GetQrCode() QRCodeParameters {
	if o == nil || utils.IsNil(o.QrCode.Get()) {
		var ret QRCodeParameters
		return ret
	}
	return *o.QrCode.Get()
}

// GetQrCodeOk returns a tuple with the QrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodParameters) GetQrCodeOk() (*QRCodeParameters, bool) {
	if o == nil {
		return nil, false
	}
	return o.QrCode.Get(), o.QrCode.IsSet()
}

// HasQrCode returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasQrCode() bool {
	if o != nil && o.QrCode.IsSet() {
		return true
	}

	return false
}

// SetQrCode gets a reference to the given NullableQRCodeParameters and assigns it to the QrCode field.
func (o *PaymentMethodParameters) SetQrCode(v QRCodeParameters) {
	o.QrCode.Set(&v)
}
// SetQrCodeNil sets the value for QrCode to be an explicit nil
func (o *PaymentMethodParameters) SetQrCodeNil() {
	o.QrCode.Set(nil)
}

// UnsetQrCode ensures that no value is present for QrCode, not even an explicit nil
func (o *PaymentMethodParameters) UnsetQrCode() {
	o.QrCode.Unset()
}

func (o PaymentMethodParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["reusability"] = o.Reusability
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
    }
	if !utils.IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
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
	return toSerialize, nil
}

type NullablePaymentMethodParameters struct {
	value *PaymentMethodParameters
	isSet bool
}

func (v NullablePaymentMethodParameters) Get() *PaymentMethodParameters {
	return v.value
}

func (v *NullablePaymentMethodParameters) Set(val *PaymentMethodParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodParameters(val *PaymentMethodParameters) *NullablePaymentMethodParameters {
	return &NullablePaymentMethodParameters{value: val, isSet: true}
}

func (v NullablePaymentMethodParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



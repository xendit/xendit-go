/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.99.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the PaymentMethodParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentMethodParameters{}

// PaymentMethodParameters struct for PaymentMethodParameters
type PaymentMethodParameters struct {
	Type PaymentMethodType `json:"type"`
	Country NullableString `json:"country,omitempty"`
	Reusability PaymentMethodReusability `json:"reusability"`
	CustomerId NullableString `json:"customer_id,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Card *CardParameters `json:"card,omitempty"`
	DirectDebit *DirectDebitParameters `json:"direct_debit,omitempty"`
	Ewallet *EWalletParameters `json:"ewallet,omitempty"`
	OverTheCounter *OverTheCounterParameters `json:"over_the_counter,omitempty"`
	VirtualAccount *VirtualAccountParameters `json:"virtual_account,omitempty"`
	QrCode *QRCodeParameters `json:"qr_code,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	BillingInformation NullableBillingInformation `json:"billing_information,omitempty"`
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

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodParameters) GetCountry() string {
	if o == nil || utils.IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodParameters) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *PaymentMethodParameters) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *PaymentMethodParameters) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *PaymentMethodParameters) UnsetCountry() {
	o.Country.Unset()
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

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodParameters) GetCustomerId() string {
	if o == nil || utils.IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodParameters) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *PaymentMethodParameters) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *PaymentMethodParameters) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *PaymentMethodParameters) UnsetCustomerId() {
	o.CustomerId.Unset()
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

// GetCard returns the Card field value if set, zero value otherwise.
func (o *PaymentMethodParameters) GetCard() CardParameters {
	if o == nil || utils.IsNil(o.Card) {
		var ret CardParameters
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodParameters) GetCardOk() (*CardParameters, bool) {
	if o == nil || utils.IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasCard() bool {
	if o != nil && !utils.IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given CardParameters and assigns it to the Card field.
func (o *PaymentMethodParameters) SetCard(v CardParameters) {
	o.Card = &v
}

// GetDirectDebit returns the DirectDebit field value if set, zero value otherwise.
func (o *PaymentMethodParameters) GetDirectDebit() DirectDebitParameters {
	if o == nil || utils.IsNil(o.DirectDebit) {
		var ret DirectDebitParameters
		return ret
	}
	return *o.DirectDebit
}

// GetDirectDebitOk returns a tuple with the DirectDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodParameters) GetDirectDebitOk() (*DirectDebitParameters, bool) {
	if o == nil || utils.IsNil(o.DirectDebit) {
		return nil, false
	}
	return o.DirectDebit, true
}

// HasDirectDebit returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasDirectDebit() bool {
	if o != nil && !utils.IsNil(o.DirectDebit) {
		return true
	}

	return false
}

// SetDirectDebit gets a reference to the given DirectDebitParameters and assigns it to the DirectDebit field.
func (o *PaymentMethodParameters) SetDirectDebit(v DirectDebitParameters) {
	o.DirectDebit = &v
}

// GetEwallet returns the Ewallet field value if set, zero value otherwise.
func (o *PaymentMethodParameters) GetEwallet() EWalletParameters {
	if o == nil || utils.IsNil(o.Ewallet) {
		var ret EWalletParameters
		return ret
	}
	return *o.Ewallet
}

// GetEwalletOk returns a tuple with the Ewallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodParameters) GetEwalletOk() (*EWalletParameters, bool) {
	if o == nil || utils.IsNil(o.Ewallet) {
		return nil, false
	}
	return o.Ewallet, true
}

// HasEwallet returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasEwallet() bool {
	if o != nil && !utils.IsNil(o.Ewallet) {
		return true
	}

	return false
}

// SetEwallet gets a reference to the given EWalletParameters and assigns it to the Ewallet field.
func (o *PaymentMethodParameters) SetEwallet(v EWalletParameters) {
	o.Ewallet = &v
}

// GetOverTheCounter returns the OverTheCounter field value if set, zero value otherwise.
func (o *PaymentMethodParameters) GetOverTheCounter() OverTheCounterParameters {
	if o == nil || utils.IsNil(o.OverTheCounter) {
		var ret OverTheCounterParameters
		return ret
	}
	return *o.OverTheCounter
}

// GetOverTheCounterOk returns a tuple with the OverTheCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodParameters) GetOverTheCounterOk() (*OverTheCounterParameters, bool) {
	if o == nil || utils.IsNil(o.OverTheCounter) {
		return nil, false
	}
	return o.OverTheCounter, true
}

// HasOverTheCounter returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasOverTheCounter() bool {
	if o != nil && !utils.IsNil(o.OverTheCounter) {
		return true
	}

	return false
}

// SetOverTheCounter gets a reference to the given OverTheCounterParameters and assigns it to the OverTheCounter field.
func (o *PaymentMethodParameters) SetOverTheCounter(v OverTheCounterParameters) {
	o.OverTheCounter = &v
}

// GetVirtualAccount returns the VirtualAccount field value if set, zero value otherwise.
func (o *PaymentMethodParameters) GetVirtualAccount() VirtualAccountParameters {
	if o == nil || utils.IsNil(o.VirtualAccount) {
		var ret VirtualAccountParameters
		return ret
	}
	return *o.VirtualAccount
}

// GetVirtualAccountOk returns a tuple with the VirtualAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodParameters) GetVirtualAccountOk() (*VirtualAccountParameters, bool) {
	if o == nil || utils.IsNil(o.VirtualAccount) {
		return nil, false
	}
	return o.VirtualAccount, true
}

// HasVirtualAccount returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasVirtualAccount() bool {
	if o != nil && !utils.IsNil(o.VirtualAccount) {
		return true
	}

	return false
}

// SetVirtualAccount gets a reference to the given VirtualAccountParameters and assigns it to the VirtualAccount field.
func (o *PaymentMethodParameters) SetVirtualAccount(v VirtualAccountParameters) {
	o.VirtualAccount = &v
}

// GetQrCode returns the QrCode field value if set, zero value otherwise.
func (o *PaymentMethodParameters) GetQrCode() QRCodeParameters {
	if o == nil || utils.IsNil(o.QrCode) {
		var ret QRCodeParameters
		return ret
	}
	return *o.QrCode
}

// GetQrCodeOk returns a tuple with the QrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodParameters) GetQrCodeOk() (*QRCodeParameters, bool) {
	if o == nil || utils.IsNil(o.QrCode) {
		return nil, false
	}
	return o.QrCode, true
}

// HasQrCode returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasQrCode() bool {
	if o != nil && !utils.IsNil(o.QrCode) {
		return true
	}

	return false
}

// SetQrCode gets a reference to the given QRCodeParameters and assigns it to the QrCode field.
func (o *PaymentMethodParameters) SetQrCode(v QRCodeParameters) {
	o.QrCode = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodParameters) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodParameters) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasMetadata() bool {
	if o != nil && utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PaymentMethodParameters) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetBillingInformation returns the BillingInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodParameters) GetBillingInformation() BillingInformation {
	if o == nil || utils.IsNil(o.BillingInformation.Get()) {
		var ret BillingInformation
		return ret
	}
	return *o.BillingInformation.Get()
}

// GetBillingInformationOk returns a tuple with the BillingInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodParameters) GetBillingInformationOk() (*BillingInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingInformation.Get(), o.BillingInformation.IsSet()
}

// HasBillingInformation returns a boolean if a field has been set.
func (o *PaymentMethodParameters) HasBillingInformation() bool {
	if o != nil && o.BillingInformation.IsSet() {
		return true
	}

	return false
}

// SetBillingInformation gets a reference to the given NullableBillingInformation and assigns it to the BillingInformation field.
func (o *PaymentMethodParameters) SetBillingInformation(v BillingInformation) {
	o.BillingInformation.Set(&v)
}
// SetBillingInformationNil sets the value for BillingInformation to be an explicit nil
func (o *PaymentMethodParameters) SetBillingInformationNil() {
	o.BillingInformation.Set(nil)
}

// UnsetBillingInformation ensures that no value is present for BillingInformation, not even an explicit nil
func (o *PaymentMethodParameters) UnsetBillingInformation() {
	o.BillingInformation.Unset()
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
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
    }
	toSerialize["reusability"] = o.Reusability
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
    }
	if !utils.IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
    }
	if !utils.IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !utils.IsNil(o.DirectDebit) {
		toSerialize["direct_debit"] = o.DirectDebit
	}
	if !utils.IsNil(o.Ewallet) {
		toSerialize["ewallet"] = o.Ewallet
	}
	if !utils.IsNil(o.OverTheCounter) {
		toSerialize["over_the_counter"] = o.OverTheCounter
	}
	if !utils.IsNil(o.VirtualAccount) {
		toSerialize["virtual_account"] = o.VirtualAccount
	}
	if !utils.IsNil(o.QrCode) {
		toSerialize["qr_code"] = o.QrCode
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
    }
	if o.BillingInformation.IsSet() {
		toSerialize["billing_information"] = o.BillingInformation.Get()
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



/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.2
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the PaymentRequestParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestParameters{}

// PaymentRequestParameters struct for PaymentRequestParameters
type PaymentRequestParameters struct {
	ReferenceId *string `json:"reference_id,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	Currency PaymentRequestCurrency `json:"currency"`
	PaymentMethod *PaymentMethodParameters `json:"payment_method,omitempty"`
	Description NullableString `json:"description,omitempty"`
	CaptureMethod NullablePaymentRequestCaptureMethod `json:"capture_method,omitempty"`
	Initiator NullablePaymentRequestInitiator `json:"initiator,omitempty"`
	PaymentMethodId *string `json:"payment_method_id,omitempty"`
	ChannelProperties *PaymentRequestParametersChannelProperties `json:"channel_properties,omitempty"`
	ShippingInformation NullablePaymentRequestShippingInformation `json:"shipping_information,omitempty"`
	Items []PaymentRequestBasketItem `json:"items,omitempty"`
	CustomerId NullableString `json:"customer_id,omitempty"`
	Customer map[string]interface{} `json:"customer,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPaymentRequestParameters instantiates a new PaymentRequestParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestParameters(currency PaymentRequestCurrency) *PaymentRequestParameters {
	this := PaymentRequestParameters{}
	this.Currency = currency
	return &this
}

// NewPaymentRequestParametersWithDefaults instantiates a new PaymentRequestParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestParametersWithDefaults() *PaymentRequestParameters {
	this := PaymentRequestParameters{}
	return &this
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *PaymentRequestParameters) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParameters) GetReferenceIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasReferenceId() bool {
	if o != nil && !utils.IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *PaymentRequestParameters) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentRequestParameters) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParameters) GetAmountOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *PaymentRequestParameters) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value
func (o *PaymentRequestParameters) GetCurrency() PaymentRequestCurrency {
	if o == nil {
		var ret PaymentRequestCurrency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestParameters) GetCurrencyOk() (*PaymentRequestCurrency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentRequestParameters) SetCurrency(v PaymentRequestCurrency) {
	o.Currency = v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *PaymentRequestParameters) GetPaymentMethod() PaymentMethodParameters {
	if o == nil || utils.IsNil(o.PaymentMethod) {
		var ret PaymentMethodParameters
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParameters) GetPaymentMethodOk() (*PaymentMethodParameters, bool) {
	if o == nil || utils.IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasPaymentMethod() bool {
	if o != nil && !utils.IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given PaymentMethodParameters and assigns it to the PaymentMethod field.
func (o *PaymentRequestParameters) SetPaymentMethod(v PaymentMethodParameters) {
	o.PaymentMethod = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestParameters) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestParameters) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PaymentRequestParameters) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PaymentRequestParameters) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PaymentRequestParameters) UnsetDescription() {
	o.Description.Unset()
}

// GetCaptureMethod returns the CaptureMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestParameters) GetCaptureMethod() PaymentRequestCaptureMethod {
	if o == nil || utils.IsNil(o.CaptureMethod.Get()) {
		var ret PaymentRequestCaptureMethod
		return ret
	}
	return *o.CaptureMethod.Get()
}

// GetCaptureMethodOk returns a tuple with the CaptureMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestParameters) GetCaptureMethodOk() (*PaymentRequestCaptureMethod, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaptureMethod.Get(), o.CaptureMethod.IsSet()
}

// HasCaptureMethod returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasCaptureMethod() bool {
	if o != nil && o.CaptureMethod.IsSet() {
		return true
	}

	return false
}

// SetCaptureMethod gets a reference to the given NullablePaymentRequestCaptureMethod and assigns it to the CaptureMethod field.
func (o *PaymentRequestParameters) SetCaptureMethod(v PaymentRequestCaptureMethod) {
	o.CaptureMethod.Set(&v)
}
// SetCaptureMethodNil sets the value for CaptureMethod to be an explicit nil
func (o *PaymentRequestParameters) SetCaptureMethodNil() {
	o.CaptureMethod.Set(nil)
}

// UnsetCaptureMethod ensures that no value is present for CaptureMethod, not even an explicit nil
func (o *PaymentRequestParameters) UnsetCaptureMethod() {
	o.CaptureMethod.Unset()
}

// GetInitiator returns the Initiator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestParameters) GetInitiator() PaymentRequestInitiator {
	if o == nil || utils.IsNil(o.Initiator.Get()) {
		var ret PaymentRequestInitiator
		return ret
	}
	return *o.Initiator.Get()
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestParameters) GetInitiatorOk() (*PaymentRequestInitiator, bool) {
	if o == nil {
		return nil, false
	}
	return o.Initiator.Get(), o.Initiator.IsSet()
}

// HasInitiator returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasInitiator() bool {
	if o != nil && o.Initiator.IsSet() {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given NullablePaymentRequestInitiator and assigns it to the Initiator field.
func (o *PaymentRequestParameters) SetInitiator(v PaymentRequestInitiator) {
	o.Initiator.Set(&v)
}
// SetInitiatorNil sets the value for Initiator to be an explicit nil
func (o *PaymentRequestParameters) SetInitiatorNil() {
	o.Initiator.Set(nil)
}

// UnsetInitiator ensures that no value is present for Initiator, not even an explicit nil
func (o *PaymentRequestParameters) UnsetInitiator() {
	o.Initiator.Unset()
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *PaymentRequestParameters) GetPaymentMethodId() string {
	if o == nil || utils.IsNil(o.PaymentMethodId) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParameters) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasPaymentMethodId() bool {
	if o != nil && !utils.IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *PaymentRequestParameters) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *PaymentRequestParameters) GetChannelProperties() PaymentRequestParametersChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret PaymentRequestParametersChannelProperties
		return ret
	}
	return *o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParameters) GetChannelPropertiesOk() (*PaymentRequestParametersChannelProperties, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given PaymentRequestParametersChannelProperties and assigns it to the ChannelProperties field.
func (o *PaymentRequestParameters) SetChannelProperties(v PaymentRequestParametersChannelProperties) {
	o.ChannelProperties = &v
}

// GetShippingInformation returns the ShippingInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestParameters) GetShippingInformation() PaymentRequestShippingInformation {
	if o == nil || utils.IsNil(o.ShippingInformation.Get()) {
		var ret PaymentRequestShippingInformation
		return ret
	}
	return *o.ShippingInformation.Get()
}

// GetShippingInformationOk returns a tuple with the ShippingInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestParameters) GetShippingInformationOk() (*PaymentRequestShippingInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShippingInformation.Get(), o.ShippingInformation.IsSet()
}

// HasShippingInformation returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasShippingInformation() bool {
	if o != nil && o.ShippingInformation.IsSet() {
		return true
	}

	return false
}

// SetShippingInformation gets a reference to the given NullablePaymentRequestShippingInformation and assigns it to the ShippingInformation field.
func (o *PaymentRequestParameters) SetShippingInformation(v PaymentRequestShippingInformation) {
	o.ShippingInformation.Set(&v)
}
// SetShippingInformationNil sets the value for ShippingInformation to be an explicit nil
func (o *PaymentRequestParameters) SetShippingInformationNil() {
	o.ShippingInformation.Set(nil)
}

// UnsetShippingInformation ensures that no value is present for ShippingInformation, not even an explicit nil
func (o *PaymentRequestParameters) UnsetShippingInformation() {
	o.ShippingInformation.Unset()
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestParameters) GetItems() []PaymentRequestBasketItem {
	if o == nil {
		var ret []PaymentRequestBasketItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestParameters) GetItemsOk() ([]PaymentRequestBasketItem, bool) {
	if o == nil || utils.IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasItems() bool {
	if o != nil && utils.IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PaymentRequestBasketItem and assigns it to the Items field.
func (o *PaymentRequestParameters) SetItems(v []PaymentRequestBasketItem) {
	o.Items = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestParameters) GetCustomerId() string {
	if o == nil || utils.IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestParameters) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *PaymentRequestParameters) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *PaymentRequestParameters) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *PaymentRequestParameters) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetCustomer returns the Customer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestParameters) GetCustomer() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestParameters) GetCustomerOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Customer) {
		return map[string]interface{}{}, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasCustomer() bool {
	if o != nil && utils.IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given map[string]interface{} and assigns it to the Customer field.
func (o *PaymentRequestParameters) SetCustomer(v map[string]interface{}) {
	o.Customer = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestParameters) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestParameters) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentRequestParameters) HasMetadata() bool {
	if o != nil && utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PaymentRequestParameters) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PaymentRequestParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["currency"] = o.Currency
	if !utils.IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
    }
	if o.CaptureMethod.IsSet() {
		toSerialize["capture_method"] = o.CaptureMethod.Get()
    }
	if o.Initiator.IsSet() {
		toSerialize["initiator"] = o.Initiator.Get()
    }
	if !utils.IsNil(o.PaymentMethodId) {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	if o.ShippingInformation.IsSet() {
		toSerialize["shipping_information"] = o.ShippingInformation.Get()
    }
	if o.Items != nil {
		toSerialize["items"] = o.Items
    }
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
    }
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
    }
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
    }
	return toSerialize, nil
}

type NullablePaymentRequestParameters struct {
	value *PaymentRequestParameters
	isSet bool
}

func (v NullablePaymentRequestParameters) Get() *PaymentRequestParameters {
	return v.value
}

func (v *NullablePaymentRequestParameters) Set(val *PaymentRequestParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestParameters(val *PaymentRequestParameters) *NullablePaymentRequestParameters {
	return &NullablePaymentRequestParameters{value: val, isSet: true}
}

func (v NullablePaymentRequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



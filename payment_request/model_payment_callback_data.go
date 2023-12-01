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

// checks if the PaymentCallbackData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentCallbackData{}

// PaymentCallbackData Represents the actual funds transaction/attempt made to a payment method
type PaymentCallbackData struct {
	Id string `json:"id"`
	PaymentRequestId NullableString `json:"payment_request_id,omitempty"`
	ReferenceId string `json:"reference_id"`
	CustomerId NullableString `json:"customer_id,omitempty"`
	Currency string `json:"currency"`
	Amount float64 `json:"amount"`
	Country string `json:"country"`
	Status string `json:"status"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	ChannelProperties NullablePaymentRequestChannelProperties `json:"channel_properties,omitempty"`
	PaymentDetail map[string]interface{} `json:"payment_detail,omitempty"`
	FailureCode map[string]interface{} `json:"failure_code,omitempty"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPaymentCallbackData instantiates a new PaymentCallbackData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentCallbackData(id string, referenceId string, currency string, amount float64, country string, status string, paymentMethod PaymentMethod, created string, updated string) *PaymentCallbackData {
	this := PaymentCallbackData{}
	this.Id = id
	this.ReferenceId = referenceId
	this.Currency = currency
	this.Amount = amount
	this.Country = country
	this.Status = status
	this.PaymentMethod = paymentMethod
	this.Created = created
	this.Updated = updated
	return &this
}

// NewPaymentCallbackDataWithDefaults instantiates a new PaymentCallbackData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentCallbackDataWithDefaults() *PaymentCallbackData {
	this := PaymentCallbackData{}
	return &this
}

// GetId returns the Id field value
func (o *PaymentCallbackData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentCallbackData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentCallbackData) SetId(v string) {
	o.Id = v
}

// GetPaymentRequestId returns the PaymentRequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentCallbackData) GetPaymentRequestId() string {
	if o == nil || utils.IsNil(o.PaymentRequestId.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentRequestId.Get()
}

// GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentCallbackData) GetPaymentRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentRequestId.Get(), o.PaymentRequestId.IsSet()
}

// HasPaymentRequestId returns a boolean if a field has been set.
func (o *PaymentCallbackData) HasPaymentRequestId() bool {
	if o != nil && o.PaymentRequestId.IsSet() {
		return true
	}

	return false
}

// SetPaymentRequestId gets a reference to the given NullableString and assigns it to the PaymentRequestId field.
func (o *PaymentCallbackData) SetPaymentRequestId(v string) {
	o.PaymentRequestId.Set(&v)
}
// SetPaymentRequestIdNil sets the value for PaymentRequestId to be an explicit nil
func (o *PaymentCallbackData) SetPaymentRequestIdNil() {
	o.PaymentRequestId.Set(nil)
}

// UnsetPaymentRequestId ensures that no value is present for PaymentRequestId, not even an explicit nil
func (o *PaymentCallbackData) UnsetPaymentRequestId() {
	o.PaymentRequestId.Unset()
}

// GetReferenceId returns the ReferenceId field value
func (o *PaymentCallbackData) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *PaymentCallbackData) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *PaymentCallbackData) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentCallbackData) GetCustomerId() string {
	if o == nil || utils.IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentCallbackData) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *PaymentCallbackData) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *PaymentCallbackData) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *PaymentCallbackData) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *PaymentCallbackData) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetCurrency returns the Currency field value
func (o *PaymentCallbackData) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentCallbackData) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentCallbackData) SetCurrency(v string) {
	o.Currency = v
}

// GetAmount returns the Amount field value
func (o *PaymentCallbackData) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentCallbackData) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentCallbackData) SetAmount(v float64) {
	o.Amount = v
}

// GetCountry returns the Country field value
func (o *PaymentCallbackData) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *PaymentCallbackData) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *PaymentCallbackData) SetCountry(v string) {
	o.Country = v
}

// GetStatus returns the Status field value
func (o *PaymentCallbackData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentCallbackData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentCallbackData) SetStatus(v string) {
	o.Status = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *PaymentCallbackData) GetPaymentMethod() PaymentMethod {
	if o == nil {
		var ret PaymentMethod
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *PaymentCallbackData) GetPaymentMethodOk() (*PaymentMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *PaymentCallbackData) SetPaymentMethod(v PaymentMethod) {
	o.PaymentMethod = v
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentCallbackData) GetChannelProperties() PaymentRequestChannelProperties {
	if o == nil || utils.IsNil(o.ChannelProperties.Get()) {
		var ret PaymentRequestChannelProperties
		return ret
	}
	return *o.ChannelProperties.Get()
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentCallbackData) GetChannelPropertiesOk() (*PaymentRequestChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChannelProperties.Get(), o.ChannelProperties.IsSet()
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *PaymentCallbackData) HasChannelProperties() bool {
	if o != nil && o.ChannelProperties.IsSet() {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given NullablePaymentRequestChannelProperties and assigns it to the ChannelProperties field.
func (o *PaymentCallbackData) SetChannelProperties(v PaymentRequestChannelProperties) {
	o.ChannelProperties.Set(&v)
}
// SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil
func (o *PaymentCallbackData) SetChannelPropertiesNil() {
	o.ChannelProperties.Set(nil)
}

// UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
func (o *PaymentCallbackData) UnsetChannelProperties() {
	o.ChannelProperties.Unset()
}

// GetPaymentDetail returns the PaymentDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentCallbackData) GetPaymentDetail() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PaymentDetail
}

// GetPaymentDetailOk returns a tuple with the PaymentDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentCallbackData) GetPaymentDetailOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.PaymentDetail) {
		return map[string]interface{}{}, false
	}
	return o.PaymentDetail, true
}

// HasPaymentDetail returns a boolean if a field has been set.
func (o *PaymentCallbackData) HasPaymentDetail() bool {
	if o != nil && utils.IsNil(o.PaymentDetail) {
		return true
	}

	return false
}

// SetPaymentDetail gets a reference to the given map[string]interface{} and assigns it to the PaymentDetail field.
func (o *PaymentCallbackData) SetPaymentDetail(v map[string]interface{}) {
	o.PaymentDetail = v
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentCallbackData) GetFailureCode() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentCallbackData) GetFailureCodeOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.FailureCode) {
		return map[string]interface{}{}, false
	}
	return o.FailureCode, true
}

// HasFailureCode returns a boolean if a field has been set.
func (o *PaymentCallbackData) HasFailureCode() bool {
	if o != nil && utils.IsNil(o.FailureCode) {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given map[string]interface{} and assigns it to the FailureCode field.
func (o *PaymentCallbackData) SetFailureCode(v map[string]interface{}) {
	o.FailureCode = v
}

// GetCreated returns the Created field value
func (o *PaymentCallbackData) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *PaymentCallbackData) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *PaymentCallbackData) SetCreated(v string) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *PaymentCallbackData) GetUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *PaymentCallbackData) GetUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *PaymentCallbackData) SetUpdated(v string) {
	o.Updated = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentCallbackData) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentCallbackData) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentCallbackData) HasMetadata() bool {
	if o != nil && utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PaymentCallbackData) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PaymentCallbackData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentCallbackData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.PaymentRequestId.IsSet() {
		toSerialize["payment_request_id"] = o.PaymentRequestId.Get()
    }
	toSerialize["reference_id"] = o.ReferenceId
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
    }
	toSerialize["currency"] = o.Currency
	toSerialize["amount"] = o.Amount
	toSerialize["country"] = o.Country
	toSerialize["status"] = o.Status
	toSerialize["payment_method"] = o.PaymentMethod
	if o.ChannelProperties.IsSet() {
		toSerialize["channel_properties"] = o.ChannelProperties.Get()
    }
	if o.PaymentDetail != nil {
		toSerialize["payment_detail"] = o.PaymentDetail
    }
	if o.FailureCode != nil {
		toSerialize["failure_code"] = o.FailureCode
    }
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
    }
	return toSerialize, nil
}

type NullablePaymentCallbackData struct {
	value *PaymentCallbackData
	isSet bool
}

func (v NullablePaymentCallbackData) Get() *PaymentCallbackData {
	return v.value
}

func (v *NullablePaymentCallbackData) Set(val *PaymentCallbackData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentCallbackData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentCallbackData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentCallbackData(val *PaymentCallbackData) *NullablePaymentCallbackData {
	return &NullablePaymentCallbackData{value: val, isSet: true}
}

func (v NullablePaymentCallbackData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentCallbackData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



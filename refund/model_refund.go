/*
Refund Service

This API is used for the unified refund service

API version: 1.3.4
*/


package refund

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the Refund type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Refund{}

// Refund struct for Refund
type Refund struct {
	Id *string `json:"id,omitempty"`
	PaymentRequestId *string `json:"payment_request_id,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	ChannelCode *string `json:"channel_code,omitempty"`
	Country *string `json:"country,omitempty"`
	Currency *string `json:"currency,omitempty"`
	ReferenceId NullableString `json:"reference_id,omitempty"`
	FailureCode NullableString `json:"failure_code,omitempty"`
	RefundFeeAmount NullableFloat64 `json:"refund_fee_amount,omitempty"`
	Created *string `json:"created,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewRefund instantiates a new Refund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefund() *Refund {
	this := Refund{}
	return &this
}

// NewRefundWithDefaults instantiates a new Refund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundWithDefaults() *Refund {
	this := Refund{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Refund) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Refund) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Refund) SetId(v string) {
	o.Id = &v
}

// GetPaymentRequestId returns the PaymentRequestId field value if set, zero value otherwise.
func (o *Refund) GetPaymentRequestId() string {
	if o == nil || utils.IsNil(o.PaymentRequestId) {
		var ret string
		return ret
	}
	return *o.PaymentRequestId
}

// GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetPaymentRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentRequestId) {
		return nil, false
	}
	return o.PaymentRequestId, true
}

// HasPaymentRequestId returns a boolean if a field has been set.
func (o *Refund) HasPaymentRequestId() bool {
	if o != nil && !utils.IsNil(o.PaymentRequestId) {
		return true
	}

	return false
}

// SetPaymentRequestId gets a reference to the given string and assigns it to the PaymentRequestId field.
func (o *Refund) SetPaymentRequestId(v string) {
	o.PaymentRequestId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Refund) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetAmountOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Refund) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *Refund) SetAmount(v float64) {
	o.Amount = &v
}

// GetChannelCode returns the ChannelCode field value if set, zero value otherwise.
func (o *Refund) GetChannelCode() string {
	if o == nil || utils.IsNil(o.ChannelCode) {
		var ret string
		return ret
	}
	return *o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetChannelCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ChannelCode) {
		return nil, false
	}
	return o.ChannelCode, true
}

// HasChannelCode returns a boolean if a field has been set.
func (o *Refund) HasChannelCode() bool {
	if o != nil && !utils.IsNil(o.ChannelCode) {
		return true
	}

	return false
}

// SetChannelCode gets a reference to the given string and assigns it to the ChannelCode field.
func (o *Refund) SetChannelCode(v string) {
	o.ChannelCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Refund) GetCountry() string {
	if o == nil || utils.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetCountryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Refund) HasCountry() bool {
	if o != nil && !utils.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Refund) SetCountry(v string) {
	o.Country = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Refund) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Refund) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Refund) SetCurrency(v string) {
	o.Currency = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Refund) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// HasReferenceId returns a boolean if a field has been set.
func (o *Refund) HasReferenceId() bool {
	if o != nil && o.ReferenceId.IsSet() {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given NullableString and assigns it to the ReferenceId field.
func (o *Refund) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}
// SetReferenceIdNil sets the value for ReferenceId to be an explicit nil
func (o *Refund) SetReferenceIdNil() {
	o.ReferenceId.Set(nil)
}

// UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
func (o *Refund) UnsetReferenceId() {
	o.ReferenceId.Unset()
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Refund) GetFailureCode() string {
	if o == nil || utils.IsNil(o.FailureCode.Get()) {
		var ret string
		return ret
	}
	return *o.FailureCode.Get()
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetFailureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureCode.Get(), o.FailureCode.IsSet()
}

// HasFailureCode returns a boolean if a field has been set.
func (o *Refund) HasFailureCode() bool {
	if o != nil && o.FailureCode.IsSet() {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given NullableString and assigns it to the FailureCode field.
func (o *Refund) SetFailureCode(v string) {
	o.FailureCode.Set(&v)
}
// SetFailureCodeNil sets the value for FailureCode to be an explicit nil
func (o *Refund) SetFailureCodeNil() {
	o.FailureCode.Set(nil)
}

// UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
func (o *Refund) UnsetFailureCode() {
	o.FailureCode.Unset()
}

// GetRefundFeeAmount returns the RefundFeeAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Refund) GetRefundFeeAmount() float64 {
	if o == nil || utils.IsNil(o.RefundFeeAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.RefundFeeAmount.Get()
}

// GetRefundFeeAmountOk returns a tuple with the RefundFeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetRefundFeeAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefundFeeAmount.Get(), o.RefundFeeAmount.IsSet()
}

// HasRefundFeeAmount returns a boolean if a field has been set.
func (o *Refund) HasRefundFeeAmount() bool {
	if o != nil && o.RefundFeeAmount.IsSet() {
		return true
	}

	return false
}

// SetRefundFeeAmount gets a reference to the given NullableFloat64 and assigns it to the RefundFeeAmount field.
func (o *Refund) SetRefundFeeAmount(v float64) {
	o.RefundFeeAmount.Set(&v)
}
// SetRefundFeeAmountNil sets the value for RefundFeeAmount to be an explicit nil
func (o *Refund) SetRefundFeeAmountNil() {
	o.RefundFeeAmount.Set(nil)
}

// UnsetRefundFeeAmount ensures that no value is present for RefundFeeAmount, not even an explicit nil
func (o *Refund) UnsetRefundFeeAmount() {
	o.RefundFeeAmount.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Refund) GetCreated() string {
	if o == nil || utils.IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetCreatedOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Refund) HasCreated() bool {
	if o != nil && !utils.IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Refund) SetCreated(v string) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Refund) GetUpdated() string {
	if o == nil || utils.IsNil(o.Updated) {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Refund) GetUpdatedOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Refund) HasUpdated() bool {
	if o != nil && !utils.IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *Refund) SetUpdated(v string) {
	o.Updated = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Refund) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Refund) HasMetadata() bool {
	if o != nil && utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Refund) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o Refund) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Refund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.PaymentRequestId) {
		toSerialize["payment_request_id"] = o.PaymentRequestId
	}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !utils.IsNil(o.ChannelCode) {
		toSerialize["channel_code"] = o.ChannelCode
	}
	if !utils.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if o.ReferenceId.IsSet() {
		toSerialize["reference_id"] = o.ReferenceId.Get()
    }
	if o.FailureCode.IsSet() {
		toSerialize["failure_code"] = o.FailureCode.Get()
    }
	if o.RefundFeeAmount.IsSet() {
		toSerialize["refund_fee_amount"] = o.RefundFeeAmount.Get()
    }
	if !utils.IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !utils.IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
    }
	return toSerialize, nil
}

type NullableRefund struct {
	value *Refund
	isSet bool
}

func (v NullableRefund) Get() *Refund {
	return v.value
}

func (v *NullableRefund) Set(val *Refund) {
	v.value = val
	v.isSet = true
}

func (v NullableRefund) IsSet() bool {
	return v.isSet
}

func (v *NullableRefund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefund(val *Refund) *NullableRefund {
	return &NullableRefund{value: val, isSet: true}
}

func (v NullableRefund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



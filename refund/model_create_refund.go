/*
Refund Service

This API is used for the unified refund service

API version: 1.2.3
*/


package refund

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the CreateRefund type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateRefund{}

// CreateRefund struct for CreateRefund
type CreateRefund struct {
	PaymentRequestId *string `json:"payment_request_id,omitempty"`
	InvoiceId *string `json:"invoice_id,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewCreateRefund instantiates a new CreateRefund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRefund() *CreateRefund {
	this := CreateRefund{}
	return &this
}

// NewCreateRefundWithDefaults instantiates a new CreateRefund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRefundWithDefaults() *CreateRefund {
	this := CreateRefund{}
	return &this
}

// GetPaymentRequestId returns the PaymentRequestId field value if set, zero value otherwise.
func (o *CreateRefund) GetPaymentRequestId() string {
	if o == nil || utils.IsNil(o.PaymentRequestId) {
		var ret string
		return ret
	}
	return *o.PaymentRequestId
}

// GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund) GetPaymentRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentRequestId) {
		return nil, false
	}
	return o.PaymentRequestId, true
}

// HasPaymentRequestId returns a boolean if a field has been set.
func (o *CreateRefund) HasPaymentRequestId() bool {
	if o != nil && !utils.IsNil(o.PaymentRequestId) {
		return true
	}

	return false
}

// SetPaymentRequestId gets a reference to the given string and assigns it to the PaymentRequestId field.
func (o *CreateRefund) SetPaymentRequestId(v string) {
	o.PaymentRequestId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *CreateRefund) GetInvoiceId() string {
	if o == nil || utils.IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund) GetInvoiceIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *CreateRefund) HasInvoiceId() bool {
	if o != nil && !utils.IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *CreateRefund) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *CreateRefund) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund) GetReferenceIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *CreateRefund) HasReferenceId() bool {
	if o != nil && !utils.IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *CreateRefund) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CreateRefund) GetAmount() float64 {
	if o == nil || utils.IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund) GetAmountOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CreateRefund) HasAmount() bool {
	if o != nil && !utils.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *CreateRefund) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CreateRefund) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CreateRefund) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CreateRefund) SetCurrency(v string) {
	o.Currency = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CreateRefund) GetReason() string {
	if o == nil || utils.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRefund) GetReasonOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CreateRefund) HasReason() bool {
	if o != nil && !utils.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CreateRefund) SetReason(v string) {
	o.Reason = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRefund) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRefund) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateRefund) HasMetadata() bool {
	if o != nil && utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CreateRefund) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CreateRefund) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRefund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PaymentRequestId) {
		toSerialize["payment_request_id"] = o.PaymentRequestId
	}
	if !utils.IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if !utils.IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !utils.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !utils.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableCreateRefund struct {
	value *CreateRefund
	isSet bool
}

func (v NullableCreateRefund) Get() *CreateRefund {
	return v.value
}

func (v *NullableCreateRefund) Set(val *CreateRefund) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRefund) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRefund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRefund(val *CreateRefund) *NullableCreateRefund {
	return &NullableCreateRefund{value: val, isSet: true}
}

func (v NullableCreateRefund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRefund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



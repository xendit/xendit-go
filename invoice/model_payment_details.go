/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.7.6
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the PaymentDetails type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentDetails{}

// PaymentDetails An object representing payment details.
type PaymentDetails struct {
	// The unique identifier or reference ID associated with the payment receipt.
	ReceiptId *string `json:"receipt_id,omitempty"`
	// The source or method of payment.
	Source *string `json:"source,omitempty"`
}

// NewPaymentDetails instantiates a new PaymentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentDetails() *PaymentDetails {
	this := PaymentDetails{}
	return &this
}

// NewPaymentDetailsWithDefaults instantiates a new PaymentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentDetailsWithDefaults() *PaymentDetails {
	this := PaymentDetails{}
	return &this
}

// GetReceiptId returns the ReceiptId field value if set, zero value otherwise.
func (o *PaymentDetails) GetReceiptId() string {
	if o == nil || utils.IsNil(o.ReceiptId) {
		var ret string
		return ret
	}
	return *o.ReceiptId
}

// GetReceiptIdOk returns a tuple with the ReceiptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetails) GetReceiptIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReceiptId) {
		return nil, false
	}
	return o.ReceiptId, true
}

// HasReceiptId returns a boolean if a field has been set.
func (o *PaymentDetails) HasReceiptId() bool {
	if o != nil && !utils.IsNil(o.ReceiptId) {
		return true
	}

	return false
}

// SetReceiptId gets a reference to the given string and assigns it to the ReceiptId field.
func (o *PaymentDetails) SetReceiptId(v string) {
	o.ReceiptId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PaymentDetails) GetSource() string {
	if o == nil || utils.IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentDetails) GetSourceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PaymentDetails) HasSource() bool {
	if o != nil && !utils.IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PaymentDetails) SetSource(v string) {
	o.Source = &v
}

func (o PaymentDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ReceiptId) {
		toSerialize["receipt_id"] = o.ReceiptId
	}
	if !utils.IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullablePaymentDetails struct {
	value *PaymentDetails
	isSet bool
}

func (v NullablePaymentDetails) Get() *PaymentDetails {
	return v.value
}

func (v *NullablePaymentDetails) Set(val *PaymentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentDetails(val *PaymentDetails) *NullablePaymentDetails {
	return &NullablePaymentDetails{value: val, isSet: true}
}

func (v NullablePaymentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



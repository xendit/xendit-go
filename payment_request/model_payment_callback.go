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

// checks if the PaymentCallback type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentCallback{}

// PaymentCallback Callback for successful or failed payments made via the Payments API
type PaymentCallback struct {
	// Identifies the event that triggered a notification to the merchant
	Event string `json:"event"`
	// business_id
	BusinessId string `json:"business_id"`
	Created string `json:"created"`
	Data *PaymentCallbackData `json:"data,omitempty"`
}

// NewPaymentCallback instantiates a new PaymentCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentCallback(event string, businessId string, created string) *PaymentCallback {
	this := PaymentCallback{}
	this.Event = event
	this.BusinessId = businessId
	this.Created = created
	return &this
}

// NewPaymentCallbackWithDefaults instantiates a new PaymentCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentCallbackWithDefaults() *PaymentCallback {
	this := PaymentCallback{}
	return &this
}

// GetEvent returns the Event field value
func (o *PaymentCallback) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *PaymentCallback) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *PaymentCallback) SetEvent(v string) {
	o.Event = v
}

// GetBusinessId returns the BusinessId field value
func (o *PaymentCallback) GetBusinessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *PaymentCallback) GetBusinessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *PaymentCallback) SetBusinessId(v string) {
	o.BusinessId = v
}

// GetCreated returns the Created field value
func (o *PaymentCallback) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *PaymentCallback) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *PaymentCallback) SetCreated(v string) {
	o.Created = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PaymentCallback) GetData() PaymentCallbackData {
	if o == nil || utils.IsNil(o.Data) {
		var ret PaymentCallbackData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentCallback) GetDataOk() (*PaymentCallbackData, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PaymentCallback) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PaymentCallbackData and assigns it to the Data field.
func (o *PaymentCallback) SetData(v PaymentCallbackData) {
	o.Data = &v
}

func (o PaymentCallback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentCallback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["business_id"] = o.BusinessId
	toSerialize["created"] = o.Created
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePaymentCallback struct {
	value *PaymentCallback
	isSet bool
}

func (v NullablePaymentCallback) Get() *PaymentCallback {
	return v.value
}

func (v *NullablePaymentCallback) Set(val *PaymentCallback) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentCallback) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentCallback(val *PaymentCallback) *NullablePaymentCallback {
	return &NullablePaymentCallback{value: val, isSet: true}
}

func (v NullablePaymentCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



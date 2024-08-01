/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the PaymentMethodCallback type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentMethodCallback{}

// PaymentMethodCallback Callback for active or expired E-Wallet or Direct Debit account linking, Virtual Accounts or QR strings
type PaymentMethodCallback struct {
	// Identifies the event that triggered a notification to the merchant
	Event string `json:"event"`
	// business_id
	BusinessId string `json:"business_id"`
	Created string `json:"created"`
	Data *PaymentMethod `json:"data,omitempty"`
}

// NewPaymentMethodCallback instantiates a new PaymentMethodCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCallback(event string, businessId string, created string) *PaymentMethodCallback {
	this := PaymentMethodCallback{}
	this.Event = event
	this.BusinessId = businessId
	this.Created = created
	return &this
}

// NewPaymentMethodCallbackWithDefaults instantiates a new PaymentMethodCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCallbackWithDefaults() *PaymentMethodCallback {
	this := PaymentMethodCallback{}
	return &this
}

// GetEvent returns the Event field value
func (o *PaymentMethodCallback) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCallback) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *PaymentMethodCallback) SetEvent(v string) {
	o.Event = v
}

// GetBusinessId returns the BusinessId field value
func (o *PaymentMethodCallback) GetBusinessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCallback) GetBusinessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *PaymentMethodCallback) SetBusinessId(v string) {
	o.BusinessId = v
}

// GetCreated returns the Created field value
func (o *PaymentMethodCallback) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCallback) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *PaymentMethodCallback) SetCreated(v string) {
	o.Created = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PaymentMethodCallback) GetData() PaymentMethod {
	if o == nil || utils.IsNil(o.Data) {
		var ret PaymentMethod
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCallback) GetDataOk() (*PaymentMethod, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PaymentMethodCallback) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PaymentMethod and assigns it to the Data field.
func (o *PaymentMethodCallback) SetData(v PaymentMethod) {
	o.Data = &v
}

func (o PaymentMethodCallback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCallback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["business_id"] = o.BusinessId
	toSerialize["created"] = o.Created
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePaymentMethodCallback struct {
	value *PaymentMethodCallback
	isSet bool
}

func (v NullablePaymentMethodCallback) Get() *PaymentMethodCallback {
	return v.value
}

func (v *NullablePaymentMethodCallback) Set(val *PaymentMethodCallback) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCallback) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCallback(val *PaymentMethodCallback) *NullablePaymentMethodCallback {
	return &NullablePaymentMethodCallback{value: val, isSet: true}
}

func (v NullablePaymentMethodCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



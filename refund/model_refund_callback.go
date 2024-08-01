/*
Refund Service

This API is used for the unified refund service

API version: 1.3.4
*/


package refund

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the RefundCallback type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RefundCallback{}

// RefundCallback Callback for successful or failed Refunds made via the Payments API
type RefundCallback struct {
	// Identifies the event that triggered a notification to the merchant
	Event string `json:"event"`
	// business_id
	BusinessId string `json:"business_id"`
	Created string `json:"created"`
	Data *RefundCallbackData `json:"data,omitempty"`
}

// NewRefundCallback instantiates a new RefundCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundCallback(event string, businessId string, created string) *RefundCallback {
	this := RefundCallback{}
	this.Event = event
	this.BusinessId = businessId
	this.Created = created
	return &this
}

// NewRefundCallbackWithDefaults instantiates a new RefundCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundCallbackWithDefaults() *RefundCallback {
	this := RefundCallback{}
	return &this
}

// GetEvent returns the Event field value
func (o *RefundCallback) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *RefundCallback) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *RefundCallback) SetEvent(v string) {
	o.Event = v
}

// GetBusinessId returns the BusinessId field value
func (o *RefundCallback) GetBusinessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *RefundCallback) GetBusinessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *RefundCallback) SetBusinessId(v string) {
	o.BusinessId = v
}

// GetCreated returns the Created field value
func (o *RefundCallback) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RefundCallback) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RefundCallback) SetCreated(v string) {
	o.Created = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RefundCallback) GetData() RefundCallbackData {
	if o == nil || utils.IsNil(o.Data) {
		var ret RefundCallbackData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefundCallback) GetDataOk() (*RefundCallbackData, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RefundCallback) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RefundCallbackData and assigns it to the Data field.
func (o *RefundCallback) SetData(v RefundCallbackData) {
	o.Data = &v
}

func (o RefundCallback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundCallback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["business_id"] = o.BusinessId
	toSerialize["created"] = o.Created
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRefundCallback struct {
	value *RefundCallback
	isSet bool
}

func (v NullableRefundCallback) Get() *RefundCallback {
	return v.value
}

func (v *NullableRefundCallback) Set(val *RefundCallback) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundCallback(val *RefundCallback) *NullableRefundCallback {
	return &NullableRefundCallback{value: val, isSet: true}
}

func (v NullableRefundCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



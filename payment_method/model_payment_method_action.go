/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the PaymentMethodAction type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentMethodAction{}

// PaymentMethodAction struct for PaymentMethodAction
type PaymentMethodAction struct {
	Action *string `json:"action,omitempty"`
	Method *string `json:"method,omitempty"`
	Url *string `json:"url,omitempty"`
	UrlType *string `json:"url_type,omitempty"`
}

// NewPaymentMethodAction instantiates a new PaymentMethodAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodAction() *PaymentMethodAction {
	this := PaymentMethodAction{}
	return &this
}

// NewPaymentMethodActionWithDefaults instantiates a new PaymentMethodAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodActionWithDefaults() *PaymentMethodAction {
	this := PaymentMethodAction{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PaymentMethodAction) GetAction() string {
	if o == nil || utils.IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodAction) GetActionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PaymentMethodAction) HasAction() bool {
	if o != nil && !utils.IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PaymentMethodAction) SetAction(v string) {
	o.Action = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PaymentMethodAction) GetMethod() string {
	if o == nil || utils.IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodAction) GetMethodOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PaymentMethodAction) HasMethod() bool {
	if o != nil && !utils.IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PaymentMethodAction) SetMethod(v string) {
	o.Method = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PaymentMethodAction) GetUrl() string {
	if o == nil || utils.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodAction) GetUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PaymentMethodAction) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PaymentMethodAction) SetUrl(v string) {
	o.Url = &v
}

// GetUrlType returns the UrlType field value if set, zero value otherwise.
func (o *PaymentMethodAction) GetUrlType() string {
	if o == nil || utils.IsNil(o.UrlType) {
		var ret string
		return ret
	}
	return *o.UrlType
}

// GetUrlTypeOk returns a tuple with the UrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodAction) GetUrlTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.UrlType) {
		return nil, false
	}
	return o.UrlType, true
}

// HasUrlType returns a boolean if a field has been set.
func (o *PaymentMethodAction) HasUrlType() bool {
	if o != nil && !utils.IsNil(o.UrlType) {
		return true
	}

	return false
}

// SetUrlType gets a reference to the given string and assigns it to the UrlType field.
func (o *PaymentMethodAction) SetUrlType(v string) {
	o.UrlType = &v
}

func (o PaymentMethodAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !utils.IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !utils.IsNil(o.UrlType) {
		toSerialize["url_type"] = o.UrlType
	}
	return toSerialize, nil
}

type NullablePaymentMethodAction struct {
	value *PaymentMethodAction
	isSet bool
}

func (v NullablePaymentMethodAction) Get() *PaymentMethodAction {
	return v.value
}

func (v *NullablePaymentMethodAction) Set(val *PaymentMethodAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodAction(val *PaymentMethodAction) *NullablePaymentMethodAction {
	return &NullablePaymentMethodAction{value: val, isSet: true}
}

func (v NullablePaymentMethodAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



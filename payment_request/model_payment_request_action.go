/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.1
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the PaymentRequestAction type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestAction{}

// PaymentRequestAction struct for PaymentRequestAction
type PaymentRequestAction struct {
	Action string `json:"action"`
	UrlType string `json:"url_type"`
	Method NullableString `json:"method"`
	Url NullableString `json:"url"`
	QrCode NullableString `json:"qr_code"`
}

// NewPaymentRequestAction instantiates a new PaymentRequestAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestAction(action string, urlType string, method NullableString, url NullableString, qrCode NullableString) *PaymentRequestAction {
	this := PaymentRequestAction{}
	this.Action = action
	this.UrlType = urlType
	this.Method = method
	this.Url = url
	this.QrCode = qrCode
	return &this
}

// NewPaymentRequestActionWithDefaults instantiates a new PaymentRequestAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestActionWithDefaults() *PaymentRequestAction {
	this := PaymentRequestAction{}
	return &this
}

// GetAction returns the Action field value
func (o *PaymentRequestAction) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestAction) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PaymentRequestAction) SetAction(v string) {
	o.Action = v
}

// GetUrlType returns the UrlType field value
func (o *PaymentRequestAction) GetUrlType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlType
}

// GetUrlTypeOk returns a tuple with the UrlType field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestAction) GetUrlTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlType, true
}

// SetUrlType sets field value
func (o *PaymentRequestAction) SetUrlType(v string) {
	o.UrlType = v
}

// GetMethod returns the Method field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentRequestAction) GetMethod() string {
	if o == nil || o.Method.Get() == nil {
		var ret string
		return ret
	}

	return *o.Method.Get()
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestAction) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Method.Get(), o.Method.IsSet()
}

// SetMethod sets field value
func (o *PaymentRequestAction) SetMethod(v string) {
	o.Method.Set(&v)
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentRequestAction) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestAction) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *PaymentRequestAction) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetQrCode returns the QrCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentRequestAction) GetQrCode() string {
	if o == nil || o.QrCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.QrCode.Get()
}

// GetQrCodeOk returns a tuple with the QrCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestAction) GetQrCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QrCode.Get(), o.QrCode.IsSet()
}

// SetQrCode sets field value
func (o *PaymentRequestAction) SetQrCode(v string) {
	o.QrCode.Set(&v)
}

func (o PaymentRequestAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
    if o.Action != "AUTH" && o.Action != "RESEND_AUTH" && o.Action != "CAPTURE" && o.Action != "CANCEL" && o.Action != "PRESENT_TO_CUSTOMER" {
        toSerialize["action"] = nil
        return toSerialize, utils.NewError("invalid value for Action when marshalling to JSON, must be one of AUTH, RESEND_AUTH, CAPTURE, CANCEL, PRESENT_TO_CUSTOMER")
    }
	toSerialize["url_type"] = o.UrlType
    if o.UrlType != "API" && o.UrlType != "WEB" && o.UrlType != "MOBILE" && o.UrlType != "DEEPLINK" {
        toSerialize["url_type"] = nil
        return toSerialize, utils.NewError("invalid value for UrlType when marshalling to JSON, must be one of API, WEB, MOBILE, DEEPLINK")
    }
	toSerialize["method"] = o.Method.Get()

	toSerialize["url"] = o.Url.Get()

	toSerialize["qr_code"] = o.QrCode.Get()

	return toSerialize, nil
}

type NullablePaymentRequestAction struct {
	value *PaymentRequestAction
	isSet bool
}

func (v NullablePaymentRequestAction) Get() *PaymentRequestAction {
	return v.value
}

func (v *NullablePaymentRequestAction) Set(val *PaymentRequestAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestAction(val *PaymentRequestAction) *NullablePaymentRequestAction {
	return &NullablePaymentRequestAction{value: val, isSet: true}
}

func (v NullablePaymentRequestAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



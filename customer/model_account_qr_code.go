/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the AccountQRCode type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AccountQRCode{}

// AccountQRCode struct for AccountQRCode
type AccountQRCode struct {
	// String representation of the QR Code image
	QrString *string `json:"qr_string,omitempty"`
}

// NewAccountQRCode instantiates a new AccountQRCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountQRCode() *AccountQRCode {
	this := AccountQRCode{}
	return &this
}

// NewAccountQRCodeWithDefaults instantiates a new AccountQRCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountQRCodeWithDefaults() *AccountQRCode {
	this := AccountQRCode{}
	return &this
}

// GetQrString returns the QrString field value if set, zero value otherwise.
func (o *AccountQRCode) GetQrString() string {
	if o == nil || utils.IsNil(o.QrString) {
		var ret string
		return ret
	}
	return *o.QrString
}

// GetQrStringOk returns a tuple with the QrString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountQRCode) GetQrStringOk() (*string, bool) {
	if o == nil || utils.IsNil(o.QrString) {
		return nil, false
	}
	return o.QrString, true
}

// HasQrString returns a boolean if a field has been set.
func (o *AccountQRCode) HasQrString() bool {
	if o != nil && !utils.IsNil(o.QrString) {
		return true
	}

	return false
}

// SetQrString gets a reference to the given string and assigns it to the QrString field.
func (o *AccountQRCode) SetQrString(v string) {
	o.QrString = &v
}

func (o AccountQRCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountQRCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.QrString) {
		toSerialize["qr_string"] = o.QrString
	}
	return toSerialize, nil
}

type NullableAccountQRCode struct {
	value *AccountQRCode
	isSet bool
}

func (v NullableAccountQRCode) Get() *AccountQRCode {
	return v.value
}

func (v *NullableAccountQRCode) Set(val *AccountQRCode) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountQRCode) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountQRCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountQRCode(val *AccountQRCode) *NullableAccountQRCode {
	return &NullableAccountQRCode{value: val, isSet: true}
}

func (v NullableAccountQRCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountQRCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



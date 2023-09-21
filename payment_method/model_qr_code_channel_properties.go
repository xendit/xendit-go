/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.87.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
	"time"
)

// checks if the QRCodeChannelProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QRCodeChannelProperties{}

// QRCodeChannelProperties QR Code Channel Properties
type QRCodeChannelProperties struct {
	// QR string to be rendered for display to end users. QR string to image rendering are commonly available in software libraries (e.g Nodejs, PHP, Java)
	QrString *string `json:"qr_string,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// NewQRCodeChannelProperties instantiates a new QRCodeChannelProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQRCodeChannelProperties() *QRCodeChannelProperties {
	this := QRCodeChannelProperties{}
	return &this
}

// NewQRCodeChannelPropertiesWithDefaults instantiates a new QRCodeChannelProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQRCodeChannelPropertiesWithDefaults() *QRCodeChannelProperties {
	this := QRCodeChannelProperties{}
	return &this
}

// GetQrString returns the QrString field value if set, zero value otherwise.
func (o *QRCodeChannelProperties) GetQrString() string {
	if o == nil || utils.IsNil(o.QrString) {
		var ret string
		return ret
	}
	return *o.QrString
}

// GetQrStringOk returns a tuple with the QrString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QRCodeChannelProperties) GetQrStringOk() (*string, bool) {
	if o == nil || utils.IsNil(o.QrString) {
		return nil, false
	}
	return o.QrString, true
}

// HasQrString returns a boolean if a field has been set.
func (o *QRCodeChannelProperties) HasQrString() bool {
	if o != nil && !utils.IsNil(o.QrString) {
		return true
	}

	return false
}

// SetQrString gets a reference to the given string and assigns it to the QrString field.
func (o *QRCodeChannelProperties) SetQrString(v string) {
	o.QrString = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *QRCodeChannelProperties) GetExpiresAt() time.Time {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QRCodeChannelProperties) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *QRCodeChannelProperties) HasExpiresAt() bool {
	if o != nil && !utils.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *QRCodeChannelProperties) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o QRCodeChannelProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QRCodeChannelProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.QrString) {
		toSerialize["qr_string"] = o.QrString
	}
	if !utils.IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableQRCodeChannelProperties struct {
	value *QRCodeChannelProperties
	isSet bool
}

func (v NullableQRCodeChannelProperties) Get() *QRCodeChannelProperties {
	return v.value
}

func (v *NullableQRCodeChannelProperties) Set(val *QRCodeChannelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableQRCodeChannelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableQRCodeChannelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQRCodeChannelProperties(val *QRCodeChannelProperties) *NullableQRCodeChannelProperties {
	return &NullableQRCodeChannelProperties{value: val, isSet: true}
}

func (v NullableQRCodeChannelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQRCodeChannelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



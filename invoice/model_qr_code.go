/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.8.7
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the QrCode type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QrCode{}

// QrCode An object representing QR code details for invoices.
type QrCode struct {
	QrCodeType QrCodeType `json:"qr_code_type"`
}

// NewQrCode instantiates a new QrCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQrCode(qrCodeType QrCodeType) *QrCode {
	this := QrCode{}
	this.QrCodeType = qrCodeType
	return &this
}

// NewQrCodeWithDefaults instantiates a new QrCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQrCodeWithDefaults() *QrCode {
	this := QrCode{}
	return &this
}

// GetQrCodeType returns the QrCodeType field value
func (o *QrCode) GetQrCodeType() QrCodeType {
	if o == nil {
		var ret QrCodeType
		return ret
	}

	return o.QrCodeType
}

// GetQrCodeTypeOk returns a tuple with the QrCodeType field value
// and a boolean to check if the value has been set.
func (o *QrCode) GetQrCodeTypeOk() (*QrCodeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QrCodeType, true
}

// SetQrCodeType sets field value
func (o *QrCode) SetQrCodeType(v QrCodeType) {
	o.QrCodeType = v
}

func (o QrCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QrCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["qr_code_type"] = o.QrCodeType
	return toSerialize, nil
}

type NullableQrCode struct {
	value *QrCode
	isSet bool
}

func (v NullableQrCode) Get() *QrCode {
	return v.value
}

func (v *NullableQrCode) Set(val *QrCode) {
	v.value = val
	v.isSet = true
}

func (v NullableQrCode) IsSet() bool {
	return v.isSet
}

func (v *NullableQrCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQrCode(val *QrCode) *NullableQrCode {
	return &NullableQrCode{value: val, isSet: true}
}

func (v NullableQrCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQrCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the PaymentRequestShippingInformation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestShippingInformation{}

// PaymentRequestShippingInformation struct for PaymentRequestShippingInformation
type PaymentRequestShippingInformation struct {
	Country string `json:"country"`
	StreetLine1 *string `json:"street_line1,omitempty"`
	StreetLine2 *string `json:"street_line2,omitempty"`
	City *string `json:"city,omitempty"`
	ProvinceState *string `json:"province_state,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
}

// NewPaymentRequestShippingInformation instantiates a new PaymentRequestShippingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestShippingInformation(country string) *PaymentRequestShippingInformation {
	this := PaymentRequestShippingInformation{}
	this.Country = country
	return &this
}

// NewPaymentRequestShippingInformationWithDefaults instantiates a new PaymentRequestShippingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestShippingInformationWithDefaults() *PaymentRequestShippingInformation {
	this := PaymentRequestShippingInformation{}
	return &this
}

// GetCountry returns the Country field value
func (o *PaymentRequestShippingInformation) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestShippingInformation) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *PaymentRequestShippingInformation) SetCountry(v string) {
	o.Country = v
}

// GetStreetLine1 returns the StreetLine1 field value if set, zero value otherwise.
func (o *PaymentRequestShippingInformation) GetStreetLine1() string {
	if o == nil || utils.IsNil(o.StreetLine1) {
		var ret string
		return ret
	}
	return *o.StreetLine1
}

// GetStreetLine1Ok returns a tuple with the StreetLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestShippingInformation) GetStreetLine1Ok() (*string, bool) {
	if o == nil || utils.IsNil(o.StreetLine1) {
		return nil, false
	}
	return o.StreetLine1, true
}

// HasStreetLine1 returns a boolean if a field has been set.
func (o *PaymentRequestShippingInformation) HasStreetLine1() bool {
	if o != nil && !utils.IsNil(o.StreetLine1) {
		return true
	}

	return false
}

// SetStreetLine1 gets a reference to the given string and assigns it to the StreetLine1 field.
func (o *PaymentRequestShippingInformation) SetStreetLine1(v string) {
	o.StreetLine1 = &v
}

// GetStreetLine2 returns the StreetLine2 field value if set, zero value otherwise.
func (o *PaymentRequestShippingInformation) GetStreetLine2() string {
	if o == nil || utils.IsNil(o.StreetLine2) {
		var ret string
		return ret
	}
	return *o.StreetLine2
}

// GetStreetLine2Ok returns a tuple with the StreetLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestShippingInformation) GetStreetLine2Ok() (*string, bool) {
	if o == nil || utils.IsNil(o.StreetLine2) {
		return nil, false
	}
	return o.StreetLine2, true
}

// HasStreetLine2 returns a boolean if a field has been set.
func (o *PaymentRequestShippingInformation) HasStreetLine2() bool {
	if o != nil && !utils.IsNil(o.StreetLine2) {
		return true
	}

	return false
}

// SetStreetLine2 gets a reference to the given string and assigns it to the StreetLine2 field.
func (o *PaymentRequestShippingInformation) SetStreetLine2(v string) {
	o.StreetLine2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *PaymentRequestShippingInformation) GetCity() string {
	if o == nil || utils.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestShippingInformation) GetCityOk() (*string, bool) {
	if o == nil || utils.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *PaymentRequestShippingInformation) HasCity() bool {
	if o != nil && !utils.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *PaymentRequestShippingInformation) SetCity(v string) {
	o.City = &v
}

// GetProvinceState returns the ProvinceState field value if set, zero value otherwise.
func (o *PaymentRequestShippingInformation) GetProvinceState() string {
	if o == nil || utils.IsNil(o.ProvinceState) {
		var ret string
		return ret
	}
	return *o.ProvinceState
}

// GetProvinceStateOk returns a tuple with the ProvinceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestShippingInformation) GetProvinceStateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ProvinceState) {
		return nil, false
	}
	return o.ProvinceState, true
}

// HasProvinceState returns a boolean if a field has been set.
func (o *PaymentRequestShippingInformation) HasProvinceState() bool {
	if o != nil && !utils.IsNil(o.ProvinceState) {
		return true
	}

	return false
}

// SetProvinceState gets a reference to the given string and assigns it to the ProvinceState field.
func (o *PaymentRequestShippingInformation) SetProvinceState(v string) {
	o.ProvinceState = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *PaymentRequestShippingInformation) GetPostalCode() string {
	if o == nil || utils.IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestShippingInformation) GetPostalCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *PaymentRequestShippingInformation) HasPostalCode() bool {
	if o != nil && !utils.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *PaymentRequestShippingInformation) SetPostalCode(v string) {
	o.PostalCode = &v
}

func (o PaymentRequestShippingInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestShippingInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country"] = o.Country
	if !utils.IsNil(o.StreetLine1) {
		toSerialize["street_line1"] = o.StreetLine1
	}
	if !utils.IsNil(o.StreetLine2) {
		toSerialize["street_line2"] = o.StreetLine2
	}
	if !utils.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !utils.IsNil(o.ProvinceState) {
		toSerialize["province_state"] = o.ProvinceState
	}
	if !utils.IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	return toSerialize, nil
}

type NullablePaymentRequestShippingInformation struct {
	value *PaymentRequestShippingInformation
	isSet bool
}

func (v NullablePaymentRequestShippingInformation) Get() *PaymentRequestShippingInformation {
	return v.value
}

func (v *NullablePaymentRequestShippingInformation) Set(val *PaymentRequestShippingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestShippingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestShippingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestShippingInformation(val *PaymentRequestShippingInformation) *NullablePaymentRequestShippingInformation {
	return &NullablePaymentRequestShippingInformation{value: val, isSet: true}
}

func (v NullablePaymentRequestShippingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestShippingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



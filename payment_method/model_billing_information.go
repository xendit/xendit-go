/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.87.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the BillingInformation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BillingInformation{}

// BillingInformation Billing Information
type BillingInformation struct {
	Country string `json:"country"`
	StreetLine1 NullableString `json:"street_line1,omitempty"`
	StreetLine2 NullableString `json:"street_line2,omitempty"`
	City NullableString `json:"city,omitempty"`
	ProvinceState NullableString `json:"province_state,omitempty"`
	PostalCode NullableString `json:"postal_code,omitempty"`
}

// NewBillingInformation instantiates a new BillingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInformation(country string) *BillingInformation {
	this := BillingInformation{}
	this.Country = country
	return &this
}

// NewBillingInformationWithDefaults instantiates a new BillingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInformationWithDefaults() *BillingInformation {
	this := BillingInformation{}
	return &this
}

// GetCountry returns the Country field value
func (o *BillingInformation) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *BillingInformation) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *BillingInformation) SetCountry(v string) {
	o.Country = v
}

// GetStreetLine1 returns the StreetLine1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInformation) GetStreetLine1() string {
	if o == nil || utils.IsNil(o.StreetLine1.Get()) {
		var ret string
		return ret
	}
	return *o.StreetLine1.Get()
}

// GetStreetLine1Ok returns a tuple with the StreetLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInformation) GetStreetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreetLine1.Get(), o.StreetLine1.IsSet()
}

// HasStreetLine1 returns a boolean if a field has been set.
func (o *BillingInformation) HasStreetLine1() bool {
	if o != nil && o.StreetLine1.IsSet() {
		return true
	}

	return false
}

// SetStreetLine1 gets a reference to the given NullableString and assigns it to the StreetLine1 field.
func (o *BillingInformation) SetStreetLine1(v string) {
	o.StreetLine1.Set(&v)
}
// SetStreetLine1Nil sets the value for StreetLine1 to be an explicit nil
func (o *BillingInformation) SetStreetLine1Nil() {
	o.StreetLine1.Set(nil)
}

// UnsetStreetLine1 ensures that no value is present for StreetLine1, not even an explicit nil
func (o *BillingInformation) UnsetStreetLine1() {
	o.StreetLine1.Unset()
}

// GetStreetLine2 returns the StreetLine2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInformation) GetStreetLine2() string {
	if o == nil || utils.IsNil(o.StreetLine2.Get()) {
		var ret string
		return ret
	}
	return *o.StreetLine2.Get()
}

// GetStreetLine2Ok returns a tuple with the StreetLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInformation) GetStreetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreetLine2.Get(), o.StreetLine2.IsSet()
}

// HasStreetLine2 returns a boolean if a field has been set.
func (o *BillingInformation) HasStreetLine2() bool {
	if o != nil && o.StreetLine2.IsSet() {
		return true
	}

	return false
}

// SetStreetLine2 gets a reference to the given NullableString and assigns it to the StreetLine2 field.
func (o *BillingInformation) SetStreetLine2(v string) {
	o.StreetLine2.Set(&v)
}
// SetStreetLine2Nil sets the value for StreetLine2 to be an explicit nil
func (o *BillingInformation) SetStreetLine2Nil() {
	o.StreetLine2.Set(nil)
}

// UnsetStreetLine2 ensures that no value is present for StreetLine2, not even an explicit nil
func (o *BillingInformation) UnsetStreetLine2() {
	o.StreetLine2.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInformation) GetCity() string {
	if o == nil || utils.IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInformation) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *BillingInformation) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *BillingInformation) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *BillingInformation) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *BillingInformation) UnsetCity() {
	o.City.Unset()
}

// GetProvinceState returns the ProvinceState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInformation) GetProvinceState() string {
	if o == nil || utils.IsNil(o.ProvinceState.Get()) {
		var ret string
		return ret
	}
	return *o.ProvinceState.Get()
}

// GetProvinceStateOk returns a tuple with the ProvinceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInformation) GetProvinceStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvinceState.Get(), o.ProvinceState.IsSet()
}

// HasProvinceState returns a boolean if a field has been set.
func (o *BillingInformation) HasProvinceState() bool {
	if o != nil && o.ProvinceState.IsSet() {
		return true
	}

	return false
}

// SetProvinceState gets a reference to the given NullableString and assigns it to the ProvinceState field.
func (o *BillingInformation) SetProvinceState(v string) {
	o.ProvinceState.Set(&v)
}
// SetProvinceStateNil sets the value for ProvinceState to be an explicit nil
func (o *BillingInformation) SetProvinceStateNil() {
	o.ProvinceState.Set(nil)
}

// UnsetProvinceState ensures that no value is present for ProvinceState, not even an explicit nil
func (o *BillingInformation) UnsetProvinceState() {
	o.ProvinceState.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInformation) GetPostalCode() string {
	if o == nil || utils.IsNil(o.PostalCode.Get()) {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInformation) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *BillingInformation) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *BillingInformation) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *BillingInformation) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *BillingInformation) UnsetPostalCode() {
	o.PostalCode.Unset()
}

func (o BillingInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country"] = o.Country
	if o.StreetLine1.IsSet() {
		toSerialize["street_line1"] = o.StreetLine1.Get()
	}
	if o.StreetLine2.IsSet() {
		toSerialize["street_line2"] = o.StreetLine2.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.ProvinceState.IsSet() {
		toSerialize["province_state"] = o.ProvinceState.Get()
	}
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	return toSerialize, nil
}

type NullableBillingInformation struct {
	value *BillingInformation
	isSet bool
}

func (v NullableBillingInformation) Get() *BillingInformation {
	return v.value
}

func (v *NullableBillingInformation) Set(val *BillingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInformation(val *BillingInformation) *NullableBillingInformation {
	return &NullableBillingInformation{value: val, isSet: true}
}

func (v NullableBillingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



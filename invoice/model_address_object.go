/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.7.6
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the AddressObject type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressObject{}

// AddressObject An object representing an address with various properties.
type AddressObject struct {
	// The country where the address is located.
	Country NullableString `json:"country,omitempty"`
	// The first line of the street address.
	StreetLine1 NullableString `json:"street_line1,omitempty"`
	// The second line of the street address.
	StreetLine2 NullableString `json:"street_line2,omitempty"`
	// The city or locality within the address.
	City NullableString `json:"city,omitempty"`
	// The province or region within the country.
	Province NullableString `json:"province,omitempty"`
	// The state or administrative division within the country.
	State NullableString `json:"state,omitempty"`
	// The postal code or ZIP code for the address.
	PostalCode NullableString `json:"postal_code,omitempty"`
}

// NewAddressObject instantiates a new AddressObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressObject() *AddressObject {
	this := AddressObject{}
	return &this
}

// NewAddressObjectWithDefaults instantiates a new AddressObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressObjectWithDefaults() *AddressObject {
	this := AddressObject{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressObject) GetCountry() string {
	if o == nil || utils.IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressObject) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *AddressObject) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *AddressObject) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *AddressObject) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *AddressObject) UnsetCountry() {
	o.Country.Unset()
}

// GetStreetLine1 returns the StreetLine1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressObject) GetStreetLine1() string {
	if o == nil || utils.IsNil(o.StreetLine1.Get()) {
		var ret string
		return ret
	}
	return *o.StreetLine1.Get()
}

// GetStreetLine1Ok returns a tuple with the StreetLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressObject) GetStreetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreetLine1.Get(), o.StreetLine1.IsSet()
}

// HasStreetLine1 returns a boolean if a field has been set.
func (o *AddressObject) HasStreetLine1() bool {
	if o != nil && o.StreetLine1.IsSet() {
		return true
	}

	return false
}

// SetStreetLine1 gets a reference to the given NullableString and assigns it to the StreetLine1 field.
func (o *AddressObject) SetStreetLine1(v string) {
	o.StreetLine1.Set(&v)
}
// SetStreetLine1Nil sets the value for StreetLine1 to be an explicit nil
func (o *AddressObject) SetStreetLine1Nil() {
	o.StreetLine1.Set(nil)
}

// UnsetStreetLine1 ensures that no value is present for StreetLine1, not even an explicit nil
func (o *AddressObject) UnsetStreetLine1() {
	o.StreetLine1.Unset()
}

// GetStreetLine2 returns the StreetLine2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressObject) GetStreetLine2() string {
	if o == nil || utils.IsNil(o.StreetLine2.Get()) {
		var ret string
		return ret
	}
	return *o.StreetLine2.Get()
}

// GetStreetLine2Ok returns a tuple with the StreetLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressObject) GetStreetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreetLine2.Get(), o.StreetLine2.IsSet()
}

// HasStreetLine2 returns a boolean if a field has been set.
func (o *AddressObject) HasStreetLine2() bool {
	if o != nil && o.StreetLine2.IsSet() {
		return true
	}

	return false
}

// SetStreetLine2 gets a reference to the given NullableString and assigns it to the StreetLine2 field.
func (o *AddressObject) SetStreetLine2(v string) {
	o.StreetLine2.Set(&v)
}
// SetStreetLine2Nil sets the value for StreetLine2 to be an explicit nil
func (o *AddressObject) SetStreetLine2Nil() {
	o.StreetLine2.Set(nil)
}

// UnsetStreetLine2 ensures that no value is present for StreetLine2, not even an explicit nil
func (o *AddressObject) UnsetStreetLine2() {
	o.StreetLine2.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressObject) GetCity() string {
	if o == nil || utils.IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressObject) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *AddressObject) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *AddressObject) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *AddressObject) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *AddressObject) UnsetCity() {
	o.City.Unset()
}

// GetProvince returns the Province field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressObject) GetProvince() string {
	if o == nil || utils.IsNil(o.Province.Get()) {
		var ret string
		return ret
	}
	return *o.Province.Get()
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressObject) GetProvinceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Province.Get(), o.Province.IsSet()
}

// HasProvince returns a boolean if a field has been set.
func (o *AddressObject) HasProvince() bool {
	if o != nil && o.Province.IsSet() {
		return true
	}

	return false
}

// SetProvince gets a reference to the given NullableString and assigns it to the Province field.
func (o *AddressObject) SetProvince(v string) {
	o.Province.Set(&v)
}
// SetProvinceNil sets the value for Province to be an explicit nil
func (o *AddressObject) SetProvinceNil() {
	o.Province.Set(nil)
}

// UnsetProvince ensures that no value is present for Province, not even an explicit nil
func (o *AddressObject) UnsetProvince() {
	o.Province.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressObject) GetState() string {
	if o == nil || utils.IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressObject) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *AddressObject) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *AddressObject) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *AddressObject) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *AddressObject) UnsetState() {
	o.State.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressObject) GetPostalCode() string {
	if o == nil || utils.IsNil(o.PostalCode.Get()) {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressObject) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AddressObject) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *AddressObject) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *AddressObject) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *AddressObject) UnsetPostalCode() {
	o.PostalCode.Unset()
}

func (o AddressObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
    }
	if o.StreetLine1.IsSet() {
		toSerialize["street_line1"] = o.StreetLine1.Get()
    }
	if o.StreetLine2.IsSet() {
		toSerialize["street_line2"] = o.StreetLine2.Get()
    }
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
    }
	if o.Province.IsSet() {
		toSerialize["province"] = o.Province.Get()
    }
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
    }
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
    }
	return toSerialize, nil
}

type NullableAddressObject struct {
	value *AddressObject
	isSet bool
}

func (v NullableAddressObject) Get() *AddressObject {
	return v.value
}

func (v *NullableAddressObject) Set(val *AddressObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressObject(val *AddressObject) *NullableAddressObject {
	return &NullableAddressObject{value: val, isSet: true}
}

func (v NullableAddressObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



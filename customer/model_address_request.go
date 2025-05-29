/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the AddressRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AddressRequest{}

// AddressRequest struct for AddressRequest
type AddressRequest struct {
	// Home, work or provincial
	Category *string `json:"category,omitempty"`
	// ISO3166-2 country code
	CountryCode NullableString `json:"country_code,omitempty"`
	ProvinceState *string `json:"province_state,omitempty"`
	City *string `json:"city,omitempty"`
	Suburb *string `json:"suburb,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	Line1 *string `json:"line_1,omitempty"`
	Line2 *string `json:"line_2,omitempty"`
	Status NullableAddressStatus `json:"status,omitempty"`
	IsPrimary *bool `json:"is_primary,omitempty"`
}

// NewAddressRequest instantiates a new AddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressRequest() *AddressRequest {
	this := AddressRequest{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// NewAddressRequestWithDefaults instantiates a new AddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressRequestWithDefaults() *AddressRequest {
	this := AddressRequest{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AddressRequest) GetCategory() string {
	if o == nil || utils.IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetCategoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AddressRequest) HasCategory() bool {
	if o != nil && !utils.IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AddressRequest) SetCategory(v string) {
	o.Category = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressRequest) GetCountryCode() string {
	if o == nil || utils.IsNil(o.CountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressRequest) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *AddressRequest) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *AddressRequest) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}
// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *AddressRequest) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *AddressRequest) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetProvinceState returns the ProvinceState field value if set, zero value otherwise.
func (o *AddressRequest) GetProvinceState() string {
	if o == nil || utils.IsNil(o.ProvinceState) {
		var ret string
		return ret
	}
	return *o.ProvinceState
}

// GetProvinceStateOk returns a tuple with the ProvinceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetProvinceStateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ProvinceState) {
		return nil, false
	}
	return o.ProvinceState, true
}

// HasProvinceState returns a boolean if a field has been set.
func (o *AddressRequest) HasProvinceState() bool {
	if o != nil && !utils.IsNil(o.ProvinceState) {
		return true
	}

	return false
}

// SetProvinceState gets a reference to the given string and assigns it to the ProvinceState field.
func (o *AddressRequest) SetProvinceState(v string) {
	o.ProvinceState = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AddressRequest) GetCity() string {
	if o == nil || utils.IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetCityOk() (*string, bool) {
	if o == nil || utils.IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AddressRequest) HasCity() bool {
	if o != nil && !utils.IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AddressRequest) SetCity(v string) {
	o.City = &v
}

// GetSuburb returns the Suburb field value if set, zero value otherwise.
func (o *AddressRequest) GetSuburb() string {
	if o == nil || utils.IsNil(o.Suburb) {
		var ret string
		return ret
	}
	return *o.Suburb
}

// GetSuburbOk returns a tuple with the Suburb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetSuburbOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Suburb) {
		return nil, false
	}
	return o.Suburb, true
}

// HasSuburb returns a boolean if a field has been set.
func (o *AddressRequest) HasSuburb() bool {
	if o != nil && !utils.IsNil(o.Suburb) {
		return true
	}

	return false
}

// SetSuburb gets a reference to the given string and assigns it to the Suburb field.
func (o *AddressRequest) SetSuburb(v string) {
	o.Suburb = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AddressRequest) GetPostalCode() string {
	if o == nil || utils.IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetPostalCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AddressRequest) HasPostalCode() bool {
	if o != nil && !utils.IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AddressRequest) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *AddressRequest) GetLine1() string {
	if o == nil || utils.IsNil(o.Line1) {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetLine1Ok() (*string, bool) {
	if o == nil || utils.IsNil(o.Line1) {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *AddressRequest) HasLine1() bool {
	if o != nil && !utils.IsNil(o.Line1) {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *AddressRequest) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *AddressRequest) GetLine2() string {
	if o == nil || utils.IsNil(o.Line2) {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetLine2Ok() (*string, bool) {
	if o == nil || utils.IsNil(o.Line2) {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *AddressRequest) HasLine2() bool {
	if o != nil && !utils.IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *AddressRequest) SetLine2(v string) {
	o.Line2 = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressRequest) GetStatus() AddressStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret AddressStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressRequest) GetStatusOk() (*AddressStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *AddressRequest) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableAddressStatus and assigns it to the Status field.
func (o *AddressRequest) SetStatus(v AddressStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *AddressRequest) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *AddressRequest) UnsetStatus() {
	o.Status.Unset()
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *AddressRequest) GetIsPrimary() bool {
	if o == nil || utils.IsNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsPrimary) {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *AddressRequest) HasIsPrimary() bool {
	if o != nil && !utils.IsNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *AddressRequest) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

func (o AddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if o.CountryCode.IsSet() {
		toSerialize["country_code"] = o.CountryCode.Get()
    }
	if !utils.IsNil(o.ProvinceState) {
		toSerialize["province_state"] = o.ProvinceState
	}
	if !utils.IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !utils.IsNil(o.Suburb) {
		toSerialize["suburb"] = o.Suburb
	}
	if !utils.IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !utils.IsNil(o.Line1) {
		toSerialize["line_1"] = o.Line1
	}
	if !utils.IsNil(o.Line2) {
		toSerialize["line_2"] = o.Line2
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
    }
	if !utils.IsNil(o.IsPrimary) {
		toSerialize["is_primary"] = o.IsPrimary
	}
	return toSerialize, nil
}

type NullableAddressRequest struct {
	value *AddressRequest
	isSet bool
}

func (v NullableAddressRequest) Get() *AddressRequest {
	return v.value
}

func (v *NullableAddressRequest) Set(val *AddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressRequest(val *AddressRequest) *NullableAddressRequest {
	return &NullableAddressRequest{value: val, isSet: true}
}

func (v NullableAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



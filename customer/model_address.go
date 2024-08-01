/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
	"time"
)

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Address{}

// Address struct for Address
type Address struct {
	Id *string `json:"id,omitempty"`
	Category NullableString `json:"category"`
	Country string `json:"country"`
	ProvinceState NullableString `json:"province_state"`
	City NullableString `json:"city"`
	PostalCode NullableString `json:"postal_code"`
	StreetLine1 NullableString `json:"street_line1"`
	StreetLine2 NullableString `json:"street_line2"`
	Status NullableAddressStatus `json:"status,omitempty"`
	IsPrimary NullableBool `json:"is_primary"`
	Meta map[string]interface{} `json:"meta,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress(category NullableString, country string, provinceState NullableString, city NullableString, postalCode NullableString, streetLine1 NullableString, streetLine2 NullableString, isPrimary NullableBool) *Address {
	this := Address{}
	this.Category = category
	this.Country = country
	this.ProvinceState = provinceState
	this.City = city
	this.PostalCode = postalCode
	this.StreetLine1 = streetLine1
	this.StreetLine2 = streetLine2
	this.IsPrimary = isPrimary
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Address) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Address) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Address) SetId(v string) {
	o.Id = &v
}

// GetCategory returns the Category field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Address) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}

	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// SetCategory sets field value
func (o *Address) SetCategory(v string) {
	o.Category.Set(&v)
}

// GetCountry returns the Country field value
func (o *Address) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *Address) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *Address) SetCountry(v string) {
	o.Country = v
}

// GetProvinceState returns the ProvinceState field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Address) GetProvinceState() string {
	if o == nil || o.ProvinceState.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProvinceState.Get()
}

// GetProvinceStateOk returns a tuple with the ProvinceState field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetProvinceStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvinceState.Get(), o.ProvinceState.IsSet()
}

// SetProvinceState sets field value
func (o *Address) SetProvinceState(v string) {
	o.ProvinceState.Set(&v)
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Address) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *Address) SetCity(v string) {
	o.City.Set(&v)
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Address) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *Address) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

// GetStreetLine1 returns the StreetLine1 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Address) GetStreetLine1() string {
	if o == nil || o.StreetLine1.Get() == nil {
		var ret string
		return ret
	}

	return *o.StreetLine1.Get()
}

// GetStreetLine1Ok returns a tuple with the StreetLine1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStreetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreetLine1.Get(), o.StreetLine1.IsSet()
}

// SetStreetLine1 sets field value
func (o *Address) SetStreetLine1(v string) {
	o.StreetLine1.Set(&v)
}

// GetStreetLine2 returns the StreetLine2 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Address) GetStreetLine2() string {
	if o == nil || o.StreetLine2.Get() == nil {
		var ret string
		return ret
	}

	return *o.StreetLine2.Get()
}

// GetStreetLine2Ok returns a tuple with the StreetLine2 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStreetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreetLine2.Get(), o.StreetLine2.IsSet()
}

// SetStreetLine2 sets field value
func (o *Address) SetStreetLine2(v string) {
	o.StreetLine2.Set(&v)
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetStatus() AddressStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret AddressStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStatusOk() (*AddressStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Address) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableAddressStatus and assigns it to the Status field.
func (o *Address) SetStatus(v AddressStatus) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Address) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Address) UnsetStatus() {
	o.Status.Unset()
}

// GetIsPrimary returns the IsPrimary field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *Address) GetIsPrimary() bool {
	if o == nil || o.IsPrimary.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsPrimary.Get()
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetIsPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPrimary.Get(), o.IsPrimary.IsSet()
}

// SetIsPrimary sets field value
func (o *Address) SetIsPrimary(v bool) {
	o.IsPrimary.Set(&v)
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetMetaOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Meta) {
		return map[string]interface{}{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Address) HasMeta() bool {
	if o != nil && utils.IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *Address) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Address) GetCreated() time.Time {
	if o == nil || utils.IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCreatedOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Address) HasCreated() bool {
	if o != nil && !utils.IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Address) SetCreated(v time.Time) {
	o.Created = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Address) GetUpdated() time.Time {
	if o == nil || utils.IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Address) HasUpdated() bool {
	if o != nil && !utils.IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Address) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["category"] = o.Category.Get()

	toSerialize["country"] = o.Country
	toSerialize["province_state"] = o.ProvinceState.Get()

	toSerialize["city"] = o.City.Get()

	toSerialize["postal_code"] = o.PostalCode.Get()

	toSerialize["street_line1"] = o.StreetLine1.Get()

	toSerialize["street_line2"] = o.StreetLine2.Get()

	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
    }
	toSerialize["is_primary"] = o.IsPrimary.Get()

	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
    }
	if !utils.IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !utils.IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	return toSerialize, nil
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



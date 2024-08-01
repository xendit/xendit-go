/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the BusinessDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BusinessDetail{}

// BusinessDetail struct for BusinessDetail
type BusinessDetail struct {
	BusinessName *string `json:"business_name,omitempty"`
	BusinessType NullableString `json:"business_type,omitempty"`
	DateOfRegistration NullableString `json:"date_of_registration,omitempty"`
	NatureOfBusiness NullableString `json:"nature_of_business,omitempty"`
	// ISO3166-2 country code
	BusinessDomicile NullableString `json:"business_domicile,omitempty"`
}

// NewBusinessDetail instantiates a new BusinessDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessDetail() *BusinessDetail {
	this := BusinessDetail{}
	return &this
}

// NewBusinessDetailWithDefaults instantiates a new BusinessDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessDetailWithDefaults() *BusinessDetail {
	this := BusinessDetail{}
	return &this
}

// GetBusinessName returns the BusinessName field value if set, zero value otherwise.
func (o *BusinessDetail) GetBusinessName() string {
	if o == nil || utils.IsNil(o.BusinessName) {
		var ret string
		return ret
	}
	return *o.BusinessName
}

// GetBusinessNameOk returns a tuple with the BusinessName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessDetail) GetBusinessNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BusinessName) {
		return nil, false
	}
	return o.BusinessName, true
}

// HasBusinessName returns a boolean if a field has been set.
func (o *BusinessDetail) HasBusinessName() bool {
	if o != nil && !utils.IsNil(o.BusinessName) {
		return true
	}

	return false
}

// SetBusinessName gets a reference to the given string and assigns it to the BusinessName field.
func (o *BusinessDetail) SetBusinessName(v string) {
	o.BusinessName = &v
}

// GetBusinessType returns the BusinessType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessDetail) GetBusinessType() string {
	if o == nil || utils.IsNil(o.BusinessType.Get()) {
		var ret string
		return ret
	}
	return *o.BusinessType.Get()
}

// GetBusinessTypeOk returns a tuple with the BusinessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessDetail) GetBusinessTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessType.Get(), o.BusinessType.IsSet()
}

// HasBusinessType returns a boolean if a field has been set.
func (o *BusinessDetail) HasBusinessType() bool {
	if o != nil && o.BusinessType.IsSet() {
		return true
	}

	return false
}

// SetBusinessType gets a reference to the given NullableString and assigns it to the BusinessType field.
func (o *BusinessDetail) SetBusinessType(v string) {
	o.BusinessType.Set(&v)
}
// SetBusinessTypeNil sets the value for BusinessType to be an explicit nil
func (o *BusinessDetail) SetBusinessTypeNil() {
	o.BusinessType.Set(nil)
}

// UnsetBusinessType ensures that no value is present for BusinessType, not even an explicit nil
func (o *BusinessDetail) UnsetBusinessType() {
	o.BusinessType.Unset()
}

// GetDateOfRegistration returns the DateOfRegistration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessDetail) GetDateOfRegistration() string {
	if o == nil || utils.IsNil(o.DateOfRegistration.Get()) {
		var ret string
		return ret
	}
	return *o.DateOfRegistration.Get()
}

// GetDateOfRegistrationOk returns a tuple with the DateOfRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessDetail) GetDateOfRegistrationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateOfRegistration.Get(), o.DateOfRegistration.IsSet()
}

// HasDateOfRegistration returns a boolean if a field has been set.
func (o *BusinessDetail) HasDateOfRegistration() bool {
	if o != nil && o.DateOfRegistration.IsSet() {
		return true
	}

	return false
}

// SetDateOfRegistration gets a reference to the given NullableString and assigns it to the DateOfRegistration field.
func (o *BusinessDetail) SetDateOfRegistration(v string) {
	o.DateOfRegistration.Set(&v)
}
// SetDateOfRegistrationNil sets the value for DateOfRegistration to be an explicit nil
func (o *BusinessDetail) SetDateOfRegistrationNil() {
	o.DateOfRegistration.Set(nil)
}

// UnsetDateOfRegistration ensures that no value is present for DateOfRegistration, not even an explicit nil
func (o *BusinessDetail) UnsetDateOfRegistration() {
	o.DateOfRegistration.Unset()
}

// GetNatureOfBusiness returns the NatureOfBusiness field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessDetail) GetNatureOfBusiness() string {
	if o == nil || utils.IsNil(o.NatureOfBusiness.Get()) {
		var ret string
		return ret
	}
	return *o.NatureOfBusiness.Get()
}

// GetNatureOfBusinessOk returns a tuple with the NatureOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessDetail) GetNatureOfBusinessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NatureOfBusiness.Get(), o.NatureOfBusiness.IsSet()
}

// HasNatureOfBusiness returns a boolean if a field has been set.
func (o *BusinessDetail) HasNatureOfBusiness() bool {
	if o != nil && o.NatureOfBusiness.IsSet() {
		return true
	}

	return false
}

// SetNatureOfBusiness gets a reference to the given NullableString and assigns it to the NatureOfBusiness field.
func (o *BusinessDetail) SetNatureOfBusiness(v string) {
	o.NatureOfBusiness.Set(&v)
}
// SetNatureOfBusinessNil sets the value for NatureOfBusiness to be an explicit nil
func (o *BusinessDetail) SetNatureOfBusinessNil() {
	o.NatureOfBusiness.Set(nil)
}

// UnsetNatureOfBusiness ensures that no value is present for NatureOfBusiness, not even an explicit nil
func (o *BusinessDetail) UnsetNatureOfBusiness() {
	o.NatureOfBusiness.Unset()
}

// GetBusinessDomicile returns the BusinessDomicile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BusinessDetail) GetBusinessDomicile() string {
	if o == nil || utils.IsNil(o.BusinessDomicile.Get()) {
		var ret string
		return ret
	}
	return *o.BusinessDomicile.Get()
}

// GetBusinessDomicileOk returns a tuple with the BusinessDomicile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BusinessDetail) GetBusinessDomicileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BusinessDomicile.Get(), o.BusinessDomicile.IsSet()
}

// HasBusinessDomicile returns a boolean if a field has been set.
func (o *BusinessDetail) HasBusinessDomicile() bool {
	if o != nil && o.BusinessDomicile.IsSet() {
		return true
	}

	return false
}

// SetBusinessDomicile gets a reference to the given NullableString and assigns it to the BusinessDomicile field.
func (o *BusinessDetail) SetBusinessDomicile(v string) {
	o.BusinessDomicile.Set(&v)
}
// SetBusinessDomicileNil sets the value for BusinessDomicile to be an explicit nil
func (o *BusinessDetail) SetBusinessDomicileNil() {
	o.BusinessDomicile.Set(nil)
}

// UnsetBusinessDomicile ensures that no value is present for BusinessDomicile, not even an explicit nil
func (o *BusinessDetail) UnsetBusinessDomicile() {
	o.BusinessDomicile.Unset()
}

func (o BusinessDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.BusinessName) {
		toSerialize["business_name"] = o.BusinessName
	}
	if o.BusinessType.IsSet() {
		toSerialize["business_type"] = o.BusinessType.Get()
        if o.BusinessType.Get() != nil && (*o.BusinessType.Get() != "CORPORATION" && *o.BusinessType.Get() != "SOLE_PROPRIETOR" && *o.BusinessType.Get() != "PARTNERSHIP" && *o.BusinessType.Get() != "COOPERATIVE" && *o.BusinessType.Get() != "TRUST" && *o.BusinessType.Get() != "NON_PROFIT" && *o.BusinessType.Get() != "GOVERNMENT") {
            toSerialize["business_type"] = nil
            return toSerialize, utils.NewError("invalid value for BusinessType when marshalling to JSON, must be one of CORPORATION, SOLE_PROPRIETOR, PARTNERSHIP, COOPERATIVE, TRUST, NON_PROFIT, GOVERNMENT")
        }
    }
	if o.DateOfRegistration.IsSet() {
		toSerialize["date_of_registration"] = o.DateOfRegistration.Get()
    }
	if o.NatureOfBusiness.IsSet() {
		toSerialize["nature_of_business"] = o.NatureOfBusiness.Get()
    }
	if o.BusinessDomicile.IsSet() {
		toSerialize["business_domicile"] = o.BusinessDomicile.Get()
    }
	return toSerialize, nil
}

type NullableBusinessDetail struct {
	value *BusinessDetail
	isSet bool
}

func (v NullableBusinessDetail) Get() *BusinessDetail {
	return v.value
}

func (v *NullableBusinessDetail) Set(val *BusinessDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessDetail(val *BusinessDetail) *NullableBusinessDetail {
	return &NullableBusinessDetail{value: val, isSet: true}
}

func (v NullableBusinessDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



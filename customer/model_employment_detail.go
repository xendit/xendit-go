/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the EmploymentDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EmploymentDetail{}

// EmploymentDetail struct for EmploymentDetail
type EmploymentDetail struct {
	// Name of employer
	EmployerName NullableString `json:"employer_name,omitempty"`
	// Industry or nature of business
	NatureOfBusiness NullableString `json:"nature_of_business,omitempty"`
	// Occupation or title
	RoleDescription NullableString `json:"role_description,omitempty"`
}

// NewEmploymentDetail instantiates a new EmploymentDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmploymentDetail() *EmploymentDetail {
	this := EmploymentDetail{}
	return &this
}

// NewEmploymentDetailWithDefaults instantiates a new EmploymentDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmploymentDetailWithDefaults() *EmploymentDetail {
	this := EmploymentDetail{}
	return &this
}

// GetEmployerName returns the EmployerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmploymentDetail) GetEmployerName() string {
	if o == nil || utils.IsNil(o.EmployerName.Get()) {
		var ret string
		return ret
	}
	return *o.EmployerName.Get()
}

// GetEmployerNameOk returns a tuple with the EmployerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentDetail) GetEmployerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmployerName.Get(), o.EmployerName.IsSet()
}

// HasEmployerName returns a boolean if a field has been set.
func (o *EmploymentDetail) HasEmployerName() bool {
	if o != nil && o.EmployerName.IsSet() {
		return true
	}

	return false
}

// SetEmployerName gets a reference to the given NullableString and assigns it to the EmployerName field.
func (o *EmploymentDetail) SetEmployerName(v string) {
	o.EmployerName.Set(&v)
}
// SetEmployerNameNil sets the value for EmployerName to be an explicit nil
func (o *EmploymentDetail) SetEmployerNameNil() {
	o.EmployerName.Set(nil)
}

// UnsetEmployerName ensures that no value is present for EmployerName, not even an explicit nil
func (o *EmploymentDetail) UnsetEmployerName() {
	o.EmployerName.Unset()
}

// GetNatureOfBusiness returns the NatureOfBusiness field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmploymentDetail) GetNatureOfBusiness() string {
	if o == nil || utils.IsNil(o.NatureOfBusiness.Get()) {
		var ret string
		return ret
	}
	return *o.NatureOfBusiness.Get()
}

// GetNatureOfBusinessOk returns a tuple with the NatureOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentDetail) GetNatureOfBusinessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NatureOfBusiness.Get(), o.NatureOfBusiness.IsSet()
}

// HasNatureOfBusiness returns a boolean if a field has been set.
func (o *EmploymentDetail) HasNatureOfBusiness() bool {
	if o != nil && o.NatureOfBusiness.IsSet() {
		return true
	}

	return false
}

// SetNatureOfBusiness gets a reference to the given NullableString and assigns it to the NatureOfBusiness field.
func (o *EmploymentDetail) SetNatureOfBusiness(v string) {
	o.NatureOfBusiness.Set(&v)
}
// SetNatureOfBusinessNil sets the value for NatureOfBusiness to be an explicit nil
func (o *EmploymentDetail) SetNatureOfBusinessNil() {
	o.NatureOfBusiness.Set(nil)
}

// UnsetNatureOfBusiness ensures that no value is present for NatureOfBusiness, not even an explicit nil
func (o *EmploymentDetail) UnsetNatureOfBusiness() {
	o.NatureOfBusiness.Unset()
}

// GetRoleDescription returns the RoleDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmploymentDetail) GetRoleDescription() string {
	if o == nil || utils.IsNil(o.RoleDescription.Get()) {
		var ret string
		return ret
	}
	return *o.RoleDescription.Get()
}

// GetRoleDescriptionOk returns a tuple with the RoleDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentDetail) GetRoleDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleDescription.Get(), o.RoleDescription.IsSet()
}

// HasRoleDescription returns a boolean if a field has been set.
func (o *EmploymentDetail) HasRoleDescription() bool {
	if o != nil && o.RoleDescription.IsSet() {
		return true
	}

	return false
}

// SetRoleDescription gets a reference to the given NullableString and assigns it to the RoleDescription field.
func (o *EmploymentDetail) SetRoleDescription(v string) {
	o.RoleDescription.Set(&v)
}
// SetRoleDescriptionNil sets the value for RoleDescription to be an explicit nil
func (o *EmploymentDetail) SetRoleDescriptionNil() {
	o.RoleDescription.Set(nil)
}

// UnsetRoleDescription ensures that no value is present for RoleDescription, not even an explicit nil
func (o *EmploymentDetail) UnsetRoleDescription() {
	o.RoleDescription.Unset()
}

func (o EmploymentDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmploymentDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EmployerName.IsSet() {
		toSerialize["employer_name"] = o.EmployerName.Get()
    }
	if o.NatureOfBusiness.IsSet() {
		toSerialize["nature_of_business"] = o.NatureOfBusiness.Get()
    }
	if o.RoleDescription.IsSet() {
		toSerialize["role_description"] = o.RoleDescription.Get()
    }
	return toSerialize, nil
}

type NullableEmploymentDetail struct {
	value *EmploymentDetail
	isSet bool
}

func (v NullableEmploymentDetail) Get() *EmploymentDetail {
	return v.value
}

func (v *NullableEmploymentDetail) Set(val *EmploymentDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableEmploymentDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableEmploymentDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmploymentDetail(val *EmploymentDetail) *NullableEmploymentDetail {
	return &NullableEmploymentDetail{value: val, isSet: true}
}

func (v NullableEmploymentDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmploymentDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the IdentityAccountRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IdentityAccountRequest{}

// IdentityAccountRequest struct for IdentityAccountRequest
type IdentityAccountRequest struct {
	Type *IdentityAccountType `json:"type,omitempty"`
	// The issuing institution associated with the account (e.g., OCBC, GOPAY, 7-11). If adding financial accounts that Xendit supports, we recommend you use the channel_name found at https://xendit.github.io/apireference/#payment-channels for this field
	Company *string `json:"company,omitempty"`
	// Free text description of this account
	Description *string `json:"description,omitempty"`
	// ISO3166-2 country code
	Country NullableString `json:"country,omitempty"`
	Properties *IdentityAccountRequestProperties `json:"properties,omitempty"`
}

// NewIdentityAccountRequest instantiates a new IdentityAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAccountRequest() *IdentityAccountRequest {
	this := IdentityAccountRequest{}
	return &this
}

// NewIdentityAccountRequestWithDefaults instantiates a new IdentityAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAccountRequestWithDefaults() *IdentityAccountRequest {
	this := IdentityAccountRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityAccountRequest) GetType() IdentityAccountType {
	if o == nil || utils.IsNil(o.Type) {
		var ret IdentityAccountType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAccountRequest) GetTypeOk() (*IdentityAccountType, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityAccountRequest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given IdentityAccountType and assigns it to the Type field.
func (o *IdentityAccountRequest) SetType(v IdentityAccountType) {
	o.Type = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *IdentityAccountRequest) GetCompany() string {
	if o == nil || utils.IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAccountRequest) GetCompanyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *IdentityAccountRequest) HasCompany() bool {
	if o != nil && !utils.IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *IdentityAccountRequest) SetCompany(v string) {
	o.Company = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentityAccountRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityAccountRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentityAccountRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityAccountRequest) GetCountry() string {
	if o == nil || utils.IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityAccountRequest) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *IdentityAccountRequest) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *IdentityAccountRequest) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *IdentityAccountRequest) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *IdentityAccountRequest) UnsetCountry() {
	o.Country.Unset()
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *IdentityAccountRequest) GetProperties() IdentityAccountRequestProperties {
	if o == nil || utils.IsNil(o.Properties) {
		var ret IdentityAccountRequestProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAccountRequest) GetPropertiesOk() (*IdentityAccountRequestProperties, bool) {
	if o == nil || utils.IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *IdentityAccountRequest) HasProperties() bool {
	if o != nil && !utils.IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given IdentityAccountRequestProperties and assigns it to the Properties field.
func (o *IdentityAccountRequest) SetProperties(v IdentityAccountRequestProperties) {
	o.Properties = &v
}

func (o IdentityAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
    }
	if !utils.IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableIdentityAccountRequest struct {
	value *IdentityAccountRequest
	isSet bool
}

func (v NullableIdentityAccountRequest) Get() *IdentityAccountRequest {
	return v.value
}

func (v *NullableIdentityAccountRequest) Set(val *IdentityAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAccountRequest(val *IdentityAccountRequest) *NullableIdentityAccountRequest {
	return &NullableIdentityAccountRequest{value: val, isSet: true}
}

func (v NullableIdentityAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



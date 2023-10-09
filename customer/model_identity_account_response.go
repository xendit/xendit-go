/*
XENDIT SDK openapi spec

XENDIT SDK openapi spec

API version: 1.0.0
*/


package customer

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
	"time"
)

// checks if the IdentityAccountResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IdentityAccountResponse{}

// IdentityAccountResponse struct for IdentityAccountResponse
type IdentityAccountResponse struct {
	Id *string `json:"id,omitempty"`
	Code NullableString `json:"code,omitempty"`
	Company NullableString `json:"company"`
	Description NullableString `json:"description"`
	// ISO3166-2 country code
	Country NullableString `json:"country"`
	HolderName NullableString `json:"holder_name,omitempty"`
	Type NullableString `json:"type"`
	Properties IdentityAccountResponseProperties `json:"properties"`
	Created *time.Time `json:"created,omitempty"`
}

// NewIdentityAccountResponse instantiates a new IdentityAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAccountResponse(company NullableString, description NullableString, country NullableString, type_ NullableString, properties IdentityAccountResponseProperties) *IdentityAccountResponse {
	this := IdentityAccountResponse{}
	this.Company = company
	this.Description = description
	this.Country = country
	this.Type = type_
	this.Properties = properties
	return &this
}

// NewIdentityAccountResponseWithDefaults instantiates a new IdentityAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAccountResponseWithDefaults() *IdentityAccountResponse {
	this := IdentityAccountResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityAccountResponse) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAccountResponse) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityAccountResponse) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityAccountResponse) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityAccountResponse) GetCode() string {
	if o == nil || utils.IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityAccountResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *IdentityAccountResponse) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *IdentityAccountResponse) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *IdentityAccountResponse) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *IdentityAccountResponse) UnsetCode() {
	o.Code.Unset()
}

// GetCompany returns the Company field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityAccountResponse) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}

	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityAccountResponse) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// SetCompany sets field value
func (o *IdentityAccountResponse) SetCompany(v string) {
	o.Company.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityAccountResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityAccountResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *IdentityAccountResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityAccountResponse) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityAccountResponse) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *IdentityAccountResponse) SetCountry(v string) {
	o.Country.Set(&v)
}

// GetHolderName returns the HolderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityAccountResponse) GetHolderName() string {
	if o == nil || utils.IsNil(o.HolderName.Get()) {
		var ret string
		return ret
	}
	return *o.HolderName.Get()
}

// GetHolderNameOk returns a tuple with the HolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityAccountResponse) GetHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HolderName.Get(), o.HolderName.IsSet()
}

// HasHolderName returns a boolean if a field has been set.
func (o *IdentityAccountResponse) HasHolderName() bool {
	if o != nil && o.HolderName.IsSet() {
		return true
	}

	return false
}

// SetHolderName gets a reference to the given NullableString and assigns it to the HolderName field.
func (o *IdentityAccountResponse) SetHolderName(v string) {
	o.HolderName.Set(&v)
}
// SetHolderNameNil sets the value for HolderName to be an explicit nil
func (o *IdentityAccountResponse) SetHolderNameNil() {
	o.HolderName.Set(nil)
}

// UnsetHolderName ensures that no value is present for HolderName, not even an explicit nil
func (o *IdentityAccountResponse) UnsetHolderName() {
	o.HolderName.Unset()
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IdentityAccountResponse) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityAccountResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *IdentityAccountResponse) SetType(v string) {
	o.Type.Set(&v)
}

// GetProperties returns the Properties field value
func (o *IdentityAccountResponse) GetProperties() IdentityAccountResponseProperties {
	if o == nil {
		var ret IdentityAccountResponseProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *IdentityAccountResponse) GetPropertiesOk() (*IdentityAccountResponseProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *IdentityAccountResponse) SetProperties(v IdentityAccountResponseProperties) {
	o.Properties = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IdentityAccountResponse) GetCreated() time.Time {
	if o == nil || utils.IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAccountResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IdentityAccountResponse) HasCreated() bool {
	if o != nil && !utils.IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IdentityAccountResponse) SetCreated(v time.Time) {
	o.Created = &v
}

func (o IdentityAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
    }
	toSerialize["company"] = o.Company.Get()

	toSerialize["description"] = o.Description.Get()

	toSerialize["country"] = o.Country.Get()

	if o.HolderName.IsSet() {
		toSerialize["holder_name"] = o.HolderName.Get()
    }
	toSerialize["type"] = o.Type.Get()
    if o.Type.Get() != nil && (*o.Type.Get() != "BANK_ACCOUNT" && *o.Type.Get() != "EWALLET" && *o.Type.Get() != "CREDIT_CARD" && *o.Type.Get() != "OTC" && *o.Type.Get() != "QR_CODE" && *o.Type.Get() != "CARDLESS_CREDIT") {
        toSerialize["type"] = nil
        return toSerialize, utils.NewError("invalid value for Type when marshalling to JSON, must be one of BANK_ACCOUNT, EWALLET, CREDIT_CARD, OTC, QR_CODE, CARDLESS_CREDIT")
    }

	toSerialize["properties"] = o.Properties
	if !utils.IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	return toSerialize, nil
}

type NullableIdentityAccountResponse struct {
	value *IdentityAccountResponse
	isSet bool
}

func (v NullableIdentityAccountResponse) Get() *IdentityAccountResponse {
	return v.value
}

func (v *NullableIdentityAccountResponse) Set(val *IdentityAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAccountResponse(val *IdentityAccountResponse) *NullableIdentityAccountResponse {
	return &NullableIdentityAccountResponse{value: val, isSet: true}
}

func (v NullableIdentityAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the KYCDocumentRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &KYCDocumentRequest{}

// KYCDocumentRequest struct for KYCDocumentRequest
type KYCDocumentRequest struct {
	// ISO3166-2 country code
	Country NullableString `json:"country,omitempty"`
	Type *KYCDocumentType `json:"type,omitempty"`
	SubType NullableKYCDocumentSubType `json:"sub_type,omitempty"`
	DocumentName *string `json:"document_name,omitempty"`
	DocumentNumber *string `json:"document_number,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	HolderName *string `json:"holder_name,omitempty"`
	DocumentImages []string `json:"document_images,omitempty"`
}

// NewKYCDocumentRequest instantiates a new KYCDocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKYCDocumentRequest() *KYCDocumentRequest {
	this := KYCDocumentRequest{}
	return &this
}

// NewKYCDocumentRequestWithDefaults instantiates a new KYCDocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKYCDocumentRequestWithDefaults() *KYCDocumentRequest {
	this := KYCDocumentRequest{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KYCDocumentRequest) GetCountry() string {
	if o == nil || utils.IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KYCDocumentRequest) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *KYCDocumentRequest) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *KYCDocumentRequest) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *KYCDocumentRequest) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *KYCDocumentRequest) UnsetCountry() {
	o.Country.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KYCDocumentRequest) GetType() KYCDocumentType {
	if o == nil || utils.IsNil(o.Type) {
		var ret KYCDocumentType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCDocumentRequest) GetTypeOk() (*KYCDocumentType, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KYCDocumentRequest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given KYCDocumentType and assigns it to the Type field.
func (o *KYCDocumentRequest) SetType(v KYCDocumentType) {
	o.Type = &v
}

// GetSubType returns the SubType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KYCDocumentRequest) GetSubType() KYCDocumentSubType {
	if o == nil || utils.IsNil(o.SubType.Get()) {
		var ret KYCDocumentSubType
		return ret
	}
	return *o.SubType.Get()
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KYCDocumentRequest) GetSubTypeOk() (*KYCDocumentSubType, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubType.Get(), o.SubType.IsSet()
}

// HasSubType returns a boolean if a field has been set.
func (o *KYCDocumentRequest) HasSubType() bool {
	if o != nil && o.SubType.IsSet() {
		return true
	}

	return false
}

// SetSubType gets a reference to the given NullableKYCDocumentSubType and assigns it to the SubType field.
func (o *KYCDocumentRequest) SetSubType(v KYCDocumentSubType) {
	o.SubType.Set(&v)
}
// SetSubTypeNil sets the value for SubType to be an explicit nil
func (o *KYCDocumentRequest) SetSubTypeNil() {
	o.SubType.Set(nil)
}

// UnsetSubType ensures that no value is present for SubType, not even an explicit nil
func (o *KYCDocumentRequest) UnsetSubType() {
	o.SubType.Unset()
}

// GetDocumentName returns the DocumentName field value if set, zero value otherwise.
func (o *KYCDocumentRequest) GetDocumentName() string {
	if o == nil || utils.IsNil(o.DocumentName) {
		var ret string
		return ret
	}
	return *o.DocumentName
}

// GetDocumentNameOk returns a tuple with the DocumentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCDocumentRequest) GetDocumentNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DocumentName) {
		return nil, false
	}
	return o.DocumentName, true
}

// HasDocumentName returns a boolean if a field has been set.
func (o *KYCDocumentRequest) HasDocumentName() bool {
	if o != nil && !utils.IsNil(o.DocumentName) {
		return true
	}

	return false
}

// SetDocumentName gets a reference to the given string and assigns it to the DocumentName field.
func (o *KYCDocumentRequest) SetDocumentName(v string) {
	o.DocumentName = &v
}

// GetDocumentNumber returns the DocumentNumber field value if set, zero value otherwise.
func (o *KYCDocumentRequest) GetDocumentNumber() string {
	if o == nil || utils.IsNil(o.DocumentNumber) {
		var ret string
		return ret
	}
	return *o.DocumentNumber
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCDocumentRequest) GetDocumentNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DocumentNumber) {
		return nil, false
	}
	return o.DocumentNumber, true
}

// HasDocumentNumber returns a boolean if a field has been set.
func (o *KYCDocumentRequest) HasDocumentNumber() bool {
	if o != nil && !utils.IsNil(o.DocumentNumber) {
		return true
	}

	return false
}

// SetDocumentNumber gets a reference to the given string and assigns it to the DocumentNumber field.
func (o *KYCDocumentRequest) SetDocumentNumber(v string) {
	o.DocumentNumber = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *KYCDocumentRequest) GetExpiresAt() string {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCDocumentRequest) GetExpiresAtOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *KYCDocumentRequest) HasExpiresAt() bool {
	if o != nil && !utils.IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *KYCDocumentRequest) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetHolderName returns the HolderName field value if set, zero value otherwise.
func (o *KYCDocumentRequest) GetHolderName() string {
	if o == nil || utils.IsNil(o.HolderName) {
		var ret string
		return ret
	}
	return *o.HolderName
}

// GetHolderNameOk returns a tuple with the HolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCDocumentRequest) GetHolderNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.HolderName) {
		return nil, false
	}
	return o.HolderName, true
}

// HasHolderName returns a boolean if a field has been set.
func (o *KYCDocumentRequest) HasHolderName() bool {
	if o != nil && !utils.IsNil(o.HolderName) {
		return true
	}

	return false
}

// SetHolderName gets a reference to the given string and assigns it to the HolderName field.
func (o *KYCDocumentRequest) SetHolderName(v string) {
	o.HolderName = &v
}

// GetDocumentImages returns the DocumentImages field value if set, zero value otherwise.
func (o *KYCDocumentRequest) GetDocumentImages() []string {
	if o == nil || utils.IsNil(o.DocumentImages) {
		var ret []string
		return ret
	}
	return o.DocumentImages
}

// GetDocumentImagesOk returns a tuple with the DocumentImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KYCDocumentRequest) GetDocumentImagesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.DocumentImages) {
		return nil, false
	}
	return o.DocumentImages, true
}

// HasDocumentImages returns a boolean if a field has been set.
func (o *KYCDocumentRequest) HasDocumentImages() bool {
	if o != nil && !utils.IsNil(o.DocumentImages) {
		return true
	}

	return false
}

// SetDocumentImages gets a reference to the given []string and assigns it to the DocumentImages field.
func (o *KYCDocumentRequest) SetDocumentImages(v []string) {
	o.DocumentImages = v
}

func (o KYCDocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KYCDocumentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
    }
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.SubType.IsSet() {
		toSerialize["sub_type"] = o.SubType.Get()
    }
	if !utils.IsNil(o.DocumentName) {
		toSerialize["document_name"] = o.DocumentName
	}
	if !utils.IsNil(o.DocumentNumber) {
		toSerialize["document_number"] = o.DocumentNumber
	}
	if !utils.IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !utils.IsNil(o.HolderName) {
		toSerialize["holder_name"] = o.HolderName
	}
	if !utils.IsNil(o.DocumentImages) {
		toSerialize["document_images"] = o.DocumentImages
	}
	return toSerialize, nil
}

type NullableKYCDocumentRequest struct {
	value *KYCDocumentRequest
	isSet bool
}

func (v NullableKYCDocumentRequest) Get() *KYCDocumentRequest {
	return v.value
}

func (v *NullableKYCDocumentRequest) Set(val *KYCDocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCDocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCDocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCDocumentRequest(val *KYCDocumentRequest) *NullableKYCDocumentRequest {
	return &NullableKYCDocumentRequest{value: val, isSet: true}
}

func (v NullableKYCDocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCDocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



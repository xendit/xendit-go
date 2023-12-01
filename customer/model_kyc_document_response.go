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

// checks if the KYCDocumentResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &KYCDocumentResponse{}

// KYCDocumentResponse struct for KYCDocumentResponse
type KYCDocumentResponse struct {
	Country string `json:"country"`
	Type KYCDocumentType `json:"type"`
	SubType KYCDocumentSubType `json:"sub_type"`
	DocumentName NullableString `json:"document_name"`
	DocumentNumber NullableString `json:"document_number"`
	ExpiresAt NullableString `json:"expires_at"`
	HolderName NullableString `json:"holder_name"`
	DocumentImages []string `json:"document_images"`
}

// NewKYCDocumentResponse instantiates a new KYCDocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKYCDocumentResponse(country string, type_ KYCDocumentType, subType KYCDocumentSubType, documentName NullableString, documentNumber NullableString, expiresAt NullableString, holderName NullableString, documentImages []string) *KYCDocumentResponse {
	this := KYCDocumentResponse{}
	this.Country = country
	this.Type = type_
	this.SubType = subType
	this.DocumentName = documentName
	this.DocumentNumber = documentNumber
	this.ExpiresAt = expiresAt
	this.HolderName = holderName
	this.DocumentImages = documentImages
	return &this
}

// NewKYCDocumentResponseWithDefaults instantiates a new KYCDocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKYCDocumentResponseWithDefaults() *KYCDocumentResponse {
	this := KYCDocumentResponse{}
	return &this
}

// GetCountry returns the Country field value
func (o *KYCDocumentResponse) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *KYCDocumentResponse) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *KYCDocumentResponse) SetCountry(v string) {
	o.Country = v
}

// GetType returns the Type field value
func (o *KYCDocumentResponse) GetType() KYCDocumentType {
	if o == nil {
		var ret KYCDocumentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KYCDocumentResponse) GetTypeOk() (*KYCDocumentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KYCDocumentResponse) SetType(v KYCDocumentType) {
	o.Type = v
}

// GetSubType returns the SubType field value
func (o *KYCDocumentResponse) GetSubType() KYCDocumentSubType {
	if o == nil {
		var ret KYCDocumentSubType
		return ret
	}

	return o.SubType
}

// GetSubTypeOk returns a tuple with the SubType field value
// and a boolean to check if the value has been set.
func (o *KYCDocumentResponse) GetSubTypeOk() (*KYCDocumentSubType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubType, true
}

// SetSubType sets field value
func (o *KYCDocumentResponse) SetSubType(v KYCDocumentSubType) {
	o.SubType = v
}

// GetDocumentName returns the DocumentName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KYCDocumentResponse) GetDocumentName() string {
	if o == nil || o.DocumentName.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentName.Get()
}

// GetDocumentNameOk returns a tuple with the DocumentName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KYCDocumentResponse) GetDocumentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentName.Get(), o.DocumentName.IsSet()
}

// SetDocumentName sets field value
func (o *KYCDocumentResponse) SetDocumentName(v string) {
	o.DocumentName.Set(&v)
}

// GetDocumentNumber returns the DocumentNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KYCDocumentResponse) GetDocumentNumber() string {
	if o == nil || o.DocumentNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentNumber.Get()
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KYCDocumentResponse) GetDocumentNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentNumber.Get(), o.DocumentNumber.IsSet()
}

// SetDocumentNumber sets field value
func (o *KYCDocumentResponse) SetDocumentNumber(v string) {
	o.DocumentNumber.Set(&v)
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KYCDocumentResponse) GetExpiresAt() string {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KYCDocumentResponse) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *KYCDocumentResponse) SetExpiresAt(v string) {
	o.ExpiresAt.Set(&v)
}

// GetHolderName returns the HolderName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *KYCDocumentResponse) GetHolderName() string {
	if o == nil || o.HolderName.Get() == nil {
		var ret string
		return ret
	}

	return *o.HolderName.Get()
}

// GetHolderNameOk returns a tuple with the HolderName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KYCDocumentResponse) GetHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HolderName.Get(), o.HolderName.IsSet()
}

// SetHolderName sets field value
func (o *KYCDocumentResponse) SetHolderName(v string) {
	o.HolderName.Set(&v)
}

// GetDocumentImages returns the DocumentImages field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *KYCDocumentResponse) GetDocumentImages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DocumentImages
}

// GetDocumentImagesOk returns a tuple with the DocumentImages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KYCDocumentResponse) GetDocumentImagesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.DocumentImages) {
		return nil, false
	}
	return o.DocumentImages, true
}

// SetDocumentImages sets field value
func (o *KYCDocumentResponse) SetDocumentImages(v []string) {
	o.DocumentImages = v
}

func (o KYCDocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KYCDocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country"] = o.Country
	toSerialize["type"] = o.Type
	toSerialize["sub_type"] = o.SubType
	toSerialize["document_name"] = o.DocumentName.Get()

	toSerialize["document_number"] = o.DocumentNumber.Get()

	toSerialize["expires_at"] = o.ExpiresAt.Get()

	toSerialize["holder_name"] = o.HolderName.Get()

	if o.DocumentImages != nil {
		toSerialize["document_images"] = o.DocumentImages
    }
	return toSerialize, nil
}

type NullableKYCDocumentResponse struct {
	value *KYCDocumentResponse
	isSet bool
}

func (v NullableKYCDocumentResponse) Get() *KYCDocumentResponse {
	return v.value
}

func (v *NullableKYCDocumentResponse) Set(val *KYCDocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKYCDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKYCDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKYCDocumentResponse(val *KYCDocumentResponse) *NullableKYCDocumentResponse {
	return &NullableKYCDocumentResponse{value: val, isSet: true}
}

func (v NullableKYCDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKYCDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



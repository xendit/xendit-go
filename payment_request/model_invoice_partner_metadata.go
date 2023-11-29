/*
Payment Requests

This API is used for Payment Requests

API version: 1.45.2
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the InvoicePartnerMetadata type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InvoicePartnerMetadata{}

// InvoicePartnerMetadata struct for InvoicePartnerMetadata
type InvoicePartnerMetadata struct {
	Notes *string `json:"notes,omitempty"`
}

// NewInvoicePartnerMetadata instantiates a new InvoicePartnerMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicePartnerMetadata() *InvoicePartnerMetadata {
	this := InvoicePartnerMetadata{}
	return &this
}

// NewInvoicePartnerMetadataWithDefaults instantiates a new InvoicePartnerMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicePartnerMetadataWithDefaults() *InvoicePartnerMetadata {
	this := InvoicePartnerMetadata{}
	return &this
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InvoicePartnerMetadata) GetNotes() string {
	if o == nil || utils.IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicePartnerMetadata) GetNotesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InvoicePartnerMetadata) HasNotes() bool {
	if o != nil && !utils.IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InvoicePartnerMetadata) SetNotes(v string) {
	o.Notes = &v
}

func (o InvoicePartnerMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoicePartnerMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableInvoicePartnerMetadata struct {
	value *InvoicePartnerMetadata
	isSet bool
}

func (v NullableInvoicePartnerMetadata) Get() *InvoicePartnerMetadata {
	return v.value
}

func (v *NullableInvoicePartnerMetadata) Set(val *InvoicePartnerMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicePartnerMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicePartnerMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicePartnerMetadata(val *InvoicePartnerMetadata) *NullableInvoicePartnerMetadata {
	return &NullableInvoicePartnerMetadata{value: val, isSet: true}
}

func (v NullableInvoicePartnerMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicePartnerMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



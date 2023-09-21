/*
Payout Service

This API allows Xendit to send money from an account to a channel (banks, eWallets, retail outlets) from across regions

API version: 1.0.0
*/


package payout

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the GetPayouts200ResponseLinks type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetPayouts200ResponseLinks{}

// GetPayouts200ResponseLinks struct for GetPayouts200ResponseLinks
type GetPayouts200ResponseLinks struct {
	Href *string `json:"href,omitempty"`
	Rel *string `json:"rel,omitempty"`
	Method *string `json:"method,omitempty"`
}

// NewGetPayouts200ResponseLinks instantiates a new GetPayouts200ResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayouts200ResponseLinks() *GetPayouts200ResponseLinks {
	this := GetPayouts200ResponseLinks{}
	return &this
}

// NewGetPayouts200ResponseLinksWithDefaults instantiates a new GetPayouts200ResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayouts200ResponseLinksWithDefaults() *GetPayouts200ResponseLinks {
	this := GetPayouts200ResponseLinks{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *GetPayouts200ResponseLinks) GetHref() string {
	if o == nil || utils.IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayouts200ResponseLinks) GetHrefOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *GetPayouts200ResponseLinks) HasHref() bool {
	if o != nil && !utils.IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *GetPayouts200ResponseLinks) SetHref(v string) {
	o.Href = &v
}

// GetRel returns the Rel field value if set, zero value otherwise.
func (o *GetPayouts200ResponseLinks) GetRel() string {
	if o == nil || utils.IsNil(o.Rel) {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayouts200ResponseLinks) GetRelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Rel) {
		return nil, false
	}
	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *GetPayouts200ResponseLinks) HasRel() bool {
	if o != nil && !utils.IsNil(o.Rel) {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *GetPayouts200ResponseLinks) SetRel(v string) {
	o.Rel = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *GetPayouts200ResponseLinks) GetMethod() string {
	if o == nil || utils.IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayouts200ResponseLinks) GetMethodOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *GetPayouts200ResponseLinks) HasMethod() bool {
	if o != nil && !utils.IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *GetPayouts200ResponseLinks) SetMethod(v string) {
	o.Method = &v
}

func (o GetPayouts200ResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayouts200ResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !utils.IsNil(o.Rel) {
		toSerialize["rel"] = o.Rel
	}
	if !utils.IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	return toSerialize, nil
}

type NullableGetPayouts200ResponseLinks struct {
	value *GetPayouts200ResponseLinks
	isSet bool
}

func (v NullableGetPayouts200ResponseLinks) Get() *GetPayouts200ResponseLinks {
	return v.value
}

func (v *NullableGetPayouts200ResponseLinks) Set(val *GetPayouts200ResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayouts200ResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayouts200ResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayouts200ResponseLinks(val *GetPayouts200ResponseLinks) *NullableGetPayouts200ResponseLinks {
	return &NullableGetPayouts200ResponseLinks{value: val, isSet: true}
}

func (v NullableGetPayouts200ResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayouts200ResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



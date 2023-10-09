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

// checks if the AccountCard type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AccountCard{}

// AccountCard struct for AccountCard
type AccountCard struct {
	// The token id returned in tokenisation
	TokenId *string `json:"token_id,omitempty"`
}

// NewAccountCard instantiates a new AccountCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCard() *AccountCard {
	this := AccountCard{}
	return &this
}

// NewAccountCardWithDefaults instantiates a new AccountCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCardWithDefaults() *AccountCard {
	this := AccountCard{}
	return &this
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *AccountCard) GetTokenId() string {
	if o == nil || utils.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCard) GetTokenIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *AccountCard) HasTokenId() bool {
	if o != nil && !utils.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *AccountCard) SetTokenId(v string) {
	o.TokenId = &v
}

func (o AccountCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	return toSerialize, nil
}

type NullableAccountCard struct {
	value *AccountCard
	isSet bool
}

func (v NullableAccountCard) Get() *AccountCard {
	return v.value
}

func (v *NullableAccountCard) Set(val *AccountCard) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCard) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCard(val *AccountCard) *NullableAccountCard {
	return &NullableAccountCard{value: val, isSet: true}
}

func (v NullableAccountCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



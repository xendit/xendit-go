/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.91.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the EWalletChannelProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EWalletChannelProperties{}

// EWalletChannelProperties EWallet Channel Properties
type EWalletChannelProperties struct {
	// URL where the end-customer is redirected if the authorization is successful
	SuccessReturnUrl *string `json:"success_return_url,omitempty"`
	// URL where the end-customer is redirected if the authorization failed
	FailureReturnUrl *string `json:"failure_return_url,omitempty"`
	// URL where the end-customer is redirected if the authorization cancelled
	CancelReturnUrl *string `json:"cancel_return_url,omitempty"`
	// Mobile number of customer in E.164 format (e.g. +628123123123). For OVO one time payment use only.
	MobileNumber *string `json:"mobile_number,omitempty"`
	// REDEEM_NONE will not use any point, REDEEM_ALL will use all available points before cash balance is used. For OVO and ShopeePay tokenized payment use only.
	RedeemPoints *string `json:"redeem_points,omitempty"`
	// Available for JENIUSPAY only
	Cashtag *string `json:"cashtag,omitempty"`
}

// NewEWalletChannelProperties instantiates a new EWalletChannelProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEWalletChannelProperties() *EWalletChannelProperties {
	this := EWalletChannelProperties{}
	return &this
}

// NewEWalletChannelPropertiesWithDefaults instantiates a new EWalletChannelProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEWalletChannelPropertiesWithDefaults() *EWalletChannelProperties {
	this := EWalletChannelProperties{}
	return &this
}

// GetSuccessReturnUrl returns the SuccessReturnUrl field value if set, zero value otherwise.
func (o *EWalletChannelProperties) GetSuccessReturnUrl() string {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		var ret string
		return ret
	}
	return *o.SuccessReturnUrl
}

// GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWalletChannelProperties) GetSuccessReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		return nil, false
	}
	return o.SuccessReturnUrl, true
}

// HasSuccessReturnUrl returns a boolean if a field has been set.
func (o *EWalletChannelProperties) HasSuccessReturnUrl() bool {
	if o != nil && !utils.IsNil(o.SuccessReturnUrl) {
		return true
	}

	return false
}

// SetSuccessReturnUrl gets a reference to the given string and assigns it to the SuccessReturnUrl field.
func (o *EWalletChannelProperties) SetSuccessReturnUrl(v string) {
	o.SuccessReturnUrl = &v
}

// GetFailureReturnUrl returns the FailureReturnUrl field value if set, zero value otherwise.
func (o *EWalletChannelProperties) GetFailureReturnUrl() string {
	if o == nil || utils.IsNil(o.FailureReturnUrl) {
		var ret string
		return ret
	}
	return *o.FailureReturnUrl
}

// GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWalletChannelProperties) GetFailureReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FailureReturnUrl) {
		return nil, false
	}
	return o.FailureReturnUrl, true
}

// HasFailureReturnUrl returns a boolean if a field has been set.
func (o *EWalletChannelProperties) HasFailureReturnUrl() bool {
	if o != nil && !utils.IsNil(o.FailureReturnUrl) {
		return true
	}

	return false
}

// SetFailureReturnUrl gets a reference to the given string and assigns it to the FailureReturnUrl field.
func (o *EWalletChannelProperties) SetFailureReturnUrl(v string) {
	o.FailureReturnUrl = &v
}

// GetCancelReturnUrl returns the CancelReturnUrl field value if set, zero value otherwise.
func (o *EWalletChannelProperties) GetCancelReturnUrl() string {
	if o == nil || utils.IsNil(o.CancelReturnUrl) {
		var ret string
		return ret
	}
	return *o.CancelReturnUrl
}

// GetCancelReturnUrlOk returns a tuple with the CancelReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWalletChannelProperties) GetCancelReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CancelReturnUrl) {
		return nil, false
	}
	return o.CancelReturnUrl, true
}

// HasCancelReturnUrl returns a boolean if a field has been set.
func (o *EWalletChannelProperties) HasCancelReturnUrl() bool {
	if o != nil && !utils.IsNil(o.CancelReturnUrl) {
		return true
	}

	return false
}

// SetCancelReturnUrl gets a reference to the given string and assigns it to the CancelReturnUrl field.
func (o *EWalletChannelProperties) SetCancelReturnUrl(v string) {
	o.CancelReturnUrl = &v
}

// GetMobileNumber returns the MobileNumber field value if set, zero value otherwise.
func (o *EWalletChannelProperties) GetMobileNumber() string {
	if o == nil || utils.IsNil(o.MobileNumber) {
		var ret string
		return ret
	}
	return *o.MobileNumber
}

// GetMobileNumberOk returns a tuple with the MobileNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWalletChannelProperties) GetMobileNumberOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MobileNumber) {
		return nil, false
	}
	return o.MobileNumber, true
}

// HasMobileNumber returns a boolean if a field has been set.
func (o *EWalletChannelProperties) HasMobileNumber() bool {
	if o != nil && !utils.IsNil(o.MobileNumber) {
		return true
	}

	return false
}

// SetMobileNumber gets a reference to the given string and assigns it to the MobileNumber field.
func (o *EWalletChannelProperties) SetMobileNumber(v string) {
	o.MobileNumber = &v
}

// GetRedeemPoints returns the RedeemPoints field value if set, zero value otherwise.
func (o *EWalletChannelProperties) GetRedeemPoints() string {
	if o == nil || utils.IsNil(o.RedeemPoints) {
		var ret string
		return ret
	}
	return *o.RedeemPoints
}

// GetRedeemPointsOk returns a tuple with the RedeemPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWalletChannelProperties) GetRedeemPointsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RedeemPoints) {
		return nil, false
	}
	return o.RedeemPoints, true
}

// HasRedeemPoints returns a boolean if a field has been set.
func (o *EWalletChannelProperties) HasRedeemPoints() bool {
	if o != nil && !utils.IsNil(o.RedeemPoints) {
		return true
	}

	return false
}

// SetRedeemPoints gets a reference to the given string and assigns it to the RedeemPoints field.
func (o *EWalletChannelProperties) SetRedeemPoints(v string) {
	o.RedeemPoints = &v
}

// GetCashtag returns the Cashtag field value if set, zero value otherwise.
func (o *EWalletChannelProperties) GetCashtag() string {
	if o == nil || utils.IsNil(o.Cashtag) {
		var ret string
		return ret
	}
	return *o.Cashtag
}

// GetCashtagOk returns a tuple with the Cashtag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EWalletChannelProperties) GetCashtagOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Cashtag) {
		return nil, false
	}
	return o.Cashtag, true
}

// HasCashtag returns a boolean if a field has been set.
func (o *EWalletChannelProperties) HasCashtag() bool {
	if o != nil && !utils.IsNil(o.Cashtag) {
		return true
	}

	return false
}

// SetCashtag gets a reference to the given string and assigns it to the Cashtag field.
func (o *EWalletChannelProperties) SetCashtag(v string) {
	o.Cashtag = &v
}

func (o EWalletChannelProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EWalletChannelProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.SuccessReturnUrl) {
		toSerialize["success_return_url"] = o.SuccessReturnUrl
	}
	if !utils.IsNil(o.FailureReturnUrl) {
		toSerialize["failure_return_url"] = o.FailureReturnUrl
	}
	if !utils.IsNil(o.CancelReturnUrl) {
		toSerialize["cancel_return_url"] = o.CancelReturnUrl
	}
	if !utils.IsNil(o.MobileNumber) {
		toSerialize["mobile_number"] = o.MobileNumber
	}
	if !utils.IsNil(o.RedeemPoints) {
		toSerialize["redeem_points"] = o.RedeemPoints
	}
	if !utils.IsNil(o.Cashtag) {
		toSerialize["cashtag"] = o.Cashtag
	}
	return toSerialize, nil
}

type NullableEWalletChannelProperties struct {
	value *EWalletChannelProperties
	isSet bool
}

func (v NullableEWalletChannelProperties) Get() *EWalletChannelProperties {
	return v.value
}

func (v *NullableEWalletChannelProperties) Set(val *EWalletChannelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableEWalletChannelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableEWalletChannelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEWalletChannelProperties(val *EWalletChannelProperties) *NullableEWalletChannelProperties {
	return &NullableEWalletChannelProperties{value: val, isSet: true}
}

func (v NullableEWalletChannelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEWalletChannelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



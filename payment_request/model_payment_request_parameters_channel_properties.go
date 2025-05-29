/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the PaymentRequestParametersChannelProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestParametersChannelProperties{}

// PaymentRequestParametersChannelProperties struct for PaymentRequestParametersChannelProperties
type PaymentRequestParametersChannelProperties struct {
	// URL where the end-customer is redirected if the authorization is successful
	SuccessReturnUrl *string `json:"success_return_url,omitempty"`
	// URL where the end-customer is redirected if the authorization failed
	FailureReturnUrl *string `json:"failure_return_url,omitempty"`
	// URL where the end-customer is redirected if the authorization cancelled
	CancelReturnUrl *string `json:"cancel_return_url,omitempty"`
	// URL where the end-customer is redirected if the authorization is pending
	PendingReturnUrl *string `json:"pending_return_url,omitempty"`
	// REDEEM_NONE will not use any point, REDEEM_ALL will use all available points before cash balance is used. For OVO and ShopeePay tokenized payment use only.
	RedeemPoints *string `json:"redeem_points,omitempty"`
	// Toggle used to require end-customer to input undergo OTP validation before completing a payment. OTP will always be required for transactions greater than 1,000,000 IDR. For BRI tokenized payment use only.
	RequireAuth *bool `json:"require_auth,omitempty"`
	// Tag for a Merchant ID that you want to associate this payment with. For merchants using their own MIDs to specify which MID they want to use 
	MerchantIdTag *string `json:"merchant_id_tag,omitempty"`
	// Type of “credential-on-file” / “card-on-file” payment being made. Indicate that this payment uses a previously linked Payment Method for charging.
	CardonfileType NullableString `json:"cardonfile_type,omitempty"`
	// Three digit code written on the back of the card (usually called CVV/CVN).
	Cvv *string `json:"cvv,omitempty"`
}

// NewPaymentRequestParametersChannelProperties instantiates a new PaymentRequestParametersChannelProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestParametersChannelProperties() *PaymentRequestParametersChannelProperties {
	this := PaymentRequestParametersChannelProperties{}
	return &this
}

// NewPaymentRequestParametersChannelPropertiesWithDefaults instantiates a new PaymentRequestParametersChannelProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestParametersChannelPropertiesWithDefaults() *PaymentRequestParametersChannelProperties {
	this := PaymentRequestParametersChannelProperties{}
	return &this
}

// GetSuccessReturnUrl returns the SuccessReturnUrl field value if set, zero value otherwise.
func (o *PaymentRequestParametersChannelProperties) GetSuccessReturnUrl() string {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		var ret string
		return ret
	}
	return *o.SuccessReturnUrl
}

// GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParametersChannelProperties) GetSuccessReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SuccessReturnUrl) {
		return nil, false
	}
	return o.SuccessReturnUrl, true
}

// HasSuccessReturnUrl returns a boolean if a field has been set.
func (o *PaymentRequestParametersChannelProperties) HasSuccessReturnUrl() bool {
	if o != nil && !utils.IsNil(o.SuccessReturnUrl) {
		return true
	}

	return false
}

// SetSuccessReturnUrl gets a reference to the given string and assigns it to the SuccessReturnUrl field.
func (o *PaymentRequestParametersChannelProperties) SetSuccessReturnUrl(v string) {
	o.SuccessReturnUrl = &v
}

// GetFailureReturnUrl returns the FailureReturnUrl field value if set, zero value otherwise.
func (o *PaymentRequestParametersChannelProperties) GetFailureReturnUrl() string {
	if o == nil || utils.IsNil(o.FailureReturnUrl) {
		var ret string
		return ret
	}
	return *o.FailureReturnUrl
}

// GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParametersChannelProperties) GetFailureReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FailureReturnUrl) {
		return nil, false
	}
	return o.FailureReturnUrl, true
}

// HasFailureReturnUrl returns a boolean if a field has been set.
func (o *PaymentRequestParametersChannelProperties) HasFailureReturnUrl() bool {
	if o != nil && !utils.IsNil(o.FailureReturnUrl) {
		return true
	}

	return false
}

// SetFailureReturnUrl gets a reference to the given string and assigns it to the FailureReturnUrl field.
func (o *PaymentRequestParametersChannelProperties) SetFailureReturnUrl(v string) {
	o.FailureReturnUrl = &v
}

// GetCancelReturnUrl returns the CancelReturnUrl field value if set, zero value otherwise.
func (o *PaymentRequestParametersChannelProperties) GetCancelReturnUrl() string {
	if o == nil || utils.IsNil(o.CancelReturnUrl) {
		var ret string
		return ret
	}
	return *o.CancelReturnUrl
}

// GetCancelReturnUrlOk returns a tuple with the CancelReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParametersChannelProperties) GetCancelReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CancelReturnUrl) {
		return nil, false
	}
	return o.CancelReturnUrl, true
}

// HasCancelReturnUrl returns a boolean if a field has been set.
func (o *PaymentRequestParametersChannelProperties) HasCancelReturnUrl() bool {
	if o != nil && !utils.IsNil(o.CancelReturnUrl) {
		return true
	}

	return false
}

// SetCancelReturnUrl gets a reference to the given string and assigns it to the CancelReturnUrl field.
func (o *PaymentRequestParametersChannelProperties) SetCancelReturnUrl(v string) {
	o.CancelReturnUrl = &v
}

// GetPendingReturnUrl returns the PendingReturnUrl field value if set, zero value otherwise.
func (o *PaymentRequestParametersChannelProperties) GetPendingReturnUrl() string {
	if o == nil || utils.IsNil(o.PendingReturnUrl) {
		var ret string
		return ret
	}
	return *o.PendingReturnUrl
}

// GetPendingReturnUrlOk returns a tuple with the PendingReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParametersChannelProperties) GetPendingReturnUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PendingReturnUrl) {
		return nil, false
	}
	return o.PendingReturnUrl, true
}

// HasPendingReturnUrl returns a boolean if a field has been set.
func (o *PaymentRequestParametersChannelProperties) HasPendingReturnUrl() bool {
	if o != nil && !utils.IsNil(o.PendingReturnUrl) {
		return true
	}

	return false
}

// SetPendingReturnUrl gets a reference to the given string and assigns it to the PendingReturnUrl field.
func (o *PaymentRequestParametersChannelProperties) SetPendingReturnUrl(v string) {
	o.PendingReturnUrl = &v
}

// GetRedeemPoints returns the RedeemPoints field value if set, zero value otherwise.
func (o *PaymentRequestParametersChannelProperties) GetRedeemPoints() string {
	if o == nil || utils.IsNil(o.RedeemPoints) {
		var ret string
		return ret
	}
	return *o.RedeemPoints
}

// GetRedeemPointsOk returns a tuple with the RedeemPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParametersChannelProperties) GetRedeemPointsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RedeemPoints) {
		return nil, false
	}
	return o.RedeemPoints, true
}

// HasRedeemPoints returns a boolean if a field has been set.
func (o *PaymentRequestParametersChannelProperties) HasRedeemPoints() bool {
	if o != nil && !utils.IsNil(o.RedeemPoints) {
		return true
	}

	return false
}

// SetRedeemPoints gets a reference to the given string and assigns it to the RedeemPoints field.
func (o *PaymentRequestParametersChannelProperties) SetRedeemPoints(v string) {
	o.RedeemPoints = &v
}

// GetRequireAuth returns the RequireAuth field value if set, zero value otherwise.
func (o *PaymentRequestParametersChannelProperties) GetRequireAuth() bool {
	if o == nil || utils.IsNil(o.RequireAuth) {
		var ret bool
		return ret
	}
	return *o.RequireAuth
}

// GetRequireAuthOk returns a tuple with the RequireAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParametersChannelProperties) GetRequireAuthOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RequireAuth) {
		return nil, false
	}
	return o.RequireAuth, true
}

// HasRequireAuth returns a boolean if a field has been set.
func (o *PaymentRequestParametersChannelProperties) HasRequireAuth() bool {
	if o != nil && !utils.IsNil(o.RequireAuth) {
		return true
	}

	return false
}

// SetRequireAuth gets a reference to the given bool and assigns it to the RequireAuth field.
func (o *PaymentRequestParametersChannelProperties) SetRequireAuth(v bool) {
	o.RequireAuth = &v
}

// GetMerchantIdTag returns the MerchantIdTag field value if set, zero value otherwise.
func (o *PaymentRequestParametersChannelProperties) GetMerchantIdTag() string {
	if o == nil || utils.IsNil(o.MerchantIdTag) {
		var ret string
		return ret
	}
	return *o.MerchantIdTag
}

// GetMerchantIdTagOk returns a tuple with the MerchantIdTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParametersChannelProperties) GetMerchantIdTagOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MerchantIdTag) {
		return nil, false
	}
	return o.MerchantIdTag, true
}

// HasMerchantIdTag returns a boolean if a field has been set.
func (o *PaymentRequestParametersChannelProperties) HasMerchantIdTag() bool {
	if o != nil && !utils.IsNil(o.MerchantIdTag) {
		return true
	}

	return false
}

// SetMerchantIdTag gets a reference to the given string and assigns it to the MerchantIdTag field.
func (o *PaymentRequestParametersChannelProperties) SetMerchantIdTag(v string) {
	o.MerchantIdTag = &v
}

// GetCardonfileType returns the CardonfileType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestParametersChannelProperties) GetCardonfileType() string {
	if o == nil || utils.IsNil(o.CardonfileType.Get()) {
		var ret string
		return ret
	}
	return *o.CardonfileType.Get()
}

// GetCardonfileTypeOk returns a tuple with the CardonfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestParametersChannelProperties) GetCardonfileTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardonfileType.Get(), o.CardonfileType.IsSet()
}

// HasCardonfileType returns a boolean if a field has been set.
func (o *PaymentRequestParametersChannelProperties) HasCardonfileType() bool {
	if o != nil && o.CardonfileType.IsSet() {
		return true
	}

	return false
}

// SetCardonfileType gets a reference to the given NullableString and assigns it to the CardonfileType field.
func (o *PaymentRequestParametersChannelProperties) SetCardonfileType(v string) {
	o.CardonfileType.Set(&v)
}
// SetCardonfileTypeNil sets the value for CardonfileType to be an explicit nil
func (o *PaymentRequestParametersChannelProperties) SetCardonfileTypeNil() {
	o.CardonfileType.Set(nil)
}

// UnsetCardonfileType ensures that no value is present for CardonfileType, not even an explicit nil
func (o *PaymentRequestParametersChannelProperties) UnsetCardonfileType() {
	o.CardonfileType.Unset()
}

// GetCvv returns the Cvv field value if set, zero value otherwise.
func (o *PaymentRequestParametersChannelProperties) GetCvv() string {
	if o == nil || utils.IsNil(o.Cvv) {
		var ret string
		return ret
	}
	return *o.Cvv
}

// GetCvvOk returns a tuple with the Cvv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestParametersChannelProperties) GetCvvOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Cvv) {
		return nil, false
	}
	return o.Cvv, true
}

// HasCvv returns a boolean if a field has been set.
func (o *PaymentRequestParametersChannelProperties) HasCvv() bool {
	if o != nil && !utils.IsNil(o.Cvv) {
		return true
	}

	return false
}

// SetCvv gets a reference to the given string and assigns it to the Cvv field.
func (o *PaymentRequestParametersChannelProperties) SetCvv(v string) {
	o.Cvv = &v
}

func (o PaymentRequestParametersChannelProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestParametersChannelProperties) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.PendingReturnUrl) {
		toSerialize["pending_return_url"] = o.PendingReturnUrl
	}
	if !utils.IsNil(o.RedeemPoints) {
		toSerialize["redeem_points"] = o.RedeemPoints
	}
	if !utils.IsNil(o.RequireAuth) {
		toSerialize["require_auth"] = o.RequireAuth
	}
	if !utils.IsNil(o.MerchantIdTag) {
		toSerialize["merchant_id_tag"] = o.MerchantIdTag
	}
	if o.CardonfileType.IsSet() {
		toSerialize["cardonfile_type"] = o.CardonfileType.Get()
    }
	if !utils.IsNil(o.Cvv) {
		toSerialize["cvv"] = o.Cvv
	}
	return toSerialize, nil
}

type NullablePaymentRequestParametersChannelProperties struct {
	value *PaymentRequestParametersChannelProperties
	isSet bool
}

func (v NullablePaymentRequestParametersChannelProperties) Get() *PaymentRequestParametersChannelProperties {
	return v.value
}

func (v *NullablePaymentRequestParametersChannelProperties) Set(val *PaymentRequestParametersChannelProperties) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestParametersChannelProperties) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestParametersChannelProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestParametersChannelProperties(val *PaymentRequestParametersChannelProperties) *NullablePaymentRequestParametersChannelProperties {
	return &NullablePaymentRequestParametersChannelProperties{value: val, isSet: true}
}

func (v NullablePaymentRequestParametersChannelProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestParametersChannelProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.128.0
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the PaymentMethodUpdateParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentMethodUpdateParameters{}

// PaymentMethodUpdateParameters struct for PaymentMethodUpdateParameters
type PaymentMethodUpdateParameters struct {
	Description *string `json:"description,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
	Reusability *PaymentMethodReusability `json:"reusability,omitempty"`
	Status *PaymentMethodStatus `json:"status,omitempty"`
	OverTheCounter *OverTheCounterUpdateParameters `json:"over_the_counter,omitempty"`
	VirtualAccount *VirtualAccountUpdateParameters `json:"virtual_account,omitempty"`
}

// NewPaymentMethodUpdateParameters instantiates a new PaymentMethodUpdateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodUpdateParameters() *PaymentMethodUpdateParameters {
	this := PaymentMethodUpdateParameters{}
	return &this
}

// NewPaymentMethodUpdateParametersWithDefaults instantiates a new PaymentMethodUpdateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodUpdateParametersWithDefaults() *PaymentMethodUpdateParameters {
	this := PaymentMethodUpdateParameters{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentMethodUpdateParameters) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateParameters) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentMethodUpdateParameters) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentMethodUpdateParameters) SetDescription(v string) {
	o.Description = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *PaymentMethodUpdateParameters) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateParameters) GetReferenceIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *PaymentMethodUpdateParameters) HasReferenceId() bool {
	if o != nil && !utils.IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *PaymentMethodUpdateParameters) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetReusability returns the Reusability field value if set, zero value otherwise.
func (o *PaymentMethodUpdateParameters) GetReusability() PaymentMethodReusability {
	if o == nil || utils.IsNil(o.Reusability) {
		var ret PaymentMethodReusability
		return ret
	}
	return *o.Reusability
}

// GetReusabilityOk returns a tuple with the Reusability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateParameters) GetReusabilityOk() (*PaymentMethodReusability, bool) {
	if o == nil || utils.IsNil(o.Reusability) {
		return nil, false
	}
	return o.Reusability, true
}

// HasReusability returns a boolean if a field has been set.
func (o *PaymentMethodUpdateParameters) HasReusability() bool {
	if o != nil && !utils.IsNil(o.Reusability) {
		return true
	}

	return false
}

// SetReusability gets a reference to the given PaymentMethodReusability and assigns it to the Reusability field.
func (o *PaymentMethodUpdateParameters) SetReusability(v PaymentMethodReusability) {
	o.Reusability = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentMethodUpdateParameters) GetStatus() PaymentMethodStatus {
	if o == nil || utils.IsNil(o.Status) {
		var ret PaymentMethodStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateParameters) GetStatusOk() (*PaymentMethodStatus, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentMethodUpdateParameters) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PaymentMethodStatus and assigns it to the Status field.
func (o *PaymentMethodUpdateParameters) SetStatus(v PaymentMethodStatus) {
	o.Status = &v
}

// GetOverTheCounter returns the OverTheCounter field value if set, zero value otherwise.
func (o *PaymentMethodUpdateParameters) GetOverTheCounter() OverTheCounterUpdateParameters {
	if o == nil || utils.IsNil(o.OverTheCounter) {
		var ret OverTheCounterUpdateParameters
		return ret
	}
	return *o.OverTheCounter
}

// GetOverTheCounterOk returns a tuple with the OverTheCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateParameters) GetOverTheCounterOk() (*OverTheCounterUpdateParameters, bool) {
	if o == nil || utils.IsNil(o.OverTheCounter) {
		return nil, false
	}
	return o.OverTheCounter, true
}

// HasOverTheCounter returns a boolean if a field has been set.
func (o *PaymentMethodUpdateParameters) HasOverTheCounter() bool {
	if o != nil && !utils.IsNil(o.OverTheCounter) {
		return true
	}

	return false
}

// SetOverTheCounter gets a reference to the given OverTheCounterUpdateParameters and assigns it to the OverTheCounter field.
func (o *PaymentMethodUpdateParameters) SetOverTheCounter(v OverTheCounterUpdateParameters) {
	o.OverTheCounter = &v
}

// GetVirtualAccount returns the VirtualAccount field value if set, zero value otherwise.
func (o *PaymentMethodUpdateParameters) GetVirtualAccount() VirtualAccountUpdateParameters {
	if o == nil || utils.IsNil(o.VirtualAccount) {
		var ret VirtualAccountUpdateParameters
		return ret
	}
	return *o.VirtualAccount
}

// GetVirtualAccountOk returns a tuple with the VirtualAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodUpdateParameters) GetVirtualAccountOk() (*VirtualAccountUpdateParameters, bool) {
	if o == nil || utils.IsNil(o.VirtualAccount) {
		return nil, false
	}
	return o.VirtualAccount, true
}

// HasVirtualAccount returns a boolean if a field has been set.
func (o *PaymentMethodUpdateParameters) HasVirtualAccount() bool {
	if o != nil && !utils.IsNil(o.VirtualAccount) {
		return true
	}

	return false
}

// SetVirtualAccount gets a reference to the given VirtualAccountUpdateParameters and assigns it to the VirtualAccount field.
func (o *PaymentMethodUpdateParameters) SetVirtualAccount(v VirtualAccountUpdateParameters) {
	o.VirtualAccount = &v
}

func (o PaymentMethodUpdateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodUpdateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !utils.IsNil(o.Reusability) {
		toSerialize["reusability"] = o.Reusability
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.OverTheCounter) {
		toSerialize["over_the_counter"] = o.OverTheCounter
	}
	if !utils.IsNil(o.VirtualAccount) {
		toSerialize["virtual_account"] = o.VirtualAccount
	}
	return toSerialize, nil
}

type NullablePaymentMethodUpdateParameters struct {
	value *PaymentMethodUpdateParameters
	isSet bool
}

func (v NullablePaymentMethodUpdateParameters) Get() *PaymentMethodUpdateParameters {
	return v.value
}

func (v *NullablePaymentMethodUpdateParameters) Set(val *PaymentMethodUpdateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodUpdateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodUpdateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodUpdateParameters(val *PaymentMethodUpdateParameters) *NullablePaymentMethodUpdateParameters {
	return &NullablePaymentMethodUpdateParameters{value: val, isSet: true}
}

func (v NullablePaymentMethodUpdateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodUpdateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



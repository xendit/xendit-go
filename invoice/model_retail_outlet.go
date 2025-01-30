/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.8.7
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the RetailOutlet type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RetailOutlet{}

// RetailOutlet An object representing retail outlet details for invoices.
type RetailOutlet struct {
	RetailOutletName RetailOutletName `json:"retail_outlet_name"`
	// The payment code.
	PaymentCode *string `json:"payment_code,omitempty"`
	// The transfer amount.
	TransferAmount *float64 `json:"transfer_amount,omitempty"`
	// The name of the merchant.
	MerchantName *string `json:"merchant_name,omitempty"`
}

// NewRetailOutlet instantiates a new RetailOutlet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetailOutlet(retailOutletName RetailOutletName) *RetailOutlet {
	this := RetailOutlet{}
	this.RetailOutletName = retailOutletName
	return &this
}

// NewRetailOutletWithDefaults instantiates a new RetailOutlet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetailOutletWithDefaults() *RetailOutlet {
	this := RetailOutlet{}
	return &this
}

// GetRetailOutletName returns the RetailOutletName field value
func (o *RetailOutlet) GetRetailOutletName() RetailOutletName {
	if o == nil {
		var ret RetailOutletName
		return ret
	}

	return o.RetailOutletName
}

// GetRetailOutletNameOk returns a tuple with the RetailOutletName field value
// and a boolean to check if the value has been set.
func (o *RetailOutlet) GetRetailOutletNameOk() (*RetailOutletName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetailOutletName, true
}

// SetRetailOutletName sets field value
func (o *RetailOutlet) SetRetailOutletName(v RetailOutletName) {
	o.RetailOutletName = v
}

// GetPaymentCode returns the PaymentCode field value if set, zero value otherwise.
func (o *RetailOutlet) GetPaymentCode() string {
	if o == nil || utils.IsNil(o.PaymentCode) {
		var ret string
		return ret
	}
	return *o.PaymentCode
}

// GetPaymentCodeOk returns a tuple with the PaymentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetailOutlet) GetPaymentCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PaymentCode) {
		return nil, false
	}
	return o.PaymentCode, true
}

// HasPaymentCode returns a boolean if a field has been set.
func (o *RetailOutlet) HasPaymentCode() bool {
	if o != nil && !utils.IsNil(o.PaymentCode) {
		return true
	}

	return false
}

// SetPaymentCode gets a reference to the given string and assigns it to the PaymentCode field.
func (o *RetailOutlet) SetPaymentCode(v string) {
	o.PaymentCode = &v
}

// GetTransferAmount returns the TransferAmount field value if set, zero value otherwise.
func (o *RetailOutlet) GetTransferAmount() float64 {
	if o == nil || utils.IsNil(o.TransferAmount) {
		var ret float64
		return ret
	}
	return *o.TransferAmount
}

// GetTransferAmountOk returns a tuple with the TransferAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetailOutlet) GetTransferAmountOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.TransferAmount) {
		return nil, false
	}
	return o.TransferAmount, true
}

// HasTransferAmount returns a boolean if a field has been set.
func (o *RetailOutlet) HasTransferAmount() bool {
	if o != nil && !utils.IsNil(o.TransferAmount) {
		return true
	}

	return false
}

// SetTransferAmount gets a reference to the given float64 and assigns it to the TransferAmount field.
func (o *RetailOutlet) SetTransferAmount(v float64) {
	o.TransferAmount = &v
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise.
func (o *RetailOutlet) GetMerchantName() string {
	if o == nil || utils.IsNil(o.MerchantName) {
		var ret string
		return ret
	}
	return *o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetailOutlet) GetMerchantNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MerchantName) {
		return nil, false
	}
	return o.MerchantName, true
}

// HasMerchantName returns a boolean if a field has been set.
func (o *RetailOutlet) HasMerchantName() bool {
	if o != nil && !utils.IsNil(o.MerchantName) {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given string and assigns it to the MerchantName field.
func (o *RetailOutlet) SetMerchantName(v string) {
	o.MerchantName = &v
}

func (o RetailOutlet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetailOutlet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["retail_outlet_name"] = o.RetailOutletName
	if !utils.IsNil(o.PaymentCode) {
		toSerialize["payment_code"] = o.PaymentCode
	}
	if !utils.IsNil(o.TransferAmount) {
		toSerialize["transfer_amount"] = o.TransferAmount
	}
	if !utils.IsNil(o.MerchantName) {
		toSerialize["merchant_name"] = o.MerchantName
	}
	return toSerialize, nil
}

type NullableRetailOutlet struct {
	value *RetailOutlet
	isSet bool
}

func (v NullableRetailOutlet) Get() *RetailOutlet {
	return v.value
}

func (v *NullableRetailOutlet) Set(val *RetailOutlet) {
	v.value = val
	v.isSet = true
}

func (v NullableRetailOutlet) IsSet() bool {
	return v.isSet
}

func (v *NullableRetailOutlet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetailOutlet(val *RetailOutlet) *NullableRetailOutlet {
	return &NullableRetailOutlet{value: val, isSet: true}
}

func (v NullableRetailOutlet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetailOutlet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



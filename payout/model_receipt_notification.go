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

// checks if the ReceiptNotification type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReceiptNotification{}

// ReceiptNotification Additional notification for completed payout
type ReceiptNotification struct {
	// Valid email address to send the payout receipt
	EmailTo []string `json:"email_to,omitempty"`
	// Valid email address to cc the payout receipt
	EmailCc []string `json:"email_cc,omitempty"`
	// Valid email address to bcc the payout receipt
	EmailBcc []string `json:"email_bcc,omitempty"`
}

// NewReceiptNotification instantiates a new ReceiptNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiptNotification() *ReceiptNotification {
	this := ReceiptNotification{}
	return &this
}

// NewReceiptNotificationWithDefaults instantiates a new ReceiptNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiptNotificationWithDefaults() *ReceiptNotification {
	this := ReceiptNotification{}
	return &this
}

// GetEmailTo returns the EmailTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptNotification) GetEmailTo() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EmailTo
}

// GetEmailToOk returns a tuple with the EmailTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptNotification) GetEmailToOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.EmailTo) {
		return nil, false
	}
	return o.EmailTo, true
}

// HasEmailTo returns a boolean if a field has been set.
func (o *ReceiptNotification) HasEmailTo() bool {
	if o != nil && utils.IsNil(o.EmailTo) {
		return true
	}

	return false
}

// SetEmailTo gets a reference to the given []string and assigns it to the EmailTo field.
func (o *ReceiptNotification) SetEmailTo(v []string) {
	o.EmailTo = v
}

// GetEmailCc returns the EmailCc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptNotification) GetEmailCc() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EmailCc
}

// GetEmailCcOk returns a tuple with the EmailCc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptNotification) GetEmailCcOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.EmailCc) {
		return nil, false
	}
	return o.EmailCc, true
}

// HasEmailCc returns a boolean if a field has been set.
func (o *ReceiptNotification) HasEmailCc() bool {
	if o != nil && utils.IsNil(o.EmailCc) {
		return true
	}

	return false
}

// SetEmailCc gets a reference to the given []string and assigns it to the EmailCc field.
func (o *ReceiptNotification) SetEmailCc(v []string) {
	o.EmailCc = v
}

// GetEmailBcc returns the EmailBcc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReceiptNotification) GetEmailBcc() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EmailBcc
}

// GetEmailBccOk returns a tuple with the EmailBcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReceiptNotification) GetEmailBccOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.EmailBcc) {
		return nil, false
	}
	return o.EmailBcc, true
}

// HasEmailBcc returns a boolean if a field has been set.
func (o *ReceiptNotification) HasEmailBcc() bool {
	if o != nil && utils.IsNil(o.EmailBcc) {
		return true
	}

	return false
}

// SetEmailBcc gets a reference to the given []string and assigns it to the EmailBcc field.
func (o *ReceiptNotification) SetEmailBcc(v []string) {
	o.EmailBcc = v
}

func (o ReceiptNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiptNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailTo != nil {
		toSerialize["email_to"] = o.EmailTo
	}
	if o.EmailCc != nil {
		toSerialize["email_cc"] = o.EmailCc
	}
	if o.EmailBcc != nil {
		toSerialize["email_bcc"] = o.EmailBcc
	}
	return toSerialize, nil
}

type NullableReceiptNotification struct {
	value *ReceiptNotification
	isSet bool
}

func (v NullableReceiptNotification) Get() *ReceiptNotification {
	return v.value
}

func (v *NullableReceiptNotification) Set(val *ReceiptNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiptNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiptNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiptNotification(val *ReceiptNotification) *NullableReceiptNotification {
	return &NullableReceiptNotification{value: val, isSet: true}
}

func (v NullableReceiptNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiptNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



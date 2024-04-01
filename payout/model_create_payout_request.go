/*
Payout Service

This API allows Xendit to send money from an account to a channel (banks, eWallets, retail outlets) from across regions

API version: 1.0.0
*/


package payout

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v5/utils"
)

// checks if the CreatePayoutRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreatePayoutRequest{}

// CreatePayoutRequest Information needed by Xendit to send money to the destination channel provided
type CreatePayoutRequest struct {
	// A client defined payout identifier
	ReferenceId string `json:"reference_id"`
	// Channel code of selected destination bank or e-wallet
	ChannelCode string `json:"channel_code"`
	ChannelProperties DigitalPayoutChannelProperties `json:"channel_properties"`
	// Amount to be sent to the destination account and should be a multiple of the minimum increment for the selected channel
	Amount float32 `json:"amount"`
	// Description to send with the payout, the recipient may see this e.g., in their bank statement (if supported) or in email receipts we send on your behalf
	Description *string `json:"description,omitempty"`
	// Currency of the destination channel using ISO-4217 currency code
	Currency string `json:"currency"`
	ReceiptNotification *ReceiptNotification `json:"receipt_notification,omitempty"`
	// Object of additional information you may use
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewCreatePayoutRequest instantiates a new CreatePayoutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePayoutRequest(referenceId string, channelCode string, channelProperties DigitalPayoutChannelProperties, amount float32, currency string) *CreatePayoutRequest {
	this := CreatePayoutRequest{}
	this.ReferenceId = referenceId
	this.ChannelCode = channelCode
	this.ChannelProperties = channelProperties
	this.Amount = amount
	this.Currency = currency
	return &this
}

// NewCreatePayoutRequestWithDefaults instantiates a new CreatePayoutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePayoutRequestWithDefaults() *CreatePayoutRequest {
	this := CreatePayoutRequest{}
	return &this
}

// GetReferenceId returns the ReferenceId field value
func (o *CreatePayoutRequest) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *CreatePayoutRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *CreatePayoutRequest) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetChannelCode returns the ChannelCode field value
func (o *CreatePayoutRequest) GetChannelCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *CreatePayoutRequest) GetChannelCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *CreatePayoutRequest) SetChannelCode(v string) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value
func (o *CreatePayoutRequest) GetChannelProperties() DigitalPayoutChannelProperties {
	if o == nil {
		var ret DigitalPayoutChannelProperties
		return ret
	}

	return o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
func (o *CreatePayoutRequest) GetChannelPropertiesOk() (*DigitalPayoutChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelProperties, true
}

// SetChannelProperties sets field value
func (o *CreatePayoutRequest) SetChannelProperties(v DigitalPayoutChannelProperties) {
	o.ChannelProperties = v
}

// GetAmount returns the Amount field value
func (o *CreatePayoutRequest) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreatePayoutRequest) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreatePayoutRequest) SetAmount(v float32) {
	o.Amount = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreatePayoutRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayoutRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreatePayoutRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreatePayoutRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCurrency returns the Currency field value
func (o *CreatePayoutRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreatePayoutRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreatePayoutRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetReceiptNotification returns the ReceiptNotification field value if set, zero value otherwise.
func (o *CreatePayoutRequest) GetReceiptNotification() ReceiptNotification {
	if o == nil || utils.IsNil(o.ReceiptNotification) {
		var ret ReceiptNotification
		return ret
	}
	return *o.ReceiptNotification
}

// GetReceiptNotificationOk returns a tuple with the ReceiptNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayoutRequest) GetReceiptNotificationOk() (*ReceiptNotification, bool) {
	if o == nil || utils.IsNil(o.ReceiptNotification) {
		return nil, false
	}
	return o.ReceiptNotification, true
}

// HasReceiptNotification returns a boolean if a field has been set.
func (o *CreatePayoutRequest) HasReceiptNotification() bool {
	if o != nil && !utils.IsNil(o.ReceiptNotification) {
		return true
	}

	return false
}

// SetReceiptNotification gets a reference to the given ReceiptNotification and assigns it to the ReceiptNotification field.
func (o *CreatePayoutRequest) SetReceiptNotification(v ReceiptNotification) {
	o.ReceiptNotification = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreatePayoutRequest) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePayoutRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreatePayoutRequest) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CreatePayoutRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CreatePayoutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePayoutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reference_id"] = o.ReferenceId
	toSerialize["channel_code"] = o.ChannelCode
	toSerialize["channel_properties"] = o.ChannelProperties
	toSerialize["amount"] = o.Amount
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["currency"] = o.Currency
	if !utils.IsNil(o.ReceiptNotification) {
		toSerialize["receipt_notification"] = o.ReceiptNotification
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableCreatePayoutRequest struct {
	value *CreatePayoutRequest
	isSet bool
}

func (v NullableCreatePayoutRequest) Get() *CreatePayoutRequest {
	return v.value
}

func (v *NullableCreatePayoutRequest) Set(val *CreatePayoutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePayoutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePayoutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePayoutRequest(val *CreatePayoutRequest) *NullableCreatePayoutRequest {
	return &NullableCreatePayoutRequest{value: val, isSet: true}
}

func (v NullableCreatePayoutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePayoutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



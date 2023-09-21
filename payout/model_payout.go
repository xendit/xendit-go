/*
Payout Service

This API allows Xendit to send money from an account to a channel (banks, eWallets, retail outlets) from across regions

API version: 1.0.0
*/


package payout

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
	"time"
)

// checks if the Payout type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Payout{}

// Payout struct for Payout
type Payout struct {
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
	// Xendit-generated unique identifier for each payout
	Id string `json:"id"`
	// The time payout was created on Xendit's system, in ISO 8601 format
	Created time.Time `json:"created"`
	// The time payout was last updated on Xendit's system, in ISO 8601 format
	Updated time.Time `json:"updated"`
	// Xendit Business ID
	BusinessId string `json:"business_id"`
	// Status of payout
	Status string `json:"status"`
	// If the Payout failed, we include a failure code for more details on the failure.
	FailureCode *string `json:"failure_code,omitempty"`
	// Our estimated time on to when your payout is reflected to the destination account
	EstimatedArrivalTime *time.Time `json:"estimated_arrival_time,omitempty"`
}

// NewPayout instantiates a new Payout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayout(referenceId string, channelCode string, channelProperties DigitalPayoutChannelProperties, amount float32, currency string, id string, created time.Time, updated time.Time, businessId string, status string) *Payout {
	this := Payout{}
	this.ReferenceId = referenceId
	this.ChannelCode = channelCode
	this.ChannelProperties = channelProperties
	this.Amount = amount
	this.Currency = currency
	this.Id = id
	this.Created = created
	this.Updated = updated
	this.BusinessId = businessId
	this.Status = status
	return &this
}

// NewPayoutWithDefaults instantiates a new Payout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutWithDefaults() *Payout {
	this := Payout{}
	return &this
}

// GetReferenceId returns the ReferenceId field value
func (o *Payout) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *Payout) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *Payout) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetChannelCode returns the ChannelCode field value
func (o *Payout) GetChannelCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *Payout) GetChannelCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *Payout) SetChannelCode(v string) {
	o.ChannelCode = v
}

// GetChannelProperties returns the ChannelProperties field value
func (o *Payout) GetChannelProperties() DigitalPayoutChannelProperties {
	if o == nil {
		var ret DigitalPayoutChannelProperties
		return ret
	}

	return o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
func (o *Payout) GetChannelPropertiesOk() (*DigitalPayoutChannelProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelProperties, true
}

// SetChannelProperties sets field value
func (o *Payout) SetChannelProperties(v DigitalPayoutChannelProperties) {
	o.ChannelProperties = v
}

// GetAmount returns the Amount field value
func (o *Payout) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Payout) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Payout) SetAmount(v float32) {
	o.Amount = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Payout) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payout) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Payout) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Payout) SetDescription(v string) {
	o.Description = &v
}

// GetCurrency returns the Currency field value
func (o *Payout) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Payout) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Payout) SetCurrency(v string) {
	o.Currency = v
}

// GetReceiptNotification returns the ReceiptNotification field value if set, zero value otherwise.
func (o *Payout) GetReceiptNotification() ReceiptNotification {
	if o == nil || utils.IsNil(o.ReceiptNotification) {
		var ret ReceiptNotification
		return ret
	}
	return *o.ReceiptNotification
}

// GetReceiptNotificationOk returns a tuple with the ReceiptNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payout) GetReceiptNotificationOk() (*ReceiptNotification, bool) {
	if o == nil || utils.IsNil(o.ReceiptNotification) {
		return nil, false
	}
	return o.ReceiptNotification, true
}

// HasReceiptNotification returns a boolean if a field has been set.
func (o *Payout) HasReceiptNotification() bool {
	if o != nil && !utils.IsNil(o.ReceiptNotification) {
		return true
	}

	return false
}

// SetReceiptNotification gets a reference to the given ReceiptNotification and assigns it to the ReceiptNotification field.
func (o *Payout) SetReceiptNotification(v ReceiptNotification) {
	o.ReceiptNotification = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Payout) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payout) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Payout) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Payout) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetId returns the Id field value
func (o *Payout) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Payout) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Payout) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Payout) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Payout) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Payout) SetCreated(v time.Time) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *Payout) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *Payout) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *Payout) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetBusinessId returns the BusinessId field value
func (o *Payout) GetBusinessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *Payout) GetBusinessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *Payout) SetBusinessId(v string) {
	o.BusinessId = v
}

// GetStatus returns the Status field value
func (o *Payout) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Payout) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Payout) SetStatus(v string) {
	o.Status = v
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise.
func (o *Payout) GetFailureCode() string {
	if o == nil || utils.IsNil(o.FailureCode) {
		var ret string
		return ret
	}
	return *o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payout) GetFailureCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FailureCode) {
		return nil, false
	}
	return o.FailureCode, true
}

// HasFailureCode returns a boolean if a field has been set.
func (o *Payout) HasFailureCode() bool {
	if o != nil && !utils.IsNil(o.FailureCode) {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given string and assigns it to the FailureCode field.
func (o *Payout) SetFailureCode(v string) {
	o.FailureCode = &v
}

// GetEstimatedArrivalTime returns the EstimatedArrivalTime field value if set, zero value otherwise.
func (o *Payout) GetEstimatedArrivalTime() time.Time {
	if o == nil || utils.IsNil(o.EstimatedArrivalTime) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedArrivalTime
}

// GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payout) GetEstimatedArrivalTimeOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EstimatedArrivalTime) {
		return nil, false
	}
	return o.EstimatedArrivalTime, true
}

// HasEstimatedArrivalTime returns a boolean if a field has been set.
func (o *Payout) HasEstimatedArrivalTime() bool {
	if o != nil && !utils.IsNil(o.EstimatedArrivalTime) {
		return true
	}

	return false
}

// SetEstimatedArrivalTime gets a reference to the given time.Time and assigns it to the EstimatedArrivalTime field.
func (o *Payout) SetEstimatedArrivalTime(v time.Time) {
	o.EstimatedArrivalTime = &v
}

func (o Payout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payout) ToMap() (map[string]interface{}, error) {
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
	toSerialize["id"] = o.Id
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	toSerialize["business_id"] = o.BusinessId
	toSerialize["status"] = o.Status
	if !utils.IsNil(o.FailureCode) {
		toSerialize["failure_code"] = o.FailureCode
	}
	if !utils.IsNil(o.EstimatedArrivalTime) {
		toSerialize["estimated_arrival_time"] = o.EstimatedArrivalTime
	}
	return toSerialize, nil
}

type NullablePayout struct {
	value *Payout
	isSet bool
}

func (v NullablePayout) Get() *Payout {
	return v.value
}

func (v *NullablePayout) Set(val *Payout) {
	v.value = val
	v.isSet = true
}

func (v NullablePayout) IsSet() bool {
	return v.isSet
}

func (v *NullablePayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayout(val *Payout) *NullablePayout {
	return &NullablePayout{value: val, isSet: true}
}

func (v NullablePayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



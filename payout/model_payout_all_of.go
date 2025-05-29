/*
Payout Service

This API allows Xendit to send money from an account to a channel (banks, eWallets, retail outlets) from across regions

API version: 1.0.0
*/


package payout

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
	"time"
)

// checks if the PayoutAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PayoutAllOf{}

// PayoutAllOf struct for PayoutAllOf
type PayoutAllOf struct {
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

// NewPayoutAllOf instantiates a new PayoutAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutAllOf(id string, created time.Time, updated time.Time, businessId string, status string) *PayoutAllOf {
	this := PayoutAllOf{}
	this.Id = id
	this.Created = created
	this.Updated = updated
	this.BusinessId = businessId
	this.Status = status
	return &this
}

// NewPayoutAllOfWithDefaults instantiates a new PayoutAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutAllOfWithDefaults() *PayoutAllOf {
	this := PayoutAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *PayoutAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PayoutAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PayoutAllOf) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *PayoutAllOf) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *PayoutAllOf) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *PayoutAllOf) SetCreated(v time.Time) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *PayoutAllOf) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *PayoutAllOf) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *PayoutAllOf) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetBusinessId returns the BusinessId field value
func (o *PayoutAllOf) GetBusinessId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value
// and a boolean to check if the value has been set.
func (o *PayoutAllOf) GetBusinessIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessId, true
}

// SetBusinessId sets field value
func (o *PayoutAllOf) SetBusinessId(v string) {
	o.BusinessId = v
}

// GetStatus returns the Status field value
func (o *PayoutAllOf) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PayoutAllOf) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PayoutAllOf) SetStatus(v string) {
	o.Status = v
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise.
func (o *PayoutAllOf) GetFailureCode() string {
	if o == nil || utils.IsNil(o.FailureCode) {
		var ret string
		return ret
	}
	return *o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutAllOf) GetFailureCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FailureCode) {
		return nil, false
	}
	return o.FailureCode, true
}

// HasFailureCode returns a boolean if a field has been set.
func (o *PayoutAllOf) HasFailureCode() bool {
	if o != nil && !utils.IsNil(o.FailureCode) {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given string and assigns it to the FailureCode field.
func (o *PayoutAllOf) SetFailureCode(v string) {
	o.FailureCode = &v
}

// GetEstimatedArrivalTime returns the EstimatedArrivalTime field value if set, zero value otherwise.
func (o *PayoutAllOf) GetEstimatedArrivalTime() time.Time {
	if o == nil || utils.IsNil(o.EstimatedArrivalTime) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedArrivalTime
}

// GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutAllOf) GetEstimatedArrivalTimeOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EstimatedArrivalTime) {
		return nil, false
	}
	return o.EstimatedArrivalTime, true
}

// HasEstimatedArrivalTime returns a boolean if a field has been set.
func (o *PayoutAllOf) HasEstimatedArrivalTime() bool {
	if o != nil && !utils.IsNil(o.EstimatedArrivalTime) {
		return true
	}

	return false
}

// SetEstimatedArrivalTime gets a reference to the given time.Time and assigns it to the EstimatedArrivalTime field.
func (o *PayoutAllOf) SetEstimatedArrivalTime(v time.Time) {
	o.EstimatedArrivalTime = &v
}

func (o PayoutAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayoutAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	toSerialize["business_id"] = o.BusinessId
	toSerialize["status"] = o.Status
    if o.Status != "SUCCEEDED" && o.Status != "FAILED" && o.Status != "ACCEPTED" && o.Status != "REQUESTED" && o.Status != "LOCKED" && o.Status != "CANCELLED" && o.Status != "REVERSED" {
        toSerialize["status"] = nil
        return toSerialize, utils.NewError("invalid value for Status when marshalling to JSON, must be one of SUCCEEDED, FAILED, ACCEPTED, REQUESTED, LOCKED, CANCELLED, REVERSED")
    }
	if !utils.IsNil(o.FailureCode) {
		toSerialize["failure_code"] = o.FailureCode
	}
	if !utils.IsNil(o.EstimatedArrivalTime) {
		toSerialize["estimated_arrival_time"] = o.EstimatedArrivalTime
	}
	return toSerialize, nil
}

type NullablePayoutAllOf struct {
	value *PayoutAllOf
	isSet bool
}

func (v NullablePayoutAllOf) Get() *PayoutAllOf {
	return v.value
}

func (v *NullablePayoutAllOf) Set(val *PayoutAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutAllOf(val *PayoutAllOf) *NullablePayoutAllOf {
	return &NullablePayoutAllOf{value: val, isSet: true}
}

func (v NullablePayoutAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



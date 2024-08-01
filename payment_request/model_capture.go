/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the Capture type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Capture{}

// Capture struct for Capture
type Capture struct {
	Id string `json:"id"`
	PaymentRequestId string `json:"payment_request_id"`
	PaymentId string `json:"payment_id"`
	ReferenceId string `json:"reference_id"`
	Currency string `json:"currency"`
	AuthorizedAmount float64 `json:"authorized_amount"`
	CapturedAmount float64 `json:"captured_amount"`
	Status string `json:"status"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	FailureCode NullableString `json:"failure_code"`
	CustomerId NullableString `json:"customer_id"`
	Metadata map[string]interface{} `json:"metadata"`
	ChannelProperties map[string]interface{} `json:"channel_properties"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

// NewCapture instantiates a new Capture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapture(id string, paymentRequestId string, paymentId string, referenceId string, currency string, authorizedAmount float64, capturedAmount float64, status string, paymentMethod PaymentMethod, failureCode NullableString, customerId NullableString, metadata map[string]interface{}, channelProperties map[string]interface{}, created string, updated string) *Capture {
	this := Capture{}
	this.Id = id
	this.PaymentRequestId = paymentRequestId
	this.PaymentId = paymentId
	this.ReferenceId = referenceId
	this.Currency = currency
	this.AuthorizedAmount = authorizedAmount
	this.CapturedAmount = capturedAmount
	this.Status = status
	this.PaymentMethod = paymentMethod
	this.FailureCode = failureCode
	this.CustomerId = customerId
	this.Metadata = metadata
	this.ChannelProperties = channelProperties
	this.Created = created
	this.Updated = updated
	return &this
}

// NewCaptureWithDefaults instantiates a new Capture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureWithDefaults() *Capture {
	this := Capture{}
	return &this
}

// GetId returns the Id field value
func (o *Capture) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Capture) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Capture) SetId(v string) {
	o.Id = v
}

// GetPaymentRequestId returns the PaymentRequestId field value
func (o *Capture) GetPaymentRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentRequestId
}

// GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field value
// and a boolean to check if the value has been set.
func (o *Capture) GetPaymentRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentRequestId, true
}

// SetPaymentRequestId sets field value
func (o *Capture) SetPaymentRequestId(v string) {
	o.PaymentRequestId = v
}

// GetPaymentId returns the PaymentId field value
func (o *Capture) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *Capture) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *Capture) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetReferenceId returns the ReferenceId field value
func (o *Capture) GetReferenceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *Capture) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *Capture) SetReferenceId(v string) {
	o.ReferenceId = v
}

// GetCurrency returns the Currency field value
func (o *Capture) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Capture) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Capture) SetCurrency(v string) {
	o.Currency = v
}

// GetAuthorizedAmount returns the AuthorizedAmount field value
func (o *Capture) GetAuthorizedAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AuthorizedAmount
}

// GetAuthorizedAmountOk returns a tuple with the AuthorizedAmount field value
// and a boolean to check if the value has been set.
func (o *Capture) GetAuthorizedAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizedAmount, true
}

// SetAuthorizedAmount sets field value
func (o *Capture) SetAuthorizedAmount(v float64) {
	o.AuthorizedAmount = v
}

// GetCapturedAmount returns the CapturedAmount field value
func (o *Capture) GetCapturedAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CapturedAmount
}

// GetCapturedAmountOk returns a tuple with the CapturedAmount field value
// and a boolean to check if the value has been set.
func (o *Capture) GetCapturedAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedAmount, true
}

// SetCapturedAmount sets field value
func (o *Capture) SetCapturedAmount(v float64) {
	o.CapturedAmount = v
}

// GetStatus returns the Status field value
func (o *Capture) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Capture) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Capture) SetStatus(v string) {
	o.Status = v
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *Capture) GetPaymentMethod() PaymentMethod {
	if o == nil {
		var ret PaymentMethod
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *Capture) GetPaymentMethodOk() (*PaymentMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *Capture) SetPaymentMethod(v PaymentMethod) {
	o.PaymentMethod = v
}

// GetFailureCode returns the FailureCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Capture) GetFailureCode() string {
	if o == nil || o.FailureCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.FailureCode.Get()
}

// GetFailureCodeOk returns a tuple with the FailureCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Capture) GetFailureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureCode.Get(), o.FailureCode.IsSet()
}

// SetFailureCode sets field value
func (o *Capture) SetFailureCode(v string) {
	o.FailureCode.Set(&v)
}

// GetCustomerId returns the CustomerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Capture) GetCustomerId() string {
	if o == nil || o.CustomerId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Capture) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// SetCustomerId sets field value
func (o *Capture) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *Capture) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Capture) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Capture) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetChannelProperties returns the ChannelProperties field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *Capture) GetChannelProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Capture) GetChannelPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return map[string]interface{}{}, false
	}
	return o.ChannelProperties, true
}

// SetChannelProperties sets field value
func (o *Capture) SetChannelProperties(v map[string]interface{}) {
	o.ChannelProperties = v
}

// GetCreated returns the Created field value
func (o *Capture) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Capture) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Capture) SetCreated(v string) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *Capture) GetUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *Capture) GetUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *Capture) SetUpdated(v string) {
	o.Updated = v
}

func (o Capture) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Capture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["payment_request_id"] = o.PaymentRequestId
	toSerialize["payment_id"] = o.PaymentId
	toSerialize["reference_id"] = o.ReferenceId
	toSerialize["currency"] = o.Currency
	toSerialize["authorized_amount"] = o.AuthorizedAmount
	toSerialize["captured_amount"] = o.CapturedAmount
	toSerialize["status"] = o.Status
    if o.Status != "SUCCEEDED" && o.Status != "FAILED" {
        toSerialize["status"] = nil
        return toSerialize, utils.NewError("invalid value for Status when marshalling to JSON, must be one of SUCCEEDED, FAILED")
    }
	toSerialize["payment_method"] = o.PaymentMethod
	toSerialize["failure_code"] = o.FailureCode.Get()

	toSerialize["customer_id"] = o.CustomerId.Get()

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
    }
	if o.ChannelProperties != nil {
		toSerialize["channel_properties"] = o.ChannelProperties
    }
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	return toSerialize, nil
}

type NullableCapture struct {
	value *Capture
	isSet bool
}

func (v NullableCapture) Get() *Capture {
	return v.value
}

func (v *NullableCapture) Set(val *Capture) {
	v.value = val
	v.isSet = true
}

func (v NullableCapture) IsSet() bool {
	return v.isSet
}

func (v *NullableCapture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapture(val *Capture) *NullableCapture {
	return &NullableCapture{value: val, isSet: true}
}

func (v NullableCapture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



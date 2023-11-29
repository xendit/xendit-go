/*
Refund Service

This API is used for the unified refund service

API version: 1.3.4
*/


package refund

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the RefundCallbackData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RefundCallbackData{}

// RefundCallbackData struct for RefundCallbackData
type RefundCallbackData struct {
	Id string `json:"id"`
	PaymentRequestId string `json:"payment_request_id"`
	InvoiceId NullableString `json:"invoice_id,omitempty"`
	PaymentMethodType string `json:"payment_method_type"`
	Amount float64 `json:"amount"`
	ChannelCode string `json:"channel_code"`
	Status string `json:"status"`
	Reason string `json:"reason"`
	Country string `json:"country"`
	Currency string `json:"currency"`
	ReferenceId NullableString `json:"reference_id,omitempty"`
	FailureCode NullableString `json:"failure_code,omitempty"`
	RefundFeeAmount NullableFloat64 `json:"refund_fee_amount,omitempty"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewRefundCallbackData instantiates a new RefundCallbackData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefundCallbackData(id string, paymentRequestId string, paymentMethodType string, amount float64, channelCode string, status string, reason string, country string, currency string, created string, updated string) *RefundCallbackData {
	this := RefundCallbackData{}
	this.Id = id
	this.PaymentRequestId = paymentRequestId
	this.PaymentMethodType = paymentMethodType
	this.Amount = amount
	this.ChannelCode = channelCode
	this.Status = status
	this.Reason = reason
	this.Country = country
	this.Currency = currency
	this.Created = created
	this.Updated = updated
	return &this
}

// NewRefundCallbackDataWithDefaults instantiates a new RefundCallbackData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundCallbackDataWithDefaults() *RefundCallbackData {
	this := RefundCallbackData{}
	return &this
}

// GetId returns the Id field value
func (o *RefundCallbackData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RefundCallbackData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RefundCallbackData) SetId(v string) {
	o.Id = v
}

// GetPaymentRequestId returns the PaymentRequestId field value
func (o *RefundCallbackData) GetPaymentRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentRequestId
}

// GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field value
// and a boolean to check if the value has been set.
func (o *RefundCallbackData) GetPaymentRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentRequestId, true
}

// SetPaymentRequestId sets field value
func (o *RefundCallbackData) SetPaymentRequestId(v string) {
	o.PaymentRequestId = v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefundCallbackData) GetInvoiceId() string {
	if o == nil || utils.IsNil(o.InvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefundCallbackData) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *RefundCallbackData) HasInvoiceId() bool {
	if o != nil && o.InvoiceId.IsSet() {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given NullableString and assigns it to the InvoiceId field.
func (o *RefundCallbackData) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}
// SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil
func (o *RefundCallbackData) SetInvoiceIdNil() {
	o.InvoiceId.Set(nil)
}

// UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
func (o *RefundCallbackData) UnsetInvoiceId() {
	o.InvoiceId.Unset()
}

// GetPaymentMethodType returns the PaymentMethodType field value
func (o *RefundCallbackData) GetPaymentMethodType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value
// and a boolean to check if the value has been set.
func (o *RefundCallbackData) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethodType, true
}

// SetPaymentMethodType sets field value
func (o *RefundCallbackData) SetPaymentMethodType(v string) {
	o.PaymentMethodType = v
}

// GetAmount returns the Amount field value
func (o *RefundCallbackData) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *RefundCallbackData) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *RefundCallbackData) SetAmount(v float64) {
	o.Amount = v
}

// GetChannelCode returns the ChannelCode field value
func (o *RefundCallbackData) GetChannelCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value
// and a boolean to check if the value has been set.
func (o *RefundCallbackData) GetChannelCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelCode, true
}

// SetChannelCode sets field value
func (o *RefundCallbackData) SetChannelCode(v string) {
	o.ChannelCode = v
}

// GetStatus returns the Status field value
func (o *RefundCallbackData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RefundCallbackData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RefundCallbackData) SetStatus(v string) {
	o.Status = v
}

// GetReason returns the Reason field value
func (o *RefundCallbackData) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *RefundCallbackData) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *RefundCallbackData) SetReason(v string) {
	o.Reason = v
}

// GetCountry returns the Country field value
func (o *RefundCallbackData) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *RefundCallbackData) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *RefundCallbackData) SetCountry(v string) {
	o.Country = v
}

// GetCurrency returns the Currency field value
func (o *RefundCallbackData) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *RefundCallbackData) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *RefundCallbackData) SetCurrency(v string) {
	o.Currency = v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefundCallbackData) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId.Get()) {
		var ret string
		return ret
	}
	return *o.ReferenceId.Get()
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefundCallbackData) GetReferenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceId.Get(), o.ReferenceId.IsSet()
}

// HasReferenceId returns a boolean if a field has been set.
func (o *RefundCallbackData) HasReferenceId() bool {
	if o != nil && o.ReferenceId.IsSet() {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given NullableString and assigns it to the ReferenceId field.
func (o *RefundCallbackData) SetReferenceId(v string) {
	o.ReferenceId.Set(&v)
}
// SetReferenceIdNil sets the value for ReferenceId to be an explicit nil
func (o *RefundCallbackData) SetReferenceIdNil() {
	o.ReferenceId.Set(nil)
}

// UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
func (o *RefundCallbackData) UnsetReferenceId() {
	o.ReferenceId.Unset()
}

// GetFailureCode returns the FailureCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefundCallbackData) GetFailureCode() string {
	if o == nil || utils.IsNil(o.FailureCode.Get()) {
		var ret string
		return ret
	}
	return *o.FailureCode.Get()
}

// GetFailureCodeOk returns a tuple with the FailureCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefundCallbackData) GetFailureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureCode.Get(), o.FailureCode.IsSet()
}

// HasFailureCode returns a boolean if a field has been set.
func (o *RefundCallbackData) HasFailureCode() bool {
	if o != nil && o.FailureCode.IsSet() {
		return true
	}

	return false
}

// SetFailureCode gets a reference to the given NullableString and assigns it to the FailureCode field.
func (o *RefundCallbackData) SetFailureCode(v string) {
	o.FailureCode.Set(&v)
}
// SetFailureCodeNil sets the value for FailureCode to be an explicit nil
func (o *RefundCallbackData) SetFailureCodeNil() {
	o.FailureCode.Set(nil)
}

// UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
func (o *RefundCallbackData) UnsetFailureCode() {
	o.FailureCode.Unset()
}

// GetRefundFeeAmount returns the RefundFeeAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefundCallbackData) GetRefundFeeAmount() float64 {
	if o == nil || utils.IsNil(o.RefundFeeAmount.Get()) {
		var ret float64
		return ret
	}
	return *o.RefundFeeAmount.Get()
}

// GetRefundFeeAmountOk returns a tuple with the RefundFeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefundCallbackData) GetRefundFeeAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefundFeeAmount.Get(), o.RefundFeeAmount.IsSet()
}

// HasRefundFeeAmount returns a boolean if a field has been set.
func (o *RefundCallbackData) HasRefundFeeAmount() bool {
	if o != nil && o.RefundFeeAmount.IsSet() {
		return true
	}

	return false
}

// SetRefundFeeAmount gets a reference to the given NullableFloat64 and assigns it to the RefundFeeAmount field.
func (o *RefundCallbackData) SetRefundFeeAmount(v float64) {
	o.RefundFeeAmount.Set(&v)
}
// SetRefundFeeAmountNil sets the value for RefundFeeAmount to be an explicit nil
func (o *RefundCallbackData) SetRefundFeeAmountNil() {
	o.RefundFeeAmount.Set(nil)
}

// UnsetRefundFeeAmount ensures that no value is present for RefundFeeAmount, not even an explicit nil
func (o *RefundCallbackData) UnsetRefundFeeAmount() {
	o.RefundFeeAmount.Unset()
}

// GetCreated returns the Created field value
func (o *RefundCallbackData) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RefundCallbackData) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RefundCallbackData) SetCreated(v string) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *RefundCallbackData) GetUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *RefundCallbackData) GetUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *RefundCallbackData) SetUpdated(v string) {
	o.Updated = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefundCallbackData) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefundCallbackData) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RefundCallbackData) HasMetadata() bool {
	if o != nil && utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RefundCallbackData) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o RefundCallbackData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefundCallbackData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["payment_request_id"] = o.PaymentRequestId
	if o.InvoiceId.IsSet() {
		toSerialize["invoice_id"] = o.InvoiceId.Get()
    }
	toSerialize["payment_method_type"] = o.PaymentMethodType
	toSerialize["amount"] = o.Amount
	toSerialize["channel_code"] = o.ChannelCode
	toSerialize["status"] = o.Status
	toSerialize["reason"] = o.Reason
	toSerialize["country"] = o.Country
	toSerialize["currency"] = o.Currency
	if o.ReferenceId.IsSet() {
		toSerialize["reference_id"] = o.ReferenceId.Get()
    }
	if o.FailureCode.IsSet() {
		toSerialize["failure_code"] = o.FailureCode.Get()
    }
	if o.RefundFeeAmount.IsSet() {
		toSerialize["refund_fee_amount"] = o.RefundFeeAmount.Get()
    }
	toSerialize["created"] = o.Created
	toSerialize["updated"] = o.Updated
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
    }
	return toSerialize, nil
}

type NullableRefundCallbackData struct {
	value *RefundCallbackData
	isSet bool
}

func (v NullableRefundCallbackData) Get() *RefundCallbackData {
	return v.value
}

func (v *NullableRefundCallbackData) Set(val *RefundCallbackData) {
	v.value = val
	v.isSet = true
}

func (v NullableRefundCallbackData) IsSet() bool {
	return v.isSet
}

func (v *NullableRefundCallbackData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefundCallbackData(val *RefundCallbackData) *NullableRefundCallbackData {
	return &NullableRefundCallbackData{value: val, isSet: true}
}

func (v NullableRefundCallbackData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefundCallbackData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



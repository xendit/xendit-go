# Refund


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Id** | Pointer to **string** |  |  |  |
| **PaymentRequestId** | Pointer to **string** |  |  |  |
| **Amount** | Pointer to **float64** |  |  |  |
| **ChannelCode** | Pointer to **string** |  |  |  |
| **Country** | Pointer to **string** |  |  |  |
| **Currency** | Pointer to **string** |  |  |  |
| **ReferenceId** | Pointer to **NullableString** |  |  |  |
| **FailureCode** | Pointer to **NullableString** |  |  |  |
| **RefundFeeAmount** | Pointer to **NullableFloat64** |  |  |  |
| **Created** | Pointer to **string** |  |  |  |
| **Updated** | Pointer to **string** |  |  |  |
| **Metadata** | Pointer to **map[string]interface{}** |  |  |  |

## Methods

### NewRefund

`func NewRefund() *Refund`

NewRefund instantiates a new Refund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundWithDefaults

`func NewRefundWithDefaults() *Refund`

NewRefundWithDefaults instantiates a new Refund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Refund) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Refund) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Refund) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Refund) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaymentRequestId

`func (o *Refund) GetPaymentRequestId() string`

GetPaymentRequestId returns the PaymentRequestId field if non-nil, zero value otherwise.

### GetPaymentRequestIdOk

`func (o *Refund) GetPaymentRequestIdOk() (*string, bool)`

GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestId

`func (o *Refund) SetPaymentRequestId(v string)`

SetPaymentRequestId sets PaymentRequestId field to given value.

### HasPaymentRequestId

`func (o *Refund) HasPaymentRequestId() bool`

HasPaymentRequestId returns a boolean if a field has been set.

### GetAmount

`func (o *Refund) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Refund) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Refund) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Refund) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChannelCode

`func (o *Refund) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *Refund) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *Refund) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *Refund) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### GetCountry

`func (o *Refund) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Refund) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Refund) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Refund) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCurrency

`func (o *Refund) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Refund) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Refund) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Refund) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetReferenceId

`func (o *Refund) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Refund) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Refund) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Refund) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *Refund) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *Refund) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetFailureCode

`func (o *Refund) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *Refund) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *Refund) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *Refund) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *Refund) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *Refund) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetRefundFeeAmount

`func (o *Refund) GetRefundFeeAmount() float64`

GetRefundFeeAmount returns the RefundFeeAmount field if non-nil, zero value otherwise.

### GetRefundFeeAmountOk

`func (o *Refund) GetRefundFeeAmountOk() (*float64, bool)`

GetRefundFeeAmountOk returns a tuple with the RefundFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundFeeAmount

`func (o *Refund) SetRefundFeeAmount(v float64)`

SetRefundFeeAmount sets RefundFeeAmount field to given value.

### HasRefundFeeAmount

`func (o *Refund) HasRefundFeeAmount() bool`

HasRefundFeeAmount returns a boolean if a field has been set.

### SetRefundFeeAmountNil

`func (o *Refund) SetRefundFeeAmountNil(b bool)`

 SetRefundFeeAmountNil sets the value for RefundFeeAmount to be an explicit nil

### UnsetRefundFeeAmount
`func (o *Refund) UnsetRefundFeeAmount()`

UnsetRefundFeeAmount ensures that no value is present for RefundFeeAmount, not even an explicit nil
### GetCreated

`func (o *Refund) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Refund) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Refund) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Refund) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Refund) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Refund) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Refund) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Refund) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetMetadata

`func (o *Refund) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Refund) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Refund) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Refund) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Refund) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Refund) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to README]](../../README.md)



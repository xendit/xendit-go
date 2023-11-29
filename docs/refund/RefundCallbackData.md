# RefundCallbackData


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Id** | **string** | ☑️ |  |  |
| **PaymentRequestId** | **string** | ☑️ |  |  |
| **InvoiceId** | Pointer to **NullableString** |  |  |  |
| **PaymentMethodType** | **string** | ☑️ |  |  |
| **Amount** | **float64** | ☑️ |  |  |
| **ChannelCode** | **string** | ☑️ |  |  |
| **Status** | **string** | ☑️ |  |  |
| **Reason** | **string** | ☑️ |  |  |
| **Country** | **string** | ☑️ |  |  |
| **Currency** | **string** | ☑️ |  |  |
| **ReferenceId** | Pointer to **NullableString** |  |  |  |
| **FailureCode** | Pointer to **NullableString** |  |  |  |
| **RefundFeeAmount** | Pointer to **NullableFloat64** |  |  |  |
| **Created** | **string** | ☑️ |  |  |
| **Updated** | **string** | ☑️ |  |  |
| **Metadata** | Pointer to **map[string]interface{}** |  |  |  |

## Methods

### NewRefundCallbackData

`func NewRefundCallbackData(id string, paymentRequestId string, paymentMethodType string, amount float64, channelCode string, status string, reason string, country string, currency string, created string, updated string, ) *RefundCallbackData`

NewRefundCallbackData instantiates a new RefundCallbackData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundCallbackDataWithDefaults

`func NewRefundCallbackDataWithDefaults() *RefundCallbackData`

NewRefundCallbackDataWithDefaults instantiates a new RefundCallbackData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RefundCallbackData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RefundCallbackData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RefundCallbackData) SetId(v string)`

SetId sets Id field to given value.


### GetPaymentRequestId

`func (o *RefundCallbackData) GetPaymentRequestId() string`

GetPaymentRequestId returns the PaymentRequestId field if non-nil, zero value otherwise.

### GetPaymentRequestIdOk

`func (o *RefundCallbackData) GetPaymentRequestIdOk() (*string, bool)`

GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestId

`func (o *RefundCallbackData) SetPaymentRequestId(v string)`

SetPaymentRequestId sets PaymentRequestId field to given value.


### GetInvoiceId

`func (o *RefundCallbackData) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *RefundCallbackData) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *RefundCallbackData) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *RefundCallbackData) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *RefundCallbackData) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *RefundCallbackData) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetPaymentMethodType

`func (o *RefundCallbackData) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *RefundCallbackData) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *RefundCallbackData) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.


### GetAmount

`func (o *RefundCallbackData) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RefundCallbackData) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RefundCallbackData) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetChannelCode

`func (o *RefundCallbackData) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *RefundCallbackData) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *RefundCallbackData) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.


### GetStatus

`func (o *RefundCallbackData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RefundCallbackData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RefundCallbackData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *RefundCallbackData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RefundCallbackData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RefundCallbackData) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCountry

`func (o *RefundCallbackData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *RefundCallbackData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *RefundCallbackData) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCurrency

`func (o *RefundCallbackData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RefundCallbackData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RefundCallbackData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReferenceId

`func (o *RefundCallbackData) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *RefundCallbackData) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *RefundCallbackData) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *RefundCallbackData) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### SetReferenceIdNil

`func (o *RefundCallbackData) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *RefundCallbackData) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetFailureCode

`func (o *RefundCallbackData) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *RefundCallbackData) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *RefundCallbackData) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *RefundCallbackData) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *RefundCallbackData) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *RefundCallbackData) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetRefundFeeAmount

`func (o *RefundCallbackData) GetRefundFeeAmount() float64`

GetRefundFeeAmount returns the RefundFeeAmount field if non-nil, zero value otherwise.

### GetRefundFeeAmountOk

`func (o *RefundCallbackData) GetRefundFeeAmountOk() (*float64, bool)`

GetRefundFeeAmountOk returns a tuple with the RefundFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundFeeAmount

`func (o *RefundCallbackData) SetRefundFeeAmount(v float64)`

SetRefundFeeAmount sets RefundFeeAmount field to given value.

### HasRefundFeeAmount

`func (o *RefundCallbackData) HasRefundFeeAmount() bool`

HasRefundFeeAmount returns a boolean if a field has been set.

### SetRefundFeeAmountNil

`func (o *RefundCallbackData) SetRefundFeeAmountNil(b bool)`

 SetRefundFeeAmountNil sets the value for RefundFeeAmount to be an explicit nil

### UnsetRefundFeeAmount
`func (o *RefundCallbackData) UnsetRefundFeeAmount()`

UnsetRefundFeeAmount ensures that no value is present for RefundFeeAmount, not even an explicit nil
### GetCreated

`func (o *RefundCallbackData) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RefundCallbackData) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RefundCallbackData) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *RefundCallbackData) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RefundCallbackData) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RefundCallbackData) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetMetadata

`func (o *RefundCallbackData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RefundCallbackData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RefundCallbackData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RefundCallbackData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *RefundCallbackData) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *RefundCallbackData) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to README]](../../README.md)



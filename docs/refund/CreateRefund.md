# CreateRefund


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **PaymentRequestId** | Pointer to **string** |  |  |  |
| **InvoiceId** | Pointer to **string** |  |  |  |
| **ReferenceId** | Pointer to **string** |  |  |  |
| **Amount** | Pointer to **float64** |  |  |  |
| **Currency** | Pointer to **string** |  |  |  |
| **Reason** | Pointer to **string** |  |  |  |
| **Metadata** | Pointer to **map[string]interface{}** |  |  |  |

## Methods

### NewCreateRefund

`func NewCreateRefund() *CreateRefund`

NewCreateRefund instantiates a new CreateRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRefundWithDefaults

`func NewCreateRefundWithDefaults() *CreateRefund`

NewCreateRefundWithDefaults instantiates a new CreateRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentRequestId

`func (o *CreateRefund) GetPaymentRequestId() string`

GetPaymentRequestId returns the PaymentRequestId field if non-nil, zero value otherwise.

### GetPaymentRequestIdOk

`func (o *CreateRefund) GetPaymentRequestIdOk() (*string, bool)`

GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestId

`func (o *CreateRefund) SetPaymentRequestId(v string)`

SetPaymentRequestId sets PaymentRequestId field to given value.

### HasPaymentRequestId

`func (o *CreateRefund) HasPaymentRequestId() bool`

HasPaymentRequestId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *CreateRefund) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreateRefund) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreateRefund) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *CreateRefund) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetReferenceId

`func (o *CreateRefund) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CreateRefund) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CreateRefund) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *CreateRefund) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetAmount

`func (o *CreateRefund) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateRefund) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateRefund) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateRefund) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateRefund) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateRefund) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateRefund) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateRefund) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetReason

`func (o *CreateRefund) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateRefund) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateRefund) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateRefund) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateRefund) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateRefund) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateRefund) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateRefund) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateRefund) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateRefund) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to README]](../../README.md)



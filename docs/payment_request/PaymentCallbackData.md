# PaymentCallbackData
Represents the actual funds transaction/attempt made to a payment method

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Id** | **string** | ☑️ |  |  |
| **PaymentRequestId** | Pointer to **NullableString** |  |  |  |
| **ReferenceId** | **string** | ☑️ |  |  |
| **CustomerId** | Pointer to **NullableString** |  |  |  |
| **Currency** | **string** | ☑️ |  |  |
| **Amount** | **float64** | ☑️ |  |  |
| **Country** | **string** | ☑️ |  |  |
| **Status** | **string** | ☑️ |  |  |
| **PaymentMethod** | [**PaymentMethod**](PaymentMethod.md) | ☑️ |  |  |
| **ChannelProperties** | Pointer to [**NullablePaymentRequestChannelProperties**](PaymentRequestChannelProperties.md) |  |  |  |
| **PaymentDetail** | Pointer to **map[string]interface{}** |  |  |  |
| **FailureCode** | Pointer to **NullableString** |  |  |  |
| **Created** | **string** | ☑️ |  |  |
| **Updated** | **string** | ☑️ |  |  |
| **Metadata** | Pointer to **map[string]interface{}** |  |  |  |

## Methods

### NewPaymentCallbackData

`func NewPaymentCallbackData(id string, referenceId string, currency string, amount float64, country string, status string, paymentMethod PaymentMethod, created string, updated string, ) *PaymentCallbackData`

NewPaymentCallbackData instantiates a new PaymentCallbackData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCallbackDataWithDefaults

`func NewPaymentCallbackDataWithDefaults() *PaymentCallbackData`

NewPaymentCallbackDataWithDefaults instantiates a new PaymentCallbackData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentCallbackData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentCallbackData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentCallbackData) SetId(v string)`

SetId sets Id field to given value.


### GetPaymentRequestId

`func (o *PaymentCallbackData) GetPaymentRequestId() string`

GetPaymentRequestId returns the PaymentRequestId field if non-nil, zero value otherwise.

### GetPaymentRequestIdOk

`func (o *PaymentCallbackData) GetPaymentRequestIdOk() (*string, bool)`

GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestId

`func (o *PaymentCallbackData) SetPaymentRequestId(v string)`

SetPaymentRequestId sets PaymentRequestId field to given value.

### HasPaymentRequestId

`func (o *PaymentCallbackData) HasPaymentRequestId() bool`

HasPaymentRequestId returns a boolean if a field has been set.

### SetPaymentRequestIdNil

`func (o *PaymentCallbackData) SetPaymentRequestIdNil(b bool)`

 SetPaymentRequestIdNil sets the value for PaymentRequestId to be an explicit nil

### UnsetPaymentRequestId
`func (o *PaymentCallbackData) UnsetPaymentRequestId()`

UnsetPaymentRequestId ensures that no value is present for PaymentRequestId, not even an explicit nil
### GetReferenceId

`func (o *PaymentCallbackData) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PaymentCallbackData) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PaymentCallbackData) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetCustomerId

`func (o *PaymentCallbackData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentCallbackData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentCallbackData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PaymentCallbackData) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *PaymentCallbackData) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *PaymentCallbackData) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCurrency

`func (o *PaymentCallbackData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentCallbackData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentCallbackData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAmount

`func (o *PaymentCallbackData) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentCallbackData) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentCallbackData) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCountry

`func (o *PaymentCallbackData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentCallbackData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentCallbackData) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetStatus

`func (o *PaymentCallbackData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentCallbackData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentCallbackData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPaymentMethod

`func (o *PaymentCallbackData) GetPaymentMethod() PaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentCallbackData) GetPaymentMethodOk() (*PaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentCallbackData) SetPaymentMethod(v PaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetChannelProperties

`func (o *PaymentCallbackData) GetChannelProperties() PaymentRequestChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PaymentCallbackData) GetChannelPropertiesOk() (*PaymentRequestChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PaymentCallbackData) SetChannelProperties(v PaymentRequestChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *PaymentCallbackData) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### SetChannelPropertiesNil

`func (o *PaymentCallbackData) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *PaymentCallbackData) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetPaymentDetail

`func (o *PaymentCallbackData) GetPaymentDetail() map[string]interface{}`

GetPaymentDetail returns the PaymentDetail field if non-nil, zero value otherwise.

### GetPaymentDetailOk

`func (o *PaymentCallbackData) GetPaymentDetailOk() (*map[string]interface{}, bool)`

GetPaymentDetailOk returns a tuple with the PaymentDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDetail

`func (o *PaymentCallbackData) SetPaymentDetail(v map[string]interface{})`

SetPaymentDetail sets PaymentDetail field to given value.

### HasPaymentDetail

`func (o *PaymentCallbackData) HasPaymentDetail() bool`

HasPaymentDetail returns a boolean if a field has been set.

### SetPaymentDetailNil

`func (o *PaymentCallbackData) SetPaymentDetailNil(b bool)`

 SetPaymentDetailNil sets the value for PaymentDetail to be an explicit nil

### UnsetPaymentDetail
`func (o *PaymentCallbackData) UnsetPaymentDetail()`

UnsetPaymentDetail ensures that no value is present for PaymentDetail, not even an explicit nil
### GetFailureCode

`func (o *PaymentCallbackData) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *PaymentCallbackData) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *PaymentCallbackData) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *PaymentCallbackData) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### SetFailureCodeNil

`func (o *PaymentCallbackData) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *PaymentCallbackData) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetCreated

`func (o *PaymentCallbackData) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PaymentCallbackData) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PaymentCallbackData) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *PaymentCallbackData) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PaymentCallbackData) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PaymentCallbackData) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetMetadata

`func (o *PaymentCallbackData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentCallbackData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentCallbackData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentCallbackData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PaymentCallbackData) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PaymentCallbackData) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to README]](../../README.md)



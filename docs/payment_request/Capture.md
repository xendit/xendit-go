# Capture

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Id** | **string** |  |  |
| **PaymentRequestId** | **string** |  |  |
| **PaymentId** | **string** |  |  |
| **ReferenceId** | **string** |  |  |
| **Currency** | **string** |  |  |
| **AuthorizedAmount** | **float64** |  |  |
| **CapturedAmount** | **float64** |  |  |
| **Status** | **string** |  |  |
| **PaymentMethod** | [**PaymentMethod**](PaymentMethod.md) |  |  |
| **FailureCode** | **NullableString** |  |  |
| **CustomerId** | **NullableString** |  |  |
| **Metadata** | **map[string]interface{}** |  |  |
| **ChannelProperties** | **map[string]interface{}** |  |  |
| **Created** | **string** |  |  |
| **Updated** | **string** |  |  |

## Methods

### NewCapture

`func NewCapture(id string, paymentRequestId string, paymentId string, referenceId string, currency string, authorizedAmount float64, capturedAmount float64, status string, paymentMethod PaymentMethod, failureCode NullableString, customerId NullableString, metadata map[string]interface{}, channelProperties map[string]interface{}, created string, updated string, ) *Capture`

NewCapture instantiates a new Capture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureWithDefaults

`func NewCaptureWithDefaults() *Capture`

NewCaptureWithDefaults instantiates a new Capture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Capture) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Capture) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Capture) SetId(v string)`

SetId sets Id field to given value.


### GetPaymentRequestId

`func (o *Capture) GetPaymentRequestId() string`

GetPaymentRequestId returns the PaymentRequestId field if non-nil, zero value otherwise.

### GetPaymentRequestIdOk

`func (o *Capture) GetPaymentRequestIdOk() (*string, bool)`

GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestId

`func (o *Capture) SetPaymentRequestId(v string)`

SetPaymentRequestId sets PaymentRequestId field to given value.


### GetPaymentId

`func (o *Capture) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *Capture) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *Capture) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetReferenceId

`func (o *Capture) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Capture) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Capture) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetCurrency

`func (o *Capture) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Capture) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Capture) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAuthorizedAmount

`func (o *Capture) GetAuthorizedAmount() float64`

GetAuthorizedAmount returns the AuthorizedAmount field if non-nil, zero value otherwise.

### GetAuthorizedAmountOk

`func (o *Capture) GetAuthorizedAmountOk() (*float64, bool)`

GetAuthorizedAmountOk returns a tuple with the AuthorizedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedAmount

`func (o *Capture) SetAuthorizedAmount(v float64)`

SetAuthorizedAmount sets AuthorizedAmount field to given value.


### GetCapturedAmount

`func (o *Capture) GetCapturedAmount() float64`

GetCapturedAmount returns the CapturedAmount field if non-nil, zero value otherwise.

### GetCapturedAmountOk

`func (o *Capture) GetCapturedAmountOk() (*float64, bool)`

GetCapturedAmountOk returns a tuple with the CapturedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAmount

`func (o *Capture) SetCapturedAmount(v float64)`

SetCapturedAmount sets CapturedAmount field to given value.


### GetStatus

`func (o *Capture) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Capture) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Capture) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPaymentMethod

`func (o *Capture) GetPaymentMethod() PaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *Capture) GetPaymentMethodOk() (*PaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *Capture) SetPaymentMethod(v PaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetFailureCode

`func (o *Capture) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *Capture) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *Capture) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.


### SetFailureCodeNil

`func (o *Capture) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *Capture) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetCustomerId

`func (o *Capture) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Capture) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Capture) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *Capture) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *Capture) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetMetadata

`func (o *Capture) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Capture) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Capture) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *Capture) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Capture) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetChannelProperties

`func (o *Capture) GetChannelProperties() map[string]interface{}`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *Capture) GetChannelPropertiesOk() (*map[string]interface{}, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *Capture) SetChannelProperties(v map[string]interface{})`

SetChannelProperties sets ChannelProperties field to given value.


### SetChannelPropertiesNil

`func (o *Capture) SetChannelPropertiesNil(b bool)`

 SetChannelPropertiesNil sets the value for ChannelProperties to be an explicit nil

### UnsetChannelProperties
`func (o *Capture) UnsetChannelProperties()`

UnsetChannelProperties ensures that no value is present for ChannelProperties, not even an explicit nil
### GetCreated

`func (o *Capture) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Capture) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Capture) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *Capture) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Capture) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Capture) SetUpdated(v string)`

SetUpdated sets Updated field to given value.



[[Back to README]](../../README.md)



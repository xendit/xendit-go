# PaymentCallback
Callback for successful or failed payments made via the Payments API

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Event** | **string** | ☑️ | Identifies the event that triggered a notification to the merchant |  |
| **BusinessId** | **string** | ☑️ | business_id |  |
| **Created** | **string** | ☑️ |  |  |
| **Data** | Pointer to [**PaymentCallbackData**](PaymentCallbackData.md) |  |  |  |

## Methods

### NewPaymentCallback

`func NewPaymentCallback(event string, businessId string, created string, ) *PaymentCallback`

NewPaymentCallback instantiates a new PaymentCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentCallbackWithDefaults

`func NewPaymentCallbackWithDefaults() *PaymentCallback`

NewPaymentCallbackWithDefaults instantiates a new PaymentCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *PaymentCallback) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PaymentCallback) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PaymentCallback) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetBusinessId

`func (o *PaymentCallback) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *PaymentCallback) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *PaymentCallback) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.


### GetCreated

`func (o *PaymentCallback) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PaymentCallback) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PaymentCallback) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetData

`func (o *PaymentCallback) GetData() PaymentCallbackData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentCallback) GetDataOk() (*PaymentCallbackData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentCallback) SetData(v PaymentCallbackData)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentCallback) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../../README.md)



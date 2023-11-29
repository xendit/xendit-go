# RefundCallback
Callback for successful or failed Refunds made via the Payments API

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Event** | **string** | ☑️ | Identifies the event that triggered a notification to the merchant |  |
| **BusinessId** | **string** | ☑️ | business_id |  |
| **Created** | **string** | ☑️ |  |  |
| **Data** | Pointer to [**RefundCallbackData**](RefundCallbackData.md) |  |  |  |

## Methods

### NewRefundCallback

`func NewRefundCallback(event string, businessId string, created string, ) *RefundCallback`

NewRefundCallback instantiates a new RefundCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundCallbackWithDefaults

`func NewRefundCallbackWithDefaults() *RefundCallback`

NewRefundCallbackWithDefaults instantiates a new RefundCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *RefundCallback) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RefundCallback) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RefundCallback) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetBusinessId

`func (o *RefundCallback) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *RefundCallback) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *RefundCallback) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.


### GetCreated

`func (o *RefundCallback) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RefundCallback) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RefundCallback) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetData

`func (o *RefundCallback) GetData() RefundCallbackData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RefundCallback) GetDataOk() (*RefundCallbackData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RefundCallback) SetData(v RefundCallbackData)`

SetData sets Data field to given value.

### HasData

`func (o *RefundCallback) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../../README.md)



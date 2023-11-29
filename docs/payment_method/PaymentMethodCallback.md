# PaymentMethodCallback
Callback for active or expired E-Wallet or Direct Debit account linking, Virtual Accounts or QR strings

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Event** | **string** | ☑️ | Identifies the event that triggered a notification to the merchant |  |
| **BusinessId** | **string** | ☑️ | business_id |  |
| **Created** | **string** | ☑️ |  |  |
| **Data** | Pointer to [**PaymentMethod**](PaymentMethod.md) |  |  |  |

## Methods

### NewPaymentMethodCallback

`func NewPaymentMethodCallback(event string, businessId string, created string, ) *PaymentMethodCallback`

NewPaymentMethodCallback instantiates a new PaymentMethodCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCallbackWithDefaults

`func NewPaymentMethodCallbackWithDefaults() *PaymentMethodCallback`

NewPaymentMethodCallbackWithDefaults instantiates a new PaymentMethodCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *PaymentMethodCallback) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PaymentMethodCallback) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PaymentMethodCallback) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetBusinessId

`func (o *PaymentMethodCallback) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *PaymentMethodCallback) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *PaymentMethodCallback) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.


### GetCreated

`func (o *PaymentMethodCallback) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PaymentMethodCallback) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PaymentMethodCallback) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetData

`func (o *PaymentMethodCallback) GetData() PaymentMethod`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentMethodCallback) GetDataOk() (*PaymentMethod, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentMethodCallback) SetData(v PaymentMethod)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentMethodCallback) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../../README.md)



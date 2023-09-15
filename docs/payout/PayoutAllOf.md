# PayoutAllOf

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Id** | **string** | Xendit-generated unique identifier for each payout |  |
| **Created** | **time.Time** | The time payout was created on Xendit&#39;s system, in ISO 8601 format |  |
| **Updated** | **time.Time** | The time payout was last updated on Xendit&#39;s system, in ISO 8601 format |  |
| **BusinessId** | **string** | Xendit Business ID |  |
| **Status** | **string** | Status of payout |  |
| **FailureCode** | Pointer to **string** | If the Payout failed, we include a failure code for more details on the failure. | [optional]  |
| **EstimatedArrivalTime** | Pointer to **time.Time** | Our estimated time on to when your payout is reflected to the destination account | [optional]  |

## Methods

### NewPayoutAllOf

`func NewPayoutAllOf(id string, created time.Time, updated time.Time, businessId string, status string, ) *PayoutAllOf`

NewPayoutAllOf instantiates a new PayoutAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutAllOfWithDefaults

`func NewPayoutAllOfWithDefaults() *PayoutAllOf`

NewPayoutAllOfWithDefaults instantiates a new PayoutAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PayoutAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayoutAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayoutAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *PayoutAllOf) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PayoutAllOf) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PayoutAllOf) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *PayoutAllOf) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PayoutAllOf) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PayoutAllOf) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetBusinessId

`func (o *PayoutAllOf) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *PayoutAllOf) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *PayoutAllOf) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.


### GetStatus

`func (o *PayoutAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFailureCode

`func (o *PayoutAllOf) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *PayoutAllOf) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *PayoutAllOf) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *PayoutAllOf) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetEstimatedArrivalTime

`func (o *PayoutAllOf) GetEstimatedArrivalTime() time.Time`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *PayoutAllOf) GetEstimatedArrivalTimeOk() (*time.Time, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *PayoutAllOf) SetEstimatedArrivalTime(v time.Time)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.

### HasEstimatedArrivalTime

`func (o *PayoutAllOf) HasEstimatedArrivalTime() bool`

HasEstimatedArrivalTime returns a boolean if a field has been set.


[[Back to README]](../../README.md)



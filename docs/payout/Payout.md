# Payout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **string** | A client defined payout identifier | 
**ChannelCode** | **string** | Channel code of selected destination bank or e-wallet | 
**ChannelProperties** | [**DigitalPayoutChannelProperties**](DigitalPayoutChannelProperties.md) |  | 
**Amount** | **float32** | Amount to be sent to the destination account and should be a multiple of the minimum increment for the selected channel | 
**Description** | Pointer to **string** | Description to send with the payout, the recipient may see this e.g., in their bank statement (if supported) or in email receipts we send on your behalf | [optional] 
**Currency** | **string** | Currency of the destination channel using ISO-4217 currency code | 
**ReceiptNotification** | Pointer to [**ReceiptNotification**](ReceiptNotification.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Object of additional information you may use | [optional] 
**Id** | **string** | Xendit-generated unique identifier for each payout | 
**Created** | **time.Time** | The time payout was created on Xendit&#39;s system, in ISO 8601 format | 
**Updated** | **time.Time** | The time payout was last updated on Xendit&#39;s system, in ISO 8601 format | 
**BusinessId** | **string** | Xendit Business ID | 
**Status** | **string** | Status of payout | 
**FailureCode** | Pointer to **string** | If the Payout failed, we include a failure code for more details on the failure. | [optional] 
**EstimatedArrivalTime** | Pointer to **time.Time** | Our estimated time on to when your payout is reflected to the destination account | [optional] 

## Methods

### NewPayout

`func NewPayout(referenceId string, channelCode string, channelProperties DigitalPayoutChannelProperties, amount float32, currency string, id string, created time.Time, updated time.Time, businessId string, status string, ) *Payout`

NewPayout instantiates a new Payout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutWithDefaults

`func NewPayoutWithDefaults() *Payout`

NewPayoutWithDefaults instantiates a new Payout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *Payout) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Payout) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Payout) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetChannelCode

`func (o *Payout) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *Payout) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *Payout) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *Payout) GetChannelProperties() DigitalPayoutChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *Payout) GetChannelPropertiesOk() (*DigitalPayoutChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *Payout) SetChannelProperties(v DigitalPayoutChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### GetAmount

`func (o *Payout) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payout) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payout) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *Payout) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Payout) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Payout) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Payout) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCurrency

`func (o *Payout) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Payout) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Payout) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReceiptNotification

`func (o *Payout) GetReceiptNotification() ReceiptNotification`

GetReceiptNotification returns the ReceiptNotification field if non-nil, zero value otherwise.

### GetReceiptNotificationOk

`func (o *Payout) GetReceiptNotificationOk() (*ReceiptNotification, bool)`

GetReceiptNotificationOk returns a tuple with the ReceiptNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptNotification

`func (o *Payout) SetReceiptNotification(v ReceiptNotification)`

SetReceiptNotification sets ReceiptNotification field to given value.

### HasReceiptNotification

`func (o *Payout) HasReceiptNotification() bool`

HasReceiptNotification returns a boolean if a field has been set.

### GetMetadata

`func (o *Payout) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Payout) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Payout) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Payout) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *Payout) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payout) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payout) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *Payout) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Payout) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Payout) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *Payout) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Payout) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Payout) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetBusinessId

`func (o *Payout) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *Payout) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *Payout) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.


### GetStatus

`func (o *Payout) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payout) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payout) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFailureCode

`func (o *Payout) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *Payout) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *Payout) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *Payout) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetEstimatedArrivalTime

`func (o *Payout) GetEstimatedArrivalTime() time.Time`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *Payout) GetEstimatedArrivalTimeOk() (*time.Time, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *Payout) SetEstimatedArrivalTime(v time.Time)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.

### HasEstimatedArrivalTime

`func (o *Payout) HasEstimatedArrivalTime() bool`

HasEstimatedArrivalTime returns a boolean if a field has been set.


[[Back to README]](../../README.md)



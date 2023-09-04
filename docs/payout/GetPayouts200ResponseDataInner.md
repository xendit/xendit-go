# GetPayouts200ResponseDataInner

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

### NewGetPayouts200ResponseDataInner

`func NewGetPayouts200ResponseDataInner(referenceId string, channelCode string, channelProperties DigitalPayoutChannelProperties, amount float32, currency string, id string, created time.Time, updated time.Time, businessId string, status string, ) *GetPayouts200ResponseDataInner`

NewGetPayouts200ResponseDataInner instantiates a new GetPayouts200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayouts200ResponseDataInnerWithDefaults

`func NewGetPayouts200ResponseDataInnerWithDefaults() *GetPayouts200ResponseDataInner`

NewGetPayouts200ResponseDataInnerWithDefaults instantiates a new GetPayouts200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *GetPayouts200ResponseDataInner) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *GetPayouts200ResponseDataInner) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *GetPayouts200ResponseDataInner) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetChannelCode

`func (o *GetPayouts200ResponseDataInner) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *GetPayouts200ResponseDataInner) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *GetPayouts200ResponseDataInner) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *GetPayouts200ResponseDataInner) GetChannelProperties() DigitalPayoutChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *GetPayouts200ResponseDataInner) GetChannelPropertiesOk() (*DigitalPayoutChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *GetPayouts200ResponseDataInner) SetChannelProperties(v DigitalPayoutChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### GetAmount

`func (o *GetPayouts200ResponseDataInner) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetPayouts200ResponseDataInner) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetPayouts200ResponseDataInner) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *GetPayouts200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetPayouts200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetPayouts200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetPayouts200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCurrency

`func (o *GetPayouts200ResponseDataInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetPayouts200ResponseDataInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetPayouts200ResponseDataInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReceiptNotification

`func (o *GetPayouts200ResponseDataInner) GetReceiptNotification() ReceiptNotification`

GetReceiptNotification returns the ReceiptNotification field if non-nil, zero value otherwise.

### GetReceiptNotificationOk

`func (o *GetPayouts200ResponseDataInner) GetReceiptNotificationOk() (*ReceiptNotification, bool)`

GetReceiptNotificationOk returns a tuple with the ReceiptNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptNotification

`func (o *GetPayouts200ResponseDataInner) SetReceiptNotification(v ReceiptNotification)`

SetReceiptNotification sets ReceiptNotification field to given value.

### HasReceiptNotification

`func (o *GetPayouts200ResponseDataInner) HasReceiptNotification() bool`

HasReceiptNotification returns a boolean if a field has been set.

### GetMetadata

`func (o *GetPayouts200ResponseDataInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetPayouts200ResponseDataInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetPayouts200ResponseDataInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetPayouts200ResponseDataInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *GetPayouts200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPayouts200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPayouts200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *GetPayouts200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetPayouts200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetPayouts200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *GetPayouts200ResponseDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetPayouts200ResponseDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetPayouts200ResponseDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.


### GetBusinessId

`func (o *GetPayouts200ResponseDataInner) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *GetPayouts200ResponseDataInner) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *GetPayouts200ResponseDataInner) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.


### GetStatus

`func (o *GetPayouts200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPayouts200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPayouts200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFailureCode

`func (o *GetPayouts200ResponseDataInner) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *GetPayouts200ResponseDataInner) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *GetPayouts200ResponseDataInner) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.

### HasFailureCode

`func (o *GetPayouts200ResponseDataInner) HasFailureCode() bool`

HasFailureCode returns a boolean if a field has been set.

### GetEstimatedArrivalTime

`func (o *GetPayouts200ResponseDataInner) GetEstimatedArrivalTime() time.Time`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *GetPayouts200ResponseDataInner) GetEstimatedArrivalTimeOk() (*time.Time, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *GetPayouts200ResponseDataInner) SetEstimatedArrivalTime(v time.Time)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.

### HasEstimatedArrivalTime

`func (o *GetPayouts200ResponseDataInner) HasEstimatedArrivalTime() bool`

HasEstimatedArrivalTime returns a boolean if a field has been set.


[[Back to README]](../../README.md)



# CreatePayoutRequest

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

## Methods

### NewCreatePayoutRequest

`func NewCreatePayoutRequest(referenceId string, channelCode string, channelProperties DigitalPayoutChannelProperties, amount float32, currency string, ) *CreatePayoutRequest`

NewCreatePayoutRequest instantiates a new CreatePayoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayoutRequestWithDefaults

`func NewCreatePayoutRequestWithDefaults() *CreatePayoutRequest`

NewCreatePayoutRequestWithDefaults instantiates a new CreatePayoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *CreatePayoutRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *CreatePayoutRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *CreatePayoutRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetChannelCode

`func (o *CreatePayoutRequest) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *CreatePayoutRequest) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *CreatePayoutRequest) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.


### GetChannelProperties

`func (o *CreatePayoutRequest) GetChannelProperties() DigitalPayoutChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *CreatePayoutRequest) GetChannelPropertiesOk() (*DigitalPayoutChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *CreatePayoutRequest) SetChannelProperties(v DigitalPayoutChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.


### GetAmount

`func (o *CreatePayoutRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreatePayoutRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreatePayoutRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *CreatePayoutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePayoutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePayoutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePayoutRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCurrency

`func (o *CreatePayoutRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePayoutRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePayoutRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReceiptNotification

`func (o *CreatePayoutRequest) GetReceiptNotification() ReceiptNotification`

GetReceiptNotification returns the ReceiptNotification field if non-nil, zero value otherwise.

### GetReceiptNotificationOk

`func (o *CreatePayoutRequest) GetReceiptNotificationOk() (*ReceiptNotification, bool)`

GetReceiptNotificationOk returns a tuple with the ReceiptNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptNotification

`func (o *CreatePayoutRequest) SetReceiptNotification(v ReceiptNotification)`

SetReceiptNotification sets ReceiptNotification field to given value.

### HasReceiptNotification

`func (o *CreatePayoutRequest) HasReceiptNotification() bool`

HasReceiptNotification returns a boolean if a field has been set.

### GetMetadata

`func (o *CreatePayoutRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePayoutRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePayoutRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreatePayoutRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to README]](../../README.md)



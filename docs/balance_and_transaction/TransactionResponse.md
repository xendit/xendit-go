# TransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of a transaction. It will have &#x60;txn_&#x60; as prefix | 
**ProductId** | **string** | The product_id of the transaction. Product id will have a different prefix for each product. You can use this id to match the transaction from this API to each product API. | 
**Type** | [**TransactionResponseType**](TransactionResponseType.md) |  | 
**Status** | [**TransactionStatuses**](TransactionStatuses.md) |  | 
**ChannelCategory** | [**ChannelsCategories**](ChannelsCategories.md) |  | 
**ChannelCode** | **NullableString** | The channel of the transaction that is used. See [channel codes](https://docs.xendit.co/xendisburse/channel-codes) for the list of available per channel categories. | 
**AccountIdentifier** | **NullableString** | Account identifier of transaction. The format will be different from each channel. | 
**ReferenceId** | **string** | customer supplied reference/external_id | 
**Currency** | [**Currency**](Currency.md) |  | 
**Amount** | **float32** | The transaction amount. The number of decimal places will be different for each currency according to ISO 4217. | 
**Cashflow** | **string** | Representing whether the transaction is money in or money out For transfer, the transfer out side it will shows up as money out and on transfer in side in will shows up as money-in. Available values are &#x60;MONEY_IN&#x60; for money in and &#x60;MONEY_OUT&#x60; for money out. | 
**SettlementStatus** | Pointer to **NullableString** | The settlement status of the transaction. &#x60;PENDING&#x60; - Transaction amount has not been settled to merchant&#39;s balance. &#x60;SETTLED&#x60; - Transaction has been settled to merchant&#39;s balance | [optional] 
**EstimatedSettlementTime** | Pointer to **NullableTime** | Estimated settlement time will only apply to money-in transactions. For money-out transaction, the value will be &#x60;NULL&#x60;. Estimated settlement time in which transaction amount will be settled to merchant&#39;s balance. | [optional] 
**BusinessId** | **string** | The id of business where this transaction belong to | 
**Fee** | [**FeeResponse**](FeeResponse.md) |  | 
**Created** | **time.Time** | Transaction created timestamp (UTC+0) | 
**Updated** | **time.Time** | Transaction updated timestamp (UTC+0) | 

## Methods

### NewTransactionResponse

`func NewTransactionResponse(id string, productId string, type_ TransactionResponseType, status TransactionStatuses, channelCategory ChannelsCategories, channelCode NullableString, accountIdentifier NullableString, referenceId string, currency Currency, amount float32, cashflow string, businessId string, fee FeeResponse, created time.Time, updated time.Time, ) *TransactionResponse`

NewTransactionResponse instantiates a new TransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResponseWithDefaults

`func NewTransactionResponseWithDefaults() *TransactionResponse`

NewTransactionResponseWithDefaults instantiates a new TransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetProductId

`func (o *TransactionResponse) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *TransactionResponse) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *TransactionResponse) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetType

`func (o *TransactionResponse) GetType() TransactionResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionResponse) GetTypeOk() (*TransactionResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionResponse) SetType(v TransactionResponseType)`

SetType sets Type field to given value.


### GetStatus

`func (o *TransactionResponse) GetStatus() TransactionStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionResponse) GetStatusOk() (*TransactionStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionResponse) SetStatus(v TransactionStatuses)`

SetStatus sets Status field to given value.


### GetChannelCategory

`func (o *TransactionResponse) GetChannelCategory() ChannelsCategories`

GetChannelCategory returns the ChannelCategory field if non-nil, zero value otherwise.

### GetChannelCategoryOk

`func (o *TransactionResponse) GetChannelCategoryOk() (*ChannelsCategories, bool)`

GetChannelCategoryOk returns a tuple with the ChannelCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCategory

`func (o *TransactionResponse) SetChannelCategory(v ChannelsCategories)`

SetChannelCategory sets ChannelCategory field to given value.


### GetChannelCode

`func (o *TransactionResponse) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *TransactionResponse) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *TransactionResponse) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.


### SetChannelCodeNil

`func (o *TransactionResponse) SetChannelCodeNil(b bool)`

 SetChannelCodeNil sets the value for ChannelCode to be an explicit nil

### UnsetChannelCode
`func (o *TransactionResponse) UnsetChannelCode()`

UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
### GetAccountIdentifier

`func (o *TransactionResponse) GetAccountIdentifier() string`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *TransactionResponse) GetAccountIdentifierOk() (*string, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *TransactionResponse) SetAccountIdentifier(v string)`

SetAccountIdentifier sets AccountIdentifier field to given value.


### SetAccountIdentifierNil

`func (o *TransactionResponse) SetAccountIdentifierNil(b bool)`

 SetAccountIdentifierNil sets the value for AccountIdentifier to be an explicit nil

### UnsetAccountIdentifier
`func (o *TransactionResponse) UnsetAccountIdentifier()`

UnsetAccountIdentifier ensures that no value is present for AccountIdentifier, not even an explicit nil
### GetReferenceId

`func (o *TransactionResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *TransactionResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *TransactionResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetCurrency

`func (o *TransactionResponse) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionResponse) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionResponse) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetAmount

`func (o *TransactionResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetCashflow

`func (o *TransactionResponse) GetCashflow() string`

GetCashflow returns the Cashflow field if non-nil, zero value otherwise.

### GetCashflowOk

`func (o *TransactionResponse) GetCashflowOk() (*string, bool)`

GetCashflowOk returns a tuple with the Cashflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashflow

`func (o *TransactionResponse) SetCashflow(v string)`

SetCashflow sets Cashflow field to given value.


### GetSettlementStatus

`func (o *TransactionResponse) GetSettlementStatus() string`

GetSettlementStatus returns the SettlementStatus field if non-nil, zero value otherwise.

### GetSettlementStatusOk

`func (o *TransactionResponse) GetSettlementStatusOk() (*string, bool)`

GetSettlementStatusOk returns a tuple with the SettlementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementStatus

`func (o *TransactionResponse) SetSettlementStatus(v string)`

SetSettlementStatus sets SettlementStatus field to given value.

### HasSettlementStatus

`func (o *TransactionResponse) HasSettlementStatus() bool`

HasSettlementStatus returns a boolean if a field has been set.

### SetSettlementStatusNil

`func (o *TransactionResponse) SetSettlementStatusNil(b bool)`

 SetSettlementStatusNil sets the value for SettlementStatus to be an explicit nil

### UnsetSettlementStatus
`func (o *TransactionResponse) UnsetSettlementStatus()`

UnsetSettlementStatus ensures that no value is present for SettlementStatus, not even an explicit nil
### GetEstimatedSettlementTime

`func (o *TransactionResponse) GetEstimatedSettlementTime() time.Time`

GetEstimatedSettlementTime returns the EstimatedSettlementTime field if non-nil, zero value otherwise.

### GetEstimatedSettlementTimeOk

`func (o *TransactionResponse) GetEstimatedSettlementTimeOk() (*time.Time, bool)`

GetEstimatedSettlementTimeOk returns a tuple with the EstimatedSettlementTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSettlementTime

`func (o *TransactionResponse) SetEstimatedSettlementTime(v time.Time)`

SetEstimatedSettlementTime sets EstimatedSettlementTime field to given value.

### HasEstimatedSettlementTime

`func (o *TransactionResponse) HasEstimatedSettlementTime() bool`

HasEstimatedSettlementTime returns a boolean if a field has been set.

### SetEstimatedSettlementTimeNil

`func (o *TransactionResponse) SetEstimatedSettlementTimeNil(b bool)`

 SetEstimatedSettlementTimeNil sets the value for EstimatedSettlementTime to be an explicit nil

### UnsetEstimatedSettlementTime
`func (o *TransactionResponse) UnsetEstimatedSettlementTime()`

UnsetEstimatedSettlementTime ensures that no value is present for EstimatedSettlementTime, not even an explicit nil
### GetBusinessId

`func (o *TransactionResponse) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *TransactionResponse) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *TransactionResponse) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.


### GetFee

`func (o *TransactionResponse) GetFee() FeeResponse`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TransactionResponse) GetFeeOk() (*FeeResponse, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TransactionResponse) SetFee(v FeeResponse)`

SetFee sets Fee field to given value.


### GetCreated

`func (o *TransactionResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TransactionResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TransactionResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *TransactionResponse) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TransactionResponse) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TransactionResponse) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.



[[Back to README]](../../README.md)



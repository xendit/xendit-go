# PaymentMethod

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Id** | **string** |  |  |
| **Type** | [**PaymentMethodType**](PaymentMethodType.md) |  |  |
| **Created** | Pointer to **string** |  | [optional]  |
| **Updated** | Pointer to **string** |  | [optional]  |
| **Description** | Pointer to **NullableString** |  | [optional]  |
| **ReferenceId** | Pointer to **string** |  | [optional]  |
| **Card** | Pointer to [**NullableCard**](Card.md) |  | [optional]  |
| **DirectDebit** | Pointer to [**NullableDirectDebit**](DirectDebit.md) |  | [optional]  |
| **Ewallet** | Pointer to [**NullableEWallet**](EWallet.md) |  | [optional]  |
| **OverTheCounter** | Pointer to [**NullableOverTheCounter**](OverTheCounter.md) |  | [optional]  |
| **VirtualAccount** | Pointer to [**NullableVirtualAccount**](VirtualAccount.md) |  | [optional]  |
| **QrCode** | Pointer to [**NullableQRCode**](QRCode.md) |  | [optional]  |
| **Reusability** | [**PaymentMethodReusability**](PaymentMethodReusability.md) |  |  |
| **Status** | [**PaymentMethodStatus**](PaymentMethodStatus.md) |  |  |
| **Metadata** | Pointer to **map[string]interface{}** |  | [optional]  |

## Methods

### NewPaymentMethod

`func NewPaymentMethod(id string, type_ PaymentMethodType, reusability PaymentMethodReusability, status PaymentMethodStatus, ) *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PaymentMethod) GetType() PaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*PaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v PaymentMethodType)`

SetType sets Type field to given value.


### GetCreated

`func (o *PaymentMethod) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PaymentMethod) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PaymentMethod) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PaymentMethod) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *PaymentMethod) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PaymentMethod) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PaymentMethod) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PaymentMethod) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentMethod) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentMethod) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentMethod) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentMethod) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentMethod) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentMethod) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetReferenceId

`func (o *PaymentMethod) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PaymentMethod) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PaymentMethod) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PaymentMethod) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetCard

`func (o *PaymentMethod) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaymentMethod) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaymentMethod) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *PaymentMethod) HasCard() bool`

HasCard returns a boolean if a field has been set.

### SetCardNil

`func (o *PaymentMethod) SetCardNil(b bool)`

 SetCardNil sets the value for Card to be an explicit nil

### UnsetCard
`func (o *PaymentMethod) UnsetCard()`

UnsetCard ensures that no value is present for Card, not even an explicit nil
### GetDirectDebit

`func (o *PaymentMethod) GetDirectDebit() DirectDebit`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *PaymentMethod) GetDirectDebitOk() (*DirectDebit, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *PaymentMethod) SetDirectDebit(v DirectDebit)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *PaymentMethod) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### SetDirectDebitNil

`func (o *PaymentMethod) SetDirectDebitNil(b bool)`

 SetDirectDebitNil sets the value for DirectDebit to be an explicit nil

### UnsetDirectDebit
`func (o *PaymentMethod) UnsetDirectDebit()`

UnsetDirectDebit ensures that no value is present for DirectDebit, not even an explicit nil
### GetEwallet

`func (o *PaymentMethod) GetEwallet() EWallet`

GetEwallet returns the Ewallet field if non-nil, zero value otherwise.

### GetEwalletOk

`func (o *PaymentMethod) GetEwalletOk() (*EWallet, bool)`

GetEwalletOk returns a tuple with the Ewallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwallet

`func (o *PaymentMethod) SetEwallet(v EWallet)`

SetEwallet sets Ewallet field to given value.

### HasEwallet

`func (o *PaymentMethod) HasEwallet() bool`

HasEwallet returns a boolean if a field has been set.

### SetEwalletNil

`func (o *PaymentMethod) SetEwalletNil(b bool)`

 SetEwalletNil sets the value for Ewallet to be an explicit nil

### UnsetEwallet
`func (o *PaymentMethod) UnsetEwallet()`

UnsetEwallet ensures that no value is present for Ewallet, not even an explicit nil
### GetOverTheCounter

`func (o *PaymentMethod) GetOverTheCounter() OverTheCounter`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *PaymentMethod) GetOverTheCounterOk() (*OverTheCounter, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *PaymentMethod) SetOverTheCounter(v OverTheCounter)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *PaymentMethod) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### SetOverTheCounterNil

`func (o *PaymentMethod) SetOverTheCounterNil(b bool)`

 SetOverTheCounterNil sets the value for OverTheCounter to be an explicit nil

### UnsetOverTheCounter
`func (o *PaymentMethod) UnsetOverTheCounter()`

UnsetOverTheCounter ensures that no value is present for OverTheCounter, not even an explicit nil
### GetVirtualAccount

`func (o *PaymentMethod) GetVirtualAccount() VirtualAccount`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *PaymentMethod) GetVirtualAccountOk() (*VirtualAccount, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *PaymentMethod) SetVirtualAccount(v VirtualAccount)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *PaymentMethod) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.

### SetVirtualAccountNil

`func (o *PaymentMethod) SetVirtualAccountNil(b bool)`

 SetVirtualAccountNil sets the value for VirtualAccount to be an explicit nil

### UnsetVirtualAccount
`func (o *PaymentMethod) UnsetVirtualAccount()`

UnsetVirtualAccount ensures that no value is present for VirtualAccount, not even an explicit nil
### GetQrCode

`func (o *PaymentMethod) GetQrCode() QRCode`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *PaymentMethod) GetQrCodeOk() (*QRCode, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *PaymentMethod) SetQrCode(v QRCode)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *PaymentMethod) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### SetQrCodeNil

`func (o *PaymentMethod) SetQrCodeNil(b bool)`

 SetQrCodeNil sets the value for QrCode to be an explicit nil

### UnsetQrCode
`func (o *PaymentMethod) UnsetQrCode()`

UnsetQrCode ensures that no value is present for QrCode, not even an explicit nil
### GetReusability

`func (o *PaymentMethod) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *PaymentMethod) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *PaymentMethod) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.


### GetStatus

`func (o *PaymentMethod) GetStatus() PaymentMethodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentMethod) GetStatusOk() (*PaymentMethodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentMethod) SetStatus(v PaymentMethodStatus)`

SetStatus sets Status field to given value.


### GetMetadata

`func (o *PaymentMethod) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentMethod) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentMethod) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentMethod) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PaymentMethod) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PaymentMethod) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to README]](../../README.md)



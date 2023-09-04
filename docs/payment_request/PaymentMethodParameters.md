# PaymentMethodParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PaymentMethodType**](PaymentMethodType.md) |  | 
**Reusability** | [**PaymentMethodReusability**](PaymentMethodReusability.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**DirectDebit** | Pointer to [**NullableDirectDebitParameters**](DirectDebitParameters.md) |  | [optional] 
**Ewallet** | Pointer to [**NullableEWalletParameters**](EWalletParameters.md) |  | [optional] 
**OverTheCounter** | Pointer to [**NullableOverTheCounterParameters**](OverTheCounterParameters.md) |  | [optional] 
**VirtualAccount** | Pointer to [**NullableVirtualAccountParameters**](VirtualAccountParameters.md) |  | [optional] 
**QrCode** | Pointer to [**NullableQRCodeParameters**](QRCodeParameters.md) |  | [optional] 

## Methods

### NewPaymentMethodParameters

`func NewPaymentMethodParameters(type_ PaymentMethodType, reusability PaymentMethodReusability, ) *PaymentMethodParameters`

NewPaymentMethodParameters instantiates a new PaymentMethodParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodParametersWithDefaults

`func NewPaymentMethodParametersWithDefaults() *PaymentMethodParameters`

NewPaymentMethodParametersWithDefaults instantiates a new PaymentMethodParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodParameters) GetType() PaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodParameters) GetTypeOk() (*PaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodParameters) SetType(v PaymentMethodType)`

SetType sets Type field to given value.


### GetReusability

`func (o *PaymentMethodParameters) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *PaymentMethodParameters) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *PaymentMethodParameters) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.


### GetDescription

`func (o *PaymentMethodParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentMethodParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentMethodParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentMethodParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentMethodParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentMethodParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetReferenceId

`func (o *PaymentMethodParameters) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PaymentMethodParameters) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PaymentMethodParameters) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PaymentMethodParameters) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetDirectDebit

`func (o *PaymentMethodParameters) GetDirectDebit() DirectDebitParameters`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *PaymentMethodParameters) GetDirectDebitOk() (*DirectDebitParameters, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *PaymentMethodParameters) SetDirectDebit(v DirectDebitParameters)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *PaymentMethodParameters) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### SetDirectDebitNil

`func (o *PaymentMethodParameters) SetDirectDebitNil(b bool)`

 SetDirectDebitNil sets the value for DirectDebit to be an explicit nil

### UnsetDirectDebit
`func (o *PaymentMethodParameters) UnsetDirectDebit()`

UnsetDirectDebit ensures that no value is present for DirectDebit, not even an explicit nil
### GetEwallet

`func (o *PaymentMethodParameters) GetEwallet() EWalletParameters`

GetEwallet returns the Ewallet field if non-nil, zero value otherwise.

### GetEwalletOk

`func (o *PaymentMethodParameters) GetEwalletOk() (*EWalletParameters, bool)`

GetEwalletOk returns a tuple with the Ewallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwallet

`func (o *PaymentMethodParameters) SetEwallet(v EWalletParameters)`

SetEwallet sets Ewallet field to given value.

### HasEwallet

`func (o *PaymentMethodParameters) HasEwallet() bool`

HasEwallet returns a boolean if a field has been set.

### SetEwalletNil

`func (o *PaymentMethodParameters) SetEwalletNil(b bool)`

 SetEwalletNil sets the value for Ewallet to be an explicit nil

### UnsetEwallet
`func (o *PaymentMethodParameters) UnsetEwallet()`

UnsetEwallet ensures that no value is present for Ewallet, not even an explicit nil
### GetOverTheCounter

`func (o *PaymentMethodParameters) GetOverTheCounter() OverTheCounterParameters`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *PaymentMethodParameters) GetOverTheCounterOk() (*OverTheCounterParameters, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *PaymentMethodParameters) SetOverTheCounter(v OverTheCounterParameters)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *PaymentMethodParameters) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### SetOverTheCounterNil

`func (o *PaymentMethodParameters) SetOverTheCounterNil(b bool)`

 SetOverTheCounterNil sets the value for OverTheCounter to be an explicit nil

### UnsetOverTheCounter
`func (o *PaymentMethodParameters) UnsetOverTheCounter()`

UnsetOverTheCounter ensures that no value is present for OverTheCounter, not even an explicit nil
### GetVirtualAccount

`func (o *PaymentMethodParameters) GetVirtualAccount() VirtualAccountParameters`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *PaymentMethodParameters) GetVirtualAccountOk() (*VirtualAccountParameters, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *PaymentMethodParameters) SetVirtualAccount(v VirtualAccountParameters)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *PaymentMethodParameters) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.

### SetVirtualAccountNil

`func (o *PaymentMethodParameters) SetVirtualAccountNil(b bool)`

 SetVirtualAccountNil sets the value for VirtualAccount to be an explicit nil

### UnsetVirtualAccount
`func (o *PaymentMethodParameters) UnsetVirtualAccount()`

UnsetVirtualAccount ensures that no value is present for VirtualAccount, not even an explicit nil
### GetQrCode

`func (o *PaymentMethodParameters) GetQrCode() QRCodeParameters`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *PaymentMethodParameters) GetQrCodeOk() (*QRCodeParameters, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *PaymentMethodParameters) SetQrCode(v QRCodeParameters)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *PaymentMethodParameters) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### SetQrCodeNil

`func (o *PaymentMethodParameters) SetQrCodeNil(b bool)`

 SetQrCodeNil sets the value for QrCode to be an explicit nil

### UnsetQrCode
`func (o *PaymentMethodParameters) UnsetQrCode()`

UnsetQrCode ensures that no value is present for QrCode, not even an explicit nil

[[Back to README]](../../README.md)



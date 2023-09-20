# EWalletChannelProperties

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **SuccessReturnUrl** | Pointer to **string** | URL where the end-customer is redirected if the authorization is successful | [optional]  |
| **FailureReturnUrl** | Pointer to **string** | URL where the end-customer is redirected if the authorization failed | [optional]  |
| **CancelReturnUrl** | Pointer to **string** | URL where the end-customer is redirected if the authorization cancelled | [optional]  |
| **MobileNumber** | Pointer to **string** | Mobile number of customer in E.164 format (e.g. +628123123123). For OVO one time payment use only. | [optional]  |
| **RedeemPoints** | Pointer to **string** | REDEEM_NONE will not use any point, REDEEM_ALL will use all available points before cash balance is used. For OVO and ShopeePay tokenized payment use only. | [optional]  |
| **Cashtag** | Pointer to **string** | Available for JENIUSPAY only | [optional]  |

## Methods

### NewEWalletChannelProperties

`func NewEWalletChannelProperties() *EWalletChannelProperties`

NewEWalletChannelProperties instantiates a new EWalletChannelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEWalletChannelPropertiesWithDefaults

`func NewEWalletChannelPropertiesWithDefaults() *EWalletChannelProperties`

NewEWalletChannelPropertiesWithDefaults instantiates a new EWalletChannelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessReturnUrl

`func (o *EWalletChannelProperties) GetSuccessReturnUrl() string`

GetSuccessReturnUrl returns the SuccessReturnUrl field if non-nil, zero value otherwise.

### GetSuccessReturnUrlOk

`func (o *EWalletChannelProperties) GetSuccessReturnUrlOk() (*string, bool)`

GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessReturnUrl

`func (o *EWalletChannelProperties) SetSuccessReturnUrl(v string)`

SetSuccessReturnUrl sets SuccessReturnUrl field to given value.

### HasSuccessReturnUrl

`func (o *EWalletChannelProperties) HasSuccessReturnUrl() bool`

HasSuccessReturnUrl returns a boolean if a field has been set.

### GetFailureReturnUrl

`func (o *EWalletChannelProperties) GetFailureReturnUrl() string`

GetFailureReturnUrl returns the FailureReturnUrl field if non-nil, zero value otherwise.

### GetFailureReturnUrlOk

`func (o *EWalletChannelProperties) GetFailureReturnUrlOk() (*string, bool)`

GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReturnUrl

`func (o *EWalletChannelProperties) SetFailureReturnUrl(v string)`

SetFailureReturnUrl sets FailureReturnUrl field to given value.

### HasFailureReturnUrl

`func (o *EWalletChannelProperties) HasFailureReturnUrl() bool`

HasFailureReturnUrl returns a boolean if a field has been set.

### GetCancelReturnUrl

`func (o *EWalletChannelProperties) GetCancelReturnUrl() string`

GetCancelReturnUrl returns the CancelReturnUrl field if non-nil, zero value otherwise.

### GetCancelReturnUrlOk

`func (o *EWalletChannelProperties) GetCancelReturnUrlOk() (*string, bool)`

GetCancelReturnUrlOk returns a tuple with the CancelReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReturnUrl

`func (o *EWalletChannelProperties) SetCancelReturnUrl(v string)`

SetCancelReturnUrl sets CancelReturnUrl field to given value.

### HasCancelReturnUrl

`func (o *EWalletChannelProperties) HasCancelReturnUrl() bool`

HasCancelReturnUrl returns a boolean if a field has been set.

### GetMobileNumber

`func (o *EWalletChannelProperties) GetMobileNumber() string`

GetMobileNumber returns the MobileNumber field if non-nil, zero value otherwise.

### GetMobileNumberOk

`func (o *EWalletChannelProperties) GetMobileNumberOk() (*string, bool)`

GetMobileNumberOk returns a tuple with the MobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumber

`func (o *EWalletChannelProperties) SetMobileNumber(v string)`

SetMobileNumber sets MobileNumber field to given value.

### HasMobileNumber

`func (o *EWalletChannelProperties) HasMobileNumber() bool`

HasMobileNumber returns a boolean if a field has been set.

### GetRedeemPoints

`func (o *EWalletChannelProperties) GetRedeemPoints() string`

GetRedeemPoints returns the RedeemPoints field if non-nil, zero value otherwise.

### GetRedeemPointsOk

`func (o *EWalletChannelProperties) GetRedeemPointsOk() (*string, bool)`

GetRedeemPointsOk returns a tuple with the RedeemPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemPoints

`func (o *EWalletChannelProperties) SetRedeemPoints(v string)`

SetRedeemPoints sets RedeemPoints field to given value.

### HasRedeemPoints

`func (o *EWalletChannelProperties) HasRedeemPoints() bool`

HasRedeemPoints returns a boolean if a field has been set.

### GetCashtag

`func (o *EWalletChannelProperties) GetCashtag() string`

GetCashtag returns the Cashtag field if non-nil, zero value otherwise.

### GetCashtagOk

`func (o *EWalletChannelProperties) GetCashtagOk() (*string, bool)`

GetCashtagOk returns a tuple with the Cashtag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashtag

`func (o *EWalletChannelProperties) SetCashtag(v string)`

SetCashtag sets Cashtag field to given value.

### HasCashtag

`func (o *EWalletChannelProperties) HasCashtag() bool`

HasCashtag returns a boolean if a field has been set.


[[Back to README]](../../README.md)



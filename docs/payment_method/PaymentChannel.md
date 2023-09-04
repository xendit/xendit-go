# PaymentChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | Pointer to **string** | The specific Xendit code used to identify the partner channel | [optional] 
**Type** | Pointer to **string** | The payment method type | [optional] 
**Country** | Pointer to **string** | The country where the channel operates  in ISO 3166-2 country code | [optional] 
**ChannelName** | Pointer to **string** | Official parter name of the payment channel | [optional] 
**ChannelProperties** | Pointer to [**[]ChannelProperty**](ChannelProperty.md) |  | [optional] 
**LogoUrl** | Pointer to **string** | If available, this contains a URL to the logo of the partner channel | [optional] 
**AmountLimits** | Pointer to [**[]ChannelAmountLimits**](ChannelAmountLimits.md) |  | [optional] 

## Methods

### NewPaymentChannel

`func NewPaymentChannel() *PaymentChannel`

NewPaymentChannel instantiates a new PaymentChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentChannelWithDefaults

`func NewPaymentChannelWithDefaults() *PaymentChannel`

NewPaymentChannelWithDefaults instantiates a new PaymentChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *PaymentChannel) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *PaymentChannel) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *PaymentChannel) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *PaymentChannel) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### GetType

`func (o *PaymentChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentChannel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentChannel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCountry

`func (o *PaymentChannel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentChannel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentChannel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentChannel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetChannelName

`func (o *PaymentChannel) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *PaymentChannel) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *PaymentChannel) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *PaymentChannel) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetChannelProperties

`func (o *PaymentChannel) GetChannelProperties() []ChannelProperty`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *PaymentChannel) GetChannelPropertiesOk() (*[]ChannelProperty, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *PaymentChannel) SetChannelProperties(v []ChannelProperty)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *PaymentChannel) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.

### GetLogoUrl

`func (o *PaymentChannel) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *PaymentChannel) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *PaymentChannel) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *PaymentChannel) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetAmountLimits

`func (o *PaymentChannel) GetAmountLimits() []ChannelAmountLimits`

GetAmountLimits returns the AmountLimits field if non-nil, zero value otherwise.

### GetAmountLimitsOk

`func (o *PaymentChannel) GetAmountLimitsOk() (*[]ChannelAmountLimits, bool)`

GetAmountLimitsOk returns a tuple with the AmountLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimits

`func (o *PaymentChannel) SetAmountLimits(v []ChannelAmountLimits)`

SetAmountLimits sets AmountLimits field to given value.

### HasAmountLimits

`func (o *PaymentChannel) HasAmountLimits() bool`

HasAmountLimits returns a boolean if a field has been set.


[[Back to README]](../../README.md)



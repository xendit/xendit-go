# QRCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelCode** | Pointer to [**NullableQRCodeChannelCode**](QRCodeChannelCode.md) |  | [optional] 
**ChannelProperties** | Pointer to [**QRCodeChannelProperties**](QRCodeChannelProperties.md) |  | [optional] 

## Methods

### NewQRCode

`func NewQRCode() *QRCode`

NewQRCode instantiates a new QRCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQRCodeWithDefaults

`func NewQRCodeWithDefaults() *QRCode`

NewQRCodeWithDefaults instantiates a new QRCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelCode

`func (o *QRCode) GetChannelCode() QRCodeChannelCode`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *QRCode) GetChannelCodeOk() (*QRCodeChannelCode, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *QRCode) SetChannelCode(v QRCodeChannelCode)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *QRCode) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### SetChannelCodeNil

`func (o *QRCode) SetChannelCodeNil(b bool)`

 SetChannelCodeNil sets the value for ChannelCode to be an explicit nil

### UnsetChannelCode
`func (o *QRCode) UnsetChannelCode()`

UnsetChannelCode ensures that no value is present for ChannelCode, not even an explicit nil
### GetChannelProperties

`func (o *QRCode) GetChannelProperties() QRCodeChannelProperties`

GetChannelProperties returns the ChannelProperties field if non-nil, zero value otherwise.

### GetChannelPropertiesOk

`func (o *QRCode) GetChannelPropertiesOk() (*QRCodeChannelProperties, bool)`

GetChannelPropertiesOk returns a tuple with the ChannelProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelProperties

`func (o *QRCode) SetChannelProperties(v QRCodeChannelProperties)`

SetChannelProperties sets ChannelProperties field to given value.

### HasChannelProperties

`func (o *QRCode) HasChannelProperties() bool`

HasChannelProperties returns a boolean if a field has been set.


[[Back to README]](../../README.md)



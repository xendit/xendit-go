# QRCodeChannelProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QrString** | Pointer to **string** | QR string to be rendered for display to end users. QR string to image rendering are commonly available in software libraries (e.g Nodejs, PHP, Java) | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewQRCodeChannelProperties

`func NewQRCodeChannelProperties() *QRCodeChannelProperties`

NewQRCodeChannelProperties instantiates a new QRCodeChannelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQRCodeChannelPropertiesWithDefaults

`func NewQRCodeChannelPropertiesWithDefaults() *QRCodeChannelProperties`

NewQRCodeChannelPropertiesWithDefaults instantiates a new QRCodeChannelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQrString

`func (o *QRCodeChannelProperties) GetQrString() string`

GetQrString returns the QrString field if non-nil, zero value otherwise.

### GetQrStringOk

`func (o *QRCodeChannelProperties) GetQrStringOk() (*string, bool)`

GetQrStringOk returns a tuple with the QrString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrString

`func (o *QRCodeChannelProperties) SetQrString(v string)`

SetQrString sets QrString field to given value.

### HasQrString

`func (o *QRCodeChannelProperties) HasQrString() bool`

HasQrString returns a boolean if a field has been set.

### GetExpiresAt

`func (o *QRCodeChannelProperties) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *QRCodeChannelProperties) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *QRCodeChannelProperties) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *QRCodeChannelProperties) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to README]](../../README.md)



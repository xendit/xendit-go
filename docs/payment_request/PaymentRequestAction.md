# PaymentRequestAction


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Action** | **string** | ☑️ |  |  |
| **UrlType** | **string** | ☑️ |  |  |
| **Method** | **NullableString** | ☑️ |  |  |
| **Url** | **NullableString** | ☑️ |  |  |
| **QrCode** | **NullableString** | ☑️ |  |  |

## Methods

### NewPaymentRequestAction

`func NewPaymentRequestAction(action string, urlType string, method NullableString, url NullableString, qrCode NullableString, ) *PaymentRequestAction`

NewPaymentRequestAction instantiates a new PaymentRequestAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestActionWithDefaults

`func NewPaymentRequestActionWithDefaults() *PaymentRequestAction`

NewPaymentRequestActionWithDefaults instantiates a new PaymentRequestAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PaymentRequestAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PaymentRequestAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PaymentRequestAction) SetAction(v string)`

SetAction sets Action field to given value.


### GetUrlType

`func (o *PaymentRequestAction) GetUrlType() string`

GetUrlType returns the UrlType field if non-nil, zero value otherwise.

### GetUrlTypeOk

`func (o *PaymentRequestAction) GetUrlTypeOk() (*string, bool)`

GetUrlTypeOk returns a tuple with the UrlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlType

`func (o *PaymentRequestAction) SetUrlType(v string)`

SetUrlType sets UrlType field to given value.


### GetMethod

`func (o *PaymentRequestAction) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *PaymentRequestAction) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *PaymentRequestAction) SetMethod(v string)`

SetMethod sets Method field to given value.


### SetMethodNil

`func (o *PaymentRequestAction) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *PaymentRequestAction) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetUrl

`func (o *PaymentRequestAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PaymentRequestAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PaymentRequestAction) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *PaymentRequestAction) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PaymentRequestAction) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetQrCode

`func (o *PaymentRequestAction) GetQrCode() string`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *PaymentRequestAction) GetQrCodeOk() (*string, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *PaymentRequestAction) SetQrCode(v string)`

SetQrCode sets QrCode field to given value.


### SetQrCodeNil

`func (o *PaymentRequestAction) SetQrCodeNil(b bool)`

 SetQrCodeNil sets the value for QrCode to be an explicit nil

### UnsetQrCode
`func (o *PaymentRequestAction) UnsetQrCode()`

UnsetQrCode ensures that no value is present for QrCode, not even an explicit nil

[[Back to README]](../../README.md)



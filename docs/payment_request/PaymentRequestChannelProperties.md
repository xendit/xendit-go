# PaymentRequestChannelProperties


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **SuccessReturnUrl** | Pointer to **string** |  | URL where the end-customer is redirected if the authorization is successful |  |
| **FailureReturnUrl** | Pointer to **string** |  | URL where the end-customer is redirected if the authorization failed |  |
| **CancelReturnUrl** | Pointer to **string** |  | URL where the end-customer is redirected if the authorization cancelled |  |
| **PendingReturnUrl** | Pointer to **string** |  | URL where the end-customer is redirected if the authorization is pending |  |
| **RedeemPoints** | Pointer to **string** |  | REDEEM_NONE will not use any point, REDEEM_ALL will use all available points before cash balance is used. For OVO and ShopeePay tokenized payment use only. |  |
| **RequireAuth** | Pointer to **bool** |  | Toggle used to require end-customer to input undergo OTP validation before completing a payment. OTP will always be required for transactions greater than 1,000,000 IDR. For BRI tokenized payment use only. |  |
| **MerchantIdTag** | Pointer to **string** |  | Tag for a Merchant ID that you want to associate this payment with. For merchants using their own MIDs to specify which MID they want to use  |  |
| **CardonfileType** | Pointer to **NullableString** |  | Type of “credential-on-file” / “card-on-file” payment being made. Indicate that this payment uses a previously linked Payment Method for charging. |  |

## Methods

### NewPaymentRequestChannelProperties

`func NewPaymentRequestChannelProperties() *PaymentRequestChannelProperties`

NewPaymentRequestChannelProperties instantiates a new PaymentRequestChannelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestChannelPropertiesWithDefaults

`func NewPaymentRequestChannelPropertiesWithDefaults() *PaymentRequestChannelProperties`

NewPaymentRequestChannelPropertiesWithDefaults instantiates a new PaymentRequestChannelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessReturnUrl

`func (o *PaymentRequestChannelProperties) GetSuccessReturnUrl() string`

GetSuccessReturnUrl returns the SuccessReturnUrl field if non-nil, zero value otherwise.

### GetSuccessReturnUrlOk

`func (o *PaymentRequestChannelProperties) GetSuccessReturnUrlOk() (*string, bool)`

GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessReturnUrl

`func (o *PaymentRequestChannelProperties) SetSuccessReturnUrl(v string)`

SetSuccessReturnUrl sets SuccessReturnUrl field to given value.

### HasSuccessReturnUrl

`func (o *PaymentRequestChannelProperties) HasSuccessReturnUrl() bool`

HasSuccessReturnUrl returns a boolean if a field has been set.

### GetFailureReturnUrl

`func (o *PaymentRequestChannelProperties) GetFailureReturnUrl() string`

GetFailureReturnUrl returns the FailureReturnUrl field if non-nil, zero value otherwise.

### GetFailureReturnUrlOk

`func (o *PaymentRequestChannelProperties) GetFailureReturnUrlOk() (*string, bool)`

GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReturnUrl

`func (o *PaymentRequestChannelProperties) SetFailureReturnUrl(v string)`

SetFailureReturnUrl sets FailureReturnUrl field to given value.

### HasFailureReturnUrl

`func (o *PaymentRequestChannelProperties) HasFailureReturnUrl() bool`

HasFailureReturnUrl returns a boolean if a field has been set.

### GetCancelReturnUrl

`func (o *PaymentRequestChannelProperties) GetCancelReturnUrl() string`

GetCancelReturnUrl returns the CancelReturnUrl field if non-nil, zero value otherwise.

### GetCancelReturnUrlOk

`func (o *PaymentRequestChannelProperties) GetCancelReturnUrlOk() (*string, bool)`

GetCancelReturnUrlOk returns a tuple with the CancelReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReturnUrl

`func (o *PaymentRequestChannelProperties) SetCancelReturnUrl(v string)`

SetCancelReturnUrl sets CancelReturnUrl field to given value.

### HasCancelReturnUrl

`func (o *PaymentRequestChannelProperties) HasCancelReturnUrl() bool`

HasCancelReturnUrl returns a boolean if a field has been set.

### GetPendingReturnUrl

`func (o *PaymentRequestChannelProperties) GetPendingReturnUrl() string`

GetPendingReturnUrl returns the PendingReturnUrl field if non-nil, zero value otherwise.

### GetPendingReturnUrlOk

`func (o *PaymentRequestChannelProperties) GetPendingReturnUrlOk() (*string, bool)`

GetPendingReturnUrlOk returns a tuple with the PendingReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingReturnUrl

`func (o *PaymentRequestChannelProperties) SetPendingReturnUrl(v string)`

SetPendingReturnUrl sets PendingReturnUrl field to given value.

### HasPendingReturnUrl

`func (o *PaymentRequestChannelProperties) HasPendingReturnUrl() bool`

HasPendingReturnUrl returns a boolean if a field has been set.

### GetRedeemPoints

`func (o *PaymentRequestChannelProperties) GetRedeemPoints() string`

GetRedeemPoints returns the RedeemPoints field if non-nil, zero value otherwise.

### GetRedeemPointsOk

`func (o *PaymentRequestChannelProperties) GetRedeemPointsOk() (*string, bool)`

GetRedeemPointsOk returns a tuple with the RedeemPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemPoints

`func (o *PaymentRequestChannelProperties) SetRedeemPoints(v string)`

SetRedeemPoints sets RedeemPoints field to given value.

### HasRedeemPoints

`func (o *PaymentRequestChannelProperties) HasRedeemPoints() bool`

HasRedeemPoints returns a boolean if a field has been set.

### GetRequireAuth

`func (o *PaymentRequestChannelProperties) GetRequireAuth() bool`

GetRequireAuth returns the RequireAuth field if non-nil, zero value otherwise.

### GetRequireAuthOk

`func (o *PaymentRequestChannelProperties) GetRequireAuthOk() (*bool, bool)`

GetRequireAuthOk returns a tuple with the RequireAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuth

`func (o *PaymentRequestChannelProperties) SetRequireAuth(v bool)`

SetRequireAuth sets RequireAuth field to given value.

### HasRequireAuth

`func (o *PaymentRequestChannelProperties) HasRequireAuth() bool`

HasRequireAuth returns a boolean if a field has been set.

### GetMerchantIdTag

`func (o *PaymentRequestChannelProperties) GetMerchantIdTag() string`

GetMerchantIdTag returns the MerchantIdTag field if non-nil, zero value otherwise.

### GetMerchantIdTagOk

`func (o *PaymentRequestChannelProperties) GetMerchantIdTagOk() (*string, bool)`

GetMerchantIdTagOk returns a tuple with the MerchantIdTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantIdTag

`func (o *PaymentRequestChannelProperties) SetMerchantIdTag(v string)`

SetMerchantIdTag sets MerchantIdTag field to given value.

### HasMerchantIdTag

`func (o *PaymentRequestChannelProperties) HasMerchantIdTag() bool`

HasMerchantIdTag returns a boolean if a field has been set.

### GetCardonfileType

`func (o *PaymentRequestChannelProperties) GetCardonfileType() string`

GetCardonfileType returns the CardonfileType field if non-nil, zero value otherwise.

### GetCardonfileTypeOk

`func (o *PaymentRequestChannelProperties) GetCardonfileTypeOk() (*string, bool)`

GetCardonfileTypeOk returns a tuple with the CardonfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardonfileType

`func (o *PaymentRequestChannelProperties) SetCardonfileType(v string)`

SetCardonfileType sets CardonfileType field to given value.

### HasCardonfileType

`func (o *PaymentRequestChannelProperties) HasCardonfileType() bool`

HasCardonfileType returns a boolean if a field has been set.

### SetCardonfileTypeNil

`func (o *PaymentRequestChannelProperties) SetCardonfileTypeNil(b bool)`

 SetCardonfileTypeNil sets the value for CardonfileType to be an explicit nil

### UnsetCardonfileType
`func (o *PaymentRequestChannelProperties) UnsetCardonfileType()`

UnsetCardonfileType ensures that no value is present for CardonfileType, not even an explicit nil

[[Back to README]](../../README.md)



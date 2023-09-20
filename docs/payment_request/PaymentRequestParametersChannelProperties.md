# PaymentRequestParametersChannelProperties

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **SuccessReturnUrl** | Pointer to **string** | URL where the end-customer is redirected if the authorization is successful | [optional]  |
| **FailureReturnUrl** | Pointer to **string** | URL where the end-customer is redirected if the authorization failed | [optional]  |
| **CancelReturnUrl** | Pointer to **string** | URL where the end-customer is redirected if the authorization cancelled | [optional]  |
| **RedeemPoints** | Pointer to **string** | REDEEM_NONE will not use any point, REDEEM_ALL will use all available points before cash balance is used. For OVO and ShopeePay tokenized payment use only. | [optional]  |
| **RequireAuth** | Pointer to **bool** | Toggle used to require end-customer to input undergo OTP validation before completing a payment. OTP will always be required for transactions greater than 1,000,000 IDR. For BRI tokenized payment use only. | [optional]  |
| **MerchantIdTag** | Pointer to **string** | Tag for a Merchant ID that you want to associate this payment with. For merchants using their own MIDs to specify which MID they want to use  | [optional]  |
| **CardonfileType** | Pointer to **NullableString** | Type of “credential-on-file” / “card-on-file” payment being made. Indicate that this payment uses a previously linked Payment Method for charging. | [optional]  |
| **Cvv** | Pointer to **string** | Three digit code written on the back of the card (usually called CVV/CVN). | [optional]  |

## Methods

### NewPaymentRequestParametersChannelProperties

`func NewPaymentRequestParametersChannelProperties() *PaymentRequestParametersChannelProperties`

NewPaymentRequestParametersChannelProperties instantiates a new PaymentRequestParametersChannelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestParametersChannelPropertiesWithDefaults

`func NewPaymentRequestParametersChannelPropertiesWithDefaults() *PaymentRequestParametersChannelProperties`

NewPaymentRequestParametersChannelPropertiesWithDefaults instantiates a new PaymentRequestParametersChannelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessReturnUrl

`func (o *PaymentRequestParametersChannelProperties) GetSuccessReturnUrl() string`

GetSuccessReturnUrl returns the SuccessReturnUrl field if non-nil, zero value otherwise.

### GetSuccessReturnUrlOk

`func (o *PaymentRequestParametersChannelProperties) GetSuccessReturnUrlOk() (*string, bool)`

GetSuccessReturnUrlOk returns a tuple with the SuccessReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessReturnUrl

`func (o *PaymentRequestParametersChannelProperties) SetSuccessReturnUrl(v string)`

SetSuccessReturnUrl sets SuccessReturnUrl field to given value.

### HasSuccessReturnUrl

`func (o *PaymentRequestParametersChannelProperties) HasSuccessReturnUrl() bool`

HasSuccessReturnUrl returns a boolean if a field has been set.

### GetFailureReturnUrl

`func (o *PaymentRequestParametersChannelProperties) GetFailureReturnUrl() string`

GetFailureReturnUrl returns the FailureReturnUrl field if non-nil, zero value otherwise.

### GetFailureReturnUrlOk

`func (o *PaymentRequestParametersChannelProperties) GetFailureReturnUrlOk() (*string, bool)`

GetFailureReturnUrlOk returns a tuple with the FailureReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReturnUrl

`func (o *PaymentRequestParametersChannelProperties) SetFailureReturnUrl(v string)`

SetFailureReturnUrl sets FailureReturnUrl field to given value.

### HasFailureReturnUrl

`func (o *PaymentRequestParametersChannelProperties) HasFailureReturnUrl() bool`

HasFailureReturnUrl returns a boolean if a field has been set.

### GetCancelReturnUrl

`func (o *PaymentRequestParametersChannelProperties) GetCancelReturnUrl() string`

GetCancelReturnUrl returns the CancelReturnUrl field if non-nil, zero value otherwise.

### GetCancelReturnUrlOk

`func (o *PaymentRequestParametersChannelProperties) GetCancelReturnUrlOk() (*string, bool)`

GetCancelReturnUrlOk returns a tuple with the CancelReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReturnUrl

`func (o *PaymentRequestParametersChannelProperties) SetCancelReturnUrl(v string)`

SetCancelReturnUrl sets CancelReturnUrl field to given value.

### HasCancelReturnUrl

`func (o *PaymentRequestParametersChannelProperties) HasCancelReturnUrl() bool`

HasCancelReturnUrl returns a boolean if a field has been set.

### GetRedeemPoints

`func (o *PaymentRequestParametersChannelProperties) GetRedeemPoints() string`

GetRedeemPoints returns the RedeemPoints field if non-nil, zero value otherwise.

### GetRedeemPointsOk

`func (o *PaymentRequestParametersChannelProperties) GetRedeemPointsOk() (*string, bool)`

GetRedeemPointsOk returns a tuple with the RedeemPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemPoints

`func (o *PaymentRequestParametersChannelProperties) SetRedeemPoints(v string)`

SetRedeemPoints sets RedeemPoints field to given value.

### HasRedeemPoints

`func (o *PaymentRequestParametersChannelProperties) HasRedeemPoints() bool`

HasRedeemPoints returns a boolean if a field has been set.

### GetRequireAuth

`func (o *PaymentRequestParametersChannelProperties) GetRequireAuth() bool`

GetRequireAuth returns the RequireAuth field if non-nil, zero value otherwise.

### GetRequireAuthOk

`func (o *PaymentRequestParametersChannelProperties) GetRequireAuthOk() (*bool, bool)`

GetRequireAuthOk returns a tuple with the RequireAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuth

`func (o *PaymentRequestParametersChannelProperties) SetRequireAuth(v bool)`

SetRequireAuth sets RequireAuth field to given value.

### HasRequireAuth

`func (o *PaymentRequestParametersChannelProperties) HasRequireAuth() bool`

HasRequireAuth returns a boolean if a field has been set.

### GetMerchantIdTag

`func (o *PaymentRequestParametersChannelProperties) GetMerchantIdTag() string`

GetMerchantIdTag returns the MerchantIdTag field if non-nil, zero value otherwise.

### GetMerchantIdTagOk

`func (o *PaymentRequestParametersChannelProperties) GetMerchantIdTagOk() (*string, bool)`

GetMerchantIdTagOk returns a tuple with the MerchantIdTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantIdTag

`func (o *PaymentRequestParametersChannelProperties) SetMerchantIdTag(v string)`

SetMerchantIdTag sets MerchantIdTag field to given value.

### HasMerchantIdTag

`func (o *PaymentRequestParametersChannelProperties) HasMerchantIdTag() bool`

HasMerchantIdTag returns a boolean if a field has been set.

### GetCardonfileType

`func (o *PaymentRequestParametersChannelProperties) GetCardonfileType() string`

GetCardonfileType returns the CardonfileType field if non-nil, zero value otherwise.

### GetCardonfileTypeOk

`func (o *PaymentRequestParametersChannelProperties) GetCardonfileTypeOk() (*string, bool)`

GetCardonfileTypeOk returns a tuple with the CardonfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardonfileType

`func (o *PaymentRequestParametersChannelProperties) SetCardonfileType(v string)`

SetCardonfileType sets CardonfileType field to given value.

### HasCardonfileType

`func (o *PaymentRequestParametersChannelProperties) HasCardonfileType() bool`

HasCardonfileType returns a boolean if a field has been set.

### SetCardonfileTypeNil

`func (o *PaymentRequestParametersChannelProperties) SetCardonfileTypeNil(b bool)`

 SetCardonfileTypeNil sets the value for CardonfileType to be an explicit nil

### UnsetCardonfileType
`func (o *PaymentRequestParametersChannelProperties) UnsetCardonfileType()`

UnsetCardonfileType ensures that no value is present for CardonfileType, not even an explicit nil
### GetCvv

`func (o *PaymentRequestParametersChannelProperties) GetCvv() string`

GetCvv returns the Cvv field if non-nil, zero value otherwise.

### GetCvvOk

`func (o *PaymentRequestParametersChannelProperties) GetCvvOk() (*string, bool)`

GetCvvOk returns a tuple with the Cvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv

`func (o *PaymentRequestParametersChannelProperties) SetCvv(v string)`

SetCvv sets Cvv field to given value.

### HasCvv

`func (o *PaymentRequestParametersChannelProperties) HasCvv() bool`

HasCvv returns a boolean if a field has been set.


[[Back to README]](../../README.md)



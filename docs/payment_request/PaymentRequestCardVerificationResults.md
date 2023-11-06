# PaymentRequestCardVerificationResults


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ThreeDSecure** | [**NullablePaymentRequestCardVerificationResultsThreeDeeSecure**](PaymentRequestCardVerificationResultsThreeDeeSecure.md) | ☑️ |  |  |
| **CvvResult** | Pointer to **NullableString** |  |  |  |
| **AddressVerificationResult** | Pointer to **NullableString** |  |  |  |

## Methods

### NewPaymentRequestCardVerificationResults

`func NewPaymentRequestCardVerificationResults(threeDSecure NullablePaymentRequestCardVerificationResultsThreeDeeSecure, ) *PaymentRequestCardVerificationResults`

NewPaymentRequestCardVerificationResults instantiates a new PaymentRequestCardVerificationResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestCardVerificationResultsWithDefaults

`func NewPaymentRequestCardVerificationResultsWithDefaults() *PaymentRequestCardVerificationResults`

NewPaymentRequestCardVerificationResultsWithDefaults instantiates a new PaymentRequestCardVerificationResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreeDSecure

`func (o *PaymentRequestCardVerificationResults) GetThreeDSecure() PaymentRequestCardVerificationResultsThreeDeeSecure`

GetThreeDSecure returns the ThreeDSecure field if non-nil, zero value otherwise.

### GetThreeDSecureOk

`func (o *PaymentRequestCardVerificationResults) GetThreeDSecureOk() (*PaymentRequestCardVerificationResultsThreeDeeSecure, bool)`

GetThreeDSecureOk returns a tuple with the ThreeDSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecure

`func (o *PaymentRequestCardVerificationResults) SetThreeDSecure(v PaymentRequestCardVerificationResultsThreeDeeSecure)`

SetThreeDSecure sets ThreeDSecure field to given value.


### SetThreeDSecureNil

`func (o *PaymentRequestCardVerificationResults) SetThreeDSecureNil(b bool)`

 SetThreeDSecureNil sets the value for ThreeDSecure to be an explicit nil

### UnsetThreeDSecure
`func (o *PaymentRequestCardVerificationResults) UnsetThreeDSecure()`

UnsetThreeDSecure ensures that no value is present for ThreeDSecure, not even an explicit nil
### GetCvvResult

`func (o *PaymentRequestCardVerificationResults) GetCvvResult() string`

GetCvvResult returns the CvvResult field if non-nil, zero value otherwise.

### GetCvvResultOk

`func (o *PaymentRequestCardVerificationResults) GetCvvResultOk() (*string, bool)`

GetCvvResultOk returns a tuple with the CvvResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvResult

`func (o *PaymentRequestCardVerificationResults) SetCvvResult(v string)`

SetCvvResult sets CvvResult field to given value.

### HasCvvResult

`func (o *PaymentRequestCardVerificationResults) HasCvvResult() bool`

HasCvvResult returns a boolean if a field has been set.

### SetCvvResultNil

`func (o *PaymentRequestCardVerificationResults) SetCvvResultNil(b bool)`

 SetCvvResultNil sets the value for CvvResult to be an explicit nil

### UnsetCvvResult
`func (o *PaymentRequestCardVerificationResults) UnsetCvvResult()`

UnsetCvvResult ensures that no value is present for CvvResult, not even an explicit nil
### GetAddressVerificationResult

`func (o *PaymentRequestCardVerificationResults) GetAddressVerificationResult() string`

GetAddressVerificationResult returns the AddressVerificationResult field if non-nil, zero value otherwise.

### GetAddressVerificationResultOk

`func (o *PaymentRequestCardVerificationResults) GetAddressVerificationResultOk() (*string, bool)`

GetAddressVerificationResultOk returns a tuple with the AddressVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerificationResult

`func (o *PaymentRequestCardVerificationResults) SetAddressVerificationResult(v string)`

SetAddressVerificationResult sets AddressVerificationResult field to given value.

### HasAddressVerificationResult

`func (o *PaymentRequestCardVerificationResults) HasAddressVerificationResult() bool`

HasAddressVerificationResult returns a boolean if a field has been set.

### SetAddressVerificationResultNil

`func (o *PaymentRequestCardVerificationResults) SetAddressVerificationResultNil(b bool)`

 SetAddressVerificationResultNil sets the value for AddressVerificationResult to be an explicit nil

### UnsetAddressVerificationResult
`func (o *PaymentRequestCardVerificationResults) UnsetAddressVerificationResult()`

UnsetAddressVerificationResult ensures that no value is present for AddressVerificationResult, not even an explicit nil

[[Back to README]](../../README.md)



# CardVerificationResults

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **ThreeDSecure** | Pointer to [**CardVerificationResultsThreeDSecure**](CardVerificationResultsThreeDSecure.md) |  | [optional]  |
| **CvvResult** | Pointer to **NullableString** |  | [optional]  |
| **AddressVerificationResult** | Pointer to **NullableString** |  | [optional]  |

## Methods

### NewCardVerificationResults

`func NewCardVerificationResults() *CardVerificationResults`

NewCardVerificationResults instantiates a new CardVerificationResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardVerificationResultsWithDefaults

`func NewCardVerificationResultsWithDefaults() *CardVerificationResults`

NewCardVerificationResultsWithDefaults instantiates a new CardVerificationResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreeDSecure

`func (o *CardVerificationResults) GetThreeDSecure() CardVerificationResultsThreeDSecure`

GetThreeDSecure returns the ThreeDSecure field if non-nil, zero value otherwise.

### GetThreeDSecureOk

`func (o *CardVerificationResults) GetThreeDSecureOk() (*CardVerificationResultsThreeDSecure, bool)`

GetThreeDSecureOk returns a tuple with the ThreeDSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDSecure

`func (o *CardVerificationResults) SetThreeDSecure(v CardVerificationResultsThreeDSecure)`

SetThreeDSecure sets ThreeDSecure field to given value.

### HasThreeDSecure

`func (o *CardVerificationResults) HasThreeDSecure() bool`

HasThreeDSecure returns a boolean if a field has been set.

### GetCvvResult

`func (o *CardVerificationResults) GetCvvResult() string`

GetCvvResult returns the CvvResult field if non-nil, zero value otherwise.

### GetCvvResultOk

`func (o *CardVerificationResults) GetCvvResultOk() (*string, bool)`

GetCvvResultOk returns a tuple with the CvvResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvResult

`func (o *CardVerificationResults) SetCvvResult(v string)`

SetCvvResult sets CvvResult field to given value.

### HasCvvResult

`func (o *CardVerificationResults) HasCvvResult() bool`

HasCvvResult returns a boolean if a field has been set.

### SetCvvResultNil

`func (o *CardVerificationResults) SetCvvResultNil(b bool)`

 SetCvvResultNil sets the value for CvvResult to be an explicit nil

### UnsetCvvResult
`func (o *CardVerificationResults) UnsetCvvResult()`

UnsetCvvResult ensures that no value is present for CvvResult, not even an explicit nil
### GetAddressVerificationResult

`func (o *CardVerificationResults) GetAddressVerificationResult() string`

GetAddressVerificationResult returns the AddressVerificationResult field if non-nil, zero value otherwise.

### GetAddressVerificationResultOk

`func (o *CardVerificationResults) GetAddressVerificationResultOk() (*string, bool)`

GetAddressVerificationResultOk returns a tuple with the AddressVerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressVerificationResult

`func (o *CardVerificationResults) SetAddressVerificationResult(v string)`

SetAddressVerificationResult sets AddressVerificationResult field to given value.

### HasAddressVerificationResult

`func (o *CardVerificationResults) HasAddressVerificationResult() bool`

HasAddressVerificationResult returns a boolean if a field has been set.

### SetAddressVerificationResultNil

`func (o *CardVerificationResults) SetAddressVerificationResultNil(b bool)`

 SetAddressVerificationResultNil sets the value for AddressVerificationResult to be an explicit nil

### UnsetAddressVerificationResult
`func (o *CardVerificationResults) UnsetAddressVerificationResult()`

UnsetAddressVerificationResult ensures that no value is present for AddressVerificationResult, not even an explicit nil

[[Back to README]](../../README.md)



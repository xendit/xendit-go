# PaymentMethodUpdateParameters


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Description** | Pointer to **string** |  |  |  |
| **ReferenceId** | Pointer to **string** |  |  |  |
| **Reusability** | Pointer to [**PaymentMethodReusability**](PaymentMethodReusability.md) |  |  |  |
| **Status** | Pointer to [**PaymentMethodStatus**](PaymentMethodStatus.md) |  |  |  |
| **OverTheCounter** | Pointer to [**OverTheCounterUpdateParameters**](OverTheCounterUpdateParameters.md) |  |  |  |
| **VirtualAccount** | Pointer to [**VirtualAccountUpdateParameters**](VirtualAccountUpdateParameters.md) |  |  |  |

## Methods

### NewPaymentMethodUpdateParameters

`func NewPaymentMethodUpdateParameters() *PaymentMethodUpdateParameters`

NewPaymentMethodUpdateParameters instantiates a new PaymentMethodUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodUpdateParametersWithDefaults

`func NewPaymentMethodUpdateParametersWithDefaults() *PaymentMethodUpdateParameters`

NewPaymentMethodUpdateParametersWithDefaults instantiates a new PaymentMethodUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PaymentMethodUpdateParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentMethodUpdateParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentMethodUpdateParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentMethodUpdateParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReferenceId

`func (o *PaymentMethodUpdateParameters) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PaymentMethodUpdateParameters) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PaymentMethodUpdateParameters) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PaymentMethodUpdateParameters) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetReusability

`func (o *PaymentMethodUpdateParameters) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *PaymentMethodUpdateParameters) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *PaymentMethodUpdateParameters) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.

### HasReusability

`func (o *PaymentMethodUpdateParameters) HasReusability() bool`

HasReusability returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentMethodUpdateParameters) GetStatus() PaymentMethodStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentMethodUpdateParameters) GetStatusOk() (*PaymentMethodStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentMethodUpdateParameters) SetStatus(v PaymentMethodStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentMethodUpdateParameters) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOverTheCounter

`func (o *PaymentMethodUpdateParameters) GetOverTheCounter() OverTheCounterUpdateParameters`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *PaymentMethodUpdateParameters) GetOverTheCounterOk() (*OverTheCounterUpdateParameters, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *PaymentMethodUpdateParameters) SetOverTheCounter(v OverTheCounterUpdateParameters)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *PaymentMethodUpdateParameters) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### GetVirtualAccount

`func (o *PaymentMethodUpdateParameters) GetVirtualAccount() VirtualAccountUpdateParameters`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *PaymentMethodUpdateParameters) GetVirtualAccountOk() (*VirtualAccountUpdateParameters, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *PaymentMethodUpdateParameters) SetVirtualAccount(v VirtualAccountUpdateParameters)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *PaymentMethodUpdateParameters) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.


[[Back to README]](../../README.md)



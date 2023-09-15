# VirtualAccountChannelPropertiesPatch

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **ExpiresAt** | Pointer to **time.Time** | The date and time in ISO 8601 UTC+0 when the virtual account number will be expired. Default: The default expiration date will be 31 years from creation date. | [optional]  |
| **SuggestedAmount** | Pointer to **float64** | The suggested amount you want to assign. Note: Suggested amounts is the amounts that can see as a suggestion, but user can still put any numbers (only supported for Mandiri and BRI) | [optional]  |

## Methods

### NewVirtualAccountChannelPropertiesPatch

`func NewVirtualAccountChannelPropertiesPatch() *VirtualAccountChannelPropertiesPatch`

NewVirtualAccountChannelPropertiesPatch instantiates a new VirtualAccountChannelPropertiesPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAccountChannelPropertiesPatchWithDefaults

`func NewVirtualAccountChannelPropertiesPatchWithDefaults() *VirtualAccountChannelPropertiesPatch`

NewVirtualAccountChannelPropertiesPatchWithDefaults instantiates a new VirtualAccountChannelPropertiesPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *VirtualAccountChannelPropertiesPatch) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *VirtualAccountChannelPropertiesPatch) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *VirtualAccountChannelPropertiesPatch) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *VirtualAccountChannelPropertiesPatch) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSuggestedAmount

`func (o *VirtualAccountChannelPropertiesPatch) GetSuggestedAmount() float64`

GetSuggestedAmount returns the SuggestedAmount field if non-nil, zero value otherwise.

### GetSuggestedAmountOk

`func (o *VirtualAccountChannelPropertiesPatch) GetSuggestedAmountOk() (*float64, bool)`

GetSuggestedAmountOk returns a tuple with the SuggestedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedAmount

`func (o *VirtualAccountChannelPropertiesPatch) SetSuggestedAmount(v float64)`

SetSuggestedAmount sets SuggestedAmount field to given value.

### HasSuggestedAmount

`func (o *VirtualAccountChannelPropertiesPatch) HasSuggestedAmount() bool`

HasSuggestedAmount returns a boolean if a field has been set.


[[Back to README]](../../README.md)



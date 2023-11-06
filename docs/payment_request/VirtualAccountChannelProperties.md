# VirtualAccountChannelProperties
Virtual Account Channel Properties

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **CustomerName** | **string** | ☑️ | Name of customer. |  |
| **VirtualAccountNumber** | Pointer to **string** |  | You can assign specific Virtual Account number using this parameter. If you do not send one, one will be picked at random. Make sure the number you specify is within your Virtual Account range. |  |
| **ExpiresAt** | Pointer to **time.Time** |  | The date and time in ISO 8601 UTC+0 when the virtual account number will be expired. Default: The default expiration date will be 31 years from creation date. |  |
| **SuggestedAmount** | Pointer to **float64** |  | The suggested amount you want to assign. Note: Suggested amounts is the amounts that can see as a suggestion, but user can still put any numbers (only supported for Mandiri and BRI) |  |

## Methods

### NewVirtualAccountChannelProperties

`func NewVirtualAccountChannelProperties(customerName string, ) *VirtualAccountChannelProperties`

NewVirtualAccountChannelProperties instantiates a new VirtualAccountChannelProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAccountChannelPropertiesWithDefaults

`func NewVirtualAccountChannelPropertiesWithDefaults() *VirtualAccountChannelProperties`

NewVirtualAccountChannelPropertiesWithDefaults instantiates a new VirtualAccountChannelProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerName

`func (o *VirtualAccountChannelProperties) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *VirtualAccountChannelProperties) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *VirtualAccountChannelProperties) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.


### GetVirtualAccountNumber

`func (o *VirtualAccountChannelProperties) GetVirtualAccountNumber() string`

GetVirtualAccountNumber returns the VirtualAccountNumber field if non-nil, zero value otherwise.

### GetVirtualAccountNumberOk

`func (o *VirtualAccountChannelProperties) GetVirtualAccountNumberOk() (*string, bool)`

GetVirtualAccountNumberOk returns a tuple with the VirtualAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccountNumber

`func (o *VirtualAccountChannelProperties) SetVirtualAccountNumber(v string)`

SetVirtualAccountNumber sets VirtualAccountNumber field to given value.

### HasVirtualAccountNumber

`func (o *VirtualAccountChannelProperties) HasVirtualAccountNumber() bool`

HasVirtualAccountNumber returns a boolean if a field has been set.

### GetExpiresAt

`func (o *VirtualAccountChannelProperties) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *VirtualAccountChannelProperties) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *VirtualAccountChannelProperties) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *VirtualAccountChannelProperties) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSuggestedAmount

`func (o *VirtualAccountChannelProperties) GetSuggestedAmount() float64`

GetSuggestedAmount returns the SuggestedAmount field if non-nil, zero value otherwise.

### GetSuggestedAmountOk

`func (o *VirtualAccountChannelProperties) GetSuggestedAmountOk() (*float64, bool)`

GetSuggestedAmountOk returns a tuple with the SuggestedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedAmount

`func (o *VirtualAccountChannelProperties) SetSuggestedAmount(v float64)`

SetSuggestedAmount sets SuggestedAmount field to given value.

### HasSuggestedAmount

`func (o *VirtualAccountChannelProperties) HasSuggestedAmount() bool`

HasSuggestedAmount returns a boolean if a field has been set.


[[Back to README]](../../README.md)



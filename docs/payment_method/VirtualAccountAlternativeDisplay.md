# VirtualAccountAlternativeDisplay
Alternative Display Object

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Type** | Pointer to **string** |  | Type of the alternative display |  |
| **Data** | Pointer to **string** |  | Data payload of the given alternative display |  |

## Methods

### NewVirtualAccountAlternativeDisplay

`func NewVirtualAccountAlternativeDisplay() *VirtualAccountAlternativeDisplay`

NewVirtualAccountAlternativeDisplay instantiates a new VirtualAccountAlternativeDisplay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAccountAlternativeDisplayWithDefaults

`func NewVirtualAccountAlternativeDisplayWithDefaults() *VirtualAccountAlternativeDisplay`

NewVirtualAccountAlternativeDisplayWithDefaults instantiates a new VirtualAccountAlternativeDisplay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VirtualAccountAlternativeDisplay) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualAccountAlternativeDisplay) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualAccountAlternativeDisplay) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualAccountAlternativeDisplay) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *VirtualAccountAlternativeDisplay) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VirtualAccountAlternativeDisplay) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VirtualAccountAlternativeDisplay) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *VirtualAccountAlternativeDisplay) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../../README.md)



# AccountCard


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **TokenId** | Pointer to **string** |  | The token id returned in tokenisation |  |

## Methods

### NewAccountCard

`func NewAccountCard() *AccountCard`

NewAccountCard instantiates a new AccountCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCardWithDefaults

`func NewAccountCardWithDefaults() *AccountCard`

NewAccountCardWithDefaults instantiates a new AccountCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *AccountCard) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AccountCard) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AccountCard) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *AccountCard) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to README]](../../README.md)



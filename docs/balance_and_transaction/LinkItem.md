# LinkItem


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Href** | **string** | ☑️ | URI of target, this will be to the next link. |  |
| **Rel** | **string** | ☑️ | The relationship between source and target. The value will be &#x60;next&#x60;. |  |
| **Method** | **string** | ☑️ | The HTTP method, the value will be &#x60;GET&#x60;. |  |

## Methods

### NewLinkItem

`func NewLinkItem(href string, rel string, method string, ) *LinkItem`

NewLinkItem instantiates a new LinkItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkItemWithDefaults

`func NewLinkItemWithDefaults() *LinkItem`

NewLinkItemWithDefaults instantiates a new LinkItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *LinkItem) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LinkItem) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LinkItem) SetHref(v string)`

SetHref sets Href field to given value.


### GetRel

`func (o *LinkItem) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *LinkItem) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *LinkItem) SetRel(v string)`

SetRel sets Rel field to given value.


### GetMethod

`func (o *LinkItem) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LinkItem) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LinkItem) SetMethod(v string)`

SetMethod sets Method field to given value.



[[Back to README]](../../README.md)



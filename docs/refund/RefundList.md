# RefundList


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Data** | [**Refund[]**](Refund.md) | ☑️ |  |  |
| **HasMore** | Pointer to **bool** |  |  |  |

## Methods

### NewRefundList

`func NewRefundList(data []Refund, ) *RefundList`

NewRefundList instantiates a new RefundList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundListWithDefaults

`func NewRefundListWithDefaults() *RefundList`

NewRefundListWithDefaults instantiates a new RefundList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RefundList) GetData() []Refund`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RefundList) GetDataOk() (*[]Refund, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RefundList) SetData(v []Refund)`

SetData sets Data field to given value.


### GetHasMore

`func (o *RefundList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *RefundList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *RefundList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *RefundList) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to README]](../../README.md)



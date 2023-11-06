# PaymentMethodList


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Data** | [**PaymentMethod[]**](PaymentMethod.md) | ☑️ |  |  |
| **HasMore** | Pointer to **bool** |  |  |  |

## Methods

### NewPaymentMethodList

`func NewPaymentMethodList(data []PaymentMethod, ) *PaymentMethodList`

NewPaymentMethodList instantiates a new PaymentMethodList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodListWithDefaults

`func NewPaymentMethodListWithDefaults() *PaymentMethodList`

NewPaymentMethodListWithDefaults instantiates a new PaymentMethodList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentMethodList) GetData() []PaymentMethod`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentMethodList) GetDataOk() (*[]PaymentMethod, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentMethodList) SetData(v []PaymentMethod)`

SetData sets Data field to given value.


### GetHasMore

`func (o *PaymentMethodList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PaymentMethodList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PaymentMethodList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *PaymentMethodList) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to README]](../../README.md)



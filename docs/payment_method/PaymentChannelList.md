# PaymentChannelList

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Data** | Pointer to [**PaymentChannel[]**](PaymentChannel.md) | Array of resources that match the provided filters | [optional]  |
| **Links** | Pointer to [**PaymentChannelListLinksInner[]**](PaymentChannelListLinksInner.md) | Array of objects that can be used to assist on navigating through the data | [optional]  |
| **HasMore** | Pointer to **bool** | Indicates whether there are more items in the list | [optional]  |

## Methods

### NewPaymentChannelList

`func NewPaymentChannelList() *PaymentChannelList`

NewPaymentChannelList instantiates a new PaymentChannelList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentChannelListWithDefaults

`func NewPaymentChannelListWithDefaults() *PaymentChannelList`

NewPaymentChannelListWithDefaults instantiates a new PaymentChannelList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentChannelList) GetData() []PaymentChannel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentChannelList) GetDataOk() (*[]PaymentChannel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentChannelList) SetData(v []PaymentChannel)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentChannelList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *PaymentChannelList) GetLinks() []PaymentChannelListLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaymentChannelList) GetLinksOk() (*[]PaymentChannelListLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaymentChannelList) SetLinks(v []PaymentChannelListLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaymentChannelList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetHasMore

`func (o *PaymentChannelList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PaymentChannelList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PaymentChannelList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *PaymentChannelList) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to README]](../../README.md)



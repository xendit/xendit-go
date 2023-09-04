# TransactionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | Indicates whether there are more items to be queried with &#x60;after_id&#x60; of the last item from the current result. Use the &#x60;links&#x60; to follow to the next result. | 
**Links** | Pointer to [**[]LinkItem**](LinkItem.md) | The links to the next page based on LinkItem if there is next result. | [optional] 
**Data** | [**[]TransactionResponse**](TransactionResponse.md) |  | 

## Methods

### NewTransactionsResponse

`func NewTransactionsResponse(hasMore bool, data []TransactionResponse, ) *TransactionsResponse`

NewTransactionsResponse instantiates a new TransactionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsResponseWithDefaults

`func NewTransactionsResponseWithDefaults() *TransactionsResponse`

NewTransactionsResponseWithDefaults instantiates a new TransactionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *TransactionsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *TransactionsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *TransactionsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetLinks

`func (o *TransactionsResponse) GetLinks() []LinkItem`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TransactionsResponse) GetLinksOk() (*[]LinkItem, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TransactionsResponse) SetLinks(v []LinkItem)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TransactionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *TransactionsResponse) GetData() []TransactionResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionsResponse) GetDataOk() (*[]TransactionResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionsResponse) SetData(v []TransactionResponse)`

SetData sets Data field to given value.



[[Back to README]](../../README.md)



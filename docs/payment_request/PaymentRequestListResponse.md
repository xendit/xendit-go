# PaymentRequestListResponse

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Data** | [**PaymentRequest[]**](PaymentRequest.md) |  |  |
| **HasMore** | **bool** |  |  |

## Methods

### NewPaymentRequestListResponse

`func NewPaymentRequestListResponse(data []PaymentRequest, hasMore bool, ) *PaymentRequestListResponse`

NewPaymentRequestListResponse instantiates a new PaymentRequestListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestListResponseWithDefaults

`func NewPaymentRequestListResponseWithDefaults() *PaymentRequestListResponse`

NewPaymentRequestListResponseWithDefaults instantiates a new PaymentRequestListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentRequestListResponse) GetData() []PaymentRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentRequestListResponse) GetDataOk() (*[]PaymentRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentRequestListResponse) SetData(v []PaymentRequest)`

SetData sets Data field to given value.


### GetHasMore

`func (o *PaymentRequestListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PaymentRequestListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PaymentRequestListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to README]](../../README.md)



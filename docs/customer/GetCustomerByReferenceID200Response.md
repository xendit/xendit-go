# GetCustomerByReferenceID200Response

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **HasMore** | Pointer to **bool** |  | [optional]  |
| **Data** | Pointer to [**Customer[]**](Customer.md) |  | [optional]  |

## Methods

### NewGetCustomerByReferenceID200Response

`func NewGetCustomerByReferenceID200Response() *GetCustomerByReferenceID200Response`

NewGetCustomerByReferenceID200Response instantiates a new GetCustomerByReferenceID200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerByReferenceID200ResponseWithDefaults

`func NewGetCustomerByReferenceID200ResponseWithDefaults() *GetCustomerByReferenceID200Response`

NewGetCustomerByReferenceID200ResponseWithDefaults instantiates a new GetCustomerByReferenceID200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetCustomerByReferenceID200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCustomerByReferenceID200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCustomerByReferenceID200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetCustomerByReferenceID200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetData

`func (o *GetCustomerByReferenceID200Response) GetData() []Customer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCustomerByReferenceID200Response) GetDataOk() (*[]Customer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCustomerByReferenceID200Response) SetData(v []Customer)`

SetData sets Data field to given value.

### HasData

`func (o *GetCustomerByReferenceID200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../../README.md)



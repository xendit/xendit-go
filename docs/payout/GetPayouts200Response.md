# GetPayouts200Response

## Properties

| Name | Type | Description | Notes |
| ------------ | ------------- | ------------- | ------------- |
| **Data** | Pointer to [**GetPayouts200ResponseDataInner[]**](GetPayouts200ResponseDataInner.md) |  | [optional]  |
| **HasMore** | Pointer to **bool** |  | [optional]  |
| **Links** | Pointer to [**GetPayouts200ResponseLinks**](GetPayouts200ResponseLinks.md) |  | [optional]  |

## Methods

### NewGetPayouts200Response

`func NewGetPayouts200Response() *GetPayouts200Response`

NewGetPayouts200Response instantiates a new GetPayouts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayouts200ResponseWithDefaults

`func NewGetPayouts200ResponseWithDefaults() *GetPayouts200Response`

NewGetPayouts200ResponseWithDefaults instantiates a new GetPayouts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPayouts200Response) GetData() []GetPayouts200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPayouts200Response) GetDataOk() (*[]GetPayouts200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPayouts200Response) SetData(v []GetPayouts200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetPayouts200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetHasMore

`func (o *GetPayouts200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetPayouts200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetPayouts200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetPayouts200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetLinks

`func (o *GetPayouts200Response) GetLinks() GetPayouts200ResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetPayouts200Response) GetLinksOk() (*GetPayouts200ResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetPayouts200Response) SetLinks(v GetPayouts200ResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetPayouts200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to README]](../../README.md)



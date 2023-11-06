# CaptureListResponse


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Data** | [**Capture[]**](Capture.md) | ☑️ |  |  |
| **HasMore** | **bool** | ☑️ |  |  |

## Methods

### NewCaptureListResponse

`func NewCaptureListResponse(data []Capture, hasMore bool, ) *CaptureListResponse`

NewCaptureListResponse instantiates a new CaptureListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureListResponseWithDefaults

`func NewCaptureListResponseWithDefaults() *CaptureListResponse`

NewCaptureListResponseWithDefaults instantiates a new CaptureListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CaptureListResponse) GetData() []Capture`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CaptureListResponse) GetDataOk() (*[]Capture, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CaptureListResponse) SetData(v []Capture)`

SetData sets Data field to given value.


### GetHasMore

`func (o *CaptureListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *CaptureListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *CaptureListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to README]](../../README.md)



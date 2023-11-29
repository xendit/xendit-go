# InvoiceCallbackItem
An object representing an item within an invoice.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Name** | **string** | ☑️ | The name of the item. |  |
| **Price** | **float32** | ☑️ | The price of the item. |  |
| **Quantity** | **float32** | ☑️ | The quantity of the item. Must be greater than or equal to 0. |  |
| **Url** | Pointer to **string** |  | The URL associated with the item. |  |
| **Category** | Pointer to **string** |  | The category of the item. |  |

## Methods

### NewInvoiceCallbackItem

`func NewInvoiceCallbackItem(name string, price float32, quantity float32, ) *InvoiceCallbackItem`

NewInvoiceCallbackItem instantiates a new InvoiceCallbackItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCallbackItemWithDefaults

`func NewInvoiceCallbackItemWithDefaults() *InvoiceCallbackItem`

NewInvoiceCallbackItemWithDefaults instantiates a new InvoiceCallbackItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InvoiceCallbackItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceCallbackItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceCallbackItem) SetName(v string)`

SetName sets Name field to given value.


### GetPrice

`func (o *InvoiceCallbackItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceCallbackItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceCallbackItem) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetQuantity

`func (o *InvoiceCallbackItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceCallbackItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceCallbackItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUrl

`func (o *InvoiceCallbackItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InvoiceCallbackItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InvoiceCallbackItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InvoiceCallbackItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCategory

`func (o *InvoiceCallbackItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InvoiceCallbackItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InvoiceCallbackItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InvoiceCallbackItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to README]](../../README.md)



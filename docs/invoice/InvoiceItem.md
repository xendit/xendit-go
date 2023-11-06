# InvoiceItem
An object representing an item within an invoice.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Name** | **string** | ☑️ | The name of the item. |  |
| **Price** | **float32** | ☑️ | The price of the item. |  |
| **Quantity** | **float32** | ☑️ | The quantity of the item. Must be greater than or equal to 0. |  |
| **ReferenceId** | Pointer to **string** |  | The reference ID of the item. |  |
| **Url** | Pointer to **string** |  | The URL associated with the item. |  |
| **Category** | Pointer to **string** |  | The category of the item. |  |

## Methods

### NewInvoiceItem

`func NewInvoiceItem(name string, price float32, quantity float32, ) *InvoiceItem`

NewInvoiceItem instantiates a new InvoiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceItemWithDefaults

`func NewInvoiceItemWithDefaults() *InvoiceItem`

NewInvoiceItemWithDefaults instantiates a new InvoiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InvoiceItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceItem) SetName(v string)`

SetName sets Name field to given value.


### GetPrice

`func (o *InvoiceItem) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoiceItem) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoiceItem) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetQuantity

`func (o *InvoiceItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetReferenceId

`func (o *InvoiceItem) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InvoiceItem) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InvoiceItem) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *InvoiceItem) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetUrl

`func (o *InvoiceItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InvoiceItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InvoiceItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InvoiceItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCategory

`func (o *InvoiceItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InvoiceItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InvoiceItem) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InvoiceItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to README]](../../README.md)



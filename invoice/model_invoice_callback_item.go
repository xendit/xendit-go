/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.8.7
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v6/utils"
)

// checks if the InvoiceCallbackItem type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InvoiceCallbackItem{}

// InvoiceCallbackItem An object representing an item within an invoice.
type InvoiceCallbackItem struct {
	// The name of the item.
	Name string `json:"name"`
	// The price of the item.
	Price float32 `json:"price"`
	// The quantity of the item. Must be greater than or equal to 0.
	Quantity float32 `json:"quantity"`
	// The URL associated with the item.
	Url *string `json:"url,omitempty"`
	// The category of the item.
	Category *string `json:"category,omitempty"`
}

// NewInvoiceCallbackItem instantiates a new InvoiceCallbackItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceCallbackItem(name string, price float32, quantity float32) *InvoiceCallbackItem {
	this := InvoiceCallbackItem{}
	this.Name = name
	this.Price = price
	this.Quantity = quantity
	return &this
}

// NewInvoiceCallbackItemWithDefaults instantiates a new InvoiceCallbackItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceCallbackItemWithDefaults() *InvoiceCallbackItem {
	this := InvoiceCallbackItem{}
	return &this
}

// GetName returns the Name field value
func (o *InvoiceCallbackItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallbackItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InvoiceCallbackItem) SetName(v string) {
	o.Name = v
}

// GetPrice returns the Price field value
func (o *InvoiceCallbackItem) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallbackItem) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *InvoiceCallbackItem) SetPrice(v float32) {
	o.Price = v
}

// GetQuantity returns the Quantity field value
func (o *InvoiceCallbackItem) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *InvoiceCallbackItem) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *InvoiceCallbackItem) SetQuantity(v float32) {
	o.Quantity = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InvoiceCallbackItem) GetUrl() string {
	if o == nil || utils.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallbackItem) GetUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InvoiceCallbackItem) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InvoiceCallbackItem) SetUrl(v string) {
	o.Url = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InvoiceCallbackItem) GetCategory() string {
	if o == nil || utils.IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceCallbackItem) GetCategoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InvoiceCallbackItem) HasCategory() bool {
	if o != nil && !utils.IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InvoiceCallbackItem) SetCategory(v string) {
	o.Category = &v
}

func (o InvoiceCallbackItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceCallbackItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["price"] = o.Price
	toSerialize["quantity"] = o.Quantity
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !utils.IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	return toSerialize, nil
}

type NullableInvoiceCallbackItem struct {
	value *InvoiceCallbackItem
	isSet bool
}

func (v NullableInvoiceCallbackItem) Get() *InvoiceCallbackItem {
	return v.value
}

func (v *NullableInvoiceCallbackItem) Set(val *InvoiceCallbackItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceCallbackItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceCallbackItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceCallbackItem(val *InvoiceCallbackItem) *NullableInvoiceCallbackItem {
	return &NullableInvoiceCallbackItem{value: val, isSet: true}
}

func (v NullableInvoiceCallbackItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceCallbackItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



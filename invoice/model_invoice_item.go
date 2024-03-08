/*
xendit-invoice-service

xendit-invoice-service descriptions

API version: 1.7.6
*/


package invoice

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v4/utils"
)

// checks if the InvoiceItem type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InvoiceItem{}

// InvoiceItem An object representing an item within an invoice.
type InvoiceItem struct {
	// The name of the item.
	Name string `json:"name"`
	// The price of the item.
	Price float32 `json:"price"`
	// The quantity of the item. Must be greater than or equal to 0.
	Quantity float32 `json:"quantity"`
	// The reference ID of the item.
	ReferenceId *string `json:"reference_id,omitempty"`
	// The URL associated with the item.
	Url *string `json:"url,omitempty"`
	// The category of the item.
	Category *string `json:"category,omitempty"`
}

// NewInvoiceItem instantiates a new InvoiceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceItem(name string, price float32, quantity float32) *InvoiceItem {
	this := InvoiceItem{}
	this.Name = name
	this.Price = price
	this.Quantity = quantity
	return &this
}

// NewInvoiceItemWithDefaults instantiates a new InvoiceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceItemWithDefaults() *InvoiceItem {
	this := InvoiceItem{}
	return &this
}

// GetName returns the Name field value
func (o *InvoiceItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InvoiceItem) SetName(v string) {
	o.Name = v
}

// GetPrice returns the Price field value
func (o *InvoiceItem) GetPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *InvoiceItem) SetPrice(v float32) {
	o.Price = v
}

// GetQuantity returns the Quantity field value
func (o *InvoiceItem) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *InvoiceItem) SetQuantity(v float32) {
	o.Quantity = v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *InvoiceItem) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetReferenceIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *InvoiceItem) HasReferenceId() bool {
	if o != nil && !utils.IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *InvoiceItem) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InvoiceItem) GetUrl() string {
	if o == nil || utils.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InvoiceItem) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InvoiceItem) SetUrl(v string) {
	o.Url = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InvoiceItem) GetCategory() string {
	if o == nil || utils.IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetCategoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InvoiceItem) HasCategory() bool {
	if o != nil && !utils.IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InvoiceItem) SetCategory(v string) {
	o.Category = &v
}

func (o InvoiceItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["price"] = o.Price
	toSerialize["quantity"] = o.Quantity
	if !utils.IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !utils.IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	return toSerialize, nil
}

type NullableInvoiceItem struct {
	value *InvoiceItem
	isSet bool
}

func (v NullableInvoiceItem) Get() *InvoiceItem {
	return v.value
}

func (v *NullableInvoiceItem) Set(val *InvoiceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceItem(val *InvoiceItem) *NullableInvoiceItem {
	return &NullableInvoiceItem{value: val, isSet: true}
}

func (v NullableInvoiceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



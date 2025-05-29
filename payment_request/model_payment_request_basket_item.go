/*
Payment Requests

This API is used for Payment Requests

API version: 1.70.0
*/


package payment_request

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v7/utils"
)

// checks if the PaymentRequestBasketItem type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentRequestBasketItem{}

// PaymentRequestBasketItem struct for PaymentRequestBasketItem
type PaymentRequestBasketItem struct {
	ReferenceId *string `json:"reference_id,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	Category string `json:"category"`
	SubCategory *string `json:"sub_category,omitempty"`
	Currency string `json:"currency"`
	Quantity float64 `json:"quantity"`
	Price float64 `json:"price"`
	PayerChargedCurrency *string `json:"payer_charged_currency,omitempty"`
	PayerChargedPrice *float64 `json:"payer_charged_price,omitempty"`
	Url *string `json:"url,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPaymentRequestBasketItem instantiates a new PaymentRequestBasketItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestBasketItem(name string, category string, currency string, quantity float64, price float64) *PaymentRequestBasketItem {
	this := PaymentRequestBasketItem{}
	this.Name = name
	this.Category = category
	this.Currency = currency
	this.Quantity = quantity
	this.Price = price
	return &this
}

// NewPaymentRequestBasketItemWithDefaults instantiates a new PaymentRequestBasketItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestBasketItemWithDefaults() *PaymentRequestBasketItem {
	this := PaymentRequestBasketItem{}
	return &this
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *PaymentRequestBasketItem) GetReferenceId() string {
	if o == nil || utils.IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetReferenceIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *PaymentRequestBasketItem) HasReferenceId() bool {
	if o != nil && !utils.IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *PaymentRequestBasketItem) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetName returns the Name field value
func (o *PaymentRequestBasketItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PaymentRequestBasketItem) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentRequestBasketItem) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentRequestBasketItem) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentRequestBasketItem) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentRequestBasketItem) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentRequestBasketItem) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentRequestBasketItem) SetType(v string) {
	o.Type = &v
}

// GetCategory returns the Category field value
func (o *PaymentRequestBasketItem) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *PaymentRequestBasketItem) SetCategory(v string) {
	o.Category = v
}

// GetSubCategory returns the SubCategory field value if set, zero value otherwise.
func (o *PaymentRequestBasketItem) GetSubCategory() string {
	if o == nil || utils.IsNil(o.SubCategory) {
		var ret string
		return ret
	}
	return *o.SubCategory
}

// GetSubCategoryOk returns a tuple with the SubCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetSubCategoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SubCategory) {
		return nil, false
	}
	return o.SubCategory, true
}

// HasSubCategory returns a boolean if a field has been set.
func (o *PaymentRequestBasketItem) HasSubCategory() bool {
	if o != nil && !utils.IsNil(o.SubCategory) {
		return true
	}

	return false
}

// SetSubCategory gets a reference to the given string and assigns it to the SubCategory field.
func (o *PaymentRequestBasketItem) SetSubCategory(v string) {
	o.SubCategory = &v
}

// GetCurrency returns the Currency field value
func (o *PaymentRequestBasketItem) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentRequestBasketItem) SetCurrency(v string) {
	o.Currency = v
}

// GetQuantity returns the Quantity field value
func (o *PaymentRequestBasketItem) GetQuantity() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetQuantityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *PaymentRequestBasketItem) SetQuantity(v float64) {
	o.Quantity = v
}

// GetPrice returns the Price field value
func (o *PaymentRequestBasketItem) GetPrice() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *PaymentRequestBasketItem) SetPrice(v float64) {
	o.Price = v
}

// GetPayerChargedCurrency returns the PayerChargedCurrency field value if set, zero value otherwise.
func (o *PaymentRequestBasketItem) GetPayerChargedCurrency() string {
	if o == nil || utils.IsNil(o.PayerChargedCurrency) {
		var ret string
		return ret
	}
	return *o.PayerChargedCurrency
}

// GetPayerChargedCurrencyOk returns a tuple with the PayerChargedCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetPayerChargedCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PayerChargedCurrency) {
		return nil, false
	}
	return o.PayerChargedCurrency, true
}

// HasPayerChargedCurrency returns a boolean if a field has been set.
func (o *PaymentRequestBasketItem) HasPayerChargedCurrency() bool {
	if o != nil && !utils.IsNil(o.PayerChargedCurrency) {
		return true
	}

	return false
}

// SetPayerChargedCurrency gets a reference to the given string and assigns it to the PayerChargedCurrency field.
func (o *PaymentRequestBasketItem) SetPayerChargedCurrency(v string) {
	o.PayerChargedCurrency = &v
}

// GetPayerChargedPrice returns the PayerChargedPrice field value if set, zero value otherwise.
func (o *PaymentRequestBasketItem) GetPayerChargedPrice() float64 {
	if o == nil || utils.IsNil(o.PayerChargedPrice) {
		var ret float64
		return ret
	}
	return *o.PayerChargedPrice
}

// GetPayerChargedPriceOk returns a tuple with the PayerChargedPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetPayerChargedPriceOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.PayerChargedPrice) {
		return nil, false
	}
	return o.PayerChargedPrice, true
}

// HasPayerChargedPrice returns a boolean if a field has been set.
func (o *PaymentRequestBasketItem) HasPayerChargedPrice() bool {
	if o != nil && !utils.IsNil(o.PayerChargedPrice) {
		return true
	}

	return false
}

// SetPayerChargedPrice gets a reference to the given float64 and assigns it to the PayerChargedPrice field.
func (o *PaymentRequestBasketItem) SetPayerChargedPrice(v float64) {
	o.PayerChargedPrice = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PaymentRequestBasketItem) GetUrl() string {
	if o == nil || utils.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PaymentRequestBasketItem) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PaymentRequestBasketItem) SetUrl(v string) {
	o.Url = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PaymentRequestBasketItem) GetMetadata() map[string]interface{} {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentRequestBasketItem) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentRequestBasketItem) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PaymentRequestBasketItem) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PaymentRequestBasketItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentRequestBasketItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	toSerialize["name"] = o.Name
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["category"] = o.Category
	if !utils.IsNil(o.SubCategory) {
		toSerialize["sub_category"] = o.SubCategory
	}
	toSerialize["currency"] = o.Currency
	toSerialize["quantity"] = o.Quantity
	toSerialize["price"] = o.Price
	if !utils.IsNil(o.PayerChargedCurrency) {
		toSerialize["payer_charged_currency"] = o.PayerChargedCurrency
	}
	if !utils.IsNil(o.PayerChargedPrice) {
		toSerialize["payer_charged_price"] = o.PayerChargedPrice
	}
	if !utils.IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePaymentRequestBasketItem struct {
	value *PaymentRequestBasketItem
	isSet bool
}

func (v NullablePaymentRequestBasketItem) Get() *PaymentRequestBasketItem {
	return v.value
}

func (v *NullablePaymentRequestBasketItem) Set(val *PaymentRequestBasketItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestBasketItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestBasketItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestBasketItem(val *PaymentRequestBasketItem) *NullablePaymentRequestBasketItem {
	return &NullablePaymentRequestBasketItem{value: val, isSet: true}
}

func (v NullablePaymentRequestBasketItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestBasketItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



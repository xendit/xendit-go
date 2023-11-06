# RetailOutlet
An object representing retail outlet details for invoices.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **RetailOutletName** | [**RetailOutletName**](RetailOutletName.md) | ☑️ |  |  |
| **PaymentCode** | Pointer to **string** |  | The payment code. |  |
| **TransferAmount** | Pointer to **float32** |  | The transfer amount. |  |
| **MerchantName** | Pointer to **string** |  | The name of the merchant. |  |

## Methods

### NewRetailOutlet

`func NewRetailOutlet(retailOutletName RetailOutletName, ) *RetailOutlet`

NewRetailOutlet instantiates a new RetailOutlet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetailOutletWithDefaults

`func NewRetailOutletWithDefaults() *RetailOutlet`

NewRetailOutletWithDefaults instantiates a new RetailOutlet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetailOutletName

`func (o *RetailOutlet) GetRetailOutletName() RetailOutletName`

GetRetailOutletName returns the RetailOutletName field if non-nil, zero value otherwise.

### GetRetailOutletNameOk

`func (o *RetailOutlet) GetRetailOutletNameOk() (*RetailOutletName, bool)`

GetRetailOutletNameOk returns a tuple with the RetailOutletName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailOutletName

`func (o *RetailOutlet) SetRetailOutletName(v RetailOutletName)`

SetRetailOutletName sets RetailOutletName field to given value.


### GetPaymentCode

`func (o *RetailOutlet) GetPaymentCode() string`

GetPaymentCode returns the PaymentCode field if non-nil, zero value otherwise.

### GetPaymentCodeOk

`func (o *RetailOutlet) GetPaymentCodeOk() (*string, bool)`

GetPaymentCodeOk returns a tuple with the PaymentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCode

`func (o *RetailOutlet) SetPaymentCode(v string)`

SetPaymentCode sets PaymentCode field to given value.

### HasPaymentCode

`func (o *RetailOutlet) HasPaymentCode() bool`

HasPaymentCode returns a boolean if a field has been set.

### GetTransferAmount

`func (o *RetailOutlet) GetTransferAmount() float32`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *RetailOutlet) GetTransferAmountOk() (*float32, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *RetailOutlet) SetTransferAmount(v float32)`

SetTransferAmount sets TransferAmount field to given value.

### HasTransferAmount

`func (o *RetailOutlet) HasTransferAmount() bool`

HasTransferAmount returns a boolean if a field has been set.

### GetMerchantName

`func (o *RetailOutlet) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *RetailOutlet) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *RetailOutlet) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *RetailOutlet) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.


[[Back to README]](../../README.md)



# FeeResponse


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **XenditFee** | **float32** | ☑️ | Amount of the Xendit fee for this transaction. |  |
| **ValueAddedTax** | **float32** | ☑️ | Amount of the VAT for this transaction. |  |
| **XenditWithholdingTax** | Pointer to **float32** |  | Amount of the Xendit Withholding Tax for this transaction if applicable. See [Tax Documentation](https://docs.xendit.co/fees-and-vat#vat) for more information. |  |
| **ThirdPartyWithholdingTax** | Pointer to **float32** |  | Amount of the 3rd Party Withholding Tax for this transaction if applicable. 3rd party example: Bank  |  |
| **Status** | Pointer to **string** |  |  |  |

## Methods

### NewFeeResponse

`func NewFeeResponse(xenditFee float32, valueAddedTax float32, ) *FeeResponse`

NewFeeResponse instantiates a new FeeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeResponseWithDefaults

`func NewFeeResponseWithDefaults() *FeeResponse`

NewFeeResponseWithDefaults instantiates a new FeeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXenditFee

`func (o *FeeResponse) GetXenditFee() float32`

GetXenditFee returns the XenditFee field if non-nil, zero value otherwise.

### GetXenditFeeOk

`func (o *FeeResponse) GetXenditFeeOk() (*float32, bool)`

GetXenditFeeOk returns a tuple with the XenditFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXenditFee

`func (o *FeeResponse) SetXenditFee(v float32)`

SetXenditFee sets XenditFee field to given value.


### GetValueAddedTax

`func (o *FeeResponse) GetValueAddedTax() float32`

GetValueAddedTax returns the ValueAddedTax field if non-nil, zero value otherwise.

### GetValueAddedTaxOk

`func (o *FeeResponse) GetValueAddedTaxOk() (*float32, bool)`

GetValueAddedTaxOk returns a tuple with the ValueAddedTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAddedTax

`func (o *FeeResponse) SetValueAddedTax(v float32)`

SetValueAddedTax sets ValueAddedTax field to given value.


### GetXenditWithholdingTax

`func (o *FeeResponse) GetXenditWithholdingTax() float32`

GetXenditWithholdingTax returns the XenditWithholdingTax field if non-nil, zero value otherwise.

### GetXenditWithholdingTaxOk

`func (o *FeeResponse) GetXenditWithholdingTaxOk() (*float32, bool)`

GetXenditWithholdingTaxOk returns a tuple with the XenditWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXenditWithholdingTax

`func (o *FeeResponse) SetXenditWithholdingTax(v float32)`

SetXenditWithholdingTax sets XenditWithholdingTax field to given value.

### HasXenditWithholdingTax

`func (o *FeeResponse) HasXenditWithholdingTax() bool`

HasXenditWithholdingTax returns a boolean if a field has been set.

### GetThirdPartyWithholdingTax

`func (o *FeeResponse) GetThirdPartyWithholdingTax() float32`

GetThirdPartyWithholdingTax returns the ThirdPartyWithholdingTax field if non-nil, zero value otherwise.

### GetThirdPartyWithholdingTaxOk

`func (o *FeeResponse) GetThirdPartyWithholdingTaxOk() (*float32, bool)`

GetThirdPartyWithholdingTaxOk returns a tuple with the ThirdPartyWithholdingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyWithholdingTax

`func (o *FeeResponse) SetThirdPartyWithholdingTax(v float32)`

SetThirdPartyWithholdingTax sets ThirdPartyWithholdingTax field to given value.

### HasThirdPartyWithholdingTax

`func (o *FeeResponse) HasThirdPartyWithholdingTax() bool`

HasThirdPartyWithholdingTax returns a boolean if a field has been set.

### GetStatus

`func (o *FeeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FeeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FeeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FeeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to README]](../../README.md)



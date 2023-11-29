# PaymentDetails
An object representing payment details.

## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **ReceiptId** | Pointer to **string** |  | The unique identifier or reference ID associated with the payment receipt. |  |
| **Source** | Pointer to **string** |  | The source or method of payment. |  |

## Methods

### NewPaymentDetails

`func NewPaymentDetails() *PaymentDetails`

NewPaymentDetails instantiates a new PaymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentDetailsWithDefaults

`func NewPaymentDetailsWithDefaults() *PaymentDetails`

NewPaymentDetailsWithDefaults instantiates a new PaymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiptId

`func (o *PaymentDetails) GetReceiptId() string`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *PaymentDetails) GetReceiptIdOk() (*string, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *PaymentDetails) SetReceiptId(v string)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *PaymentDetails) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### GetSource

`func (o *PaymentDetails) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PaymentDetails) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PaymentDetails) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PaymentDetails) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to README]](../../README.md)



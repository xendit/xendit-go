# PaymentMethodParameters


## Properties
| Name | Type | Required | Description | Examples |
|------------|:-------------:|:-------------:|-------------|:-------------:|
| **Type** | [**PaymentMethodType**](PaymentMethodType.md) | ☑️ |  |  |
| **Country** | Pointer to **NullableString** |  |  |  |
| **Reusability** | [**PaymentMethodReusability**](PaymentMethodReusability.md) | ☑️ |  |  |
| **CustomerId** | Pointer to **NullableString** |  |  |  |
| **ReferenceId** | Pointer to **string** |  |  |  |
| **Description** | Pointer to **NullableString** |  |  |  |
| **Card** | Pointer to [**CardParameters**](CardParameters.md) |  |  |  |
| **DirectDebit** | Pointer to [**DirectDebitParameters**](DirectDebitParameters.md) |  |  |  |
| **Ewallet** | Pointer to [**EWalletParameters**](EWalletParameters.md) |  |  |  |
| **OverTheCounter** | Pointer to [**OverTheCounterParameters**](OverTheCounterParameters.md) |  |  |  |
| **VirtualAccount** | Pointer to [**VirtualAccountParameters**](VirtualAccountParameters.md) |  |  |  |
| **QrCode** | Pointer to [**QRCodeParameters**](QRCodeParameters.md) |  |  |  |
| **Metadata** | Pointer to **map[string]interface{}** |  |  |  |
| **BillingInformation** | Pointer to [**NullableBillingInformation**](BillingInformation.md) |  |  |  |

## Methods

### NewPaymentMethodParameters

`func NewPaymentMethodParameters(type_ PaymentMethodType, reusability PaymentMethodReusability, ) *PaymentMethodParameters`

NewPaymentMethodParameters instantiates a new PaymentMethodParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodParametersWithDefaults

`func NewPaymentMethodParametersWithDefaults() *PaymentMethodParameters`

NewPaymentMethodParametersWithDefaults instantiates a new PaymentMethodParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodParameters) GetType() PaymentMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodParameters) GetTypeOk() (*PaymentMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodParameters) SetType(v PaymentMethodType)`

SetType sets Type field to given value.


### GetCountry

`func (o *PaymentMethodParameters) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentMethodParameters) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentMethodParameters) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentMethodParameters) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PaymentMethodParameters) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PaymentMethodParameters) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetReusability

`func (o *PaymentMethodParameters) GetReusability() PaymentMethodReusability`

GetReusability returns the Reusability field if non-nil, zero value otherwise.

### GetReusabilityOk

`func (o *PaymentMethodParameters) GetReusabilityOk() (*PaymentMethodReusability, bool)`

GetReusabilityOk returns a tuple with the Reusability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusability

`func (o *PaymentMethodParameters) SetReusability(v PaymentMethodReusability)`

SetReusability sets Reusability field to given value.


### GetCustomerId

`func (o *PaymentMethodParameters) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentMethodParameters) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentMethodParameters) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *PaymentMethodParameters) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *PaymentMethodParameters) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *PaymentMethodParameters) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetReferenceId

`func (o *PaymentMethodParameters) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *PaymentMethodParameters) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *PaymentMethodParameters) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *PaymentMethodParameters) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentMethodParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentMethodParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentMethodParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentMethodParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentMethodParameters) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentMethodParameters) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCard

`func (o *PaymentMethodParameters) GetCard() CardParameters`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaymentMethodParameters) GetCardOk() (*CardParameters, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaymentMethodParameters) SetCard(v CardParameters)`

SetCard sets Card field to given value.

### HasCard

`func (o *PaymentMethodParameters) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetDirectDebit

`func (o *PaymentMethodParameters) GetDirectDebit() DirectDebitParameters`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *PaymentMethodParameters) GetDirectDebitOk() (*DirectDebitParameters, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *PaymentMethodParameters) SetDirectDebit(v DirectDebitParameters)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *PaymentMethodParameters) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.

### GetEwallet

`func (o *PaymentMethodParameters) GetEwallet() EWalletParameters`

GetEwallet returns the Ewallet field if non-nil, zero value otherwise.

### GetEwalletOk

`func (o *PaymentMethodParameters) GetEwalletOk() (*EWalletParameters, bool)`

GetEwalletOk returns a tuple with the Ewallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEwallet

`func (o *PaymentMethodParameters) SetEwallet(v EWalletParameters)`

SetEwallet sets Ewallet field to given value.

### HasEwallet

`func (o *PaymentMethodParameters) HasEwallet() bool`

HasEwallet returns a boolean if a field has been set.

### GetOverTheCounter

`func (o *PaymentMethodParameters) GetOverTheCounter() OverTheCounterParameters`

GetOverTheCounter returns the OverTheCounter field if non-nil, zero value otherwise.

### GetOverTheCounterOk

`func (o *PaymentMethodParameters) GetOverTheCounterOk() (*OverTheCounterParameters, bool)`

GetOverTheCounterOk returns a tuple with the OverTheCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverTheCounter

`func (o *PaymentMethodParameters) SetOverTheCounter(v OverTheCounterParameters)`

SetOverTheCounter sets OverTheCounter field to given value.

### HasOverTheCounter

`func (o *PaymentMethodParameters) HasOverTheCounter() bool`

HasOverTheCounter returns a boolean if a field has been set.

### GetVirtualAccount

`func (o *PaymentMethodParameters) GetVirtualAccount() VirtualAccountParameters`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *PaymentMethodParameters) GetVirtualAccountOk() (*VirtualAccountParameters, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *PaymentMethodParameters) SetVirtualAccount(v VirtualAccountParameters)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *PaymentMethodParameters) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.

### GetQrCode

`func (o *PaymentMethodParameters) GetQrCode() QRCodeParameters`

GetQrCode returns the QrCode field if non-nil, zero value otherwise.

### GetQrCodeOk

`func (o *PaymentMethodParameters) GetQrCodeOk() (*QRCodeParameters, bool)`

GetQrCodeOk returns a tuple with the QrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCode

`func (o *PaymentMethodParameters) SetQrCode(v QRCodeParameters)`

SetQrCode sets QrCode field to given value.

### HasQrCode

`func (o *PaymentMethodParameters) HasQrCode() bool`

HasQrCode returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentMethodParameters) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentMethodParameters) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentMethodParameters) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentMethodParameters) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PaymentMethodParameters) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PaymentMethodParameters) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetBillingInformation

`func (o *PaymentMethodParameters) GetBillingInformation() BillingInformation`

GetBillingInformation returns the BillingInformation field if non-nil, zero value otherwise.

### GetBillingInformationOk

`func (o *PaymentMethodParameters) GetBillingInformationOk() (*BillingInformation, bool)`

GetBillingInformationOk returns a tuple with the BillingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInformation

`func (o *PaymentMethodParameters) SetBillingInformation(v BillingInformation)`

SetBillingInformation sets BillingInformation field to given value.

### HasBillingInformation

`func (o *PaymentMethodParameters) HasBillingInformation() bool`

HasBillingInformation returns a boolean if a field has been set.

### SetBillingInformationNil

`func (o *PaymentMethodParameters) SetBillingInformationNil(b bool)`

 SetBillingInformationNil sets the value for BillingInformation to be an explicit nil

### UnsetBillingInformation
`func (o *PaymentMethodParameters) UnsetBillingInformation()`

UnsetBillingInformation ensures that no value is present for BillingInformation, not even an explicit nil

[[Back to README]](../../README.md)



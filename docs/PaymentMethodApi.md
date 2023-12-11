# PaymentMethodApi


You can use the APIs below to interface with Xendit's `PaymentMethodApi`.
To start using the API, you need to configure the secret key and initiate the client instance.

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v4"
)

func main() {
    xenditClient := xendit.NewClient("API-KEY")
}
```

All URIs are relative to *https://api.xendit.co*

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**CreatePaymentMethod**](PaymentMethodApi.md#createpaymentmethod-function) | **Post** /v2/payment_methods | Creates payment method |
| [**GetPaymentMethodByID**](PaymentMethodApi.md#getpaymentmethodbyid-function) | **Get** /v2/payment_methods/{paymentMethodId} | Get payment method by ID |
| [**GetPaymentsByPaymentMethodId**](PaymentMethodApi.md#getpaymentsbypaymentmethodid-function) | **Get** /v2/payment_methods/{paymentMethodId}/payments | Returns payments with matching PaymentMethodID. |
| [**PatchPaymentMethod**](PaymentMethodApi.md#patchpaymentmethod-function) | **Patch** /v2/payment_methods/{paymentMethodId} | Patch payment methods |
| [**GetAllPaymentMethods**](PaymentMethodApi.md#getallpaymentmethods-function) | **Get** /v2/payment_methods | Get all payment methods by filters |
| [**ExpirePaymentMethod**](PaymentMethodApi.md#expirepaymentmethod-function) | **Post** /v2/payment_methods/{paymentMethodId}/expire | Expires a payment method |
| [**AuthPaymentMethod**](PaymentMethodApi.md#authpaymentmethod-function) | **Post** /v2/payment_methods/{paymentMethodId}/auth | Validate a payment method&#39;s linking OTP |
| [**SimulatePayment**](PaymentMethodApi.md#simulatepayment-function) | **Post** /v2/payment_methods/{paymentMethodId}/payments/simulate | Makes payment with matching PaymentMethodID. |



## `CreatePaymentMethod()` Function

Creates payment method



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `CreatePaymentMethod` |
| Path Parameters  |  [CreatePaymentMethodPathParams](#request-parameters--CreatePaymentMethodPathParams)	 |
| Request Parameters  |  [CreatePaymentMethodRequestParams](#request-parameters--CreatePaymentMethodRequestParams)	 |
| Return Type  | [**PaymentMethod**](payment_method/PaymentMethod.md) |

### Path Parameters - CreatePaymentMethodPathParams
This endpoint does not need any path parameter.


### Request Parameters - CreatePaymentMethodRequestParams

Parameters that are passed through a pointer to a apiCreatePaymentMethodRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **forUserId** |**string**|  |  | 
|  **paymentMethodParameters** |[**PaymentMethodParameters**](payment_method/PaymentMethodParameters.md)|  |  | 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v4"
    payment_method "github.com/xendit/xendit-go/v4/payment_method"
)

func main() {
    
    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    paymentMethodParameters := *payment_method.NewPaymentMethodParameters(payment_method.PaymentMethodType("CARD"), payment_method.PaymentMethodReusability("MULTIPLE_USE")) // [OPTIONAL] | PaymentMethodParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.CreatePaymentMethod(context.Background()).
        ForUserId(forUserId).
        PaymentMethodParameters(paymentMethodParameters). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodApi.CreatePaymentMethod``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodApi.CreatePaymentMethod`: %v\n", resp)
}
```

## `GetPaymentMethodByID()` Function

Get payment method by ID



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetPaymentMethodByID` |
| Path Parameters  |  [GetPaymentMethodByIDPathParams](#request-parameters--GetPaymentMethodByIDPathParams)	 |
| Request Parameters  |  [GetPaymentMethodByIDRequestParams](#request-parameters--GetPaymentMethodByIDRequestParams)	 |
| Return Type  | [**PaymentMethod**](payment_method/PaymentMethod.md) |

### Path Parameters - GetPaymentMethodByIDPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentMethodId** | **string** |  | ☑️ |  | 

### Request Parameters - GetPaymentMethodByIDRequestParams

Parameters that are passed through a pointer to a apiGetPaymentMethodByIDRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **forUserId** |**string**|  |  | 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v4"
    payment_method "github.com/xendit/xendit-go/v4/payment_method"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.GetPaymentMethodByID(context.Background(), paymentMethodId).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodApi.GetPaymentMethodByID``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentMethodByID`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodApi.GetPaymentMethodByID`: %v\n", resp)
}
```

## `GetPaymentsByPaymentMethodId()` Function

Returns payments with matching PaymentMethodID.



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetPaymentsByPaymentMethodId` |
| Path Parameters  |  [GetPaymentsByPaymentMethodIdPathParams](#request-parameters--GetPaymentsByPaymentMethodIdPathParams)	 |
| Request Parameters  |  [GetPaymentsByPaymentMethodIdRequestParams](#request-parameters--GetPaymentsByPaymentMethodIdRequestParams)	 |
| Return Type  | **map[string]interface{}** |

### Path Parameters - GetPaymentsByPaymentMethodIdPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentMethodId** | **string** |  | ☑️ |  | 

### Request Parameters - GetPaymentsByPaymentMethodIdRequestParams

Parameters that are passed through a pointer to a apiGetPaymentsByPaymentMethodIdRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **forUserId** |**string**|  |  | 
|  **paymentRequestId** |**string[]**|  |  | 
|  **paymentMethodId2** |**string[]**|  |  | 
|  **referenceId** |**string[]**|  |  | 
|  **paymentMethodType** |[**PaymentMethodType[]**](payment_method/PaymentMethodType.md)|  |  | 
|  **channelCode** |**string[]**|  |  | 
|  **status** |**string[]**|  |  | 
|  **currency** |**string[]**|  |  | 
|  **createdGte** |**time.Time**|  |  | 
|  **createdLte** |**time.Time**|  |  | 
|  **updatedGte** |**time.Time**|  |  | 
|  **updatedLte** |**time.Time**|  |  | 
|  **limit** |**int32**|  |  | 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
        "time"
    xendit "github.com/xendit/xendit-go/v4"
    payment_method "github.com/xendit/xendit-go/v4/payment_method"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    paymentRequestId := []string{"Inner_example"} // [OPTIONAL] | []string

    paymentMethodId2 := []string{"Inner_example"} // [OPTIONAL] | []string

    referenceId := []string{"Inner_example"} // [OPTIONAL] | []string

    paymentMethodType := []payment_method.PaymentMethodType{payment_method.PaymentMethodType("CARD")} // [OPTIONAL] | []PaymentMethodType

    channelCode := []string{"Inner_example"} // [OPTIONAL] | []string

    status := []string{"Inner_example"} // [OPTIONAL] | []string

    currency := []string{"Inner_example"} // [OPTIONAL] | []string

    createdGte := time.Now() // [OPTIONAL] | time.Time

    createdLte := time.Now() // [OPTIONAL] | time.Time

    updatedGte := time.Now() // [OPTIONAL] | time.Time

    updatedLte := time.Now() // [OPTIONAL] | time.Time

    limit := int32(56) // [OPTIONAL] | int32

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.GetPaymentsByPaymentMethodId(context.Background(), paymentMethodId).
        ForUserId(forUserId).
        PaymentRequestId(paymentRequestId).
        PaymentMethodId2(paymentMethodId2).
        ReferenceId(referenceId).
        PaymentMethodType(paymentMethodType).
        ChannelCode(channelCode).
        Status(status).
        Currency(currency).
        CreatedGte(createdGte).
        CreatedLte(createdLte).
        UpdatedGte(updatedGte).
        UpdatedLte(updatedLte).
        Limit(limit). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodApi.GetPaymentsByPaymentMethodId``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentsByPaymentMethodId`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodApi.GetPaymentsByPaymentMethodId`: %v\n", resp)
}
```

## `PatchPaymentMethod()` Function

Patch payment methods



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `PatchPaymentMethod` |
| Path Parameters  |  [PatchPaymentMethodPathParams](#request-parameters--PatchPaymentMethodPathParams)	 |
| Request Parameters  |  [PatchPaymentMethodRequestParams](#request-parameters--PatchPaymentMethodRequestParams)	 |
| Return Type  | [**PaymentMethod**](payment_method/PaymentMethod.md) |

### Path Parameters - PatchPaymentMethodPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentMethodId** | **string** |  | ☑️ |  | 

### Request Parameters - PatchPaymentMethodRequestParams

Parameters that are passed through a pointer to a apiPatchPaymentMethodRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **forUserId** |**string**|  |  | 
|  **paymentMethodUpdateParameters** |[**PaymentMethodUpdateParameters**](payment_method/PaymentMethodUpdateParameters.md)|  |  | 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v4"
    payment_method "github.com/xendit/xendit-go/v4/payment_method"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    paymentMethodUpdateParameters := *payment_method.NewPaymentMethodUpdateParameters() // [OPTIONAL] | PaymentMethodUpdateParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.PatchPaymentMethod(context.Background(), paymentMethodId).
        ForUserId(forUserId).
        PaymentMethodUpdateParameters(paymentMethodUpdateParameters). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodApi.PatchPaymentMethod``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodApi.PatchPaymentMethod`: %v\n", resp)
}
```

## `GetAllPaymentMethods()` Function

Get all payment methods by filters



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetAllPaymentMethods` |
| Path Parameters  |  [GetAllPaymentMethodsPathParams](#request-parameters--GetAllPaymentMethodsPathParams)	 |
| Request Parameters  |  [GetAllPaymentMethodsRequestParams](#request-parameters--GetAllPaymentMethodsRequestParams)	 |
| Return Type  | [**PaymentMethodList**](payment_method/PaymentMethodList.md) |

### Path Parameters - GetAllPaymentMethodsPathParams
This endpoint does not need any path parameter.


### Request Parameters - GetAllPaymentMethodsRequestParams

Parameters that are passed through a pointer to a apiGetAllPaymentMethodsRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **forUserId** |**string**|  |  | 
|  **id** |**string[]**|  |  | 
|  **type_** |**string[]**|  |  | 
|  **status** |[**PaymentMethodStatus[]**](payment_method/PaymentMethodStatus.md)|  |  | 
|  **reusability** |[**PaymentMethodReusability**](payment_method/PaymentMethodReusabilitypayment_method/.md)|  |  | 
|  **customerId** |**string**|  |  | 
|  **referenceId** |**string**|  |  | 
|  **afterId** |**string**|  |  | 
|  **beforeId** |**string**|  |  | 
|  **limit** |**int32**|  |  | 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v4"
    payment_method "github.com/xendit/xendit-go/v4/payment_method"
)

func main() {
    
    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    id := []string{"Inner_example"} // [OPTIONAL] | []string

    type_ := []string{"Inner_example"} // [OPTIONAL] | []string

    status := []payment_method.PaymentMethodStatus{payment_method.PaymentMethodStatus("ACTIVE")} // [OPTIONAL] | []PaymentMethodStatus

    reusability := payment_method.PaymentMethodReusability("MULTIPLE_USE") // [OPTIONAL] | PaymentMethodReusability

    customerId := "customerId_example" // [OPTIONAL] | string

    referenceId := "referenceId_example" // [OPTIONAL] | string

    afterId := "afterId_example" // [OPTIONAL] | string

    beforeId := "beforeId_example" // [OPTIONAL] | string

    limit := int32(56) // [OPTIONAL] | int32

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.GetAllPaymentMethods(context.Background()).
        ForUserId(forUserId).
        Id(id).
        Type_(type_).
        Status(status).
        Reusability(reusability).
        CustomerId(customerId).
        ReferenceId(referenceId).
        AfterId(afterId).
        BeforeId(beforeId).
        Limit(limit). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodApi.GetAllPaymentMethods``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPaymentMethods`: PaymentMethodList
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodApi.GetAllPaymentMethods`: %v\n", resp)
}
```

## `ExpirePaymentMethod()` Function

Expires a payment method



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `ExpirePaymentMethod` |
| Path Parameters  |  [ExpirePaymentMethodPathParams](#request-parameters--ExpirePaymentMethodPathParams)	 |
| Request Parameters  |  [ExpirePaymentMethodRequestParams](#request-parameters--ExpirePaymentMethodRequestParams)	 |
| Return Type  | [**PaymentMethod**](payment_method/PaymentMethod.md) |

### Path Parameters - ExpirePaymentMethodPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentMethodId** | **string** |  | ☑️ |  | 

### Request Parameters - ExpirePaymentMethodRequestParams

Parameters that are passed through a pointer to a apiExpirePaymentMethodRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **forUserId** |**string**|  |  | 
|  **paymentMethodExpireParameters** |[**PaymentMethodExpireParameters**](payment_method/PaymentMethodExpireParameters.md)|  |  | 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v4"
    payment_method "github.com/xendit/xendit-go/v4/payment_method"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    paymentMethodExpireParameters := *payment_method.NewPaymentMethodExpireParameters() // [OPTIONAL] | PaymentMethodExpireParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.ExpirePaymentMethod(context.Background(), paymentMethodId).
        ForUserId(forUserId).
        PaymentMethodExpireParameters(paymentMethodExpireParameters). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodApi.ExpirePaymentMethod``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpirePaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodApi.ExpirePaymentMethod`: %v\n", resp)
}
```

## `AuthPaymentMethod()` Function

Validate a payment method's linking OTP



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `AuthPaymentMethod` |
| Path Parameters  |  [AuthPaymentMethodPathParams](#request-parameters--AuthPaymentMethodPathParams)	 |
| Request Parameters  |  [AuthPaymentMethodRequestParams](#request-parameters--AuthPaymentMethodRequestParams)	 |
| Return Type  | [**PaymentMethod**](payment_method/PaymentMethod.md) |

### Path Parameters - AuthPaymentMethodPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentMethodId** | **string** |  | ☑️ |  | 

### Request Parameters - AuthPaymentMethodRequestParams

Parameters that are passed through a pointer to a apiAuthPaymentMethodRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **forUserId** |**string**|  |  | 
|  **paymentMethodAuthParameters** |[**PaymentMethodAuthParameters**](payment_method/PaymentMethodAuthParameters.md)|  |  | 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v4"
    payment_method "github.com/xendit/xendit-go/v4/payment_method"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    paymentMethodAuthParameters := *payment_method.NewPaymentMethodAuthParameters("AuthCode_example") // [OPTIONAL] | PaymentMethodAuthParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.AuthPaymentMethod(context.Background(), paymentMethodId).
        ForUserId(forUserId).
        PaymentMethodAuthParameters(paymentMethodAuthParameters). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodApi.AuthPaymentMethod``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodApi.AuthPaymentMethod`: %v\n", resp)
}
```

## `SimulatePayment()` Function

Makes payment with matching PaymentMethodID.



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `SimulatePayment` |
| Path Parameters  |  [SimulatePaymentPathParams](#request-parameters--SimulatePaymentPathParams)	 |
| Request Parameters  |  [SimulatePaymentRequestParams](#request-parameters--SimulatePaymentRequestParams)	 |
| Return Type  |  (empty response body) |

### Path Parameters - SimulatePaymentPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentMethodId** | **string** |  | ☑️ |  | 

### Request Parameters - SimulatePaymentRequestParams

Parameters that are passed through a pointer to a apiSimulatePaymentRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **simulatePaymentRequest** |[**SimulatePaymentRequest**](payment_method/SimulatePaymentRequest.md)|  |  | 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v4"
    payment_method "github.com/xendit/xendit-go/v4/payment_method"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    simulatePaymentRequest := *payment_method.NewSimulatePaymentRequest() // [OPTIONAL] | SimulatePaymentRequest

    xenditClient := xendit.NewClient("API-KEY")

    r, err := xenditClient.PaymentMethodApi.SimulatePayment(context.Background(), paymentMethodId).
        SimulatePaymentRequest(simulatePaymentRequest). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodApi.SimulatePayment``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

## Callback Objects
Use the following callback objects provided by Xendit to receive callbacks (also known as webhooks) that Xendit sends you on events, such as successful payments. Note that the example is meant to illustrate the contents of the callback object -- you will not need to instantiate these objects in practice
### PaymentMethodCallback Object
>Callback for active or expired E-Wallet or Direct Debit account linking, Virtual Accounts or QR strings

Model Documentation: [PaymentMethodCallback](payment_method/PaymentMethodCallback.md)
#### Usage Example
Note that the example is meant to illustrate the contents of the callback object -- you will not need to instantiate these objects in practice
```go
PaymentMethodCallbackJson := map[string]interface{}{
  "event" : "payment_method.activated",
  "data" : {
    "id" : "pm-6ff0b6f2-f5de-457f-b08f-bc98fbae485a",
    "card" : null,
    "type" : "DIRECT_DEBIT",
    "status" : "ACTIVE",
    "country" : "PH",
    "created" : "2022-08-12T13:30:26.579048Z",
    "ewallet" : null,
    "qr_code" : null,
    "updated" : "2022-08-12T13:30:58.908220358Z",
    "metadata" : null,
    "customer_id" : "e2878b4c-d57e-4a2c-922d-c0313c2800a3",
    "description" : null,
    "reusability" : "MULTIPLE_USE",
    "direct_debit" : {
      "type" : "BANK_ACCOUNT",
      "debit_card" : null,
      "bank_account" : {
        "bank_account_hash" : "b4dfa99c9b60c77f2e3962b73c098945",
        "masked_bank_account_number" : "XXXXXX1234"
      },
      "channel_code" : "BPI",
      "channel_properties" : {
        "failure_return_url" : "https://your-redirect-website.com/failure",
        "success_return_url" : "https://your-redirect-website.com/success"
      }
    },
    "failure_code" : null,
    "reference_id" : "620b9df4-fe69-4bfd-b9d4-5cba6861db8a",
    "virtual_account" : null,
    "over_the_counter" : null,
    "billing_information" : null,
    "direct_bank_transfer" : null
  },
  "created" : "2022-08-12T13:30:59.074277334Z",
  "business_id" : "5f27a14a9bf05c73dd040bc8"
}
jsonData, _ := json.Marshal(PaymentMethodCallbackJson)
```

You may then use the callback object in your webhook or callback handler like so,
```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v4"
    payment_method "github.com/xendit/xendit-go/v4/payment_method"
)

func main() {
    // get callback object here
    // define jsonData

    // unmarshal callback object jsonData
    var PaymentMethodCallback payment_method.PaymentMethodCallback
    err := json.Unmarshal(jsonData, &PaymentMethodCallback)
    if err == nil {
        fmt.Fprintf(os.Stdout, "Callback Object ID: %v\n", PaymentMethodCallback.GetId())
        // do things here with the callback
    }
}
```

[[Back to README]](../README.md)

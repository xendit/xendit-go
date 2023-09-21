# xendit\PaymentMethodApi

All URIs are relative to *https://api.xendit.co*

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**AuthPaymentMethod**](PaymentMethodApi.md#AuthPaymentMethod) | **Post** /v2/payment_methods/{paymentMethodId}/auth | Validate a payment method&#39;s linking OTP |
| [**CreatePaymentMethod**](PaymentMethodApi.md#CreatePaymentMethod) | **Post** /v2/payment_methods | Creates payment method |
| [**ExpirePaymentMethod**](PaymentMethodApi.md#ExpirePaymentMethod) | **Post** /v2/payment_methods/{paymentMethodId}/expire | Expires a payment method |
| [**GetAllPaymentChannels**](PaymentMethodApi.md#GetAllPaymentChannels) | **Get** /v2/payment_methods/channels | Get all payment channels |
| [**GetAllPaymentMethods**](PaymentMethodApi.md#GetAllPaymentMethods) | **Get** /v2/payment_methods | Get all payment methods by filters |
| [**GetPaymentMethodByID**](PaymentMethodApi.md#GetPaymentMethodByID) | **Get** /v2/payment_methods/{paymentMethodId} | Get payment method by ID |
| [**GetPaymentsByPaymentMethodId**](PaymentMethodApi.md#GetPaymentsByPaymentMethodId) | **Get** /v2/payment_methods/{paymentMethodId}/payments | Returns payments with matching PaymentMethodID. |
| [**PatchPaymentMethod**](PaymentMethodApi.md#PatchPaymentMethod) | **Patch** /v2/payment_methods/{paymentMethodId} | Patch payment methods |
| [**SimulatePayment**](PaymentMethodApi.md#SimulatePayment) | **Post** /v2/payment_methods/{paymentMethodId}/payments/simulate | Makes payment with matching PaymentMethodID. |



## AuthPaymentMethod

Validate a payment method's linking OTP



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_method "github.com/xendit/xendit-go/v3/payment_method"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    paymentMethodAuthParameters := *payment_method.NewPaymentMethodAuthParameters("AuthCode_example") // [OPTIONAL] | PaymentMethodAuthParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.AuthPaymentMethod(context.Background(), paymentMethodId).
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

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **paymentMethodId** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthPaymentMethodRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **paymentMethodAuthParameters** |[**PaymentMethodAuthParameters**](payment_method/PaymentMethodAuthParameters.md)|  |  | 

### Return type

[**PaymentMethod**](payment_method/PaymentMethod.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## CreatePaymentMethod

Creates payment method



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_method "github.com/xendit/xendit-go/v3/payment_method"
)

func main() {
    
    paymentMethodParameters := *payment_method.NewPaymentMethodParameters(payment_method.PaymentMethodType("CARD"), payment_method.PaymentMethodReusability("MULTIPLE_USE")) // [OPTIONAL] | PaymentMethodParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.CreatePaymentMethod(context.Background()).
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

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentMethodRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
|  **paymentMethodParameters** |[**PaymentMethodParameters**](payment_method/PaymentMethodParameters.md)|  |  | 

### Return type

[**PaymentMethod**](payment_method/PaymentMethod.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## ExpirePaymentMethod

Expires a payment method



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_method "github.com/xendit/xendit-go/v3/payment_method"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    paymentMethodExpireParameters := *payment_method.NewPaymentMethodExpireParameters() // [OPTIONAL] | PaymentMethodExpireParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.ExpirePaymentMethod(context.Background(), paymentMethodId).
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

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **paymentMethodId** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpirePaymentMethodRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **paymentMethodExpireParameters** |[**PaymentMethodExpireParameters**](payment_method/PaymentMethodExpireParameters.md)|  |  | 

### Return type

[**PaymentMethod**](payment_method/PaymentMethod.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAllPaymentChannels

Get all payment channels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_method "github.com/xendit/xendit-go/v3/payment_method"
)

func main() {
    
    // (default to true)
    isActivated := true // [OPTIONAL] | bool

    type_ := "DIRECT_DEBIT" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.GetAllPaymentChannels(context.Background()).
        IsActivated(isActivated).
        Type_(type_). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodApi.GetAllPaymentChannels``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPaymentChannels`: PaymentChannelList
    fmt.Fprintf(os.Stdout, "Response from `PaymentMethodApi.GetAllPaymentChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPaymentChannelsRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
|  **isActivated** |**bool**|  | [default to true] | 
|  **type_** |**string**|  |  | 

### Return type

[**PaymentChannelList**](payment_method/PaymentChannelList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAllPaymentMethods

Get all payment methods by filters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_method "github.com/xendit/xendit-go/v3/payment_method"
)

func main() {
    
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

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPaymentMethodsRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
|  **id** |**string[]**|  |  | 
|  **type_** |**string[]**|  |  | 
|  **status** |[**PaymentMethodStatus[]**](payment_method/PaymentMethodStatus.md)|  |  | 
|  **reusability** |[**PaymentMethodReusability**](payment_method/PaymentMethodReusabilitypayment_method/.md)|  |  | 
|  **customerId** |**string**|  |  | 
|  **referenceId** |**string**|  |  | 
|  **afterId** |**string**|  |  | 
|  **beforeId** |**string**|  |  | 
|  **limit** |**int32**|  |  | 

### Return type

[**PaymentMethodList**](payment_method/PaymentMethodList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPaymentMethodByID

Get payment method by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_method "github.com/xendit/xendit-go/v3/payment_method"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.GetPaymentMethodByID(context.Background(), paymentMethodId). // [OPTIONAL]
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

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **paymentMethodId** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentMethodByIDRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 

### Return type

[**PaymentMethod**](payment_method/PaymentMethod.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPaymentsByPaymentMethodId

Returns payments with matching PaymentMethodID.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    xendit "github.com/xendit/xendit-go/v3"
    payment_method "github.com/xendit/xendit-go/v3/payment_method"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

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

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **paymentMethodId** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsByPaymentMethodIdRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
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

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## PatchPaymentMethod

Patch payment methods



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_method "github.com/xendit/xendit-go/v3/payment_method"
)

func main() {
    
    paymentMethodId := "pm-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    paymentMethodUpdateParameters := *payment_method.NewPaymentMethodUpdateParameters() // [OPTIONAL] | PaymentMethodUpdateParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentMethodApi.PatchPaymentMethod(context.Background(), paymentMethodId).
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

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **paymentMethodId** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPaymentMethodRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **paymentMethodUpdateParameters** |[**PaymentMethodUpdateParameters**](payment_method/PaymentMethodUpdateParameters.md)|  |  | 

### Return type

[**PaymentMethod**](payment_method/PaymentMethod.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## SimulatePayment

Makes payment with matching PaymentMethodID.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_method "github.com/xendit/xendit-go/v3/payment_method"
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

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **paymentMethodId** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSimulatePaymentRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **simulatePaymentRequest** |[**SimulatePaymentRequest**](payment_method/SimulatePaymentRequest.md)|  |  | 

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


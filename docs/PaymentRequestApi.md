# xendit\PaymentRequestApi

All URIs are relative to *https://api.xendit.co*

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**AuthorizePaymentRequest**](PaymentRequestApi.md#AuthorizePaymentRequest) | **Post** /payment_requests/{paymentRequestId}/auth | Payment Request Authorize |
| [**CapturePaymentRequest**](PaymentRequestApi.md#CapturePaymentRequest) | **Post** /payment_requests/{paymentRequestId}/captures | Payment Request Capture |
| [**CreatePaymentRequest**](PaymentRequestApi.md#CreatePaymentRequest) | **Post** /payment_requests | Create Payment Request |
| [**GetAllPaymentRequests**](PaymentRequestApi.md#GetAllPaymentRequests) | **Get** /payment_requests | Get all payment requests by filter |
| [**GetPaymentRequestByID**](PaymentRequestApi.md#GetPaymentRequestByID) | **Get** /payment_requests/{paymentRequestId} | Get payment request by ID |
| [**GetPaymentRequestCaptures**](PaymentRequestApi.md#GetPaymentRequestCaptures) | **Get** /payment_requests/{paymentRequestId}/captures | Get Payment Request Capture |
| [**ResendPaymentRequestAuth**](PaymentRequestApi.md#ResendPaymentRequestAuth) | **Post** /payment_requests/{paymentRequestId}/auth/resend | Payment Request Resend Auth |



## AuthorizePaymentRequest

Payment Request Authorize



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_request "github.com/xendit/xendit-go/v3/payment_request"
)

func main() {
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    paymentRequestAuthParameters := *payment_request.NewPaymentRequestAuthParameters("AuthCode_example") // [OPTIONAL] | PaymentRequestAuthParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.AuthorizePaymentRequest(context.Background(), paymentRequestId).
        PaymentRequestAuthParameters(paymentRequestAuthParameters). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestApi.AuthorizePaymentRequest``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthorizePaymentRequest`: PaymentRequest
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestApi.AuthorizePaymentRequest`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **paymentRequestId** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizePaymentRequestRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **paymentRequestAuthParameters** |[**PaymentRequestAuthParameters**](payment_request/PaymentRequestAuthParameters.md)|  |  | 

### Return type

[**PaymentRequest**](payment_request/PaymentRequest.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## CapturePaymentRequest

Payment Request Capture



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_request "github.com/xendit/xendit-go/v3/payment_request"
)

func main() {
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    captureParameters := *payment_request.NewCaptureParameters(float64(123)) // [OPTIONAL] | CaptureParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.CapturePaymentRequest(context.Background(), paymentRequestId).
        CaptureParameters(captureParameters). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestApi.CapturePaymentRequest``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CapturePaymentRequest`: Capture
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestApi.CapturePaymentRequest`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **paymentRequestId** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCapturePaymentRequestRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **captureParameters** |[**CaptureParameters**](payment_request/CaptureParameters.md)|  |  | 

### Return type

[**Capture**](payment_request/Capture.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## CreatePaymentRequest

Create Payment Request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_request "github.com/xendit/xendit-go/v3/payment_request"
)

func main() {
    
    idempotencyKey := "5f9a3fbd571a1c4068aa40ce" // [OPTIONAL] | string

    paymentRequestParameters := *payment_request.NewPaymentRequestParameters(payment_request.PaymentRequestCurrency("IDR")) // [OPTIONAL] | PaymentRequestParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.CreatePaymentRequest(context.Background()).
        IdempotencyKey(idempotencyKey).
        PaymentRequestParameters(paymentRequestParameters). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestApi.CreatePaymentRequest``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentRequest`: PaymentRequest
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestApi.CreatePaymentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRequestRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
|  **idempotencyKey** |**string**|  |  | 
|  **paymentRequestParameters** |[**PaymentRequestParameters**](payment_request/PaymentRequestParameters.md)|  |  | 

### Return type

[**PaymentRequest**](payment_request/PaymentRequest.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAllPaymentRequests

Get all payment requests by filter



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_request "github.com/xendit/xendit-go/v3/payment_request"
)

func main() {
    
    referenceId := []string{"Inner_example"} // [OPTIONAL] | []string

    id := []string{"Inner_example"} // [OPTIONAL] | []string

    customerId := []string{"Inner_example"} // [OPTIONAL] | []string

    limit := int32(56) // [OPTIONAL] | int32

    beforeId := "beforeId_example" // [OPTIONAL] | string

    afterId := "afterId_example" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.GetAllPaymentRequests(context.Background()).
        ReferenceId(referenceId).
        Id(id).
        CustomerId(customerId).
        Limit(limit).
        BeforeId(beforeId).
        AfterId(afterId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestApi.GetAllPaymentRequests``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPaymentRequests`: PaymentRequestListResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestApi.GetAllPaymentRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPaymentRequestsRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
|  **referenceId** |**string[]**|  |  | 
|  **id** |**string[]**|  |  | 
|  **customerId** |**string[]**|  |  | 
|  **limit** |**int32**|  |  | 
|  **beforeId** |**string**|  |  | 
|  **afterId** |**string**|  |  | 

### Return type

[**PaymentRequestListResponse**](payment_request/PaymentRequestListResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPaymentRequestByID

Get payment request by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_request "github.com/xendit/xendit-go/v3/payment_request"
)

func main() {
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.GetPaymentRequestByID(context.Background(), paymentRequestId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestApi.GetPaymentRequestByID``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentRequestByID`: PaymentRequest
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestApi.GetPaymentRequestByID`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **paymentRequestId** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRequestByIDRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 

### Return type

[**PaymentRequest**](payment_request/PaymentRequest.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPaymentRequestCaptures

Get Payment Request Capture



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_request "github.com/xendit/xendit-go/v3/payment_request"
)

func main() {
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    limit := int32(56) // [OPTIONAL] | int32

    afterId := "afterId_example" // [OPTIONAL] | string

    beforeId := "beforeId_example" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.GetPaymentRequestCaptures(context.Background(), paymentRequestId).
        Limit(limit).
        AfterId(afterId).
        BeforeId(beforeId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestApi.GetPaymentRequestCaptures``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentRequestCaptures`: CaptureListResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestApi.GetPaymentRequestCaptures`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **paymentRequestId** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentRequestCapturesRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **limit** |**int32**|  |  | 
|  **afterId** |**string**|  |  | 
|  **beforeId** |**string**|  |  | 

### Return type

[**CaptureListResponse**](payment_request/CaptureListResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## ResendPaymentRequestAuth

Payment Request Resend Auth



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payment_request "github.com/xendit/xendit-go/v3/payment_request"
)

func main() {
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.ResendPaymentRequestAuth(context.Background(), paymentRequestId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestApi.ResendPaymentRequestAuth``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendPaymentRequestAuth`: PaymentRequest
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestApi.ResendPaymentRequestAuth`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **paymentRequestId** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendPaymentRequestAuthRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 

### Return type

[**PaymentRequest**](payment_request/PaymentRequest.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


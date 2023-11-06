# PaymentRequestApi


You can use the APIs below to interface with Xendit's `PaymentRequestApi`.
To start using the API, you need to configure the secret key and initiate the client instance.

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
)

func main() {
    xenditClient := xendit.NewClient("API-KEY")
}
```

All URIs are relative to *https://api.xendit.co*

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**CreatePaymentRequest**](PaymentRequestApi.md#createpaymentrequest-function) | **Post** /payment_requests | Create Payment Request |
| [**GetPaymentRequestByID**](PaymentRequestApi.md#getpaymentrequestbyid-function) | **Get** /payment_requests/{paymentRequestId} | Get payment request by ID |
| [**GetPaymentRequestCaptures**](PaymentRequestApi.md#getpaymentrequestcaptures-function) | **Get** /payment_requests/{paymentRequestId}/captures | Get Payment Request Capture |
| [**GetAllPaymentRequests**](PaymentRequestApi.md#getallpaymentrequests-function) | **Get** /payment_requests | Get all payment requests by filter |
| [**CapturePaymentRequest**](PaymentRequestApi.md#capturepaymentrequest-function) | **Post** /payment_requests/{paymentRequestId}/captures | Payment Request Capture |
| [**AuthorizePaymentRequest**](PaymentRequestApi.md#authorizepaymentrequest-function) | **Post** /payment_requests/{paymentRequestId}/auth | Payment Request Authorize |
| [**ResendPaymentRequestAuth**](PaymentRequestApi.md#resendpaymentrequestauth-function) | **Post** /payment_requests/{paymentRequestId}/auth/resend | Payment Request Resend Auth |



## `CreatePaymentRequest()` Function

Create Payment Request



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `CreatePaymentRequest` |
| Path Parameters  |  [CreatePaymentRequestPathParams](#request-parameters--CreatePaymentRequestPathParams)	 |
| Request Parameters  |  [CreatePaymentRequestRequestParams](#request-parameters--CreatePaymentRequestRequestParams)	 |
| Return Type  | [**PaymentRequest**](payment_request/PaymentRequest.md) |

### Path Parameters - CreatePaymentRequestPathParams
This endpoint does not need any path parameter.


### Request Parameters - CreatePaymentRequestRequestParams

Parameters that are passed through a pointer to a apiCreatePaymentRequestRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **idempotencyKey** |**string**|  |  | 
|  **forUserId** |**string**|  |  | 
|  **paymentRequestParameters** |[**PaymentRequestParameters**](payment_request/PaymentRequestParameters.md)|  |  | 

### Usage Example

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

    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    paymentRequestParameters := *payment_request.NewPaymentRequestParameters(payment_request.PaymentRequestCurrency("IDR")) // [OPTIONAL] | PaymentRequestParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.CreatePaymentRequest(context.Background()).
        IdempotencyKey(idempotencyKey).
        ForUserId(forUserId).
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

## `GetPaymentRequestByID()` Function

Get payment request by ID



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetPaymentRequestByID` |
| Path Parameters  |  [GetPaymentRequestByIDPathParams](#request-parameters--GetPaymentRequestByIDPathParams)	 |
| Request Parameters  |  [GetPaymentRequestByIDRequestParams](#request-parameters--GetPaymentRequestByIDRequestParams)	 |
| Return Type  | [**PaymentRequest**](payment_request/PaymentRequest.md) |

### Path Parameters - GetPaymentRequestByIDPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentRequestId** | **string** |  | ☑️ |  | 

### Request Parameters - GetPaymentRequestByIDRequestParams

Parameters that are passed through a pointer to a apiGetPaymentRequestByIDRequest struct via the builder pattern

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
    xendit "github.com/xendit/xendit-go/v3"
    payment_request "github.com/xendit/xendit-go/v3/payment_request"
)

func main() {
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.GetPaymentRequestByID(context.Background(), paymentRequestId).
        ForUserId(forUserId). // [OPTIONAL]
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

## `GetPaymentRequestCaptures()` Function

Get Payment Request Capture



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetPaymentRequestCaptures` |
| Path Parameters  |  [GetPaymentRequestCapturesPathParams](#request-parameters--GetPaymentRequestCapturesPathParams)	 |
| Request Parameters  |  [GetPaymentRequestCapturesRequestParams](#request-parameters--GetPaymentRequestCapturesRequestParams)	 |
| Return Type  | [**CaptureListResponse**](payment_request/CaptureListResponse.md) |

### Path Parameters - GetPaymentRequestCapturesPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentRequestId** | **string** |  | ☑️ |  | 

### Request Parameters - GetPaymentRequestCapturesRequestParams

Parameters that are passed through a pointer to a apiGetPaymentRequestCapturesRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **forUserId** |**string**|  |  | 
|  **limit** |**int32**|  |  | 

### Usage Example

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

    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    limit := int32(56) // [OPTIONAL] | int32

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.GetPaymentRequestCaptures(context.Background(), paymentRequestId).
        ForUserId(forUserId).
        Limit(limit). // [OPTIONAL]
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

## `GetAllPaymentRequests()` Function

Get all payment requests by filter



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetAllPaymentRequests` |
| Path Parameters  |  [GetAllPaymentRequestsPathParams](#request-parameters--GetAllPaymentRequestsPathParams)	 |
| Request Parameters  |  [GetAllPaymentRequestsRequestParams](#request-parameters--GetAllPaymentRequestsRequestParams)	 |
| Return Type  | [**PaymentRequestListResponse**](payment_request/PaymentRequestListResponse.md) |

### Path Parameters - GetAllPaymentRequestsPathParams
This endpoint does not need any path parameter.


### Request Parameters - GetAllPaymentRequestsRequestParams

Parameters that are passed through a pointer to a apiGetAllPaymentRequestsRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **forUserId** |**string**|  |  | 
|  **referenceId** |**string[]**|  |  | 
|  **id** |**string[]**|  |  | 
|  **customerId** |**string[]**|  |  | 
|  **limit** |**int32**|  |  | 
|  **beforeId** |**string**|  |  | 
|  **afterId** |**string**|  |  | 

### Usage Example

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
    
    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    referenceId := []string{"Inner_example"} // [OPTIONAL] | []string

    id := []string{"Inner_example"} // [OPTIONAL] | []string

    customerId := []string{"Inner_example"} // [OPTIONAL] | []string

    limit := int32(56) // [OPTIONAL] | int32

    beforeId := "beforeId_example" // [OPTIONAL] | string

    afterId := "afterId_example" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.GetAllPaymentRequests(context.Background()).
        ForUserId(forUserId).
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

## `CapturePaymentRequest()` Function

Payment Request Capture



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `CapturePaymentRequest` |
| Path Parameters  |  [CapturePaymentRequestPathParams](#request-parameters--CapturePaymentRequestPathParams)	 |
| Request Parameters  |  [CapturePaymentRequestRequestParams](#request-parameters--CapturePaymentRequestRequestParams)	 |
| Return Type  | [**Capture**](payment_request/Capture.md) |

### Path Parameters - CapturePaymentRequestPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentRequestId** | **string** |  | ☑️ |  | 

### Request Parameters - CapturePaymentRequestRequestParams

Parameters that are passed through a pointer to a apiCapturePaymentRequestRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **forUserId** |**string**|  |  | 
|  **captureParameters** |[**CaptureParameters**](payment_request/CaptureParameters.md)|  |  | 

### Usage Example

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

    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    captureParameters := *payment_request.NewCaptureParameters(float64(123)) // [OPTIONAL] | CaptureParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.CapturePaymentRequest(context.Background(), paymentRequestId).
        ForUserId(forUserId).
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

## `AuthorizePaymentRequest()` Function

Payment Request Authorize



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `AuthorizePaymentRequest` |
| Path Parameters  |  [AuthorizePaymentRequestPathParams](#request-parameters--AuthorizePaymentRequestPathParams)	 |
| Request Parameters  |  [AuthorizePaymentRequestRequestParams](#request-parameters--AuthorizePaymentRequestRequestParams)	 |
| Return Type  | [**PaymentRequest**](payment_request/PaymentRequest.md) |

### Path Parameters - AuthorizePaymentRequestPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentRequestId** | **string** |  | ☑️ |  | 

### Request Parameters - AuthorizePaymentRequestRequestParams

Parameters that are passed through a pointer to a apiAuthorizePaymentRequestRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **forUserId** |**string**|  |  | 
|  **paymentRequestAuthParameters** |[**PaymentRequestAuthParameters**](payment_request/PaymentRequestAuthParameters.md)|  |  | 

### Usage Example

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

    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    paymentRequestAuthParameters := *payment_request.NewPaymentRequestAuthParameters("AuthCode_example") // [OPTIONAL] | PaymentRequestAuthParameters

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.AuthorizePaymentRequest(context.Background(), paymentRequestId).
        ForUserId(forUserId).
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

## `ResendPaymentRequestAuth()` Function

Payment Request Resend Auth



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `ResendPaymentRequestAuth` |
| Path Parameters  |  [ResendPaymentRequestAuthPathParams](#request-parameters--ResendPaymentRequestAuthPathParams)	 |
| Request Parameters  |  [ResendPaymentRequestAuthRequestParams](#request-parameters--ResendPaymentRequestAuthRequestParams)	 |
| Return Type  | [**PaymentRequest**](payment_request/PaymentRequest.md) |

### Path Parameters - ResendPaymentRequestAuthPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentRequestId** | **string** |  | ☑️ |  | 

### Request Parameters - ResendPaymentRequestAuthRequestParams

Parameters that are passed through a pointer to a apiResendPaymentRequestAuthRequest struct via the builder pattern

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
    xendit "github.com/xendit/xendit-go/v3"
    payment_request "github.com/xendit/xendit-go/v3/payment_request"
)

func main() {
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    forUserId := "5f9a3fbd571a1c4068aa40cf" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.ResendPaymentRequestAuth(context.Background(), paymentRequestId).
        ForUserId(forUserId). // [OPTIONAL]
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

[[Back to README]](../README.md)

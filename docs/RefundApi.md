# xendit\RefundApi

All URIs are relative to *https://api.xendit.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRefund**](RefundApi.md#CancelRefund) | **Post** /refunds/{refundID}/cancel | 
[**CreateRefund**](RefundApi.md#CreateRefund) | **Post** /refunds | 
[**GetAllRefunds**](RefundApi.md#GetAllRefunds) | **Get** /refunds/ | 
[**GetRefund**](RefundApi.md#GetRefund) | **Get** /refunds/{refundID} | 



## CancelRefund



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    refund "github.com/xendit/xendit-go/v3/refund"
)

func main() {
    
    refundID := "rfd-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    idempotencyKey := "9797b5a6-54ad-4511-80a4-ec451346808b" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.RefundApi.CancelRefund(context.Background(), refundID).
        IdempotencyKey(idempotencyKey). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundApi.CancelRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelRefund`: Refund
    fmt.Fprintf(os.Stdout, "Response from `RefundApi.CancelRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refundID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

### Return type

[**Refund**](refund/Refund.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## CreateRefund



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    refund "github.com/xendit/xendit-go/v3/refund"
)

func main() {
    
    idempotencyKey := "9797b5a6-54ad-4511-80a4-ec451346808b" // [OPTIONAL] | string

    createRefund := *refund.NewCreateRefund() // [OPTIONAL] | CreateRefund

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.RefundApi.CreateRefund(context.Background()).
        IdempotencyKey(idempotencyKey).
        CreateRefund(createRefund). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundApi.CreateRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRefund`: Refund
    fmt.Fprintf(os.Stdout, "Response from `RefundApi.CreateRefund`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** |  | 
 **createRefund** | [**CreateRefund**](CreateRefund.md) |  | 

### Return type

[**Refund**](refund/Refund.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAllRefunds



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    refund "github.com/xendit/xendit-go/v3/refund"
)

func main() {
    
    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.RefundApi.GetAllRefunds(context.Background()). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundApi.GetAllRefunds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRefunds`: RefundList
    fmt.Fprintf(os.Stdout, "Response from `RefundApi.GetAllRefunds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRefundsRequest struct via the builder pattern


### Return type

[**RefundList**](refund/RefundList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetRefund



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    refund "github.com/xendit/xendit-go/v3/refund"
)

func main() {
    
    refundID := "rfd-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    idempotencyKey := "9797b5a6-54ad-4511-80a4-ec451346808b" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.RefundApi.GetRefund(context.Background(), refundID).
        IdempotencyKey(idempotencyKey). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundApi.GetRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRefund`: Refund
    fmt.Fprintf(os.Stdout, "Response from `RefundApi.GetRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refundID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** |  | 

### Return type

[**Refund**](refund/Refund.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


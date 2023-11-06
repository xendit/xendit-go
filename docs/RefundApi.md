# RefundApi


You can use the APIs below to interface with Xendit's `RefundApi`.
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
| [**CreateRefund**](RefundApi.md#createrefund-function) | **Post** /refunds |  |
| [**GetRefund**](RefundApi.md#getrefund-function) | **Get** /refunds/{refundID} |  |
| [**GetAllRefunds**](RefundApi.md#getallrefunds-function) | **Get** /refunds |  |
| [**CancelRefund**](RefundApi.md#cancelrefund-function) | **Post** /refunds/{refundID}/cancel |  |



## `CreateRefund()` Function



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `CreateRefund` |
| Path Parameters  |  [CreateRefundPathParams](#request-parameters--CreateRefundPathParams)	 |
| Request Parameters  |  [CreateRefundRequestParams](#request-parameters--CreateRefundRequestParams)	 |
| Return Type  | [**Refund**](refund/Refund.md) |

### Path Parameters - CreateRefundPathParams
This endpoint does not need any path parameter.


### Request Parameters - CreateRefundRequestParams

Parameters that are passed through a pointer to a apiCreateRefundRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **idempotencyKey** |**string**|  |  | 
|  **forUserId** |**string**|  |  | 
|  **createRefund** |[**CreateRefund**](refund/CreateRefund.md)|  |  | 

### Usage Example

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

    forUserId := "5f9a3fbd571a1c4068aa40ce" // [OPTIONAL] | string

    createRefund := *refund.NewCreateRefund() // [OPTIONAL] | CreateRefund

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.RefundApi.CreateRefund(context.Background()).
        IdempotencyKey(idempotencyKey).
        ForUserId(forUserId).
        CreateRefund(createRefund). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundApi.CreateRefund``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRefund`: Refund
    fmt.Fprintf(os.Stdout, "Response from `RefundApi.CreateRefund`: %v\n", resp)
}
```

## `GetRefund()` Function



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetRefund` |
| Path Parameters  |  [GetRefundPathParams](#request-parameters--GetRefundPathParams)	 |
| Request Parameters  |  [GetRefundRequestParams](#request-parameters--GetRefundRequestParams)	 |
| Return Type  | [**Refund**](refund/Refund.md) |

### Path Parameters - GetRefundPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **refundID** | **string** |  | ☑️ |  | 

### Request Parameters - GetRefundRequestParams

Parameters that are passed through a pointer to a apiGetRefundRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **idempotencyKey** |**string**|  |  | 
|  **forUserId** |**string**|  |  | 

### Usage Example

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

    forUserId := "5f9a3fbd571a1c4068aa40ce" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.RefundApi.GetRefund(context.Background(), refundID).
        IdempotencyKey(idempotencyKey).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundApi.GetRefund``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRefund`: Refund
    fmt.Fprintf(os.Stdout, "Response from `RefundApi.GetRefund`: %v\n", resp)
}
```

## `GetAllRefunds()` Function



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetAllRefunds` |
| Path Parameters  |  [GetAllRefundsPathParams](#request-parameters--GetAllRefundsPathParams)	 |
| Request Parameters  |  [GetAllRefundsRequestParams](#request-parameters--GetAllRefundsRequestParams)	 |
| Return Type  | [**RefundList**](refund/RefundList.md) |

### Path Parameters - GetAllRefundsPathParams
This endpoint does not need any path parameter.


### Request Parameters - GetAllRefundsRequestParams

Parameters that are passed through a pointer to a apiGetAllRefundsRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **forUserId** |**string**|  |  | 
|  **paymentRequestId** |**string**|  |  | 
|  **invoiceId** |**string**|  |  | 
|  **paymentMethodType** |**string**|  |  | 
|  **channelCode** |**string**|  |  | 
|  **limit** |**float32**|  |  | 
|  **afterId** |**string**|  |  | 
|  **beforeId** |**string**|  |  | 

### Usage Example

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
    
    forUserId := "5f9a3fbd571a1c4068aa40ce" // [OPTIONAL] | string

    paymentRequestId := "paymentRequestId_example" // [OPTIONAL] | string

    invoiceId := "invoiceId_example" // [OPTIONAL] | string

    paymentMethodType := "paymentMethodType_example" // [OPTIONAL] | string

    channelCode := "channelCode_example" // [OPTIONAL] | string

    limit := float32(8.14) // [OPTIONAL] | float32

    afterId := "afterId_example" // [OPTIONAL] | string

    beforeId := "beforeId_example" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.RefundApi.GetAllRefunds(context.Background()).
        ForUserId(forUserId).
        PaymentRequestId(paymentRequestId).
        InvoiceId(invoiceId).
        PaymentMethodType(paymentMethodType).
        ChannelCode(channelCode).
        Limit(limit).
        AfterId(afterId).
        BeforeId(beforeId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundApi.GetAllRefunds``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRefunds`: RefundList
    fmt.Fprintf(os.Stdout, "Response from `RefundApi.GetAllRefunds`: %v\n", resp)
}
```

## `CancelRefund()` Function



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `CancelRefund` |
| Path Parameters  |  [CancelRefundPathParams](#request-parameters--CancelRefundPathParams)	 |
| Request Parameters  |  [CancelRefundRequestParams](#request-parameters--CancelRefundRequestParams)	 |
| Return Type  | [**Refund**](refund/Refund.md) |

### Path Parameters - CancelRefundPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **refundID** | **string** |  | ☑️ |  | 

### Request Parameters - CancelRefundRequestParams

Parameters that are passed through a pointer to a apiCancelRefundRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **idempotencyKey** |**string**|  |  | 
|  **forUserId** |**string**|  |  | 

### Usage Example

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

    forUserId := "5f9a3fbd571a1c4068aa40ce" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.RefundApi.CancelRefund(context.Background(), refundID).
        IdempotencyKey(idempotencyKey).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RefundApi.CancelRefund``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelRefund`: Refund
    fmt.Fprintf(os.Stdout, "Response from `RefundApi.CancelRefund`: %v\n", resp)
}
```

[[Back to README]](../README.md)

# RefundApi


You can use the APIs below to interface with Xendit's `RefundApi`.
To start using the API, you need to configure the secret key and initiate the client instance.

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v7"
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
    xendit "github.com/xendit/xendit-go/v7"
    refund "github.com/xendit/xendit-go/v7/refund"
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
    xendit "github.com/xendit/xendit-go/v7"
    refund "github.com/xendit/xendit-go/v7/refund"
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
    xendit "github.com/xendit/xendit-go/v7"
    refund "github.com/xendit/xendit-go/v7/refund"
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
    xendit "github.com/xendit/xendit-go/v7"
    refund "github.com/xendit/xendit-go/v7/refund"
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

## Callback Objects
Use the following callback objects provided by Xendit to receive callbacks (also known as webhooks) that Xendit sends you on events, such as successful payments. Note that the example is meant to illustrate the contents of the callback object -- you will not need to instantiate these objects in practice
### RefundCallback Object
>Callback for successful or failed Refunds made via the Payments API

Model Documentation: [RefundCallback](refund/RefundCallback.md)
#### Usage Example
Note that the example is meant to illustrate the contents of the callback object -- you will not need to instantiate these objects in practice
```go
RefundCallbackJson := `{
  "event" : "refund.succeeded",
  "business_id" : "5f27a14a9bf05c73dd040bc8",
  "created" : "2020-08-29T09:12:33.001Z",
  "data" : {
    "id" : "rfd-6f4a377d-a201-437f-9119-f8b00cbbe857",
    "payment_id" : "ddpy-3cd658ae-25b9-4659-aa36-596ae41a809f",
    "invoice_id" : null,
    "amount" : 10000,
    "payment_method_type" : "DIRECT_DEBIT",
    "channel_code" : "BPI",
    "currency" : "PHP",
    "status" : "SUCCEEDED",
    "reason" : "CANCELLATION",
    "reference_id" : "b2756a1e-e6cd-4352-9a68-0483aa2b6a2",
    "failure_code" : null,
    "refund_fee_amount" : null,
    "created" : "2020-08-30T09:12:33.001Z",
    "updated" : "2020-08-30T09:12:33.001Z",
    "metadata" : null
  }
}`
```

You may then use the callback object in your webhook or callback handler like so,
```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v7"
    refund "github.com/xendit/xendit-go/v7/refund"
)

func main() {
    // get callback object here
    // define jsonData

    // unmarshal callback object jsonData
    var RefundCallback refund.RefundCallback
    err := json.Unmarshal([]byte(RefundCallbackJson), &RefundCallback)
    if err == nil {
        fmt.Fprintf(os.Stdout, "Callback Object ID: %v\n", RefundCallback.GetId())
        // do things here with the callback
    }
}
```

[[Back to README]](../README.md)

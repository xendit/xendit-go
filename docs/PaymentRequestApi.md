# PaymentRequestApi


You can use the APIs below to interface with Xendit's `PaymentRequestApi`.
To start using the API, you need to configure the secret key and initiate the client instance.

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v5"
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
| [**SimulatePaymentRequestPayment**](PaymentRequestApi.md#simulatepaymentrequestpayment-function) | **Post** /payment_requests/{paymentRequestId}/payments/simulate | Payment Request Simulate Payment |



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
    xendit "github.com/xendit/xendit-go/v5"
    payment_request "github.com/xendit/xendit-go/v5/payment_request"
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
    xendit "github.com/xendit/xendit-go/v5"
    payment_request "github.com/xendit/xendit-go/v5/payment_request"
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
    xendit "github.com/xendit/xendit-go/v5"
    payment_request "github.com/xendit/xendit-go/v5/payment_request"
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
    xendit "github.com/xendit/xendit-go/v5"
    payment_request "github.com/xendit/xendit-go/v5/payment_request"
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
    xendit "github.com/xendit/xendit-go/v5"
    payment_request "github.com/xendit/xendit-go/v5/payment_request"
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
    xendit "github.com/xendit/xendit-go/v5"
    payment_request "github.com/xendit/xendit-go/v5/payment_request"
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
    xendit "github.com/xendit/xendit-go/v5"
    payment_request "github.com/xendit/xendit-go/v5/payment_request"
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

## `SimulatePaymentRequestPayment()` Function

Payment Request Simulate Payment



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `SimulatePaymentRequestPayment` |
| Path Parameters  |  [SimulatePaymentRequestPaymentPathParams](#request-parameters--SimulatePaymentRequestPaymentPathParams)	 |
| Request Parameters  |  [SimulatePaymentRequestPaymentRequestParams](#request-parameters--SimulatePaymentRequestPaymentRequestParams)	 |
| Return Type  | [**PaymentSimulation**](payment_request/PaymentSimulation.md) |

### Path Parameters - SimulatePaymentRequestPaymentPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **paymentRequestId** | **string** |  | ☑️ |  | 

### Request Parameters - SimulatePaymentRequestPaymentRequestParams

Parameters that are passed through a pointer to a apiSimulatePaymentRequestPaymentRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v5"
    payment_request "github.com/xendit/xendit-go/v5/payment_request"
)

func main() {
    
    paymentRequestId := "pr-1fdaf346-dd2e-4b6c-b938-124c7167a822" // [REQUIRED] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PaymentRequestApi.SimulatePaymentRequestPayment(context.Background(), paymentRequestId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestApi.SimulatePaymentRequestPayment``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimulatePaymentRequestPayment`: PaymentSimulation
    fmt.Fprintf(os.Stdout, "Response from `PaymentRequestApi.SimulatePaymentRequestPayment`: %v\n", resp)
}
```

## Callback Objects
Use the following callback objects provided by Xendit to receive callbacks (also known as webhooks) that Xendit sends you on events, such as successful payments. Note that the example is meant to illustrate the contents of the callback object -- you will not need to instantiate these objects in practice
### PaymentCallback Object
>Callback for successful or failed payments made via the Payments API

Model Documentation: [PaymentCallback](payment_request/PaymentCallback.md)
#### Usage Example
Note that the example is meant to illustrate the contents of the callback object -- you will not need to instantiate these objects in practice
```go
PaymentCallbackJson := `{
  "event" : "payment.succeeded",
  "data" : {
    "id" : "ddpy-3cd658ae-25b9-4659-aa36-596ae41a809f",
    "amount" : 1000,
    "status" : "SUCCEEDED",
    "country" : "PH",
    "created" : "2022-08-12T13:30:40.9209Z",
    "updated" : "2022-08-12T13:30:58.729373Z",
    "currency" : "PHP",
    "metadata" : {
      "sku" : "ABCDEFGH"
    },
    "customer_id" : "c832697e-a62d-46fa-a383-24930b155e81",
    "reference_id" : "25cfd0f9-baee-44ca-9a12-6debe03f3c22",
    "payment_method" : {
      "id" : "pm-951b1ad9-1fbb-4724-a744-8956ab6ed17f",
      "card" : null,
      "type" : "DIRECT_DEBIT",
      "status" : "ACTIVE",
      "created" : "2022-08-12T13:30:26.579048Z",
      "ewallet" : null,
      "qr_code" : null,
      "updated" : "2022-08-12T13:30:40.221525Z",
      "metadata" : {
        "sku" : "ABCDEFGH"
      },
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
      "reference_id" : "620b9df4-fe69-4bfd-b9d4-5cba6861db8a",
      "virtual_account" : null,
      "over_the_counter" : null,
      "direct_bank_transfer" : null
    },
    "description" : null,
    "failure_code" : null,
    "payment_detail" : null,
    "channel_properties" : null,
    "payment_request_id" : "pr-5b26cae1-545b-49e9-855e-f85128f3e705"
  },
  "created" : "2022-08-12T13:30:58.986Z",
  "business_id" : "5f27a14a9bf05c73dd040bc8",
  "api_version" : null
}`
```

You may then use the callback object in your webhook or callback handler like so,
```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v5"
    payment_request "github.com/xendit/xendit-go/v5/payment_request"
)

func main() {
    // get callback object here
    // define jsonData

    // unmarshal callback object jsonData
    var PaymentCallback payment_request.PaymentCallback
    err := json.Unmarshal([]byte(PaymentCallbackJson), &PaymentCallback)
    if err == nil {
        fmt.Fprintf(os.Stdout, "Callback Object ID: %v\n", PaymentCallback.GetId())
        // do things here with the callback
    }
}
```

[[Back to README]](../README.md)

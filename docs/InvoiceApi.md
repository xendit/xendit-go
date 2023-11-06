# InvoiceApi


You can use the APIs below to interface with Xendit's `InvoiceApi`.
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
| [**CreateInvoice**](InvoiceApi.md#createinvoice-function) | **Post** /v2/invoices/ | Create an invoice |
| [**GetInvoiceById**](InvoiceApi.md#getinvoicebyid-function) | **Get** /v2/invoices/{invoice_id} | Get invoice by invoice id |
| [**GetInvoices**](InvoiceApi.md#getinvoices-function) | **Get** /v2/invoices | Get all Invoices |
| [**ExpireInvoice**](InvoiceApi.md#expireinvoice-function) | **Post** /invoices/{invoice_id}/expire! | Manually expire an invoice |



## `CreateInvoice()` Function

Create an invoice

| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `CreateInvoice` |
| Path Parameters  |  [CreateInvoicePathParams](#request-parameters--CreateInvoicePathParams)	 |
| Request Parameters  |  [CreateInvoiceRequestParams](#request-parameters--CreateInvoiceRequestParams)	 |
| Return Type  | [**Invoice**](invoice/Invoice.md) |

### Path Parameters - CreateInvoicePathParams
This endpoint does not need any path parameter.


### Request Parameters - CreateInvoiceRequestParams

Parameters that are passed through a pointer to a apiCreateInvoiceRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **createInvoiceRequest** |[**CreateInvoiceRequest**](invoice/CreateInvoiceRequest.md)| ☑️ |  | 
|  **forUserId** |**string**|  |  | 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    invoice "github.com/xendit/xendit-go/v3/invoice"
)

func main() {
    
    createInvoiceRequest := *invoice.NewCreateInvoiceRequest("ExternalId_example", float32(123)) // [REQUIRED] | CreateInvoiceRequest

    // Business ID of the sub-account merchant (XP feature)
    forUserId := "62efe4c33e45694d63f585f8" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.InvoiceApi.CreateInvoice(context.Background()).
        CreateInvoiceRequest(createInvoiceRequest).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.CreateInvoice``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.CreateInvoice`: %v\n", resp)
}
```

## `GetInvoiceById()` Function

Get invoice by invoice id

| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetInvoiceById` |
| Path Parameters  |  [GetInvoiceByIdPathParams](#request-parameters--GetInvoiceByIdPathParams)	 |
| Request Parameters  |  [GetInvoiceByIdRequestParams](#request-parameters--GetInvoiceByIdRequestParams)	 |
| Return Type  | [**Invoice**](invoice/Invoice.md) |

### Path Parameters - GetInvoiceByIdPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **invoiceId** | **string** | Invoice ID | ☑️ |  | 

### Request Parameters - GetInvoiceByIdRequestParams

Parameters that are passed through a pointer to a apiGetInvoiceByIdRequest struct via the builder pattern

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
    invoice "github.com/xendit/xendit-go/v3/invoice"
)

func main() {
    
    // Invoice ID
    invoiceId := "62efe4c33e45294d63f585f2" // [REQUIRED] | string

    // Business ID of the sub-account merchant (XP feature)
    forUserId := "62efe4c33e45694d63f585f8" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.InvoiceApi.GetInvoiceById(context.Background(), invoiceId).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoiceById``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoiceById`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoiceById`: %v\n", resp)
}
```

## `GetInvoices()` Function

Get all Invoices

| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetInvoices` |
| Path Parameters  |  [GetInvoicesPathParams](#request-parameters--GetInvoicesPathParams)	 |
| Request Parameters  |  [GetInvoicesRequestParams](#request-parameters--GetInvoicesRequestParams)	 |
| Return Type  | [**[]Invoice**](invoice/Invoice.md) |

### Path Parameters - GetInvoicesPathParams
This endpoint does not need any path parameter.


### Request Parameters - GetInvoicesRequestParams

Parameters that are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **forUserId** |**string**|  |  | 
|  **externalId** |**string**|  |  | 
|  **statuses** |[**InvoiceStatus[]**](invoice/InvoiceStatus.md)|  |  | 
|  **limit** |**float32**|  |  | 
|  **createdAfter** |**time.Time**|  |  | 
|  **createdBefore** |**time.Time**|  |  | 
|  **paidAfter** |**time.Time**|  |  | 
|  **paidBefore** |**time.Time**|  |  | 
|  **expiredAfter** |**time.Time**|  |  | 
|  **expiredBefore** |**time.Time**|  |  | 
|  **lastInvoice** |**string**|  |  | 
|  **clientTypes** |[**InvoiceClientType[]**](invoice/InvoiceClientType.md)|  |  | 
|  **paymentChannels** |**string[]**|  |  | 
|  **onDemandLink** |**string**|  |  | 
|  **recurringPaymentId** |**string**|  |  | 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
        "time"
    xendit "github.com/xendit/xendit-go/v3"
    invoice "github.com/xendit/xendit-go/v3/invoice"
)

func main() {
    
    // Business ID of the sub-account merchant (XP feature)
    forUserId := "62efe4c33e45694d63f585f8" // [OPTIONAL] | string

    externalId := "test-external" // [OPTIONAL] | string

    statuses := []invoice.InvoiceStatus{invoice.InvoiceStatus("PENDING")} // [OPTIONAL] | []InvoiceStatus

    limit := float32(10) // [OPTIONAL] | float32

    createdAfter := time.Now() // [OPTIONAL] | time.Time

    createdBefore := time.Now() // [OPTIONAL] | time.Time

    paidAfter := time.Now() // [OPTIONAL] | time.Time

    paidBefore := time.Now() // [OPTIONAL] | time.Time

    expiredAfter := time.Now() // [OPTIONAL] | time.Time

    expiredBefore := time.Now() // [OPTIONAL] | time.Time

    lastInvoice := "62efe4c33e45294d63f585f2" // [OPTIONAL] | string

    clientTypes := []invoice.InvoiceClientType{invoice.InvoiceClientType("DASHBOARD")} // [OPTIONAL] | []InvoiceClientType

    paymentChannels := []string{"Inner_example"} // [OPTIONAL] | []string

    onDemandLink := "test-link" // [OPTIONAL] | string

    recurringPaymentId := "62efe4c33e45294d63f585f2" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.InvoiceApi.GetInvoices(context.Background()).
        ForUserId(forUserId).
        ExternalId(externalId).
        Statuses(statuses).
        Limit(limit).
        CreatedAfter(createdAfter).
        CreatedBefore(createdBefore).
        PaidAfter(paidAfter).
        PaidBefore(paidBefore).
        ExpiredAfter(expiredAfter).
        ExpiredBefore(expiredBefore).
        LastInvoice(lastInvoice).
        ClientTypes(clientTypes).
        PaymentChannels(paymentChannels).
        OnDemandLink(onDemandLink).
        RecurringPaymentId(recurringPaymentId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.GetInvoices``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoices`: []Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.GetInvoices`: %v\n", resp)
}
```

## `ExpireInvoice()` Function

Manually expire an invoice

| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `ExpireInvoice` |
| Path Parameters  |  [ExpireInvoicePathParams](#request-parameters--ExpireInvoicePathParams)	 |
| Request Parameters  |  [ExpireInvoiceRequestParams](#request-parameters--ExpireInvoiceRequestParams)	 |
| Return Type  | [**Invoice**](invoice/Invoice.md) |

### Path Parameters - ExpireInvoicePathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **invoiceId** | **string** | Invoice ID to be expired | ☑️ |  | 

### Request Parameters - ExpireInvoiceRequestParams

Parameters that are passed through a pointer to a apiExpireInvoiceRequest struct via the builder pattern

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
    invoice "github.com/xendit/xendit-go/v3/invoice"
)

func main() {
    
    // Invoice ID to be expired
    invoiceId := "5f4708b7bd394b0400b96276" // [REQUIRED] | string

    // Business ID of the sub-account merchant (XP feature)
    forUserId := "62efe4c33e45694d63f585f8" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.InvoiceApi.ExpireInvoice(context.Background(), invoiceId).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.ExpireInvoice``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpireInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.ExpireInvoice`: %v\n", resp)
}
```

[[Back to README]](../README.md)

# xendit\InvoiceApi

All URIs are relative to *https://api.xendit.co*

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**CreateInvoice**](InvoiceApi.md#CreateInvoice) | **Post** /v2/invoices/ | Create an invoice |
| [**GetInvoiceById**](InvoiceApi.md#GetInvoiceById) | **Get** /v2/invoices/{invoice_id} | Get invoice by invoice id |
| [**GetInvoices**](InvoiceApi.md#GetInvoices) | **Get** /v2/invoices | Get all Invoices |
| [**ExpireInvoice**](InvoiceApi.md#ExpireInvoice) | **Post** /invoices/{invoice_id}/expire! | Manually expire an invoice |



## CreateInvoice

Create an invoice

### Example

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

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvoiceRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
|  **createInvoiceRequest** |[**CreateInvoiceRequest**](invoice/CreateInvoiceRequest.md)|  |  | 
|  **forUserId** |**string**| Business ID of the sub-account merchant (XP feature) |  | 

### Return type

[**Invoice**](invoice/Invoice.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetInvoiceById

Get invoice by invoice id

### Example

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

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **invoiceId** | **string** | Invoice ID |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceByIdRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **forUserId** |**string**| Business ID of the sub-account merchant (XP feature) |  | 

### Return type

[**Invoice**](invoice/Invoice.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetInvoices

Get all Invoices

### Example

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

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
|  **forUserId** |**string**| Business ID of the sub-account merchant (XP feature) |  | 
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

### Return type

[**[]Invoice**](invoice/Invoice.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## ExpireInvoice

Manually expire an invoice

### Example

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

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **invoiceId** | **string** | Invoice ID to be expired |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpireInvoiceRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **forUserId** |**string**| Business ID of the sub-account merchant (XP feature) |  | 

### Return type

[**Invoice**](invoice/Invoice.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


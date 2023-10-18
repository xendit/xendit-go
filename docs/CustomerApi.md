# xendit\CustomerApi

All URIs are relative to *https://api.xendit.co*

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**CreateCustomer**](CustomerApi.md#CreateCustomer) | **Post** /customers | Create Customer |
| [**GetCustomer**](CustomerApi.md#GetCustomer) | **Get** /customers/{id} | Get Customer By ID |
| [**GetCustomerByReferenceID**](CustomerApi.md#GetCustomerByReferenceID) | **Get** /customers | GET customers by reference id |
| [**UpdateCustomer**](CustomerApi.md#UpdateCustomer) | **Patch** /customers/{id} | Update End Customer Resource |



## CreateCustomer

Create Customer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    customer "github.com/xendit/xendit-go/v3/customer"
)

func main() {
    
    // A unique key to prevent processing duplicate requests.
    idempotencyKey := "idempotency-123" // [OPTIONAL] | string

    // The sub-account user-id that you want to make this transaction for.
    forUserId := "user-1" // [OPTIONAL] | string

    // Request object for end customer object
    customerRequest := *customer.NewCustomerRequest("ReferenceId_example") // [OPTIONAL] | CustomerRequest

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.CustomerApi.CreateCustomer(context.Background()).
        IdempotencyKey(idempotencyKey).
        ForUserId(forUserId).
        CustomerRequest(customerRequest). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.CreateCustomer``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
|  **idempotencyKey** |**string**| A unique key to prevent processing duplicate requests. |  | 
|  **forUserId** |**string**| The sub-account user-id that you want to make this transaction for. |  | 
|  **customerRequest** |[**CustomerRequest**](customer/CustomerRequest.md)| Request object for end customer object |  | 

### Return type

[**Customer**](customer/Customer.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetCustomer

Get Customer By ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    customer "github.com/xendit/xendit-go/v3/customer"
)

func main() {
    
    // End customer resource id
    id := "d290f1ee-6c54-4b01-90e6-d701748f0851" // [REQUIRED] | string

    // The sub-account user-id that you want to make this transaction for.
    forUserId := "user-1" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.CustomerApi.GetCustomer(context.Background(), id).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.GetCustomer``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **id** | **string** | End customer resource id |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **forUserId** |**string**| The sub-account user-id that you want to make this transaction for. |  | 

### Return type

[**Customer**](customer/Customer.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetCustomerByReferenceID

GET customers by reference id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    customer "github.com/xendit/xendit-go/v3/customer"
)

func main() {
    
    // Merchant's reference of end customer
    referenceId := "referenceId_example" // [REQUIRED] | string

    // The sub-account user-id that you want to make this transaction for.
    forUserId := "user-1" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.CustomerApi.GetCustomerByReferenceID(context.Background()).
        ReferenceId(referenceId).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.GetCustomerByReferenceID``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerByReferenceID`: GetCustomerByReferenceID200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.GetCustomerByReferenceID`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerByReferenceIDRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
|  **referenceId** |**string**| Merchant&#39;s reference of end customer |  | 
|  **forUserId** |**string**| The sub-account user-id that you want to make this transaction for. |  | 

### Return type

[**GetCustomerByReferenceID200Response**](customer/GetCustomerByReferenceID200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## UpdateCustomer

Update End Customer Resource



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    customer "github.com/xendit/xendit-go/v3/customer"
)

func main() {
    
    // End customer resource id
    id := "d290f1ee-6c54-4b01-90e6-d701748f0851" // [REQUIRED] | string

    // The sub-account user-id that you want to make this transaction for.
    forUserId := "user-1" // [OPTIONAL] | string

    // Update Request for end customer object
    patchCustomer := *customer.NewPatchCustomer() // [OPTIONAL] | PatchCustomer

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.CustomerApi.UpdateCustomer(context.Background(), id).
        ForUserId(forUserId).
        PatchCustomer(patchCustomer). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerApi.UpdateCustomer``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomer`: Customer
    fmt.Fprintf(os.Stdout, "Response from `CustomerApi.UpdateCustomer`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **id** | **string** | End customer resource id |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **forUserId** |**string**| The sub-account user-id that you want to make this transaction for. |  | 
|  **patchCustomer** |[**PatchCustomer**](customer/PatchCustomer.md)| Update Request for end customer object |  | 

### Return type

[**Customer**](customer/Customer.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


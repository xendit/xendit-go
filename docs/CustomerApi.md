# CustomerApi


You can use the APIs below to interface with Xendit's `CustomerApi`.
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
| [**CreateCustomer**](CustomerApi.md#createcustomer-function) | **Post** /customers | Create Customer |
| [**GetCustomer**](CustomerApi.md#getcustomer-function) | **Get** /customers/{id} | Get Customer By ID |
| [**GetCustomerByReferenceID**](CustomerApi.md#getcustomerbyreferenceid-function) | **Get** /customers | GET customers by reference id |
| [**UpdateCustomer**](CustomerApi.md#updatecustomer-function) | **Patch** /customers/{id} | Update End Customer Resource |



## `CreateCustomer()` Function

Create Customer



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `CreateCustomer` |
| Path Parameters  |  [CreateCustomerPathParams](#request-parameters--CreateCustomerPathParams)	 |
| Request Parameters  |  [CreateCustomerRequestParams](#request-parameters--CreateCustomerRequestParams)	 |
| Return Type  | [**Customer**](customer/Customer.md) |

### Path Parameters - CreateCustomerPathParams
This endpoint does not need any path parameter.


### Request Parameters - CreateCustomerRequestParams

Parameters that are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **idempotencyKey** |**string**|  |  | 
|  **forUserId** |**string**|  |  | 
|  **customerRequest** |[**CustomerRequest**](customer/CustomerRequest.md)|  |  | 

### Usage Example

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

## `GetCustomer()` Function

Get Customer By ID



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetCustomer` |
| Path Parameters  |  [GetCustomerPathParams](#request-parameters--GetCustomerPathParams)	 |
| Request Parameters  |  [GetCustomerRequestParams](#request-parameters--GetCustomerRequestParams)	 |
| Return Type  | [**Customer**](customer/Customer.md) |

### Path Parameters - GetCustomerPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **id** | **string** | End customer resource id | ☑️ |  | 

### Request Parameters - GetCustomerRequestParams

Parameters that are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern

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

## `GetCustomerByReferenceID()` Function

GET customers by reference id



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetCustomerByReferenceID` |
| Path Parameters  |  [GetCustomerByReferenceIDPathParams](#request-parameters--GetCustomerByReferenceIDPathParams)	 |
| Request Parameters  |  [GetCustomerByReferenceIDRequestParams](#request-parameters--GetCustomerByReferenceIDRequestParams)	 |
| Return Type  | [**GetCustomerByReferenceID200Response**](customer/GetCustomerByReferenceID200Response.md) |

### Path Parameters - GetCustomerByReferenceIDPathParams
This endpoint does not need any path parameter.


### Request Parameters - GetCustomerByReferenceIDRequestParams

Parameters that are passed through a pointer to a apiGetCustomerByReferenceIDRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **referenceId** |**string**| ☑️ |  | 
|  **forUserId** |**string**|  |  | 

### Usage Example

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

## `UpdateCustomer()` Function

Update End Customer Resource



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `UpdateCustomer` |
| Path Parameters  |  [UpdateCustomerPathParams](#request-parameters--UpdateCustomerPathParams)	 |
| Request Parameters  |  [UpdateCustomerRequestParams](#request-parameters--UpdateCustomerRequestParams)	 |
| Return Type  | [**Customer**](customer/Customer.md) |

### Path Parameters - UpdateCustomerPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **id** | **string** | End customer resource id | ☑️ |  | 

### Request Parameters - UpdateCustomerRequestParams

Parameters that are passed through a pointer to a apiUpdateCustomerRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
| 
|  **forUserId** |**string**|  |  | 
|  **patchCustomer** |[**PatchCustomer**](customer/PatchCustomer.md)|  |  | 

### Usage Example

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


[[Back to README]](../README.md)

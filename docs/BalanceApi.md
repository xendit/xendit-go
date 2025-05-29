# BalanceApi


You can use the APIs below to interface with Xendit's `BalanceApi`.
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
| [**GetBalance**](BalanceApi.md#getbalance-function) | **Get** /balance | Retrieves balances for a business, default to CASH type |



## `GetBalance()` Function

Retrieves balances for a business, default to CASH type



| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetBalance` |
| Path Parameters  |  [GetBalancePathParams](#request-parameters--GetBalancePathParams)	 |
| Request Parameters  |  [GetBalanceRequestParams](#request-parameters--GetBalanceRequestParams)	 |
| Return Type  | [**Balance**](balance_and_transaction/Balance.md) |

### Path Parameters - GetBalancePathParams
This endpoint does not need any path parameter.


### Request Parameters - GetBalanceRequestParams

Parameters that are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **accountType** |**string**|  | [&quot;CASH&quot;] | 
|  **currency** |**string**|  |  | 
|  **atTimestamp** |**time.Time**|  |  | 
|  **forUserId** |**string**|  |  | 

### Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
        "time"
    xendit "github.com/xendit/xendit-go/v7"
    balance_and_transaction "github.com/xendit/xendit-go/v7/balance_and_transaction"
)

func main() {
    
    // The selected balance type
    // (default to "CASH")
    accountType := "CASH" // [OPTIONAL] | string

    // Currency for filter for customers with multi currency accounts
    currency := "IDR" // [OPTIONAL] | string

    // The timestamp you want to use as the limit for balance retrieval
    atTimestamp := time.Now() // [OPTIONAL] | time.Time

    // The sub-account user-id that you want to make this transaction for. This header
    // is only used if you have access to xenPlatform. See xenPlatform for more
    // information
    forUserId := "5dbf20d7c8eb0c0896f811b6" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.BalanceApi.GetBalance(context.Background()).
        AccountType(accountType).
        Currency(currency).
        AtTimestamp(atTimestamp).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalanceApi.GetBalance``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalance`: Balance
    fmt.Fprintf(os.Stdout, "Response from `BalanceApi.GetBalance`: %v\n", resp)
}
```


[[Back to README]](../README.md)

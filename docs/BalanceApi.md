# xendit\BalanceApi

All URIs are relative to *https://api.xendit.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBalance**](BalanceApi.md#GetBalance) | **Get** /balance | Retrieves balances for a business, default to CASH type



## GetBalance

Retrieves balances for a business, default to CASH type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    balance_and_transaction "github.com/xendit/xendit-go/v3/balance_and_transaction"
)

func main() {
    
    // The selected balance type 
    // (default to "CASH")
    accountType := "CASH" // [OPTIONAL] | string

    // Currency for filter for customers with multi currency accounts
    currency := "IDR" // [OPTIONAL] | string

    // The sub-account user-id that you want to make this transaction for. This header
    // is only used if you have access to xenPlatform. See xenPlatform for more
    // information
    forUserId := "5dbf20d7c8eb0c0896f811b6" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.BalanceApi.GetBalance(context.Background()).
        AccountType(accountType).
        Currency(currency).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalanceApi.GetBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBalance`: Balance
    fmt.Fprintf(os.Stdout, "Response from `BalanceApi.GetBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountType** | **string** | The selected balance type | [default to &quot;CASH&quot;]
 **currency** | **string** | Currency for filter for customers with multi currency accounts | 
 **forUserId** | **string** | The sub-account user-id that you want to make this transaction for. This header is only used if you have access to xenPlatform. See xenPlatform for more information | 

### Return type

[**Balance**](balance_and_transaction/Balance.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


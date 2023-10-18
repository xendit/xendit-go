# xendit\TransactionApi

All URIs are relative to *https://api.xendit.co*

| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**GetTransactionByID**](TransactionApi.md#GetTransactionByID) | **Get** /transactions/{id} | Get a transaction based on its id |
| [**GetAllTransactions**](TransactionApi.md#GetAllTransactions) | **Get** /transactions | Get a list of transactions |



## GetTransactionByID

Get a transaction based on its id



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
    
    id := "id_example" // [REQUIRED] | string

    // The sub-account user-id that you want to make this transaction for. This header
    // is only used if you have access to xenPlatform. See xenPlatform for more
    // information
    forUserId := "5dbf20d7c8eb0c0896f811b6" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.TransactionApi.GetTransactionByID(context.Background(), id).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.GetTransactionByID``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionByID`: TransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.GetTransactionByID`: %v\n", resp)
}
```

### Path Parameters


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | -------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| | 
| **id** | **string** |  |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByIDRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
| 
|  **forUserId** |**string**| The sub-account user-id that you want to make this transaction for. This header is only used if you have access to xenPlatform. See xenPlatform for more information |  | 

### Return type

[**TransactionResponse**](balance_and_transaction/TransactionResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetAllTransactions

Get a list of transactions



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
    
    // The sub-account user-id that you want to make this transaction for. This header
    // is only used if you have access to xenPlatform. See xenPlatform for more
    // information
    forUserId := "5dbf20d7c8eb0c0896f811b6" // [OPTIONAL] | string

    // Transaction types that will be included in the result. Default is to include all
    // transaction types
    types := []balance_and_transaction.TransactionTypes{balance_and_transaction.TransactionTypes("BATCH_DISBURSEMENT")} // [OPTIONAL] | []TransactionTypes

    // Status of the transaction. Default is to include all status.
    statuses := []balance_and_transaction.TransactionStatuses{balance_and_transaction.TransactionStatuses("SUCCESS")} // [OPTIONAL] | []TransactionStatuses

    // Payment channels in which the transaction is carried out. Default is to include
    // all channels.
    channelCategories := []balance_and_transaction.ChannelsCategories{balance_and_transaction.ChannelsCategories("BANK")} // [OPTIONAL] | []ChannelsCategories

    // To filter the result for transactions with matching reference given (case
    // sensitive)
    referenceId := "ref23232" // [OPTIONAL] | string

    // To filter the result for transactions with matching product_id (a.k.a payment_id)
    // given (case sensitive)
    productId := "d290f1ee-6c54-4b01-90e6-d701748f0701" // [OPTIONAL] | string

    // Account identifier of transaction. The format will be different from each
    // channel. For example, on `BANK` channel it will be account number and on `CARD`
    // it will be masked card number.
    accountIdentifier := "123123123" // [OPTIONAL] | string

    // Specific transaction amount to search for
    amount := float32(100) // [OPTIONAL] | float32

    currency := balance_and_transaction.Currency("IDR") // [OPTIONAL] | Currency

    // Filter time of transaction by created date. If not specified will list all dates.
    created := map[string][]balance_and_transaction.DateRangeFilter{"key": map[string]interface{}{ ... }} // [OPTIONAL] | DateRangeFilter

    // Filter time of transaction by updated date. If not specified will list all dates.
    updated := map[string][]balance_and_transaction.DateRangeFilter{"key": map[string]interface{}{ ... }} // [OPTIONAL] | DateRangeFilter

    // number of items in the result per page. Another name for \"results_per_page\"
    // (default to 10)
    limit := float32(10) // [OPTIONAL] | float32

    afterId := "afterId_example" // [OPTIONAL] | string

    beforeId := "beforeId_example" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.TransactionApi.GetAllTransactions(context.Background()).
        ForUserId(forUserId).
        Types(types).
        Statuses(statuses).
        ChannelCategories(channelCategories).
        ReferenceId(referenceId).
        ProductId(productId).
        AccountIdentifier(accountIdentifier).
        Amount(amount).
        Currency(currency).
        Created(created).
        Updated(updated).
        Limit(limit).
        AfterId(afterId).
        BeforeId(beforeId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionApi.GetAllTransactions``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTransactions`: TransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionApi.GetAllTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTransactionsRequest struct via the builder pattern


| Name | Type | Description  | Notes |
| ------------- | ------------- | ------------- | ------------- |
|  **forUserId** |**string**| The sub-account user-id that you want to make this transaction for. This header is only used if you have access to xenPlatform. See xenPlatform for more information |  | 
|  **types** |[**TransactionTypes[]**](balance_and_transaction/TransactionTypes.md)| Transaction types that will be included in the result. Default is to include all transaction types |  | 
|  **statuses** |[**TransactionStatuses[]**](balance_and_transaction/TransactionStatuses.md)| Status of the transaction. Default is to include all status. |  | 
|  **channelCategories** |[**ChannelsCategories[]**](balance_and_transaction/ChannelsCategories.md)| Payment channels in which the transaction is carried out. Default is to include all channels. |  | 
|  **referenceId** |**string**| To filter the result for transactions with matching reference given (case sensitive) |  | 
|  **productId** |**string**| To filter the result for transactions with matching product_id (a.k.a payment_id) given (case sensitive) |  | 
|  **accountIdentifier** |**string**| Account identifier of transaction. The format will be different from each channel. For example, on &#x60;BANK&#x60; channel it will be account number and on &#x60;CARD&#x60; it will be masked card number. |  | 
|  **amount** |**float32**| Specific transaction amount to search for |  | 
|  **currency** |[**Currency**](balance_and_transaction/Currencybalance_and_transaction/.md)|  |  | 
|  **created** |[**DateRangeFilter**](balance_and_transaction/DateRangeFilterbalance_and_transaction/.md)| Filter time of transaction by created date. If not specified will list all dates. |  | 
|  **updated** |[**DateRangeFilter**](balance_and_transaction/DateRangeFilterbalance_and_transaction/.md)| Filter time of transaction by updated date. If not specified will list all dates. |  | 
|  **limit** |**float32**| number of items in the result per page. Another name for \&quot;results_per_page\&quot; | [default to 10] | 
|  **afterId** |**string**|  |  | 
|  **beforeId** |**string**|  |  | 

### Return type

[**TransactionsResponse**](balance_and_transaction/TransactionsResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


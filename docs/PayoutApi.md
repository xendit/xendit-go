# xendit\PayoutApi

All URIs are relative to *https://api.xendit.co*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPayout**](PayoutApi.md#CancelPayout) | **Post** /v2/payouts/{id}/cancel | API to cancel requested payouts that have not yet been sent to partner banks and e-wallets. Cancellation is possible if the payout has not been sent out via our partner and when payout status is ACCEPTED.
[**CreatePayout**](PayoutApi.md#CreatePayout) | **Post** /v2/payouts | API to send money at scale to bank accounts &amp; eWallets
[**GetPayoutById**](PayoutApi.md#GetPayoutById) | **Get** /v2/payouts/{id} | API to fetch the current status, or details of the payout
[**GetPayoutChannels**](PayoutApi.md#GetPayoutChannels) | **Get** /payouts_channels | API providing the current list of banks and e-wallets we support for payouts for both regions
[**GetPayouts**](PayoutApi.md#GetPayouts) | **Get** /v2/payouts | API to retrieve all matching payouts with reference ID



## CancelPayout

API to cancel requested payouts that have not yet been sent to partner banks and e-wallets. Cancellation is possible if the payout has not been sent out via our partner and when payout status is ACCEPTED.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payout "github.com/xendit/xendit-go/v3/payout"
)

func main() {
    
    // Payout id returned from the response of /v2/payouts
    id := "disb-7baa7335-a0b2-4678-bb8c-318c0167f332" // [REQUIRED] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PayoutApi.CancelPayout(context.Background(), id). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.CancelPayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelPayout`: GetPayouts200ResponseDataInner
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.CancelPayout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payout id returned from the response of /v2/payouts | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPayouts200ResponseDataInner**](payout/GetPayouts200ResponseDataInner.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## CreatePayout

API to send money at scale to bank accounts & eWallets

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payout "github.com/xendit/xendit-go/v3/payout"
)

func main() {
    
    // A unique key to prevent duplicate requests from pushing through our system. No
    // expiration.
    idempotencyKey := "DISB-1234" // [REQUIRED] | string

    // The sub-account user-id that you want to make this transaction for. This header
    // is only used if you have access to xenPlatform. See xenPlatform for more
    // information.
    forUserId := "5dbf20d7c8eb0c0896f811b6" // [OPTIONAL] | string

    createPayoutRequest := *payout.NewCreatePayoutRequest("DISB-001", "PH_BDO", *payout.NewDigitalPayoutChannelProperties("9999999999"), float32(15000.05), "PHP") // [OPTIONAL] | CreatePayoutRequest

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PayoutApi.CreatePayout(context.Background()).
        IdempotencyKey(idempotencyKey).
        ForUserId(forUserId).
        CreatePayoutRequest(createPayoutRequest). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.CreatePayout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePayout`: GetPayouts200ResponseDataInner
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.CreatePayout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePayoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique key to prevent duplicate requests from pushing through our system. No expiration. | 
 **forUserId** | **string** | The sub-account user-id that you want to make this transaction for. This header is only used if you have access to xenPlatform. See xenPlatform for more information. | 
 **createPayoutRequest** | [**CreatePayoutRequest**](CreatePayoutRequest.md) |  | 

### Return type

[**GetPayouts200ResponseDataInner**](payout/GetPayouts200ResponseDataInner.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPayoutById

API to fetch the current status, or details of the payout

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payout "github.com/xendit/xendit-go/v3/payout"
)

func main() {
    
    // Payout id returned from the response of /v2/payouts
    id := "disb-7baa7335-a0b2-4678-bb8c-318c0167f332" // [REQUIRED] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PayoutApi.GetPayoutById(context.Background(), id). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.GetPayoutById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutById`: GetPayouts200ResponseDataInner
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.GetPayoutById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payout id returned from the response of /v2/payouts | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPayouts200ResponseDataInner**](payout/GetPayouts200ResponseDataInner.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPayoutChannels

API providing the current list of banks and e-wallets we support for payouts for both regions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payout "github.com/xendit/xendit-go/v3/payout"
)

func main() {
    
    // Filter channels by currency from ISO-4217 values
    currency := "IDR, PHP" // [OPTIONAL] | string

    // Filter channels by category
    channelCategory := []payout.ChannelCategory{payout.ChannelCategory("BANK")} // [OPTIONAL] | []ChannelCategory

    // Filter channels by channel code, prefixed by ISO-3166 country code
    channelCode := "ID_MANDIRI, PH_GCASH" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PayoutApi.GetPayoutChannels(context.Background()).
        Currency(currency).
        ChannelCategory(channelCategory).
        ChannelCode(channelCode). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.GetPayoutChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutChannels`: []Channel
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.GetPayoutChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Filter channels by currency from ISO-4217 values | 
 **channelCategory** | [**[]ChannelCategory**](payout/ChannelCategory.md) | Filter channels by category | 
 **channelCode** | **string** | Filter channels by channel code, prefixed by ISO-3166 country code | 

### Return type

[**[]Channel**](payout/Channel.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


## GetPayouts

API to retrieve all matching payouts with reference ID

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    xendit "github.com/xendit/xendit-go/v3"
    payout "github.com/xendit/xendit-go/v3/payout"
)

func main() {
    
    // Reference_id provided when creating the payout
    referenceId := "DISB-123" // [REQUIRED] | string

    // Number of records to fetch per API call
    limit := float32(10) // [OPTIONAL] | float32

    // Used to fetch record after this payout unique id
    afterId := "disb-7baa7335-a0b2-4678-bb8c-318c0167f332" // [OPTIONAL] | string

    // Used to fetch record before this payout unique id
    beforeId := "disb-7baa7335-a0b2-4678-bb8c-318c0167f332" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PayoutApi.GetPayouts(context.Background()).
        ReferenceId(referenceId).
        Limit(limit).
        AfterId(afterId).
        BeforeId(beforeId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.GetPayouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayouts`: GetPayouts200Response
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.GetPayouts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referenceId** | **string** | Reference_id provided when creating the payout | 
 **limit** | **float32** | Number of records to fetch per API call | 
 **afterId** | **string** | Used to fetch record after this payout unique id | 
 **beforeId** | **string** | Used to fetch record before this payout unique id | 

### Return type

[**GetPayouts200Response**](payout/GetPayouts200Response.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#)
[[Back to README]](../README.md)


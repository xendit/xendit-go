# PayoutApi


You can use the APIs below to interface with Xendit's `PayoutApi`.
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
| [**CreatePayout**](PayoutApi.md#createpayout-function) | **Post** /v2/payouts | API to send money at scale to bank accounts &amp; eWallets |
| [**GetPayoutById**](PayoutApi.md#getpayoutbyid-function) | **Get** /v2/payouts/{id} | API to fetch the current status, or details of the payout |
| [**GetPayoutChannels**](PayoutApi.md#getpayoutchannels-function) | **Get** /payouts_channels | API providing the current list of banks and e-wallets we support for payouts for both regions |
| [**GetPayouts**](PayoutApi.md#getpayouts-function) | **Get** /v2/payouts | API to retrieve all matching payouts with reference ID |
| [**CancelPayout**](PayoutApi.md#cancelpayout-function) | **Post** /v2/payouts/{id}/cancel | API to cancel requested payouts that have not yet been sent to partner banks and e-wallets. Cancellation is possible if the payout has not been sent out via our partner and when payout status is ACCEPTED. |



## `CreatePayout()` Function

API to send money at scale to bank accounts & eWallets

| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `CreatePayout` |
| Path Parameters  |  [CreatePayoutPathParams](#request-parameters--CreatePayoutPathParams)	 |
| Request Parameters  |  [CreatePayoutRequestParams](#request-parameters--CreatePayoutRequestParams)	 |
| Return Type  | [**GetPayouts200ResponseDataInner**](payout/GetPayouts200ResponseDataInner.md) |

### Path Parameters - CreatePayoutPathParams
This endpoint does not need any path parameter.


### Request Parameters - CreatePayoutRequestParams

Parameters that are passed through a pointer to a apiCreatePayoutRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **idempotencyKey** |**string**| ☑️ |  | 
|  **forUserId** |**string**|  |  | 
|  **createPayoutRequest** |[**CreatePayoutRequest**](payout/CreatePayoutRequest.md)|  |  | 

### Usage Example

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
    forUserId := "5f9a3fbd571a1c4068aa40ce" // [OPTIONAL] | string

    createPayoutRequest := *payout.NewCreatePayoutRequest("DISB-001", "PH_BDO", *payout.NewDigitalPayoutChannelProperties("9999999999"), float32(15000.05), "PHP") // [OPTIONAL] | CreatePayoutRequest

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PayoutApi.CreatePayout(context.Background()).
        IdempotencyKey(idempotencyKey).
        ForUserId(forUserId).
        CreatePayoutRequest(createPayoutRequest). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.CreatePayout``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePayout`: GetPayouts200ResponseDataInner
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.CreatePayout`: %v\n", resp)
}
```

## `GetPayoutById()` Function

API to fetch the current status, or details of the payout

| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetPayoutById` |
| Path Parameters  |  [GetPayoutByIdPathParams](#request-parameters--GetPayoutByIdPathParams)	 |
| Request Parameters  |  [GetPayoutByIdRequestParams](#request-parameters--GetPayoutByIdRequestParams)	 |
| Return Type  | [**GetPayouts200ResponseDataInner**](payout/GetPayouts200ResponseDataInner.md) |

### Path Parameters - GetPayoutByIdPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **id** | **string** | Payout id returned from the response of /v2/payouts | ☑️ |  | 

### Request Parameters - GetPayoutByIdRequestParams

Parameters that are passed through a pointer to a apiGetPayoutByIdRequest struct via the builder pattern

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
    payout "github.com/xendit/xendit-go/v3/payout"
)

func main() {
    
    // Payout id returned from the response of /v2/payouts
    id := "disb-7baa7335-a0b2-4678-bb8c-318c0167f332" // [REQUIRED] | string

    // The sub-account user-id that you want to make this transaction for. This header
    // is only used if you have access to xenPlatform. See xenPlatform for more
    // information.
    forUserId := "5f9a3fbd571a1c4068aa40ce" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PayoutApi.GetPayoutById(context.Background(), id).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.GetPayoutById``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutById`: GetPayouts200ResponseDataInner
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.GetPayoutById`: %v\n", resp)
}
```

## `GetPayoutChannels()` Function

API providing the current list of banks and e-wallets we support for payouts for both regions

| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetPayoutChannels` |
| Path Parameters  |  [GetPayoutChannelsPathParams](#request-parameters--GetPayoutChannelsPathParams)	 |
| Request Parameters  |  [GetPayoutChannelsRequestParams](#request-parameters--GetPayoutChannelsRequestParams)	 |
| Return Type  | [**[]Channel**](payout/Channel.md) |

### Path Parameters - GetPayoutChannelsPathParams
This endpoint does not need any path parameter.


### Request Parameters - GetPayoutChannelsRequestParams

Parameters that are passed through a pointer to a apiGetPayoutChannelsRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **currency** |**string**|  |  | 
|  **channelCategory** |[**ChannelCategory[]**](payout/ChannelCategory.md)|  |  | 
|  **channelCode** |**string**|  |  | 
|  **forUserId** |**string**|  |  | 

### Usage Example

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

    // The sub-account user-id that you want to make this transaction for. This header
    // is only used if you have access to xenPlatform. See xenPlatform for more
    // information.
    forUserId := "5f9a3fbd571a1c4068aa40ce" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PayoutApi.GetPayoutChannels(context.Background()).
        Currency(currency).
        ChannelCategory(channelCategory).
        ChannelCode(channelCode).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.GetPayoutChannels``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutChannels`: []Channel
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.GetPayoutChannels`: %v\n", resp)
}
```

## `GetPayouts()` Function

API to retrieve all matching payouts with reference ID

| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `GetPayouts` |
| Path Parameters  |  [GetPayoutsPathParams](#request-parameters--GetPayoutsPathParams)	 |
| Request Parameters  |  [GetPayoutsRequestParams](#request-parameters--GetPayoutsRequestParams)	 |
| Return Type  | [**GetPayouts200Response**](payout/GetPayouts200Response.md) |

### Path Parameters - GetPayoutsPathParams
This endpoint does not need any path parameter.


### Request Parameters - GetPayoutsRequestParams

Parameters that are passed through a pointer to a apiGetPayoutsRequest struct via the builder pattern

|Name | Type | Required |Default |
|-------------|:-------------:|:-------------:|-------------|
|  **referenceId** |**string**| ☑️ |  | 
|  **limit** |**float32**|  |  | 
|  **afterId** |**string**|  |  | 
|  **beforeId** |**string**|  |  | 
|  **forUserId** |**string**|  |  | 

### Usage Example

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

    // The sub-account user-id that you want to make this transaction for. This header
    // is only used if you have access to xenPlatform. See xenPlatform for more
    // information.
    forUserId := "5f9a3fbd571a1c4068aa40ce" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PayoutApi.GetPayouts(context.Background()).
        ReferenceId(referenceId).
        Limit(limit).
        AfterId(afterId).
        BeforeId(beforeId).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.GetPayouts``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayouts`: GetPayouts200Response
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.GetPayouts`: %v\n", resp)
}
```

## `CancelPayout()` Function

API to cancel requested payouts that have not yet been sent to partner banks and e-wallets. Cancellation is possible if the payout has not been sent out via our partner and when payout status is ACCEPTED.

| Name          |    Value 	     |
|--------------------|:-------------:|
| Function Name | `CancelPayout` |
| Path Parameters  |  [CancelPayoutPathParams](#request-parameters--CancelPayoutPathParams)	 |
| Request Parameters  |  [CancelPayoutRequestParams](#request-parameters--CancelPayoutRequestParams)	 |
| Return Type  | [**GetPayouts200ResponseDataInner**](payout/GetPayouts200ResponseDataInner.md) |

### Path Parameters - CancelPayoutPathParams


| Name | Type | Description | Required  | Default |
| ------------- |:-------------:| ------------- |:-------------:|-------------|
| **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.| ☑️ |  | 
| **id** | **string** | Payout id returned from the response of /v2/payouts | ☑️ |  | 

### Request Parameters - CancelPayoutRequestParams

Parameters that are passed through a pointer to a apiCancelPayoutRequest struct via the builder pattern

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
    payout "github.com/xendit/xendit-go/v3/payout"
)

func main() {
    
    // Payout id returned from the response of /v2/payouts
    id := "disb-7baa7335-a0b2-4678-bb8c-318c0167f332" // [REQUIRED] | string

    // The sub-account user-id that you want to make this transaction for. This header
    // is only used if you have access to xenPlatform. See xenPlatform for more
    // information.
    forUserId := "5f9a3fbd571a1c4068aa40ce" // [OPTIONAL] | string

    xenditClient := xendit.NewClient("API-KEY")

    resp, r, err := xenditClient.PayoutApi.CancelPayout(context.Background(), id).
        ForUserId(forUserId). // [OPTIONAL]
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.CancelPayout``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelPayout`: GetPayouts200ResponseDataInner
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.CancelPayout`: %v\n", resp)
}
```


[[Back to README]](../README.md)

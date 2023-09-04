![Xendit PHP SDK](docs/header.jpg "Xendit Go SDK")

# Xendit Go SDK

The official Xendit Go SDK provides a simple and convenient way to call Xendit's REST API
in applications written in Go.

* Package version: 3.0.0-beta.0

# Getting Started

## Installation

Install xendit-go with:

```shell
go get github.com/xendit/xendit-go/v3
```

Put the package under your project folder and add the following in import:

```golang
import xendit "github.com/xendit/xendit-go/v3"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Authorization

The SDK needs to be instantiated using your secret API key obtained from the [Xendit Dashboard](https://dashboard.xendit.co/settings/developers#api-keys).
You can sign up for a free Dashboard account [here](https://dashboard.xendit.co/register).

```golang
xnd := xendit.NewClient("API-KEY")
```

# Documentation

Find detailed API infomration and examples for each of our product's by clicking the links below,

* [Balance](docs/BalanceApi.md)
* [Invoice](docs/InvoiceApi.md)
* [PaymentMethod](docs/PaymentMethodApi.md)
* [PaymentRequest](docs/PaymentRequestApi.md)
* [Payout](docs/PayoutApi.md)
* [Refund](docs/RefundApi.md)
* [Transaction](docs/TransactionApi.md)

All URIs are relative to *https://api.xendit.co*.  For more information about our API, please refer to *https://developers.xendit.co/*.

Further Reading

* [Xendit Docs](https://docs.xendit.co/)
* [Xendit API Reference](https://developers.xendit.co/)
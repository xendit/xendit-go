# Xendit API Go Client

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/xendit/xendit-go)

This library is the abstraction of Xendit API for access from applications written with Go.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Documentation](#documentation)
- [Installation](#installation)
  - [Go Module Support](#go-module-support)
- [Usage](#usage)
  - [Without Client](#without-client)
  - [With Client](#with-client)
  - [Sub-Packages Documentations](#sub-packages-documentations)
    - [Invoice](#invoice)
    - [E-Wallet](#e-wallet)
    - [Balance](#balance)
    - [Virtual Account](#virtual-account)
    - [Retail Outlet](#retail-outlet)
    - [Disbursement](#disbursement)
    - [Card](#card)
    - [Payout](#payout)
    - [Recurring Payment](#recurring-payment)
    - [Cardless Credit](#cardless-credit)
- [Contribute](#contribute)
  - [Test](#test)
    - [Run all tests](#run-all-tests)
    - [Run tests for a package](#run-tests-for-a-package)
    - [Run a single test](#run-a-single-test)
  - [Pre-commit](#pre-commit)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Documentation

For the API documentation, check [Xendit API Reference](https://xendit.github.io/apireference).

For the details of this library, see the [GoDoc](https://pkg.go.dev/github.com/xendit/xendit-go).

## Installation

Install xendit-go with:

```sh
go get -u github.com/xendit/xendit-go
```

Then, import it using:

```go
import (
    "github.com/xendit/xendit-go"
    "github.com/xendit/xendit-go/$product$"
)
```

with `$product$` is the product of Xendit such as `invoice` and `balance`.

### Go Module Support

This library can also be included via Go modules. To do so, require xendit-go in `go.mod` with a version like so:

```go
module github.com/my/package

go 1.13

require (
  github.com/xendit/xendit-go v1.0.0
)
```

And use the same style of import paths as above:

```go
import (
    "github.com/xendit/xendit-go"
    "github.com/xendit/xendit-go/$product$"
)
```

with `$product$` is the product of Xendit such as `invoice` and `balance`.

## Usage

The following pattern is applied throughout the library for a given `$product$`:

### Without Client

If you're only dealing with a single `secret key`, you can simply import the packages required for the products you're interacting with without the need to create a client.

```go
import (
    "github.com/xendit/xendit-go"
    "github.com/xendit/xendit-go/$product$"
)

// Setup
xendit.Opt.SecretKey = "examplesecretkey"

xendit.SetAPIRequester(apiRequester) // optional, useful for mocking

// Create
resp, err := $product$.Create($product$.CreateParams)

// Get
resp, err := $product$.Get($product$.GetParams)

// GetAll
resp, err := $product$.GetAll($product$.GetAllParams)
```

### With Client

If you're dealing with multiple `secret key`s, it is recommended you use `client.API`. This allows you to create as many clients as needed, each with their own individual key.

```go
import (
    "github.com/xendit/xendit-go"
    "github.com/xendit/xendit-go/client"
)

// Basic setup
client := client.New("examplesecretkey")

// or with optional, useful-for-mocking `exampleAPIRequester`
client := client.New("examplesecretkey").WithAPIRequester(exampleAPIRequester)

// Create
resp, err := client.$product$.Create($product$.CreateParams)

// Get
resp, err := client.$product$.Get($product$.GetParams)

// GetAll
resp, err := client.$product$.GetAll($product$.GetAllParams)
```

### Sub-Packages Documentations

The following is a list of pointers to documentations for sub-packages of [xendit-go](https://github.com/xendit/xendit-go).

#### Invoice

Find invoice documentation in [go.dev](https://pkg.go.dev/github.com/xendit/xendit-go/invoice).

#### E-Wallet

Find ewallet documentation in [go.dev](https://pkg.go.dev/github.com/xendit/xendit-go/ewallet).

#### Balance

Find balance documentation in [go.dev](https://pkg.go.dev/github.com/xendit/xendit-go/balance).

#### Virtual Account

Find virtualaccount documentation in [go.dev](https://pkg.go.dev/github.com/xendit/xendit-go/virtualaccount).

#### Retail Outlet

Find retailoutlet documentation in [go.dev](https://pkg.go.dev/github.com/xendit/xendit-go/retailoutlet).

#### Disbursement

Find disbursement documentation in [go.dev](https://pkg.go.dev/github.com/xendit/xendit-go/disbursement).

#### Card

Find card documentation in [go.dev](https://pkg.go.dev/github.com/xendit/xendit-go/card).

#### Payout

Find payout documentation in [go.dev](https://pkg.go.dev/github.com/xendit/xendit-go/payout).

#### Recurring Payment

Find recurringpayment documentation in [go.dev](https://pkg.go.dev/github.com/xendit/xendit-go/recurringpayment).

#### Cardless Credit

Find cardlesscredit documentation in [go.dev](https://pkg.go.dev/github.com/xendit/xendit-go/cardlesscredit).

## Contribute

For any requests, bugs, or comments, please [open an issue](https://github.com/xendit/xendit-go/issues/new) or submit a pull request.

### Test

After modifying the code, please make sure that the code passes all test cases.

#### Run all tests

```sh
go test ./...
```

#### Run tests for a package

```sh
go test ./invoice
```

#### Run a single test

```sh
go test ./invoice -run TestCreateInvoice
```

### Pre-commit

Before making any commits, please install pre-commit.
To install pre-commit, follow the [installation steps](https://pre-commit.com/#install).

After installing the pre-commit, please install the needed dependencies:

```sh
make init
```

After the code passes everything, please [submit a pull request](https://github.com/xendit/xendit-go/pulls).

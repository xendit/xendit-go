# Xendit API Go Client

This library is the abstraction of Xendit API for access from applications written with Go.

---

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Documentation](#documentation)
- [Installation](#installation)
  - [Go Module Support](#go-module-support)
- [Usage](#usage)
  - [Without Client](#without-client)
  - [With Client](#with-client)
  - [Packages Method Signatures](#packages-method-signatures)
    - [Invoice](#invoice)
- [Contribute](#contribute)
  - [Test](#test)
    - [Run all tests](#run-all-tests)
    - [Run tests for a package](#run-tests-for-a-package)
    - [Run a single test](#run-a-single-test)
  - [Pre-commit](#pre-commit)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

---

## Documentation

For the API documentation, check [Xendit API Reference](https://xendit.github.io/apireference).

For the details of this library, see the [GoDoc](http://godoc.org/github.com/xendit/xendit-go).

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
resp, err := $product$.Get($product$.GetAllParams)
```

### With Client

If you're dealing with multiple `secret key`s, it is recommended you use `client.API`. This allows you to create as many clients as needed, each with their own individual key.

```go
import (
    "github.com/xendit/xendit-go"
    "github.com/xendit/xendit-go/client"
)

// Setup
client := client.New("examplepublickey", "examplesecretkey", "", exampleAPIRequester)
// the `exampleAPIRequester is optional, useful for mocking

// Create
resp, err := client.$product$.Create($product$.CreateParams)

// Get
resp, err := client.$product$.Get($product$.GetParams)

// GetAll
resp, err := client.$product$.Get($product$.GetAllParams)
```

### Packages Method Signatures

The following is a list of method signatures for each packages (for quick reference)

#### Invoice

Create Invoice

```go
invoice.Create(data *invoice.CreateParams) (*xendit.Invoice, *xendit.Error)
```

Get Invoice

```go
invoice.Get(data *invoice.GetParams) (*xendit.Invoice, *xendit.Error)
```

Expire Invoice

```go
invoice.Expire(data *invoice.ExpireParams) (*xendit.Invoice, *xendit.Error)
```

GetAll Invoice

```go
invoice.GetAll(data *invoice.GetAll) (*xendit.Invoice, *xendit.Error)
```

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

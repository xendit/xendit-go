# Xendit API Go Client

This library is the abstraction of Xendit API for access from applications written with Go.

---

<!--toc-->

---

## Documentation

For the API documentation, check [Xendit API Reference](https://xendit.github.io/apireference).

For the details of this library, see the GoDoc.

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

To submit a pull request, please clone this project with https:

```sh
git clone https://github.com/xendit/xendit-go.git
```

or with ssh:

```sh
git clone git@github.com:xendit/xendit-go.git
```

After modifying the code, please make sure that the code passes all test cases.

### Test

#### Run all tests:

```sh
go test ./...
```

#### Run tests for a packages:

```sh
go test ./invoice
```

#### Run a single test:

```sh
go test ./invoice -run TestCreateInvoice
```

After passing all the test cases, please [submit a pull request](https://github.com/xendit/xendit-go/pulls).

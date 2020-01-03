package invoice

import (
	"context"
	"errors"
	"fmt"

	"github.com/xendit/xendit-go"

	validator "github.com/go-playground/validator"
)

// Client is the client used to invoke invoice API.
type Client struct {
	Opt            *xendit.Option
	XdAPIRequester xendit.XdAPIRequester
}

// CreateInvoice creates new invoice
func CreateInvoice(data *xendit.CreateInvoiceParams) (*xendit.Invoice, error) {
	return CreateInvoiceWithContext(nil, data)
}

// CreateInvoice creates new invoice
func (c Client) CreateInvoice(data *xendit.CreateInvoiceParams) (*xendit.Invoice, error) {
	return c.CreateInvoiceWithContext(nil, data)
}

// CreateInvoiceWithContext creates new invoice with context
func CreateInvoiceWithContext(ctx context.Context, data *xendit.CreateInvoiceParams) (*xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateInvoiceWithContext(ctx, data)
}

// CreateInvoiceWithContext creates new invoice with context
func (c Client) CreateInvoiceWithContext(ctx context.Context, data *xendit.CreateInvoiceParams) (*xendit.Invoice, error) {
	response := &xendit.Invoice{}

	v := validator.New()
	if err := v.Struct(data); err != nil {
		return nil, err
	}

	err := c.XdAPIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/v2/invoices", c.Opt.XenditURL),
		c.Opt.SecretKey,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetInvoice gets one invoice
func GetInvoice(invoiceID string) (*xendit.Invoice, error) {
	return GetInvoiceWithContext(nil, invoiceID)
}

// GetInvoice gets one invoice
func (c Client) GetInvoice(invoiceID string) (*xendit.Invoice, error) {
	return c.GetInvoiceWithContext(nil, invoiceID)
}

// GetInvoiceWithContext gets one invoice with context
func GetInvoiceWithContext(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetInvoiceWithContext(ctx, invoiceID)
}

// GetInvoiceWithContext gets one invoice with context
func (c Client) GetInvoiceWithContext(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	response := &xendit.Invoice{}

	err := c.XdAPIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/invoices/%s", c.Opt.XenditURL, invoiceID),
		c.Opt.SecretKey,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ExpireInvoice expire the created invoice
func ExpireInvoice(invoiceID string) (*xendit.Invoice, error) {
	return ExpireInvoiceWithContext(nil, invoiceID)
}

// ExpireInvoice expire the created invoice
func (c Client) ExpireInvoice(invoiceID string) (*xendit.Invoice, error) {
	return c.ExpireInvoiceWithContext(nil, invoiceID)
}

// ExpireInvoiceWithContext expire the created invoice with context
func ExpireInvoiceWithContext(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ExpireInvoiceWithContext(ctx, invoiceID)
}

// ExpireInvoiceWithContext expire the created invoice with context
func (c Client) ExpireInvoiceWithContext(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	response := &xendit.Invoice{}

	err := c.XdAPIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/invoices/%s/expire!", c.Opt.XenditURL, invoiceID),
		c.Opt.SecretKey,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAllInvoices gets all invoices with conditions
func GetAllInvoices(data *xendit.GetAllInvoicesParams) ([]xendit.Invoice, error) {
	return GetAllInvoicesWithContext(nil, data)
}

// GetAllInvoices gets all invoices with conditions
func (c Client) GetAllInvoices(data *xendit.GetAllInvoicesParams) ([]xendit.Invoice, error) {
	return c.GetAllInvoicesWithContext(nil, data)
}

// GetAllInvoicesWithContext gets all invoices with conditions
func GetAllInvoicesWithContext(ctx context.Context, data *xendit.GetAllInvoicesParams) ([]xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetAllInvoicesWithContext(ctx, data)
}

// GetAllInvoicesWithContext gets all invoices with conditions
func (c Client) GetAllInvoicesWithContext(ctx context.Context, data *xendit.GetAllInvoicesParams) ([]xendit.Invoice, error) {
	response := []xendit.Invoice{}

	err := c.XdAPIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/invoices", c.Opt.XenditURL),
		c.Opt.SecretKey,
		data,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func getClient() (Client, error) {
	if xendit.Opt.SecretKey == "" {
		return Client{}, errors.New("secret key is not allowed to be empty")
	}

	return Client{
		Opt:            &xendit.Opt,
		XdAPIRequester: xendit.GetXdAPIRequester(),
	}, nil
}

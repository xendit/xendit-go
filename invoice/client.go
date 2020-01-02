package invoice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xendit/xendit-go"

	validator "github.com/go-playground/validator"
)

// Client is the client used to invoke invoice API.
type Client struct {
	Opt           *xendit.Option
	HTTPRequester xendit.HTTPRequester
}

// CreateInvoice creates new invoice
func CreateInvoice(ctx context.Context, data *xendit.CreateInvoiceParams) (*xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateInvoice(ctx, data)
}

// CreateInvoice creates new invoice
func (c Client) CreateInvoice(ctx context.Context, data *xendit.CreateInvoiceParams) (*xendit.Invoice, error) {
	var response *xendit.Invoice

	v := validator.New()
	if err := v.Struct(data); err != nil {
		return nil, err
	}

	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	respBody, err := c.HTTPRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/v2/invoices", c.Opt.XenditURL),
		c.Opt.SecretKey,
		string(reqBody),
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetInvoice gets one invoice
func GetInvoice(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetInvoice(ctx, invoiceID)
}

// GetInvoice gets one invoice
func (c Client) GetInvoice(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	var response *xendit.Invoice

	respBody, err := c.HTTPRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/invoices/%s", c.Opt.XenditURL, invoiceID),
		c.Opt.SecretKey,
		"",
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// ExpireInvoice expire the created invoice
func ExpireInvoice(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ExpireInvoice(ctx, invoiceID)
}

// ExpireInvoice expire the created invoice
func (c Client) ExpireInvoice(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	var response *xendit.Invoice

	respBody, err := c.HTTPRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/invoices/%s/expire!", c.Opt.XenditURL, invoiceID),
		c.Opt.SecretKey,
		"",
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetAllInvoices gets all invoices with conditions
func GetAllInvoices(ctx context.Context, data *xendit.GetAllInvoicesParams) ([]xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetAllInvoices(ctx, data)
}

// GetAllInvoices gets all invoices with conditions
func (c Client) GetAllInvoices(ctx context.Context, data *xendit.GetAllInvoicesParams) ([]xendit.Invoice, error) {
	var responses []xendit.Invoice

	reqParams, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	respBody, err := c.HTTPRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/invoices", c.Opt.XenditURL),
		c.Opt.SecretKey,
		string(reqParams),
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(respBody, &responses); err != nil {
		return nil, err
	}

	return responses, nil
}

func getClient() (Client, error) {
	if xendit.Opt.SecretKey == "" {
		return Client{}, errors.New("secret key is not allowed to be empty")
	}

	return Client{
		Opt:           &xendit.Opt,
		HTTPRequester: xendit.GetHTTPRequester(),
	}, nil
}

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
	XdAPIRequester xendit.APIRequester
}

// Create creates new invoice
func Create(data *CreateParams) (*xendit.Invoice, error) {
	return CreateWithContext(context.Background(), data)
}

// Create creates new invoice
func (c Client) Create(data *CreateParams) (*xendit.Invoice, error) {
	return c.CreateWithContext(context.Background(), data)
}

// CreateWithContext creates new invoice with context
func CreateWithContext(ctx context.Context, data *CreateParams) (*xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateWithContext(ctx, data)
}

// CreateWithContext creates new invoice with context
func (c Client) CreateWithContext(ctx context.Context, data *CreateParams) (*xendit.Invoice, error) {
	v := validator.New()
	if err := v.Struct(data); err != nil {
		return nil, err
	}

	response := &xendit.Invoice{}

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

// Get gets one invoice
func Get(invoiceID string) (*xendit.Invoice, error) {
	return GetWithContext(context.Background(), invoiceID)
}

// Get gets one invoice
func (c Client) Get(invoiceID string) (*xendit.Invoice, error) {
	return c.GetWithContext(context.Background(), invoiceID)
}

// GetWithContext gets one invoice with context
func GetWithContext(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetWithContext(ctx, invoiceID)
}

// GetWithContext gets one invoice with context
func (c Client) GetWithContext(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	v := validator.New()
	if err := v.Var(invoiceID, "required"); err != nil {
		return nil, err
	}

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

// Expire expire the created invoice
func Expire(invoiceID string) (*xendit.Invoice, error) {
	return ExpireWithContext(context.Background(), invoiceID)
}

// Expire expire the created invoice
func (c Client) Expire(invoiceID string) (*xendit.Invoice, error) {
	return c.ExpireWithContext(context.Background(), invoiceID)
}

// ExpireWithContext expire the created invoice with context
func ExpireWithContext(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ExpireWithContext(ctx, invoiceID)
}

// ExpireWithContext expire the created invoice with context
func (c Client) ExpireWithContext(ctx context.Context, invoiceID string) (*xendit.Invoice, error) {
	v := validator.New()
	if err := v.Var(invoiceID, "required"); err != nil {
		return nil, err
	}

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

// GetAll gets all invoices with conditions
func GetAll(data *GetAllParams) ([]xendit.Invoice, error) {
	return GetAllWithContext(context.Background(), data)
}

// GetAll gets all invoices with conditions
func (c Client) GetAll(data *GetAllParams) ([]xendit.Invoice, error) {
	return c.GetAllWithContext(context.Background(), data)
}

// GetAllWithContext gets all invoices with conditions
func GetAllWithContext(ctx context.Context, data *GetAllParams) ([]xendit.Invoice, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetAllWithContext(ctx, data)
}

// GetAllWithContext gets all invoices with conditions
func (c Client) GetAllWithContext(ctx context.Context, data *GetAllParams) ([]xendit.Invoice, error) {
	response := []xendit.Invoice{}
	var queryString string

	if data != nil {
		queryString = data.QueryString()
	}

	err := c.XdAPIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/invoices?%s", c.Opt.XenditURL, queryString),
		c.Opt.SecretKey,
		data,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func getClient() (*Client, error) {
	if xendit.Opt.SecretKey == "" {
		return nil, errors.New("secret key is not allowed to be empty")
	}

	return &Client{
		Opt:            &xendit.Opt,
		XdAPIRequester: xendit.GetAPIRequester(),
	}, nil
}

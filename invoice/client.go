package invoice

import (
	"context"
	"fmt"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke invoice API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// Create creates new invoice
func Create(data *CreateParams) (*xendit.Invoice, *xendit.Error) {
	return CreateWithContext(context.Background(), data)
}

// Create creates new invoice
func (c Client) Create(data *CreateParams) (*xendit.Invoice, *xendit.Error) {
	return c.CreateWithContext(context.Background(), data)
}

// CreateWithContext creates new invoice with context
func CreateWithContext(ctx context.Context, data *CreateParams) (*xendit.Invoice, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateWithContext(ctx, data)
}

// CreateWithContext creates new invoice with context
func (c Client) CreateWithContext(ctx context.Context, data *CreateParams) (*xendit.Invoice, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Invoice{}

	err := c.APIRequester.Call(
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
func Get(data *GetParams) (*xendit.Invoice, *xendit.Error) {
	return GetWithContext(context.Background(), data)
}

// Get gets one invoice
func (c Client) Get(data *GetParams) (*xendit.Invoice, *xendit.Error) {
	return c.GetWithContext(context.Background(), data)
}

// GetWithContext gets one invoice with context
func GetWithContext(ctx context.Context, data *GetParams) (*xendit.Invoice, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetWithContext(ctx, data)
}

// GetWithContext gets one invoice with context
func (c Client) GetWithContext(ctx context.Context, data *GetParams) (*xendit.Invoice, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Invoice{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/invoices/%s", c.Opt.XenditURL, data.ID),
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
func Expire(data *ExpireParams) (*xendit.Invoice, *xendit.Error) {
	return ExpireWithContext(context.Background(), data)
}

// Expire expire the created invoice
func (c Client) Expire(data *ExpireParams) (*xendit.Invoice, *xendit.Error) {
	return c.ExpireWithContext(context.Background(), data)
}

// ExpireWithContext expire the created invoice with context
func ExpireWithContext(ctx context.Context, data *ExpireParams) (*xendit.Invoice, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ExpireWithContext(ctx, data)
}

// ExpireWithContext expire the created invoice with context
func (c Client) ExpireWithContext(ctx context.Context, data *ExpireParams) (*xendit.Invoice, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Invoice{}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/invoices/%s/expire!", c.Opt.XenditURL, data.ID),
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
func GetAll(data *GetAllParams) ([]xendit.Invoice, *xendit.Error) {
	return GetAllWithContext(context.Background(), data)
}

// GetAll gets all invoices with conditions
func (c Client) GetAll(data *GetAllParams) ([]xendit.Invoice, *xendit.Error) {
	return c.GetAllWithContext(context.Background(), data)
}

// GetAllWithContext gets all invoices with conditions
func GetAllWithContext(ctx context.Context, data *GetAllParams) ([]xendit.Invoice, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetAllWithContext(ctx, data)
}

// GetAllWithContext gets all invoices with conditions
func (c Client) GetAllWithContext(ctx context.Context, data *GetAllParams) ([]xendit.Invoice, *xendit.Error) {
	response := []xendit.Invoice{}
	var queryString string

	if data != nil {
		queryString = data.QueryString()
	}

	err := c.APIRequester.Call(
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

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}

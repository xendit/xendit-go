package customer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke customer API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// CreateCustomer creates new customer
func (c *Client) CreateCustomer(data *CreateCustomerParams) (*xendit.Customer, *xendit.Error) {
	return c.CreateCustomerWithContext(context.Background(), data)
}

// CreateCustomerWithContext creates new customer
func (c *Client) CreateCustomerWithContext(ctx context.Context, data *CreateCustomerParams) (*xendit.Customer, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Customer{}
	header := &http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/customers", c.Opt.XenditURL),
		c.Opt.SecretKey,
		header,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetCustomerByReferenceID gets customer by reference ID
func (c *Client) GetCustomerByReferenceID(data *GetCustomerByReferenceIDParams) ([]xendit.Customer, *xendit.Error) {
	return c.GetCustomerByReferenceIDWithContext(context.Background(), data)
}

// GetCustomerByReferenceIDWithContext gets customer by reference ID
func (c *Client) GetCustomerByReferenceIDWithContext(ctx context.Context, data *GetCustomerByReferenceIDParams) ([]xendit.Customer, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := []xendit.Customer{}
	var queryString string

	if data != nil {
		queryString = data.QueryString()
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/customers?%s", c.Opt.XenditURL, queryString),
		c.Opt.SecretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

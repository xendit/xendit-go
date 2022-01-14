package paymentmethod

import (
	"context"
	"fmt"
	"net/http"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke e-wallet API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// CreatePaymentMethod creates new payment method
func (c *Client) CreatePaymentMethod(data *CreatePaymentMethodParams) (*xendit.PaymentMethod, *xendit.Error) {
	return c.CreatePaymentMethodWithContext(context.Background(), data)
}

// CreatePaymentMethodWithContext creates new payment method
func (c *Client) CreatePaymentMethodWithContext(ctx context.Context, data *CreatePaymentMethodParams) (*xendit.PaymentMethod, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.PaymentMethod{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/payment_methods", c.Opt.XenditURL),
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

// GetPaymentMethodsByCustomerIDParams gets payment methods by customer ID
func (c *Client) GetPaymentMethodsByCustomerID(data *GetPaymentMethodsByCustomerIDParams) ([]xendit.PaymentMethod, *xendit.Error) {
	return c.GetPaymentMethodsByCustomerIDWithContext(context.Background(), data)
}

// GetPaymentMethodsByCustomerIDWithContext gets payment methods by customer ID
func (c *Client) GetPaymentMethodsByCustomerIDWithContext(ctx context.Context, data *GetPaymentMethodsByCustomerIDParams) ([]xendit.PaymentMethod, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := []xendit.PaymentMethod{}
	var queryString string

	if data != nil {
		queryString = data.QueryString()
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/payment_methods?%s", c.Opt.XenditURL, queryString),
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

package virtualaccount

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

// CreateFixed creates new fixed virtual account
func (c Client) CreateFixed(data *CreateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	return c.CreateFixedWithContext(context.Background(), data)
}

// CreateFixedWithContext creates new fixed virtual account with context
func (c *Client) CreateFixedWithContext(ctx context.Context, data *CreateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.VirtualAccount{}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/callback_virtual_accounts", c.Opt.XenditURL),
		c.Opt.SecretKey,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetFixed gets one fixed virtual account
func (c *Client) GetFixed(data *GetFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	return c.GetFixedWithContext(context.Background(), data)
}

// GetFixedWithContext gets one invoice with context
func (c *Client) GetFixedWithContext(ctx context.Context, data *GetFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.VirtualAccount{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/callback_virtual_accounts/%s", c.Opt.XenditURL, data.ID),
		c.Opt.SecretKey,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateFixed updates one fixed virtual account
func (c Client) UpdateFixed(data *UpdateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	return c.UpdateFixedWithContext(context.Background(), data)
}

// UpdateFixedWithContext updates one fixed virtual account with context
func (c *Client) UpdateFixedWithContext(ctx context.Context, data *UpdateFixedParams) (*xendit.VirtualAccount, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.VirtualAccount{}

	err := c.APIRequester.Call(
		ctx,
		"PATCH",
		fmt.Sprintf("%s/callback_virtual_accounts/%s", c.Opt.XenditURL, data.ID),
		c.Opt.SecretKey,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAvailableBanks gets available virtual account banks
func (c *Client) GetAvailableBanks() ([]xendit.VirtualAccountBank, *xendit.Error) {
	return c.GetAvailableBanksWithContext(context.Background())
}

// GetAvailableBanksWithContext gets available virtual account banks with context
func (c *Client) GetAvailableBanksWithContext(ctx context.Context) ([]xendit.VirtualAccountBank, *xendit.Error) {
	response := []xendit.VirtualAccountBank{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/available_virtual_account_banks", c.Opt.XenditURL),
		c.Opt.SecretKey,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetPayment gets one fixed virtual account payment
func (c *Client) GetPayment(data *GetPaymentParams) (*xendit.VirtualAccountPayment, *xendit.Error) {
	return c.GetPaymentWithContext(context.Background(), data)
}

// GetPaymentWithContext gets one fixed virtual account payment with context
func (c *Client) GetPaymentWithContext(ctx context.Context, data *GetPaymentParams) (*xendit.VirtualAccountPayment, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.VirtualAccountPayment{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/callback_virtual_account_payments/payment_id=%s", c.Opt.XenditURL, data.PaymentID),
		c.Opt.SecretKey,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

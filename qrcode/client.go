package qrcode

import (
	"context"
	"fmt"
	"net/http"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke QRCode API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// CreateQRCode creates new QR Code
func (c *Client) CreateQRCode(data *CreateQRCodeParams) (*xendit.QRCode, *xendit.Error) {
	return c.CreateQRCodeWithContext(context.Background(), data)
}

// CreateQRCodeWithContext creates new QR Code with context
func (c *Client) CreateQRCodeWithContext(ctx context.Context, data *CreateQRCodeParams) (*xendit.QRCode, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.QRCode{}
	header := &http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/qr_codes", c.Opt.XenditURL),
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

// GetQRCode gets a single QRCode
func (c *Client) GetQRCode(data *GetQRCodeParams) (*xendit.QRCode, *xendit.Error) {
	return c.GetQRCodeWithContext(context.Background(), data)
}

// GetQRCodeWithContext gets a single QRCode with context
func (c *Client) GetQRCodeWithContext(ctx context.Context, data *GetQRCodeParams) (*xendit.QRCode, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.QRCode{}
	header := &http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/qr_codes/%s", c.Opt.XenditURL, data.ExternalID),
		c.Opt.SecretKey,
		header,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetQRCodePayments gets a list of QR Code payments
func (c *Client) GetQRCodePayments(data *GetQRCodePaymentsParams) ([]xendit.QRCodePayments, *xendit.Error) {
	return c.GetQRCodePaymentsWithContext(context.Background(), data)
}

// GetQRCodePaymentsWithContext gets a list of QR Code payments with context
func (c *Client) GetQRCodePaymentsWithContext(ctx context.Context, data *GetQRCodePaymentsParams) ([]xendit.QRCodePayments, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := []xendit.QRCodePayments{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/qr_codes/payments?%s", c.Opt.XenditURL, data.QueryString()),
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

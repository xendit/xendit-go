package disbursementph

import (
	"context"
	"fmt"
	"net/http"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke invoice API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// Create creates new PHP disbursement
func (c *Client) Create(data *CreateParams) (*xendit.DisbursementPh, *xendit.Error) {
	return c.CreateWithContext(context.Background(), data)
}

// CreateWithContext creates new disbursement with context
func (c *Client) CreateWithContext(ctx context.Context, data *CreateParams) (*xendit.DisbursementPh, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.DisbursementPh{}
	header := http.Header{}

	if data.IdempotencyKey != "" {
		header.Add("XENDIT-IDEMPOTENCY-KEY", data.IdempotencyKey)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/disbursements", c.Opt.XenditURL),
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

// GetByID gets a disbursement by id
func (c *Client) GetByID(data *GetByIDParams) (*xendit.DisbursementPh, *xendit.Error) {
	return c.GetByIDWithContext(context.Background(), data)
}

// GetByIDWithContext gets a disbursement by id with context
func (c *Client) GetByIDWithContext(ctx context.Context, data *GetByIDParams) (*xendit.DisbursementPh, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.DisbursementPh{}
	header := http.Header{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/disbursements/%s", c.Opt.XenditURL, data.DisbursementID),
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

// GetByReferenceID gets a disbursement by id
func (c *Client) GetByReferenceID(data *GetByReferenceIDParams) ([]xendit.DisbursementPh, *xendit.Error) {
	return c.GetByReferenceIDWithContext(context.Background(), data)
}

// GetByReferenceIDWithContext gets a disbursement by id with context
func (c *Client) GetByReferenceIDWithContext(ctx context.Context, data *GetByReferenceIDParams) ([]xendit.DisbursementPh, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := []xendit.DisbursementPh{}
	header := http.Header{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/disbursements?%s", c.Opt.XenditURL, data.QueryString()),
		c.Opt.SecretKey,
		header,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetDisbursementChannels gets available disbursement channels
func (c *Client) GetDisbursementChannels() ([]xendit.DisbursementChannel, *xendit.Error) {
	return c.GetDisbursementChannelsWithContext(context.Background())
}

// GetDisbursementChannelsWithContext gets available disbursement channels with context
func (c *Client) GetDisbursementChannelsWithContext(ctx context.Context) ([]xendit.DisbursementChannel, *xendit.Error) {
	response := []xendit.DisbursementChannel{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/disbursement-channels", c.Opt.XenditURL),
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

// GetDisbursementChannelsByChannelCategory gets available disbursement channels by ChannelCategory
func (c *Client) GetDisbursementChannelsByChannelCategory(data *GetByChannelCategoryParams) ([]xendit.DisbursementChannel, *xendit.Error) {
	return c.GetDisbursementChannelsByChannelCategoryWithContext(context.Background(), data)
}

// GetDisbursementChannelsByChannelCategoryWithContext gets available disbursement channels by ChannelCategory with context
func (c *Client) GetDisbursementChannelsByChannelCategoryWithContext(ctx context.Context, data *GetByChannelCategoryParams) ([]xendit.DisbursementChannel, *xendit.Error) {
	response := []xendit.DisbursementChannel{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/disbursement-channels?%s", c.Opt.XenditURL, data.QueryString()),
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

// GetDisbursementChannelsByChannelCode gets available disbursement channels by ChannelCode
func (c *Client) GetDisbursementChannelsByChannelCode(data *GetByChannelCodeParams) ([]xendit.DisbursementChannel, *xendit.Error) {
	return c.GetDisbursementChannelsByChannelCodeWithContext(context.Background(), data)
}

// GetDisbursementChannelsByChannelCodeWithContext gets available disbursement channels by ChannelCode with context
func (c *Client) GetDisbursementChannelsByChannelCodeWithContext(ctx context.Context, data *GetByChannelCodeParams) ([]xendit.DisbursementChannel, *xendit.Error) {
	response := []xendit.DisbursementChannel{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/disbursement-channels?%s", c.Opt.XenditURL, data.QueryString()),
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

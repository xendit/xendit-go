package promotion

import (
	"context"
	"fmt"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke ewallet API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// CreatePromotion creates new promotion
func (c *Client) CreatePromotion(data *CreatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	return c.CreatePromotionWithContext(context.Background(), data)
}

// CreatePromotionWithContext creates new promotion with context
func (c *Client) CreatePromotionWithContext(ctx context.Context, data *CreatePromotionParams) (*xendit.Promotion, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Promotion{}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/promotions", c.Opt.XenditURL),
		c.Opt.SecretKey,
		nil,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetPromotions gets promotions
func (c *Client) GetPromotions(data *GetPromotionsParams) ([]xendit.Promotion, *xendit.Error) {
	return c.GetPromotionsWithContext(context.Background(), data)
}

// GetPromotionsWithContext gets promotions with context
func (c *Client) GetPromotionsWithContext(ctx context.Context, data *GetPromotionsParams) ([]xendit.Promotion, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := []xendit.Promotion{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/promotions?%s", c.Opt.XenditURL, data.QueryString()),
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

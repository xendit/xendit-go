package card

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

// CreateAuthorization creates new card authorization
func (c Client) CreateAuthorization(data *CreateAuthorizationParams) (*xendit.CardCreateAuthorizationResponse, *xendit.Error) {
	return c.CreateAuthorizationWithContext(context.Background(), data)
}

// CreateAuthorizationWithContext creates new card authorization with context
func (c Client) CreateAuthorizationWithContext(ctx context.Context, data *CreateAuthorizationParams) (*xendit.CardCreateAuthorizationResponse, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.CardCreateAuthorizationResponse{}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/credit_card_charges", c.Opt.XenditURL),
		c.Opt.SecretKey,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ReverseAuthorization reverses a card authorization
func (c Client) ReverseAuthorization(data *ReverseAuthorizationParams) (*xendit.CardReverseAuthorizationResponse, *xendit.Error) {
	return c.ReverseAuthorizationWithContext(context.Background(), data)
}

// ReverseAuthorizationWithContext reverses a card authorization with context
func (c Client) ReverseAuthorizationWithContext(ctx context.Context, data *ReverseAuthorizationParams) (*xendit.CardReverseAuthorizationResponse, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.CardReverseAuthorizationResponse{}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/credit_card_charges/%s/auth_reversal", c.Opt.XenditURL, data.ChargeID),
		c.Opt.SecretKey,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

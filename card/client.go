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

// CreateCharge creates new card charge
func (c Client) CreateCharge(data *CreateChargeParams) (*xendit.CardCharge, *xendit.Error) {
	return c.CreateChargeWithContext(context.Background(), data)
}

// CreateChargeWithContext creates new card charge with context
func (c Client) CreateChargeWithContext(ctx context.Context, data *CreateChargeParams) (*xendit.CardCharge, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.CardCharge{}

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

// CaptureCharge captures a card charge
func (c Client) CaptureCharge(data *CaptureChargeParams) (*xendit.CardCharge, *xendit.Error) {
	return c.CaptureChargeWithContext(context.Background(), data)
}

// CaptureChargeWithContext captures a card charge with context
func (c Client) CaptureChargeWithContext(ctx context.Context, data *CaptureChargeParams) (*xendit.CardCharge, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.CardCharge{}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/credit_card_charges/%s/capture", c.Opt.XenditURL, data.ChargeID),
		c.Opt.SecretKey,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetCharge gets a card charge
func (c Client) GetCharge(data *CaptureChargeParams) (*xendit.CardCharge, *xendit.Error) {
	return c.CaptureChargeWithContext(context.Background(), data)
}

// GetChargeWithContext gets a card charge with context
func (c Client) GetChargeWithContext(ctx context.Context, data *GetChargeParams) (*xendit.CardCharge, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.CardCharge{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/credit_card_charges/%s", c.Opt.XenditURL, data.ChargeID),
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

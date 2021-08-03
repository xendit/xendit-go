package directdebitpayment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke direct debit (linked account) API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

type DirectDebitPaymentResponse struct {
	ID                     string                         `json:"id"`
	ReferenceID            string                         `json:"reference_id"`
	ChannelCode            xendit.ChannelCodeEnum         `json:"channel_code"`
	PaymentMethodID        string                         `json:"payment_method_id"`
	Currency               string                         `json:"currency"`
	Amount                 float64                        `json:"amount"`
	Description            string                         `json:"description"`
	Status                 string                         `json:"status"`
	FailureCode            string                         `json:"failure_code"`
	IsOTPRequired          bool                           `json:"is_otp_required"`
	OTPMobileNumber        string                         `json:"otp_mobile_number"`
	OTPExpirationTimestamp string                         `json:"otp_expiration_timestamp"`
	Created                string                         `json:"created"`
	Updated                string                         `json:"updated"`
	Basket                 []xendit.DirectDebitBasketItem `json:"basket"`
	Metadata               map[string]interface{}         `json:"metadata"`
	Device                 xendit.DirectDebitDevice       `json:"device"`
	RefundedAmount         float64                        `json:"refunded_amount"`
	Refunds                xendit.DirectDebitRefunds      `json:"refunds"`
	SuccessRedirectURL     string                         `json:"success_redirect_url"`
	CheckoutURL            string                         `json:"checkout_url"`
	FailureRedirectURL     string                         `json:"failure_redirect_url"`
	RequiredAction         string                         `json:"required_action"`
}

func (r *DirectDebitPaymentResponse) toDirectDebitPaymentResponse() *xendit.DirectDebitPayment {
	return &xendit.DirectDebitPayment{
		ID:                     r.ID,
		ReferenceID:            r.ReferenceID,
		ChannelCode:            r.ChannelCode,
		PaymentMethodID:        r.PaymentMethodID,
		Currency:               r.Currency,
		Amount:                 r.Amount,
		Description:            r.Description,
		Status:                 r.Status,
		FailureCode:            r.FailureCode,
		IsOTPRequired:          r.IsOTPRequired,
		OTPMobileNumber:        r.OTPMobileNumber,
		OTPExpirationTimestamp: r.OTPExpirationTimestamp,
		Created:                r.Created,
		Updated:                r.Updated,
		Basket:                 r.Basket,
		Metadata:               r.Metadata,
		Device:                 r.Device,
		RefundedAmount:         r.RefundedAmount,
		Refunds:                r.Refunds,
		SuccessRedirectURL:     r.SuccessRedirectURL,
		CheckoutURL:            r.CheckoutURL,
		FailureRedirectURL:     r.FailureRedirectURL,
		RequiredAction:         r.RequiredAction,
	}
}

// CreateDirectDebitPayment created new direct debit payment
func (c *Client) CreateDirectDebitPayment(data *CreateDirectDebitPaymentParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	return c.CreateDirectDebitPaymentWithContext(context.Background(), data)
}

// CreateDirectDebitPaymentWithContext created new direct debit payment
func (c *Client) CreateDirectDebitPaymentWithContext(ctx context.Context, data *CreateDirectDebitPaymentParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.DirectDebitPayment{}
	header := &http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	header.Add("Idempotency-key", data.IdempotencyKey)

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/direct_debits", c.Opt.XenditURL),
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

// ValidateOTPForDirectDebitPayment validate OTP for direct debit payment
func (c *Client) ValidateOTPForDirectDebitPayment(data *ValidateOTPForDirectDebitPaymentParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	return c.ValidateOTPForDirectDebitPaymentWithContext(context.Background(), data)
}

// ValidateOTPForDirectDebitPayment validate OTP for direct debit payment
func (c *Client) ValidateOTPForDirectDebitPaymentWithContext(ctx context.Context, data *ValidateOTPForDirectDebitPaymentParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.DirectDebitPayment{}
	header := &http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/direct_debits/%s/validate_otp/", c.Opt.XenditURL, data.DirectDebitID),
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

// GetDirectDebitPaymentStatusByID gets direct debit payment status by ID
func (c *Client) GetDirectDebitPaymentStatusByID(data *GetDirectDebitPaymentStatusByIDParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	return c.GetDirectDebitPaymentStatusByIDWithContext(context.Background(), data)
}

// GetDirectDebitPaymentStatusByIDWithContext gets direct debit payment status by ID
func (c *Client) GetDirectDebitPaymentStatusByIDWithContext(ctx context.Context, data *GetDirectDebitPaymentStatusByIDParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}
	tempResponse := &DirectDebitPaymentResponse{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/direct_debits/%s/", c.Opt.XenditURL, data.ID),
		c.Opt.SecretKey,
		nil,
		nil,
		tempResponse,
	)
	if err != nil {
		return nil, err
	}

	response := tempResponse.toDirectDebitPaymentResponse()

	return response, nil
}

// GetDirectDebitPaymentStatusByReferenceID gets direct debit payment status by reference ID
func (c *Client) GetDirectDebitPaymentStatusByReferenceID(data *GetDirectDebitPaymentStatusByReferenceIDParams) ([]xendit.DirectDebitPayment, *xendit.Error) {
	return c.GetDirectDebitPaymentStatusByReferenceIDWithContext(context.Background(), data)
}

// GetDirectDebitPaymentStatusByReferenceIDWithContext gets direct debit payment status by reference ID
func (c *Client) GetDirectDebitPaymentStatusByReferenceIDWithContext(ctx context.Context, data *GetDirectDebitPaymentStatusByReferenceIDParams) ([]xendit.DirectDebitPayment, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}
	response := []xendit.DirectDebitPayment{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/direct_debits?%s", c.Opt.XenditURL, data.QueryString()),
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

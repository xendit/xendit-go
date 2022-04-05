package ewallet

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
)

// Client is the client used to invoke e-wallet API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// getPaymentStatusResponse is e-wallet data that is contained in API response of Get Payment Status.
// It exists because the type of `Amount` in Get Payment Status json response is string,
// different from the CreatePayment
type getPaymentStatusResponse struct {
	EWalletType     xendit.EWalletTypeEnum `json:"ewallet_type"`
	ExternalID      string                 `json:"external_id"`
	Amount          float64                `json:"amount"`
	TransactionDate *time.Time             `json:"transaction_date,omitempty"`
	CheckoutURL     string                 `json:"checkout_url,omitempty"`
	BusinessID      string                 `json:"business_id,omitempty"`
}

type EWalletChargeResponse struct {
	ID                 string                     `json:"id"`
	BusinessID         string                     `json:"business_id"`
	ReferenceID        string                     `json:"reference_id"`
	Status             string                     `json:"status"`
	Currency           string                     `json:"currency"`
	ChargeAmount       float64                    `json:"charge_amount"`
	CaptureAmount      float64                    `json:"capture_amount"`
	CheckoutMethod     string                     `json:"checkout_method"`
	ChannelCode        string                     `json:"channel_code"`
	ChannelProperties  map[string]string          `json:"channel_properties"`
	Actions            map[string]string          `json:"actions"`
	IsRedirectRequired bool                       `json:"is_redirect_required"`
	CallbackURL        string                     `json:"callback_url"`
	Created            string                     `json:"created"`
	Updated            string                     `json:"updated"`
	VoidedAt           string                     `json:"voided_at"`
	CaptureNow         bool                       `json:"capture_now"`
	CustomerID         string                     `json:"customer_id"`
	PaymentMethodID    string                     `json:"payment_method_id"`
	FailureCode        string                     `json:"failure_code"`
	Basket             []xendit.EWalletBasketItem `json:"basket"`
	Metadata           map[string]interface{}     `json:"metadata"`
}

func (r *getPaymentStatusResponse) toEwalletResponse() *xendit.EWallet {
	return &xendit.EWallet{
		EWalletType:     r.EWalletType,
		ExternalID:      r.ExternalID,
		Amount:          r.Amount,
		TransactionDate: r.TransactionDate,
		CheckoutURL:     r.CheckoutURL,
		BusinessID:      r.BusinessID,
	}
}

func (r *EWalletChargeResponse) toEWalletChargeResponse() *xendit.EWalletCharge {
	return &xendit.EWalletCharge{
		ID:                 r.ID,
		BusinessID:         r.BusinessID,
		ReferenceID:        r.ReferenceID,
		Status:             r.Status,
		Currency:           r.Currency,
		ChargeAmount:       r.ChargeAmount,
		CaptureAmount:      r.CaptureAmount,
		CheckoutMethod:     r.CheckoutMethod,
		ChannelCode:        r.ChannelCode,
		ChannelProperties:  r.ChannelProperties,
		Actions:            r.Actions,
		IsRedirectRequired: r.IsRedirectRequired,
		CallbackURL:        r.CallbackURL,
		Created:            r.Created,
		Updated:            r.Updated,
		VoidedAt:           r.VoidedAt,
		CaptureNow:         r.CaptureNow,
		CustomerID:         r.CustomerID,
		PaymentMethodID:    r.PaymentMethodID,
		FailureCode:        r.FailureCode,
		Basket:             r.Basket,
		Metadata:           r.Metadata,
	}
}

// CreatePayment creates new payment
func (c *Client) CreatePayment(data *CreatePaymentParams) (*xendit.EWallet, *xendit.Error) {
	return c.CreatePaymentWithContext(context.Background(), data)
}

// CreatePaymentWithContext creates new payment
func (c *Client) CreatePaymentWithContext(ctx context.Context, data *CreatePaymentParams) (*xendit.EWallet, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.EWallet{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.XApiVersion != "" {
		header.Add("X-API-VERSION", data.XApiVersion)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/ewallets", c.Opt.XenditURL),
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

// GetPaymentStatus gets one payment with its status
func (c *Client) GetPaymentStatus(data *GetPaymentStatusParams) (*xendit.EWallet, *xendit.Error) {
	return c.GetPaymentStatusWithContext(context.Background(), data)
}

// GetPaymentStatusWithContext gets one payment with its status
func (c *Client) GetPaymentStatusWithContext(ctx context.Context, data *GetPaymentStatusParams) (*xendit.EWallet, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	tempResponse := &getPaymentStatusResponse{}
	var queryString string

	if data != nil {
		queryString = data.QueryString()
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/ewallets?%s", c.Opt.XenditURL, queryString),
		c.Opt.SecretKey,
		nil,
		nil,
		tempResponse,
	)
	if err != nil {
		return nil, err
	}

	response := tempResponse.toEwalletResponse()

	return response, nil
}

// CreateEWalletCharge creates new e-wallet charge
func (c *Client) CreateEWalletCharge(data *CreateEWalletChargeParams) (*xendit.EWalletCharge, *xendit.Error) {
	return c.CreateEWalletChargeWithContext(context.Background(), data)
}

// CreateEWalletChargeWithContext creates new e-wallet charge
func (c *Client) CreateEWalletChargeWithContext(ctx context.Context, data *CreateEWalletChargeParams) (*xendit.EWalletCharge, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.EWalletCharge{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.WithFeeRule != "" {
		header.Add("with-fee-rule", data.WithFeeRule)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/ewallets/charges", c.Opt.XenditURL),
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

// GetEWalletChargeStatus gets one e-wallet charge with its status
func (c *Client) GetEWalletChargeStatus(data *GetEWalletChargeStatusParams) (*xendit.EWalletCharge, *xendit.Error) {
	return c.GetEWalletChargeStatusWithContext(context.Background(), data)
}

// GetPaymentStatusWithContext gets one payment with its status
func (c *Client) GetEWalletChargeStatusWithContext(ctx context.Context, data *GetEWalletChargeStatusParams) (*xendit.EWalletCharge, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	tempResponse := &EWalletChargeResponse{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/ewallets/charges/%s", c.Opt.XenditURL, data.ChargeID),
		c.Opt.SecretKey,
		header,
		nil,
		tempResponse,
	)
	if err != nil {
		return nil, err
	}

	response := tempResponse.toEWalletChargeResponse()

	return response, nil
}

package pmsv2

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/xendit/xendit-go"
)

type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

func (c *Client) CreatePaymentMethod(data *CreatePaymentMethodParams) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return c.CreatePaymentMethodWithContext(context.Background(), data)
}

func (c *Client) CreatePaymentMethodWithContext(ctx context.Context, data *CreatePaymentMethodParams) (*xendit.PaymentMethodResponse, *xendit.Error) {
	response := &xendit.PaymentMethodResponse{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	if data.IdempotencyKey != "" {
		header.Add("idempotency-key", data.IdempotencyKey)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/v2/payment_method", c.Opt.XenditURL),
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

func (c *Client) AuthPaymentMethod(data *ValidateOTPRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return c.AuthPaymentMethodWithContext(context.Background(), data)
}

func (c *Client) AuthPaymentMethodWithContext(ctx context.Context, data *ValidateOTPRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	response := &xendit.PaymentMethodResponse{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	if data.IdempotencyKey != "" {
		header.Add("idempotency-key", data.IdempotencyKey)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/v2/payment_method/%s/auth", c.Opt.XenditURL, data.ID),
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

func (c *Client) ExpirePaymentMethod(data *ExpireRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return c.ExpirePaymentMethodWithContext(context.Background(), data)
}

func (c *Client) ExpirePaymentMethodWithContext(ctx context.Context, data *ExpireRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	response := &xendit.PaymentMethodResponse{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	if data.IdempotencyKey != "" {
		header.Add("idempotency-key", data.IdempotencyKey)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/v2/payment_method/%s/expire", c.Opt.XenditURL, data.ID),
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

func (c *Client) RetrievePaymentMethod(data *RetrievePaymentMethodRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return c.RetrievePaymentMethodWithContext(context.Background(), data)
}

func (c *Client) RetrievePaymentMethodWithContext(ctx context.Context, data *RetrievePaymentMethodRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {

	response := &xendit.PaymentMethodResponse{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/payment_method/%s", c.Opt.XenditURL, data.ID),
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

func (c *Client) RetrieveAllPaymentMethods(data *RetrieveAllPaymentMethodsRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return c.RetrieveAllPaymentMethodsWithContext(context.Background(), data)
}

func (c *Client) RetrieveAllPaymentMethodsWithContext(ctx context.Context, data *RetrieveAllPaymentMethodsRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {

	response := &xendit.PaymentMethodResponse{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	urlValues := &url.Values{}
	for _, item := range data.ID {
		urlValues.Add("id", string(item))
	}
	for _, item := range data.Status {
		urlValues.Add("status", string(item))
	}
	addUrlValue(urlValues, "type", string(data.Type))
	addUrlValue(urlValues, "reusability", string(data.Reusability))
	addUrlValue(urlValues, "customer_id", data.CustomerID)
	addUrlValue(urlValues, "business_id", data.BusinessID)
	addUrlValue(urlValues, "after_id", data.AfterID)
	addUrlValue(urlValues, "before_id", data.BeforeID)
	if data.Limit != 0 {
		addUrlValue(urlValues, "limit", strconv.Itoa(data.Limit))
	}
	queryString := urlValues.Encode()

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/payment_method?%s", c.Opt.XenditURL, queryString),
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

func (c *Client) UpdatePaymentMethod(data *UpdateRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return c.UpdatePaymentMethodWithContext(context.Background(), data)
}

func (c *Client) UpdatePaymentMethodWithContext(ctx context.Context, data *UpdateRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	response := &xendit.PaymentMethodResponse{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	if data.IdempotencyKey != "" {
		header.Add("idempotency-key", data.IdempotencyKey)
	}

	err := c.APIRequester.Call(
		ctx,
		"PATCH",
		fmt.Sprintf("%s/v2/payment_method/%s", c.Opt.XenditURL, data.ID),
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

func (c *Client) RetrievePayments(data *RetrievePaymentsRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {
	return c.RetrievePaymentsWithContext(context.Background(), data)
}

func (c *Client) RetrievePaymentsWithContext(ctx context.Context, data *RetrievePaymentsRequest) (*xendit.PaymentMethodResponse, *xendit.Error) {

	response := &xendit.PaymentMethodResponse{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	urlValues := &url.Values{}
	if !data.CreatedGTE.IsZero() {
		addUrlValue(urlValues, "created[gte]", data.CreatedGTE.Format(time.RFC3339))
	}
	if !data.CreatedLTE.IsZero() {
		addUrlValue(urlValues, "created[lte]", data.CreatedLTE.Format(time.RFC3339))
	}
	if !data.UpdatedGTE.IsZero() {
		addUrlValue(urlValues, "created[lte]", data.UpdatedGTE.Format(time.RFC3339))
	}
	if !data.UpdatedLTE.IsZero() {
		addUrlValue(urlValues, "created[lte]", data.UpdatedLTE.Format(time.RFC3339))
	}

	for _, item := range data.PaymentMethodType {
		urlValues.Add("payment_method_type", string(item))
	}
	for _, item := range data.PaymentMethodID {
		urlValues.Add("payment_method_id", string(item))
	}
	for _, item := range data.PaymentRequestID {
		urlValues.Add("payment_request_id", string(item))
	}
	for _, item := range data.ChannelCode {
		urlValues.Add("channel_code", string(item))
	}
	for _, item := range data.Status {
		urlValues.Add("status", string(item))
	}
	for _, item := range data.Currency {
		urlValues.Add("currency", string(item))
	}
	for _, item := range data.ReferenceID {
		urlValues.Add("reference_id", string(item))
	}
	addUrlValue(urlValues, "after_id", data.AfterID)
	addUrlValue(urlValues, "before_id", data.BeforeID)
	if data.Limit != 0 {
		addUrlValue(urlValues, "limit", strconv.Itoa(data.Limit))
	}

	queryString := urlValues.Encode()

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/payment_method/%s/payments?%s", c.Opt.XenditURL, data.ID, queryString),
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

func addUrlValue(urlValues *url.Values, key string, data string) {
	if data != "" {
		urlValues.Add(key, string(data))
	}
}

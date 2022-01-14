package linkedaccount

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

// InitializeLinkedAccountTokenization initialize new linked account tokenization
func (c *Client) InitializeLinkedAccountTokenization(data *InitializeLinkedAccountTokenizationParams) (*xendit.InitializedLinkedAccount, *xendit.Error) {
	return c.InitializeLinkedAccountTokenizationWithContext(context.Background(), data)
}

// InitializeLinkedAccountTokenization initialize new linked account tokenization
func (c *Client) InitializeLinkedAccountTokenizationWithContext(ctx context.Context, data *InitializeLinkedAccountTokenizationParams) (*xendit.InitializedLinkedAccount, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.InitializedLinkedAccount{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/linked_account_tokens/auth", c.Opt.XenditURL),
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

// ValidateOTPForLinkedAccount validate OTP for linked account
func (c *Client) ValidateOTPForLinkedAccount(data *ValidateOTPForLinkedAccountParams) (*xendit.ValidatedLinkedAccount, *xendit.Error) {
	return c.ValidateOTPForLinkedAccountWithContext(context.Background(), data)
}

// ValidateOTPForLinkedAccount validate OTP for linked account
func (c *Client) ValidateOTPForLinkedAccountWithContext(ctx context.Context, data *ValidateOTPForLinkedAccountParams) (*xendit.ValidatedLinkedAccount, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.ValidatedLinkedAccount{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/linked_account_tokens/%s/validate_otp", c.Opt.XenditURL, data.LinkedAccountTokenID),
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

// RetrieveAccessibleLinkedAccounts gets accessible linked accounts
func (c *Client) RetrieveAccessibleLinkedAccounts(data *RetrieveAccessibleLinkedAccountParams) ([]xendit.AccessibleLinkedAccount, *xendit.Error) {
	return c.RetrieveAccessibleLinkedAccountsWithContext(context.Background(), data)
}

// RetrieveAccessibleLinkedAccountsWithContext gets accessible linked accounts
func (c *Client) RetrieveAccessibleLinkedAccountsWithContext(ctx context.Context, data *RetrieveAccessibleLinkedAccountParams) ([]xendit.AccessibleLinkedAccount, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := []xendit.AccessibleLinkedAccount{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/linked_account_tokens/%s/accounts", c.Opt.XenditURL, data.LinkedAccountTokenID),
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

// UnbindLinkedAccountToken unbind a successful linked account token
func (c *Client) UnbindLinkedAccountToken(data *UnbindLinkedAccountTokenParams) (*xendit.UnbindedLinkedAccount, *xendit.Error) {
	return c.UnbindLinkedAccountTokenWithContext(context.Background(), data)
}

// UnbindLinkedAccountTokenWithContext unbind a successful linked account token
func (c *Client) UnbindLinkedAccountTokenWithContext(ctx context.Context, data *UnbindLinkedAccountTokenParams) (*xendit.UnbindedLinkedAccount, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.UnbindedLinkedAccount{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"DELETE",
		fmt.Sprintf("%s/linked_account_tokens/%s", c.Opt.XenditURL, data.LinkedAccountTokenID),
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

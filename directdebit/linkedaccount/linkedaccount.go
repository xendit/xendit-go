package linkedaccount

import (
	"context"

	"github.com/xendit/xendit-go"
)

// InitializeLinkedAccountTokenization initialize new linked account tokenization
func InitializeLinkedAccountTokenization(data *InitializeLinkedAccountTokenizationParams) (*xendit.InitializedLinkedAccount, *xendit.Error) {
	return InitializeLinkedAccountTokenizationWithContext(context.Background(), data)
}

// InitializeLinkedAccountTokenizationWithContext initialize new linked account tokenization
func InitializeLinkedAccountTokenizationWithContext(ctx context.Context, data *InitializeLinkedAccountTokenizationParams) (*xendit.InitializedLinkedAccount, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.InitializeLinkedAccountTokenizationWithContext(ctx, data)
}

// ValidateOTPForLinkedAccount validate OTP for linked account
func ValidateOTPForLinkedAccount(data *ValidateOTPForLinkedAccountParams) (*xendit.ValidatedLinkedAccount, *xendit.Error) {
	return ValidateOTPForLinkedAccountWithContext(context.Background(), data)
}

// ValidateOTPForLinkedAccountWithContext validate OTP for linked account
func ValidateOTPForLinkedAccountWithContext(ctx context.Context, data *ValidateOTPForLinkedAccountParams) (*xendit.ValidatedLinkedAccount, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ValidateOTPForLinkedAccountWithContext(ctx, data)
}

// RetrieveAccessibleLinkedAccounts gets accessible linked accounts
func RetrieveAccessibleLinkedAccounts(data *RetrieveAccessibleLinkedAccountParams) ([]xendit.AccessibleLinkedAccount, *xendit.Error) {
	return RetrieveAccessibleLinkedAccountsWithContext(context.Background(), data)
}

// RetrieveAccessibleLinkedAccountsWithContext gets accessible linked accounts
func RetrieveAccessibleLinkedAccountsWithContext(ctx context.Context, data *RetrieveAccessibleLinkedAccountParams) ([]xendit.AccessibleLinkedAccount, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.RetrieveAccessibleLinkedAccountsWithContext(ctx, data)
}

// UnbindLinkedAccountToken unbind a successful linked account token
func UnbindLinkedAccountToken(data *UnbindLinkedAccountTokenParams) (*xendit.UnbindedLinkedAccount, *xendit.Error) {
	return UnbindLinkedAccountTokenWithContext(context.Background(), data)
}

// UnbindLinkedAccountTokenWithContext unbind a successful linked account token
func UnbindLinkedAccountTokenWithContext(ctx context.Context, data *UnbindLinkedAccountTokenParams) (*xendit.UnbindedLinkedAccount, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.UnbindLinkedAccountTokenWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}

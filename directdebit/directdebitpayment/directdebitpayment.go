package directdebitpayment

import (
	"context"

	"github.com/xendit/xendit-go"
)

// CreateDirectDebitPayment created new direct debit payment
func CreateDirectDebitPayment(data *CreateDirectDebitPaymentParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	return CreateDirectDebitPaymentWithContext(context.Background(), data)
}

// CreateDirectDebitPaymentWithContext created new direct debit payment
func CreateDirectDebitPaymentWithContext(ctx context.Context, data *CreateDirectDebitPaymentParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateDirectDebitPaymentWithContext(ctx, data)
}

// ValidateOTPForDirectDebitPayment validate OTP for direct debit payment
func ValidateOTPForDirectDebitPayment(data *ValidateOTPForDirectDebitPaymentParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	return ValidateOTPForDirectDebitPaymentWithContext(context.Background(), data)
}

// ValidateOTPForDirectDebitPaymentWithContext validate OTP for direct debit payment
func ValidateOTPForDirectDebitPaymentWithContext(ctx context.Context, data *ValidateOTPForDirectDebitPaymentParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ValidateOTPForDirectDebitPaymentWithContext(ctx, data)
}

// GetDirectDebitPaymentStatusByID gets direct debit payment status by ID
func GetDirectDebitPaymentStatusByID(data *GetDirectDebitPaymentStatusByIDParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	return GetDirectDebitPaymentStatusByIDWithContext(context.Background(), data)
}

// GetDirectDebitPaymentStatusByIDWithContext gets direct debit payment status by ID
func GetDirectDebitPaymentStatusByIDWithContext(ctx context.Context, data *GetDirectDebitPaymentStatusByIDParams) (*xendit.DirectDebitPayment, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetDirectDebitPaymentStatusByIDWithContext(ctx, data)
}

// GetDirectDebitPaymentStatusByReferenceID gets direct debit payment status by reference ID
func GetDirectDebitPaymentStatusByReferenceID(data *GetDirectDebitPaymentStatusByReferenceIDParams) ([]xendit.DirectDebitPayment, *xendit.Error) {
	return GetDirectDebitPaymentStatusByReferenceIDWithContext(context.Background(), data)
}

// GetDirectDebitPaymentStatusByReferenceIDWithContext gets direct debit payment status by reference ID
func GetDirectDebitPaymentStatusByReferenceIDWithContext(ctx context.Context, data *GetDirectDebitPaymentStatusByReferenceIDParams) ([]xendit.DirectDebitPayment, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetDirectDebitPaymentStatusByReferenceIDWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}

package linkedaccount

import (
	"github.com/xendit/xendit-go"
)

// InitializeLinkedAccountTokenizationParams contains parameters for InitializeLinkedAccountTokenization
type InitializeLinkedAccountTokenizationParams struct {
	ForUserID   string                 `json:"-"`
	CustomerID  string                 `json:"customer_id" validate:"required"`
	ChannelCode xendit.ChannelCodeEnum `json:"channel_code" validate:"required"`
	Properties  map[string]interface{} `json:"properties,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// ValidateOTPParams contains parameters for ValidateOTPForLinkedAccount
type ValidateOTPForLinkedAccountParams struct {
	ForUserID            string `json:"-"`
	LinkedAccountTokenID string `json:"linked_account_token_id" validate:"required"`
	OTPCode              string `json:"otp_code" validate:"required"`
}

type RetrieveAccessibleLinkedAccountParams struct {
	ForUserID            string `json:"-"`
	LinkedAccountTokenID string `json:"linked_account_token_id" validate:"required"`
}

type UnbindLinkedAccountTokenParams struct {
	ForUserID            string `json:"-"`
	LinkedAccountTokenID string `json:"linked_account_token_id" validate:"required"`
}

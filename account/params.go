package account

import "github.com/xendit/xendit-go"

type CreateParams struct {
	Email         string               `json:"email" validate:"required"`
	Type          xendit.AccountType   `json:"type" validate:"required"`
	PublicProfile xendit.PublicProfile `json:"public_profile"`
}

type GetParams struct {
	ID string `json:"id" validate:"required"`
}

package directdebit

type ChannelCode string
type ChannelProperties map[string]interface{}
type Type string

const (
	// debit card
	BRI        ChannelCode = "BRI"
	BCAOneklik ChannelCode = "BCA_ONEKLIK"
	CIMBNIAGA  ChannelCode = "CIMBNIAGA"

	// bank account
	BPI       ChannelCode = "BPI"
	Unionbank ChannelCode = "UBP"
	RCBC      ChannelCode = "RCBC"
	BDO       ChannelCode = "BDO"
	MTB       ChannelCode = "MTB"
	CHINABANK ChannelCode = "CHINABANK"
	Mandiri   ChannelCode = "MANDIRI"

	BBL ChannelCode = "BBL"
	SCB ChannelCode = "SCB"
	KTB ChannelCode = "KTB"
	BAY ChannelCode = "BAY"

	// bank redirect
	BCAKlikpay ChannelCode = "BCA_KLIKPAY"
)

const (
	BankAccountType  Type = "BANK_ACCOUNT"
	DebitCardType    Type = "DEBIT_CARD"
	BankRedirectType Type = "BANK_REDIRECT"
)

type CreateMethod struct {
	LinkedAccountTokenID string            `json:"linked_account_token_id"`
	LinkedAccountID      string            `json:"linked_account_id"`
	ChannelCode          ChannelCode       `json:"channel_code"`
	ChannelProperties    ChannelProperties `json:"channel_properties"`
	Type                 Type              `json:"type"`
	AccessToken          *string           `json:"access_token"`
}

type BankAccount struct {
	AccountNumber *string `json:"masked_bank_account_number"`
	AccountHash   *string `json:"bank_account_hash"`
}

type DebitCard struct {
	MobileNumber *string `json:"mobile_number"`
	CardLastFour *string `json:"card_last_four"`
	CardExpiry   *string `json:"card_expiry"`
	Email        *string `json:"email"`
}

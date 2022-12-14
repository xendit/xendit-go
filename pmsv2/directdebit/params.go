package directdebit

type ChannelCode string
type ChannelProperties map[string]interface{}
type Type string

const (
	// debit card
	BRI       ChannelCode = "BRI"
	CIMBNIAGA ChannelCode = "CIMBNIAGA"

	// bank account
	BPI       ChannelCode = "BPI"
	Unionbank ChannelCode = "UBP"
	RCBC      ChannelCode = "RCBC"
	CHINABANK ChannelCode = "CHINABANK"
	Mandiri   ChannelCode = "MANDIRI"

	BBL ChannelCode = "BBL"
	SCB ChannelCode = "SCB"
	KTB ChannelCode = "KTB"
	BAY ChannelCode = "BAY"
)

const (
	BankAccountType  Type = "BANK_ACCOUNT"
	DebitCardType    Type = "DEBIT_CARD"
	BankRedirectType Type = "BANK_REDIRECT"
)

type CreateMethod struct {
	ChannelCode       ChannelCode       `json:"channel_code"`
	ChannelProperties ChannelProperties `json:"channel_properties"`
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

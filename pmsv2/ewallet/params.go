package ewallet

type ChannelCode string

type ChannelProperties map[string]interface{}

const (
	Grabpay   ChannelCode = "GRABPAY"
	Paymaya   ChannelCode = "PAYMAYA"
	Gcash     ChannelCode = "GCASH"
	OVO       ChannelCode = "OVO"
	DANA      ChannelCode = "DANA"
	LinkAja   ChannelCode = "LINKAJA"
	Shopeepay ChannelCode = "SHOPEEPAY"
	Sakuku    ChannelCode = "SAKUKU"
	Nexcash   ChannelCode = "NEXCASH"
	Astrapay  ChannelCode = "ASTRAPAY"
	Jeniuspay ChannelCode = "JENIUSPAY"
)

type CreateMethod struct {
	LinkedAccountTokenID string            `json:"linked_account_token_id"`
	LinkedAccountID      string            `json:"linked_account_id"`
	ChannelCode          ChannelCode       `json:"channel_code"`
	ChannelProperties    ChannelProperties `json:"channel_properties"`
	AccessToken          *string           `json:"access_token"`
}

type Account struct {
	Name           *string `json:"name"`
	AccountDetails *string `json:"account_details"`
	Balance        float64 `json:"balance"`
	PointBalance   float64 `json:"point_balance"`
}

package xendit

// ChannelCodeEnum constants are the available channel code
type ChannelCodeEnum string

// AccountType constants are the available account type
type AccountTypeEnum string

// This consists the values that ChannelCodeEnum can take
const (
	DC_BRI ChannelCodeEnum = "DC_BRI"
	BA_BPI ChannelCodeEnum = "BA_BPI"
	BA_UBP ChannelCodeEnum = "BA_UBP"
)

// This consists the values that AccountTypeEnum can take
const (
	DEBIT_CARD   AccountTypeEnum = "DEBIT_CARD"
	BANK_ACCOUNT AccountTypeEnum = "BANK_ACCOUNT"
)

type InitializedLinkedAccount struct {
	ID            string                 `json:"id"`
	CustomerID    string                 `json:"customer_id"`
	ChannelCode   ChannelCodeEnum        `json:"channel_code"`
	AuthorizerURL string                 `json:"authorizer_url,omitempty"`
	Status        string                 `json:"status,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

type ValidatedLinkedAccount struct {
	ID          string          `json:"id"`
	CustomerID  string          `json:"customer_id"`
	ChannelCode ChannelCodeEnum `json:"channel_code"`
	Status      string          `json:"status"`
}

type AccessibleLinkedAccount struct {
	ID          string                 `json:"id"`
	ChannelCode ChannelCodeEnum        `json:"channel_code"`
	AccountType AccountTypeEnum        `json:"type"`
	Properties  map[string]interface{} `json:"properties"`
}

type UnbindedLinkedAccount struct {
	ID        string `json:"id"`
	IsDeleted bool   `json:"is_deleted"`
}

type PaymentMethod struct {
	ID         string                 `json:"id"`
	Type       AccountTypeEnum        `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	CustomerID string                 `json:"customer_id"`
	Status     string                 `json:"status"`
	Created    string                 `json:"created"`
	Updated    string                 `json:"updated"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

type DirectDebitBasketItem struct {
	ReferenceID string                 `json:"reference_id"`
	Name        string                 `json:"name"`
	Market      string                 `json:"market"`
	Type        string                 `json:"type"`
	Description string                 `json:"description,omitempty"`
	Category    string                 `json:"category,omitempty"`
	SubCategory string                 `json:"sub_category,omitempty"`
	Price       float64                `json:"price,omitempty"`
	URL         string                 `json:"url,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Quantity    int                    `json:"quantity,omitempty"`
}

type DirectDebitDevice struct {
	ID        string `json:"id"`
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	ADID      string `json:"ad_id,omitempty"`
	Imei      string `json:"imei,omitempty"`
}

type DirectDebitRefunds struct {
	Data    []string `json:"data"`
	HasMore bool     `json:"has_more"`
	URL     string   `json:"url"`
}

type DirectDebitPayment struct {
	ID                     string                  `json:"id"`
	ReferenceID            string                  `json:"reference_id"`
	ChannelCode            ChannelCodeEnum         `json:"channel_code"`
	PaymentMethodID        string                  `json:"payment_method_id"`
	Currency               string                  `json:"currency"`
	Amount                 float64                 `json:"amount"`
	Description            string                  `json:"description"`
	Status                 string                  `json:"status"`
	FailureCode            string                  `json:"failure_code"`
	IsOTPRequired          bool                    `json:"is_otp_required"`
	OTPMobileNumber        string                  `json:"otp_mobile_number"`
	OTPExpirationTimestamp string                  `json:"otp_expiration_timestamp"`
	Created                string                  `json:"created"`
	Updated                string                  `json:"updated"`
	Basket                 []DirectDebitBasketItem `json:"basket"`
	Metadata               map[string]interface{}  `json:"metadata"`
	Device                 DirectDebitDevice       `json:"device"`
	RefundedAmount         float64                 `json:"refunded_amount"`
	Refunds                DirectDebitRefunds      `json:"refunds"`
	SuccessRedirectURL     string                  `json:"success_redirect_url"`
	CheckoutURL            string                  `json:"checkout_url"`
	FailureRedirectURL     string                  `json:"failure_redirect_url"`
	RequiredAction         string                  `json:"required_action"`
}

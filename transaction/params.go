package transaction

// GetTransactionParams contains parameters for GetTransaction
type GetTransactionParams struct {
	ForUserID     string `json:"-"`
	TransactionID string `json:"transaction_id" validate:"required"`
}

// GetListTransactionParams contains parameters for GetListTransaction
type GetListTransactionParams struct {
	ForUserID         string
	Types             []string `url:"types"`
	Statuses          []string `url:"statuses"`
	ChannelCategories []string `url:"channel_categories"`
	ReferenceID       string   `url:"reference_id"`
	ProductID         string   `url:"product_id"`
	AccountIdentifier string   `url:"account_identifier"`
	Currency          string   `url:"currency"`
	Amount            float64  `url:"Amount"`
	CreatedGte        string   `url:"created[gte]"`
	CreatedLte        string   `url:"created[lte]"`
	UpdatedGte        string   `url:"updated[gte]"`
	UpdatedLte        string   `url:"updated[lte]"`
	Limit             int      `url:"limit"`
	AfterID           string   `url:"after_id"`
	BeforeID          string   `url:"before_id"`
}

package xendit

// CustomerAddress contains data from Xendit's API response of customer's Customer Addres requests.
// For more details see https://xendit.github.io/apireference/?bash#customers.
// For documentation of subpackage customer, checkout https://pkg.go.dev/github.com/xendit/xendit-go/customer/
type CustomerAddress struct {
	Country     string `json:"country" validate:"required"`
	StreetLine1 string `json:"street_line1,omitempty"`
	StreetLine2 string `json:"street_line2,omitempty"`
	City        string `json:"city,omitempty"`
	Province    string `json:"province,omitempty"`
	State       string `json:"state,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	Category    string `json:"category,omitempty"`
	IsPreferred bool   `json:"is_preferred,omitempty"`
}

// Customer contains data from Xendit's API response of customer related requests.
// For more details see https://xendit.github.io/apireference/?bash#customers.
// For documentation of subpackage customer, checkout https://pkg.go.dev/github.com/xendit/xendit-go/customer/
type Customer struct {
	ID           string                 `json:"id"`
	ReferenceID  string                 `json:"reference_id"`
	MobileNumber string                 `json:"mobile_number,omitempty"`
	Email        string                 `json:"email,omitempty"`
	GivenNames   string                 `json:"given_names"`
	MiddleName   string                 `json:"middle_name"`
	Surname      string                 `json:"surname"`
	Description  string                 `json:"description,omitempty"`
	PhoneNumber  string                 `json:"phone_number"`
	Nationality  string                 `json:"nationality"`
	Addresses    []CustomerAddress      `json:"addresses"`
	DateOfBirth  string                 `json:"date_of_birth"`
	Metadata     map[string]interface{} `json:"metadata"`
}

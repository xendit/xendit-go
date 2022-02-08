package xendit

import "time"

type AccountType string
type Status string

const (
	OWNED   AccountType = "OWNED"
	MANAGED AccountType = "MANAGED"
)
const (
	LIVE          Status = "LIVE"
	AWAITING_DOCS Status = "AWAITING_DOCS"
	REGISTERED    Status = "REGISTERED"
	INVITED       Status = "INVITED"
)

type PublicProfile struct {
	BusinessName string `json:"business_name,omitempty"`
}

type Account struct {
	ID            string        `json:"id,omitempty"`
	Created       *time.Time    `json:"created,omitempty"`
	Updated       *time.Time    `json:"updated,omitempty"`
	Type          AccountType   `json:"type,omitempty"`
	Email         string        `json:"email,omitempty"`
	PublicProfile PublicProfile `json:"public_profile,omitempty"`
	Country       string        `json:"country,omitempty"`
	Status        Status        `json:"status,omitempty"`
}

package xendit

import "time"

type Report struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Filter     Filter     `json:"filter"`
	Format     string     `json:"format"`
	Status     string     `json:"status"`
	Url        string     `json:"url"`
	Currency   string     `json:"currency"`
	BusinessID string     `json:"business_id"`
	Created    *time.Time `json:"created"`
	Updated    *time.Time `json:"updated"`
}

type Filter struct {
	From string `json:"from"`
	To   string `json:"to"`
}

package xendit

import "time"

// QRCodeType constants are the available QR Code type
type QRCodeType string

// This consists the values that QRCodeType can have
const (
	DynamicQRCode QRCodeType = "DYNAMIC"
	StaticQRCode  QRCodeType = "STATIC"
)

// QRCode contains data from Xendit's API response of QR code related request.
// For more details see https://developers.xendit.co/api-reference/#create-qr-code
type QRCode struct {
	ID          string     `json:"id"`
	ExternalID  string     `json:"external_id"`
	Amount      float64    `json:"amount"`
	QRString    string     `json:"qr_string"`
	CallbackURL string     `json:"callback_url"`
	Type        QRCodeType `json:"type"`
	Status      string     `json:"status"`
	Created     *time.Time `json:"created"`
	Updated     *time.Time `json:"updated"`
}

// QRDetail contains data from qr field from Xendit's API response of QRCode payments related request
// For more details see https://developers.xendit.co/api-reference/#get-payments-by-external-id
type QRDetail struct {
	ID         string     `json:"id"`
	ExternalID string     `json:"external_id"`
	QRString   string     `json:"qr_string"`
	Type       QRCodeType `json:"type"`
}

// QRCodePayments contains data from Xendit's API response of QRCode payments related request.
// For more details see https://developers.xendit.co/api-reference/#get-payments-by-external-id
type QRCodePayments struct {
	ID      string     `json:"id"`
	Amount  float64    `json:"amount"`
	Created *time.Time `json:"created"`
	Status  string     `json:"status"`
	QRCode  QRDetail   `json:"qr_code"`
}

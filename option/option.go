package option

// Option is ...
type Option struct {
	PublicKey string // customer's public API key
	SecretKey string // customer's secret API key
	XenditURL string // should there be a need to override API base URL
}

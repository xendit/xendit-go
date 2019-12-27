package xendit

// PublicKey ...
var PublicKey string

// SecretKey ...
var SecretKey string

// XenditURL ...
var XenditURL string = "https://api.xendit.co"

var defaultHTTPRequester HTTPRequester = HTTPRequesterImplementation{}

// Option is ...
type Option struct {
	PublicKey string // customer's public API key
	SecretKey string // customer's secret API key
	XenditURL string // should there be a need to override API base URL
}

// GetHTTPRequester returns the defaultHTTPRequester
func GetHTTPRequester() HTTPRequester {
	return defaultHTTPRequester
}

// SetHTTPRequester sets the defaultHTTPRequester
func SetHTTPRequester(httpRequester HTTPRequester) {
	defaultHTTPRequester = httpRequester
}

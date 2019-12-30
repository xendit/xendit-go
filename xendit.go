package xendit

// Opt is the default Option for the API call without API client
var Opt Option = Option{
	XenditURL: "https://api.xendit.co",
}

var defaultHTTPRequester HTTPRequester = HTTPRequesterImplementation{}

// Option is the wrap of the parameters needed for the API call
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

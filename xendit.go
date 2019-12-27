package xendit

// Option is ...
type Option struct {
	PublicKey string // customer's public API key
	SecretKey string // customer's secret API key
	XenditURL string // should there be a need to override API base URL
}

// GetHTTPRequester gets the default implementation of HTTPRequester
func GetHTTPRequester() HTTPRequester {
	httpRequester := HTTPRequesterImplementation{}

	return httpRequester
}

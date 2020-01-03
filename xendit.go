package xendit

import "sync"

// Opt is the default Option for the API call without API client
var Opt Option = Option{
	XenditURL: "https://api.xendit.co",
}

var httpRequesterWrapper HTTPRequesterWrapper = HTTPRequesterWrapper{
	httpRequester: HTTPRequesterImplementation{},
}

// Option is the wrap of the parameters needed for the API call
type Option struct {
	PublicKey string // customer's public API key
	SecretKey string // customer's secret API key
	XenditURL string // should there be a need to override API base URL
}

// HTTPRequesterWrapper is the HTTPRequester with locker for setting the HTTPRequester
type HTTPRequesterWrapper struct {
	httpRequester HTTPRequester
	mu            sync.RWMutex
}

// GetHTTPRequester returns the httpRequester
func GetHTTPRequester() HTTPRequester {
	return httpRequesterWrapper.httpRequester
}

// SetHTTPRequester sets the httpRequester
func SetHTTPRequester(httpRequester HTTPRequester) {
	httpRequesterWrapper.mu.Lock()
	defer httpRequesterWrapper.mu.Unlock()

	httpRequesterWrapper.httpRequester = httpRequester
}

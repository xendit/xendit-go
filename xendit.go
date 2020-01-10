package xendit

import (
	"net/http"
	"sync"
)

// Opt is the default Option for the API call without API client
var Opt Option = Option{
	XenditURL: "https://api.xendit.co",
}

var xdAPIRequesterWrapper XdAPIRequesterWrapper = XdAPIRequesterWrapper{}

var httpClient *http.Client = &http.Client{}

// Option is the wrap of the parameters needed for the API call
type Option struct {
	PublicKey string // customer's public API key
	SecretKey string // customer's secret API key
	XenditURL string // should there be a need to override API base URL
}

// XdAPIRequesterWrapper is the XdAPIRequester with locker for setting the XdAPIRequester
type XdAPIRequesterWrapper struct {
	xdAPIRequester APIRequester
	mu             sync.RWMutex
}

// GetAPIRequester returns the xendit APIRequester
// if it is already created, it will return the created one
// else, it will create a default implementation
func GetAPIRequester() APIRequester {
	if xdAPIRequesterWrapper.xdAPIRequester != nil {
		return xdAPIRequesterWrapper.xdAPIRequester
	}

	xdAPIRequesterWrapper.xdAPIRequester = &APIRequesterImplementation{
		HTTPClient: httpClient,
	}

	return xdAPIRequesterWrapper.xdAPIRequester
}

// SetAPIRequester sets the APIRequester
func SetAPIRequester(xdAPIRequester APIRequester) {
	xdAPIRequesterWrapper.mu.Lock()
	defer xdAPIRequesterWrapper.mu.Unlock()

	xdAPIRequesterWrapper.xdAPIRequester = xdAPIRequester
}

// SetHTTPClient sets the httpClient
func SetHTTPClient(newHTTPClient *http.Client) {
	httpClient = newHTTPClient
}

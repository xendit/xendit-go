package xendit

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

// APIRequester abstraction of HTTP Client that will make API calls to Xendit backend
// `body` is POST-requests' bodies if applicable.
// `result` pointer to value which response string will be unmarshalled to.
type APIRequester interface {
	Call(ctx context.Context, method string, url string, secretKey string, body interface{}, result interface{}) error
}

// APIRequesterImplementation is the default implementation of XdAPIRequester
type APIRequesterImplementation struct {
	HTTPClient *http.Client
}

// Call makes HTTP requests with JSON-format body.
// `body` is POST-requests' bodies if applicable.
// `result` pointer to value which response string will be unmarshalled to.
func (h *APIRequesterImplementation) Call(ctx context.Context, method string, url string, secretKey string, body interface{}, result interface{}) error {
	reqBody := []byte("")
	var req *http.Request
	var err error

	isParamsNil := body == nil || (reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil())

	if !isParamsNil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	req, err = http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return err
	}

	req.SetBasicAuth(secretKey, "")
	req.Header.Set("Content-Type", "application/json")

	if err := h.doRequest(req, result); err != nil {
		return err
	}

	return nil
}

func (h *APIRequesterImplementation) doRequest(req *http.Request, result interface{}) error {
	resp, err := h.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return err
	}

	return nil
}

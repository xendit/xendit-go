package xendit

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// HTTPRequester is an interface for making http requests
// This interface exists to enable mocking for testing
type HTTPRequester interface {
	Call(ctx context.Context, method string, path string, secretKey string, params interface{}, result interface{}) error
}

// HTTPRequesterImplementation is the default implementation of HTTPRequester
type HTTPRequesterImplementation struct {
}

// Call creates a HTTP request
func (h HTTPRequesterImplementation) Call(ctx context.Context, method string, path string, secretKey string, params interface{}, result interface{}) error {
	var body string

	if params != nil {
		reqBody, err := json.Marshal(params)
		if err != nil {
			return err
		}
		body = string(reqBody)

		if method == "GET" {
			// move body to url
			path = h.createGETURL(path, body)
			body = ""
		}
	}

	req, err := h.newRequest(ctx, method, path, secretKey, body)
	if err != nil {
		return err
	}

	if err := h.doRequest(req, result); err != nil {
		return err
	}

	return nil
}

func (h HTTPRequesterImplementation) newRequest(ctx context.Context, method string, path string, secretKey string, body string) (*http.Request, error) {
	var req *http.Request
	var err error

	if ctx != nil {
		req, err = http.NewRequestWithContext(
			ctx,
			method,
			path,
			bytes.NewBuffer([]byte(body)),
		)
	} else {
		req, err = http.NewRequest(
			method,
			path,
			bytes.NewBuffer([]byte(body)),
		)
	}
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(secretKey, "")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (h HTTPRequesterImplementation) doRequest(req *http.Request, result interface{}) error {
	client := &http.Client{}
	resp, err := client.Do(req)
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

// createGETURL creates an url with query string for HTTP GET request
func (h HTTPRequesterImplementation) createGETURL(path string, body string) string {
	base, _ := url.Parse(path)

	if body != "" && body != "null" {
		paramList := convertJSONStringToStringSlice(body[1 : len(body)-1])

		queryParams := url.Values{}
		for _, param := range paramList {
			if param[1][0] == '"' {
				param[1] = param[1][1 : len(param[1])-1]
			}
			queryParams.Add(param[0][1:len(param[0])-1], param[1])
		}
		base.RawQuery = queryParams.Encode()
	}

	return base.String()
}

func removeFromSlice(slice []string, idx int) []string {
	return append(slice[:idx], slice[idx+1:]...)
}

// convertJSONStringToStringSlice converts from JSON string to slice: [[key][value],...]
func convertJSONStringToStringSlice(stringParams string) [][]string {
	var stringSlice [][]string

	stringParamsList := strings.Split(stringParams, ",")

	for i := 0; i < len(stringParamsList); i++ {
		if strings.Contains(stringParamsList[i], "[") {
			stringParamsList[i] = stringParamsList[i] + stringParamsList[i+1]
			stringParamsList = removeFromSlice(stringParamsList, i+1)
			i++
		}

		stringSlice = append(stringSlice, strings.SplitN(stringParamsList[i], ":", 2))
	}

	return stringSlice
}

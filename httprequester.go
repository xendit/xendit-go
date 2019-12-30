package xendit

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strings"
)

// HTTPRequester is an interface for making http requests
// This interface exists to enable mocking for testing
type HTTPRequester interface {
	Call(ctx context.Context, method string, url string, secretKey string, paramsString string) ([]byte, error)
}

// HTTPRequesterImplementation is the default implementation of HTTPRequester
type HTTPRequesterImplementation struct {
}

// Call makes http request
func (h HTTPRequesterImplementation) Call(ctx context.Context, method string, url string, secretKey string, paramsString string) ([]byte, error) {
	var req *http.Request
	var err error

	if method == "GET" {
		req, err = h.getReq(ctx, url, paramsString)
	} else if method == "POST" {
		req, err = h.postReq(ctx, url, paramsString)
	}
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(secretKey, "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (h HTTPRequesterImplementation) getReq(ctx context.Context, url string, paramsString string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}

	if paramsString != "" {
		paramList := convertJSONStringToStringSlice(paramsString[1 : len(paramsString)-1])

		queryParams := req.URL.Query()
		for _, param := range paramList {
			queryParams.Add(param[0], param[1])
		}
		req.URL.RawQuery = queryParams.Encode()
	}

	return req, nil
}

func (h HTTPRequesterImplementation) postReq(ctx context.Context, url string, paramsString string) (*http.Request, error) {
	var req *http.Request
	var err error

	if paramsString == "" {
		req, err = http.NewRequestWithContext(
			ctx,
			"POST",
			url,
			nil,
		)
	} else {
		req, err = http.NewRequestWithContext(
			ctx,
			"POST",
			url,
			bytes.NewBuffer([]byte(paramsString)),
		)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func removeFromSlice(slice []string, idx int) []string {
	return append(slice[:idx], slice[idx+1:]...)
}

func convertJSONStringToStringSlice(stringParams string) [][]string {
	var stringSlice [][]string

	stringParamsList := strings.Split(stringParams, ",")

	loopEnd := len(stringParamsList)
	for i := 0; i < loopEnd; i++ {
		if strings.Contains(stringParamsList[i], "[") {
			stringParamsList[i] = stringParamsList[i] + stringParamsList[i+1]
			stringParamsList = removeFromSlice(stringParamsList, i+1)
			loopEnd--
		}

		stringSlice = append(stringSlice, strings.SplitN(stringParamsList[i], ":", 2))
	}

	return stringSlice
}

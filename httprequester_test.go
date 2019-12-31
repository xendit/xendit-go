package xendit

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestingContext() (context.Context, func()) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	return ctx, cancel
}

func TestHTTPRequesterImplementationCall(t *testing.T) {
	httpRequester := HTTPRequesterImplementation{}

	ctx, cancel := getTestingContext()
	defer cancel()

	respBody, err := httpRequester.Call(ctx, "GET", "https://www.xendit.co/", "", "")

	assert.Nil(t, err)
	assert.NotNil(t, respBody)
}

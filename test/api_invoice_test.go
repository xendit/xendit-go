package xendit

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	xendit "github.com/xendit/xendit-go/v5"
	"github.com/xendit/xendit-go/v5/invoice"
)

func Test_xendit_InvoiceApiService(t *testing.T) {
	err := godotenv.Load(".env.api_test")
	if err != nil {
		log.Fatal("Error loading .env.api_test file")
	}

	apiKey := os.Getenv("DEVELOPMENT_API_KEY")
	if apiKey == "" {
		t.Skip("DEVELOPMENT_API_KEY not set")
	}

	apiClient := xendit.NewClient(apiKey)

	t.Run("Test InvoiceApiService CreateInvoice", func(t *testing.T) {
		createInvoiceRequest := *invoice.NewCreateInvoiceRequest(time.Now().Format("20060102150405"), float64(10000))
		resp, httpRes, err := apiClient.InvoiceApi.
			CreateInvoice(context.Background()).
			CreateInvoiceRequest(createInvoiceRequest).
			Execute()

		if err != nil {
			fmt.Fprintf(os.Stdout, "CreateInvoice Err`: %v\n", err)
		} else {
			r, _ := resp.MarshalJSON()
			fmt.Fprintf(os.Stdout, "CreateInvoice Res`: %v\n", string(r))

			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
			assert.Equal(t, invoice.InvoiceStatus("PENDING"), resp.Status)
		}
	})

	t.Run("Test InvoiceApiService GetInvoiceById", func(t *testing.T) {

		var invoiceId string = "654a103b5e6dfa587b6025c3"

		resp, httpRes, err := apiClient.InvoiceApi.GetInvoiceById(context.Background(), invoiceId).Execute()

		if err != nil {
			fmt.Fprintf(os.Stdout, "GetInvoiceById Err`: %v\n", err)
		} else {
			r, _ := resp.MarshalJSON()
			fmt.Fprintf(os.Stdout, "GetInvoiceById Res`: %v\n", string(r))

			require.NotNil(t, resp)
			assert.Equal(t, 200, httpRes.StatusCode)
		}

	})
}

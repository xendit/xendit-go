package xendit

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/joho/godotenv"
	xendit "github.com/xendit/xendit-go/v3"
	payment_method "github.com/xendit/xendit-go/v3/payment_method"
	"github.com/xendit/xendit-go/v3/payment_request"
)

func Test_xendit_PaymentAPIService(t *testing.T) {

	err := godotenv.Load(".env.api_test")
	if err != nil {
		log.Fatal("Error loading .env.api_test file")
	}

	apiKey := os.Getenv("DEVELOPMENT_API_KEY")
	if apiKey == "" {
		t.Skip("DEVELOPMENT_API_KEY not set")
	}
	apiClient := xendit.NewClient(apiKey)

	t.Run("Test Create Card Payment", func(t *testing.T) {
		// Create Payment Method
		paymentMethodParameters := *payment_method.
			NewPaymentMethodParameters(
				payment_method.PaymentMethodType("CARD"),
				payment_method.PaymentMethodReusability("ONE_TIME_USE"))

		cardHolderName := "John Doe"
		cvv := "123"
		successReturnUrl := "https://redirect.me/goodstuff"
		failureReturnUrl := "https://redirect.me/badstuff"
		cardParams := payment_method.CardParameters{
			Currency: "IDR",
			CardInformation: &payment_method.CardParametersCardInformation{
				CardNumber: "4000000000001091",
				ExpiryMonth: "12",
				ExpiryYear: "2027",
				CardholderName: *payment_method.NewNullableString(&cardHolderName),
				Cvv: *payment_method.NewNullableString(&cvv),
			},
			ChannelProperties: *payment_method.NewNullableCardChannelProperties(&payment_method.CardChannelProperties{
				SuccessReturnUrl: *payment_method.NewNullableString(&successReturnUrl),
				FailureReturnUrl: *payment_method.NewNullableString(&failureReturnUrl),
			}),

		}

		paymentMethodParameters.SetDescription("This is a description")
		paymentMethodParameters.SetCard(cardParams)

		createPMResp, createPMhttpRes, createPMErr := apiClient.PaymentMethodApi.
			CreatePaymentMethod(context.Background()).
			PaymentMethodParameters(paymentMethodParameters).
			Execute()

		if createPMErr != nil {
			fmt.Fprintf(os.Stdout, "Create Card Payment Method Err`: %v\n", createPMErr)
		} else {
			fmt.Fprintf(os.Stdout, "Create Card Payment Method Resp`: %v\n", createPMResp)

			require.NotNil(t, createPMResp)
			assert.Equal(t, 201, createPMhttpRes.StatusCode)
			assert.Equal(t, "PENDING", createPMResp.Status.String())
		}

		// Create Payment Request
		referenceId := time.Now().Format("20060102150405")
		amount := float64(10000)
		paymentRequestParameters := payment_request.PaymentRequestParameters{
			ReferenceId: &referenceId,
			Amount: &amount,
			Currency: "IDR",
			PaymentMethodId: &createPMResp.Id,
		}

		createPRResp, createPRHttpRes, createPRErr := apiClient.PaymentRequestApi.
			CreatePaymentRequest(context.Background()).
			PaymentRequestParameters(paymentRequestParameters).
			Execute()

		if createPRErr != nil {
			fmt.Fprintf(os.Stdout, "Create Card Payment Request Err`: %v\n", createPRErr)
		} else {
			fmt.Fprintf(os.Stdout, "Create Card Payment Request Resp`: %v\n", createPRResp)

			require.NotNil(t, createPRResp)
			assert.Equal(t, 201, createPRHttpRes.StatusCode)
			assert.Equal(t, "REQUIRES_ACTION", createPRResp.Status.String())
			assert.True(t, len(createPRResp.Actions) > 0)
		}
	})

	t.Run("Test Create Direct Debit Payment", func(t *testing.T) {
		// Create Payment Method
		paymentMethodParameters := *payment_method.
			NewPaymentMethodParameters(
				payment_method.PaymentMethodType("DIRECT_DEBIT"),
				payment_method.PaymentMethodReusability("ONE_TIME_USE"))

		mobileNumber := "+62818555988"
		cardlastFour := "8888"
		cardExpiry := "06/24"
		email := "email@email.com"

		directDebitParams := payment_method.DirectDebitParameters{
			ChannelCode: payment_method.DIRECTDEBITCHANNELCODE_BRI,
			ChannelProperties: *payment_method.NewNullableDirectDebitChannelProperties(&payment_method.DirectDebitChannelProperties{
				MobileNumber: *payment_method.NewNullableString(&mobileNumber),
				CardLastFour: *payment_method.NewNullableString(&cardlastFour),
				CardExpiry: *payment_method.NewNullableString(&cardExpiry),
				Email: *payment_method.NewNullableString(&email),
			}),
		}

		paymentMethodParameters.SetDescription("This is a description")
		paymentMethodParameters.SetDirectDebit(directDebitParams)
		paymentMethodParameters.SetCustomerId("cust-59660fb7-dcf2-4bb9-b864-f69b081219d7")

		createPMResp, createPMhttpRes, createPMErr := apiClient.PaymentMethodApi.
			CreatePaymentMethod(context.Background()).
			PaymentMethodParameters(paymentMethodParameters).
			Execute()

		if createPMErr != nil {
			fmt.Fprintf(os.Stdout, "Create DIRECT DEBIT Payment Method Err`: %v\n", createPMErr)
		} else {
			fmt.Fprintf(os.Stdout, "Create DIRECT DEBIT Payment Method Resp`: %v\n", createPMResp)

			require.NotNil(t, createPMResp)
			assert.Equal(t, 201, createPMhttpRes.StatusCode)
			assert.Equal(t, "REQUIRES_ACTION", createPMResp.Status.String())
			assert.True(t, len(createPMResp.Actions) > 0)
		}

		// Authenticate Payment Method
		paymentMethodAuthParameters := *payment_method.NewPaymentMethodAuthParameters("333000")
		authPMResp, authPMHttpResp, authPMErr := apiClient.PaymentMethodApi.AuthPaymentMethod(context.Background(), createPMResp.Id).
        	PaymentMethodAuthParameters(paymentMethodAuthParameters).
        	Execute()

		fmt.Println(authPMResp, authPMHttpResp,authPMErr)
		if authPMErr != nil {
			fmt.Fprintf(os.Stdout, "Authenticate DIRECT DEBIT Payment Method Err`: %v\n", authPMErr)
		} else {
			fmt.Fprintf(os.Stdout, "Authenticate DIRECT DEBIT Payment Method Resp`: %v\n", authPMResp)

			require.NotNil(t, authPMResp)
			assert.Equal(t, 200, authPMHttpResp.StatusCode)
			assert.Equal(t, "ACTIVE", authPMResp.Status.String())
		}

		// Create Payment Request
		referenceId := time.Now().Format("20060102150405")
		amount := float64(10000)
		paymentRequestParameters := payment_request.PaymentRequestParameters{
			ReferenceId: &referenceId,
			Amount: &amount,
			Currency: "IDR",
			PaymentMethodId: &createPMResp.Id,
		}

		createPRResp, createPRHttpRes, createPRErr := apiClient.PaymentRequestApi.
			CreatePaymentRequest(context.Background()).
			PaymentRequestParameters(paymentRequestParameters).
			Execute()

		if createPRErr != nil {
			fmt.Fprintf(os.Stdout, "Create DIRECT DEBIT Payment Request Err`: %v\n", createPRErr)
		} else {
			fmt.Fprintf(os.Stdout, "Create DIRECT DEBIT Payment Request Resp`: %v\n", createPRResp)

			require.NotNil(t, createPRResp)
			assert.Equal(t, 201, createPRHttpRes.StatusCode)
			assert.Equal(t, "REQUIRES_ACTION", createPRResp.Status.String())
			assert.True(t, len(createPRResp.Actions) > 0)
		}
	})

	t.Run("Test Create EWallet Payment", func(t *testing.T) {
		// Create Payment Method
		paymentMethodParameters := *payment_method.
			NewPaymentMethodParameters(
				payment_method.PaymentMethodType("EWALLET"),
				payment_method.PaymentMethodReusability("ONE_TIME_USE"))

		successReturnUrl := "https://redirect.me/goodstuff"
		failureReturnUrl := "https://redirect.me/badstuff"
		cancelReturnUrl := "https://redirect.me/nostuff"
		mobileNumber := "+62818555988"

		ewalletParams := payment_method.EWalletParameters{
			ChannelCode: payment_method.EWALLETCHANNELCODE_OVO,
			ChannelProperties: &payment_method.EWalletChannelProperties{
				SuccessReturnUrl: &successReturnUrl,
				FailureReturnUrl: &failureReturnUrl,
				CancelReturnUrl: &cancelReturnUrl,
				MobileNumber: &mobileNumber,
			},
		}

		paymentMethodParameters.SetDescription("This is a description")
		paymentMethodParameters.SetEwallet(ewalletParams)
		paymentMethodParameters.SetCustomerId("cust-59660fb7-dcf2-4bb9-b864-f69b081219d7")

		createPMResp, createPMhttpRes, createPMErr := apiClient.PaymentMethodApi.
			CreatePaymentMethod(context.Background()).
			PaymentMethodParameters(paymentMethodParameters).
			Execute()

		if createPMErr != nil {
			fmt.Fprintf(os.Stdout, "Create EWallet Payment Method Err`: %v\n", createPMErr)
		} else {
			fmt.Fprintf(os.Stdout, "Create EWallet Payment Method Resp`: %v\n", createPMResp)

			require.NotNil(t, createPMResp)
			assert.Equal(t, 201, createPMhttpRes.StatusCode)
			assert.Equal(t, "ACTIVE", createPMResp.Status.String())
		}

		// Create Payment Request
		referenceId := time.Now().Format("20060102150405")
		amount := float64(10000)
		paymentRequestParameters := payment_request.PaymentRequestParameters{
			ReferenceId: &referenceId,
			Amount: &amount,
			Currency: "IDR",
			PaymentMethodId: &createPMResp.Id,
		}

		createPRResp, createPRHttpRes, createPRErr := apiClient.PaymentRequestApi.
			CreatePaymentRequest(context.Background()).
			PaymentRequestParameters(paymentRequestParameters).
			Execute()

		if createPRErr != nil {
			fmt.Fprintf(os.Stdout, "Create EWallet Payment Request Err`: %v\n", createPRErr)
		} else {
			fmt.Fprintf(os.Stdout, "Create EWallet Payment Request Resp`: %v\n", createPRResp)

			require.NotNil(t, createPRResp)
			assert.Equal(t, 201, createPRHttpRes.StatusCode)
			assert.Equal(t, "PENDING", createPRResp.Status.String())
		}
	})

	t.Run("Test Create Over The Counter Payment", func(t *testing.T) {
		// Create Payment Request with Bunddled Payment Method
		amount := float64(10000)
		overTheCounter := payment_request.OverTheCounterParameters{
			ChannelCode: payment_request.OVERTHECOUNTERCHANNELCODE_ALFAMART,
			ChannelProperties: *payment_request.NewOverTheCounterChannelProperties("John Doe"),
		}
		nullableOverTheCounter := payment_request.NewNullableOverTheCounterParameters(&overTheCounter)

		paymentMethod := payment_request.PaymentMethodParameters{
			Type: payment_request.PAYMENTMETHODTYPE_OVER_THE_COUNTER,
			Reusability: payment_request.PAYMENTMETHODREUSABILITY_ONE_TIME_USE,
			OverTheCounter: *nullableOverTheCounter,
		}
		paymentRequestParameters := payment_request.PaymentRequestParameters{
			Amount: &amount,
			Currency: "IDR",
			PaymentMethod: &paymentMethod,
		}

		createPRResp, createPRHttpRes, createPRErr := apiClient.PaymentRequestApi.
			CreatePaymentRequest(context.Background()).
			PaymentRequestParameters(paymentRequestParameters).
			Execute()

		if createPRErr != nil {
			fmt.Fprintf(os.Stdout, "Create Over The Counter Payment Request Err`: %v\n", createPRErr)
		} else {
			fmt.Fprintf(os.Stdout, "Create Over The Counter Payment Request Resp`: %v\n", createPRResp)

			require.NotNil(t, createPRResp)
			assert.Equal(t, 201, createPRHttpRes.StatusCode)
			assert.Equal(t, "PENDING", createPRResp.Status.String())
		}
	})

	t.Run("Test Create QR Payment", func(t *testing.T) {
		// Create Payment Request with Bunddled Payment Method
		amount := float64(10000)

		qrCode := payment_request.QRCodeParameters{
			ChannelCode: *payment_request.NewNullableQRCodeChannelCode(payment_request.QRCODECHANNELCODE_DANA.Ptr()),
		}

		paymentMethod := payment_request.PaymentMethodParameters{
			Type: payment_request.PAYMENTMETHODTYPE_QR_CODE,
			Reusability: payment_request.PAYMENTMETHODREUSABILITY_ONE_TIME_USE,
			QrCode: *payment_request.NewNullableQRCodeParameters(&qrCode),
		}
		paymentRequestParameters := payment_request.PaymentRequestParameters{
			Amount: &amount,
			Currency: "IDR",
			PaymentMethod: &paymentMethod,
		}

		createPRResp, createPRHttpRes, createPRErr := apiClient.PaymentRequestApi.
			CreatePaymentRequest(context.Background()).
			PaymentRequestParameters(paymentRequestParameters).
			Execute()

		if createPRErr != nil {
			fmt.Fprintf(os.Stdout, "Create QR Payment Request Err`: %v\n", createPRErr)
		} else {
			fmt.Fprintf(os.Stdout, "Create QR Payment Request Resp`: %v\n", createPRResp)

			require.NotNil(t, createPRResp)
			assert.Equal(t, 201, createPRHttpRes.StatusCode)
			assert.Equal(t, "PENDING", createPRResp.Status.String())
		}
	})

	t.Run("Test Virtual Account Payment", func(t *testing.T) {
		// Create Payment Request with Bunddled Payment Method
		amount := float64(10000)
		virtualAccount := payment_request.VirtualAccountParameters{
			ChannelCode: payment_request.VIRTUALACCOUNTCHANNELCODE_BRI,
			ChannelProperties: *payment_request.NewVirtualAccountChannelProperties("John Doe"),
		}

		paymentMethod := payment_request.PaymentMethodParameters{
			Type: payment_request.PAYMENTMETHODTYPE_QR_CODE,
			Reusability: payment_request.PAYMENTMETHODREUSABILITY_ONE_TIME_USE,
			VirtualAccount: *payment_request.NewNullableVirtualAccountParameters(&virtualAccount),
		}
		paymentRequestParameters := payment_request.PaymentRequestParameters{
			Amount: &amount,
			Currency: "IDR",
			PaymentMethod: &paymentMethod,
		}

		createPRResp, createPRHttpRes, createPRErr := apiClient.PaymentRequestApi.
			CreatePaymentRequest(context.Background()).
			PaymentRequestParameters(paymentRequestParameters).
			Execute()

		if createPRErr != nil {
			fmt.Fprintf(os.Stdout, "Create Virtual Account Request Err`: %v\n", createPRErr)
		} else {
			fmt.Fprintf(os.Stdout, "Create Virtual Account Payment Request Resp`: %v\n", createPRResp)

			require.NotNil(t, createPRResp)
			assert.Equal(t, 201, createPRHttpRes.StatusCode)
			assert.Equal(t, "PENDING", createPRResp.Status.String())
		}
	})

	t.Run("Get Payment Method By Id", func (t *testing.T){
		resp, httpResp, err := apiClient.PaymentMethodApi.GetPaymentMethodByID(
			context.Background(),
			"pm-89a09e44-3a9f-4bf3-903e-3f68ec170723",
		).Execute()

		if err != nil {
			fmt.Fprintf(os.Stdout, "Get Payment Method By Id Err`: %v\n", err)
		} else {
			fmt.Fprintf(os.Stdout, "Get Payment Method By Id Resp`: %v\n", resp)

			require.NotNil(t, resp)
			assert.Equal(t, 200, httpResp.StatusCode)
		}
	})

	t.Run("Get Payment Request By Id", func (t *testing.T){
		resp, httpResp, err := apiClient.PaymentRequestApi.GetPaymentRequestByID(
			context.Background(),
			"pr-6fd4595a-d6da-4939-9b65-b9f57ebf78dc",
		).Execute()

		if err != nil {
			fmt.Fprintf(os.Stdout, "Get Payment Request By Id Err`: %v\n", err)
		} else {
			fmt.Fprintf(os.Stdout, "Get Payment Request By Id Resp`: %v\n", resp)

			require.NotNil(t, resp)
			assert.Equal(t, 200, httpResp.StatusCode)
		}
	})
}

package disbursement_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xendit/xendit-go/disbursement"
)

func TestPayload_Completed(t *testing.T) {
	payload := `{
	    "id": "57e214ba82b034c325e84d6e",
	    "created": "2021-07-10T08:15:03.404Z",
	    "updated": "2021-07-10T08:15:03.404Z",
	    "external_id": "disbursement_123124123",
	    "user_id": "57c5aa7a36e3b6a709b6e148",
	    "amount": 150000,
	    "bank_code": "BCA",
	    "account_holder_name": "MICHAEL CHEN",
	    "disbursement_description": "Refund for shoes",
	    "status": "COMPLETED",
	    "is_instant": true
	}`

	expectedPayload := disbursement.DisbursementCallback{
		ID:                      "57e214ba82b034c325e84d6e",
		Created:                 "2021-07-10T08:15:03.404Z",
		Updated:                 "2021-07-10T08:15:03.404Z",
		ExternalID:              "disbursement_123124123",
		UserID:                  "57c5aa7a36e3b6a709b6e148",
		Amount:                  150000,
		BankCode:                "BCA",
		AccountHolderName:       "MICHAEL CHEN",
		DisbursementDescription: "Refund for shoes",
		Status:                  "COMPLETED",
		IsInstant:               true,
	}

	var actualPayload disbursement.DisbursementCallback
	err := json.Unmarshal([]byte(payload), &actualPayload)
	assert.NoError(t, err)

	assert.Equal(t, expectedPayload, actualPayload)
}

func TestPayload_Failed(t *testing.T) {
	payload := `{
	    "id": "57e214ba82b034c325e84d6e",
	    "created": "2021-07-10T08:15:03.404Z",
	    "updated": "2021-07-10T08:15:03.404Z",
	    "external_id": "disbursement_123124123",
	    "user_id": "57c5aa7a36e3b6a709b6e148",
	    "amount": 150000,
	    "bank_code": "BCA",
	    "account_holder_name": "MICHAEL CHEN",
	    "disbursement_description": "Refund for shoes",
	    "status": "FAILED",
	    "failure_code": "INVALID_DESTINATION",
	    "is_instant": true
	}`

	expectedPayload := disbursement.DisbursementCallback{
		ID:                      "57e214ba82b034c325e84d6e",
		Created:                 "2021-07-10T08:15:03.404Z",
		Updated:                 "2021-07-10T08:15:03.404Z",
		ExternalID:              "disbursement_123124123",
		UserID:                  "57c5aa7a36e3b6a709b6e148",
		Amount:                  150000,
		BankCode:                "BCA",
		AccountHolderName:       "MICHAEL CHEN",
		DisbursementDescription: "Refund for shoes",
		Status:                  "FAILED",
		FailureCode:             "INVALID_DESTINATION",
		IsInstant:               true,
	}

	var actualPayload disbursement.DisbursementCallback
	err := json.Unmarshal([]byte(payload), &actualPayload)
	assert.NoError(t, err)

	assert.Equal(t, expectedPayload, actualPayload)
}

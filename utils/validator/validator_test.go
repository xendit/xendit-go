package validator_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
	"net/http"
	"testing"
)

func TestValidateRequired(t *testing.T) {
	type testStruct struct {
		Foo string `validate:"required"`
	}
	testCases := []struct {
		desc     string
		input    *testStruct
		expected error
	}{
		{
			desc:     "should return error with missing fields' names",
			input:    &testStruct{},
			expected: errors.New("Missing required fields: 'Foo'"),
		},
		{
			desc: "should return nil on no missing fields",
			input: &testStruct{
				Foo: "Bar",
			},
			expected: nil,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := validator.ValidateRequired(context.Background(), tC.input)
			assert.EqualValues(t, tC.expected, err)
		})
	}
}

func TestAPIValidatorErr(t *testing.T) {
	t.Run("should return xendit.Error{400, `API_VALIDATION_ERROR`, msg} with message from its input error", func(t *testing.T) {
		fakeMsg := "fake validation err message"
		err := validator.APIValidatorErr(errors.New(fakeMsg))
		assert.EqualValues(t, &xendit.Error{
			Status:    http.StatusBadRequest,
			ErrorCode: xendit.APIValidationErrCode,
			Message:   fakeMsg,
		}, err)
	})
}

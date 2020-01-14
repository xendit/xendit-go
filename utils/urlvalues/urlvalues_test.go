package urlvalues_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/xendit/xendit-go/utils/urlvalues"
	"net/url"
	"testing"
	"time"
)

func TestAddTimeToURLValues(t *testing.T) {
	nonzeroTime := time.Now()

	testCases := []struct {
		desc     string
		t        time.Time
		expected string
	}{
		{
			desc:     "should add non-zero time value in RFC3399 format",
			t:        nonzeroTime,
			expected: nonzeroTime.Format(time.RFC3339),
		},
		{
			desc:     "should not add zero time value",
			t:        time.Time{},
			expected: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			u := &url.Values{}
			fieldname := "time"
			urlvalues.AddTimeToURLValues(u, tC.t, fieldname)
			assert.Equal(t, tC.expected, u.Get(fieldname))
		})
	}
}

func TestAddStringSliceToURLValues(t *testing.T) {
	testCases := []struct {
		desc                string
		sl                  []string
		fieldname           string
		expectedQueryFields map[string]string
	}{
		{
			desc:                "should append values from string slice to URL query values",
			sl:                  []string{"foo", "bar"},
			fieldname:           "test",
			expectedQueryFields: map[string]string{"test[0]": "foo", "test[1]": "bar"},
		},
		{
			desc:                "should not append any value from nil slice",
			sl:                  nil,
			fieldname:           "test",
			expectedQueryFields: map[string]string{"test[0]": ""},
		},
		{
			desc:                "should not append any value from empty slice",
			sl:                  []string{},
			fieldname:           "test",
			expectedQueryFields: map[string]string{"test[0]": ""},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			u := &url.Values{}
			urlvalues.AddStringSliceToURLValues(u, tC.sl, tC.fieldname)
			for i, v := range tC.expectedQueryFields {
				assert.Equal(t, v, u.Get(i))
			}
		})
	}
}

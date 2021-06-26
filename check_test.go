package tld_test

import (
	"testing"

	"github.com/lawzava/go-tld"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		input          string
		expectedOutput bool
	}{
		{"com", true},
		{"org", true},
		{"xyz", true},
		{"dev", true},
		{"xir", false},
		{"netlink", false},
		{"015", false},
	}

	for _, testCase := range testCases {
		res := tld.IsValid(testCase.input)

		assert.Equal(t, testCase.expectedOutput, res, testCase.input)
	}
}

package tld

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct{
		input string
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
		res := IsValid(testCase.input)

		assert.Equal(t, testCase.expectedOutput, res, testCase.input)
	}
}

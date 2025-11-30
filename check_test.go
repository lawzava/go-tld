package tld_test

import (
	"testing"

	"github.com/lawzava/go-tld"
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

	for _, tc := range testCases { //nolint:varnamelen // test case name
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()

			res := tld.IsValid(tc.input)

			if res != tc.expectedOutput {
				t.Errorf("IsValid(%q) = %v, want %v", tc.input, res, tc.expectedOutput)
			}
		})
	}
}

// BenchmarkIsValid_FirstElement benchmarks lookup of the first TLD (best case for linear search).
func BenchmarkIsValid_FirstElement(b *testing.B) {
	for b.Loop() {
		tld.IsValid("aaa")
	}
}

// BenchmarkIsValid_MiddleElement benchmarks lookup of a TLD in the middle of the list.
func BenchmarkIsValid_MiddleElement(b *testing.B) {
	for b.Loop() {
		tld.IsValid("com")
	}
}

// BenchmarkIsValid_LastElement benchmarks lookup of the last TLD (worst case for linear search).
func BenchmarkIsValid_LastElement(b *testing.B) {
	for b.Loop() {
		tld.IsValid("zw")
	}
}

// BenchmarkIsValid_NotFound benchmarks lookup of a non-existent TLD (worst case).
func BenchmarkIsValid_NotFound(b *testing.B) {
	for b.Loop() {
		tld.IsValid("notarealtld")
	}
}

// BenchmarkIsValid_Parallel benchmarks concurrent lookups.
func BenchmarkIsValid_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tld.IsValid("com")
		}
	})
}

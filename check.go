//go:generate go run internal/gen/main.go

package tld

// IsValid checks whether the supplied TLD is valid.
func IsValid(tld string) bool {
	_, ok := validTLDs[tld]
	return ok
}

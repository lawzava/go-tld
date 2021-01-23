package tld

// IsValid checks whether the supplied request is valid tld.
func IsValid(tld string) bool {
	for _, validTLD := range availableTLDs() {
		if validTLD == tld {
			return true
		}
	}

	return false
}

![GolangCI](https://github.com/lawzava/go-tld/workflows/golangci/badge.svg?branch=main)
[![Version](https://img.shields.io/github/v/release/lawzava/go-tld)](https://github.com/lawzava/go-tld/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/lawzava/go-tld)](https://goreportcard.com/report/github.com/lawzava/go-tld)
[![Coverage Status](https://coveralls.io/repos/github/lawzava/go-tld/badge.svg?branch=main)](https://coveralls.io/github/lawzava/go-tld?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/lawzava/go-tld.svg)](https://pkg.go.dev/github.com/lawzava/go-tld)

# go-tld

Minimalistic library for checking whether a string is a valid top-level domain.

The list of valid TLDs is generated from the official
[IANA list](https://data.iana.org/TLD/tlds-alpha-by-domain.txt). Each release
embeds the list as it stood when the release was cut; the header of `list.go`
records the IANA version.

## Installation

```
go get github.com/lawzava/go-tld
```

## Usage

```go
package main

import "github.com/lawzava/go-tld"

func main() {
	tld.IsValid("com") // true
	tld.IsValid("xir") // false
}
```

## Refreshing the list

```
go generate ./...
```

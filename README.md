![GolangCI](https://github.com/lawzava/go-tld/workflows/golangci/badge.svg?branch=main)
[![Version](https://img.shields.io/badge/version-v1.0.0-green.svg)](https://github.com/lawzava/go-tld/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/lawzava/go-tld)](https://goreportcard.com/report/github.com/lawzava/go-tld)
[![Coverage Status](https://coveralls.io/repos/github/lawzava/go-tld/badge.svg?branch=main)](https://coveralls.io/github/lawzava/go-tld?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/lawzava/go-tld.svg)](https://pkg.go.dev/github.com/lawzava/go-tld)

# go-tld

Minimalistic library for keeping up to date with TLD checks. 

## Installation

```
go get github.com/lawzava/go-tld
```

## Usage

```go
package main

import tld "github.com/lawzava/go-tld"

func main() {
	tld.IsValid("com") // true
	tld.IsValid("xir") // false
}
```

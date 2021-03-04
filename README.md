# jsonmap.Ordered

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/dolmen-go/jsonmap)
[![Travis-CI](https://api.travis-ci.org/dolmen-go/jsonmap.svg?branch=master)](https://travis-ci.org/dolmen-go/jsonmap)
[![Go Report Card](https://goreportcard.com/badge/github.com/dolmen-go/jsonmap)](https://goreportcard.com/report/github.com/dolmen-go/jsonmap)

Package `jsonmap` provides tools to serialize a map as JSON using a given order for keys.

```go
type Ordered struct {
    Order []string // Order of keys
    Data map[string]interface{}
}
package main

import (
	"strings"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	a := strings.Repeat("A", len(b))
	_ = copy(b, a)
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

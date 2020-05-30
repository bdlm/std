package parser

import (
	"io"
)

// Parser is the interface that wraps the standard Parse method.
type Parser interface {
	// Parse accepts a reader and parses the data it returns, returning an
	// error if parsing fails.
	//
	// Implementations must not retain r. Implementations should not retain
	// any parsed data if returning an error.
	Parse(r io.Reader) (interface{}, error)
}

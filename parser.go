package std

import "io"

/*
Parser is the interface that wraps the basic Parse method.

Parser parses the provided data p, returning an error if parsing fails.

Implementations must not retain p. Implementations should not retain
any parsed data if returning an error.
*/
type Parser interface {
	Parse(r io.Reader) (interface{}, error)
}

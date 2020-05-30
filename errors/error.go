package error

import (
	std_caller "github.com/bdlm/std/v2/caller"
)

// Error defines a robust error stack interface.
type Error interface {
	// Error implements error.
	Error() string

	// Has tests to see if the test error exists anywhere in the error
	// stack.
	Has(error) bool

	// Is tests to see if the test error matches most recent error in the
	// stack.
	Is(error) bool
}

// ErrorCaller is the interface implemented by error types that can expose
// caller information about themselves.
type ErrorCaller interface {
	// Caller returns the associated Caller instance.
	Caller() std_caller.Caller
}

package error

// Error defines a robust error stack interface.
type Error interface {
	// Caller returns the associated Caller instance.
	Caller() Caller

	// Error implements error.
	Error() string

	// Has tests to see if the test error exists anywhere in the error
	// stack.
	Has(error) bool

	// Is tests to see if the test error matches most recent error in the
	// stack.
	Is(error) bool

	// Unwrap returns the next error, if any.
	Unwrap() Error
}

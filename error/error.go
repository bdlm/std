package error

// Error defines a slightly more robust error interface than the standard
// library.
type Error interface {
	// Code returns the associated error code. A value of 0 should be
	// considered uncoded.
	Code() Code
	// Caller returns the runtime caller information for this error.
	Caller() Caller
	// Error implements standard library compatibility.
	Error() string
	// String is an alias of Error.
	String() string
	// Trace returns the full stack trace of this error.
	Trace() Trace
}

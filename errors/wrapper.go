package error

type ErrorWrapper interface {
	// WrapE returns the given error as a wrapper around the receiver.
	WrapE(error) error
}

type Wrapper interface {
	// Wrap returns a new error defined by string that wraps the receiver.
	Wrap(string) error
}

type Unwrapper interface {
	// Unwrap returns the wrapped error, if any, otherwise nil.
	Unwrap() error
}

// Package sandboxerr should do something like https://github.com/go-errors/errors
package sandboxerr

// ErrorWithMeta is an error type that includes
// some extra context information
type ErrorWithMeta interface {
	error
	StackTrace() []byte
}

// Wrap adds a stacktrace to the error. If the err is already
// wrapped, it will not be wrapped again
func Wrap(err error) ErrorWithMeta {
	// todo
	return nil
}

// ValidationError ...
type ValidationError interface {
	FieldErrors() map[string]string
	error
}

// GoValidatorError parses validation errors from
// https://github.com/asaskevich/govalidator
func GoValidatorError(err error) ValidationError {
	// todo
	return nil
}

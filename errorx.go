package errorx

import (
	"errors"
	"fmt"
	"strings"
)

// Newf returns an error according to a format specifier and param.
// Each call to New returns a distinct error value even if the text is identical.
//
// If tracing is enabled then stacktrace is attached to the returned error.
func Newf(format string, a ...any) error {
	// NOTE: go vet printf check doesn't understand inverse expression.
	if !Tracing() || strings.Contains(format, "%w") {
		return fmt.Errorf(format, a...)
	}

	msg := format
	if a != nil {
		msg = fmt.Sprintf(format, a...)
	}
	return &errorString{
		s:     msg,
		frame: caller(1),
	}
}

// Is is just [errors.Is].
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As is just [errors.As].
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Into is type-safe alternative to [errorx.As].
func Into[T error](err error) (T, bool) {
	var dst T
	ok := As(err, &dst)
	return dst, ok
}

// Unwrap is just [errors.Unwrap].
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// IsAny is just a multiple [errors.Is] calls.
func IsAny(err, target error, targets ...error) bool {
	if errors.Is(err, target) {
		return true
	}

	for _, t := range targets {
		if errors.Is(err, t) {
			return true
		}
	}
	return false
}

// Must returns value or panic if the error is non-nil.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// errorString same as [errors.errorString] but with a frame field.
type errorString struct {
	s     string
	frame Frame
}

func (e *errorString) Error() string {
	return e.s
}

func (e *errorString) Format(s fmt.State, v rune) {
	FormatError(e, s, v)
}

func (e *errorString) FormatError(p Printer) (next error) {
	p.Print(e.s)
	e.frame.Format(p)
	return nil
}

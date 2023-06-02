package errorx

import (
	"errors"
	"fmt"
	"strings"
)

func Newf(format string, a ...any) error {
	// NOTE: go vet printf check doesn't understand inverse expression.
	if !IsTracingEnabled() || strings.Contains(format, "%w") {
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

// Unwrap is just [errors.Unwrap].
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

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

package errorx

import (
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
)

var traceMode atomic.Int32

// Tracing reports whether trace is attached to the errors.
func Tracing() bool { return traceMode.Load() == 0 }

// EnableTrace collection for errors.
func EnableTrace() { traceMode.Store(0) }

// DisableTrace collection for errors.
func DisableTrace() { traceMode.Store(1) }

// Newf creates a new error with a formatted message and caller info if tracing is enabled.
// Fallback to [fmt.Errorf] if '%w' is presented and tracing is disabled.
//
// Each call to New returns a distinct error value even if the text is identical.
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
		frame: caller(),
	}
}

// Trace returns a wrapped error with a caller info if tracing is enabled.
// Returns nil if err is nil.
func Trace(err error) error {
	if err == nil {
		return nil
	}

	var frame Frame
	if Tracing() {
		frame = caller()
	}
	return &wrapError{
		err:   err,
		frame: frame,
	}
}

// Wrapf returns a wrapped error with a formatted message and caller info if tracing is enabled.
// Returns nil if err is nil.
func Wrapf(err error, format string, a ...any) error {
	if err == nil {
		return nil
	}

	var frame Frame
	if Tracing() {
		frame = caller()
	}

	msg := format
	if a != nil {
		msg = fmt.Sprintf(format, a...)
	}
	return &wrapError{
		err:   err,
		msg:   msg,
		frame: frame,
	}
}

// Is is just a [errors.Is] with multiple targets.
func Is(err, target error, targets ...error) bool {
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

// As is just [errors.As].
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Into is a typed alternative to [errorx.As].
func Into[T error](err error) (T, bool) {
	var dst T
	ok := As(err, &dst)
	return dst, ok
}

// Unwrap is just [errors.Unwrap].
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Join is just [errors.Join].
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// UnwrapMulti is just [errors.Unwrap] for errors created via [errors.Join].
// Returns slice with 1 element if a regular error is passed.
// Returns nil if err is nil.
func UnwrapMulti(err error) []error {
	if err == nil {
		return nil
	}

	// multiError represents error created by [errors.Join].
	// There is no other way to unwrap it without casting to this interface.
	type multiError interface {
		Unwrap() []error
	}

	errs, ok := err.(multiError)
	if ok {
		return errs.Unwrap()
	}
	return []error{err}
}

// Must returns value or panic if the error is non-nil.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Must2 returns second value or panic if the error is non-nil.
func Must2[T any, U any](_ T, u U, err error) U {
	if err != nil {
		panic(err)
	}
	return u
}

// Must3 returns third value or panic if the error is non-nil.
func Must3[T any, U any, V any](_ T, _ U, v V, err error) V {
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

type wrapError struct {
	err   error
	msg   string
	frame Frame
}

func (e *wrapError) Error() string {
	return fmt.Sprint(e)
}

func (e *wrapError) Format(s fmt.State, v rune) {
	FormatError(e, s, v)
}

func (e *wrapError) FormatError(p Printer) (next error) {
	p.Print(e.msg)
	e.frame.Format(p)
	return e.err
}

func (e *wrapError) Unwrap() error {
	return e.err
}

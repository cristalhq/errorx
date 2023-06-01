package errorx

import (
	"errors"
	"fmt"
	"strings"
)

func New(text string) error {
	if IsTracingEnabled() {
		return &errorString{
			s:     text,
			frame: caller(1),
		}
	}
	return errors.New(text)
}

func Newf(format string, a ...any) error {
	if IsTracingEnabled() && !strings.Contains(format, "%w") {
		return &errorString{
			s:     fmt.Sprintf(format, a...),
			frame: caller(1),
		}
	}
	return fmt.Errorf(format, a...)
}

// Is is just [errors.Is].
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As is just [errors.As].
func As(err error, target any) bool {
	return errors.As(err, target)
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

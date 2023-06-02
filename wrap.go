package errorx

import "fmt"

func Trace(err error) error {
	if err == nil {
		return nil
	}

	frame := Frame{}
	if IsTracingEnabled() {
		frame = caller(1)
	}
	return &wrapError{
		err:   err,
		frame: frame,
	}
}

func Wrapf(err error, format string, a ...any) error {
	if err == nil {
		return nil
	}

	frame := Frame{}
	if IsTracingEnabled() {
		frame = caller(1)
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

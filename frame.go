package errorx

import (
	"runtime"
)

type Frame struct {
	frames [3]uintptr
}

func (f Frame) Location() (function, file string, line int) {
	frames := runtime.CallersFrames(f.frames[:])
	if _, ok := frames.Next(); !ok {
		return "", "", 0
	}

	fr, ok := frames.Next()
	if !ok {
		return "", "", 0
	}
	return fr.Function, fr.File, fr.Line
}

func (f Frame) Format(p Printer) {
	if p.Detail() {
		function, file, line := f.Location()
		if function != "" {
			p.Printf("%s\n    ", function)
		}
		if file != "" {
			p.Printf("%s:%d\n", file, line)
		}
	}
}

// Caller returns a Frame that describes a frame on the caller's stack.
// The argument skip is the number of frames to skip over.
// Caller(0) returns the frame for the caller of Caller.
func caller() Frame {
	var s Frame
	runtime.Callers(2, s.frames[:])
	return s
}

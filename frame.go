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

func caller(skip int) Frame {
	var s Frame
	runtime.Callers(skip+1, s.frames[:])
	return s
}

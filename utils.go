package errorx

import "runtime"

func caller(skip int) Frame {
	var s Frame
	runtime.Callers(skip+1, s.frames[:])
	return s
}

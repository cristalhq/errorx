package errorx

import "sync/atomic"

var traceFlag int32

// IsTracingEnabled reports whether trace is attached to the errors.
func IsTracingEnabled() bool { return atomic.LoadInt32(&traceFlag) == 0 }

// EnableTrace collection for errors.
func EnableTrace() { atomic.StoreInt32(&traceFlag, 0) }

// DisableTrace collection for errors.
func DisableTrace() { atomic.StoreInt32(&traceFlag, 1) }

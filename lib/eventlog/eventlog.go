package eventlog

import "sync/atomic"

var enabled atomic.Bool

func Enabled() bool {
	return enabled.Load()
}

func Toggle() bool {
	next := !enabled.Load()
	enabled.Store(next)
	return next
}

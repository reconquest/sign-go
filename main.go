package sign

import (
	"os"
	"os/signal"
)

type Handler func(os.Signal) bool

// Notify calls handler when gets specified signals and pass given signal to
// handler. If handler returns false, notify stops waiting for signals.
func Notify(handler Handler, signals ...os.Signal) {
	pipe := make(chan os.Signal, 1)
	signal.Notify(pipe, signals...)

	for sign := range pipe {
		if !handler(sign) {
			return
		}
	}
}

package ctrlc

import (
	"os"
	"os/signal"
	"syscall"
)

// Watch watches when `ctrl-c` is pressed.
func Watch(fn func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		fn()
	}()
}

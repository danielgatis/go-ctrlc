package ctrlc

import (
	"os"
	"os/signal"
	"syscall"
)

func ctrlc(fn func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		fn()
	}()
}

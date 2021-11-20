package signaler

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	s = make(chan os.Signal, 1)
)

func init() {
	sigs := []os.Signal{
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
		syscall.SIGABRT,
	}
	signal.Notify(s, sigs...)
}

func Wait() os.Signal {
	return <-s
}

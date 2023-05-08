package server

import (
	"os"
	"os/signal"
)

var setonce = make(chan struct{})

var shutdownc chan os.Signal

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() <-chan struct{} {
	close(setonce) // channel cannot be close repeatedly, so there will throw a panic when called twice

	shutdownc = make(chan os.Signal, 2)

	stop := make(chan struct{})

	signal.Notify(shutdownc, shutdownSignals...)

	go func() {
		<-shutdownc
		close(stop)
		<-shutdownc
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

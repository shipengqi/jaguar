//go:build linux
// +build linux

package managers

import (
	"syscall"
	"testing"
	"time"

	"github.com/jaguar/apiskeleton/pkg/shutdown"
)

type startShutdownFunc func(sm shutdown.ShutdownManager)

func (f startShutdownFunc) StartShutdown(sm shutdown.ShutdownManager) {
	f(sm)
}

func (f startShutdownFunc) ReportError(err error) {
}

func (f startShutdownFunc) AddShutdownCallback(shutdownCallback shutdown.ShutdownCallback) {}

func waitSig(t *testing.T, c <-chan int) {
	select {
	case <-c:

	case <-time.After(1 * time.Second):
		t.Error("Timeout waiting for StartShutdown.")
	}
}

func TestStartShutdownCalledOnDefaultSignals(t *testing.T) {
	c := make(chan int, 100)

	psm := NewPosixSignalManager()
	psm.Start(startShutdownFunc(func(sm shutdown.ShutdownManager) {
		c <- 1
	}))

	time.Sleep(time.Millisecond)

	syscall.Kill(syscall.Getpid(), syscall.SIGINT)

	waitSig(t, c)

	psm.Start(startShutdownFunc(func(sm shutdown.ShutdownManager) {
		c <- 1
	}))

	time.Sleep(time.Millisecond)

	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

	waitSig(t, c)
}

func TestStartShutdownCalledCustomSignal(t *testing.T) {
	c := make(chan int, 100)

	psm := NewPosixSignalManager(syscall.SIGHUP)
	psm.Start(startShutdownFunc(func(sm shutdown.ShutdownManager) {
		c <- 1
	}))

	time.Sleep(time.Millisecond)

	syscall.Kill(syscall.Getpid(), syscall.SIGHUP)

	waitSig(t, c)
}
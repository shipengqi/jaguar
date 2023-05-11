package spinner

import (
	"testing"
	"time"
)

func TestSpinner_Start(t *testing.T) {
	s := New().WithEndMsg("\n")
	s.Start()
	time.Sleep(time.Second * 1)
	s.Stop()
}

func TestSpinner_Start2(t *testing.T) {
	s := New().WithEndMsg("\n").WithChars(4)
	s.Start()
	time.Sleep(time.Second * 5)
	s.Stop()
}

func TestSpinner_WithPrefix(t *testing.T) {
	s := New().WithPrefix("Checking ................ [ ").
		WithSuffix(" ]").
		WithEndMsg("\n")
	s.Start()
	time.Sleep(time.Second * 1)
	s.Stop()
}

func TestSpinner_Multi_Start(t *testing.T) {
	s := New().WithPrefix("Checking test1 ................ [ ").
		WithSuffix(" ]").
		WithEndMsg("\n")
	s.Start()
	time.Sleep(time.Second * 1)
	s.Stop()
	s2 := New().WithPrefix("Checking test2 ................ [ ").
		WithSuffix(" ]").
		WithEndMsg("\n")
	s2.Start()
	time.Sleep(time.Second * 1)
	s2.Stop()
}

func TestSpinner_StopWithStatus(t *testing.T) {
	s := New().WithPrefix("Checking test1 ................ [ ").
		WithSuffix(" ]").
		WithEndMsg("\n")
	s.Start()
	time.Sleep(time.Second * 1)
	s.StopWithStatus("PASS")

	s2 := New().WithPrefix("Checking test2 ................ [ ").
		WithSuffix(" ]").
		WithEndMsg("\n")
	s2.Start()
	time.Sleep(time.Second * 1)
	s2.StopWithStatus("PASS")
}

func TestSpinner_Reset(t *testing.T) {
	s := New().WithPrefix("Checking restart1 ................ [ ").
		WithSuffix(" ]").
		WithEndMsg("\n")
	s.Start()
	time.Sleep(time.Second * 1)
	s.StopWithStatus("PASS")

	s.Reset().WithPrefix("Checking restart2 ................ [ ").Start()
	time.Sleep(time.Second * 1)
	s.StopWithStatus("PASS")
}
package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// type SpySleeper struct {
// 	Calls int
// }

// func (s *SpySleeper) Sleep() {
// 	s.Calls++
// }

const writeOP = "write"
const sleepOP = "sleep"

type SuperSpy struct {
	Calls []string
}

func (ss *SuperSpy) Write(p []byte) (int, error) {
	ss.Calls = append(ss.Calls, writeOP)
	return 0, nil
}

func (ss *SuperSpy) Sleep() {
	ss.Calls = append(ss.Calls, sleepOP)

}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func TestCountdown(t *testing.T) {
	t.Run("Prints 3,2,1,Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		superSpy := &SuperSpy{}

		Countdown(buffer, superSpy)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("Sleeps before every write", func(t *testing.T) {

		superSpy := &SuperSpy{}

		Countdown(superSpy, superSpy)

		want := []string{
			"sleep", "write",
			"sleep", "write",
			"sleep", "write",
			"sleep", "write",
		}
		if !reflect.DeepEqual(want, superSpy.Calls) {
			t.Errorf("got %q want %q", superSpy.Calls, want)
		}

	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

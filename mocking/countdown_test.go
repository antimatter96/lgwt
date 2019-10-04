package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

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

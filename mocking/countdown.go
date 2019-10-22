package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintf(out, "%d\n", i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, "%s", finalWord)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration) // Function that we need to call instead of time.Sleep
}

// Exposed Sleep
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

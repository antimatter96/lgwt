package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
	}
	fmt.Fprintf(out, "%s", finalWord)
}

func main() {
	Countdown(os.Stdout)
}

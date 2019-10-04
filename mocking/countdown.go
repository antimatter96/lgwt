package main

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "%s", "3")
}

package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "%s", "3")
}

func main() {
	Countdown(os.Stdout)
}

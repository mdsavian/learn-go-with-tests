package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

func Countdown(writer io.Writer) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		time.Sleep(1 * time.Second)
	}

	fmt.Fprint(writer, finalWord)
}

func main() {
	Countdown(os.Stdout)
}

package mocking

import (
	"time"
	"os"
	"io"
	"fmt"
)


const finalWord = "Go!"
const countdownStart = 3

//Countdown to one
func Countdown(out io.Writer) {
    for i := countdownStart; i > 0; i-- {
        time.Sleep(1 * time.Second)
        fmt.Fprintln(out, i)
    }

    time.Sleep(1 * time.Second)
    fmt.Fprint(out, finalWord)
}

func main() {
    Countdown(os.Stdout)
}
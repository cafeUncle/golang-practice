package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"
)

const beginNum = 3
const startWord = "go"

func Countdown(out io.Writer, sleeper Sleeper) {
	fmt.Println(startWord)
	time.Sleep(3 * time.Second)
	for i := beginNum; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func main() {
	s := &SpySleeper{}
	Countdown(os.Stdout, s)
	buffer := bytes.Buffer{}
	Countdown(&buffer, s)
	fmt.Println(buffer.String())

	fmt.Println("calls is ", s.Calls)
}

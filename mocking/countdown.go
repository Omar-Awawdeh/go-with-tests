package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

const write = "write"
const sleep = "sleep"

type SpyCountdownOperation struct {
	Calls []string
}

func (s *SpyCountdownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountdownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const countdownStart = 3
const finalWord = "Go!"

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

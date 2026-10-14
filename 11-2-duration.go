package main

import (
	"time"
)

func main() {
	// Duration represents the elapsed time between two instants as an int64 nanosecond count.
	// A Duration type is a signed integer that represents a duration of time in nanoseconds.
	// The zero value of a Duration represents no time elapsed.

	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration1 - duration2

	println("Duration 3:", duration3)
}

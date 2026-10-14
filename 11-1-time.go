package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	// currentTime := time.Now()
	// fmt.Println("Current Time:", currentTime)

	// utc := time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
	// fmt.Println("UTC Time:", utc)
	// fmt.Println("UTC Time in Unix:", utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2023-08-15 14:30:00"
	// Parse a time string
	parsedTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error parsing time:", err.Error())
	}
	fmt.Println("Parsed Time:", parsedTime)

	// Format the current time
	// formattedTime := currentTime.Format(formatter)
	// fmt.Println("Formatted Current Time:", formattedTime)
}

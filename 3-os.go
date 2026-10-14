package main

import (
	"os"
	"fmt"
)

func main() {
	// os.Args is a slice of strings that contains the command-line arguments passed to the program.
	// The first element (os.Args[0]) is the name of the program itself.
	// The subsequent elements (os.Args[1], os.Args[2], ...) are the actual arguments provided by the user.
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error retrieving hostname:", err.Error())
	}
}
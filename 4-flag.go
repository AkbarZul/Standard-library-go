package main

import (
	"fmt"
	"flag"
)

func main() {
	// Define command-line flags
	username := flag.String("username", "defaultUser", "The username for login")
	password := flag.String("password", "defaultPassword", "The password for login")
	host := flag.String("host", "localhost", "The host address")
	port := flag.Int("port", 0, "The port number")

	// Parse the command-line flags
	flag.Parse()
	// Print the values of the flags after parsing
	fmt.Println("Before parsing flags:")
	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *host)
	fmt.Println("Port:", *port)	
}
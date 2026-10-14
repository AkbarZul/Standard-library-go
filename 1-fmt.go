package main

import "fmt"

func main() {
	firstName := "Akbar"
	lastName := "Zulfikar"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
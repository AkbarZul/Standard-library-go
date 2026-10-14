package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es")) // true
	fmt.Println(strings.Split("Akbar Zulfikar", " ")) // [Akbar Zulfikar]
	fmt.Println(strings.ToLower("AKBAR ZULFIKAR")) // akbar zulfikar
	fmt.Println(strings.ToUpper("akbar zulfikar")) // AKBAR ZULFIKAR
	fmt.Println(strings.Trim("   Akbar Zulfikar   ", " ")) // Akbar Zulfikar
	fmt.Println(strings.Replace("Akbar Zulfikar", "Zulfikar", "Zulfi", 1)) // Akbar Zulfi
	fmt.Println(strings.ReplaceAll("Akbar Zulfikar", "Zulfikar", "Zulfi")) // Akbar Zulfi
}
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(re.MatchString("EKO"))
	fmt.Println(re.MatchString("edo"))
	fmt.Println(re.MatchString("eao"))
}
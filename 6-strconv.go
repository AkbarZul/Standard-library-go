package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(boolean)
	}

	int, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(int)
	}
}
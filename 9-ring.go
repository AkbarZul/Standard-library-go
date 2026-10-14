package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	myRing := ring.New(3)
	// myRing.Value = "Akbar"
	// myRing.Next().Value = "Zulfikar"
	// myRing.Next().Next().Value = "Golang"

	// Print the values in the ring
	// for i := 0; i < myRing.Len(); i++ {
	// 	fmt.Println(myRing.Value)
	// 	myRing = myRing.Next()
	// }

	for i := 0; i < myRing.Len(); i++ {
		myRing.Value = "Hello " + strconv.Itoa(i)
		myRing = myRing.Next()
	}

	myRing.Do(func(value any) {
		fmt.Println(value)
	})
}

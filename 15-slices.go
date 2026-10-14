package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Akbar", "Zulfikar", "GO", "John", "Doe", "Python", "Jane", "Smith", "Java"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "GO"))
	fmt.Println(slices.Contains(values, 10))
	fmt.Println(slices.Index(names, "John"))
	fmt.Println(slices.Index(values, 5))
}

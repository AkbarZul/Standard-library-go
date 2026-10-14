package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int           { return len(u) }
func (u UserSlice) Less(i, j int) bool { return u[i].Age < u[j].Age }
func (u UserSlice) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

func main() {
	users := UserSlice{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	// Sort the users by age
	sort.Sort(users)

	// Print the sorted users
	for _, user := range users {
		fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age)
	}
}

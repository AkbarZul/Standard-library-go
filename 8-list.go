package main

import (
	"fmt"
	"container/list"
)

func main() {
	mylist := list.New()
	mylist.PushBack("Akbar")
	mylist.PushBack("Zulfikar")
	mylist.PushBack("Golang")

	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

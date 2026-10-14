package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Base("/home/user/documents/file.txt"))
	fmt.Println(filepath.Dir("hello/world.go"))
}

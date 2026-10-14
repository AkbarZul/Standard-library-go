package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Hello, World!\n")
	// Flush the buffer to ensure the data is written to the underlying writer
	writer.Flush()
}

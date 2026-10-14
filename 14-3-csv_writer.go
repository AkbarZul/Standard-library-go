package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	writer.Write([]string{"Akbar", "Zulfikar", "GO"})
	writer.Write([]string{"John", "Doe", "Python"})
	writer.Write([]string{"Jane", "Smith", "Java"})

	writer.Flush() // Flush the buffered data to the underlying writer
}

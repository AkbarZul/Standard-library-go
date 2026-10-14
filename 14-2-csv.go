package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Akbar,Zulfikar,GO\n" +
		"John,Doe,Python\n" +
		"Jane,Smith,Java"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println("Record:", record)

	}
}

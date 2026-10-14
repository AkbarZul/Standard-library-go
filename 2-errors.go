package main

import "errors"

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetByID(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "123" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := GetByID("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			println("Validation error occurred")
		} else if errors.Is(err, NotFoundError) {
			println("Not found error occurred")
		} else {
			println("An unknown error occurred")
		}
}

}
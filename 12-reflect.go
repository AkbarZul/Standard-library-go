package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type:", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fmt.Println("Field Name:", field.Name)
		fmt.Println("Field Type:", field.Type)

		fieldTag := field.Tag
		fmt.Println("Field Tag Required:", fieldTag.Get("required"))
		fmt.Println("Field Tag Max:", fieldTag.Get("max"))
	}
}

func IsValid(value any) bool {
	valueType := reflect.TypeOf(value)
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fieldTag := field.Tag
		if fieldTag.Get("required") == "true" {
			fieldValue := reflect.ValueOf(value).Field(i).Interface()
			return fieldValue != ""

		}
	}
	return true
}

func main() {
	sample := Sample{Name: "John"}
	person := Person{Name: "Alice"}
	readField(sample)
	readField(person)
	fmt.Println("Is Sample valid?", IsValid(sample))
	fmt.Println("Is Person valid?", IsValid(person))
}

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	b := User{"L", 28}
	getType := reflect.TypeOf(b)
	getValue := reflect.ValueOf(b)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

}

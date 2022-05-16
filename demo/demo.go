package main

import (
	"fmt"
	"github.com/json-iterator/go"
)

func main() {
	m := make(map[string]interface{})
	m["a"] = "a"
	m["b"] = 1

	toString, err := jsoniter.MarshalToString(m)
	if err != nil {

	}
	fmt.Printf("%T, %v\n", toString, toString)
	fmt.Printf("this is a main method!")
}

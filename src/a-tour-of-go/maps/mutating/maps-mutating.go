package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := make(map[string]int)

	m["1"] = 1
	fmt.Println(m["1"])

	m["1"] = 11
	fmt.Println(m["1"])

	delete(m, "1")
	fmt.Println(m["1"]) // not a good way to verify if key is present
	// because m[key] will be default int value = 0

	// we need to use syntax like this
	_, ok := m["1"] // now we retrieve key "status"
	if ok {
		fmt.Println("key `1` exists")
	} else {
		fmt.Println("key `1` not exists")
	}

	fmt.Println(reflect.TypeOf(ok)) // type of ok is `bool``
}
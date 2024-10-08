package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	/*
	Slices are like references to arrays.
	Slices does not store any data - it just describes a section
	declared in bounds of an underlying array.

	So any changes modifies not only underlying array
	but any other slice will see the change of its origin.
	*/
}

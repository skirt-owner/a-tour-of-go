package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
var pows = [][]int{
	{1, 2, 4, 8, 16, 32, 64, 128},
	{1, 3, 9, 27, 81, 243, 729, 2187},
}

func main() {
	for i, value := range pow { // return of index and corresponded value from slice (copy)
		fmt.Printf("2^%d = %d\n", i, value)
	}

	for i, value := range pows {
		fmt.Println(&value == &pows[i])
	} 

	/*
	The output here:
	false
	false

	So we can see that it actually returns a copy of an element
	*/

	for i := range pow {
		fmt.Printf("2^%d = %d\n", i, 2*i)
	} // we can omit value

	for _, value := range pow {
		fmt.Printf("2^%d = %d\n", 2, value)
	} // we can use _ to skip index or value
}
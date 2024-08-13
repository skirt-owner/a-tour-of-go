package main

import "fmt"

func main() {
	s := make([]int, 0)
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	/*
	I don't quite understand source code of aranging capacity :(
	But! Append works like so if there is a space in underlying
	array so it just puts value on the len's position (len <= cap)

	Otherwise it creats new array of capacity comparibly bigger
	based on number of values to be inserted and easier bit
	manipulation (shifting) and returns a slice of this array. 
	*/
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
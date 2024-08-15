package main

import (
	"fmt"
	"strings"
)

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q, len(q), cap(q))

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r, len(r), cap(r))

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s, len(s), cap(s))

	/*
	Slices bounds has defaults values : 0 and len of slice. 
	And i already covered cap and len of slices and arrays.
	*/

	fmt.Println(s[:3])
	fmt.Println(s[3:])
	fmt.Println(s[:])

	// zero value of slice is nil
	var b []int
	fmt.Println(b, len(b), cap(b))

	// we can also create slices with make func : make(type, len, cap)
	// this func returns initialized data structure
	make_slice := make([]int, 0, 5)
	fmt.Println(make_slice, len(make_slice), cap(make_slice))

	/*
	make allocate zeroed array and
	returns a slice reffered to that array
	*/

	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
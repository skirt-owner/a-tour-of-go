package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // array of fixed size
	fmt.Println(len(primes), cap(primes))

	var s []int = primes[1 : 4] // slice with bounds
	fmt.Println(s, len(s), cap(s))
	/*
	primes has len = 6 and cap = 6
	when we take a slice of primes we get the len equals
	to the actual length of slice which is 3 {3, 5, 7}
	but capacity stills the same as for the initial array primes
	but when we put bound like [{n, n > 0} : ] 
	we shrink the avaliable capacity by n so now slice has cap
	of 5 - not 6.
	*/
}
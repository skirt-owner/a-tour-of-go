package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// The receiver type has the literal syntax *T for some type T.
// Methods with pointer receivers can modify the value 
// to which the receiver points (as Scale does here).
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// And as non pointers receiver can be rewritten as reg func
func Scale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	// {3 4}
	// (&v).Scale(100)
	v.Scale(100)
	fmt.Println(v)
	// {300 400}

	Scale(&v, 0.1)
	fmt.Println(v)
	// {30 40}

	v1 := &Vertex{5, 5}
	v1.Scale(10)

	Scale(v1, 10)
}
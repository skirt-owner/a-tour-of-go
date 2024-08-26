package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// (v Vertex) - receiver arg.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// we can define the same as regular func
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// We can declare a method on non-struct types, too.
type MyFloat float64

func (f MyFloat) Abs() MyFloat {
	if f < 0 {
		return -f
	}
	return f
}


func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	// 5
	fmt.Println(Abs(v))
	// 5

	f := MyFloat(math.Sqrt2)
	fmt.Println(f.Abs())
	// 1.4142135623730951
}
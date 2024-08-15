package main

import "github.com/skirt-owner/a-tour-of-go/src/a-tour-of-go/exercises/slices/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	
	for i := range picture {
		picture[i] = make([]uint8, dx)
		for j := range picture[i] {
			picture[i][j] = uint8(i ^ j)
		}
	}
	
	return picture
}

func main() {
	pic.Show(Pic)
}


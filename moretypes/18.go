package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for y := range res {
		res[y] = make([]uint8, dx)
		for x := range res[y] {
			res[y][x] = func1(x, y)
		}
	}

	return res
}

// ---

func func1(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func func2(x, y int) uint8 {
	return uint8(x * y)
}

func func3(x, y int) uint8 {
	return uint8(x ^ y)
}

func printSlice(s [][]uint8) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	for i := range s {
		fmt.Printf("len=%d, cap=%d, %v\n", len(s[i]), cap(s[i]), s[i])
	}
}

func main() {
	printSlice(Pic(3, 4))

	pic.Show(Pic)
}

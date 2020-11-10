package main

import (
	"fmt"
	"math"
)

const loopNum = 5

func Sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}

	z := x / 2
	for i := 0; i < loopNum; i++ {
		z -= (z*z - x) / (2 * z)
		if z == x {
			break
		}
	}
	return z
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Sqrt(float64(i)), math.Sqrt(float64(i)))
	}
}

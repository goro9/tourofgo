package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	// infinite loop
	// return fmt.Sprintf("cannot Sqrt negative number: %v", e)
}

const loopNum = 5

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else if x == 0 {
		return 0, nil
	}

	z := x / 2
	for i := 0; i < loopNum; i++ {
		z -= (z*z - x) / (2 * z)
		if z == x {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

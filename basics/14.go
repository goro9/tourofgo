package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)

	v1 := 3.142
	fmt.Printf("v is of type %T\n", v1)

	v2 := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v2)
}

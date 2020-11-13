package main

import "fmt"

func fibonacci() func() int {
	e0 := 0
	e1 := 1
	return func() int {
		e := e0
		e0 = e1
		e1 = e + e0
		return e
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

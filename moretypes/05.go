package main

import (
	"fmt"
	"unsafe"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
	fmt.Println(unsafe.Sizeof(v1), unsafe.Sizeof(p))
}

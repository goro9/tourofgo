package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { // gofmtでセミコロン消されているくさい
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)
}

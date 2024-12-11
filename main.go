package main

import "fmt"

func main() {
	a := 5
	b := 10
	a, b = b, a
	fmt.Println(a, b)

	a = 5
	b = 10
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)
}

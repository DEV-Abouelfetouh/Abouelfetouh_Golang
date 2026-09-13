package main

import "fmt"

type MathOp func(a, b int) int

func Execute(a, b int, op MathOp) int {
	return op(a, b)
}

func main() {
	add := func(a, b int) int {
		return a + b
	}

	subtract := func(a, b int) int {
		return a - b
	}

	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Println(Execute(10, 5, add))
	fmt.Println(Execute(10, 5, subtract))
	fmt.Println(Execute(10, 5, multiply))
}

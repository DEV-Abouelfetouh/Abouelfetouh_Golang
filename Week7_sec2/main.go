package main

import "fmt"

func fibonacciSlice(n int, result *[]int) {
	*result = make([]int, n)
	if n > 0 {
		(*result)[0] = 0
	}
	if n > 1 {
		(*result)[1] = 1
	}
	for i := 2; i < n; i++ {
		(*result)[i] = (*result)[i-1] + (*result)[i-2]
	}
}

func reverseString(s *string) {
	runes := []rune(*s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*s = string(runes)
}

func main() {
	var fibs []int
	fibonacciSlice(12, &fibs)
	fmt.Println("Fibonacci:", fibs)

	word := "hello"
	reverseString(&word)
	fmt.Println("Reversed:", word)

	x := []int{1, 2, 3}
	y := &x
	x[0] = 100
	fmt.Println("x:", x)
	fmt.Println("y:", *y)
}

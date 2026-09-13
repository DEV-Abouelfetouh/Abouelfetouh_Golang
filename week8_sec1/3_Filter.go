package main

import "fmt"

func Filter(numbers []int, condition func(int) bool) []int {
	var result []int
	for _, v := range numbers {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	isEven := func(n int) bool {
		return n%2 == 0
	}

	evenNumbers := Filter(nums, isEven)
	fmt.Println(evenNumbers)
}

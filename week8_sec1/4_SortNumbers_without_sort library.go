package main

import "fmt"

func SortNumbers(numbers []int, comparator func(a, b int) bool) []int {
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)

	for i := 0; i < len(sorted); i++ {
		for j := 0; j < len(sorted)-1-i; j++ {
			if comparator(sorted[j+1], sorted[j]) {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	return sorted
}

func main() {
	original := []int{5, 2, 8, 1, 9}

	ascending := func(a, b int) bool {
		return a < b
	}

	sortedList := SortNumbers(original, ascending)

	fmt.Println(original)
	fmt.Println(sortedList)
}

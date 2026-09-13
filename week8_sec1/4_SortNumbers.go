package main

import (
	"fmt"
	"sort"
)

func SortNumbers(numbers []int, comparator func(a, b int) bool) []int {
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)

	sort.Slice(sorted, func(i, j int) bool {
		return comparator(sorted[i], sorted[j])
	})

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

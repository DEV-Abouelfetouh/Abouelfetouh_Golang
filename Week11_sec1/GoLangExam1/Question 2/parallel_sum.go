package main

import (
	"fmt"
	"sync"
)

func ParallelSum(nums []int, workers int) int {
	if len(nums) == 0 {
		return 0
	}
	if workers <= 0 {
		workers = 1
	}
	if workers > len(nums) {
		workers = len(nums)
	}

	chunkSize := (len(nums) + workers - 1) / workers
	partialSums := make([]int, workers)

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(nums) {
			end = len(nums)
		}
		if start >= len(nums) {
			break
		}

		wg.Add(1)
		go func(idx, s, e int) {
			defer wg.Done()
			sum := 0
			for _, v := range nums[s:e] {
				sum += v
			}
			partialSums[idx] = sum
		}(i, start, end)
	}

	wg.Wait()

	total := 0
	for _, sum := range partialSums {
		total += sum
	}

	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	workers := 4

	result := ParallelSum(nums, workers)
	fmt.Printf("Total Sum: %d\n", result)
}

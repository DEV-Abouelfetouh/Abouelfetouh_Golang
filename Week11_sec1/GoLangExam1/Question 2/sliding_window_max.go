package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	n := len(nums)

	for i := 0; i <= n-k; i++ {
		maxVal := nums[i]
		for j := i; j < i+k; j++ {
			if nums[j] > maxVal {
				maxVal = nums[j]
			}
		}
		result = append(result, maxVal)
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
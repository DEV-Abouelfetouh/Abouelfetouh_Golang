package main

import (
	"fmt"
	"strconv"
)

func memoizedFactorial() func(int) int {
	cache := make(map[int]int)

	var computeFactorial func(int) int
	computeFactorial = func(n int) int {
		if n <= 1 {
			return 1
		}
		if val, exists := cache[n]; exists {
			return val
		}
		result := n * computeFactorial(n-1)
		cache[n] = result
		return result
	}

	// Wrapper function that checks cache before computing
	return func(n int) int {
		if _, exists := cache[n]; exists {
			fmt.Println("[NOTIFICATION] Fetching result from cache...")
		}
		return computeFactorial(n)
	}
}

func runFactorialTask() {
	fact := memoizedFactorial()
	var input string

	fmt.Println("[NOTIFICATION] Entering Factorial Module.")

	for {
		fmt.Print("Enter a positive integer (or 'x' to return to main menu): ")
		fmt.Scan(&input)

		if input == "x" || input == "X" {
			fmt.Println("[NOTIFICATION] Exiting Factorial Module...")
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil || num < 0 {
			fmt.Println("[ERROR] Invalid input! Only non-negative integers are allowed.")
			continue
		}

		fmt.Printf("[RESULT] %d! = %d\n\n", num, fact(num))
	}
}

package main

import (
	"fmt"
	"strconv"
)

func average(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	var total float64
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

func runAverageTask() {
	var numbers []float64
	var input string

	fmt.Println("[NOTIFICATION] Entering Average Calculator Module.")

	for {
		fmt.Print("Enter a number (or 'x' to calculate average and return to main menu): ")
		fmt.Scan(&input)

		if input == "x" || input == "X" {
			if len(numbers) == 0 {
				fmt.Println("[NOTIFICATION] No numbers were entered.")
			} else {
				avg := average(numbers...)
				fmt.Printf("[FINAL RESULT] Average of %v is: %.2f\n", numbers, avg)
			}
			fmt.Println("[NOTIFICATION] Exiting Average Calculator Module...")
			break
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("[ERROR] Invalid input! Only numbers are allowed.")
			continue
		}

		numbers = append(numbers, num)
		fmt.Printf("[NOTIFICATION] Added %.2f to list. Current numbers count: %d\n\n", num, len(numbers))
	}
}

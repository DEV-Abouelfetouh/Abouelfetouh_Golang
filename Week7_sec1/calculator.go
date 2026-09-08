package main

import (
	"errors"
	"fmt"
	"strconv"
)

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+", "sum":
		return a + b, nil
	case "-", "diff":
		return a - b, nil
	case "*", "mul":
		return a * b, nil
	case "/", "div":
		if b == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("unsupported operation")
	}
}

func runCalculatorTask() {
	var input string
	fmt.Println("[NOTIFICATION] Entering Continuous Calculator Module.")

	for {
		fmt.Print("Enter initial starting number (or 'x' to return to main menu): ")
		fmt.Scan(&input)

		if input == "x" || input == "X" {
			fmt.Println("[NOTIFICATION] Exiting Calculator Module...")
			break
		}

		currentResult, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("[ERROR] Invalid input! Only numbers are allowed.")
			continue
		}

		for {
			fmt.Printf("Current Result: %.2f\n", currentResult)
			fmt.Print("Enter operation (+, -, *, /) or 'x' to finish calculation: ")
			var op string
			fmt.Scan(&op)

			if op == "x" || op == "X" {
				fmt.Printf("[FINAL RESULT] %.2f\n\n", currentResult)
				break
			}

			fmt.Print("Enter next number: ")
			fmt.Scan(&input)

			if input == "x" || input == "X" {
				fmt.Printf("[FINAL RESULT] %.2f\n\n", currentResult)
				break
			}

			nextNum, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("[ERROR] Invalid number! Skipping step.")
				continue
			}

			newResult, calcErr := calculate(currentResult, nextNum, op)
			if calcErr != nil {
				fmt.Printf("[ERROR] %v\n", calcErr)
				continue
			}

			currentResult = newResult
		}
	}
}

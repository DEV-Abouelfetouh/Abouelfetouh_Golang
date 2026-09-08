package main

import "fmt"

func main() {
	for {
		fmt.Println("\n====================================")
		fmt.Println("       TASK SELECTION MENU          ")
		fmt.Println("====================================")
		fmt.Println("1. Memoized Factorial (Closure + Map)")
		fmt.Println("2. Palindrome Checker")
		fmt.Println("3. Anagram Checker")
		fmt.Println("4. Average Calculator (Variadic Function)")
		fmt.Println("5. Calculator (Error Handling)")
		fmt.Println("0. Exit Application")
		fmt.Println("------------------------------------")

		var choice int
		fmt.Print("Select a task number: ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("[ERROR] Invalid input! Please enter a valid menu number.")
			var dump string
			fmt.Scan(&dump)
			continue
		}

		fmt.Println()
		switch choice {
		case 1:
			runFactorialTask()
		case 2:
			runPalindromeTask()
		case 3:
			runAnagramTask()
		case 4:
			runAverageTask()
		case 5:
			runCalculatorTask()
		case 0:
			fmt.Println("[NOTIFICATION] Exiting application. Goodbye!")
			return
		default:
			fmt.Println("[ERROR] Invalid choice. Please select from the menu.")
		}
	}
}

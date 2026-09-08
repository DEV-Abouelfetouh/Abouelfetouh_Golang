package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	// Convert to lowercase to ignore case sensitivity
	cleanStr := strings.ToLower(s)
	runes := []rune(cleanStr)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func runPalindromeTask() {
	var input string
	fmt.Println("[NOTIFICATION] Entering Palindrome Checker Module.")

	for {
		fmt.Print("Enter a string to check (or 'x' to return to main menu): ")
		fmt.Scan(&input)

		if input == "x" || input == "X" {
			fmt.Println("[NOTIFICATION] Exiting Palindrome Checker Module...")
			break
		}

		if isPalindrome(input) {
			fmt.Printf("[RESULT] '%s' is a palindrome!\n\n", input)
		} else {
			fmt.Printf("[RESULT] '%s' is NOT a palindrome.\n\n", input)
		}
	}
}

package main

import "fmt"

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	counts := make(map[rune]int)
	for _, char := range s1 {
		counts[char]++
	}
	for _, char := range s2 {
		counts[char]--
		if counts[char] < 0 {
			return false
		}
	}
	return true
}

func runAnagramTask() {
	var word1, word2 string
	fmt.Println("[NOTIFICATION] Entering Anagram Checker Module.")

	for {
		fmt.Print("Enter first string (or 'x' to return to main menu): ")
		fmt.Scan(&word1)

		if word1 == "x" || word1 == "X" {
			fmt.Println("[NOTIFICATION] Exiting Anagram Checker Module...")
			break
		}

		fmt.Print("Enter second string: ")
		fmt.Scan(&word2)

		if isAnagram(word1, word2) {
			fmt.Printf("[RESULT] '%s' and '%s' are anagrams!\n\n", word1, word2)
		} else {
			fmt.Printf("[RESULT] '%s' and '%s' are NOT anagrams.\n\n", word1, word2)
		}
	}
}

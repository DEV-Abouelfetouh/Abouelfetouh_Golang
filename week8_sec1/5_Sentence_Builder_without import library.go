package main

import "fmt"

func BuildSentence(words []string) string {
	var result string
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}
	return result
}

func main() {
	words := []string{"Go", "is", "an", "awesome", "language"}
	result := BuildSentence(words)
	fmt.Println(result)
}

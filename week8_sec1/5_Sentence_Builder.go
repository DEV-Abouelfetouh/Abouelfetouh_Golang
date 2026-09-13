package main

import (
	"fmt"
	"strings"
)

func BuildSentence(words []string) string {
	return strings.Join(words, " ")
}

func main() {
	words := []string{"Go", "is", "an", "awesome", "language"}
	result := BuildSentence(words)
	fmt.Println(result)
}

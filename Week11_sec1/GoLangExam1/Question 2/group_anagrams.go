package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range strs {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		key := string(runes)

		groups[key] = append(groups[key], word)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input))
}

package main

import "fmt"

func main() {
	s := "abcabcbb"

	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	lastIndexMap := make(map[rune]int)

	start := 0
	maxLength := 0
	currentLength := 0
	for end, c := range s {
		lastIndex, ok := lastIndexMap[c]
		if !ok {
			lastIndex = -1
		}
		start = max(start, lastIndex+1)

		currentLength = end - start + 1
		maxLength = max(currentLength, maxLength)

		lastIndexMap[c] = end
	}

	return maxLength
}

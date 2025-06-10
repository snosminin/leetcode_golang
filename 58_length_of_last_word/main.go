package main

import "fmt"

func main() {
	var s = "day"

	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	lastWordCount := 0

	for i := len(s) - 1; i >= 0; i-- {
		if rune(s[i]) != rune(' ') {
			lastWordCount = 0
			for j := i; j >= 0; j-- {
				if rune(s[j]) != rune(' ') {
					lastWordCount++
				} else {
					return lastWordCount
				}
				i = j
			}
		}
	}
	return lastWordCount
}
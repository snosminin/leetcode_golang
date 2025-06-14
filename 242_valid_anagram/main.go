package main

import "fmt"

func main() {
	s, t := "anagram", "nagaram"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	lettersCount := [26]int{}

	for i, sChar := range s {
		tChar := rune(t[i])
		lettersCount[sChar-rune('a')]++
		lettersCount[tChar-rune('a')]--
	}

	for _, count := range lettersCount {
		if count != 0 {
			return false
		}
	}

	return true
}
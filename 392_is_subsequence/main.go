package main

import "fmt"

func main() {
	s, t := "axc", "ahbgdc"

	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	sIndex, tIndex := 0, 0

	for sIndex < len(s) && tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		tIndex++
	}

	return sIndex == len(s)
}

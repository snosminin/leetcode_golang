package main

import "fmt"

func main() {
	s := "the sky is blue"

	fmt.Println("-", reverseWords(s), "-")
}

func reverseWords(s string) string {
	stack := [][]rune{}
	i := 0
	for i < len(s) {
		if rune(s[i]) != rune(' ') {
			word := []rune{}
			for j := i; j < len(s); j++ {
				if rune(s[j]) != rune(' ') {
					word = append(word, rune(s[j]))
				} else {
					i = j
					break
				}
				i = j
			}
			stack = append(stack, word)
		}
		i++
	}

	result := []rune{}
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1]...)
		result = append(result, rune(' '))
		stack = stack[:len(stack)-1]
	}

	return string(result[:len(result)-1])
}

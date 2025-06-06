package main

import "fmt"

type TreeNode struct {
		Val int
		Left *TreeNode
		Right *TreeNode
	}

func main() {
	s := "abb"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	if (len(s) == 0) {
		return false
	}

	converted := []rune {}
    for _, c := range s {
		if (c >= rune('a') && c <= rune('z')) {
			converted = append(converted, c)
		} else if(c >= rune('0') && c <= rune('9')) {
			converted = append(converted, c)
		} else if (c >= rune('A') && c <= rune('Z')) {
			converted = append(converted, c + 32)
		} else {
			continue
		}
	}

	if (len(converted) == 1) {
		return true
	}

	halfLength := len(converted) / 2
	isOdd := len(converted) % 2 == 1
	stack := make([]rune, halfLength) 

	i := 0
	for i < halfLength {
		stack = append(stack, converted[i])
		i++
	}

	if (isOdd) {
		i++
	}

	for i < len(converted) {
		if (stack[len(stack) - 1] != converted[i]) {
			return false
		}

		stack = stack[:len(stack) - 1]
		i++
	}
	
	return true
}
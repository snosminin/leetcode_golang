package main

import "fmt"

func main() {
	s := "{(])}"

	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	stack := []rune{}

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if (len(stack) > 0) && (c == ')' && stack[len(stack)-1] == '(') {
			stack = stack[:len(stack)-1]
		} else if (len(stack) > 0) && (c == '}' && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else if (len(stack) > 0) && (c == ']' && stack[len(stack)-1] == '[') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

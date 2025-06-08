package main

import (
	"fmt"
)

func main() {
	const number = 252

	fmt.Println(isPalindrome(number))
}

func isPalindrome(x int) bool {
	reversedNumber := 0
	copyNumber := x
	for copyNumber > 0 {
		reversedNumber = reversedNumber * 10 + copyNumber % 10 
		copyNumber /= 10
	}

	return reversedNumber == x
}
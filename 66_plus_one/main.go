package main

import "fmt"

func main() {
	var array = [...]int{9, 9, 9, 9, 9, 9, 9, 9}

	fmt.Println(plusOne(array[:]))
}

func plusOne(digits []int) []int {
	overflow := false
	digits[len(digits) - 1]++
    for i := len(digits) - 1; i >= 0; i-- {
		if(overflow){
			digits[i]++
		}

		if(digits[i] > 9){
			digits[i] -= 10 
			overflow = true
		} else {
			overflow = false
			break
		}
	}
	if(overflow) {
		digits = append([]int {1}, digits... )
	}
	return digits
}
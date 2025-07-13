package main

import "fmt"

func main() {
	var array = [...]int{2, 7, 2, 15, 34, 2, 1, 6, 2, 4, 2, 2, 2, 7, 2}

	fmt.Println(majorityElement(array[:]))
}

func majorityElement(nums []int) int {
	candidate := -1
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

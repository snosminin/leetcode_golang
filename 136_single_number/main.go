package main

import "fmt"

func main() {
	var array = [...]int{2, 2, 1}

	fmt.Println(singleNumber(array[:]))
}

func singleNumber(nums []int) int {
	var result = 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

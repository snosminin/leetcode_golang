package main

import "fmt"

func main() {
	var array = []int{2, 3, 1, 1, 4}

	fmt.Println(jump(array))
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	result, farthest, end := 0, 0, 0

	for i := range nums {
		if farthest < i+nums[i] {
			farthest = i + nums[i]
		}

		if farthest >= len(nums)-1 {
			result++
			break
		}

		if i == end {
			result++
			end = farthest
		}
	}

	return result
}

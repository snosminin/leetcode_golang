package main

import "fmt"

func main() {
	var array = []int{2,3,1,1,4}

	fmt.Println(canJump(array[:]))
}

func canJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		if maxReach < i {
			return false
		}

		if maxReach < i+nums[i] {
			maxReach = i + nums[i]
		}
	}

	return true
}
package main

import "fmt"

func main() {
	var array = []int{1, 2}

	fmt.Println(removeDuplicates(array))
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	index := 2
	for i := 2; i < len(nums); i++ {
		if nums[index-1] == nums[i] && nums[index-2] != nums[i] {
			nums[index] = nums[i]
			index++
		} else if nums[index-1] < nums[i] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

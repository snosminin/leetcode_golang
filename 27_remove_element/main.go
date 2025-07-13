package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums, val := []int{4, 5}, 4
	result := removeElement(nums, val)
	fmt.Println(nums, result)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	lastIndex := len(nums) - 1
	currentIndex := lastIndex
	result := 0

	for currentIndex >= 0 {
		if nums[currentIndex] == val {
			result++
			i := currentIndex
			for i < lastIndex && nums[i+1] != val {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				i++
			}
		}
		currentIndex--
	}
	return len(nums) - result
}

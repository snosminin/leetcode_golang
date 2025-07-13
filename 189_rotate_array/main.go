package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums, val := []int{1, 2, 3, 4, 5, 6, 7}, 3
	rotate(nums, val)
	fmt.Println(nums, val)
}

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	k = k % len(nums)
	result := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i, _ := range nums {
		nums[i] = result[i]
	}
}

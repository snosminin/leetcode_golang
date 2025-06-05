package main

import "fmt"

type TreeNode struct {
		Val int
		Left *TreeNode
		Right *TreeNode
	}

func main() {
	nums1, m, nums2, n := []int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j := 0, 0

	for i < m || j < n {
		
	}
}